package models

import (
    "github.com/astaxie/beego/orm"
    "golang.org/x/crypto/bcrypt"
)

func init() { orm.RegisterModel(new(User)) }

func (u *User) TableName() string {
    return "users"
}

type User struct {
    Id       int
    Login    string
    Password string
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CheckLoginPassword(login, password string) bool {
    var user User
    o := orm.NewOrm()
    o.QueryTable("users").Filter("Login", login).One(&user)

    if user.Id > 0 {
        if CheckPasswordHash(password, user.Password) == true {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}

func CreateSchemeIfNotExisted() bool {
    o := orm.NewOrm()
    num, _ := o.QueryTable("users").Count()

    if num == 0 {
        o.Raw("CREATE TABLE communities(id INTEGER PRIMARY KEY AUTOINCREMENT, vk_id INTEGER, `name` TEXT, `type` TEXT, `key` TEXT, `check_response` TEXT)").Exec()
        o.Raw("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, login TEXT, password TEXT, bot_id INT, bot_token TEXT, chat_id INT);").Exec()
        // set hash for password "secret"
        o.Raw(`INSERT INTO users (login, password) VALUES ('odmin','$2a$14$dmCpB2Hync5ghQwScJ/RYe24XjuisBovJDHAPPq9jlkQjpTSydSFi')`).Exec()

        return true
    } else {
        return false
    }
}
