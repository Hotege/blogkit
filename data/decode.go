package data

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "strings"
)

func DecodeArticle(path string) string {
    result := ""
    f, _ := os.Open(path)
    buffer, _ := ioutil.ReadAll(f)
    f.Close()
    var article []articleSlice
    json.Unmarshal(buffer, &article)
    for _, v := range article {
        switch v.Type {
            case "l":
                result +=
`    <h2>` + v.Value + `</h2>
`
            case "i":
                result +=
`    <img src='` + v.Value + `' /><br>
`
            case "f":
                layers := strings.Split(v.Value, "/")
                result +=
`    <span>Attach file: <a href='` + v.Value + `'>` + layers[len(layers) - 1] + `</a></span><br>
`
            case "c":
                data := strings.Split(v.Value, "\n")
                result +=
`    <b>` + data[0] + ` code</b><br>
`
                for _, l := range data[1:] {
                    result +=
`    <p>` + l + `</p>
`
                }
            case "t":
                fallthrough
            default:
                lines := strings.Split(v.Value, "\n")
                for _, l := range lines {
                    result +=
`    <p>` + l + `</p>
`
                }
        }
    }
    return result
}
