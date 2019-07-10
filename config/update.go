package config

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "time"
)

func UpdateFile() {
    output, _ := json.Marshal(Cfg)
    ioutil.WriteFile("config.json", output, 0666)
}

func getPasswordHash(password string) string {
    b := sha256.Sum256([]byte(">> primary salt begins <<" + password + ">> primary salt ends <<"))
    h := hex.EncodeToString(b[:])
    b = sha256.Sum256([]byte(">> secondary salt begins <<" + h + ">> secondary salt ends <<"))
    h = hex.EncodeToString(b[:])
    return h
}

func UpdateAdmin(mail string, name string, password string) {
    Cfg.Users[0].Mail = mail
    Cfg.Users[0].Name = name
    Cfg.Users[0].Token = getPasswordHash(password)
}

func AddNewUser(mail string, name string, password string) {
    var u User
    u.Id = Cfg.Users[len(Cfg.Users) - 1].Id + 1
    u.Mail = mail
    u.Name = name
    u.Token = getPasswordHash(password)
    u.Permissions.CreateComment = true
    Cfg.Users = append(Cfg.Users, u)
}

func CreateComment(content string, belongsTo int, repliesTo int, authorId int) {
    var c Comment
    c.Id = Cfg.Comments[len(Cfg.Comments) - 1].Id + 1
    c.Content = content
    c.BelongsTo = belongsTo
    c.RepliesTo = repliesTo
    c.AuthorId = authorId
    c.DateTime = time.Now().Format("2006/01/02-15:04:05")
    Cfg.Comments = append(Cfg.Comments, c)
}
