package apid

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ExampleParse() {
	a, _ := Parse("test_6hoOwWlutwzIKWFCp54MUb")

	// Output:
	// test ce5cc4ed-0201-4bd9-82b8-27ece33bce6b
	fmt.Println(a.Prefix, a.UUID)
}

func TestNew(t *testing.T) {
	a := New("test")

	assert.Equal(t, "test", a.Prefix)
	assert.NotEqual(t, uuid.Nil, a.UUID)
	assert.Equal(t, a.separator, DefaultSeparator)
}

func TestParse(t *testing.T) {
	var tests = []struct {
		input string
		want  *APID
		err   error
	}{
		{
			input: "test_6hoOwWlutwzIKWFCp54MUb",
			want: &APID{
				Prefix:    "test",
				UUID:      uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"),
				separator: DefaultSeparator,
			},
			err: nil,
		},
		{
			input: "test-6hoOwWlutwzIKWFCp54MUb",
			want: &APID{
				Prefix:    "test",
				UUID:      uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"),
				separator: "-",
			},
			err: nil,
		},
		{
			input: "test.6hoOwWlutwzIKWFCp54MUb",
			want: &APID{
				Prefix:    "test",
				UUID:      uuid.MustParse("ce5cc4ed-0201-4bd9-82b8-27ece33bce6b"),
				separator: ".",
			},
			err: nil,
		},
		{
			input: "test6hoOwWlutwzIKWFCp54MUb",
			want:  nil,
			err:   ErrInvalidAPID,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			got, err := Parse(tt.input)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWithSeparator(t *testing.T) {
	u := uuid.New()

	a2, err := WithSeparator(".", APID{Prefix: "test", UUID: u})
	assert.NoError(t, err)
	assert.Equal(t, &APID{
		Prefix:    "test",
		UUID:      u,
		separator: ".",
	}, a2)
}
