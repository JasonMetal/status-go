package walletevent

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// EventType type for event types.
type EventType string

// EventType prefix to be used for internal events.
// These events are not forwarded to the client, they are only used
// within status-go.
const InternalEventTypePrefix = "INT-"

func (t EventType) IsInternal() bool {
	return strings.HasPrefix(string(t), InternalEventTypePrefix)
}

// Event is a type for transfer events.
type Event struct {
	Type        EventType        `json:"type"`
	BlockNumber *big.Int         `json:"blockNumber"`
	Accounts    []common.Address `json:"accounts"`
	Message     string           `json:"message"`
	At          int64            `json:"at"`
	ChainID     uint64           `json:"chainId"`
	RequestID   *int             `json:"requestId,omitempty"`
	// For Internal events only, not serialized
	EventParams interface{}
}
