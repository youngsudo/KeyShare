// 借助 token 获取当前用户信息
getUserInformation();

function getUserInformation(){
    $.ajax({
        url: "/api/v1/userInformation",
        type: "get",
        datType: "JSON",
        headers: {
            Authorization: `Bearer ` + localStorage.token //此处放置请求到的用户token,Bearer后面必选加上一个空格
        },
        dataType: "JSON",
        success: function (res) {
            console.log(res);
        },
        error: function (err) {
            console.log(err);
        }
    })
}