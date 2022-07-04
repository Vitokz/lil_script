package staking

import (
	"context"
	"fmt"
	"github.com/Vitokz/lil_script/db"
	dbTypes "github.com/Vitokz/lil_script/db/types"
	"github.com/Vitokz/lil_script/worker/x"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	web3 "github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"
	"math/big"
	"time"
)

// msgs types
var (
	msgDelegate        = sdkTypes.MsgTypeURL(&stakingTypes.MsgDelegate{})
	msgEditValidator   = sdkTypes.MsgTypeURL(&stakingTypes.MsgEditValidator{})
	msgBeginRedelegate = sdkTypes.MsgTypeURL(&stakingTypes.MsgBeginRedelegate{})
	msgUndelegate      = sdkTypes.MsgTypeURL(&stakingTypes.MsgUndelegate{})
	msgCreateValidator = sdkTypes.MsgTypeURL(&stakingTypes.MsgCreateValidator{})
)

type Staking struct {
}

func NewStaking() x.Module {
	staking := Staking{}

	return &staking
}

func (s *Staking) GetHandler() x.HandlerI {
	return s.getHandler()
}

func (s *Staking) Msgs() []string {
	return []string{
		msgDelegate,
		msgEditValidator,
		msgBeginRedelegate,
		msgUndelegate,
		msgCreateValidator,
	}
}

func (s *Staking) GetName() string {
	return stakingTypes.ModuleName
}

// handlers ------------------------------------------------

func (s *Staking) getHandler() x.HandlerI {
	return func(ctx context.Context, web3 *web3.Client, db *db.DB, logger log.Logger, msg sdkTypes.Msg) error {
		switch msg := msg.(type) {
		case *stakingTypes.MsgCreateValidator:
			return s.msgCreateValidator()
		case *stakingTypes.MsgEditValidator:
			return s.msgEditValidator()
		case *stakingTypes.MsgDelegate:
			return s.msgDelegate(ctx, web3, db, logger, msg)
		case *stakingTypes.MsgBeginRedelegate:
			return s.msgRedelegate()
		case *stakingTypes.MsgUndelegate:
			return s.msgUndelegate()
		default:
			return errors.Errorf("unrecognized %s message type: %T", s.GetName(), msg)
		}
	}
}

func (s *Staking) msgDelegate(ctx context.Context, web3 *web3.Client, db *db.DB, logger log.Logger, msg *stakingTypes.MsgDelegate) error {
	logger.Info(fmt.Sprintf("start handle delegate tx: valAddr(%s) delAddr(%s)", msg.ValidatorAddress, msg.DelegatorAddress))

	delAddr := msg.DelegatorAddress
	delAccAddr, err := sdkTypes.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return err
	}
	delAccAddrHex, err := x.Bech32ToHex(delAddr)
	if err != nil {
		return err
	}

	height := ctx.Value("height")

	balance, err := web3.BalanceAt(ctx, common.HexToAddress(delAccAddrHex), big.NewInt(height.(int64)))
	if err != nil {
		logger.Debug(fmt.Sprintf("failed to get balances to address: %s", delAccAddrHex))
		return err
	}

	balanceStr := balance.String()

	coinBalance := dbTypes.AddressCoinBalances{
		AddressHash: delAccAddr.Bytes(),
		BlockNumber: int(height.(int64)),
	}
	var existAddr int
	query := `SELECT count(*) FROM addresses a WHERE a.hash = $1`

	err = db.DB.GetContext(ctx, &existAddr,
		query, coinBalance.AddressHash,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get address data with addr: %s", delAccAddrHex))
		return err
	}

	if existAddr == 0 {
		nonce, err := web3.NonceAt(ctx, common.HexToAddress(delAccAddrHex), big.NewInt(int64(coinBalance.BlockNumber)))
		if err != nil {
			return err
		}

		query = `INSERT INTO addresses
				  (fetched_coin_balance, fetched_coin_balance_block_number, hash, contract_code, inserted_at, updated_at, nonce,
				   decompiled, verified, gas_used, transactions_count, token_transfers_count)
			      VALUES(cast($1 AS NUMERIC), $2, $3, NUll, $4, $5, $6,
			             false, false, NULL, NULL, NULL);
	`
		_, err = db.DB.ExecContext(ctx, query,
			balanceStr, coinBalance.BlockNumber, coinBalance.AddressHash, time.Now(), time.Now(), nonce,
		)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to insert new address in db: %s", delAccAddrHex))
			return err
		}
		logger.Info(fmt.Sprintf("Add new addr in database: %s", delAccAddrHex))
	}

	//  insert user balance
	query = `INSERT INTO address_coin_balances
			  (address_hash, block_number, value, value_fetched_at, inserted_at, updated_at)
			  VALUES($1, $2, cast($3 AS NUMERIC), NULL, $4, $5) ON conflict DO NOTHING;
	`
	_, err = db.DB.ExecContext(ctx, query,
		coinBalance.AddressHash, coinBalance.BlockNumber, balanceStr,
		time.Now(), time.Now(),
	)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to insert address_coin_balance in msgDelegate: %s", delAddr))
		return err
	}

	return nil
}

func (s *Staking) msgUndelegate() error {

	return nil
}

func (s *Staking) msgEditValidator() error {

	return nil
}

func (s *Staking) msgCreateValidator() error {

	return nil
}

func (s *Staking) msgRedelegate() error {

	return nil
}
