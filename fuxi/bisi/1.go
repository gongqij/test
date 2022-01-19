package main

type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//
//func main() {
//	//指针类型的结构体对象可以同时调用结构体值类型和指针类型对应的方法。而值类型的结构体对象只能调用值类型对应的接口方法。
//	var peo People = Stduent{}
//	think := "bitch"
//	fmt.Println(peo.Speak(think))
//}
