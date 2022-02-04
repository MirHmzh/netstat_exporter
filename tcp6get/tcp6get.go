package	tcp6get

import (
	"bytes"
	"strconv"
	"github.com/cakturk/go-netstat/netstat"
)

func Init() string{
	var strbuff bytes.Buffer
	tcps, err := netstat.TCP6Socks(netstat.NoopFilter)
	if err != nil {
		panic(err)
	}
	for _, e := range tcps {
		if err != nil {
			panic(err)
		}
		if e.State == netstat.Listen {
			strbuff.WriteString("node_port_used{port="+(strconv.Itoa(int(e.LocalAddr.Port)))+",process=\""+e.Process.Name+"\",type=\"TCP\",pid="+(strconv.Itoa(e.Process.Pid))+"} 1.0\n")
		}
	}
	return strbuff.String()
}