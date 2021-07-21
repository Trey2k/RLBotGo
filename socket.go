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

type rlData interface {
	marshal() []byte
}

func InitConnection(port int) (Socket, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	socket := Socket{
		conn: conn,
	}
	return socket, err
}

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

func (socket *Socket) SetTickHandler(handler func(gameState *GameState, socket *Socket)) error {

	//payload := make([]byte, 23504) //Trey, Change me!
	gameState := &GameState{}
	gameState.MatchSettingsOK = false
	gameState.FieldInfoOK = false
	for {

		data := make([]byte, 4)
		_, err := io.ReadFull(socket.conn, data)
		if err != nil && err != io.EOF {
			return err
		}

		dataType := binary.BigEndian.Uint16(data[:2])
		dataSize := binary.BigEndian.Uint16(data[2:])
		fmt.Println(dataSize)
		payload := make([]byte, dataSize)
		_, err = io.ReadFull(socket.conn, payload)
		if err != nil && err != io.EOF {
			return err
		}

		switch dataType {
		case DataType_TickPacket:
			flatGameTick := schema.GetRootAsGameTickPacket(payload, 0)
			gameState.GameTick = &GameTickPacket{} // Restting to 0 values just in case
			gameState.GameTick.unmarshal(flatGameTick)
			handler(gameState, socket)
		case DataType_FieldInfo:
			faltFieldInfo := schema.GetRootAsFieldInfo(payload, 0)
			gameState.FieldInfoOK = true
			gameState.FieldInfo.unmarshal(faltFieldInfo)
		case DataType_MatchSettings:
			flatMatchSettings := schema.GetRootAsMatchSettings(payload, 0)
			gameState.MatchSettingsOK = true
			gameState.MatchSettigns.unmarshal(flatMatchSettings)
			// TODO: Figure out why we are not sent MatchSettings
		case DataType_BallPrediction:
			flatBallPrediction := schema.GetRootAsBallPrediction(payload, 0)
			if flatBallPrediction.SlicesLength() == 360 {
				gameState.BallPrediction.unmarshal(flatBallPrediction)
			}
			// TODO: Fix Ball predictions, faltbuffers fails GetRootAsBallPrediction for some reason
		}
	}
}
