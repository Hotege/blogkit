package render

func RenderPage(moduleId int, isLogin bool, loginId int) string {
    result :=
`<html>
<head>
    <title>BlogKit</title>
</head>
<body>
`
    result += renderLogin(isLogin, loginId)
    result += renderModule(moduleId)
    result +=
`</body>
</html>`
    return result
}
