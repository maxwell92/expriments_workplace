package main
import (
    "fmt"
    "time"
    "math/rand"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

var nameArray map[int]string
var priceArray map[int]string

func initArray() {
    nameArray = make(map[int]string)
    nameArray[0] = "iphone3gs"
    nameArray[1] = "iphone4"
    nameArray[2] = "iphone4s"
    nameArray[3] = "iphone5"
    nameArray[4] = "iphone5s"
    nameArray[5] = "iphone5c"
    nameArray[6] = "iphone6"
    nameArray[7] = "iphone6Plus"
    nameArray[8] = "iphone6s"
    nameArray[9] = "iphone6sPlus"
    nameArray[10]  = "SamsungNote2"
    nameArray[11] = "SamsungNote3"
    nameArray[12] = "SamsungNote4"
    nameArray[13] = "SamsungNote5"
    nameArray[14] = "SamsungS4"
    nameArray[15] = "SamsungS5"
    nameArray[16] = "SamsungS6"
    nameArray[17] = "SamsungS6Edge"

    priceArray = make(map[int]string)
    priceArray[0] = "500"
    priceArray[1] = "600"
    priceArray[2] = "1000"
    priceArray[3] = "1500"
    priceArray[4] = "2000"
    priceArray[5] = "1300"
    priceArray[6] = "3000"
    priceArray[7] = "3500"
    priceArray[8] = "5000"
    priceArray[9] = "5500"
    priceArray[10] = "1000"
    priceArray[11] = "1500"
    priceArray[12] = "2000"
    priceArray[13] = "2500"
    priceArray[14] = "3000"
    priceArray[15] = "3500"
    priceArray[16] = "4000"
    priceArray[17] = "4500"

    
}



func randValues() (name string, price string) {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))    
    trand := r.Intn(100)
    num := trand % 18 
    return nameArray[num], priceArray[num]
}



func main() {   

    initArray()

    db, err := sql.Open("mysql", "root:root@tcp(192.168.3.254:3306)/maxwell")
    if err != nil {
        fmt.Println("open sql database error")
        panic(err)
    }
    
//    serid := 5
//    name := "iphone5c"
//    price := 1000 
 
    for i := 0; i < 10000; i++ {
        serid := 105 + i
        name, price := randValues()
        stmt, err := db.Prepare(`INSERT phones (serid, name, price) VALUES(?, ?, ?)`)
        if err != nil {
            fmt.Println("Prepare error")
            panic(err)
        }
        _, err = stmt.Exec(serid, name, price)
        if err != nil {
            fmt.Println("Exec error")
            panic(err)
        }

    }
     
}
