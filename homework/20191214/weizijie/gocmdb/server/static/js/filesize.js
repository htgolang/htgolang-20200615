function FileSize(num) {
    units = ["B", "KB", "MB", "GB", "TB", "PB"];
    index = 0;
    while(num > 1024) {
        num /= 1024;
        index += 1
    }

    console.log(parseInt(num).toFixed(2))
    return parseInt(num).toFixed(2) + units[index];
}