$(document).ready(() => {
    $("#contact_but_input").click(() => {
        // console.log($("#name").val(),
        //     $("#email").val(),
        //     $("#address").val(),
        //     $("#your_message").val(),
        // )

        // 如果输入的是邮箱
        if (/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test($("#email").val())) {
            console.log($("#email").val());
        } else {
            alert("邮箱格式错误");
            return;
        }


        $.ajax({
            url: "/api/v1/contact",
            type: "post",
            headers: {
                Authorization: "Bearer " + localStorage.getItem("token")
            },
            datType: "JSON",
            data: JSON.stringify({
                "name": $("#name").val(),
                "email": $("#email").val(),
                "address": $("#address").val(),
                "message": $("#your_message").val(),
            }),
            success: function (res) {
                
            }

        })

    })
})