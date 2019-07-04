package render

func RenderArticle(articleId int, isLogin bool, loginId int) string {
    result :=
`<html>
<head>
    <title>BlogKit</title>
</head>
<body>
`
    result += renderLogin(isLogin, loginId)
    result += renderArticle(articleId)
    result +=
`</body>
</html>`
    return result
}
