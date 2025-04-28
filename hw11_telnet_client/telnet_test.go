package main

import (
	"bytes"
	"io"
	"net"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTelnetClient(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")
		require.NoError(t, err)
		defer func() { require.NoError(t, l.Close()) }()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			in := &bytes.Buffer{}
			out := &bytes.Buffer{}

			timeout, err := time.ParseDuration("10s")
			require.NoError(t, err)

			client := NewTelnetClient(l.Addr().String(), timeout, io.NopCloser(in), out)
			require.NoError(t, client.Connect())
			defer func() { require.NoError(t, client.Close()) }()

			in.WriteString("hello\n")
			err = client.Send()
			require.NoError(t, err)

			err = client.Receive()
			require.NoError(t, err)
			require.Equal(t, "world\n", out.String())
		}()

		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			require.NoError(t, err)
			require.NotNil(t, conn)
			defer func() { require.NoError(t, conn.Close()) }()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			require.NoError(t, err)
			require.Equal(t, "hello\n", string(request)[:n])

			n, err = conn.Write([]byte("world\n"))
			require.NoError(t, err)
			require.NotEqual(t, 0, n)
		}()

		wg.Wait()
	})
}

func TestTelnet_Connect(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:")
	require.NoError(t, err)
	defer func() { require.NoError(t, l.Close()) }()

	tests := []struct {
		name        string
		address     string
		timeout     time.Duration
		wantErr     bool
		expectedErr string
	}{
		{
			name:    "successful connection",
			address: l.Addr().String(),
			timeout: time.Second,
			wantErr: false,
		},
		{
			name:        "invalid address",
			address:     "invalid_address",
			timeout:     time.Second,
			wantErr:     true,
			expectedErr: "connection error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewTelnetClient(tt.address, tt.timeout, nil, nil)
			err := client.Connect()

			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && !strings.Contains(err.Error(), tt.expectedErr) {
				t.Errorf("Connect() error = %v, expected to contain %v", err, tt.expectedErr)
			}
		})
	}
}

// MockConn - мок реализации net.Conn для тестов.
type MockConn struct {
	ReadBuf  *bytes.Buffer
	WriteBuf *bytes.Buffer
	Closed   bool
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	return m.ReadBuf.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	return m.WriteBuf.Write(b)
}

func (m *MockConn) Close() error {
	m.Closed = true
	return nil
}

func (m *MockConn) LocalAddr() net.Addr              { return nil }
func (m *MockConn) RemoteAddr() net.Addr             { return nil }
func (m *MockConn) SetDeadline(time.Time) error      { return nil }
func (m *MockConn) SetReadDeadline(time.Time) error  { return nil }
func (m *MockConn) SetWriteDeadline(time.Time) error { return nil }

func TestTelnet_Send(t *testing.T) {
	mockConn := &MockConn{WriteBuf: &bytes.Buffer{}}
	input := strings.NewReader("test message")

	client := &Telnet{
		Conn: mockConn,
		In:   io.NopCloser(input),
	}

	err := client.Send()
	if err != nil {
		t.Errorf("Send() unexpected error: %v", err)
	}

	if mockConn.WriteBuf.String() != "test message" {
		t.Errorf("Send() wrote %q, want %q", mockConn.WriteBuf.String(), "test message")
	}
}

func TestTelnet_Receive(t *testing.T) {
	mockConn := &MockConn{ReadBuf: bytes.NewBufferString("test response")}
	output := &bytes.Buffer{}

	client := &Telnet{
		Conn: mockConn,
		Out:  output,
	}

	err := client.Receive()
	if err != nil {
		t.Errorf("Receive() unexpected error: %v", err)
	}

	if output.String() != "test response" {
		t.Errorf("Receive() read %q, want %q", output.String(), "test response")
	}
}

func TestTelnet_Close(t *testing.T) {
	mockConn := &MockConn{}
	client := &Telnet{Conn: mockConn}

	err := client.Close()
	if err != nil {
		t.Errorf("Close() unexpected error: %v", err)
	}

	if !mockConn.Closed {
		t.Error("Close() connection was not closed")
	}
}
