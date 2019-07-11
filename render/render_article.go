package render

import (
    "blogkit/config"
)

func RenderArticle(articleId string, isLogin bool, loginId string) string {
    result :=
`<html>
<head>
    <title>BlogKit</title>
</head>
<body>
`
    result += renderLogin(isLogin, loginId)
    moduleId := config.Cfg.Articles[articleId].ModuleId
    result +=
`    <a href='module?id=` + moduleId + `'><b>...(` + config.Cfg.Modules[moduleId].Name + `)</b></a><br>
`
    result +=
`    <h1>` + config.Cfg.Articles[articleId].Title + `</h1>
`
    result +=
`    <h7>` + config.Cfg.Articles[articleId].DateTime + `</h7>
`
    result += renderArticle(articleId, isLogin, loginId)
    result +=
`</body>
</html>`
    return result
}
