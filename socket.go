package RLBotGo

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"net"

	schema "github.com/Trey2k/RLBotGo/flat"
)

type RLBot struct {
	conn             net.Conn
	debugRenderGroup *RenderGroup
	PlayerIndex      int32
}

type payload struct {
	data     []byte
	dataType uint16
	dataSize uint16
}

type rlData interface {
	marshal() []byte
}

// SendQuickChat This will allow your bot to be toxic. Who doesn't want that?
func (socket *RLBot) SendQuickChat(quickChatSelection int8, teamOnly bool) error {

	quickChat := &QuickChat{
		QuickChatSelection: quickChatSelection,
		TeamOnly:           teamOnly,
		PlayerIndex:        socket.PlayerIndex,
	}
	return socket.SendMessage(DataType_QuickChat, quickChat)
}

// SendReadyMessage Send the ready message to RLBot
func (socket *RLBot) SendReadyMessage(wantsBallPredictions, wantsQuickChat, wantsGameMessages bool) error {
	readyMsg := &ReadyMessage{
		WantsBallPredictions: wantsBallPredictions,
		WantsQuickChat:       wantsQuickChat,
		WantsGameMessages:    wantsGameMessages,
	}
	return socket.SendMessage(DataType_ReadyMessage, readyMsg)
}

// SendDesiredGameState Send a specific game state. Good for testing
func (socket *RLBot) SendDesiredGameState(desiredGameState *DesiredGameState) error {
	return socket.SendMessage(DataType_DesiredGameState, desiredGameState)
}

// Connect (port int) (Socket, error) Initiate the connection to RLBot returns a socket and a error on failure.
// Default port is 23234
func Connect(port int) (*RLBot, error) {
	var index = flag.Int("player-index", 0, "The player index for the bot")
	// Go has to know about these two otherwise will fail to launch.
	flag.String("rlbot-version", "0", "RLBot version")
	flag.String("rlbot-dll-directory", "0", "RLBot DLL dir")
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	socket := &RLBot{
		conn:        conn,
		PlayerIndex: int32(*index),
	}
	return socket, err
}

// SendMessage (dataType uint16, data rlData) error Send a data payload to RLBot, returns a error on failure
func (socket *RLBot) SendMessage(dataType uint16, data rlData) error {
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

// read game data from the RLBot API
func (socket *RLBot) startReadingBytes(payloadChannel chan *payload, errChan chan error) {
	for {
		dataInfo := make([]byte, 4)
		_, err := io.ReadFull(socket.conn, dataInfo)
		if err != nil && err != io.EOF {
			errChan <- err
			break
		}

		dataType := binary.BigEndian.Uint16(dataInfo[:2])
		dataSize := binary.BigEndian.Uint16(dataInfo[2:])

		data := make([]byte, dataSize)
		_, err = io.ReadFull(socket.conn, data)
		if err != nil && err != io.EOF {
			errChan <- err
			break
		}

		payloadChannel <- &payload{
			data:     data,
			dataType: dataType,
			dataSize: dataSize,
		}

	}
}

// SetGetInput (handler func(gameState *GameState, socket *Socket) Set your tick handler function and start listening for gameTickPackets
func (socket *RLBot) SetGetInput(handler func(gameState *GameState, socket *RLBot) *ControllerState) error {

	gameState := &GameState{}
	gameState.BallPrediction = &BallPrediction{}
	gameState.FieldInfo = &FieldInfo{}
	gameState.GameTick = &GameTickPacket{}
	gameState.MatchSettings = &MatchSettings{}

	gameState.FieldInfoOK = false
	gameState.MatchSettingsOK = false

	payloadChan := make(chan *payload, 5) // Makeing a payload channel with a buffer size of 5
	errChan := make(chan error)

	go socket.startReadingBytes(payloadChan, errChan) // Start reading packets in go routine and sending them over a channel

	for {
		select {
		case err := <-errChan: // Check for a error every loop
			return err
		default:
			payload := <-payloadChan // Hold until we get a payload
			switch payload.dataType {
			case DataType_TickPacket:
				flatGameTick := schema.GetRootAsGameTickPacket(payload.data, 0)
				gameState.GameTick = &GameTickPacket{} // Resetting to 0 values just in case
				gameState.GameTick.unmarshal(flatGameTick)
				input := handler(gameState, socket)
				// Get input from handler and send it
				if input != nil {
					playerInput := &PlayerInput{
						PlayerIndex:     socket.PlayerIndex,
						ControllerState: *input,
					}

					err := socket.SendMessage(DataType_PlayerInput, playerInput)
					if err != nil {
						return err
					}
				}

			case DataType_FieldInfo:
				flatFieldInfo := schema.GetRootAsFieldInfo(payload.data, 0)
				gameState.FieldInfoOK = true
				gameState.FieldInfo.unmarshal(flatFieldInfo)

			case DataType_MatchSettings:
				flatMatchSettings := schema.GetRootAsMatchSettings(payload.data, 0)
				gameState.MatchSettingsOK = true
				gameState.MatchSettings.unmarshal(flatMatchSettings)

			case DataType_BallPrediction:
				flatBallPrediction := schema.GetRootAsBallPrediction(payload.data, 0)
				gameState.BallPrediction.unmarshal(flatBallPrediction)
			}
		}
	}
}
