package cqrs_test

import (
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"

	"github.com/golang/protobuf/ptypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtobufMarshaler(t *testing.T) {
	marshaler := cqrs.ProtobufMarshaler{}

	when, err := ptypes.TimestampProto(time.Now())
	require.NoError(t, err)

	eventToMarshal := &TestProtobufEvent{
		Id:   watermill.NewULID(),
		When: when,
	}

	msg, err := marshaler.Marshal(eventToMarshal)
	require.NoError(t, err)

	eventToUnmarshal := &TestProtobufEvent{}
	err = marshaler.Unmarshal(msg, eventToUnmarshal)
	require.NoError(t, err)

	assert.EqualValues(t, eventToMarshal.String(), eventToUnmarshal.String())
}
