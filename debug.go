package RLBotGo

func (socket *RLBot) DebugMessageAdd(text string) error {

	if socket.debugRenderGroup == nil {
		socket.debugRenderGroup = &RenderGroup{Id: 1}
	}
	renderGroup := socket.debugRenderGroup

	var message RenderMessage
	message.Color = Color{A: 255, R: 46, G: 255, B: 0}
	var start float32 = 20 * float32(len(renderGroup.RenderMessages))
	message.Start = Vector3{X: 0, Y: start, Z: 0}
	message.End = Vector3{X: 0, Y: start + 1, Z: 0}
	message.ScaleX = 1
	message.ScaleY = 1
	message.IsFilled = true
	message.RenderType = RenderType_DrawString2D
	message.Text = text
	renderGroup.RenderMessages = append(renderGroup.RenderMessages, message)

	return socket.SendMessage(DataType_RenderGroup, socket.debugRenderGroup)
}

func (socket *RLBot) DebugMessageClear() error {
	if socket.debugRenderGroup == nil {
		socket.debugRenderGroup = &RenderGroup{Id: 1}
	}
	renderGroup := socket.debugRenderGroup

	renderGroup.RenderMessages = []RenderMessage{}

	return socket.SendMessage(DataType_RenderGroup, socket.debugRenderGroup)
}
