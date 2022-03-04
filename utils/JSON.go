package utils

type JsonResponse struct {
	Code string `json:"code"`
	// key不能是interface{}  不然json无法处理报错
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}
