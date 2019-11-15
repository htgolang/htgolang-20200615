function ajax(method, url, params, callback) {
    // var ajax_callback = jQuery.post
    // if (method != "post") {
    //     ajax_callback = jQuery.get;
    // }
    
    var ajax_callback = method != "post" ? jQuery.get : jQuery.post
    ajax_callback(url, params,
        function(response) {
            switch(response["code"]) {
                case 200:
                    callback(response);
                    break;
                case 403:
                    // 5秒后自动关闭提示框
                    swal({
                        title: "",
                        text: response["text"],
                        type: "error",
                        timer: 5000,
                        showCancelButton: false,
                        showConfirmButton: false,
                    }); 
                    setTimeout(function() {
                        window.location.reload();
                    }, 5 * 1000);
                    break;
                case 400:
                    var errors = [];
                    $.each(response["result"], function (index, error) {
                        errors.push(error["Message"]);

                    });
                    if (!errors) {
                        errors.push(respone["text"]);
                    }
                    //alert(errors.join("\n"));
                    swal({
                        title: "",
                        text: errors.join("\n"),
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确认",
                        cancelButtonText: "取消",
                        closeOnConfirm: true,
                        closeOnCancel: false
                    }); 
                    break;
                default:
                    //alert(response["text"]);
                    swal({
                        title: "",
                        text: response["text"],
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确认",
                        cancelButtonText: "取消",
                        closeOnConfirm: true,
                        closeOnCancel: false
                    }); 
            }
        }, "json"
    ); 
}