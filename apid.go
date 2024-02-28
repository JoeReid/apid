package apid

import (
	"fmt"
	"math/big"
	"regexp"

	"github.com/google/uuid"
)

const DefaultSeparator = "_"

var (
	ErrInvalidAPID      = fmt.Errorf("invalid APID")
	ErrInvalidSeparator = fmt.Errorf("invalid APID separator")
)

var (
	apidRegex      = regexp.MustCompile(`^([a-zA-Z0-9]+)([^a-zA-Z0-9])([a-zA-Z0-9]+)$`)
	separatorRegex = regexp.MustCompile(`^[^a-zA-Z0-9]$`)
)

type APID struct {
	Prefix    string
	UUID      uuid.UUID
	separator string
}

func (a APID) String() string {
	return a.Prefix + a.Separator() +
		big.NewInt(0).SetBytes(a.UUID[:]).Text(62)
}

func (a APID) Separator() string {
	if a.separator == "" {
		return DefaultSeparator
	}

	return a.separator
}

func New(prefix string) *APID {
	return &APID{
		Prefix:    prefix,
		UUID:      uuid.New(),
		separator: DefaultSeparator,
	}
}

func Parse(s string) (*APID, error) {
	matches := apidRegex.FindStringSubmatch(s)
	if len(matches) != 4 {
		return nil, ErrInvalidAPID
	}

	var i big.Int
	if _, ok := i.SetString(matches[3], 62); !ok {
		return nil, ErrInvalidAPID
	}

	return &APID{
		Prefix:    matches[1],
		UUID:      uuid.UUID(i.Bytes()),
		separator: matches[2],
	}, nil
}

func WithSeparator(s string, a APID) (*APID, error) {
	if !separatorRegex.MatchString(s) {
		return nil, ErrInvalidSeparator
	}

	return &APID{
		Prefix:    a.Prefix,
		UUID:      a.UUID,
		separator: s,
	}, nil
}
