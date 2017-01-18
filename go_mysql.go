package main
import (
    "database/sql"
    "fmt"
    _"github.com/go-sql-driver/mysql"
)

func main(){
    db, _:=sql.Open("mysql","root:root@/appstore")
    rows, _:= db.Query("SELECT * FROM application")
    //fmt.Println(rows.Columns())
    fmt.Println(rows.Columns())
}



