// a golang version of multiprocess and exec shell commands
package main

import (
    "os/exec"
    "time"
)

func importSql() {
    time.Sleep(time.Second * time.Duration(30))
    _, err := exec.Command("/bin/bash", "/root/importSql.sh")
    if err != nil {
        panic(err)
    }
}

func main() {
    _, err := exec.Command("mysqld", "--initialize-insecure", "--user=root")
    if err != nil {
        panic(err)
    } 
    
    go importSql()
    
    _, err = exec.Command("mysqld")
    if err != nil {
        panic(err)
    }
}

