package response

// 成功返回值
const successCode = 0

// Result 基础序列化器
type Result struct {
	Code int32       `form:"code" json:"code" query:"code"`
	Msg  string      ` form:"msg" json:"msg" query:"msg"`
	Data interface{} `form:"data" json:"data" query:"data"`
}

// List 基础列表结构
type list struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// ListResult 列表构建器
func ListResult(items interface{}, total uint) (result Result) {
	if items == nil {
		items = make([]int, 0)
	}
	result = Success(
		list{
			Items: items,
			Total: total,
		})
	return
}

func Success(data interface{}) (result Result) {
	result.Data = data
	result.Code = successCode
	return
}

func Failed(code int32, msg string) (result Result) {
	result.Msg = msg
	result.Code = code
	return
}
