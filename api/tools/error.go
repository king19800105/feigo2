package tools

// type Err struct {
// 	Code    int
// 	Message string
// 	Result  string
// }
//
// func (e *Err) Error() string {
// 	return e.Message
// }

// // Fill the error struct with the detail error information.
// func Fill(err *Err) *Err {
// 	_, fn, line, _ := runtime.Caller(1)
//
// 	// Fill the error occurred path, line, code.
// 	err.Path = strings.Replace(fn, os.Getenv("GOPATH"), "", -1)
// 	err.Line = line
// 	return err
// }

// Abort the current request with the specified error code.
// func Abort(errStruct *Err, err error, c *gin.Context) {
// 	c.Error(err)
// 	c.Error(Fill(errStruct))
// 	c.Abort()
// }
