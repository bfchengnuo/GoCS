package main

import (
	"GoCS/study/entity"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	defer db.Close()

	queryRowDemo()
	println()

	queryMultiRowDemo()
	println()

	//insert4NamedExec()
	select4NameQuery()
}

// 事务操作
// 使用 sqlx 中提供的 db.Beginx() 和 tx.Exec() 方法
func transactionDemo2() (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}

	defer func() {
		// panic 与 recover 是 Go 的两个内置函数，这两个内置函数用于处理 Go 运行时的错误。
		// panic 用于主动抛出错误, recover 用来捕获 panic 抛出的错误。
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update user set age=20 where id=?"
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}

	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "Update user set age=50 where i=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	return err
}

func select4NameQuery() {
	// 使用 map 做命名查询
	rows, err := db.NamedQuery(`SELECT * FROM user WHERE name=:name`, map[string]interface{}{"name": "mps"})
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}

	rows.Close()

	type User struct {
		Name string `db:"name"`
		Age  int    `db:"age"`
	}
	u := User{
		Name: "mps",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(`SELECT * FROM user WHERE name=:name`, u)
	if err != nil {
		fmt.Printf("query2 failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

func insert4NamedExec() (err error) {
	_, err = db.NamedExec(`INSERT INTO user (name,age) VALUES (:name,:age)`,
		map[string]interface{}{
			"name": "沛凝",
			"age":  28,
		})
	return
}

// 插入数据，与标准接口使用一致
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "小魔女", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []entity.User
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u entity.User
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

func initDB() (err error) {
	dsn := "mps:******@tcp(localhost:3306)/skye?charset=utf8&parseTime=True"
	// 也可以使用 MustConnect 连接不成功就 panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
