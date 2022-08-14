// ajax 请求 无参数
function ajaxRequest(url,type,success){
    $.ajax({
        url: url,
        type: type,
        datType: "JSON",
        headers: {
            Authorization: `Bearer ` + localStorage.token //此处放置请求到的用户token,Bearer后面必选加上一个空格
        },
        dataType: "JSON",
        success: success
    })
}

// ajax 请求 携带参数
function ajaxResponse(url,type,data,success){
    $.ajax({
        url: url,
        type: type,
        datType: "JSON",
        headers: {
            Authorization: `Bearer ` + localStorage.token //此处放置请求到的用户token,Bearer后面必选加上一个空格
        },
        dataType: "JSON",
        // 将 JavaScript 对象或值转换为 JSON 字符串
        data: data,
        success: success
    })
}