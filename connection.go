package main

import (
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func handleConnection(conn net.Conn, sshServerConfig *ssh.ServerConfig) {
	defer conn.Close()
	serverConn, newChannels, requests, err := ssh.NewServerConn(conn, sshServerConfig)
	if err != nil {
		log.Println("Failed to establish SSH connection:", err)
		return
	}

	getLogEntry(serverConn).Infoln("SSH connection established")

	go handleGlobalRequests(requests, serverConn)

	channelID := 0
	for newChannel := range newChannels {
		go handleNewChannel(newChannel, channelMetadata{conn: serverConn, channelID: channelID})
		channelID++
	}
}
