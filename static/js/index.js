$(document).ready(()=>{
    Array.from($(".offers-content")).forEach(element => {
        // console.log(element);
        element.addEventListener("click", ()=>{
            // window.location.href = element.dataset.url;
           alert(element.dataset.url)
        });
    });
})