package main
import (
    "fmt"
    "net/http"
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    http.HandleFunc("/", Index)
    http.HandleFunc("/mushroom", Mushroom)
    http.HandleFunc("/normal", normal)
    http.HandleFunc("/database", database)
//    http.HandleFunc("/magic/{magicID}", MagicShow)
    http.ListenAndServe(":9999", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, world\n");
}

func Mushroom(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello mushroom")
}

func MagicShow(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintln(w, "%d\n", magicID)
}

func normal(w http.ResponseWriter, r *http.Request) {
    time.Sleep(time.Second * 30) 
    fmt.Fprintln(w, "waiting to deal with")
}

func database(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Connecting to MySQL...")
    db, err := sql.Open("mysql", "root:root@tcp(192.168.3.254:3306)/maxwell")
    if err != nil {
        fmt.Fprintln(w, "connecting to database error")
        panic(err)
    }
    
   // rows, err := db.Query("SELECT * FROM phones")
    rows, err := db.Query("SELECT serid, name FROM phones WHERE serid = 2")

    if err != nil {
        fmt.Fprintln(w, "selecting error")
        panic(err)
    } 

    columns, err := rows.Columns()
    if err != nil {
        fmt.Fprintln(w, "selecting error")
        panic(err)
    }
//    fmt.Fprintln(w, columns)
    values := make([]sql.RawBytes, len(columns))
    
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
    
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err)
        }
        
        var value string 
        for i, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Fprintln(w, columns[i],":", value)
        }
        fmt.Fprintln(w, "--------------------------------")
    
    }
    if err = rows.Err(); err != nil {
        panic(err)
    } 
}
