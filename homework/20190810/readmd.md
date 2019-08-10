1. 整理
    a. tcp服务端开发/客户端开发流程
    b. http服务端开发/客户端开发流程
    c. rpc服务端开发/客户端开发流程

2. 爬取 https://imsilence.github.io/ 站点所有文章的标题
    注意分页

3. 定义RPC服务&客户端
    a. 可以获取服务器当前程序目录(或自定义文件)下的文件列表(文件名,类型(文件/目录),大小,修改时间)
    b. 可以根据指定文件获取当前程序目录下的文件内容

    客户端通过命令行参数控制动作(ls/cat), ls可指定目录名 cat操作需要输入文件名


    e:/hdfs/hdfs.exe
            a.txt
            b.txt
            kk/a.txt
            kk/b.txt

    client.exe -t ls -p /
    client.exe -t ls -p /kk/

    client.exe -t cat -p a.txt


    上传/删除 自己随意

4. 复习go基础知识