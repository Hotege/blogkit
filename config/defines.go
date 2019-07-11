package config

var Cfg Config

type Config struct {
    Version string `json:Version`
    Users map[string]*User `json:Users`
    Modules map[string]*Module `json:Modules`
    Articles map[string]*Article `json:Articles`
    Comments map[string]*Comment `json:Comments`
}

type User struct {
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
    Name string `json:Name`
    Previous string `json:Previous`
}

type Article struct {
    Path string `json:Path`
    Title string `json:Title`
    AuthorId string `json:AuthorID`
    DateTime string `json:DateTime`
    ModuleId string `json:ModuleID`
}

type Comment struct {
    Content string `json:Content`
    BelongsTo string `json:BelongsTo`
    RepliesTo string `json:RepliesTo`
    AuthorId string `json:AuthorID`
    DateTime string `json:DateTime`
}
