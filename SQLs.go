package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
)

const (
	dbname = "user"
	sqlcon = "root:Qin_123456@tcp(127.0.0.1:3306)/" + dbname + "?charset=utf8"
)

var (
	db  *sql.DB
	err error
)

type Record []map[string]string

func init() {
	//打开数据库

	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err = sql.Open("mysql", sqlcon)

	if err != nil {

		fmt.Println(err)

	}
}

func Execute(q string) error {
	_, err = db.Exec(q)
	return err
}

func GetRecord(query string) Record {
	//fmt.Println(query)
	rows, _ := db.Query(query)
	//返回所有列
	//fmt.Println(rows)
	cols, _ := rows.Columns()

	//这里表示一行所有列的值，用[]byte表示
	result := Record{}
	vals := make([]string, len(cols))
	//fmt.Println(len(vals))

	//这里表示一行填充数据

	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里

	for k, _ := range vals {

		scans[k] = &vals[k]

	}

	//	i := 0

	for rows.Next() {

		//填充数据

		rows.Scan(scans...)

		//每行数据

		row := make(map[string]string)

		//把vals中的数据复制到row中

		for k, v := range vals {

			key := cols[k]

			//这里把[]byte数据转成string

			row[key] = v

		}

		//放入结果集
		result = append(result, row)
	}
	return result
}
