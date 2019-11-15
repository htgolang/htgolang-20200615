function htmlEncode(str) {
    if(typeof(str) == "undefined") {
        return "";
    }

    if(typeof(str) != "string") {
        str = str.toString();
    }
    return str.replace(/&/g,'&amp;', /'/g, '&#39；').replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/"/g, '&quot;');
}