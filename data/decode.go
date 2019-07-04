package data

import (
    "os"
    "io/ioutil"
)

func DecodeArticle(path string) string {
    result := ""
    f, _ := os.Open(path)
    data, _ := ioutil.ReadAll(f)
    result += string(data)
    f.Close()
    return result
}
