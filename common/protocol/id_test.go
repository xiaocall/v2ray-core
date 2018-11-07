package protocol_test

import (
	"testing"

	"v2ray.com/core/common/compare"
	. "v2ray.com/core/common/protocol"
	"v2ray.com/core/common/uuid"
	. "v2ray.com/ext/assert"
)

func TestCmdKey(t *testing.T) {
	assert := With(t)

	id := NewID(uuid.New())
	assert(compare.BytesAll(id.CmdKey(), 0), IsFalse)
}

func TestIdEquals(t *testing.T) {
	assert := With(t)

	id1 := NewID(uuid.New())
	id2 := NewID(id1.UUID())

	assert(id1.Equals(id2), IsTrue)
	assert(id1.String(), Equals, id2.String())
}
