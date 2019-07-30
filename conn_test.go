package ipset

import (
	"testing"

	"github.com/mdlayher/netlink"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type queryMock struct {
	mock.Mock
}

func (q queryMock) Query(nlm netlink.Message) ([]netlink.Message, error) {
	args := q.Called(nlm.Data)
	return args.Get(0).([]netlink.Message), args.Error(1)
}

func TestConn_Create(t *testing.T) {
	assert2 := assert.New(t)

	m := new(queryMock)

	data := []byte{
		0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x01, 0x00, 0x06, 0x00, 0x00, 0x00, 0x08, 0x00, 0x02, 0x00,
		0x66, 0x6f, 0x6f, 0x00, 0x0d, 0x00, 0x03, 0x00, 0x68, 0x61, 0x73, 0x68, 0x3a, 0x6d, 0x61, 0x63,
		0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x00, 0x05, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x07, 0x80,
	}
	m.On("Query", data).Return([]netlink.Message{}, nil)

	c := Conn{Conn: m}
	assert2.NoError(c.Create("foo", "hash:mac", 0, 0))

	m.AssertExpectations(t)
}
