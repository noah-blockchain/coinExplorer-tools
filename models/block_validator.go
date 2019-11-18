package models

import "time"

type BlockValidator struct {
	tableName   struct{}  `sql:"block_validator"`
	BlockID     uint64    `json:"block_id"`
	ValidatorID uint64    `json:"validator_id"`
	CreatedAt   time.Time `json:"created_at"`
	Signed      bool      `json:"signed"`
	Validator   Validator `json:"validator"`
}
