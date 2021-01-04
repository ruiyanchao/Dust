package main

// 匿名函数，具名函数。
func main(){

}

//具名函数
func addFunc(a,b int64) int64{
	return a+b
}

//匿名函数
var add = func(a ,b int64) int64 {
	return a+b
}

// 多个参数和多个返回值
func func1()(int , int){
	return 1,2
}

// 多个参数和多个返回值 (预先定义好)
func func3()(a int , b int){
	a = 1
	b = 2
	return a,b
}

//可变数量的参数  ... 代表解包
func func2(a int, more ...int)int{
	for _, v := range more {
		a += v
	}
	return a
}

// defer 延迟执行一个匿名函数 （在函数执行完之后，才会执行）
func func4() (v int) {
	defer func(){ v++ } ()
	return 42
}

/**
 * 方法
 */
// 文件对象
type File struct {
	fd int
}

// 打开文件
func OpenFile(name string) (f *File, err error) {
	// ...
	return
}

// 关闭文件
func CloseFile(f *File) (err error) {
	// ...
	return
}

// 读文件数据
func ReadFile(f *File, offset int64, data []byte) (fd int) {
	// ...
	return
}

/**
 * 接口  + 继承
 */
type A interface {
	run()
}

type B struct {
	name string
}

func(b *B)run(){

}

type C struct {
	B
}

func(c *C)do(){
	c.run()
}




