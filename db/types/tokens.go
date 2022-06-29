package types

import "time"

type TokenInstance struct {
	TokenID                  int         `db:"token_id"`
	TokenContractAddressHash []byte      `db:"token_contract_address_hash"`
	Metadata                 interface{} `db:"metadata"`
	InsertedAt               time.Time   `db:"inserted_at"`
	UpdatedAt                time.Time   `db:"updated_at"`
	Error                    *string     `db:"error"`
}

type TokenTransfer struct {
	TransactionHash          []byte    `db:"transaction_hash"`
	LogIndex                 int       `db:"log_index"`
	FromAddressHash          []byte    `db:"from_address_hash"`
	ToAddressHash            []byte    `db:"to_address_hash"`
	Amount                   *float64  `db:"amount"`
	TokenID                  *float64  `db:"token_id"`
	TokenContractAddressHash []byte    `db:"token_contract_address_hash"`
	InsertedAt               time.Time `db:"inserted_at"`
	UpdatedAt                time.Time `db:"updated_at"`
	BlockNumber              int       `db:"block_number"`
	BlockHash                []byte    `db:"block_hash"`
	Amounts                  []float64 `db:"amounts"`
	TokenIDs                 []float64 `db:"token_i_ds"`
}

type Tokens struct {
	Name                *string   `db:"name"`
	Symbol              *string   `db:"symbol"`
	TotalSupply         *float64  `db:"total_supply"`
	Decimals            *float64  `db:"decimals"`
	Type                string    `db:"type"`
	Cataloged           bool      `db:"cataloged"`
	ContractAddressHash []byte    `db:"contract_address_hash"`
	InsertedAt          time.Time `db:"inserted_at"`
	UpdatedAt           time.Time `db:"updated_at"`
	HolderCount         *int      `db:"holder_count"`
	Bridged             *bool     `db:"bridged"`
	SkipMetadata        *bool     `db:"skip_metadata"`
}
