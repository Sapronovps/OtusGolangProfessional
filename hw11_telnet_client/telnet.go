package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	Address string
	Timeout time.Duration
	In      io.ReadCloser
	Out     io.Writer
	Conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	telnet := &Telnet{
		Address: address,
		Timeout: timeout,
		In:      in,
		Out:     out,
		Conn:    nil,
	}
	return telnet
}

func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.Address, t.Timeout)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	t.Conn = conn
	log.Println("connected to ", t.Address)

	return nil
}

func (t *Telnet) Send() error {
	_, err := io.Copy(t.Conn, t.In)
	if err != nil {
		return fmt.Errorf("send message error: %w", err)
	}
	return nil
}

func (t *Telnet) Receive() error {
	_, err := io.Copy(t.Out, t.Conn)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("receive message error: %w", err)
	}
	return nil
}

func (t *Telnet) Close() error {
	if t.Conn != nil {
		if err := t.Conn.Close(); err != nil {
			return fmt.Errorf("close connection error: %w", err)
		}
	}
	return nil
}
