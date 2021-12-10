package rewardaccount

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql","删除地址及帐号密码")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func CloseDBConnection()  {
	db.Close()
}

