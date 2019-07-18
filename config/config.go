package config

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

func LoadConfig() {
    f, _ := os.Open("config.json")
    content, _ := ioutil.ReadAll(f)
    json.Unmarshal(content, &Cfg)
    f.Close()
}

func SaveConfig() {
    content, _ := json.Marshal(Cfg)
    ioutil.WriteFile("config.json", content, 0666)
}

func GetRootComment() map[string]string {
    var result = make(map[string]string)
    for k, v := range Cfg.Comments {
        if v.RepliesTo != "-1" {
            previousId := k
            _, ok := Cfg.Comments[previousId]
            if !ok {
                continue
            }
            for Cfg.Comments[previousId].RepliesTo != "-1" {
                previousId = Cfg.Comments[previousId].RepliesTo
            }
            result[k] = previousId
        }
    }
    return result
}
