package rewardaccount

import "fmt"

func FetchRewardEventRecords(account string) (bool, []BlockEventRecord)   {
	var blockEventRecords []BlockEventRecord

	sql := "SELECT block_id FROM data_event WHERE attributes like ? AND event_id='DelegatorReward'  ORDER BY block_id DESC;"
	likeValue := "%" + account + "%"

	err := db.Select(&blockEventRecords,sql, likeValue)
	if err != nil {
		fmt.Println("查询DelegatorReward记录失败：, ", err)
		return false, blockEventRecords
	}

	return true, blockEventRecords
}