package apid

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

var ErrInvalidApid = fmt.Errorf("invalid APID format")

var Nil = APID{}

type APID struct {
	Prefix string
	UUID   uuid.UUID
}

func (a APID) String() string {
	var i big.Int
	i.SetBytes(a.UUID[:])

	return a.Prefix + "_" + i.Text(62)
}

func New(prefix string) APID {
	return APID{
		Prefix: prefix,
		UUID:   uuid.New(),
	}
}

func NewWithUUID(prefix string, uuid uuid.UUID) APID {
	return APID{
		Prefix: prefix,
		UUID:   uuid,
	}
}

func Parse(s string) (APID, error) {
	spl := strings.Split(s, "_")
	if len(spl) != 2 {
		return Nil, ErrInvalidApid
	}

	var i big.Int
	if _, ok := i.SetString(spl[1], 62); !ok {
		return Nil, ErrInvalidApid
	}

	return APID{
		Prefix: spl[0],
		UUID:   uuid.UUID(i.Bytes()),
	}, nil
}
