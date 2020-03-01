package apidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"reflect"
	"strings"
	"sync"
	"unicode"
)

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

func (api *Api) ApiDoc(reqParams, resParams interface{}) error {
	reqFieldInfo := make([]ApiFieldInfo, 0, 20)
	resFieldInfo := make([]ApiFieldInfo, 0, 20)
	reqObjectMap := make(map[int]string)
	resObjectMap := make(map[int]string)
	var reqString []string
	var resString []string
	reqKey := 0
	resKey := 0
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		myJsonEncode(reqParams, reqKey, &reqFieldInfo, reqObjectMap)
		wg.Done()
	}()
	go func() {
		myJsonEncode(resParams, resKey, &resFieldInfo, resObjectMap)
		wg.Done()
	}()

	go func() {
		reqString = paramsString(reqParams)
		wg.Done()
	}()
	go func() {
		resString = paramsString(resParams)
		wg.Done()
	}()

	wg.Wait()
	api.ApiParams = reqFieldInfo
	api.ApiSuccess = resFieldInfo
	api.ApiParamExample = reqString
	api.ApiSuccessExample = resString

	funcMap := template.FuncMap{"add": add, "unescaped": unescaped}
	t := template.Must(template.New("template.tpl").Funcs(funcMap).Parse(tpl))
	err := t.Execute(api.Output, api)
	return err
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
		return
	}

	// 反射变量
	objType = reflect.TypeOf(obj)   // 反射类型
	objValue = reflect.ValueOf(obj) // 反射值
	// 如果是指针, 需要取值
	if objType.Kind() == reflect.Ptr {
		if objValue.IsNil() {
			return // 空指针
		}

		objType = objType.Elem()   // 相当于类型为*ptr
		objValue = objValue.Elem() // 相当于值为*ptr
	}

	// 如果不是结构体, 则不需要递归处理
	if objType.Kind() != reflect.Struct {
		//if  objType.Kind() == reflect.Slice {
		//	objValue = reflect.ValueOf(obj).Elem()
		//	objType = reflect.TypeOf(objValue)
		//} else {
		//	//fmt.Println("普通值", objValue.Interface())
		//	return
		//}
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
				if objectMap[key-1] != "" {
					jsonTag = objectMap[key-1] + "." + jsonTag
				}
			}
		}
		comment := field.Tag.Get("comment")
		if strings.Contains(fieldType, "[]") && strings.Contains(fieldType, ".") {
			fieldType = "[]struct"
		} else if strings.Contains(fieldType, "*") && strings.Contains(fieldType, ".") {
			fieldType = "struct"
			objectMap[key] = jsonTag
		} else if strings.Contains(fieldType, "struct") && strings.Contains(fieldType, ".") {
			fieldType = "struct"
			objectMap[key] = jsonTag
		} else if strings.Contains(fieldType, ".") {
			fieldType = "struct"
			objectMap[key] = jsonTag
		} else {

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

func paramsString(params interface{}) (res []string) {

	b, _ := json.Marshal(params)
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "    ")
	content := strings.ReplaceAll(out.String(), "null", "{}")
	data := strings.Split(content, "\n")
	res = make([]string, 0, len(data))
	for i := 0; i < len(data); i++ {
		res = append(res, data[i])
	}
	return res
}
