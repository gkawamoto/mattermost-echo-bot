package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

const (
	MATTERMOST_API_REST  = "http://localhost:8085"
	MATTERMOST_API_WS    = "ws://localhost:8085"
	MATTERMOST_BOT_TOKEN = "<bot token>"
)

var client *model.Client4
var webSocketClient *model.WebSocketClient

func main() {
	client = model.NewAPIv4Client(MATTERMOST_API_REST)
	client.MockSession(MATTERMOST_BOT_TOKEN)

	// Lets start listening to some channels via the websocket!
	webSocketClient, err := model.NewWebSocketClient4(MATTERMOST_API_WS, client.AuthToken)
	if err != nil {
		println("We failed to connect to the web socket")
		PrintError(err)
	}

	webSocketClient.Listen()

	go func() {
		for {
			select {
			case resp := <-webSocketClient.EventChannel:
				HandleWebSocketResponse(resp)
			}
		}
	}()

	// You can block forever with
	select {}
}

func SendMessage(msg string, channelId string) {
	post := &model.Post{}
	post.ChannelId = channelId
	post.Message = msg

	if _, resp := client.CreatePost(post); resp.Error != nil {
		println("We failed to send a message to the logging channel")
		PrintError(resp.Error)
	}
}

func HandleWebSocketResponse(event *model.WebSocketEvent) {
	// Lets only reponded to messaged posted events
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	var post = model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post != nil {
		var me, _ = client.GetMe("")
		if me == nil {
			print("Error getting me")
			return
		}
		// Ignoring me
		if post.UserId == me.Id {
			return
		}
		SendMessage("echo: "+post.Message, post.ChannelId)
	}
}

func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}
