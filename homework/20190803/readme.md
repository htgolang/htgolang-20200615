作业
1. Go基本语法 => 复习 => mm
2. 点对点聊天 => 代码完成 => 相应检查
    比如命令输入空行 => 不用发送
    输入前后空格 => 去掉

3. httpserver.go 挑战
    a. 了解HTTP协议 (使用浏览器作为客户端)
        i. 客户端先发起数据
        ii. 只需要关注第一行
            GET/POST /index.html HTTP/1.1
            请求方法 请求路径 协议/版本

    需要做的
    获取请求路径 => 在进程所在目录下寻找请求路径对应文件
    找到文件 => 将文件内容输出给客户端(浏览器)
    文件不存在 => 404.html => 发送给客户端
    / => index.html => 发送给客户端

    如何使用例程去处理与客户端的请求


http

request
1: 请求行 \r\n GET/POST url HTTP/1.0
2-n: key: value \r\n请求头
\r\n
请求体


response
1: 响应行 HTTP/1.0 STATUS_CODE STATUS_TEXT
2-n: key: value 响应头
\r\n
响应体