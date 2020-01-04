gocmdb
web + mysql

1. go编译环境
2. mysql (mariadb)
3. GOPROXY=https://goproxy.io

4. 创建数据库
    create database gocmdb default charset utf8mb4;
5. 修改数据库连接配置
    conf/db.conf
    dsn=root:881019@tcp(localhost:3306)/gocmdb?charset=utf8mb4&loc=Local&parseTime=True

    用户名:密码@tcp(数据库服务地址:数据库端口)/数据库名称?
6. go run web.go -init
7. go run web.go
    http://ip:port/auth/login