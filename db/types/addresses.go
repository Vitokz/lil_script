package types

import "time"

type AddressCoinBalances struct {
	AddressHash    []byte     `db:"address_hash"`
	BlockNumber    int        `db:"block_number"`
	Value          float64    `db:"value"`
	ValueFetchedAt *time.Time `db:"value_fetched_at"`
	InsertedAt     time.Time  `db:"inserted_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
}

type AddressCoinBalancesDaily struct {
	AddressHash []byte      `db:"address_hash"`
	Day         interface{} `db:"day"` //TODO
	Value       float64     `db:"value"`
	InsertedAt  time.Time   `db:"inserted_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}

type AddressCurrentTokenBalances struct {
	ID                       int        `db:"id"`
	AddressHash              []byte     `db:"address_hash"`
	BlockNumber              int        `db:"block_number"`
	TokenContractAddressHash []byte     `db:"token_contract_address_hash"`
	Value                    *float64   `db:"value"`
	ValueFetchedAt           *time.Time `db:"value_fetched_at"`
	InsertedAt               time.Time  `db:"inserted_at"`
	UpdatedAt                time.Time  `db:"updated_at"`
	OldValue                 *float64   `db:"old_value"`
	TokenID                  *float64   `db:"token_id"`
	TokenType                *string    `db:"token_type"`
}

type AddressNames struct {
	AddressHash []byte       `db:"address_hash"`
	Name        string       `db:"name"`
	Primary     bool         `db:"primary"`
	InsertedAt  time.Time    `db:"inserted_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	Metadata    *interface{} `db:"metadata"`
}

type AddressTokenBalance struct {
	ID                       int        `db:"id"`
	AddressHash              []byte     `db:"address_hash"`
	BlockNumber              int        `db:"block_number"`
	TokenContractAddressHash []byte     `db:"token_contract_address_hash"`
	Value                    *float64   `db:"value"`
	ValueFetchedAt           *time.Time `db:"value_fetched_at"`
	InsertedAt               time.Time  `db:"inserted_at"`
	UpdatedAt                time.Time  `db:"updated_at"`
	TokenID                  *float64   `db:"token_id"`
	TokenType                *string    `db:"token_type"`
}

type Addresses struct {
	FetchedCoinBalance            *float64  `db:"fetched_coin_balance"`
	FetchedCoinBalanceBlockNumber *int      `db:"fetched_coin_balance_block_number"`
	Hash                          []byte    `db:"hash"`
	ContractCode                  *[]byte   `db:"contract_code"`
	InsertedAt                    time.Time `db:"inserted_at"`
	UpdatedAt                     time.Time `db:"updated_at"`
	Nonce                         *int      `db:"nonce"`
	Decompiled                    *bool     `db:"decompiled"`
	Verified                      *bool     `db:"verified"`
	GasUsed                       *int      `db:"gas_used"`
	TransactionsCount             *int      `db:"transactions_count"`
	TokenTransfersCount           *int      `db:"token_transfers_count"`
}

type Administrators struct {
	ID         int       `db:"id"`
	Role       string    `db:"role"`
	UserID     int       `db:"user_id"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type UserContacts struct {
	ID         int       `db:"id"`
	Email      string    `db:"email"`
	UserID     int       `db:"user_id"`
	Primary    bool      `db:"primary"`
	Verified   bool      `db:"verified"`
	InsertedAt time.Time `db:"inserted_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type Users struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	InsertedAt   time.Time `db:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
