package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func connectToServiceWithPool() interface{} {
	time.Sleep(1 * time.Second)

	return struct{}{}
}

func warmServiceCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToServiceWithPool,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func startNetworkDaemonWithPool() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceCache()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("can not listen %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v ", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()

	return &wg
}

func init() {
	daemonStarted := startNetworkDaemonWithPool()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequestWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

// go test -benchtime=10s -bench=. benchmark_with_pool_test.go
