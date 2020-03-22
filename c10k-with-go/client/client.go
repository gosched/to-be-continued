package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var connCount int64

func connect() {
	defer func() {
		atomic.AddInt64(&connCount, -1)
	}()

	atomic.AddInt64(&connCount, 1)

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}

	text := "hello"

	for {
		conn.Write([]byte(text))
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		_ = message
	}
}

func main() {
	fmt.Println("process id", os.Getpid())
	timer := time.NewTimer(time.Second * 60 * 60 * 3)
	tick1 := time.NewTicker(time.Second * 1)
	tick2 := time.NewTicker(time.Second * 3)
loop:
	for {
		select {
		case <-tick1.C:
			go connect()
		case <-tick2.C:
			fmt.Println("number of goroutines", runtime.NumGoroutine())
			fmt.Println("number of connect", atomic.LoadInt64(&connCount))
		case <-timer.C:
			tick1.Stop()
			tick2.Stop()
			timer.Stop()
			break loop
		}
	}
}

// runtime.GOMAXPROCS()
// runtime.NumGoroutine()
