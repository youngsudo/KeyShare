$(document).ready(function () {
    $("#login_btn").click(() => {
        console.log($("#account").val(),
            $("#password").val());
        let account = $("#account").val();
        let password = $("#password").val();
        let isCheck = false; // 是否勾选
        // 判断 account
        if (account == "") {
            alert("请输入账号");
            return;
        }
        // 如果输入的是地址
        if (/^(0x)?[0-9a-fA-F]{40}$/.test(account)) {
            console.log(account);
        }
        // 如果输入的是账号
        else if (/^[a-zA-Z0-9]{4,16}$/.test(account)) {
            console.log(account.toLowerCase());
        }
        // 报错
        else {
            alert("账号格式错误");
            return;
        }

        // 判断 password
        if (password == "") {
            alert("请输入密码");
            return;
        }
        if (/^[a-zA-Z0-9]{4,32}$/.test(password)) {
            console.log(password);
        } else {
            alert("密码格式错误");
            return;
        }

        // 记住我
        if ($("#checkbox").is(":checked")) {
            // 选中
            isCheck = true;
        }

        console.log(account, password, isCheck)
        $.ajax({
            url: "/api/v1/login",
            type: "post",
            datType: "JSON",
            data: JSON.stringify({
                "account": account,
                "password": password,
            }),
            success: function (res) {
                // console.log(res)
                // console.log(res.data.token)
                // token Authorization
                // 存储token
                // sessionStorage.setItem("token", res.data.token);
                // localStorage.setItem("token", JSON.stringify(res.data.token));
                // localStorage.setItem("user_type", JSON.stringify(res.data.userType));
                // 存储账号
                // localStorage.setItem("account", account);
                // 存储记住我
                // localStorage.setItem("isCheck", isCheck);
                // 跳转
                window.location.href = "/api/v1/index";

            }
        })

    })
})

// function loginType(i) {
//     console.log(i.attributes.ind.value)
//     if (i.attributes.ind.value == "login") {
//         window.location.href = "/api/v1/loginToAddress";
//     } else {
//         window.location.href = "/api/v1/login";
//     }
// }

// if($("#account").val().indexOf("@") != -1){
//     // 判断邮箱格式
//     let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/;
//     if(!reg.test($("#account").val())){
//         alert("邮箱格式不正确");
//         return;
//     }
// }else{
//     // 判断手机号格式
//     let reg = /^1[3-9]\d{9}$/;
//     if(!reg.test($("#account").val())){
//         alert("手机号格式不正确");
//         return;
//     }
// }