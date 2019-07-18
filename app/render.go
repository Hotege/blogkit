package app

import (
    "strconv"
    "html/template"
)

var renders map[string]*template.Template

func initializeRenders() {
    renders = make(map[string]*template.Template)
    renders["error"], _ = template.ParseFiles("static/templates/error.html")
    renders["initialize"], _ = template.ParseFiles("static/templates/initialize.html")
    renders["signup"], _ = template.ParseFiles("static/templates/signup.html")
    renders["page"], _ = template.ParseFiles(
        "static/templates/page.html",
        "static/templates/user.html",
        "static/templates/module.html",
    )
    renders["article"], _ = template.ParseFiles(
        "static/templates/blog.html",
        "static/templates/user.html",
        "static/templates/article.html",
    )
    renders["create"], _ = template.ParseFiles(
        "static/templates/create.html",
        "static/templates/user.html",
        "static/templates/dashboard.html",
    )
    renders["files"], _ = template.ParseFiles("static/templates/files.html")
}

type byDemical []string

func (a byDemical) Len() int {
    return len(a)
}

func (a byDemical) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a byDemical) Less(i, j int) bool {
    x, _ := strconv.Atoi(a[i])
    y, _ := strconv.Atoi(a[j])
    return x < y 
}
