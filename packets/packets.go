package packets

const (
	PACKET_CONNECTION = 0x01

	PACKET_PULSE   = 0x03
	PACKET_MESSAGE = 0x04

	PACKET_END = 0xAA
)

const (
	MESSAGE_CREATE_LOBBY = 0x01
	MESSAGE_LIST_LOBBIES = 0x07

	MESSAGE_OTHER_PEER_CONNECTED = 0x02
	MESSAGE_PEER_INFO_CHANGED    = 0x03
	MESSAGE_PEER_DISCONNECTED    = 0x04

	MESSAGE_ERROR = 0x05

	MESSAGE_CREATE_GAME = 0x06
)

const (
	CLIENT_MESSAGE_CONNECT_TO_LOBBY = 0x01
	CLIENT_MESSAGE_CHANGE_NAME      = 0x02

	CLIENT_MESSAGE_LIST_LOBBIES = 0x3
)

const (
	PEER_INFO_NAME_CHANGED = 0x01
)

const (
	ERROR_UNKNOWN = 0x01
)
