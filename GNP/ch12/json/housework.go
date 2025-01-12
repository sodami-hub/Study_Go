package json

import (
	"encoding/json"
	"io"

	"housework"
)

func Load(r io.Reader) ([]*housework.Chore, error) {
	var chores []*housework.Chore

	return chores, json.NewDecoder(r).Decode(&chores)
}

func Flush(w io.Writer, chores []*housework.Chore) error {
	return json.NewEncoder(w).Encode(chores)
}
