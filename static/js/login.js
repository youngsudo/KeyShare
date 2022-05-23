$(document).ready(function () {
    $("#login_btn").click(() => {     
        console.log($("#account").val(),
            $("#password").val());
        // 判断 account
        if ($("#account").val() == "") {
            alert("请输入账号");
            return;
        }
        // 如果输入的是地址
        if (/^(0x)?[0-9a-fA-F]{40}$/.test($("#account").val())) {
            $("#account").val($("#account").val());
        }
        // 如果输入的是账号
        else if (/^[a-zA-Z0-9]{4,16}$/.test($("#account").val())) {
            $("#account").val($("#account").val().toLowerCase());
        }
        // 报错
        else {
            alert("账号格式错误");
            return;
        }

        // 判断 password
        if ($("#password").val() == "") {
            alert("请输入密码");
            return;
        }
        if (/^[a-zA-Z0-9]{4,32}$/.test($("#password").val())) {
            $("#password").val($("#password").val());
        } else {
            alert("密码格式错误");
            return;
        }

        // 记住我
        if($("#checkbox").is(":checked")){
            // 选中
           console.log(1)
        }else{
            console.log(2)
        }

    })
})

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