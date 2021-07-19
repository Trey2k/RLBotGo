package RLBotGo

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"unsafe"

	schema "github.com/Trey2k/RLBotGo/flat"
)

type Socket struct {
	Conn    net.Conn
	BinBuff *bytes.Buffer
}

type rlData interface {
	Marshel() []byte
}

func InitConnection(port int) (Socket, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	socket := Socket{
		Conn:    conn,
		BinBuff: new(bytes.Buffer),
	}
	return socket, err
}

func (socket *Socket) SendMessage(dataType uint16, data rlData) error {
	dataTypePayload := make([]byte, 2)
	binary.BigEndian.PutUint16(dataTypePayload, dataType)

	size := make([]byte, 2)
	binary.BigEndian.PutUint16(size, uint16(unsafe.Sizeof(data)))

	bytes := append([]byte{}, dataTypePayload...)
	bytes = append(bytes, size...)

	bytes = append(bytes, data.Marshel()...)
	fmt.Println(bytes)
	_, err := socket.BinBuff.Write(bytes)
	if err != nil {
		return errors.New("rocket league is gay " + err.Error())
	}

	// Big Penis In Town

	_, err = socket.Conn.Write(socket.BinBuff.Bytes())
	return err
}

func (socket *Socket) SetTickHandler(handler func(gameTick *GameTickPacket, socket *Socket)) error {

	payload := make([]byte, 4096) //Trey, Change me!
	for {

		n, err := socket.Conn.Read(payload)
		if err != nil && err != io.EOF {
			return err
		}
		if n <= 4 { // Make sure we get a full packet
			continue
		}

		switch binary.BigEndian.Uint16(payload[:2]) {
		case 1:
			flatGameTick := schema.GetRootAsGameTickPacket(payload, 4)
			gameTick := &GameTickPacket{}
			gameTick.unmarshal(flatGameTick)
			handler(gameTick, socket)
		case 2:

		case 3:

		case 4:

		}
	}

}
