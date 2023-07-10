package scores

import (
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
)

// when adding a new req, update CalcSlots()

// ScoreReq is an interface for pairing requirement scoring
type ScoreReq interface {
	// Score() calculates a provider's score according to the requirement
	Score(stakeEntry epochstoragetypes.StakeEntry, weight uint64) uint64
	// GetName returns the unique name of the ScoreReq implementation
	GetName() string
}
