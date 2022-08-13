package main

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/nats-io/nats.go"
)

func sendFileHandler(conn *nats.Conn, topic string, r io.ReadWriter) error {
	getDataSlice := func(buff []byte, maxPayload int64) int64 {
		if int64(len(buff)) > maxPayload {
			return maxPayload
		}
		return int64(len(buff))
	}

	maxPayload := conn.MaxPayload()
	rb := bufio.NewReader(r)
	t := time.NewTicker(10 * time.Millisecond)
	var buff []byte
	for {
		select {
		case _, ok := <-t.C:
			if !ok {
				return fmt.Errorf("timer channel closed")
			}
			if len(buff) > 0 {
				length := getDataSlice(buff, maxPayload)
				if e := conn.PublishMsg(&nats.Msg{
					Subject: topic,
					Data:    buff[:length],
				}); e != nil {
					return e
				}
				n := copy(buff, buff[length:])
				buff = buff[:n]
			}
		default:
			if b, err := rb.ReadByte(); err == nil {
				buff = append(buff, b)
			} else if err == io.EOF {
				if len(buff) == 0 {
					return nil
				}
			} else {
				return err
			}
		}
	}
}

func listenTopicHandler(conn *nats.Conn, topic string, w io.ReadWriter) error {
	msgs := make(chan *nats.Msg, 2)
	sb, err := conn.ChanSubscribe(topic, msgs)
	if err != nil {
		return err
	}
	defer func() {
		sb.Unsubscribe()
		close(msgs)
	}()
	for data := range msgs {
		if _, err := w.Write(data.Data); err != nil {
			return err
		}
	}
	return nil
}

func streamHandler(conn *nats.Conn, topic string, rw io.ReadWriter) error {
	_, err := conn.JetStream()
	if err != nil {
		return err
	}
	return nil
}
