用户管理显示：
    前 -> 后
    / -> controller -> module -> view
    
    后 -> 前
    module -> controller -> view -> handle
    
    a. /users/ -> Action -> handler
    b. module -> 文件加载 -> 返回
    c. views -> user.html
    
用户列表查询
       用户输入数据(input) form
       form action=? method=?
       /usrs/query?xxx
       
用户登录
    a. 打开登录页面
        Get /usr/login => Action => Execute(/user/login.html)
    
    b. 登录流程
        方法一：获取用户名/密码 -> 验证, 找输入用户名/密码都相等的用户
        方法二：通过用户名去查找用户 -> 没找到user
                    找到 -> 判断密码是否正确
                    
        验证成功
            -> 跳转到任务列表
        验证失败
            -> 返回输入信息页面, 并提示错误, 及用户原输入信息
     
     发现用户没有登录时跳转到登录页面让进行登录。
     机制：跟踪用户状态
            
     技术：
        怎么生成Session Id => go.uuid
        存储内存/文件/数据库/reids
        怎么设置, 怎么修改response header set-cookie
        获取cookie request
         
            
退出登陆
    session 销毁
    cookie  销毁 
            
添加用户
     a 打开添加页面
     Get /usrs/create
     
     b 提交数据        
     Post /usrs/create/
     
     验证：
        用户名长度 4-12 不允许重复，不能为空
        密码长度6-30位
        出生日期60年到现在
     验证成功 添加用户 并持久化
     验证失败 返回添加页面
     
修改用户

删除用户
        
         
        
            
            
            
            
            
            
            