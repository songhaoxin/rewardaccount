package rewardaccount

type Accounts struct {
	AccountId string `db:"id"`
	Address string `db:"address"`
}

type BlockEventRecord struct {
	BlockId string `db:"block_id"`
}


type RewardRange struct {
	from int64
	to   int64
}

type ShouldRewardAccount struct {
	Account string
	Range   []RewardRange
}
