package data

import (
    "blogkit/config"
)

func CheckModuleExistById(id int) bool {
    for _, v := range config.Cfg.Modules {
        if v.Id == id {
            return true
        }
    }
    return false
}

func CheckArticleExistById(id int) bool {
    for _, v := range config.Cfg.Articles {
        if v.Id == id {
            return true
        }
    }
    return false
}

func CheckUserExistByMail(mail string) bool {
    for _, v := range config.Cfg.Users {
        if v.Mail == mail {
            return true
        }
    }
    return false
}

func CheckUserExistByName(name string) bool {
    for _, v := range config.Cfg.Users {
        if v.Name == name {
            return true
        }
    }
    return false
}

func CheckUserExistById(id int) bool {
    for _, v := range config.Cfg.Users {
        if v.Id == id {
            return true
        }
    }
    return false
}
