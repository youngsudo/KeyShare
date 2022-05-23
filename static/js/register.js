let isVerify = false; // 是否发送验证码

$(document).ready(function () {
    $("#register_btn").click(() => {
        // $("#address").val()
        // $("#account").val()
        // $("#email").val()
        // $("#password").val()
        // $("#repassword").val()
        // // 验证码
        // $("#address").val()

        // 地址
        if ($("#address").val() != "") {
            if (/^(0x)?[0-9a-fA-F]{40}$/.test($("#address").val())) {
                console.log($("#address").val())
            } else {
                alert("地址格式错误");
                return;
            }
        }
        // 助记词
        if ($("#mnemonic").val() != ""){
            console.log($("#mnemonic").val())
        }

        // 判断 account
        if ($("#account").val() == "") {
            alert("请输入账号");
            return;
        }
        // 如果输入的是账号
        if (/^[a-zA-Z0-9]{4,16}$/.test($("#account").val())) {
            console.log($("#account").val());
        } else {
            alert("账号格式错误,请输入4-16位的英文字母或数字");
            return;
        }

        // 判断 email
        if ($("#email").val() == "") {
            alert("请输入邮箱");
            return;
        }
        // 如果输入的是邮箱
        if (/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test($("#email").val())) {
            $("#email").val($("#email").val());
        } else {
            alert("邮箱格式错误");
            return;
        }

        // 判断 password
        if ($("#password").val() == "") {
            alert("请输入密码,密码长度为4-16位");
            return;
        }
        if (/^[a-zA-Z0-9]{6,16}$/.test($("#password").val())) {
            $("#password").val($("#password").val());
        } else {
            alert("密码格式错误,密码长度为6-16位");
            return;
        }

        // 判断 repassword
        if ($("#repassword").val() == "") {
            alert("请输入确认密码");
            return;
        }
        if ($("#password").val() != $("#repassword").val()) {
            alert("两次密码不一致");
            return;
        }

        // 判断 验证码
        if (isVerify === false) {
            alert("请先发送验证码");
            return;
        } else if ($("#verify").val() == "") {
            alert("请输入验证码");
            return;
        }

        $.ajax({
            url: "/api/v1/register",
            type: "post",
            datType: "JSON",
            data: JSON.stringify({
                "address": $("#address").val(),
                "mnemonic": $("#mnemonic").val(),
                "account": $("#account").val(),
                "email": $("#email").val(),
                "password": $("#password").val(),
                "re_password": $("#repassword").val(),
                "verifyCode": $("#verify").val()
            }),
            success: function (data) {
                console.log(data)
            }
        })

    })

    $("#verify_btn").click(() => {
        // 判断 email
        if ($("#email").val() == "") {
            alert("请输入邮箱");
            return;
        }
        // 如果输入的是邮箱
        if (/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test($("#email").val())) {
            $("#email").val($("#email").val());
        } else {
            alert("邮箱格式错误");
            return;
        }
        // ajax 后台发送邮件
        isVerify = true;
        alert("验证码已发送到您的邮箱");
    })

    // $("#register_type").click((i) => {
    //     alert(1)
    //     console.log(i)
    // })

})
function registerType(i) {
    // console.log(i)
    // console.log($("#register_type")[0].attributes.ind.value)
    console.log(i.attributes.ind.value)
    if (i.attributes.ind.value == "register") {
        window.location.href = "/api/v1/registerToAddress";
    } else {
        window.location.href = "/api/v1/register";
    }
}