/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/5/26 0:40
 */

package Ping

import (
	"errors"
	"net"
	"time"
	"os"
	"bytes"
	"log"
)

// taken from http://golang.org/src/pkg/net/ipraw_test.go

const (
	icmpv4EchoRequest = 8
	icmpv4EchoReply   = 0
	icmpv6EchoRequest = 128
	icmpv6EchoReply   = 129
)

type icmpMessage struct {
	Type 		int
	Code 		int
	Checksum 	int
	Body 		icmpMessageBody
}

// icmpMessageBody接口包含了2个方法
type icmpMessageBody interface {
	Len() int
	Marshal() ([]byte, error)
}

// Marshal方法返回ICMP的二进制request/reply
func (m *icmpMessage) Marshal() ([]byte, error) {
	b := []byte{byte(m.Type), byte(m.Code), 0, 0}
	if m.Body != nil && m.Body.Len() != 0 {
		mb, err := m.Body.Marshal()
		if err != nil {
			return nil, err
		}
		b = append(b, mb...)
	}
	switch m.Type {
	case icmpv6EchoRequest, icmpv6EchoReply:
		return b, nil
	}
	// checksum范围
	csumcv := len(b) - 1
	s := uint32(0)
	for i := 0; i < csumcv; i += 2 {
		s += uint32(b[i+1]) << 8 | uint32(b[i])
	}
	if csumcv & 1 == 0 {
		s += uint32(b[csumcv])
	}
	s = s >> 16 + s & 0xffff
	s = s + s >> 16
	b[2] ^= byte(^s & 0xff)
	b[3] ^= byte(^s >> 8)
	return b, nil
}

// 解析ICMP消息为icmpMessage
func parseICMPMessage(b []byte) (*icmpMessage, error) {
	msglen := len(b)
	if msglen < 4 {
		return nil, errors.New("message too short")
	}
	m := &icmpMessage{
		Type: 		int(b[0]),
		Code: 		int(b[1]),
		Checksum: 	int(b[2])<<8 | int(b[3]),
	}
	if msglen > 4 {
		var err error
		switch m.Type {
		case icmpv4EchoRequest, icmpv4EchoReply, icmpv6EchoRequest, icmpv6EchoReply:
			m.Body, err = parseICMPEcho(b[4:])
			if err != nil {
				return nil, err
			}
		}
	}
	return m, nil
}

// icmpEcho代表一个ICMP echo request/reply消息体
type icmpEcho struct {
	ID 		int
	Seq 	int
	Data 	[]byte
}

func (p *icmpEcho) Len() int {
	if p == nil {
		return 0
	}
	return 4 + len(p.Data)
}

// Marshal方法返回二进制编码的ICMP echo request/reply消息体
func (p *icmpEcho) Marshal() ([]byte, error) {
	b := make([]byte, 4+len(p.Data))
	b[0], b[1] = byte(p.ID >> 8), byte(p.ID & 0xff)
	b[2], b[3] = byte(p.Seq >> 8), byte(p.Seq & 0xff)
	copy(b[4:], p.Data)
	return b, nil
}

// parseICMPEcho把b解析成一个ICMP echo request/reply消息体
func parseICMPEcho(b []byte) (*icmpEcho, error) {
	bodylen := len(b)
	p := &icmpEcho{
		ID: int(b[0])<<8 | int(b[1]),
		Seq: int(b[2])<<8 | int(b[3]),
	}
	if bodylen > 4 {
		p.Data = make([]byte, bodylen-4)
		copy(p.Data, b[4:])
	}
	return p, nil
}

func Ping(address string, timeout int) error {
	c, err := net.Dial("ip4:icmp", address)
	if err != nil {
		return err
	}
	c.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	defer c.Close()

	typ := icmpv4EchoRequest
	xid, xseq := os.Getpid()&0xffff, 1
	wb, err := (&icmpMessage{
		Type: typ, Code: 0,
		Body: &icmpEcho{
			ID: xid, Seq: xseq,
			Data: bytes.Repeat([]byte("Go Go Gadget Ping!!!"), 3),
		},
	}).Marshal()
	if err != nil {
		return err
	}
	if _, err = c.Write(wb); err != nil {
		return err
	}
	var m *icmpMessage
	rb := make([]byte, 20+len(wb))
	for {
		if _, err = c.Read(rb); err != nil {
			return err
		}
		rb = ipv4Payload(rb)
		if m, err = parseICMPMessage(rb); err != nil {
			return err
		}
		switch m.Type {
		case icmpv4EchoRequest, icmpv6EchoRequest:
			continue
		}
		break
	}
	return nil
}

func ipv4Payload(b []byte) []byte {
	if len(b) < 20 {
		return b
	}
	hdrlen := int(b[0]&0x0f) << 2
	return b[hdrlen:]
}