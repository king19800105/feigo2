package internal

// 0表示成功，非0表示失败
// 101 ~ 199 为通用的系统错误码
const (
	SUCCESS = 0 // 成功
	REQUEST = 101 + iota // 请求参数不合法
	PARSE    // 数据解析错误
	_
	_
	INNER    // 内部服务错误
	ACCOUNT  // 账号信息错误
)
