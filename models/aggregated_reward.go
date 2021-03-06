package models

import "time"

type AggregatedReward struct {
	FromBlockID uint64    `json:"from_block_id" sql:",pk"`
	ToBlockID   uint64    `json:"to_block_id"`
	AddressID   uint64    `json:"address_id"    sql:",pk"`
	ValidatorID uint64    `json:"validator_id"  sql:",pk"`
	Role        string    `json:"role"          sql:",pk"`
	Amount      string    `json:"amount"        sql:"type:numeric(70)"`
	TimeID      time.Time `json:"time_id"`
	FromBlock   *Block     //Relation has one to Blocks
	ToBlock     *Block     //Relation has one to Blocks
	Address     *Address   //Relation has one to Addresses
	Validator   *Validator //Relation has one to Validators
}
