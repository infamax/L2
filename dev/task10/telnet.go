package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "waiting time")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("expected two arguments host and port")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	if *timeout <= 0 {
		log.Fatal("incorrect timeout, expected positive timeout")
	}

	tc := New(host, port, *timeout)
	if err := tc.Open(); err != nil {
		log.Fatal(err)
	}
	defer tc.Close()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	errorChan := make(chan error, 1)
	go func() {
		if _, err := io.Copy(os.Stdout, tc); err != nil {
			errorChan <- err
			return
		}
	}()

	go func() {
		if _, err := io.Copy(tc, os.Stdin); err != nil {
			errorChan <- err
			return
		}
	}()

	select {
	case err := <-errorChan:
		log.Fatal(err)
	case <-signalChan:
		fmt.Println("Interrupted!")
	}
}

type telnetClient struct {
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

func New(host, port string, timeout time.Duration) *telnetClient {
	return &telnetClient{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

func (tc *telnetClient) Open() error {
	if tc.conn != nil {
		return errors.New("tcp connection already opened")
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(tc.host, tc.port), tc.timeout)
	if err != nil {
		return err
	}
	tc.conn = conn
	return nil
}

func (tc *telnetClient) Close() error {
	if tc.conn != nil {
		return tc.conn.Close()
	}
	return nil
}

func (tc *telnetClient) Read(b []byte) (int, error) {
	return tc.conn.Read(b)
}

func (tc *telnetClient) Write(b []byte) (int, error) {
	return tc.conn.Write(b)
}
