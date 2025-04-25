package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "how long to wait")
	flag.Parse()

	host := flag.Arg(0)
	port := flag.Arg(1)

	in := &bytes.Buffer{}

	telnetClient := NewTelnetClient(net.JoinHostPort(host, port), timeout, io.NopCloser(in), os.Stdout)
	err := telnetClient.Connect()
	if err != nil {
		fmt.Printf("connect telnet client error: %v\n", err)
	}
	defer telnetClient.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go sendToConnection(ctx, cancel, &wg, telnetClient, in)

	go readFromConnection(ctx, &wg, telnetClient)

	wg.Wait()
}

func sendToConnection(ctx context.Context, cancel func(), wg *sync.WaitGroup, t TelnetClient, in *bytes.Buffer) {
	go func() {
		<-ctx.Done()
		t.Close()
		wg.Done()
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		resp, err := reader.ReadString('\n')
		if err != nil {
			cancel()
			return
		}
		in.WriteString(resp)
		if err := t.Send(); err != nil {
			cancel()
			return
		}
	}
}

func readFromConnection(ctx context.Context, wg *sync.WaitGroup, t TelnetClient) {
	defer wg.Done()
	errChannel := make(chan error)

	for {
		select {
		case <-ctx.Done():
			return
		case errChannel <- t.Receive():
			err := <-errChannel
			if err != nil {
				return
			}
		}
	}
}
