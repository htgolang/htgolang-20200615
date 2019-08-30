1. 用户列表显示
    前 -> 后
    / => controller => model => view

    后->前
    model->controller->view->handle

    a. /users/ => Action => HandleFunc
    b. model => 文件中加载 => 返回
    c. views => users.html

2. 用户列表查询
    用户输入数据(input q) form
    form action=? method=?
    /users/ get q=xxx
    FormValue("q")
    users = GetUsers(q)


    /users/query post

    PostFormValue("q")

    GetUsers()
    for users
    users contains q

3. 用户登陆
    a. 打开登陆页面
        Get /user/login => Action => Execute(user/login.html)

    b. 登陆流程
        POSt /user/login

        方法一: 获取用户名/密码 => 验证 找输入用户名/密码都相等的用户
        [用]方法二: 通过用户名去查找用户 => 没找到 User
                    找到 => 判断密码是否正确（通过User方法来验证）

        认证成功
            => 跳到到任务列表
        认证失败
            => 返回到输入信息页面，并提示错误，及用户原输入信息

    发现用户没有登陆时跳转到登陆页面让进行登陆
    机制：跟踪用户状态
    session + cookie

    你银行办业务

    你（浏览器）                 银行(服务器)
    1. 第一次去银行              开户                     0
                                给你银行卡
    2. 第二次去银行带上卡           存1w                  1w
    3. 第3次去银行带上卡            取1k                  9k
    4. 没带银行卡                想要取钱(银行拒绝)


    cookie的存储: 在浏览器
    cookie的信息: 卡号

    浏览器                                  服务器
    1. 第一次请求                               开辟一定存储空间（编号=>session ID, 存储空间=>session）
                                               将session ID返回给客户端

                                                    response header
                                                    Set-Cookie: session=xxxxx
        浏览器接收到请求存储cookie信息

    2. 第二次请求   浏览器会读取cookie信息（session ID）并通过http请求提交给服务器      (登陆 成功 在存储空间User)
                                            获取Session ID 查找对应存储空间中的数据
    3. 以后请求中都会带cookie信息                   从存储空间中尝试获取User，如果没有获取到， 未登陆
                                                                         如果获取到，已登录


    技术:
        怎么生成Session ID => go.uuid

        存储内存/文件/数据库/redis

        怎么设置， 怎么取修改response header set-cookie
        获取cookie request

4. 退出登陆
    session销毁
    cookie销毁

5. 添加用户
     a. 打开添加页面
        Get /users/create/
     b. 提交数据
        Post /users/create/
        r.PostFormValue
        验证
            用户名长度 4-12
                不允许重复
            密码长度 6-30
            出生日期 1960 - 2019
        验证成功 添加用户并持久化
        验证失败，返回添加页面，并回显错误和输入信息