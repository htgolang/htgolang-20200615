function htmlEncode(str) {

    if(typeof(str) == "undefined" || str == null) {
        return "";
    }

    if(typeof(str) != "string") {
        str = str.toString();
    }

    return str.replace(/&/g, '&amp;')
        .replace(/'/g, '&#39;')
        .replace(/</g, '&lt;')
        .replace(/>/g, '&gt;')
        .replace(/"/g, '&quot;');
}

function FileSizeb(size) {
    if(size < 8 ){
        return size.toFixed(2) + "b";
    }
    size /= 8;
    var index = 0;
    var units = ["B","KB","MB","GB","TB","PB"];

    while(size >= 1024){
        size /= 1024;
        index += 1;
    }
    return size.toFixed(2) + units[index];

}


function ajaxRequest(method, url, params, callback) {
    console.log("2.ajaxRequest");
    jQuery.ajax({
        type: method,
        url: url,
        data: params,
        beforeSend: function(xhr) {
            xhr.setRequestHeader("X-Xsrftoken", jQuery.base64.atob(jQuery.cookie("_xsrf").split("|")[0]));
        },
        success: function(response) {
            console.log("3.success函数")
            switch(response["code"]) {
                case 200:
                    callback(response);
                    var notify = jQuery.notify(response["text"],{type:'info'});
                    setTimeout(function(){notify.update({'type': 'success', 'progress': 25});}, 3500);
                    break;
                case 400:
                    var errors = [];
                    jQuery.each(response["result"], function(k, v) {
                        errors.push(v['Message']);
                    });
                    if(!errors) {
                        errors.push(response['text']);
                    }
                    swal({
                        title: response["text"],
                        text: errors.join("\n"),
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    });
                    break;
                case 401:
                    var errors = [];
                    jQuery.each(response["result"], function(k, v) {
                        errors.push(v['Message']);
                    });
                    if(!errors) {
                        errors.push(response['text']);
                    }
                    swal({
                        title: response["text"],
                        text: errors.join("\n"),
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    });
                    window.location.replace("/")
                    break;
                case 403:
                    swal({
                        title: response["text"],
                        text: response["result"]["Message"],
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    });
                    break;
                case 500:
                    swal({
                        title: response["text"],
                        text: response["result"]["Message"],
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    });
                    break;
                default:
                    swal({
                        title: response["text"],
                        text: response["result"]["Message"],
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    });
                    break;
            }
        },
        dataType: "json"
    });
}