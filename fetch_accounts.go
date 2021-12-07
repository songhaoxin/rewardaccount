package rewardaccount

import "fmt"

func FetchAccounts() (bool, []Accounts) {
	var accounts  []Accounts
	sql := "select id,address from data_account"
	err := db.Select(&accounts,sql)
	if err != nil {
		fmt.Println("查询帐户信息时出错：, ", err)
		return false, accounts
	}
	return true, accounts
}
