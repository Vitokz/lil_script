package types

import "time"

type InternalTransactions struct {
	CallType                   *string   `db:"call_type"`
	CreatedContractCode        []byte    `db:"created_contract_code"`
	Error                      *string   `db:"error"`
	Gas                        *float64  `db:"gas"`
	GasUsed                    *float64  `db:"gas_used"`
	Index                      int       `db:"index"`
	Init                       []byte    `db:"init"`
	Input                      []byte    `db:"input"`
	Output                     []byte    `db:"output"`
	TraceAddress               int       `db:"trace_address"`
	Type                       string    `db:"type"`
	Value                      float64   `db:"value"`
	InsertedAt                 time.Time `db:"inserted_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
	CreatedContractAddressHash []byte    `db:"created_contract_address_hash"`
	FromAddressHash            []byte    `db:"from_address_hash"`
	ToAddressHash              []byte    `db:"to_address_hash"`
	TransactionHash            []byte    `db:"transaction_hash"`
	BlockNumber                *int      `db:"block_number"`
	TransactionIndex           *int      `db:"transaction_index"`
	BlockHash                  []byte    `db:"block_hash"`
	BlockIndex                 int       `db:"block_index"`
}

type Log struct {
	Data            []byte
	Index           int
	Type            *string
	FirstTopic      *string
	SecondTopic     *string
	ThirdTopic      *string
	FourthTopic     *string
	InsertedAt      time.Time
	UpdatedAt       time.Time
	AddressHash     *[]byte
	TransactionHash []byte
	BlockHash       []byte
	BlockNumber     *int
}

type TransactionFork struct {
	Hash       []byte    `db:"hash"`
	Index      int       `db:"index"`
	UncleHash  []byte    `db:"uncle_hash"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type TransactionStats struct {
	ID                   int         `db:"id"`
	Date                 interface{} `db:"date"`
	NumberOfTransactions *int        `db:"number_of_transactions"`
	GasUsed              *float64    `db:"gas_used"`
	TotalFee             *float64    `db:"total_fee"`
}

type Transaction struct {
	CumulativeGasUsed            *float64   `db:"cumulative_gas_used"`
	Error                        *string    `db:"error,omitempty"`
	Gas                          float64    `db:"gas,omitempty"`
	GasPrice                     float64    `db:"gas_price,omitempty"`
	GasUsed                      float64    `db:"gas_used,omitempty"`
	Hash                         []byte     `db:"hash,omitempty"`
	Index                        uint       `db:"index,omitempty"`
	Input                        []byte     `db:"input,omitempty"`
	Nonce                        uint64     `db:"nonce,omitempty"`
	R                            *float64   `db:"r,omitempty"`
	S                            *float64   `db:"s,omitempty"`
	Status                       *uint64    `db:"status,omitempty"`
	V                            *float64   `db:"v,omitempty"`
	Value                        float64    `db:"value,omitempty"`
	InsertedAt                   time.Time  `db:"inserted_at"`
	UpdatedAt                    *time.Time `db:"updated_at,omitempty"`
	BlockHash                    []byte     `db:"block_hash,omitempty"`
	BlockNumber                  int64      `db:"block_number,omitempty"`
	FromAddressHash              []byte     `db:"from_address_hash,omitempty"`
	ToAddressHash                []byte     `db:"to_address_hash,omitempty"`
	CreatedContractAddressHash   []byte     `db:"created_contract_address_hash,omitempty"`
	CreatedContractCodeIndexedAt *time.Time `db:"created_contract_code_indexed_at,omitempty"`
	EarliestProcessingStart      *time.Time `db:"earliest_processing_start,omitempty"`
	OldBlockHash                 []byte     `db:"old_block_hash,omitempty"`
	RevertReason                 *string    `db:"revert_reason,omitempty"`
	MaxPriorityFeePerGas         *float64   `db:"max_priority_fee_per_gas,omitempty"`
	MaxFeePerGas                 *float64   `db:"max_fee_per_gas,omitempty"`
	Type                         *uint8     `db:"eth_tx_type,omitempty"`
	HasErrorInInternalTxs        *bool      `db:"has_error_in_internal_txs,omitempty"`
}
