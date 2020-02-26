package apidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"reflect"
	"strings"
	"unicode"
)

var apiDocType = map[string]string{
	"uint8":     "Number",
	"uint16":    "Number",
	"uint32":    "Number",
	"uint64":    "Number",
	"int":       "Number",
	"int8":      "Number",
	"int16":     "Number",
	"int32":     "Number",
	"int64":     "Number",
	"*uint8":    "Number",
	"*uint16":   "Number",
	"*uint32":   "Number",
	"*uint64":   "Number",
	"*int":      "Number",
	"*int8":     "Number",
	"*int16":    "Number",
	"*int32":    "Number",
	"*int64":    "Number",
	"bool":      "Boolean",
	"string":    "String",
	"time.Time": "String",
	"float32":   "Float32",
	"float64":   "Float64",
	"struct":    "Object",
}

const temp = `/**
 * @api { {{.Method}} } {{.Route}}   {{.Desc}}
 * @apiVersion {{.Version}}
 * @apiName {{.ApiName}}
 * @apiGroup {{.ApiGroup}} {{range .ApiParams}}
 * @apiParam { {{.FieldType}} } {{ add }} {{ .Json }} {{ add }} {{ .Comment }} {{end}}
 * @apiParamExample {json} Request-Example:{{range .ApiParamExample}}
 * {{unescaped .}} {{end}}
 * {{range .ApiSuccess}}
 * @apiSuccess { {{.FieldType}} } {{ add }} {{ .Json }} {{ add }} {{.Comment}}{{end}}
 * @apiSuccessExample {json} Request-Example:{{range .ApiSuccessExample}}
 * {{unescaped .}} {{end}}
 */`

type Api struct {
	Output            io.Writer
	Method            string
	Route             string
	Desc              string
	Version           string
	ApiName           string
	ApiGroup          string
	ApiParams         []ApiFieldInfo
	ApiParamExample   []string
	ApiSuccess        []ApiFieldInfo
	ApiSuccessExample []string
}

type ApiFieldInfo struct {
	FieldType string
	Json      string
	Key       int
	Comment   string
}

func ApiDoc(api *Api, reqParams, resParams interface{}) {
	reqFieldInfo := make([]ApiFieldInfo, 0, 20)
	resFieldInfo := make([]ApiFieldInfo, 0, 20)
	reqObjectMap := make(map[int]string)
	resObjectMap := make(map[int]string)
	reqKey := 0
	resKey := 0
	myJsonEncode(reqParams, reqKey, &reqFieldInfo, reqObjectMap)
	myJsonEncode(resParams, resKey, &resFieldInfo, resObjectMap)
	api.ApiParams = reqFieldInfo
	api.ApiSuccess = resFieldInfo
	reqString, _ := paramsString(reqParams)
	resString, _ := paramsString(resParams)
	api.ApiParamExample = reqString
	api.ApiSuccessExample = resString

	funcMap := template.FuncMap{"add": add, "unescaped": unescaped}
	t := template.Must(template.New("template.tpl").Funcs(funcMap).ParseFiles("./template.tpl"))
	err := t.Execute(api.Output, api)
	if err != nil {
		log.Println("executing template:", err)
	}
}

func myJsonEncode(obj interface{}, key int, fieldInfo *[]ApiFieldInfo, objectMap map[int]string) {
	var (
		i int

		objType  reflect.Type
		objValue reflect.Value

		field      reflect.StructField
		fieldValue reflect.Value

		fieldName string
	)

	// 接口是空(没装任何东西的interface{})
	if obj == nil {
		//fmt.Println("空接口")
		return
	}

	// 反射变量
	objType = reflect.TypeOf(obj)   // 反射类型
	objValue = reflect.ValueOf(obj) // 反射值
	// 如果是指针, 需要取值
	if objType.Kind() == reflect.Ptr {
		if objValue.IsNil() { // 空指针
			//fmt.Println("空指针")
			return
		}
		objType = objType.Elem()   // 相当于类型为*ptr
		objValue = objValue.Elem() // 相当于值为*ptr
	}

	// 如果不是结构体, 则不需要递归处理
	if objType.Kind() != reflect.Struct {
		//fmt.Println("普通值", objValue.Interface())
		return
	}

	// 递归处理结构体中的字段
	key = key + 1
	for i = 0; i < objType.NumField(); i++ {
		field = objType.Field(i)       // 获取字段类型
		fieldValue = objValue.Field(i) // 获取字段的值

		// 小写字段不导出
		fieldName = field.Name
		if unicode.IsLower(rune(fieldName[0])) {
			continue
		}

		// 打印这个字段的信息
		fieldType := fmt.Sprintf(" %v", field.Type)
		fieldType = strings.Trim(fieldType, " ")
		jsonTag := field.Tag.Get("json")
		if key > 1 {
			for i := key; i > key-1; i-- {
				jsonTag = objectMap[key-1] + "." + jsonTag
			}
		}
		comment := field.Tag.Get("comment")
		if strings.Contains(fieldType, "struct") ||
			strings.Contains(fieldType, ".") {
			objectMap[key] = jsonTag
			fieldType = "struct"
		}
		//fmt.Println("Field", field.Name, "类型:", field.Type, "标签:", field.Tag)
		*fieldInfo = append(*fieldInfo, ApiFieldInfo{
			FieldType: apiDocType[fieldType],
			Json:      jsonTag,
			Comment:   comment,
			Key:       key,
		})
		//递归编码这个字段
		myJsonEncode(fieldValue.Interface(), key, fieldInfo, objectMap)
	}
}

func add() string {
	return "\t"
}

func unescaped(str string) template.HTML { return template.HTML(str) }

func paramsString(params interface{}) (res []string, err error) {

	b, err := json.Marshal(params)
	if err != nil {
		return res, fmt.Errorf("%+v", params)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return res, fmt.Errorf("%+v", params)
	}
	content := out.String()
	data := strings.Split(content, "\n")
	res = make([]string, 0, len(data))
	for i := 0; i < len(data); i++ {
		res = append(res, data[i])
	}
	return res, err
}
