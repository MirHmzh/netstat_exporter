package main

import (
	"bytes"
	"flag"
	"fmt"
	"netstat_exporter/httphandler"
)

var strbuff bytes.Buffer

func main() {
	listenPort := flag.Int("port", 9119, "Listen port")
	flag.Parse()
	fmt.Println(*listenPort)
	httphandler.Init(*listenPort)
}