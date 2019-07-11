package config

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "strconv"
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
    Cfg.Users["0"].Mail = mail
    Cfg.Users["0"].Name = name
    Cfg.Users["0"].Token = getPasswordHash(password)
}

func AddNewUser(mail string, name string, password string) {
    var u User
    u.Mail = mail
    u.Name = name
    u.Token = getPasswordHash(password)
    u.Permissions.CreateComment = true
    id := "-1"
    for k, _ := range Cfg.Users {
        kid, _ := strconv.Atoi(k)
        iid, _ := strconv.Atoi(id)
        if kid > iid {
            id = k
        }
    }
    final, _ := strconv.Atoi(id)
    Cfg.Users[strconv.Itoa(final + 1)] = &u
}

func CreateModule(name string, previous string) {
    var m Module
    m.Name = name
    m.Previous = previous
    id := "-1"
    for k, _ := range Cfg.Modules {
        kid, _ := strconv.Atoi(k)
        iid, _ := strconv.Atoi(id)
        if kid > iid {
            id = k
        }
    }
    final, _ := strconv.Atoi(id)
    Cfg.Modules[strconv.Itoa(final + 1)] = &m
}

func EditModule(id string, name string, previous string) {
    Cfg.Modules[id].Name = name
    Cfg.Modules[id].Previous = previous
}

func DeleteModule(id string) {
    for k, v := range Cfg.Modules {
        if v.Previous == id {
            DeleteModule(k)
        }
    }
    delete(Cfg.Modules, id)
}

func CreateComment(content string, belongsTo string, repliesTo string, authorId string) {
    var c Comment
    c.Content = content
    c.BelongsTo = belongsTo
    c.RepliesTo = repliesTo
    c.AuthorId = authorId
    c.DateTime = time.Now().Format("2006/01/02-15:04:05")
    id := "-1"
    for k, _ := range Cfg.Comments {
        kid, _ := strconv.Atoi(k)
        iid, _ := strconv.Atoi(id)
        if kid > iid {
            id = k
        }
    }
    final, _ := strconv.Atoi(id)
    Cfg.Comments[strconv.Itoa(final + 1)] = &c
}

func DeleteComment(id string) {
    for k, v := range Cfg.Comments {
        if v.RepliesTo == id {
            DeleteComment(k)
        }
    }
    delete(Cfg.Comments, id)
}
