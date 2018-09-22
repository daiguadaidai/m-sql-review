package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	"fmt"
	"github.com/daiguadaidai/m-sql-review/config"
)

type Instance struct {
	DBconfig *config.DBConfig
	DB *sql.DB
}

/* 新建一个数据库执行器
Params:
    _dbConfig: 数据库配置
 */
func NewInstance(_dbConfig *config.DBConfig) *Instance {
	executor := &Instance {DBconfig: _dbConfig}

	return executor
}

// 打开数据库连接
func (this *Instance) OpenDB() error {
	var err error

	this.DB, err = sql.Open("mysql", this.DBconfig.GetDataSource())
	if err != nil { // 打开数据库失败
		errMSG := fmt.Sprintf("打开数据库链接失败[%v:%v]",
			this.DBconfig.Host, this.DBconfig.Port)
		return errors.New(errMSG)
	}

	return err
}

// 关闭数据库链接
func (this *Instance) CloseDB() error {
	var err error

	if this.DB != nil {
		err = this.DB.Close()
	}

	return err
}