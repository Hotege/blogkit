package render

import (
    "strconv"
    "blogkit/config"
    "blogkit/data"
)

func renderLogin(isLogin bool, loginId int) string {
    result := ""
    if isLogin {
        result +=
`    <span><h5>` + config.Cfg.Users[loginId].Name + `</h5></span>
    <a href='logout'>logout</a>
    <hr>
`
    } else {
        result +=
`    <form action='login' method='POST'>
        <span>Username: <input name='login_username' /></span><br>
        <span>Password: <input name='login_password' type='password' /></span><br>
        <span><input type='submit' value='login' /> <a href='signup'>sign up</a></span>
    </form>
    <hr>
`
    }
    return result
}

func renderModule(id int) string {
    result := ""
    if id != 0 {
        result +=
`    <a href='module?id=` + strconv.Itoa(config.Cfg.Modules[id].Previous) + `'><h4>...(` + config.Cfg.Modules[config.Cfg.Modules[id].Previous].Name + `)</h4></a>
`
    }
    result +=
`    <h3>` + config.Cfg.Modules[id].Name + `</h3>
`
    for _, v := range config.Cfg.Articles {
        if v.ModuleId == id {
            result +=
`    <a href='article?id=` + strconv.Itoa(v.Id) + `'><b>` + v.Title + `</b></a><br>
`
        }
    }
    for _, v := range config.Cfg.Modules {
        if v.Previous == id {
            result +=
`    <a href='module?id=` + strconv.Itoa(v.Id) + `'><b>` + v.Name + `</b></a><br>
`
        }
    }
    return result
}

func renderArticle(id int) string {
    result := ""
    result +=
`` + data.DecodeArticle(config.Cfg.Articles[id].Path) + `<br>
`
    for _, v := range config.Cfg.Comments {
        if v.BelongsTo == id && v.RepliesTo == -1 {
            result +=
`<b>` + config.Cfg.Users[v.AuthorId].Name + ": " + v.Content + " - " + v.DateTime + `</b><br>
`
            for _, sub := range config.Cfg.Comments {
                if sub.BelongsTo == id && sub.RepliesTo != -1 {
                    previousId := sub.Id
                    check := false
                    for config.Cfg.Comments[previousId].RepliesTo != -1 {
                        previousId = config.Cfg.Comments[previousId].RepliesTo
                    }
                    check = config.Cfg.Comments[previousId].Id == v.Id
                    if check {
                        result +=
`<p>` + config.Cfg.Users[sub.AuthorId].Name + ": " + sub.Content + " - " + sub.DateTime + `</p>
`
                    }
                }
            }
        }
    }
    return result
}
