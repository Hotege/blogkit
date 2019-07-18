package data

type articleSlice struct {
    Type string `json:Type`
    Value string `json:Value`
    Extra interface{} `json:Extra`
}
