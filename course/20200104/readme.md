
任务
    终端
    插件
    插件参数
    状态
    开始时间
    结束时间

任务结果
    任务
    结果
    失败原因


AGENT 如何获取任务
    定时从server上要任务 每隔10s问下server，我有哪些任务（新创建）要执行
    (ENS)我去查一下，2个任务，标记任务已经返回给AGENT(执行中)，返回给AGENT

任务执行
    每个任务 => 插件
    插件管理 => 负责执行

    ENS =>（管道）=> Manager 通信


redis常用操作

auth 认证
ping 测试

针对KEY的处理
type,del,keys,ttl,expire,exists

针对每种数据类型的
    字符串
        set
        get
        mset
        mget
        setnx
        incr
        incrby
        decr
        decrby
    列表
        lpush
        rpush
        lpop
        rpop
        lrange
        llen
        brpop
        blpop
    哈希
        hset
        hget
        hmset
        hmget
        hsetnx
        hdel
        hexists
        hlen
        hgetall
    集合
        sadd
        scard
        smembers
        sismember
        srem
    有序集合
        zadd
        zcard
        zrange
        zrevrange
        zrangebyscore
        zrevrangebyscore
        zrem

    发布订阅
        subscribe
        publish
运维的


心跳 =》 db
心跳 =》 zset
uuid score time unix

zrangebyscore uuid
now - 5 now


supervisor

1. yum install/pip install
2. systemd


内推
    1. 知道部门现在在做什么
    2. 部门想做做什么(devops 规划，开发-》测试（自动化测试），测试-》上线（自动发布，审批），监控（业务数据监控，资源监控）)
    3. 自己做过哪些事情
    4.

BOSS 你想去的（大）公司（提供相应的职位，他们在干什么）

    希望我做什么
    我能做什么
    我哪里需要补充(学)

简历
面试
    1. 个人介绍
    2. 项目（设计，架构，问题），编程语言(写一段代码，编程思想，设计模式)，算法(设计，代码，数据结构，算法)，网络，操作系统


开发
    => 懂(安全 & 测试 & 运维)

运维 (devops)
    => 懂（开发 & 测试 & 运维（监控，上线））
            自动化测试 自动化上线（运维） 资源/业务系统（可用性, 响应）

    云(openstack，k8s) docker


索引
    （名词）索引 （库）
    （动词）索引 插入数据

    type （表）

增
POST
删
改
查