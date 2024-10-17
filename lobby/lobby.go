package lobby

import (
	"server/packets"
	"server/utils"
)

type Lobby struct {
	ID    byte
	Peers []Peer

	Name string
}

const MaxLobbiesCount = 16

var Lobbies []Lobby

func genereateLobbyID() byte {
	for i := 0x00; i < MaxLobbiesCount; i++ {

		hasThisID := false
		for j := 0; j < len(Lobbies); j++ {
			lobby := Lobbies[j]

			if lobby.ID == (byte)(i) {
				break
			}

			hasThisID = true
		}

		if !hasThisID {
			return (byte)(i)
		}
	}

	return 0xAA
}

func NewLobby(admin Peer, name string) (lobby *Lobby) {
	lobbyID := genereateLobbyID()

	if lobbyID == 0xAA {
		return nil
	}

	peers := make([]Peer, 1)
	peers[0] = admin

	newLobby := &Lobby{Peers: peers, ID: lobbyID, Name: name}

	Lobbies = append(Lobbies, *newLobby)

	return newLobby
}

func GetLobby(id byte) (lobby *Lobby) {
	for i := 0; i < len(Lobbies); i++ {
		lobby := Lobbies[i]

		if lobby.ID == id {
			return &lobby
		}
	}

	return nil
}

func (lobby *Lobby) AddPeer(peer *Peer) {
	message := []byte{packets.MESSAGE_OTHER_PEER_CONNECTED}

	lobby.WritePacket(packets.PACKET_MESSAGE, message)

	peer.Connected = true
	peer.Lobby = lobby.ID

	lobby.Peers = append(lobby.Peers, *peer)
}

func (lobby *Lobby) DisconnectPeer(peer *Peer) {
	new_peers := make([]Peer, len(lobby.Peers)-1)

	for i := 0; i < len(lobby.Peers); i++ {
		loby_peer := lobby.Peers[i]

		if loby_peer.ID == peer.ID {
			continue
		}

		new_peers = append(new_peers, loby_peer)

		message := []byte{packets.MESSAGE_PEER_DISCONNECTED, (byte)(peer.ID)}
		loby_peer.WritePacket(packets.PACKET_MESSAGE, message)
	}

	lobby.Peers = new_peers
}

func (lobby *Lobby) WritePacket(packet_type byte, message []byte) bool {
	for i := 0; i < len(lobby.Peers); i++ {
		peer := lobby.Peers[i]

		if err := peer.WritePacket(packet_type, message); err != nil {
			return false
		}
	}

	return true
}

func (lobby *Lobby) WritePacketExclude(packet_type byte, message []byte, exclude []int) bool {
	for i := 0; i < len(lobby.Peers); i++ {
		peer := lobby.Peers[i]

		if utils.Contains[int](exclude, peer.ID) {
			continue
		}

		if err := peer.WritePacket(packet_type, message); err != nil {
			return false
		}
	}

	return true
}
