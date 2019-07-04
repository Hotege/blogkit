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
