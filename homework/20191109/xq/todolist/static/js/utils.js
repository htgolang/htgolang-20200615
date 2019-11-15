function htmlEncode(str) {
    if (typeof(str) == "undefined"){
        return "";
    }

    if (str == null ){
        return ""
    }
    if (typeof(str) != "string"){
        str = str.toString();
    }

    return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/'/g, '&#39;').replace(/"/g, '&quota');
}