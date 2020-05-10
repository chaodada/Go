package student

import "fmt"

type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// 初始化函数  构造方法
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

// 注：我们可以把接收者类型为指针的成员方法叫做指针方法，把接收者类型为非指针的成员方法叫做值方法。

// 非指针方法 定义类方法 获取姓名
func (s Student) GetName() string {
	return s.name
}

// 指针方法 定义类方法 设置姓名
func (s *Student) SetName(name string) {
	s.name = name
}

//PHP、Java 支持默认调用类的 toString 方法以字符串格式打印类的实例，Go 语言也有类似的机制，只不过这个方法名是 String，以上面这个 Student 类型为例，我们为其编写 String 方法如下：
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}
