package internal

const (
	REQUEST_DATA_ILLEGAL = "请求参数不合法"
)

type Err struct {
	Code    int
	Message string
	Result  string
}

func (e *Err) Error() string {
	return e.Message
}

var (
	ErrReqData = &Err{
		Code: REQUEST,
		Message: "请求参数不合法，无法正确解析",
		Result: "{}",
	}
)


// var (
// 	ErrBind = &Err{
// 		Code:       "BIND_ERR",
// 		Message:    "Error occurred while binding the request body to the struct.",
// 		StatusCode: http.StatusBadRequest}
// )