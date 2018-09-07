package main

import (
	"fmt"
	"github.com/daiguadaidai/m-sql-review/service"
)

func main() {

	/*
	sql = `
        CREATE DATABASE IF NOT EXISTS yourdbname DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
    `
	*/
	sql := `
		alter table t1 
            add id varchar(20) not null default '' comment 'ni hao '
            ;
    `

    code, msg := service.Start(sql)
    fmt.Println(code, msg)

}
