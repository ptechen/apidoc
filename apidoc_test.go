package apidoc

import (
	"os"
	"testing"
)

type ReqParams struct {
	Id    int    `json:"id" comment:"test"`
	Name  string `json:"name" comment:"test"`
	inner string
	Age   int `json:"age" comment:"test"`
	Xi    struct {
		Name string `json:"name" comment:"test"`
	} `json:"xi" comment:"test"`
}

type ResParams struct {
	Id    int    `json:"id" comment:"test"`
	Name  string `json:"name" comment:"test"`
	inner string
	Age   int `json:"age" comment:"test"`
	Xi    struct {
		Name string `json:"name" comment:"test"`
		Info struct {
			Age int `json:"age" comment:"age"`
		} `json:"info" comment:"user_info"`
	} `json:"xi" comment:"test"`
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
	api.ApiDoc(req, res)
}
