package mysql

import (
	"app/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *setting.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	// 创建表
	err = createTableUser()
	return
}

// 小写的 db变量 不能对外暴露, 只能在包内部使用
// 所以在外面 关闭 db 连接的时候, 就不能使用 defer db.Close()
// 可以封装一个函数,可以在外面调用, 关闭 db 连接
func Close() {
	_ = db.Close()
}

// DROP TABLE IF EXISTS `tab_user`;
func createTableUser() error {
	sqlStr := `CREATE TABLE IF NOT EXISTS tab_user  (
	id int(0) NOT NULL AUTO_INCREMENT,
	user_id bigint(0) NOT NULL COMMENT '随机生成的用户id',
	address varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户地址',
	account varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户账号',
	email varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户邮箱',
	password varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户加密后的密码',
	privateKey varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '私钥',
	time datetime(0) NULL DEFAULT NULL COMMENT '注册时间',
	PRIMARY KEY (id, user_id, address, account, email) USING BTREE
	) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;`

	_, err := db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}
