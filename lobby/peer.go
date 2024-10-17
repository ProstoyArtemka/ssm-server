package lobby

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"server/packets"
	"server/utils"
)

type Peer struct {
	ID int

	Connection net.Conn
	Connected  bool

	Name string

	Lobby byte
}

var LastPeerID = 0

func HandleConnection(connection net.Conn, err error) {
	peer := &Peer{Connection: connection, Connected: false, ID: LastPeerID}
	LastPeerID++

	if err != nil {
		log.Fatal(err)
	}

	go peer.HandleRequest()
}

func (peer *Peer) HandleRequest() {
	reader := bufio.NewReader(peer.Connection)

	for {
		message, error := reader.ReadBytes(0xAA)

		if error != nil {
			peer.Disconnect()

			return
		}

		message_length := len(message)

		if message_length == 0 {
			return
		}

		packet_type := message[0]

		handlePacket(packet_type, message[1:message_length-1], peer)
	}
}

func handlePacket(packet_type byte, message []byte, peer *Peer) {
	if packet_type == packets.PACKET_CONNECTION {
		if len(message) == 0 {
			peer.Disconnect()

			return
		}

		connected_lobby := message[0]
		lobby := GetLobby(connected_lobby)

		if lobby == nil {
			peer.Disconnect()

			return
		}

		lobby.AddPeer(peer)
	}

	if !peer.Connected {
		peer.Disconnect()

		return
	}

	if packet_type == packets.PACKET_MESSAGE {
		peer.handleMessage(message)
	}

	if packet_type == packets.PACKET_PULSE {
		fmt.Println("Pulse!")
	}
}

func (peer *Peer) WritePacket(packet_type byte, message []byte) error {
	full_message := []byte{packet_type}
	full_message = append(full_message, message...)

	_, err := peer.Connection.Write(full_message)

	return err
}

func (peer *Peer) ChangeName(newName string) {
	peer.Name = newName
	bytesName := utils.StringToBytes(newName)

	lobby := GetLobby(peer.Lobby)

	message := []byte{packets.MESSAGE_PEER_INFO_CHANGED, packets.PEER_INFO_NAME_CHANGED}
	message = append(message, bytesName...)

	lobby.WritePacketExclude(packets.PACKET_MESSAGE, message, []int{peer.ID})
}

func (peer *Peer) Disconnect() {
	lobby := GetLobby(peer.Lobby)

	if lobby != nil {
		lobby.DisconnectPeer(peer)
	}

	message := []byte{packets.MESSAGE_ERROR, packets.ERROR_UNKNOWN}
	peer.WritePacket(packets.PACKET_MESSAGE, []byte(message))

	peer.Connection.Close()
}
