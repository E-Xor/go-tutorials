package main

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "time"
)

func main() {

  start := time.Now()

  fmt.Println("sql.Open: " + "maksim:passwd@(mysql_host:3306)/hubyak_dev?charset=utf8")

  db, err := sql.Open("mysql", "maksim:passwd@(mysql_host:3306)/hubyak_dev?charset=utf8")
  checkErr(err)

  rows, err := db.Query("SELECT id, name, city, state FROM accounts LIMIT 5")
  checkErr(err)

  for rows.Next() {
    var id int
    var name string
    var city string
    var state string

    err = rows.Scan(&id, &name, &city, &state)
    // checkErr(err) // Scan fails on NULL values trying to convert them to string

    fmt.Printf("%v | %v | %v | %v\n", id, name, city, state)
  }

  db.Close()

  diff := time.Now().Sub(start)
  fmt.Println(diff) // 22-76ms in Go, 8-21ms in rake :-)

}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}
