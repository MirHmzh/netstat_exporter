package httphandler

import (
	"fmt"
	"bytes"
	"strconv"
	"net/http"
	"netstat_exporter/tcpget"
	"netstat_exporter/udpget"
	"netstat_exporter/tcp6get"
	"netstat_exporter/udp6get"
)

var strbuff bytes.Buffer

func Init(port int) {
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func metricsHandler(w http.ResponseWriter, req *http.Request) {
	strbuff.Reset()
	strbuff.WriteString(tcpget.Init())
	strbuff.WriteString(udpget.Init())
	strbuff.WriteString(tcp6get.Init())
	strbuff.WriteString(udp6get.Init())
	fmt.Fprintf(w, strbuff.String())
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `
		<html>
			<head>
				<title>Netstat Exporter</title>
			</head>
			<body>
				<h1>Netstat Exporter</h1>
				<p><a href="/metrics">Metrics</a></p>
			</body>
		</html>
		`)
}