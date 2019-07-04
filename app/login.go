package app

import (
    "crypto/sha256"
    "encoding/hex"
    "blogkit/config"
)

func login(username string, password string) (bool, int) {
    b := sha256.Sum256([]byte(">> primary salt begins <<" + password + ">> primary salt ends <<"))
    h := hex.EncodeToString(b[:])
    b = sha256.Sum256([]byte(">> secondary salt begins <<" + h + ">> secondary salt ends <<"))
    h = hex.EncodeToString(b[:])
    for _, v := range config.Cfg.Users {
        if v.Name == username {
            if v.Token == h {
                return true, v.Id
            } else {
                return false, 0
            }
        }
    }
    return false, 0
}
