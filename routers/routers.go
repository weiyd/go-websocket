package routers

import (
	"go-websocket/api/bind2group"
	"go-websocket/api/getonlinelist"
	"go-websocket/api/register"
	"go-websocket/api/send2client"
	"go-websocket/api/send2clients"
	"go-websocket/api/send2group"
	"go-websocket/servers"
	"net/http"
)

func Init() {
	registerHandler := &register.Controller{}
	sendToClientHandler := &send2client.Controller{}
	sendToClientsHandler := &send2clients.Controller{}
	sendToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}
	getGroupListHandler := &getonlinelist.Controller{}

	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(sendToClientHandler.Run))
	http.HandleFunc("/api/send_to_clients", AccessTokenMiddleware(sendToClientsHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(sendToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))
	http.HandleFunc("/api/get_online_list", AccessTokenMiddleware(getGroupListHandler.Run))

	servers.StartWebSocket()

	go servers.WriteMessage()
}
