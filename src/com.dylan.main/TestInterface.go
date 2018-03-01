package main
/*

注意：Go语言支持的除Interface类型外的任何其它数据类型都可以定义其method（而并非只有struct才支持method），
		只不过实际项目中，method(s)多定义在struct上而已。

1.接口是类型 type
2.是一堆方法的集合
//总结 接口是拥有一堆方法的集合 类型
*/
// 如果


func main() {
p:=new(person)
p.sporter();

p.sporter()

}

// 结构体
type person struct{

}
//跑的接口
type runType interface{
	firstrun()

}
//跑的接口
type smileType interface{
	firstsmlie()
}


//
func  (* person) firstrun(){

}
func  (* person) firstsmlie(){

}


func  (* person) thief(){

}
func (*person) sporter()  {

}