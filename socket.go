package RLBotGo

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	schema "github.com/Trey2k/RLBotGo/flat"
)

type Socket struct {
	conn net.Conn
}

type payload struct {
	data     []byte
	dataType uint16
	dataSize uint16
}

type rlData interface {
	marshal() []byte
}

// InitConnection(port int) (Socket, error) Initiate the connection to RLBot returns a socket and a error on failure.
// Default port is 23234
func InitConnection(port int) (Socket, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	socket := Socket{
		conn: conn,
	}
	return socket, err
}

// SendMessage(dataType uint16, data rlData) error Send a data payload to RLBot, returns a error on failure
func (socket *Socket) SendMessage(dataType uint16, data rlData) error {
	dataTypePayload := make([]byte, 2)
	binary.BigEndian.PutUint16(dataTypePayload, dataType)

	payload := data.marshal()

	size := make([]byte, 2)
	binary.BigEndian.PutUint16(size, uint16(len(payload)))

	bytes := append([]byte{}, dataTypePayload...)
	bytes = append(bytes, size...)

	bytes = append(bytes, payload...)

	_, err := socket.conn.Write(bytes)
	return err
}

func (socket *Socket) startReadingBytes(payloadChannel chan *payload) error {
	for {
		dataInfo := make([]byte, 4)
		_, err := io.ReadFull(socket.conn, dataInfo)
		if err != nil && err != io.EOF {
			return err
		}

		dataType := binary.BigEndian.Uint16(dataInfo[:2])
		dataSize := binary.BigEndian.Uint16(dataInfo[2:])

		data := make([]byte, dataSize)
		_, err = io.ReadFull(socket.conn, data)
		if err != nil && err != io.EOF {
			return err
		}

		payloadChannel <- &payload{
			data:     data,
			dataType: dataType,
			dataSize: dataSize,
		}

	}
}

// SetTickHandler(handler func(gameState *GameState, socket *Socket) Set your tick handler function and start listening for gameTickPackets
func (socket *Socket) SetTickHandler(handler func(gameState *GameState, socket *Socket)) {

	gameState := &GameState{}
	gameState.BallPrediction = &BallPrediction{}
	gameState.FieldInfo = &FieldInfo{}
	gameState.GameTick = &GameTickPacket{}
	gameState.MatchSettigns = &MatchSettings{}

	gameState.FieldInfoOK = false
	gameState.MatchSettingsOK = false

	payloadChan := make(chan *payload, 5) // Makeing a payload channel with a buffer size of 5

	go socket.startReadingBytes(payloadChan) // Start reading packets in go routine and sending them over a channel

	for {
		payload := <-payloadChan
		switch payload.dataType {
		case DataType_TickPacket:
			flatGameTick := schema.GetRootAsGameTickPacket(payload.data, 0)
			gameState.GameTick = &GameTickPacket{} // Restting to 0 values just in case
			gameState.GameTick.unmarshal(flatGameTick)
			handler(gameState, socket)

		case DataType_FieldInfo:
			faltFieldInfo := schema.GetRootAsFieldInfo(payload.data, 0)
			gameState.FieldInfoOK = true
			gameState.FieldInfo.unmarshal(faltFieldInfo)

		case DataType_MatchSettings:
			flatMatchSettings := schema.GetRootAsMatchSettings(payload.data, 0)
			gameState.MatchSettingsOK = true
			gameState.MatchSettigns.unmarshal(flatMatchSettings)

		case DataType_BallPrediction:
			flatBallPrediction := schema.GetRootAsBallPrediction(payload.data, 0)
			gameState.BallPrediction.unmarshal(flatBallPrediction)
		}
	}
}
