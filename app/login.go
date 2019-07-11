package app

import (
    "crypto/sha256"
    "encoding/hex"
    "blogkit/config"
)

func login(username string, password string) (bool, string) {
    b := sha256.Sum256([]byte(">> primary salt begins <<" + password + ">> primary salt ends <<"))
    h := hex.EncodeToString(b[:])
    b = sha256.Sum256([]byte(">> secondary salt begins <<" + h + ">> secondary salt ends <<"))
    h = hex.EncodeToString(b[:])
    for k, v := range config.Cfg.Users {
        if v.Name == username {
            if v.Token == h {
                return true, k
            } else {
                return false, "-1"
            }
        }
    }
    return false, "-1"
}
