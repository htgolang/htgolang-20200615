{
    "code": int,
    "text" : string,
    "result" : array/object/nil
}

code :
    200 正常
    401 未登陆
    400 提交数据有错误
    500 服务端错误


Dialog
    Dialog title显示不一样 <a data-title>
    显示内部不一样 jQuery.load(url) data-url <form>
    保存 提交位置不一样 form data-url=> jQuery.serializeArray, method data-callback 字符串

    提醒 一样

    回调（创建完成以后干什么） 回调函数


    object[callback] => function() {

    }

    window

    jQuery(document).on("click", "btn-open-dialog")

    data-url, data-title


控制按钮:
    1 点击后，确认框（） data-title
    2 发起请求 + pk data-url data-pk
    3 ajax / =>提示 data-callback

    btn-control