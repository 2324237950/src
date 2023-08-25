package test

import "fmt"

type Liu struct {
	a int
}

func Newliu() *Liu {
	return &Liu{
		a: 1,
	}
}
func (l *Liu) Mymethod() {
	fmt.Print("mymethod被使用1111111111111111111111111111")
}

//注意如果要在别的，被调用的结构体和方法必须是公开的（即首字母大写），
//以便其他包能够访问和调用它们。如果结构体或方法是私有的，则只能在其所属的包中访问和使用。
