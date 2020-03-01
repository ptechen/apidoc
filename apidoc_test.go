package apidoc

import (
	"os"
	"testing"
)

type ReqParams struct {
	String string `json:"string" comment:"test string"`
	Int    int    `json:"int" comment:"test int"`
	Int8   int8   `json:"int_8" comment:"test int8"`
	Int16  int16  `json:"int_16" comment:"test int16"`
	Int32  int32  `json:"int_32" comment:"test int32"`
	Int64  int64  `json:"int_64" comment:"test int64"`

	Uint   uint   `json:"uint" comment:"test uint"`
	Uint8  uint8  `json:"uint_8" comment:"test uint8"`
	Uint16 uint16 `json:"uint_16" comment:"test uint16"`
	Uint32 uint32 `json:"uint_32" comment:"test uint32"`
	Uint64 uint64 `json:"uint_64" comment:"test uint64"`

	Bool         bool              `json:"bool" comment:"test bool"`
	Float64      float64           `json:"float_64" comment:"test float64"`
	Float32      float32           `json:"float_32" comment:"test float32"`
	MapString    map[string]string `json:"map_string" comment:"test map[string]string"`
	MapInt       map[int]int       `json:"map_int" comment:"test map[int]int"`
	MapIntString map[int]string    `json:"map_int_string" comment:"test map[int]string"`
	MapStringInt map[string]int    `json:"map_string_int" comment:"test map[string]int"`
}

type ResParams struct {
	String string `json:"string" comment:"test string"`
	Int    int    `json:"int" comment:"test int"`
	Int8   int8   `json:"int_8" comment:"test int8"`
	Int16  int16  `json:"int_16" comment:"test int16"`
	Int32  int32  `json:"int_32" comment:"test int32"`
	Int64  int64  `json:"int_64" comment:"test int64"`

	Uint   uint   `json:"uint" comment:"test uint"`
	Uint8  uint8  `json:"uint_8" comment:"test uint8"`
	Uint16 uint16 `json:"uint_16" comment:"test uint16"`
	Uint32 uint32 `json:"uint_32" comment:"test uint32"`
	Uint64 uint64 `json:"uint_64" comment:"test uint64"`

	Bool         bool              `json:"bool" comment:"test bool"`
	Float64      float64           `json:"float_64" comment:"test float64"`
	Float32      float32           `json:"float_32" comment:"test float32"`
	MapString    map[string]string `json:"map_string" comment:"test map[string]string"`
	MapInt       map[int]int       `json:"map_int" comment:"test map[int]int"`
	MapIntString map[int]string    `json:"map_int_string" comment:"test map[int]string"`
	MapStringInt map[string]int    `json:"map_string_int" comment:"test map[string]int"`
}

func TestApiDoc(t *testing.T) {
	req := ReqParams{}
	res := ResParams{}
	f, _ := os.Create("./apidoc/apidoc")
	api := &Api{
		Output:   f,
		Method:   "post",
		Route:    "/test1",
		Desc:     "123",
		Version:  "v0.0.1",
		ApiName:  "user",
		ApiGroup: "user",
	}
	api.ApiDoc(req, res)
}

type ReqParams1 struct {
	Req ReqParams `json:"req"`
}

type ResParams1 struct {
	Res ResParams `json:"res"`
}

func TestApiDoc1(t *testing.T) {
	req := ReqParams1{}
	res := ResParams1{}
	f, _ := os.Create("./apidoc/apidoc1")
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

type ReqParams2 struct {
	Req *ReqParams `json:"req"`
}

type ResParams2 struct {
	Res *ResParams `json:"res"`
}

func TestApiDoc2(t *testing.T) {
	req := ReqParams2{Req: &ReqParams{}}
	res := ResParams2{Res: &ResParams{}}
	f, _ := os.Create("./apidoc/apidoc2")
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

type ReqParams3 struct {
	ReqParams
}

type ResParams3 struct {
	ResParams
}

func TestApiDoc3(t *testing.T) {
	req := ReqParams3{}
	res := ResParams3{}
	f, _ := os.Create("./apidoc/apidoc3")
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

type ReqParams4 struct {
	Req []ReqParams
}

type ResParams4 struct {
	Res []ResParams
}

func TestApiDoc4(t *testing.T) {
	req := ReqParams4{Req: []ReqParams{}}
	res := ResParams4{Res: []ResParams{}}
	f, _ := os.Create("./apidoc/apidoc4")
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
