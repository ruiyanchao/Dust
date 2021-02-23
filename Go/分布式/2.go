package main

import (
	"fmt"
	"sync"
	"github.com/go-redis/redis"
	"time"
)

// 在单机程序并发或并行修改全局变量时，需要对修改的行为加锁
var counter int

func main(){
	noLock() // 每次结果不一样
	counter = 0
	hasLock()  // 结果正确
	// 正常运用中 我们可以使用 try lock (尝试加锁 加锁成功执行后续流程，失败也不会阻塞)
	counter = 0
	tryLock()
	// 单机中这不是一个好方案，大量的go抢锁导致cpu的无意义的资源浪费。"活锁"
	// 分布式的场景下 也有 像单机一样的问题。我们可以使用 redis 的 setnx
	redisLock()
}

//单机不加锁的情况下
func noLock(){
	var wg sync.WaitGroup
	for i := 0; i<1000;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("sign no lock:",counter)
}

//单机内加锁
func hasLock(){
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i<1000;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("sign lock:",counter)
}

//尝试加锁
func tryLock(){
	var wg sync.WaitGroup
	var l = NewLock()
	for i := 0; i<10;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock(){
				fmt.Println("lock fail")
				return
			}
			counter++
			fmt.Println("current counter",counter)
			l.UnLock()
		}()
	}
	wg.Wait()
}

// try lock
type Lock struct {
	c chan struct{}
}

func NewLock() Lock{
	var l Lock
	l.c = make(chan struct{},1)
	l.c <- struct{}{}
	return l
}

func (l Lock)Lock()bool{
	lockRe := false
	select {
	case <-l.c:
		lockRe = true
	default:

	}
	return lockRe

}

func (l Lock)UnLock(){
	l.c<- struct{}{}
}

// redis setnx 实现
func incr(){
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "jch9shshl",
		DB: 0,
	})
	var lockKey = "count_lock"
	var counterkey = "counter"

	//lock
	resp := client.SetNX(lockKey,1,5*time.Second)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess{
		fmt.Println(err,"lock result",lockSuccess)
		return
	}

	// 累加
	getResp := client.Get(counterkey)
	cntValue , err := getResp.Int64()
	if err != nil || err == redis.Nil{
		cntValue++
		resp := client.Set(counterkey,cntValue,0)
		_,err := resp.Result()
		if err !=nil{
			fmt.Println("set value err")
		}
	}
	fmt.Println("current Counter is" ,cntValue)
	delResp:= client.Del(lockKey)
	unlockSuccess ,err := delResp.Result()
	if err == nil && unlockSuccess >0{
		fmt.Println("unlock success")
	}else{
		fmt.Println("unlock failed",err)
	}
}

func redisLock(){
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
