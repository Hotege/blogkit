package app

import (
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "os"
    "strings"
    "time"
    "strconv"
    "blogkit/config"
    "blogkit/render"
    "blogkit/data"
)

func Run() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
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
            if config.Cfg.Users["0"].Token == "None" {
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
                    config.SaveConfig()
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            } else {
                http.Redirect(w, r, "/", http.StatusFound)
                return
            }
        }
        if dir[1] == "login" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                if r.Method == "POST" {
                    r.ParseForm()
                    success, id := login(r.PostForm["login_username"][0], r.PostForm["login_password"][0]) 
                    if success {
                        newCookie := http.Cookie {
                            Name: "user_login",
                            Value: "login:true,id:" + id,
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
            if config.Cfg.Users["0"].Token == "None" {
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
                    config.SaveConfig()
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if dir[1] == "logout" {
            if config.Cfg.Users["0"].Token == "None" {
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
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                r.ParseForm()
                login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
                isLogin := strings.Split(login_data[0], ":")[1]
                loginId := strings.Split(login_data[1], ":")[1]
                if _, ok := r.Form["id"]; ok {
                    id := r.Form["id"][0]
                    if r.Method == "GET" {
                        if id == "0" {
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
                    }
                    if r.Method == "POST" {
                        if isLogin != "true" {
                            io.WriteString(w, render.RenderError("Need login."))
                            return
                        }
                        if !data.CheckUserExistById(loginId) {
                            newCookie := http.Cookie {
                                Name: "user_login",
                                Value: "login:false,id:0",
                            }   
                            http.SetCookie(w, &newCookie)
                            loginCookie = &newCookie
                            io.WriteString(w, render.RenderError("User error."))
                            return
                        }
                        if data.CheckModuleExistById(id) {
                            if r.Form["do"][0] == "edit" {
                                if r.PostForm["module_edit_type"][0] == "create" {
                                    config.CreateModule(r.PostForm["module_edit_name"][0], r.PostForm["module_edit_pid"][0])
                                }
                                if r.PostForm["module_edit_type"][0] == "edit" {
                                    config.EditModule(r.PostForm["module_edit_id"][0], r.PostForm["module_edit_name"][0], r.PostForm["module_edit_pid"][0])
                                }
                            }
                            if r.Form["do"][0] == "delete" {
                                config.DeleteModule(r.PostForm["delete_id"][0])
                            }
                            if r.Form["do"][0] == "delete_article" {
                                config.DeleteArticle(r.PostForm["delete_article_id"][0])
                            }
                            config.SaveConfig()
                            http.Redirect(w, r, "/module?id=" + r.Form["id"][0], http.StatusFound)
                            return
                        } else {
                            io.WriteString(w, render.RenderError("Page not found."))
                            return
                        }
                    }
                } else {
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if strings.Split(dir[1], "?")[0] == "article" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                r.ParseForm()
                login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
                isLogin := strings.Split(login_data[0], ":")[1]
                loginId := strings.Split(login_data[1], ":")[1]
                if _, ok := r.Form["id"]; ok {
                    id := r.Form["id"][0]
                    if r.Method == "GET" {
                        if data.CheckArticleExistById(id) {
                            io.WriteString(w, render.RenderArticle(id, isLogin == "true", loginId))
                            return
                        } else {
                            io.WriteString(w, render.RenderError("Page not found."))
                            return
                        }
                    }
                    if r.Method == "POST" {
                        if isLogin != "true" {
                            io.WriteString(w, render.RenderError("Need login."))
                            return
                        }
                        if !data.CheckUserExistById(loginId) {
                            newCookie := http.Cookie {
                                Name: "user_login",
                                Value: "login:false,id:0",
                            }   
                            http.SetCookie(w, &newCookie)
                            loginCookie = &newCookie
                            io.WriteString(w, render.RenderError("User error."))
                            return
                        }
                        if data.CheckArticleExistById(id) {
                            if r.Form["do"][0] == "reply_comment" {
                                if r.PostForm["reply_type"][0] == "new" {
                                    config.CreateComment(r.PostForm["reply_comment_content"][0], id, "-1", loginId)
                                }
                                if r.PostForm["reply_type"][0] == "reply" {
                                    config.CreateComment(r.PostForm["reply_comment_content"][0], id, r.PostForm["reply_id"][0], loginId)
                                }
                            }
                            if r.Form["do"][0] == "delete_comment" {
                                config.DeleteComment(r.PostForm["delete_id"][0])
                            }
                            config.SaveConfig()
                            http.Redirect(w, r, "/article?id=" + r.Form["id"][0], http.StatusFound)
                            return
                        } else {
                            io.WriteString(w, render.RenderError("Page not found."))
                            return
                        }
                    }
                } else {
                    http.Redirect(w, r, "/", http.StatusFound)
                    return
                }
            }
        }
        if strings.Split(dir[1], "?")[0] == "create" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            } else {
                r.ParseForm()
                login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
                isLogin := strings.Split(login_data[0], ":")[1]
                loginId := strings.Split(login_data[1], ":")[1]
                if isLogin != "true" {
                    io.WriteString(w, render.RenderError("Need login."))
                    return
                }
                if !config.Cfg.Users[loginId].Permissions.CreateArticle {
                    io.WriteString(w, render.RenderError("No permission(s)."))
                    return
                }
                id := "-1"
                moduleId := "0"
                if _, okModule := r.Form["module"]; okModule {
                    moduleId = r.Form["module"][0]
                }
                if _, okDo := r.Form["do"]; okDo {
                    id = r.Form["id"][0]
                }
                if r.Method == "GET" {
                    if id == "-1" {
                        io.WriteString(w, render.RenderCreate(id, moduleId, isLogin == "true", loginId))
                        return
                    }
                }
                if r.Method == "POST" {
                    filepath := config.CreateArticle(
                        moduleId, loginId, r.PostForm["create_title"][0],
                    )
                    fmt.Println(filepath)
                    config.SaveConfig()
                    result, _ := data.EncodeArticle(r.PostForm["step_s"], r.PostForm["step_ti"], r.PostForm["step_ii_ii"], r.PostForm["step_ff_fi"], r.PostForm["step_ci"])
                    ioutil.WriteFile(filepath, result, 0666)
                }
            }
        }
        if dir[1] == "upload_images" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            }
            login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
            isLogin := strings.Split(login_data[0], ":")[1]
            loginId := strings.Split(login_data[1], ":")[1]
            if isLogin != "true" {
                io.WriteString(w, render.RenderError("Need login."))
                return
            }
            if !config.Cfg.Users[loginId].Permissions.CreateArticle {
                io.WriteString(w, render.RenderError("No permission(s)."))
                return
            }
            reader, _ := r.MultipartReader()
            filename := ""
            files := make([]string, 0);
            for {
                part, err := reader.NextPart()
                if err == io.EOF {
                    break
                }
                if part.FileName() == "" {
                } else {
                    filename = "static/files/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + loginId + "-" + part.FileName()
                    files = append(files, filename)
                    dst, _ := os.Create(filename)
                    defer dst.Close()
                    io.Copy(dst, part)
                }
            }
            io.WriteString(w, render.RenderFiles(files))
            return
        }
        if dir[1] == "upload_files" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            }   
            login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
            isLogin := strings.Split(login_data[0], ":")[1]
            loginId := strings.Split(login_data[1], ":")[1]
            if isLogin != "true" {
                io.WriteString(w, render.RenderError("Need login."))
                return
            }
            if !config.Cfg.Users[loginId].Permissions.CreateArticle {
                io.WriteString(w, render.RenderError("No permission(s)."))
                return
            }
            reader, _ := r.MultipartReader()
            files := make([]string, 0)
            for {
                part, err := reader.NextPart()
                if err == io.EOF {
                    break
                }   
                if part.FileName() == "" {
                } else {
                    filename := "static/files/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + loginId + "-" + part.FileName()
                    files = append(files, filename)
                    dst, _ := os.Create(filename)
                    defer dst.Close()
                    io.Copy(dst, part)
                }   
            }   
            io.WriteString(w, render.RenderFiles(files))
            return
        }
        if dir[1] == "" {
            if config.Cfg.Users["0"].Token == "None" {
                http.Redirect(w, r, "/initialize", http.StatusFound)
                return
            }
            login_data := strings.Split(strings.Trim(strings.Split(loginCookie.String(), "=")[1], "\""), ",")
            isLogin := strings.Split(login_data[0], ":")[1]
            loginId := strings.Split(login_data[1], ":")[1]
            io.WriteString(w, render.RenderPage("0", isLogin == "true", loginId))
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
