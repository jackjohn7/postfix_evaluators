package symbols

import (
	"strings"
)

type Text struct {
	raw      string
	split    []string
	position int
}

func New(raw string) *Text {
	return &Text{
		raw:      raw,
		split:    strings.Split(raw, " "),
		position: 0,
	}
}

func (t *Text) ReadNext() *string {
	if len(t.split) <= t.position {
		return nil
	}
	res := t.split[t.position]
	t.position += 1
	return &res
}
