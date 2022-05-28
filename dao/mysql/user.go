package mysql

import (
	"app/models"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func InsertUser(user_id int64, address, account, email, password, privatekey string) (err error) {
	sqlStr := `insert into tab_user(user_id,address,account,
		email,password,privatekey,time) values (
		:user_id,:address,:account,:email,:password,:privatekey,:time
		);`
	//执行 SQL 语句
	ret, err := db.NamedExec(sqlStr,
		map[string]interface{}{
			"user_id":    user_id,
			"address":    address,
			"account":    account,
			"email":      email,
			"password":   password,
			"privatekey": privatekey,
			"time":       time.Now(),
		})
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		zap.L().Warn("insert failed", zap.Error(err))
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		zap.L().Warn("get lastinsert ID failed", zap.Error(err))
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	zap.L().Info("insert success", zap.Int64("theID", theID))
	return nil
}

func GetUserByAddress(address string) (user *models.DBUser, err error) {
	sqlStr := `select * from tab_user where address = ?;`
	//执行 SQL 语句
	var u models.DBUser
	err = db.Get(&u, sqlStr, address)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	// fmt.Printf("User:%v\n", u)
	return &u, nil
}
