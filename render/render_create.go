package render

func RenderCreate(id string, moduleId string, isLogin bool, loginId string) string {
    result := ""
    result +=
`<html>
<head>
    <title>BlogKit</title>
</head>
<body>
`
    result += renderLogin(isLogin, loginId)
    result += renderCreate(moduleId)
    result +=
`</body>
</html>`
    return result
}
