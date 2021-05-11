package main

import (
	"fmt"
	"net"
)

// 1. Client sends syn request
// 2. Server responds with syn-ack if open, or with rst if closed
// 3. Client sends ack request and TCP connection is established
// If firewall is in the middle filtering ports, connection timeout should occure.

func worker(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("localhost:%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var OpenPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 5000; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 5000; i++ {
		port := <-results
		if port != 0 {
			OpenPorts = append(OpenPorts, port)
		}
	}

	close(ports)
	close(results)

	for _, p := range OpenPorts {
		fmt.Printf("%d open\n", p)
	}

}
