作业处理

用户管理
    一个 dialog
        创建/编辑 流程做完
        锁定/解锁/删除 流程做完

    当前登陆用户不能锁定/删除/解锁 自己
    当前用户只能查看和生成自己的Token

    修改密码
    alert => sweetalert

1. ajax请求未登录返回json
2. 创建


多云管理
管理多个云平台
    阿里云
    腾讯云
    aws
    azure
    华为云
    京东云
    青云
    Openstack
    ...

获取虚拟机
启动
停止
重启

server/cloud

plugins/aliyun
        tenant
        aws

manager
instance => vm


Type string
Name string
Init(addr region accessKey, secrectKey)
TestConnect() error
GetInstances() []*Instance
StartInstance(uuid) error
StopInstance(uuid) error
RebootInstance(uuid) error


配置信息
    认证
    地址
    Region配置

Platform
    Id
    Name
    Type
    Addr
    AccessKey
    SecrectKey
    Region
    Remark
    CreatedTime
    DeletedTime
    SyncTime
    CreateUser rel,reverse
    Status


VirtualMachine
    Platform 1: n

    UUID
    Name
    CPU
    Memeory
    OS
    PrivateAddrs
    PublicAddrs
    Status string
    VmCreatedTime
    VmExpiredTime

    CreatedTime
    DeletedTime
    UpdatedTime