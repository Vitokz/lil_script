package types

import (
	"time"
)

type BlockRewards struct {
	AddressHash []byte    `db:"address_hash"`
	AddressType string    `db:"address_type"`
	BlockHash   []byte    `db:"block_hash"`
	Reward      *float64  `db:"reward"`
	InsertedAt  time.Time `db:"inserted_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type BlockSecondDegreeRelations struct {
	NephewHash     []byte     `db:"nephew_hash"`
	UncleHash      []byte     `db:"uncle_hash"`
	UncleFetchedAt *time.Time `db:"uncle_fetched_at"`
	Index          *int       `db:"index"`
}

type Block struct {
	Consensus       bool      `db:"consensus"`
	Difficulty      *float64  `db:"difficulty"`
	GasLimit        float64   `db:"gas_limit"`
	GasUsed         float64   `db:"gas_used"`
	Hash            []byte    `db:"hash"`
	MinerHash       []byte    `db:"miner_hash"`
	Nonce           []byte    `db:"nonce"`
	Number          int       `db:"number"`
	ParentHash      []byte    `db:"parent_hash"`
	Size            *int      `db:"size"`
	Timestamp       time.Time `db:"timestamp"`
	TotalDifficulty *float64  `db:"total_difficulty"`
	InsertedAt      time.Time `db:"inserted_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	RefetchNeeded   *bool     `db:"refetch_needed"`
	BaseFeePerGas   *float64  `db:"base_fee_per_gas"`
	IsEmpty         *bool     `db:"is_empty"`
}

type EmissionRewards struct {
	BlockRange int
	Reward     float64
}

type PendingBlockTransaction struct {
	BlockHash                 []byte    `db:"block_hash"`
	InsertedAt                time.Time `db:"inserted_at"`
	UpdatedAT                 time.Time `db:"updated_at"`
	FetchInternalTransactions bool      `db:"fetch_internal_transactions"`
}
