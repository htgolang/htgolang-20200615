function htmlEncode(str) {
    if (typeof (str) == "undefined" || str == null) {
        return "";
    }

    if (typeof (str) != "string") {
        str = str.toString()
    }
    return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/'/g, '&#39;').replace(/"/g, '&quot;')
}

function ajaxRequest(method, url, params, callback) {
    jQuery.ajax({
        type: method,
        url: url,
        data: params,
        beforeSend: function (xhr) {
            xhr.setRequestHeader("X-Xsrftoken", jQuery.base64.atob(jQuery.cookie("_xsrf").split("|")[0]));
        },
        success: function (response) {
            switch (response["code"]) {
                case 200:
                    callback(response);
                    swal({
                        title: "操作成功！",
                        timer: 1000,
                        showConfirmButton: false,
                        type: "success"
                    });
                    break;
                case 400:
                    var errors = [];

                    jQuery.each(response["result"], function (index, error) {
                        errors.push(error["Message"]);
                    });

                    if (!errors) {
                        errors.push(response["Message"]);
                    }
                    swal({
                        title: "",
                        text: errors.join("\n"),
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确认",
                        cancelButtonText: "取消",
                        closeOnConfirm: true,
                        closeOnCancel: false,
                    });
                    break;
                case 401:
                    swal({
                        title: response["text"],
                        timer: 1000,
                        showConfirmButton: false,
                        type: "error"
                    });
                    window.location.replace("/");
                    break;
                case 403:
                    swal({
                        title: response["text"],
                        timer: 1000,
                        showConfirmButton: false,
                        type: "error"
                    });
                    break;
                case 500:
                    swal({
                        title: response["text"],
                        timer: 1000,
                        showConfirmButton: false,
                        type: "error"
                    });
                    break;
                default:
                    swal({
                        title: response["text"],
                        timer: 1000,
                        showConfirmButton: false,
                        type: "error"
                    });
                    break;
            }
        },
        dataType: "json"
    })
}