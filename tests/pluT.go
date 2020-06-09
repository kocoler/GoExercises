package main

import "plugin"

func main() {
	p, err := plugin.Open("")
	sb, err := p.Lookup("")
	sbs := sb.(func(string) string)
}
