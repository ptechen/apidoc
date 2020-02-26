package apidoc

import (
	"os"
	"testing"
)

type ReqParams struct {
	Id    int    `json:"id" comment:"fsfsfs"`
	Name  string `json:"name" comment:"fsfsfs"`
	inner string
	Age   int `json:"age" comment:"fsfsfs"`
	Xi    struct {
		Name string `json:"name" comment:"fsfsfs"`
	} `json:"xi" comment:"fsfsfs"`
}

type ResParams struct {
	Id    int    `json:"id" comment:"fsfsfs"`
	Name  string `json:"name" comment:"fsfsfs"`
	inner string
	Age   int `json:"age" comment:"fsfsfs"`
	Xi    struct {
		Name string `json:"name" comment:"fsfsfs"`
		Info struct {
			Age int `json:"age" comment:"age"`
		} `json:"info" comment:"user_info"`
	} `json:"xi" comment:"fsfsfs"`
}

func TestApiDoc(t *testing.T) {
	req := ReqParams{}
	res := ResParams{}
	f, _ := os.Create("./apidoc/apidoc")
	api := &Api{
		Output:   f,
		Method:   "post",
		Route:    "/test",
		Desc:     "123",
		Version:  "v0.0.1",
		ApiName:  "user",
		ApiGroup: "user",
	}
	ApiDoc(api, req, res)
}
