package main

import (
	// "encoding/binary"
	"flag"
	// "fmt"
	// "encoding/binary"
	"github.com/golang/glog"
	// "github.com/golang/protobuf/proto"
	"github.com/oikomi/FishChatServer2/libnet"
	// mybinary "github.com/oikomi/FishChatServer2/libnet/binary"
	"fmt"
	// myproto "github.com/oikomi/FishChatServer2/protocol"
	"io"
	"io/ioutil"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "127.0.0.1:17000", "echo server address")
	flag.Parse()

	// session, err := libnet.Connect("tcp", addr, libnet.Packet(2, 1024*1024, 1024, binary.BigEndian, TestCodec{}))
	session, err := libnet.Connect("tcp", addr, TestCodec{})

	if err != nil {
		glog.Error("libnet.Connect error: ", err)
		panic(err)
	}

	// go func() {
	// 	var msg string
	// 	for {
	// 		if err := session.Receive(&msg); err != nil {
	// 			glog.Error("session.Receive error: ", err)
	// 			break
	// 		}
	// 		fmt.Printf("%s\n", msg)
	// 	}
	// }()

	for {
		// msg := &myproto.ReqMsgServer{
		// 	Label: proto.String("hello"),
		// }
		var msg string
		println("------")
		if _, err := fmt.Scanf("%s\n", &msg); err != nil {
			glog.Error("fmt.Scanf error: ", err)
			break
		}

		if err = session.Send(msg); err != nil {
			glog.Error("session.Send error: ", err)
			break
		}
	}

	session.Close()
	println("bye")
}

type TestCodec struct {
}

type TestEncoder struct {
	w io.Writer
}

type TestDecoder struct {
	r io.Reader
}

func (codec TestCodec) NewEncoder(w io.Writer) libnet.Encoder {
	return &TestEncoder{w}
}

func (codec TestCodec) NewDecoder(r io.Reader) libnet.Decoder {
	return &TestDecoder{r}
}

func (encoder *TestEncoder) Encode(msg interface{}) error {
	_, err := encoder.w.Write([]byte(msg.(string)))
	return err
}

func (decoder *TestDecoder) Decode(msg interface{}) error {
	d, err := ioutil.ReadAll(decoder.r)
	if err != nil {
		return err
	}
	*(msg.(*string)) = string(d)
	return nil
}