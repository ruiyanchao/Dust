package main


// 类似于Mysql的自增ID

// 首选 Twitter 的  snowflake算法

// 首选我们确认数值是64位，int64类型。分为4个部分（不含开头的第一个bit）

// 41位表示收到的请求时的时间戳（毫秒）  +   5位数据中心的ID（业务相关）  +   5位机器实例ID  +  12位循环自增ID（位数满了后归0）

// 同一个毫秒  4096 个 够用了

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main(){
	n, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i:=0 ; i< 3;i++{
		id := n.Generate()
		fmt.Println("id",id)
	}


}

// 另外一个 github.com/sony/sonyflake  索尼公司的一个开源项目