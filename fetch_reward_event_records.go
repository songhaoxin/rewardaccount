package rewardaccount

import (
	"fmt"
	"os"
)

func FetchRewardEventRecords(account string, from string, to string) (bool, []BlockEventRecord)   {
	var blockEventRecords []BlockEventRecord

	var sql string
	if from != "" && to != "" {
		sql = "SELECT block_id FROM data_event WHERE attributes like ? AND event_id='DelegatorReward'" + " AND block_id>='" + from + "' AND block_id<='" + to + "'  ORDER BY block_id DESC;"
	} else {
		sql = "SELECT block_id FROM data_event WHERE attributes like ? AND event_id='DelegatorReward'  ORDER BY block_id DESC;"
	}

	likeValue := "%" + account + "%"

	err := db.Select(&blockEventRecords,sql, likeValue)
	if err != nil {
		fmt.Println("查询DelegatorReward记录失败：, ", err)
		return false, blockEventRecords
	}

	return true, blockEventRecords
}

func Save2File(fileName string,shouldRewardAccounts []ShouldRewardAccount)  {
	if nil == shouldRewardAccounts || 0 == len(shouldRewardAccounts) {
		return
	}

	var outputString string = ""
	for _, rewardAccount := range shouldRewardAccounts {
		rowString := ""
		for _, vs := range rewardAccount.Range {
			rowString = fmt.Sprintf("%s, %d, %d, %d, %d \n",rewardAccount.Account,vs.from,vs.to,vs.from/17280,vs.to/17280)
		}
		outputString = fmt.Sprintf("%s%s",outputString,rowString)
	}
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()
	dstFile.WriteString(outputString)

}