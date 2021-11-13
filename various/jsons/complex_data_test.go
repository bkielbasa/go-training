package jsons

import (
	"encoding/json"
	"testing"

	"github.com/matryer/is"
)

// in megahertz
type frequency int

type PC struct {
	Type string
	CPU  *frequency
	GPU  *frequency
}

func TestComplexData(t *testing.T) {
	is := is.New(t)
	d := []byte(`{"type": "pc", "speed": "1gh"}`)
	pc := PC{}

	err := json.Unmarshal(d, &pc)
	is.NoErr(err)

	is.True(pc.CPU != nil)
	is.Equal(*pc.CPU, frequency(1000))
	is.Equal(pc.GPU, nil)
}
