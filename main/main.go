package main

import (
	"fmt"
	"rewardaccount"
)

func main() {

	shouldRewardAccounts := make([]rewardaccount.ShouldRewardAccount,0)

	flag,accounts := rewardaccount.FetchAccounts()
	if flag == false{
		return
	}
	for _,account := range accounts {
		fmt.Println(account.AccountId)
		flag,records := rewardaccount.FetchRewardEventRecords(account.AccountId)
		if flag == false {
			continue
		}
		for _, record := range records {
			fmt.Println(record.BlockId)
		}
		flag,shouldRewardAccount :=  rewardaccount.VerifyBlocks(account.Address,records,17280*2)
		if flag == true {
			shouldRewardAccounts = append(shouldRewardAccounts, shouldRewardAccount)
		}
	}

	fmt.Println(shouldRewardAccounts)
	//accountsCount := len(accounts)
	//fmt.Println("帐户的数量：", accountsCount)

}
