# mattermost-echo-bot

## Overview
The most simple way to build a bot for [Mattermost](https://www.mattermost.org) (based on [Mattermost Bot Sample](https://github.com/mattermost/mattermost-bot-sample-golang)).

## Dependencies
* Go 1.12
* Mattermost 5.12

## How to run
* `go list -m all`
* Create a bot account (new feature for Mattermost 5.12)
* Fill in the constants in the beginning of the `main.go`
  * `MATTERMOST_API_REST` for the [REST API](https://api.mattermost.com)
  * `MATTERMOST_API_WS` for the [WebSocket API](https://api.mattermost.com/#tag/WebSocket)
  * `MATTERMOST_BOT_TOKEN` for the Access Token created in the bot account.
* `go run main.go`
* ???
* Profit!
