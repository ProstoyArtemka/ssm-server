package lobby

import (
	"bufio"
	"fmt"
	"net"
	"server/packets"
)

type Peer struct {
	Connection net.Conn
	Connected  bool

	Lobby []byte
}

type Lobby struct {
	ID             []byte
	ConnectedUsers []Peer
}

var Lobbies []Lobby

func genereateLobbyID() []byte {
	// Do some smart shit later

	return make([]byte, 4)
}

func NewLobby(admin Peer) (lobby *Lobby) {
	lobbyID := genereateLobbyID()

	connectedUsers := make([]Peer, 1)
	connectedUsers[0] = admin

	newLobby := &Lobby{ConnectedUsers: connectedUsers, ID: lobbyID}

	Lobbies = append(Lobbies, *newLobby)

	return newLobby
}

func (peer *Peer) HandleRequest() {
	reader := bufio.NewReader(peer.Connection)

	for {
		message, error := reader.ReadBytes(0xAA)

		if error != nil {
			peer.Connection.Close()

			return
		}

		message_length := len(message)

		if message_length == 0 {
			return
		}

		packet_type := message[0]
		HandlePacket(packet_type, message[1:message_length-1], peer)
	}
}

func HandlePacket(packet_type byte, message []byte, peer *Peer) {
	if packet_type == packets.PACKET_CONNECTION {
		peer.Connected = true

		connected_lobby := message

		fmt.Println(connected_lobby)
	}

	if !peer.Connected {
		peer.Connection.Close()

		return
	}

	if packet_type == packets.PACKET_MESSAGE {
		fmt.Println(message)
	}

	if packet_type == packets.PACKET_PULSE {
		fmt.Println("Pulse!")
	}
}
