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
                    alert("成功");
                    break;
                case 400:
                    var errors = [];
                    jQuery.each(response["result"], function(k, v) {
                        errors.push(v['Message']);
                    });
                    if(!errors) {
                        errors.push(response['text']);
                    }
                    alert(errors.join("\n"))
                    break;
                case 401:
                    alert(response['text'])
                    window.location.replace("/")
                    break;
                case 500:
                    alert(response['text']);
                    break;
                default:
                    alert(response['text']);
                    break;
            }
        },
        dataType: "json"
    });
}