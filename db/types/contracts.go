package types

import "time"

type BridgedToken struct {
	ForeignChainID                  float64   `db:"foreign_chain_id"`
	ForeignTokenContractAddressHash []byte    `db:"foreign_token_contract_address_hash"`
	HomeTokenContractAddressHash    []byte    `db:"home_token_contract_address_hash"`
	InsertedAt                      time.Time `db:"inserted_at"`
	UpdatedAt                       time.Time `db:"updated_at"`
	CustomMetadata                  *string   `db:"custom_metadata"`
	Type                            *string   `db:"type"`
	ExchangeRate                    *float64  `db:"exchange_rate"`
	LpToken                         *bool     `db:"lp_token"`
	CustomCap                       *float64  `db:"custom_cap"`
}

type ContractMethod struct {
	ID         int         `db:"id"`
	Identifier int         `db:"identifier"`
	Abi        interface{} `db:"abi"`
	Type       string      `db:"type"`
	InsertedAt time.Time   `db:"inserted_at"`
	UpdatedAt  time.Time   `db:"updated_at"`
}

type ContractVerificationStatus struct {
	Uid         string    `db:"uid"`
	Status      int       `db:"status"`
	AddressHash []byte    `db:"address_hash"`
	InsertedAt  time.Time `db:"inserted_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type DecompiledSmartContract struct {
	ID                   int       `db:"id"`
	DecompilerVersion    string    `db:"decompiler_version"`
	DecompiledSourceCode string    `db:"decompiled_source_code"`
	AddressesHash        []byte    `db:"addresses_hash"`
	InsertedAt           time.Time `db:"inserted_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}

type SmartContract struct {
	ID                   int           `db:"id"`
	Name                 string        `db:"name"`
	CompilerVersion      string        `db:"compiler_version"`
	Optimization         string        `db:"optimization"`
	ContractSourceCode   string        `db:"contract_source_code"`
	Abi                  interface{}   `db:"abi"`
	AddressHash          []byte        `db:"address_hash"`
	InsertedAt           time.Time     `db:"inserted_at"`
	UpdatedAt            time.Time     `db:"updated_at"`
	ConstructorArguments *string       `db:"constructor_arguments"`
	OptimizationRuns     *int          `db:"optimization_runs"`
	EvmVersion           *string       `db:"evm_version"`
	ExternalLibraries    []interface{} `db:"external_libraries"`
	VerifiedViaSourcify  *bool         `db:"verified_via_sourcify"`
	IsVyperContract      *bool         `db:"is_vyper_contract"`
	PartiallyVerified    *bool         `db:"partially_verified"`
	FilePath             *string       `db:"file_path"`
	IsChangedBytecode    *bool         `db:"is_changed_bytecode"`
	BytecodeCheckedAt    *time.Time    `db:"bytecode_checked_at"`
}

type SmartContractsAdditionalSource struct {
	ID                 int       `db:"id"`
	FileName           string    `db:"file_name"`
	ContractSourceCode string    `db:"contract_source_code"`
	AddressHash        []byte    `db:"address_hash"`
	InsertedAt         time.Time `db:"inserted_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
