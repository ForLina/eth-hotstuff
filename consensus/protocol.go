// BLS-Upgrade: HotStuff settings
package consensus

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Constants to match up protocol versions and messages

const (
	eth63      = 63
	eth64      = 64
	eth65      = 65
	HotStuff66 = 66
)

var (
	HotStuffProtocol = Protocol{
		Name:     "hotstuff",
		Versions: []uint{HotStuff66},
		Lengths:  map[uint]uint64{HotStuff66: 18},
	}

	// Default: Keep up-to-date with eth/protocol.go
	CliqueProtocol = Protocol{
		Name:     "clique",
		Versions: []uint{eth65, eth64, eth63},
		Lengths:  map[uint]uint64{eth65: 17, eth64: 17, eth63: 17},
	}

	// Default: Keep up-to-date with eth/protocol.go
	EthProtocol = Protocol{
		Name:     "eth",
		Versions: []uint{eth65, eth64, eth63},
		Lengths:  map[uint]uint64{eth65: 17, eth64: 17, eth63: 17},
	}

	NorewardsProtocol = Protocol{
		Name:     "Norewards",
		Versions: []uint{0},
		Lengths:  map[uint]uint64{0: 0},
	}
)

// Protocol defines the protocol of the consensus
type Protocol struct {
	// Official short name of the protocol used during capability negotiation.
	Name string
	// Supported versions of the eth protocol (first is primary).
	Versions []uint
	// Number of implemented message corresponding to different protocol versions.
	Lengths map[uint]uint64
}

// Broadcaster defines the interface to enqueue blocks to fetcher and find peer
type Broadcaster interface {
	// Enqueue add a block into fetcher queue
	Enqueue(id string, block *types.Block)
	// FindPeers retrives peers by addresses
	FindPeers(map[common.Address]bool) map[common.Address]Peer
}

// Peer defines the interface to communicate with peer
type Peer interface {
	// Send sends the message to this peer
	Send(msgcode uint64, data interface{}) error
}
