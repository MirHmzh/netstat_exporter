package main

import (
	"bytes"
	"flag"
	"strconv"
	"netstat_exporter/httphandler"
	log "github.com/jeanphorn/log4go"
)

var strbuff bytes.Buffer

func main() {
	listenPort := flag.Int("port", 9119, "Listen port")
	flag.Parse()
	log.Info("Netstat Exporter by mirhmzh running on port :"+strconv.Itoa(*listenPort))
	httphandler.Init(*listenPort)
}