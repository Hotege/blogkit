package data

import (
    "blogkit/config"
)

func CheckModuleExistById(id string) bool {
    for k, _ := range config.Cfg.Modules {
        if k == id {
            return true
        }
    }
    return false
}

func CheckArticleExistById(id string) bool {
    for k, _ := range config.Cfg.Articles {
        if k == id {
            return true
        }
    }
    return false
}

func CheckCommentExistById(id string) bool {
    for k, _ := range config.Cfg.Comments {
        if k == id {
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

func CheckUserExistById(id string) bool {
    for k, _ := range config.Cfg.Users {
        if k == id {
            return true
        }
    }
    return false
}
