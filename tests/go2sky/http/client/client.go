//
// Copyright 2021 SkyAPM org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	httpPlugin "github.com/SkyAPM/go2sky/plugins/http"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const (
	oap         = "127.0.0.1:19876"
	service     = "http-client"
	upstreamURL = "http://127.0.0.1:8080/hello"
)

func main() {
	report, err := reporter.NewGRPCReporter(oap)
	if err != nil {
		log.Fatalf("crate grpc reporter error: %v \n", err)
	}

	tracer, err := go2sky.NewTracer(service, go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	}

	client, err := NewClient(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}

	route := http.NewServeMux()
	route.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		clientReq, err1 := http.NewRequest(http.MethodGet, upstreamURL, nil)
		if err1 != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to create http request error: %v \n", err)
			return
		}
		clientReq = clientReq.WithContext(request.Context())
		res, err1 := client.Do(clientReq)
		if err1 != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to do http request error: %v \n", err)
			return
		}
		defer res.Body.Close()
		body, err1 := ioutil.ReadAll(res.Body)
		if err1 != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Printf("read http response error: %v \n", err)
			return
		}
		writer.WriteHeader(res.StatusCode)
		_, _ = writer.Write(body)
	})

	route.HandleFunc("/healthCheck", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Success"))
	})

	sm, err := httpPlugin.NewServerMiddleware(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}
	err = http.ListenAndServe(":8081", sm(route))
	if err != nil {
		log.Fatal(err)
	}
}

const componentIDGOHttpClient = 5005

type ClientConfig struct {
	name      string
	client    *http.Client
	tracer    *go2sky.Tracer
	extraTags map[string]string
}

// ClientOption allows optional configuration of Client.
type ClientOption func(*ClientConfig)

// WithOperationName override default operation name.
func WithClientOperationName(name string) ClientOption {
	return func(c *ClientConfig) {
		c.name = name
	}
}

// WithClientTag adds extra tag to client spans.
func WithClientTag(key string, value string) ClientOption {
	return func(c *ClientConfig) {
		if c.extraTags == nil {
			c.extraTags = make(map[string]string)
		}
		c.extraTags[key] = value
	}
}

// WithClient set customer http client.
func WithClient(client *http.Client) ClientOption {
	return func(c *ClientConfig) {
		c.client = client
	}
}

// NewClient returns an HTTP Client with tracer
func NewClient(tracer *go2sky.Tracer, options ...ClientOption) (*http.Client, error) {
	if tracer == nil {
		return nil, nil
	}
	co := &ClientConfig{tracer: tracer}
	for _, option := range options {
		option(co)
	}
	if co.client == nil {
		co.client = &http.Client{}
	}
	tp := &transport{
		ClientConfig: co,
		delegated:    http.DefaultTransport,
	}
	if co.client.Transport != nil {
		tp.delegated = co.client.Transport
	}
	co.client.Transport = tp
	return co.client, nil
}

type transport struct {
	*ClientConfig
	delegated http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	fmt.Println("!!!!!!", getOperationName(t.name, req), req.Host)
	span, err := t.tracer.CreateExitSpan(req.Context(), getOperationName(t.name, req), req.Host+"www", func(key, value string) error {
		fmt.Println("set!!", key, value)
		req.Header.Set(key, value)
		return nil
	})
	if err != nil {
		return t.delegated.RoundTrip(req)
	}
	defer span.End()
	span.SetComponent(componentIDGOHttpClient)
	for k, v := range t.extraTags {
		span.Tag(go2sky.Tag(k), v)
	}
	span.Tag(go2sky.TagHTTPMethod, req.Method)
	span.Tag(go2sky.TagURL, req.URL.String())
	span.SetSpanLayer(agentv3.SpanLayer_Http)
	res, err = t.delegated.RoundTrip(req)
	if err != nil {
		span.Error(time.Now(), err.Error())
		return
	}
	span.Tag(go2sky.TagStatusCode, strconv.Itoa(res.StatusCode))
	if res.StatusCode >= http.StatusBadRequest {
		span.Error(time.Now(), "Errors on handling client")
	}
	return res, nil
}

const componentIDGOHttpServer = 5004

type handler struct {
	tracer    *go2sky.Tracer
	name      string
	next      http.Handler
	extraTags map[string]string
}

// ServerOption allows Middleware to be optionally configured.
type ServerOption func(*handler)

// Tag adds extra tag to server spans.
func WithServerTag(key string, value string) ServerOption {
	return func(h *handler) {
		if h.extraTags == nil {
			h.extraTags = make(map[string]string)
		}
		h.extraTags[key] = value
	}
}

// WithOperationName override default operation name.
func WithServerOperationName(name string) ServerOption {
	return func(h *handler) {
		h.name = name
	}
}

// NewServerMiddleware returns a http.Handler middleware with tracing.
func NewServerMiddleware(tracer *go2sky.Tracer, options ...ServerOption) (func(http.Handler) http.Handler, error) {
	if tracer == nil {
		return nil, nil
	}
	return func(next http.Handler) http.Handler {
		h := &handler{
			tracer: tracer,
			next:   next,
		}
		for _, option := range options {
			option(h)
		}
		return h
	}, nil
}

// ServeHTTP implements http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	span, ctx, err := h.tracer.CreateEntrySpan(r.Context(), getOperationName(h.name, r), func(key string) (string, error) {
		fmt.Println("get!!", key)
		return r.Header.Get(key), nil
	})
	if err != nil {
		if h.next != nil {
			h.next.ServeHTTP(w, r)
		}
		return
	}
	span.SetComponent(componentIDGOHttpServer)
	for k, v := range h.extraTags {
		span.Tag(go2sky.Tag(k), v)
	}
	span.Tag(go2sky.TagHTTPMethod, r.Method)
	span.Tag(go2sky.TagURL, fmt.Sprintf("%s%s", r.Host, r.URL.Path))
	span.SetSpanLayer(agentv3.SpanLayer_Http)

	rww := &responseWriterWrapper{w: w, statusCode: 200}
	defer func() {
		code := rww.statusCode
		if code >= 400 {
			span.Error(time.Now(), "Error on handling request")
		}
		span.Tag(go2sky.TagStatusCode, strconv.Itoa(code))
		span.End()
	}()
	if h.next != nil {
		h.next.ServeHTTP(rww, r.WithContext(ctx))
	}
}

type responseWriterWrapper struct {
	w          http.ResponseWriter
	statusCode int
}

func (rww *responseWriterWrapper) Header() http.Header {
	return rww.w.Header()
}

func (rww *responseWriterWrapper) Write(bytes []byte) (int, error) {
	return rww.w.Write(bytes)
}

func (rww *responseWriterWrapper) WriteHeader(statusCode int) {
	rww.statusCode = statusCode
	rww.w.WriteHeader(statusCode)
}

func getOperationName(name string, r *http.Request) string {
	if name == "" {
		return fmt.Sprintf("/%s%s", r.Method, r.URL.Path)
	}
	return name
}

