package errno

/*
这里定义的错误码是给前端展示的。通过errno.DecodeErr方法将后端内部的错误信息转化为这些通俗易懂的提示语。
错误代码说明:
   1			  00 			  02
服务级错误		服务模块代码      具体错误代码

服务级别错误:1为系统错误，2为普通错误，通常是用户非法操作引起的

*/

var (
	//通用错误
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "服务器内部错误"}
	ParamBindError      = &Errno{Code: 10002, Message: "绑定请求的参数错误"}

	//用户相关错误
	UserNotFoundError = &Errno{Code: 20102, Message: "未找到用户"}
)
