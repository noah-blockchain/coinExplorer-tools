package models

import (
	"time"
)

type Coin struct {
	ID                    uint64     `json:"id" sql:",pk"`
	CreationAddressID     *uint64    `json:"creation_address_id"`
	CreationTransactionID *uint64    `json:"creation_transaction_id"`
	Crr                   uint64     `json:"crr"`
	Volume                string     `json:"volume"          sql:"type:numeric(70)"`
	ReserveBalance        string     `json:"reserve_balance" sql:"type:numeric(70)"`
	Price                 string     `json:"price"           sql:"type:numeric(100)"`
	Capitalization        string     `json:"capitalization"  sql:"type:numeric(100)"`
	Delegated             uint64     `json:"delegated"`
	StartVolume           string     `json:"start_volume"          sql:"type:numeric(70)"`
	StartReserveBalance   string     `json:"start_reserve_balance" sql:"type:numeric(70)"`
	StartPrice            string     `json:"start_price"           sql:"type:numeric(100)"`
	Name                  string     `json:"name"            sql:"type:varchar(255)"`
	Symbol                string     `json:"symbol"          sql:"type:varchar(20)"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at"      pg:",soft_delete"`

	Address string `json:"address" sql:"-"`
}
