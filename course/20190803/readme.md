json

gob
csv


load
store


1. 在未sleep之前 看着只执行了3，1和2未执行


锁
在对资源操作时获取锁
获取了 => 操作 => 释放锁

没获取 => 等待

读写错 RWMutex

Lock
RLock

Unlock
RUnlock

n个在写  m个在读
多个读 不需要锁
1读 / 写
RLock => RLock 不阻塞 RUnlock
RLock => Lock 阻塞
Lock => Rlock 阻塞
Lock => Lock 阻塞 Unlock


queue := make([]int, 0)
先进先出

入
queue = append(queue, x)

出
queue[0]
queue = queue[1:]

堆栈
stack := make([]int, 0)
先进后出

入
stack = append(stack, x)

出
stack[len(stack) - 1]

stack[:len(stack) -1]


// 监听  0.0.0.0:9999 端口
// 等待客户端连接 //当客户端连接完成后，给客户端发送一个当前时间 //关闭客户端连接