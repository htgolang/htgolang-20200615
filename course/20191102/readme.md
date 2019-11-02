beego.Get("/", function(c *context.Context) {

})


/users/delete/?id=1
/users/delete/1/

:name

/users/delete/:id:int/

type UserController struct {
    beego.Controller
}

beego.Router("/user/", &UserController{})

PUT /user/
GET /user/


Login
Error

Get /user/ => Login
Post /user/ => Login
其他的 /user/ => Error
PUT /user/ => Create
DELETE /user/ = > Delete

beego.Router("/user/", &UserController{}, "get,post:Login;put:Create;delete:Delete;*:Error")

beego.AutoRouter(&UserController{})

Login
Error
Create
Delete

/user/login => Login
/user/error => Error


login:
body: name, password

delete:/delete/?id=1


type LoginForm struct {
    Name string `form:"name"`
    Password string `form:"password"`
}


1. layouts 公共js/css
2. LayoutContent
    每个页面都自己的js/css
    LayoutSections

    LayoutStyles
    LayoutScripts

3. Database
    a. table => 页面 => DataTable生成分页页面数据(前端查询 js)
    b. ajax => 请求数据 => DataTable根据ajax返回数据生成分页页面数据（前端查询 js）
    c. 全后端 ajax

    jQuery.get
    jQuery.post
    url, {}, function(response) {}, "json"


    "code" : 200/400/403/500
    "text": "",
    "result": nil/[]/{}

创建
    dialog => 内容(index.html) 不发送请求
编辑
    a. dialog => 表单内容（index.html）
                数据 => 发送请求获取的 ajax => 填充到表单中
                对应关系的问题
    dialog => ajax 获取表单内容+数据 html => 放在dialog中
        jQuery(selector).load(url) ajax