package main

import (
	"flag"
	"fmt"
	"rewardaccount"
	"sync"
)

var wg sync.WaitGroup
var jobsChan = make(chan rewardaccount.ShouldRewardAccount,30)
var shouldRewardAccounts = make([]rewardaccount.ShouldRewardAccount,0)


func main() {
	var blockFrom string
	var blockTo string
	var outputFileName string
	var tasksPerGr int
	flag.StringVar(&blockFrom,"s","","扫描范围的起始区块号")
	flag.StringVar(&blockTo,"e","","扫描范围的结束区块号")
	flag.StringVar(&outputFileName,"f","should-reward-accounts.txt","保存需要补偿收益帐户信息的文件名")
	flag.IntVar(&tasksPerGr,"n",100,"每一个协程处理的任务数")
	flag.Parse()
	//fmt.Printf("blockFrom=%v blockTo=%v outputFileName=%v",blockFrom,blockTo,outputFileName)

	flag,accounts := rewardaccount.FetchAccounts()
	if flag == false{
		return
	}

	go handleRewardAccount()

	tasks := tasksPerGr;l := len(accounts);from := 0;to := 0
	if l <= 0 {
		fmt.Println("没有需要处理的帐户信息")
		return
	}
	for {
		to = from + tasks
		if to > l { // 超界了
			to = l
		}

		//fmt.Println("from:",from)
		//fmt.Println("to:",to)
		//fmt.Println(accounts[from:to])
		//fmt.Println("-------------------------")
		wg.Add(1)
		go handleAccounts(accounts[from:to],blockFrom,blockTo)
		from = to
		if from >= l {
			break
		}
	}
	wg.Wait()
	close(jobsChan)
	rewardaccount.Save2File(outputFileName,shouldRewardAccounts)

}


func handleAccounts(accounts []rewardaccount.Accounts, from string, to string)  {
	defer wg.Done()
	for _,account := range accounts {
		fmt.Println("正在处理帐户：",account.AccountId)
		flag,records := rewardaccount.FetchRewardEventRecords(account.AccountId,from,to)
		if flag == false {
			continue
		}

		flag,shouldRewardAccount :=  rewardaccount.VerifyBlocks(account.Address,records,17280*2)
		if flag == true {
			jobsChan <- shouldRewardAccount
			fmt.Println("需补偿收益帐户：",shouldRewardAccount.Account)
		}
	}
}

func handleRewardAccount()  {
	for {
		select {
			case shouldRewardAccount := <- jobsChan: {
					shouldRewardAccounts = append(shouldRewardAccounts, shouldRewardAccount)
					//fmt.Println("增加了一个帐户信息了")
					break
				}
			}
	}
}
