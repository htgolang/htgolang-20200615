


jQuery(selector).DataTable();


table = jQuery(selector).DataTable({
    dom: "lftip",
    ordering: false,
    searching: false,
    ajax: {
        url : "/User/List",
        dataSrc: "result",
        data: function(data) {
            return {};
        }
    },
    columns: [
        {
            name: "name",
            orderable: false,
            data: function(row, type, set, meta) {
                return HtmlEncode(row["name"]); //xss
            },
        },
    ]
});

布局:
    <"class" l><"#id.class">

事件:
    table.on("draw", function() {
        // 只初始化一次
    });


绑定事件：
jQuery(selector).on("click");

给动态生成的标签绑定事件
jQuery(static-tag).on("click", selector, function() {

});


static-tag => document


用户属性
Id
name varchar(32) 默认空字符串
password varchar(1024) 默认空字符串
gender int 默认0， 0 女 1 男
birthday date 允许为null, 默认值为null
tel varchar(32) 默认为空字符串
email varchar(64) 默认为空字符
addr varchar(512) 默认为空字符串
remark varchar(1024) 默认为空字符串
is_superuser bool 默认为false， true超级管理员，false普通管理员
status int 默认为0， 0表示正常，1表示锁定
created_time datetime 添加时初始化事件
updated_time datetime 更新时间，允许为null，默认为null
逻辑删除
deleted_time datetime 删除时间，允许为null，默认为null， null未删除，非null已删除


PLugin

Session

Token


curl -XPOST url -H: authentication： Token
-H "accesskey: "
-H "signature"


session 浏览器发起的 没有这个头：authentication


abc => password
salt => xyz

db: = xyz:md5(xyz:abc)



password: abc

salt:md5(salt:password)

需要获取到目前密码的salt


Task  => 创建者 => User

Task create_user (user id)

1一个用户可以有多个任务   1: n
                        1:1
                        n:m


一个用户只能有一个Token

指定orm字段关系:  rel(), reverse()
Token => User rel(one) Token 表存储User主键

User Token reverse(one)

Task => User Task.Create_User beego.Orm

User => Task beego.ORm