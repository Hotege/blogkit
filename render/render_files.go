package render

func RenderFiles(filenames []string) string {
    result := ""
    result +=
`<html>
<head>
    <title>BlogKit - Files</title>
</head>
<body>
`
    for _, v := range filenames {
        result +=
`    <input value='` + v + `' />
`
    }
    result +=
`
</body>
</html>
`
    return result
}
