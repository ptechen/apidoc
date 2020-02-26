/**
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

