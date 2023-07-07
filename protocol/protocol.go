package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	CMD_SET byte = iota
	CMD_GET
	CMD_DELETE
	CMD_HAS_KEY
	CMD_GET_KEYS
)

func ParseAction(r io.Reader) (any, error) {
	var cmd byte
	if err := binary.Read(r, binary.LittleEndian, &cmd); err != nil {
		return nil, err
	}

	switch cmd {
	case CMD_GET:
		return parseGetCommand(r), nil

	default:
		return nil, fmt.Errorf("invalid command")
	}

}

type GetRequest struct {
	key []byte
}
type GetResponse struct{}

func CreateGetRequest(key []byte) *GetRequest {
	return &GetRequest{
		key: key,
	}
}

func (r GetRequest) Serialize() []byte {
	//Creation of the buffer where we sill store the
	//serialized GetRequest
	result := new(bytes.Buffer)
	//Write the kind of request: CMD_GET
	binary.Write(result, binary.LittleEndian, CMD_GET)
	//Write the key
	keyLen := int32(len(r.key))
	binary.Write(result, binary.LittleEndian, keyLen)
	binary.Write(result, binary.LittleEndian, r.key)

	return result.Bytes()
}
func parseGetCommand(r io.Reader) *GetRequest {
	rqt := GetRequest{}
	var keyLen int32
	binary.Read(r, binary.LittleEndian, &keyLen)
	rqt.key = make([]byte, keyLen)
	binary.Read(r, binary.LittleEndian, &rqt.key)
	return &rqt
}
