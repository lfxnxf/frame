package ikio

import (
	"testing"
	"encoding/binary"
)

var (
	UIDKey = []byte("uid")
	FlowIDKey = []byte("flowid")
	CMDKey = []byte("cmd")
)

type UAMessage struct {
	UID, FlowID uint64
	CMD int32
	Body []byte
}

func (m UAMessage) Serialize() ([]byte, error) {
	return m.Body, nil
}

func (m UAMessage) Packet(pkt *RPCPacket) {
	buffer := make([]byte, 20)
	binary.BigEndian.PutUint64(buffer[0:8], m.UID)
	binary.BigEndian.PutUint64(buffer[8:16], m.FlowID)
	binary.BigEndian.PutUint32(buffer[16:20], uint32(m.CMD))
	pkt.AddHeader(UIDKey, buffer[0:8])
	pkt.AddHeader(FlowIDKey, buffer[8:16])
	pkt.AddHeader(CMDKey, buffer[16:20])
}


func TestRPCHeader(t *testing.T) {
	pkt := new(RPCPacket)
	message := UAMessage{
		UID: 17819812,
		FlowID: 1234,
		CMD: 12,
	}
	message.Packet(pkt)
	connID, ok1 := pkt.GetHeaderUint64(FlowIDKey)
	uid, ok2 := pkt.GetHeaderUint64(UIDKey)
	cmd, ok3 := pkt.GetHeaderUint32(CMDKey)
//	body := pkt.serializePacketHeader()
//	header, _ := parseHeader(body)
	t.Logf("%d %d %d %v %v %v\n", connID, uid, cmd, ok1, ok2, ok3)

	cmdh, _ := pkt.GetHeader(CMDKey)
	t.Logf("cmd %d %x", int32(binary.BigEndian.Uint32(cmdh)), cmdh)

	pkt.ForeachHeader(func(k, v []byte) error {
		t.Logf("%s %x", k, v)
		return nil
	})
}
