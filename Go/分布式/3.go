package main

// 分布式定时任务是很多公司需要的东西
// 我在的一家公司 定时任务遍布 各个服务器。每个服务器七八百个，当时接手直接懵逼。无法管理。甚至还有没人知道的定时任务。
// 最后改版了cronsun 分布式定时任务才将他们统一管理起来。


// 书中介绍的太过理论化 我这边描述一下cronsun的结构

// 它分为两部分 master worker
/**
MASTER

初始化话操作

1.  mogodb创建索引

2.  apiserver

3.  监听报警  监控节点 任务的致命错误

4.  定时清除日志
 */

/**
WORKER

初始化操作


1.节点注册
step 1 创建节点信息
/crontab/node/xx

step 2 申请固定时间的 租约
创建租约（节点的租约）

step 3 写入pid
write pid

2. 维持租约
初始化租约ID 并配置
增加 读写锁 对租约id的读取
处理续租应答 读到就续租 没就设置

3. 启动服务
这里保持节点租约续租 协程
这里任务被哪个节点获取运用了分布式锁
step 1 获得所有的节点分组    etcdctl get /crontab/group/  --prefix
step 2 获得所有任务列表    etcdctl get "/crontab/cmd/"  --prefix
step 3 甄别当前节点的任务 进行 装饰
step 4 获取该节点的最终
step 5 加入 cron
step 6 运行 cron 这里运用 Schedule 和 start的方式运行 协程的方式
step 7 利用etcd事件监听 来 解决任务的变更 立即执行 组的变更 运行中的进程数量

 go n.watchJobs()
 go n.watchExcutingProc()
 go n.watchGroups()
 go n.watchOnce()
 go n.watchCsctl()

step 8  记录mongo
step 9  阻塞运行
 */

// 节点机器时间一致很重要 踩的坑