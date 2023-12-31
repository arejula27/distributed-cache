package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	CMD_SET byte = iota
	CMD_GET
	CMD_DELETE
	CMD_HAS_KEY
	CMD_GET_KEYS
	CMD_EXIT
	CMD_ERROR
)

func ParseAction(r io.Reader) (any, error) {
	var cmd byte
	if err := binary.Read(r, binary.LittleEndian, &cmd); err != nil {
		return nil, err
	}

	switch cmd {
	case CMD_GET:
		return parseGetCommand(r), nil
	case CMD_SET:
		return parseSetCommand(r), nil

	default:
		return nil, fmt.Errorf("invalid command")
	}

}

func ParseResponse(r io.Reader) (any, error) {
	var cmd byte
	if err := binary.Read(r, binary.LittleEndian, &cmd); err != nil {
		return nil, err
	}

	switch cmd {
	case CMD_GET:
		return parseGetResponse(r), nil
	case CMD_ERROR:
		return parseErrorResponse(r), nil

	default:
		return nil, fmt.Errorf("invalid command")
	}
}

type ErrorResponse struct {
	Error error
}

func (r ErrorResponse) Serialize() []byte {
	//Creation of the buffer where we will store the
	//serialized GetRequest
	result := new(bytes.Buffer)
	//Write the kind of request: CMD_ERROR
	binary.Write(result, binary.LittleEndian, CMD_ERROR)
	//Write the error
	errorLen := int32(len(r.Error.Error()))
	binary.Write(result, binary.LittleEndian, errorLen)
	binary.Write(result, binary.LittleEndian, []byte(r.Error.Error()))

	return result.Bytes()
}
func parseErrorResponse(r io.Reader) *ErrorResponse {
	response := ErrorResponse{}
	var errorLen int32
	binary.Read(r, binary.LittleEndian, &errorLen)
	errorRaw := make([]byte, errorLen)
	binary.Read(r, binary.LittleEndian, &errorRaw)
	response.Error = errors.New(string(errorRaw))
	return &response

}

type GetRequest struct {
	key []byte
}
type GetResponse struct {
	Value []byte
}

func CreateGetRequest(key []byte) *GetRequest {
	return &GetRequest{
		key: key,
	}
}
func (r GetRequest) Key() string {
	return string(r.key)
}
func (r GetRequest) Serialize() []byte {
	//Creation of the buffer where we will store the
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

func (r GetResponse) Serialize() []byte {
	//Creation of the buffer where we will store the
	//serialized GetRequest
	result := new(bytes.Buffer)
	//Write the kind of request: CMD_GET

	binary.Write(result, binary.LittleEndian, CMD_GET)
	valueLen := int32(len(r.Value))
	binary.Write(result, binary.LittleEndian, valueLen)
	binary.Write(result, binary.LittleEndian, r.Value)

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
func parseGetResponse(r io.Reader) *GetResponse {
	response := GetResponse{}
	var valueLen int32
	binary.Read(r, binary.LittleEndian, &valueLen)
	response.Value = make([]byte, valueLen)
	binary.Read(r, binary.LittleEndian, &response.Value)
	return &response
}

type SetRequest struct {
	key   []byte
	value []byte
}
type SetResponse struct{}

func CreateSetRequest(key, value []byte) *SetRequest {
	return &SetRequest{
		key:   key,
		value: value,
	}
}
func (r SetRequest) Key() string {
	return string(r.key)
}
func (r SetRequest) Value() string {
	return string(r.value)
}
func (r *SetRequest) Serialize() []byte {
	//Creation of the buffer where we will store the
	//serialized GetRequest
	result := new(bytes.Buffer)
	//Write the kind of request: CMD_GET
	binary.Write(result, binary.LittleEndian, CMD_SET)
	//Write the key
	keyLen := int32(len(r.key))
	binary.Write(result, binary.LittleEndian, keyLen)
	binary.Write(result, binary.LittleEndian, r.key)
	//Write the value
	valueLen := int32(len(r.value))
	binary.Write(result, binary.LittleEndian, valueLen)
	binary.Write(result, binary.LittleEndian, r.value)

	return result.Bytes()
}
func parseSetCommand(r io.Reader) *SetRequest {
	rqt := SetRequest{}
	//parse key
	var keyLen int32
	binary.Read(r, binary.LittleEndian, &keyLen)
	rqt.key = make([]byte, keyLen)
	binary.Read(r, binary.LittleEndian, &rqt.key)
	//parse value
	var valueLen int32
	binary.Read(r, binary.LittleEndian, &valueLen)
	rqt.value = make([]byte, valueLen)
	binary.Read(r, binary.LittleEndian, &rqt.value)

	return &rqt
}

type ExitRequest struct{}
