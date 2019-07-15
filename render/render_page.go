package render

import (
    "blogkit/config"
)

func RenderPage(moduleId string, isLogin bool, loginId string) string {
    result :=
`<html>
<head>
    <title>BlogKit</title>
</head>
<body>
`
    result += renderLogin(isLogin, loginId)
    if isLogin {
        if config.Cfg.Users[loginId].Permissions.CreateArticle {
            result +=
`    <a href='create?moudle=` + moduleId + `'><b>New article</b></a><br>
`
        }
    }
    result += renderModule(moduleId, isLogin, loginId)
    result +=
`</body>
</html>`
    return result
}
