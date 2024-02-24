package apid

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleParse() {
	apid, _ := Parse("test_6hoOwWlutwzIKWFCp54MUb")

	// Output:
	// test ce5cc4ed-0201-4bd9-82b8-27ece33bce6b
	fmt.Println(apid.Prefix, apid.UUID)
}

func TestNew(t *testing.T) {
	apid := New("test")

	assert.Equal(t, "test", apid.Prefix)
	assert.NotNil(t, apid.UUID)
	assert.NotEqual(t, uuid.Nil, apid.UUID)
}

func TestNewWithUUID(t *testing.T) {
	uuid := uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b")
	apid := NewWithUUID("test", uuid)

	assert.Equal(t, "test", apid.Prefix)
	assert.Equal(t, uuid, apid.UUID)
}

func TestParse(t *testing.T) {
	apid, err := Parse("test_6hoOwWlutwzIKWFCp54MUb")
	require.NoError(t, err)

	assert.Equal(t, APID{
		Prefix: "test",
		UUID:   uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"),
	}, apid)
}

func TestAPID_String(t *testing.T) {
	apid := APID{
		Prefix: "test",
		UUID:   uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"),
	}

	assert.Equal(t, "test_6hoOwWlutwzIKWFCp54MUb", apid.String())
}
