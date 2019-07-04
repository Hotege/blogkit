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
