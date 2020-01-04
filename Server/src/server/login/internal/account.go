package internal

//import (
//	"fmt"
//"server/mysql"
//)

type Account struct {
	PlayerID   int32  `gorm:"primary_key"`
	PlayerName string `gorm:"not null;unique"`
}

func getAccountByAccountID(accountID string) *Account {

	var account Account
	/*
		db := mysql.MysqlDB()
		err := db.Where("Account = ?", accountID). Limit(1).Find(&account).Error
		if nil != err {
			fmt.Println(err)
			return nil
		}
		fmt.Println("password:",account.Password)
	*/
	return &account
}

func creatAccountByAccountIDAndPassword(accountID string, password string) *Account {
	var account Account
	/*
		db := mysql.MysqlDB()
		var account = Account{AccountID:accountID,Password:password}
		err := db.Create(&account).Error
		if nil != err {
			return  nil
		}
	*/

	return &account
}

func accountIsExist(playerId int32) bool {

	_, ok := playerMap[playerId]

	return ok

}

func accountNameIsExist(playerId int32, playerName string) bool {

	player, ok := playerMap[playerId]

	if ok {
		return (player.PlayerName == playerName)
	}

	return false

}

func createAccount(playerId int32, playerName string) *Account {
	account := &Account{PlayerID: playerId, PlayerName: playerName}
	playerMap[playerId] = account
	return account
}

func delAccount(playerId int32) {

	playerMap[playerId] = nil

}
