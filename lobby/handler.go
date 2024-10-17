package lobby

import (
	"server/packets"
	"server/utils"
)

type LobbyInfo struct {
	ID   byte
	Name string
}

func (peer *Peer) handleMessage(message []byte) {
	if len(message) == 0 {
		return
	}

	switch message_type := message[0]; message_type {

	case packets.CLIENT_MESSAGE_LIST_LOBBIES:
		lobbiesInfo := make([]byte, 0)

		for i := 0; i < len(Lobbies); i++ {
			lobby := Lobbies[i]
			lobbyName := utils.StringToBytes(lobby.Name)

			lobbiesInfo = append(lobbiesInfo, lobby.ID, (byte)(len(lobby.Name)))
			lobbiesInfo = append(lobbiesInfo, lobbyName...)
		}

		message := []byte{packets.MESSAGE_LIST_LOBBIES}
		message = append(message, lobbiesInfo...)

		peer.WritePacket(packets.PACKET_MESSAGE, message)

	case packets.CLIENT_MESSAGE_CONNECT_TO_LOBBY:
		lobbyID := message[0]
		lobby := GetLobby(lobbyID)

		if lobby == nil {
			peer.WritePacket(packets.PACKET_MESSAGE, []byte{packets.MESSAGE_ERROR, packets.ERROR_UNKNOWN})

			return
		}

		lobby.AddPeer(peer)

	case packets.CLIENT_MESSAGE_CHANGE_NAME:
		name := (string)(message)

		peer.ChangeName(name)
	}
}
