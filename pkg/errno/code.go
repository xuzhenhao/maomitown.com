package errno

/*
这里定义的错误码是给前端展示的。通过errno.DecodeErr方法将后端内部的错误信息转化为这些通俗易懂的提示语。
错误代码说明:
   1			  00 			  02
服务级错误		服务模块代码      具体错误代码

服务级别错误:1为系统错误，2为普通错误，通常是用户非法操作引起的

*/

var (
	//通用错误，服务代码 00
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "服务器内部错误"}
	ParamBindError      = &Errno{Code: 10002, Message: "绑定请求的参数错误"}

	ValidationError = &Errno{Code: 20001, Message: "参数校验失败"}
	DatabaseError   = &Errno{Code: 20002, Message: "数据库出错"}

	//用户模块,服务代码 01
	EncryptPwdError   = &Errno{Code: 20101, Message: "加密密码出错"}
	UserNotFoundError = &Errno{Code: 20102, Message: "未找到用户"}
	TokenInvalidError = &Errno{Code: 20103, Message: "无效Token"}
	PwdIncorrectError = &Errno{Code: 20104, Message: "密码错误"}
)
