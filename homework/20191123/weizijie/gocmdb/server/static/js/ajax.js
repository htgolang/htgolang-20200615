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
                    callback(response);
                    swal({
                        title: response["text"],
                        text: '',
                        type: "success",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    // alert("成功");
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
                        title: '',
                        text: errors.join("\n"),
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    // alert(errors.join("\n"))
                    break;
                case 401:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: false
                    }, function() {
                        window.location.replace("/");
                    });
                    // alert(response['text'])
                    // window.location.replace("/")
                    break;
                case 500:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    // alert(response['text']);
                    break;
                default:
                    swal({
                        title: response["text"],
                        text: '',
                        type: "error",
                        confirmButtonText: "确定",
                        closeOnConfirm: true
                    });
                    // alert(response['text']);
                    break;
            }
        },
        dataType: "json"
    });
}