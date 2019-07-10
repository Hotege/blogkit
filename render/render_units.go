package render

import (
    "strconv"
    "blogkit/config"
    "blogkit/data"
)

func renderLogin(isLogin bool, loginId int) string {
    result := ""
    if isLogin {
        result +=
`    <span><h5>` + config.Cfg.Users[loginId].Name + `</h5></span>
    <a href='logout'>logout</a>
    <hr>
`
    } else {
        result +=
`    <form action='login' method='POST'>
        <span>Username: <input name='login_username' /></span><br>
        <span>Password: <input name='login_password' type='password' /></span><br>
        <span><input type='submit' value='login' /> <a href='signup'>sign up</a></span>
    </form>
    <hr>
`
    }
    return result
}

func renderModule(id int, isLogin bool, loginId int) string {
    result := ""
    if id != 0 {
        previousId := config.Cfg.Modules[id].Previous
        result +=
`    <a href='module?id=` + strconv.Itoa(previousId) + `'><h4>...(` + config.Cfg.Modules[previousId].Name + `)</h4></a>
`
    }
    result +=
`    <h3>` + config.Cfg.Modules[id].Name + `</h3>
`
    for _, v := range config.Cfg.Articles {
        if v.ModuleId == id {
            result +=
`    <a href='article?id=` + strconv.Itoa(v.Id) + `'><b>` + v.Title + `</b></a>`
            result += `<span>`
            if isLogin {
                if config.Cfg.Users[loginId].Permissions.EditArticle {
                    result += ` <a href='#'>Edit</a>`
                }
                if config.Cfg.Users[loginId].Permissions.DeleteArticle {
                    result += ` <a href='#'>Delete</a>`
                }
            }
            result += `</span><br>
`
        }
    }
    for _, v := range config.Cfg.Modules {
        if v.Previous == id {
            result +=
`    <a href='module?id=` + strconv.Itoa(v.Id) + `'><b>` + v.Name + `</b></a><br>
`
        }
    }
    return result
}

func renderArticle(id int, isLogin bool, loginId int) string {
    result := ""
    result +=
`    <script>
        function setMainComment(id, tId, context) {
            var obj = document.getElementById(id);
            if (obj.style.display == "none") {
                obj.style.display = "block";
                document.getElementById(tId).innerHTML = "Hide";
            } else {
                obj.style.display = "none";
                document.getElementById(tId).innerHTML = context;
            }   
        }   
        function showReplyComment(id, hId, rId) {
            var obj = document.getElementById(id);
            document.getElementById(hId).style.display = "block";
            document.getElementById(hId).style.height = document.body.clientHeight + "px";
            obj.style.display = "block";
            document.getElementById('reply_id').value = rId;
            document.getElementById('reply_comment_content').focus();
        }
        function hideReplyComment(id, hId) {
            document.getElementById(hId).style.display = "none";
            document.getElementById(id).style.display = "none";
        }
        function submitReplyComment(id, hId) {
            hideReplyComment(id, hId);
            document.getElementById("reply_form").submit();
        }
    </script>
`
    result +=
`    <div id='hidebg' style='position: absolute; left: 0px; top: 0px; background-color: #000000; width: 100%; filter: alpha(opacity=60); opacity: 0.6; z-index: 2;'></div>
`
    result +=
`    <div id='reply_comment' style='position: absolute; left: 30%; top: 200px; background-color: #fff; border: 1px solid black; display: none; z-index: 3;'>
        <form id='reply_form' action='article?id=` +strconv.Itoa(id) + `&do=reply_comment' method='POST'>
            <span><input id='reply_id' name='reply_id' readonly='readonly' style='display: none;' />
            <div style='display: inline-block; cursor: pointer;' onclick='hideReplyComment("reply_comment", "hidebg");'>X</div></span><br>
            <span>content:</span><br>
            <textarea id='reply_comment_content' name='reply_comment_content' style='resize: none; width: 320px; height: 160px;'></textarea><br>
            <div style='display: inline-block; background-color: blue; cursor: pointer;' onclick='submitReplyComment("reply_comment", "hidebg");'>submit</div>
        </form>
    </div>
`
    result += `    <span>`
    if isLogin {
        if config.Cfg.Users[loginId].Permissions.EditArticle {
            result += ` <a href='#'>Edit</a>`
        }
        if config.Cfg.Users[loginId].Permissions.DeleteArticle {
            result += ` <a href='#'>Delete</a>`
        }
    }
    result += `</span><br>
`
    result +=
`    ` + data.DecodeArticle(config.Cfg.Articles[id].Path) + `
`
    if isLogin {
        if config.Cfg.Users[loginId].Permissions.CreateComment {
            result +=
`    <span><a id='main_comment_set' href='javascript:setMainComment("main_comment", "main_comment_set", "New comment");'>New comment</a></span><br>
    <div id='main_comment' style='display: none;'>
        <form action='article?id=` + strconv.Itoa(id) + `&do=main_comment' method='POST'>
            <textarea name='main_comment_content' style='resize: none; width: 320px; height: 160px;'></textarea><br>
            <input type='submit' value='submit' />
        </form>
    </div>
`
        }
    }
    for _, v := range config.Cfg.Comments {
        if v.BelongsTo == id && v.RepliesTo == -1 {
            result +=
`    <b>` + config.Cfg.Users[v.AuthorId].Name + ": " + v.Content + " - " + v.DateTime + `</b>`
            if isLogin {
                if config.Cfg.Users[loginId].Permissions.CreateComment {
                    replyId := "reply_" + strconv.Itoa(v.Id)
                    result += `<span> <a id='` + replyId + `' href='javascript:showReplyComment("reply_comment", "hidebg", "` + strconv.Itoa(v.Id) + `");'>Reply</a></span>`
                }
                if config.Cfg.Users[loginId].Permissions.DeleteComment {
                    result += `<span> <a href='#'>Delete</a></span>`
                }
            }
            result += `<br>
`
            for _, sub := range config.Cfg.Comments {
                if sub.BelongsTo == id && sub.RepliesTo != -1 {
                    previousId := sub.Id
                    check := false
                    for config.Cfg.Comments[previousId].RepliesTo != -1 {
                        previousId = config.Cfg.Comments[previousId].RepliesTo
                    }
                    check = config.Cfg.Comments[previousId].Id == v.Id
                    if check {
                        result +=
`    <span>` + config.Cfg.Users[sub.AuthorId].Name + ": " + sub.Content + " - " + sub.DateTime + `</span>`
                        if isLogin {
                            if config.Cfg.Users[loginId].Permissions.CreateComment {
                                replyId := "reply_" + strconv.Itoa(sub.Id)
                                result += `<span> <a id='` + replyId + `' href='javascript:showReplyComment("reply_comment", "hidebg", "` + strconv.Itoa(sub.Id) + `");'>Reply</a></span>`
                            }
                            if config.Cfg.Users[loginId].Permissions.DeleteComment {
                                result += `<span> <a href='#'>Delete</a></span>`
                            }
                        }
                        result += `<br>
`
                    }
                }
            }
        }
    }
    return result
}
