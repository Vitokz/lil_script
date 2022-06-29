package types

import "time"

type StakingPools struct {
	ID                               int       `db:"id"`
	AreDelegatorsBanned              bool      `db:"are_delegators_banned"`
	BannedDelegatorsUntil            *int      `db:"banned_delegators_until"`
	BannedUntil                      *string   `db:"banned_until"`
	BanReason                        *string   `db:"ban_reason"`
	DelegatorsCount                  *int      `db:"delegators_count"`
	IsActive                         bool      `db:"is_active"`
	IsBanned                         bool      `db:"is_banned"`
	IsDeleted                        bool      `db:"is_deleted"`
	IsUnremovable                    bool      `db:"is_unremovable"`
	IsValidator                      bool      `db:"is_validator"`
	LikeLihood                       *float64  `db:"like_lihood"`
	MiningAddressHash                []byte    `db:"mining_address_hash"`
	SelfStakedAmount                 *float64  `db:"self_staked_amount"`
	SnapshottedSelfStakedAmount      *float64  `db:"snapshotted_self_staked_amount"`
	SnapshottedTotalStakedAmount     *float64  `db:"snapshotted_total_staked_amount"`
	SnapshottedValidatorRewardRation *float64  `db:"snapshotted_validator_reward_ration"`
	StakesRatio                      *float64  `db:"stakes_ratio"`
	StakingAddressHash               []byte    `db:"staking_address_hash"`
	TotalStakedAmount                *float64  `db:"total_staked_amount"`
	ValidatorRewardPercent           *float64  `db:"validator_reward_percent"`
	ValidatorRewardRatio             *float64  `db:"validator_reward_ratio"`
	WasBannedCount                   *int      `db:"was_banned_count"`
	WasValidatorCount                *int      `db:"was_validator_count"`
	InsertedAt                       time.Time `db:"inserted_at"`
	UpdatedAt                        time.Time `db:"updated_at"`
	Name                             *string   `db:"name"`
	Description                      *string   `db:"description"`
}

type StakingPoolsDelegators struct {
	ID                        int
	AddressHash               []byte
	IsActive                  bool
	IsDeleted                 bool
	MaxOrderedWithdrawAllowed *float64
	MaxWithdrawAllowed        *float64
	OrderedWithdraw           *float64
	OrderedWithdrawEpoch      *int
	RewardRatio               *float64
	SnapshottedRewardRation   *float64
	SnapshottedStakeAmount    *float64
	StakeAmount               *float64
	StakingAddressHash        []byte
	InsertedAt                time.Time
	UpdatedAt                 time.Time
}
