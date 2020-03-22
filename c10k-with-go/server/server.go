package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync/atomic"
)

var connCount int64
var messageCount uint64

func handle(conn net.Conn) {
	atomic.AddInt64(&connCount, 1)
	fmt.Printf("connCount %d\n", connCount)
	defer func() {
		atomic.AddInt64(&connCount, -1)
		fmt.Printf("connCount %d\n", connCount)
	}()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		atomic.AddUint64(&messageCount, 1)
		conn.Write([]byte(strings.ToUpper(message) + "\n"))
	}
}

func main() {
	fmt.Println("process id", os.Getpid())
	ln, _ := net.Listen("tcp", ":8080")
	for {
		if conn, err := ln.Accept(); err == nil {
			if 5000 < atomic.LoadInt64(&connCount) {
				ln.Close()
			}
			go handle(conn)
		} else {
			fmt.Println(err)
		}
	}
}

// runtime.GOMAXPROCS()
// runtime.NumGoroutine()

/*
sudo vi /etc/security/limits.conf
gosched soft nofile 10000
gosched hard nofile 10000
*/

// ulimit -Sn
// ulimit -n 10000
// go run server.go
// lsof -p pID | wc -l
// sudo lsof -i:8080
