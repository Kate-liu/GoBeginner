package main

import (
	"fmt"
	"github.com/Kate-liu/GoBeginner/practiceproject/webtcpserver/webtcpservermetrics/frame"
	"github.com/Kate-liu/GoBeginner/practiceproject/webtcpserver/webtcpservermetrics/metrics"
	"github.com/Kate-liu/GoBeginner/practiceproject/webtcpserver/webtcpservermetrics/packet"

	"net"
)

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}

	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	metrics.ClientConnected.Inc() // 连接建立，ClientConnected加1
	defer func() {
		metrics.ClientConnected.Dec() // 连接断开，ClientConnected减1
		c.Close()
	}()
	frameCodec := frame.NewMyFrameCodec()

	for {
		// read from the connection
		// decode the frame to get the payload
		// the payload is undecoded packet
		framePayload, err := frameCodec.Decode(c)
		if err != nil {
			fmt.Println("handleConn: frame decode error:", err)
			return
		}
		metrics.ReqRecvTotal.Add(1) // 收到并解码一个消息请求，ReqRecvTotal消息计数器加1

		// do something with the packet
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
			return
		}

		// write ack frame to the connection
		err = frameCodec.Encode(c, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
		metrics.RspSendTotal.Add(1) // 返回响应后，RspSendTotal消息计数器减1
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle the new connection.
		go handleConn(c)
	}
}
