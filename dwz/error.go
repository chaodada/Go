package main

// 自定义错误处理接口Error
type Error interface {
	error        // go 系统错误接口
	Status() int //自定义方法
}

// 定义一个类
type StatusError struct {
	Code int   // 错误码
	Err  error // 继承系统错误接口
}

// 实现系统错误处理接口 error 的 Error方法
func (se StatusError) Error() string {
	return se.Err.Error()
}

// 实现自定义错误处理接口 的 Status 方法
func (se StatusError) Status() int {
	return se.Code
}
