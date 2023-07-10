package protocol_test

import (
	"bytes"
	"testing"

	"github.com/arejula27/distributed-cache/protocol"
	"github.com/stretchr/testify/assert"
)

func TestProtocol(t *testing.T) {
	t.Run("Parse get action", func(t *testing.T) {

		//Create request

		rqt := protocol.CreateGetRequest([]byte("exampleKey"))
		conn_mock := bytes.NewReader(rqt.Serialize())
		parsedAction, err := protocol.ParseAction(conn_mock)
		assert.Nil(t, err)
		assert.Equal(t, rqt, parsedAction)

	})
	t.Run("Parse set action", func(t *testing.T) {
		rqt := protocol.CreateSetRequest([]byte("exampleKey"), []byte("exampleValue"))
		conn_mock := bytes.NewReader(rqt.Serialize())
		parsedAction, err := protocol.ParseAction(conn_mock)
		assert.Nil(t, err)
		assert.Equal(t, rqt, parsedAction)
	})
}
