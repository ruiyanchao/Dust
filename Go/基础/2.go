package main

import (
	"fmt"
)

func main(){
	fmt.Println("show array")
	showArr()
	fmt.Println("show string")
	showStr()
	fmt.Println("show slice")
	showSlice()
}

func showArr(){
	//数组 数据结构（ [n]TYPE ）上讲 固定长度，连续空间，类型一致
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6   (index:param)
	//定义空数组  空数组不占用空间 可用于管道的同步操作 （ 做好的还是 nocopy ）
	var g [0]int       // 定义一个长度为0的数组
	var h = [0]int{}   // 定义一个长度为0的数组
	var i = [...]int{} // 定义一个长度为0的数组
	fmt.Println(a,b,c,d,g,h,i)

	//当数组很大时 数组的赋值会带来很大的开销，避免这样的开销可以采用指针 数组的值是整个数组 而 数组的指针并不是数组
	var e = [...]int{1, 2, 3} // a 是一个数组
	var f = &e                // b 是指向数组的指针

	fmt.Println(e[0], e[1])   // 打印数组的前2个元素
	fmt.Println(f[0], f[1])   // 通过数组指针访问数组元素的方式和数组类似
	e[0] = 4                  // 改变e 也改变了 f
	fmt.Println(e, f)

	//变量数组 用for range遍历  len 返回数组长度 cap 返回数组容量
	for range a{
		fmt.Println("i get it")
	}
	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range a {
		fmt.Printf("a[%d]: %d\n", i, v)
	}
	fmt.Println(len(a),cap(a))
}

func showStr(){
	//字符串 只读的字节数组 （只读的二进制数组）
	data := [...]byte{
		'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',
	}
	str := "hello, world"
	fmt.Println(data)
	fmt.Println([]byte(str))
	//字符串支持切片操作
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]
	fmt.Println(hello,world,s1,s2)
	//也支持len
	fmt.Println(len(str))
	// go文件都是utf8编码的字符序列 可以用for range循环直接遍历UTF8解码后的Unicode码点值
	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界

}

func showSlice(){
	//切片 动态数组 相对于数组更灵活
	var (
		a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
		d = c[:2]             // 有2个元素的切片, len为2, cap为3
		e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
		f = c[:0]             // 有0个元素的切片, len为0, cap为3
		g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)
	printSlice(a)
	printSlice(b)
	printSlice(c)
	printSlice(d)
	printSlice(e)
	printSlice(f)
	printSlice(g)
	printSlice(h)
	printSlice(i)
	//for range 打印
	//元素追加 append
	//尾部追加N个元素
	fmt.Println(a)
	a = append(a,1)
	fmt.Println(a)
	a = append(a,1,2)
	fmt.Println(a)
	a = append(a,[]int{6,7,8}...)
	fmt.Println(a)
	//在容量不足的情况下，append的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。即使容量足够，依然需要用append函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。
    //头部追加
    fmt.Println(a)
	a = append([]int{7},a...)
	fmt.Println(a)
	//中间插入
	a = append(a[:4], append([]int{8}, a[4:]...)...) //第 4 位插入 8
	fmt.Println(a)
	//可以用copy和append组合可以避免创建中间的临时切片，同样是完成添加元素的操作
	a = append(a, 0)     // 切片扩展1个空间
	fmt.Println(a)
	copy(a[len(a):], a[len(a)-1:]) // a[i:]向后移动1个位置
	a[len(a)-1] = 9             // 设置新添加的元素
	fmt.Println(a)
	//删除尾部N个数组 a = a[:len(a)-N]
	a = a[:len(a)-1]
    //切片内存溢出问题
    //（数据的传值是Go语言编程的一个哲学，虽然传值有一定的代价，但是换取的好处是切断了对原始数据的依赖）
	//切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这会延迟自动内存回收器对底层数组的回收。
}

func printSlice(a []int){
	fmt.Println(a,len(a),cap(a))
}