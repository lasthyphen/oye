// (c) 2020, Alex Willmer, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"errors"

)

// Highlighting modes available

var errUnknownHighlight = errors.New("unknown highlight")

// Highlight mode to apply to displayed logs
type Highlight int

// ToHighlight chooses a highlighting mode

func (h *Highlight) MarshalJSON() ([]byte, error) {
	switch *h {
	case Plain:
		return []byte("\"PLAIN\""), nil
	case Colors:
		return []byte("\"COLORS\""), nil
	default:
		return nil, errUnknownHighlight
	}
}
