package integer

type IntNumber interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
	Increase(i Integer)
	Decrease(i Integer)
}
type IntNumber2 interface {
	Equal(i Integer) bool
	LessThan(i Integer) bool
	MoreThan(i Integer) bool
}

//在 Go 语言中，你可以给任意类型（包括基本类型，但不包括指针类型）添加成员方法，但是如果是基本类型的话，需要借助 type 关键字对类型进行再定义，例如：
type Integer int // 对int 类型进行扩充

func (a Integer) Equal(b Integer) bool { // 扩充了 int类型  并为int 类型扩展了一个新方法
	return a == b
}

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

func (a Integer) MoreThan(b Integer) bool {
	return a > b
}

func (a *Integer) Increase(i Integer) {
	*a = *a + i
}

func (a *Integer) Decrease(i Integer) {
	*a = *a - i
}
