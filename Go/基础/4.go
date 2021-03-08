package main

func main() {

}


/**
历史的发展，并将带来新时代。
原先的单核，到现在的多核。正是在多核和网络化的时代背景下诞生的原生支持并发的编程语言GO。

常见的并发编模型中有很多，其中常见的有 多线程模型和消费并发模型
理论上两者是一致的
go 语言就是集大成者
 */

/**
go协程 是轻量级线程。在我的理解中，是一种对线程的调度 多路复用。
一个Goroutine会以一个很小的栈启动 （是不是很像线程），当遇到深度递归导致当前栈空间不足时，
Goroutine会根据 需要动态地伸缩栈的大小(主流实现中栈的最大值可达到1GB)。
因为启动的代价 很小，所以我们可以轻易地启动成千上万个Goroutine。

同时运行时也包含了一个调度器 可以在多个系统线程上调度多个Goroutine，只有当他们阻塞才会被调度。
 */