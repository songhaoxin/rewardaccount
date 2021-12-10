package rewardaccount

import "strconv"


func VerifyBlocks(account string, blockEventRecords []BlockEventRecord, blockInterval int64) (bool, ShouldRewardAccount) {
	ret := ShouldRewardAccount{}
	lenBlock := len(blockEventRecords)
	if lenBlock < 2 {
		return false, ret
	}
	ret.Account = account

	for index := 0; index < lenBlock-1; index++ {
		preBlockNumber, err := strconv.ParseInt(blockEventRecords[index].BlockId, 10, 64)
		if err != nil {
			continue
		}
		nextBlockNumber, err := strconv.ParseInt(blockEventRecords[index+1].BlockId, 10, 64)
		if err != nil {
			continue
		}

		if preBlockNumber - nextBlockNumber >= blockInterval {
			ret.Range = append(ret.Range, RewardRange{
				from: nextBlockNumber,
				to:   preBlockNumber,
			})
		}
	}
	return true, ret
}
