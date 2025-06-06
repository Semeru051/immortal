package websocket

import (
	"sync"

	"github.com/starrysilk/immortal/types/filter"
)

type clientState struct {
	challenge string
	pubkey    *string
	isKnown   *bool
	subs      map[string]filter.Filter
	*sync.RWMutex
}
