package p2p

import (
    "net"
    "log"
    "fmt"
    "flag"
)

// flag.String   f func(name string, value string, usage string) *string
var server = flag.String("-s", "http://127.0.0.1:5100", "server address")
var local  = flag.String("-c", "http://127.0.0.1:5200", "client address")
var dir    = flag.String("-d", "", "")

func main() {
    flag.Parse();
    fmt.Println(server, local, dir)

    addr, err := net.ResolveTCPAddr("tcp", *local)
    if err != nil { log.Panic(err) }
    fmt.Println(addr)
}
