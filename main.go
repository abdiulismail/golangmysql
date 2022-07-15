package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db, er := sql.Open("mysql","root:password$@tcp(127.0.0.1:3306)/golangmysql")
	if er != nil{
		panic(er.Error())
	}
	defer db.Close()

	//_, err := db.Query("INSERT INTO USERS VALUES('abdi')")
	//if err != nil{
	//	panic(er.Error())
	//}

	res, err := db.Query("select name from users")
	if err != nil{
		panic(er.Error())
	}

	for res.Next(){
		var user User
		err = res.Scan(&user.name)
		if err != nil{
			panic(er.Error())
		}
		fmt.Println(user.name)
	}
}

type User struct {
	name string
}

