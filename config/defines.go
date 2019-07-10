package config

var Cfg Config

type Config struct {
    Version string `json:Version`
    Users []User `json:Users`
    Modules []Module `json:Modules`
    Articles []Article `json:Articles`
    Comments []Comment `json:Comments`
}

type User struct {
    Id int `json:ID`
    Mail string `json:Mail`
    Name string `json:Name`
    Token string `json:Token`
    Permissions Permission `json:Permissions`
}

type Permission struct {
    CreateUser bool `json:CreateUser`
    EditUser bool `json:EditUser`
    DeleteUser bool `json:DeleteUser`
    CreateModule bool `json:CreateModule`
    EditModule bool `json:EditModule`
    DeleteModule bool `json:DeleteModule`
    CreateArticle bool `json:CreateArticle`
    EditArticle bool `json:EditArticle`
    DeleteArticle bool `json:DeleteArticle`
    CreateComment bool `json:CreateComment`
    EditComment bool `json:EditComment`
    DeleteComment bool `json:DeleteComment`
}

type Module struct {
    Id int `json:ID`
    Name string `json:Name`
    Previous int `json:Previous`
}

type Article struct {
    Id int `json:ID`
    Path string `json:Path`
    Title string `json:Title`
    AuthorId int `json:AuthorID`
    DateTime string `json:DateTime`
    ModuleId int `json:ModuleID`
}

type Comment struct {
    Id int `json:ID`
    Content string `json:Content`
    BelongsTo int `json:BelongsTo`
    RepliesTo int `json:RepliesTo`
    AuthorId int `json:AuthorID`
    DateTime string `json:DateTime`
}
