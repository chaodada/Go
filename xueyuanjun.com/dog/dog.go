package dog

import "xueyuanjun.com/animal"

type Dog struct {
	Name *animal.Animal
}

func (d Dog) FavorFood() string {
	return "骨头"
}

func (d Dog) Call() string {
	return "汪汪汪"
}
