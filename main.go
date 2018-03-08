package main

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "vknotifier/models"
    _ "vknotifier/routers"
    _ "github.com/mattn/go-sqlite3"
    "os/exec"
    "os"
)

func init() {
    dbDirectory := beego.AppConfig.String("dbDirectory")
    dbFileName := beego.AppConfig.String("dbFileName")
    _, err := os.Stat(dbDirectory)
    if os.IsNotExist(err) {
        exec.Command("/bin/sh", "-c", "sudo mkdir "+dbDirectory).Run()
        exec.Command("/bin/sh", "-c", "sudo chown $USER "+dbDirectory).Run()
    }
    _, err = os.Stat("/var/log/vkCallbackTelegramNotifier.log")
    if os.IsNotExist(err) {
        exec.Command("/bin/sh", "-c", "sudo touch /var/log/vkCallbackTelegramNotifier.log").Run()
        exec.Command("/bin/sh", "-c", "sudo chown $USER /var/log/vkCallbackTelegramNotifier.log").Run()
    }

    orm.RegisterDriver("sqlite3", orm.DRSqlite)
    err = orm.RegisterDataBase("default", "sqlite3", "file:"+dbDirectory+dbFileName)
    if err != nil {
        panic(err)
    }
    // TODO debug
    orm.Debug = true

    models.CreateSchemeIfNotExisted()
}

func main() {
    beego.Run()
}
