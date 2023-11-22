package data

import (
	"encoding/json"
	"io"
)

func FromJSON(i interface{}, w io.Reader) error {
	e := json.NewDecoder(w)

	return e.Decode(i)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
// https://pkg.go.dev/encoding/json#NewEncoder
func ToJSON(i interface{}, r io.Writer) error {
	e := json.NewEncoder(r)
	return e.Encode(i)
}
