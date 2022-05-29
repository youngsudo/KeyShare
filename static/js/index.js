$(document).ready(() => {
    Array.from($(".offers-content")).forEach(element => {
        // console.log(element);
        element.addEventListener("click", () => {
            // window.location.href = element.dataset.url;
            alert(element.dataset.url)
        });
    });
})

// $.ajax({
//     url: "/api/v1/index",
//     type: "get",
//     datType: "JSON",
//     headers: {
//         Authorization: `Bearer `+ localStorage.token //此处放置请求到的用户token,Bearer后面必选加上一个空格
//     },
//     success: function (res) {
//         console.log(res)
//         // window.location.href = "/api/v1/index";
//     }
// })