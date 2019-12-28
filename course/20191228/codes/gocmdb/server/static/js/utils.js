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


function ajaxRequest(method, url, params, callback) {
    console.log("2. ajaxRequest")
    jQuery.ajax({
        type: method,
        url: url,
        data: params,
        beforeSend: function(xhr) {
            xhr.setRequestHeader("X-Xsrftoken", jQuery.base64.atob(jQuery.cookie("_xsrf").split("|")[0]));
        },
        success: function(response) {
            console.log("3. success")
            switch(response["code"]) {
                case 200:
                    swal({
                        title: response['text'],
                        text: "",
                        type: "success",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    }, function() {
                        swal.close();
                        callback(response);
                    });
                    break;
                case 400:
                    var errors = [];
                    jQuery.each(response["result"], function(k, v) {
                        errors.push(v['Message']);
                    });
                    if(errors.length == 0) {
                        errors.push(response['text']);
                    }
                    swal({
                        title: '',
                        text: errors.join("\n"),
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    break;
                case 401:
                    swal({
                        title: response['text'],
                        text: "",
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    }, function() {
                        swal.close();
                        window.location.replace("/")
                    });
                    break;
                case 500:
                    alert(response['text']);
                    break;
                default:
                    swal({
                        title: response['text'],
                        text: "",
                        type: "error",
                        showCancelButton: false,
                        confirmButtonColor: "#DD6B55",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    break;
            }
        },
        dataType: "json"
    });
}