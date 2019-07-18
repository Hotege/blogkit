package data

import (
    "os"
    "io/ioutil"
    "encoding/json"
    "strings"
)

func DecodeArticleStruct(path string) []articleSlice {
    f, _ := os.Open(path)
    buffer, _ := ioutil.ReadAll(f)
    f.Close()
    var article []articleSlice
    json.Unmarshal(buffer, &article)
    for k, v := range article {
        switch v.Type {
            case "i":
                layers := strings.Split(v.Value, "/")
                article[k].Extra = layers[len(layers) - 1]
            case "f":
                layers := strings.Split(v.Value, "/")
                article[k].Extra = layers[len(layers) - 1]
            case "c":
                article[k].Extra = strings.Split(v.Value, "\n")
            case "t":
                fallthrough
            default:
                article[k].Extra = strings.Split(v.Value, "\n")
        }
    }
    return article
}
