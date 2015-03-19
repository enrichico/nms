package nms

import "fmt"

const (
	MAXSIZE = 65535

	VERSION_AND_TYPE_FIELD_OFFSET uint16 = 0
	LENGTH_FIELD_OFFSET           uint16 = 1
	MSG_TYPE_FIELD_OFFSET         uint16 = 3
	SOURCE_ADDRESS_FIELD_OFFSET   uint16 = 4
	TARGET_ADDRESS_FIELD_OFFSET   uint16 = 8
	SESSION_ID_FIELD_OFFSET       uint16 = 12
	MSG_ID_FIELD_OFFSET           uint16 = 14
	HOP_COUNT_FIELD_OFFSET        uint16 = 16
	TTL_FIELD_OFFSET              uint16 = 17
	CHUNK_TYPE_FIELD              uint16 = 18
	RELIABLE_FIELD                uint16 = 19
)

/**
 * NetworkMessage wraps a buffer containing the NetworkMessageService's
 * payload and the NetworkMessageService's header and provides convenient
 * methods to retrieve them.
 * When a NetworkMessage is created, a new buffer of the proper size is
 * allocated and both the header and payload are <em>copied</em> into it.
 */
type NetworkMessage interface {
	IsMalformed() bool
	FixedHeaderLength() uint16

	MetaData() []byte
	MetaDataLen() uint16
	Msg() []byte
	MsgLen() uint16
}

type networkMessage struct {
	buffer    []byte
	netMsgLen uint16
}
