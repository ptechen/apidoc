package apidoc

const tpl = `/**
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
 */

`

var apiDocType = map[string]string{
	"uint8":     "Number",
	"uint16":    "Number",
	"uint32":    "Number",
	"uint64":    "Number",
	"uint":      "Number",
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
	"float32":   "Float",
	"float64":   "Float",
	"struct":    "Object",

	"[]uint8":     "Number[]",
	"[]uint16":    "Number[]",
	"[]uint32":    "Number[]",
	"[]uint64":    "Number[]",
	"[]int":       "Number[]",
	"[]int8":      "Number[]",
	"[]int16":     "Number[]",
	"[]int32":     "Number[]",
	"[]int64":     "Number[]",
	"[]bool":      "Boolean[]",
	"[]string":    "String[]",
	"[]time.Time": "time[]",
	"[]float32":   "Float[]",
	"[]float64":   "Float[]",
	"[]struct":    "Object[]",

	"map[int]int":       "Map[Number]Int",
	"map[int]string":    "Map[Number]String",
	"map[string]int":    "Map[String]Int",
	"map[string]string": "Map[String]String",
	"map[int8]int8":     "Map[Number]Int",
	"map[int8]string":   "Map[Number]String",
	"map[string]int8":   "Map[String]Int",

	"map[int16]int16":  "Map[Number]Int",
	"map[int16]string": "Map[Number]String",
	"map[string]int16": "Map[String]Int",

	"map[int32]int32":  "Map[Number]Int",
	"map[int32]string": "Map[Number]String",
	"map[string]int32": "Map[String]Int",

	"map[int64]int64":  "Map[Number]Int",
	"map[int64]string": "Map[Number]String",
	"map[string]int64": "Map[String]Int",
}
