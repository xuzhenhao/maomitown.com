package errno

import "fmt"

// Errno 错误码结构体
type Errno struct {
	Code    int
	Message string
}

// 实现error定义的接口
func (err Errno) Error() string {
	return err.Message
}

// Err 代表一个完整的错误
type Err struct {
	Code    int
	Message string
	Err     error
}

// 实现 error 定义的接口
func (err *Err) Error() string {
	return fmt.Sprintf("Err -code: %d, message: %s,error:%s", err.Code, err.Message, err.Err)
}

// New 新建定制的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// DecodeErr 解析定制的错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}
	return InternalServerError.Code, err.Error()
}

// Add 添加一条信息
func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

// Addf 添加一条格式化信息
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}
