package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type CronJob struct {
    Id       string `json:"id"`
    BumpTime string `json:"bumpTime"`
    State    string `json:"state"`
}

type PostJob struct {
    Id      string `json:"id"`
    JobId   string `json:"jobId"`
    State   string `json:"state"`
    Url     string `json:"url"`
    Content string `json:"content"`
}

var jobs = []*CronJob{
    {
        Id:       "updateExpiredData",
            BumpTime: "0:0",
        State:    "active",
    },
    {
        Id:       "updateHotInfo",
            BumpTime: "*/1:*",
        State:    "active",
    },
}

var posts = []*PostJob{
    {
        Id:    "postUpdateExpiredData",
            JobId: "updateExpiredData",
        State: "sent",
        Url:   "http://127.0.0.1:7000/api/update-expired-data",
    },
    {
        Id:    "postUpdateHotInfo",
            JobId: "updateHotInfo",
        State: "sent",
        Url:   "http://127.0.0.1:7000/api/update-hot-info",
    },
}

func main() {
var buf []byte
var err error

if buf, err = json.Marshal(jobs); err != nil {
log.Fatal("json marshal error:", err)
}

fmt.Println(string(buf))

var msg []*CronJob
err = json.Unmarshal([]byte(string(buf)), &msg)
for _, v := range msg {
    fmt.Println(v)
}
}
