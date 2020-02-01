package net

import (
	"net"
	"time"
)

type ISocket struct {
	conn net.Conn
}
type ReadCheck func(chunk []byte) bool

func Dial(network, address string) (ISocket, error) {
	conn, err := net.Dial(network, address)
	return ISocket{conn: conn}, err
}

func (socket *ISocket) Read(b []byte) (n int, err error) {
	return socket.conn.Read(b)
}
func (socket *ISocket) Write(b []byte) (n int, err error) {
	return socket.conn.Write(b)
}
func (socket *ISocket) Close() error {
	return socket.conn.Close()
}
func (socket *ISocket) LocalAddr() net.Addr {
	return socket.conn.LocalAddr()
}
func (socket *ISocket) RemoteAddr() net.Addr {
	return socket.conn.RemoteAddr()
}
func (socket *ISocket) SetDeadline(t time.Time) error {
	return socket.conn.SetDeadline(t)
}
func (socket *ISocket) SetReadDeadline(t time.Time) error {
	return socket.conn.SetReadDeadline(t)
}
func (socket *ISocket) SetWriteDeadline(t time.Time) error {
	return socket.conn.SetWriteDeadline(t)
}

func (socket *ISocket) ReadWhile(b *[]byte, bufferSize int, check ReadCheck) (n int, err error) {
	buf := make([]byte, bufferSize)
	var bytesRead int = 0
	for {
		n, err := socket.conn.Read(buf)
		if err != nil {
			return bytesRead, err
		}
		bytesRead += n
		*b = append(*b, buf...)
		if !check(buf) {
			return bytesRead, nil
		}
	}
}

func (socket *ISocket) ReadUntil(b *[]byte, key []byte) (n int, err error) {
	var lastByte byte = 0;
	var firstByte bool = true;
	var streak int = 0;
	readCheck := func (chunk []byte) bool {
		if firstByte {
			firstByte = false;
		} else {
			if streak >= len(key) {
				return false
			} else if key[streak] == chunk[0] {
				streak++
			} else {
				streak = 0
			}
		}
		lastByte = chunk[0]
		return true
	}
	return socket.ReadWhile(b, 1, readCheck)
}
