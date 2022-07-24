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

        console.log(account, password, isCheck);

        ajaxResponse("/api/v1/login", "post", JSON.stringify({
            "account": account,
            "password": password,
        }),
        function (res) {
            if (res.msg == "success") {
                // 存储token
                localStorage.token = res.data.token;
                // window.sessionStorage.setItem("token", res.data.token);
                // 跳转到首页
                alert("登录成功");
                    window.location.href = "/api/v1/index";
            }
            else{
                alert(res.msg);
            }
        });
    })
})

$(".a_index").on("click", function (event) {
    event.preventDefault();//使a自带的方法失效，即无法调整到href中的URL
    var url = '';   //请求的URl,接口
    var xhr = new XMLHttpRequest();		//定义http请求对象
    xhr.open("post", url, true);//根据接口使用请求方式
    xhr.responseType = "blob";  // 返回类型blob
    xhr.setRequestHeader("accessToken", "此处获取token");
    xhr.setRequestHeader("Content-type", "application/application/octet-stream");
    //定义请求完成的处理函数,请求前可以增加加载框/禁用下载按钮逻辑
    xhr.onload = function () {
        if (this.status === 200) {
            //请求完成
            var blob = this.response;
            var reader = new FileReader();
            reader.readAsDataURL(blob);  // 转换为base64，可以直接放入a标签href
            reader.onload = function (e) {
                console.log(e);			//查看有没有接收到数据流
                // 转换完成，创建一个a标签用于下载
                var a = document.createElement('a');
                a.download = '';			//此处填写文件地址
                a.href = e.target.result;
                $("body").append(a);    // 修复firefox中无法触发click
                a.click();
                $(a).remove();
            }
        }
        else {
            alert("出现了未知的错误!");
        }
    }
    xhr.send();
});