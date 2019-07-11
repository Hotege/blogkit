package render

import (
    "sort"
    "strconv"
    "blogkit/config"
    "blogkit/data"
)

type ByDemical []string

func (a ByDemical) Len() int {
    return len(a)
}

func (a ByDemical) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByDemical) Less(i, j int) bool {
    x, _ := strconv.Atoi(a[i])
    y, _ := strconv.Atoi(a[j])
    return x < y
}

func renderLogin(isLogin bool, loginId string) string {
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

func renderModule(id string, isLogin bool, loginId string) string {
    result := ""
    if id != "0" {
        previousId := config.Cfg.Modules[id].Previous
        result +=
`    <a href='module?id=` + previousId + `'><h4>...(` + config.Cfg.Modules[previousId].Name + `)</h4></a>
`
    }
    if isLogin {
        result +=
`    <script>
        function showModuleEditor(id, hId, mId, pId, value, type) {
            document.getElementById(hId).style.display = "block";
            document.getElementById(hId).style.height = document.body.clientHeight + "px";
            document.getElementById(id).style.display = "block";
            document.getElementById('module_edit_id').value = mId;
            document.getElementById('module_edit_pid').value = pId;
            document.getElementById('module_edit_name').value = value;
            document.getElementById('module_edit_type').value = type;
            var options = document.getElementById('select_previous');
            for (i = 0; i < options.length; i++) {
                if (options[i].id == pId) {
                    options[i].selected = true;
                    break;
                }
            }
            document.getElementById('module_edit_name').focus();
        }
        function hideModuleEditor(id, hId) {
            document.getElementById(hId).style.display = "none";
            document.getElementById(id).style.display = "none";
        }
        function submitModuleEditor(id, hId) {
            hideModuleEditor(id, hId);
            document.getElementById("module_form").submit();
        }
        function deleteModule(mId) {
            var msg = confirm("Will you delete this module?");
            if (msg) {
                document.getElementById("delete_id").value = mId;
                document.getElementById("delete_form").submit();
            }
        }
    </script>
    <div id='hidebg' style='position: absolute; left: 0px; top: 0px; background-color: #000000; width: 100%; filter: alpha(opacity=60); opacity: 0.6; z-index: 2;'></div>
    <div id='module_editor' style='position: absolute; left: 30%; top: 200px; background-color: #fff; border: 1px solid black; display: none; z-index: 3;'>
        <form id='module_form' action='module?id=` + id + `&do=edit' method='POST'>
            <input id='module_edit_type' name='module_edit_type' readonly='readonly' style='display: none;' />
            <input id='module_edit_id' name='module_edit_id' readonly='readonly' style='display: none;' />
            <input id='module_edit_pid' name='module_edit_pid' readonly='readonly' style='display: none;' />
            <span>Module name: <input id='module_edit_name' name='module_edit_name' /></span><br>
            <span>Previous module: 
                <select id='select_previous'>
                    <option id='NONE' value='NONE'>NONE</option>
`
        for k, v := range config.Cfg.Modules {
            result +=
`                    <option id='` + k + `' value='` + v.Name + `'>` + v.Name + `</option>
`
        }
        result +=
`
                </select>
            </span><br>
            <div style='display: inline-block; background-color: blue; cursor: pointer;' onclick='submitModuleEditor("module_editor", "hidebg");'>submit</div>
            <div style='display: inline-block; background-color: pink; cursor: pointer;' onclick='hideModuleEditor("module_editor", "hidebg");'>cancel</div>
        </form>
    </div>
`
    }
    result +=
`    <h3>` + config.Cfg.Modules[id].Name + `</h3>`
    if isLogin {
        result += `    <span>`
        if config.Cfg.Users[loginId].Permissions.CreateModule {
            result += ` <a href='javascript:showModuleEditor("module_editor", "hidebg", "-1", "` + id + `", "", "create");'>Create</a>`
        }
        if config.Cfg.Users[loginId].Permissions.EditModule {
            previousId := config.Cfg.Modules[id].Previous
            result += ` <a href='javascript:showModuleEditor("module_editor", "hidebg", "` + id + `", "` + previousId + `", "` + config.Cfg.Modules[id].Name + `", "edit");'>Edit</a>`
        }
        if config.Cfg.Users[loginId].Permissions.DeleteComment {
            result +=
`    <div style='display: none;'>
        <form id='delete_form' action='module?id=` + id + `&do=delete' method='POST'>
            <input id='delete_id' name='delete_id' readonly='readonly' />
        </form>
    </div>
`
        }
        result += `</span><br>`
    }
    var keysArticles = make([]string, 0)
    for k, _ := range config.Cfg.Articles {
        keysArticles = append(keysArticles, k)
    }
    sort.Sort(ByDemical(keysArticles))
    for _, k := range keysArticles {
        v := config.Cfg.Articles[k]
        if v.ModuleId == id {
            result +=
`    <a href='article?id=` + k + `'><b>` + v.Title + `</b></a>`
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
    var keysModules = make([]string, 0)
    for k, _ := range config.Cfg.Modules {
        keysModules = append(keysModules, k)
    }
    sort.Sort(ByDemical(keysModules))
    for _, k := range keysModules {
        v := config.Cfg.Modules[k]
        if v.Previous == id {
            result +=
`    <a href='module?id=` + k + `'><b>` + v.Name + `</b></a>`
            result += `<span>`
            if isLogin {
                if config.Cfg.Users[loginId].Permissions.EditModule {
                    result += ` <a href='javascript:showModuleEditor("module_editor", "hidebg", "` + k + `", "` + v.Previous + `", "` + v.Name + `", "edit")'>Edit</a>`
                }
                if config.Cfg.Users[loginId].Permissions.DeleteModule {
                    result += ` <a href='javascript:deleteModule("` + k + `");'>Delete</a>`
                }
            }
            result += `</span><br>
`
        }
    }
    return result
}

func renderArticle(id string, isLogin bool, loginId string) string {
    result := ""
    result +=
`    <script>
        function showReplyComment(id, hId, rId, type) {
            var obj = document.getElementById(id);
            document.getElementById(hId).style.display = "block";
            document.getElementById(hId).style.height = document.body.clientHeight + "px";
            obj.style.display = "block";
            document.getElementById('reply_id').value = rId;
            document.getElementById('reply_type').value = type;
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
        function deleteComment(id) {
            var msg = confirm("Will you delete this comment?");
            if (msg) {
                document.getElementById("delete_id").value = id;
                document.getElementById("delete_form").submit();
            }
        }
    </script>
`
    if isLogin {
        if config.Cfg.Users[loginId].Permissions.CreateComment {
            result +=
`    <div id='hidebg' style='position: absolute; left: 0px; top: 0px; background-color: #000000; width: 100%; filter: alpha(opacity=60); opacity: 0.6; z-index: 2;'></div>
    <div id='reply_comment' style='position: absolute; left: 30%; top: 200px; background-color: #fff; border: 1px solid black; display: none; z-index: 3;'>
        <form id='reply_form' action='article?id=` + id + `&do=reply_comment' method='POST'>
            <span><input id='reply_id' name='reply_id' readonly='readonly' style='display: none;' />
            <input id='reply_type' name='reply_type' readonly='readonly' style='display: none;' />
            <div style='display: inline-block; cursor: pointer;' onclick='hideReplyComment("reply_comment", "hidebg");'>X</div></span><br>
            <span>content:</span><br>
            <textarea id='reply_comment_content' name='reply_comment_content' style='resize: none; width: 320px; height: 160px;'></textarea><br>
            <div style='display: inline-block; background-color: blue; cursor: pointer;' onclick='submitReplyComment("reply_comment", "hidebg");'>submit</div>
        </form>
    </div>
`
        }
        if config.Cfg.Users[loginId].Permissions.DeleteComment {
            result +=
`    <div style='display: none;'>
        <form id='delete_form' action='article?id=` + id + `&do=delete_comment' method='POST'>
            <input id='delete_id' name='delete_id' readonly='readonly' />
        </form>
    </div>
`
        }
    }
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
`    <span><a id='main_comment_set' href='javascript:showReplyComment("reply_comment", "hidebg", "-1", "new");'>New comment</a></span><br>
`
        }
    }
    var keys = make([]string, 0)
    for k, _ := range config.Cfg.Comments {
        keys = append(keys, k)
    }
    sort.Sort(ByDemical(keys))
    for _, k := range keys {
        v := config.Cfg.Comments[k]
        if v.BelongsTo == id && v.RepliesTo == "-1" {
            result +=
`    <b>` + config.Cfg.Users[v.AuthorId].Name + ": " + v.Content + " - " + v.DateTime + `</b>`
            if isLogin {
                if config.Cfg.Users[loginId].Permissions.CreateComment {
                    replyId := "reply_" + k
                    result += `<span> <a id='` + replyId + `' href='javascript:showReplyComment("reply_comment", "hidebg", "` + k + `", "reply");'>Reply</a></span>`
                }
                if config.Cfg.Users[loginId].Permissions.DeleteComment {
                    result += `<span> <a href='javascript:deleteComment(` + k + `);'>Delete</a></span>`
                }
            }
            result += `<br>
`
            for _, sk := range keys {
                sub := config.Cfg.Comments[sk]
                if sub.BelongsTo == id && sub.RepliesTo != "-1" {
                    previousId := sk
                    check := false
                    _, ok := config.Cfg.Comments[previousId]
                    if !ok {
                        continue
                    }
                    for config.Cfg.Comments[previousId].RepliesTo != "-1" {
                        previousId = config.Cfg.Comments[previousId].RepliesTo
                    }
                    if !data.CheckCommentExistById(previousId) {
                        continue
                    }
                    check = previousId == k
                    if check {
                        result +=
`    <span>` + config.Cfg.Users[sub.AuthorId].Name + ": " + sub.Content + " - " + sub.DateTime + `</span>`
                        if isLogin {
                            if config.Cfg.Users[loginId].Permissions.CreateComment {
                                replyId := "reply_" + sk
                                result +=
`<span> <a id='` + replyId + `' href='javascript:showReplyComment("reply_comment", "hidebg", "` + sk + `", "reply");'>Reply</a></span>`
                            }
                            if config.Cfg.Users[loginId].Permissions.DeleteComment {
                                result += `<span> <a href='javascript:deleteComment(` + sk + `)'>Delete</a></span>`
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
