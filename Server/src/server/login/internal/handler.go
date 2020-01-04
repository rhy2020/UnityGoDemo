package internal

import (
	"reflect"
	"server/game"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handleAuth)
}
func handleAuth(args []interface{}) {
	m := args[0].(*msg.Login)
	a := args[1].(gate.Agent)

	if accountIsExist(m.PlayerID) {
		a.WriteMsg(&msg.LoginFaild{RspHead: &msg.RspHead{ErrorId: 1, ErrorString: "玩家id已存在!"}})
		return
	}
	if accountNameIsExist(m.PlayerID, m.Name) {
		a.WriteMsg(&msg.LoginFaild{RspHead: &msg.RspHead{ErrorId: 2, ErrorString: "玩家名字已存在!"}})
		return
	}

	a.SetUserData(createAccount(m.PlayerID, m.Name))

	//game.ChanRPC.Go("CreatePlayer", a, m.PlayerID, m.Name)
	game.ChanRPC.Go("UserLogin", a, m.PlayerID, m.Name)

	a.WriteMsg(&msg.LoginSuccessfull{})

	/*
		if len(m.Account) < 2 || len(m.Account) > 12 {
			a.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_AccIDInvalid})
			return
		}

		account := getAccountByAccountID(m.Account)

		data := []byte(m.Name)
		var hash = md5.Sum(data)
		password := hex.EncodeToString(hash[:])
		if nil == account {
			//not having this account,creat account
			newAccount := creatAccountByAccountIDAndPassword(m.Account, password)
			if nil != newAccount {
				game.ChanRPC.Go("CreatePlayer", a, newAccount.PlayerID)
				game.ChanRPC.Go("UserLogin", a, newAccount.PlayerID)
			} else {
				log.Debug("create account error ", m.Account)
				a.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_InnerError})
			}

		} else {
			// match password
			if password == account.Password {
				game.ChanRPC.Go("UserLogin", a, account.PlayerID)
			} else {
				a.WriteMsg(&msg.LoginFaild{Code: msg.LoginFaild_AccountOrPasswardNotMatch})
			}
		}
	*/
}
