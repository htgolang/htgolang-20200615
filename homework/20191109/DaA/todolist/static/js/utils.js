function htmlEncode(str) {
    if(typeof(str) == "undefined"){
        return "";
    }
    if(typeof(str) != "string"){
        str = str.toString();
    }

    return str.replace(/&/g,'&amp;',/'/g,'&#39;').replace(/</g,'&lt;').replace(/>/g,'&gt;').replace(/"/g,'&quot;');
}


function ajax(method, url, params, callback) {
    /*
    var ajax_callback = jQuery.post;
    if (method != "post") {
        ajax_callback = jQuery.get;
    }
    */
    var ajax_callback = method != "post" ? jQuery.get : jQuery.post
    ajax_callback(url, params,
        function(response){
            switch(response["code"]) {
                case 200:
                    callback(response);
                    break;
                case 403:
                    swal({
                        title: "",
                        text: "请重新登录",
                        type: "error",
                        timer: 5000,
                        showConfirmButton: false
                    });
                    setTimeout(function(){
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
                    var notify = jQuery.notify(response["text"],{type:'warning'});
                    setTimeout(function(){notify.update({'type': 'danger', 'progress': 25});}, 3500);
                    //jQuery.notify({type: 'danger', message: response["text"], showProgressbar: true});
                    break;
            }
        },
        "json");
}