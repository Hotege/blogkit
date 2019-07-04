package app

import (
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
    "blogkit/config"
    "blogkit/render"
    "blogkit/data"
)

func Run() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r)
        dir := strings.Split(r.URL.String(), "/")
        if dir[1] == "favicon.ico" {
            file, _ := os.Open("favicon.ico")
            defer file.Close()
            buffer, _ := ioutil.ReadAll(file)
            w.Write(buffer)
            return
        }
        loginCookie, err := r.Cookie("user_login")
        if err != nil {
            newCookie := http.Cookie {
                Name: "user_login",
                Value: "login:false,id:0",
            }
            http.SetCookie(w, &newCookie)
            loginCookie = &newCookie
        }
        if dir[1] == "initialize" {
            if config.Cfg.Users[0].Token == "None" {
                if r.Method == "GET" {
                    io.WriteString(w, render.RenderInitialize())
                    return
                }
                if r.Method == "POST" {
                    r.ParseForm()
                    if r.PostForm["init_password"][0] != r.PostForm["init_confirm"][0] {
                        io.WriteString(w, render.RenderError("Password not confirmed."))
                        return
                    }
                    config.UpdateAdmin(r.PostForm["init_mail"][0], r.PostForm["init_name"][0], r.PostForm["init_password"][0])
                    config.UpdateFile()
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            } else {
                http.Redirect(w, r, "/", http.StatusFound)
                return
            }
        }
        if dir[1] == "login" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                if r.Method == "POST" {
                    r.ParseForm()
                    success, id := login(r.PostForm["login_username"][0], r.PostForm["login_password"][0]) 
                    if success {
                        newCookie := http.Cookie {
                            Name: "user_login",
                            Value: "login:true,id:" + strconv.Itoa(id),
                        }
                        http.SetCookie(w, &newCookie)
                        loginCookie = &newCookie
                    }
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
            http.Redirect(w, r, "/", http.StatusFound)
            return
        }
        if dir[1] == "signup" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                if r.Method == "GET" {
                    io.WriteString(w, render.RenderSignUp())
                    return
                }
                if r.Method == "POST" {
                    r.ParseForm()
                    if r.PostForm["signup_password"][0] != r.PostForm["signup_confirm"][0] {
                        io.WriteString(w, render.RenderError("Password not confirmed."))
                        return
                    }
                    if data.CheckUserExistByMail(r.PostForm["signup_mail"][0]) || data.CheckUserExistByName(r.PostForm["signup_name"][0]) {
                        io.WriteString(w, render.RenderError("Mail or name already exist."))
                        return
                    }
                    config.AddNewUser(r.PostForm["signup_mail"][0], r.PostForm["signup_name"][0], r.PostForm["signup_password"][0])
                    config.UpdateFile()
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if dir[1] == "logout" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                newCookie := http.Cookie {
                    Name: "user_login",
                    Value: "login:false,id:0",
                }
                http.SetCookie(w, &newCookie)
                loginCookie = &newCookie
                http.Redirect(w, r, "/", http.StatusFound)
                return
            }
        }
        if strings.Split(dir[1], "?")[0] == "module" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                r.ParseForm()
                login_data := strings.Split(strings.Split(loginCookie.String(), "=")[1], ",")
                isLogin := strings.Split(login_data[0], ":")[1]
                loginId, _ := strconv.Atoi(strings.Split(login_data[1], ":")[1])
                if _, ok := r.Form["id"]; ok {
                    id, _ := strconv.Atoi(r.Form["id"][0])
                    if id == 0 {
                        http.Redirect(w, r, "/", http.StatusFound)
                        return
                    }
                    if data.CheckModuleExistById(id) {
                        io.WriteString(w, render.RenderPage(id, isLogin == "true", loginId))
                        return
                    } else {
                        io.WriteString(w, render.RenderError("Page not found."))
                        return
                    }
                } else {
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if strings.Split(dir[1], "?")[0] == "article" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                r.ParseForm()
                login_data := strings.Split(strings.Split(loginCookie.String(), "=")[1], ",")
                isLogin := strings.Split(login_data[0], ":")[1]
                loginId, _ := strconv.Atoi(strings.Split(login_data[1], ":")[1])
                if _, ok := r.Form["id"]; ok {
                    id, _ := strconv.Atoi(r.Form["id"][0])
                    if data.CheckArticleExistById(id) {
                        io.WriteString(w, render.RenderArticle(id, isLogin == "true", loginId))
                        return
                    } else {
                        io.WriteString(w, render.RenderError("Page not found."))
                        return
                    }
                } else {
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if dir[1] == "" {
            if config.Cfg.Users[0].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            }
            login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
            isLogin := strings.Split(login_data[0], ":")[1]
            loginId, _ := strconv.Atoi(strings.Split(login_data[1], ":")[1])
            io.WriteString(w, render.RenderPage(0, isLogin == "true", loginId))
            return
        }
        http.Redirect(w, r, "/", http.StatusFound)
        return
    })
    err := http.ListenAndServe("0.0.0.0:80", nil)
    if err != nil {
        fmt.Println(err)
    }
}
