package main

import (
	"bufio"
	"net"
)

func drpecho(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			conn.Write([]byte("\n"))
			return
		}
		conn.Write([]byte(line + "\n"))
	}
}

func handlerror(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:2333")
	handlerror(err)
	for {
		conn, err := ln.Accept()
		handlerror(err)
		go drpecho(conn)
	}
}
