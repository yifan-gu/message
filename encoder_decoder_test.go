package message

import (
	"bytes"
	"reflect"
	"testing"

	"code.google.com/p/gogoprotobuf/proto"
	"github.com/go-epaxos/message/example"
)

// a simple test that encodes a message into a buffer
// and decodes a message out of the buffer.
func TestEncoderAndDecoder(t *testing.T) {
	buf := new(bytes.Buffer)

	inPb := &example.A{
		Description: "hello world!",
		Number:      1,
	}
	// UUID is 16 byte long
	for i := 0; i < 16; i++ {
		inPb.Id = append(inPb.Id, byte(i))
	}

	bytes, err := proto.Marshal(inPb)

	if err != nil {
		t.Fatal(err)
	}

	msg := NewMessage(0, bytes)

	e := NewMsgEncoder(buf)
	e.Encode(msg)

	outMsg := NewEmptyMessage()

	d := NewMsgDecoder(buf)
	d.Decode(outMsg)

	if !reflect.DeepEqual(msg, outMsg) {
		t.Fatal("Messages are not equal!")
	}

	outPb := new(example.A)

	proto.Unmarshal(outMsg.bytes, outPb)

	if !reflect.DeepEqual(outPb, inPb) {
		t.Fatal("Protos are not equal!")
	}
}
