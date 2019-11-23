function htmlEncode(str) {

    if(typeof(str) == "undefined" || str == null) {
        return "";
    }

    if(typeof(str) != "string") {
        str = str.toString();
    }

    return str.replace(/&/g, '&amp;', /'/g, '&#39;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
}


function ajax(method, url, params, callback) {
    // var ajax_callback = jQuery.post
    // if (method != "post") {
    //     ajax_callback = jQuery.get;
    // }
    var ajax_callback = method != "post" ? jQuery.get : jQuery.post
    // jQuery.ajax
    ajax_callback(url, params,
        function(response) {
            switch(response["code"]) {
                case 200:
                    callback(response);
                    break;
                case 403:
                    swal({
                        title: "",
                        text: response['text'],
                        type: "error",
                        timer: 5 * 1000,
                        showConfirmButton: false
                    });

                    setTimeout(function() {
                        window.location.reload();
                    }, 5 * 1000);
                    break;
                default:
                    swal({
                        title: "",
                        text: response['text'],
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确认",
                        cancelButtonText: "取消",
                        closeOnConfirm: true,
                        closeOnCancel: false
                    });
            }
        },
        "json"
    );
}