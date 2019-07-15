package data

import (
    "strings"
    "encoding/json"
)

func EncodeArticle(steps []string, texts []string, images []string, files []string, codes []string) ([]byte, error) {
    slices := make([]articleSlice, 0)
    for k, v := range steps {
        if v == "step_ot" {
            var s articleSlice
            s.Type = "t"
            s.Value = texts[k]
            slices = append(slices, s)
        }
        if v == "step_oi" {
            var s articleSlice
            s.Type = "i"
            s.Value = "/static/files/" + images[k]
            slices = append(slices, s)
        }
        if v == "step_of" {
            fs := strings.Split(files[k], "/")
            for _, f := range fs {
                var s articleSlice
                s.Type = "f"
                s.Value = "/static/files/" + f
                slices = append(slices, s)
            }
        }
        if v == "step_oc" {
            var s articleSlice
            s.Type = "c"
            s.Value = codes[k]
            slices = append(slices, s)
        }
    }
    return json.Marshal(slices)
}
