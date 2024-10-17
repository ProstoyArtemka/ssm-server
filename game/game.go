package game

import (
	"server/lobby"
	"server/packets"
)

type Game struct {
	Lobby byte
}

func NewGame(creator lobby.Peer) *Game {
	game := &Game{Lobby: creator.Lobby}

	lobby := lobby.GetLobby(creator.Lobby)
	peers := lobby.Peers

	for i := 0; i < len(peers); i++ {
		peer := peers[i]

		message := []byte{packets.MESSAGE_CREATE_GAME}

		peer.WritePacket(packets.PACKET_MESSAGE, message)
	}

	return game
}
