package data

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "strings"
)

type articleSlice struct {
    Type string `json:Type`
    Value string `json:Value`
}

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
