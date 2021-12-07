package rewardaccount

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql","deeperreader:pyV38DGJZz3Df7lwKjahw@tcp(mainnet-mysql-public.crzdyh02estr.us-east-1.rds.amazonaws.com:3306)/polkascan")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

