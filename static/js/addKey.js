$(document).ready(function () {

    // 我们将用来创建随机密码字母的所有函数名的对象
    const randomFunc = {
        lower: getRandomLower,
        upper: getRandomUpper,
        number: getRandomNumber,
        symbol: getRandomSymbol,
    };

    // Generator Functions
    // 所有负责返回随机值的函数，我们将使用这些函数来创建密码。
    function getRandomLower() {
        return String.fromCharCode(Math.floor(Math.random() * 26) + 97);
    }
    function getRandomUpper() {
        return String.fromCharCode(Math.floor(Math.random() * 26) + 65);
    }
    function getRandomNumber() {
        return String.fromCharCode(Math.floor(Math.random() * 10) + 48);
    }
    function getRandomSymbol() {
        const symbols = '~!@#$%^&*()_+{}":?><;.,';
        return symbols[Math.floor(Math.random() * symbols.length)];
    }
    // 输入滑块，将用于更改密码的长度
    const lengthEl = document.getElementById("slider");
    let passtext = $("#pass_text");
    
    // 复选框表示负责根据用户创建不同类型密码的选项
    const uppercaseEl = document.getElementById("uppercase");   // 大写字母
    const lowercaseEl = document.getElementById("lowercase");   // 小写字母
    const numberEl = document.getElementById("number");  // 数字
    const symbolEl = document.getElementById("symbol");  // 符号
    // 生成密码的按钮
    const generateBtn = document.getElementById("generate");

    // 单击“生成”时，将生成密码
    generateBtn.addEventListener("click", () => {
        const length = +lengthEl.value;
        const hasLower = lowercaseEl.checked;
        const hasUpper = uppercaseEl.checked;
        const hasNumber = numberEl.checked;
        const hasSymbol = symbolEl.checked;
        console.log(length, hasLower, hasUpper, hasNumber, hasSymbol);
        pass = generatePassword(length, hasLower, hasUpper, hasNumber, hasSymbol);
        passtext.val(pass);
    });

    // 负责生成密码并返回密码的函数。
    function generatePassword(length, lower, upper, number, symbol) {
        let generatedPassword = "";
        const typesCount = lower + upper + number + symbol;
        const typesArr = [{ lower }, { upper }, { number }, { symbol }].filter(item => Object.values(item)[0]);
        if (typesCount === 0) {
            return "";
        }
        for (let i = 0; i < length; i++) {
            typesArr.forEach(type => {
                const funcName = Object.keys(type)[0];
                generatedPassword += randomFunc[funcName]();
            });
        }
        return generatedPassword.slice(0, length);
    }

    let flexbut = $("#flex_but_div");
    passtext.focus(function () {
        // 输入框获取焦点
        if (passtext.val() != "") {
            flexbut.css("display", "block");
            // 输入框失去焦点
            // passtext.blur();
        }
    });

    var clipboard = new ClipboardJS('.btn');
    flexbut.click(() => {
        clipboard.on('success', function (e) {
            console.info('Action:', e.action);
            console.info('Text:', e.text);
            console.info('Trigger:', e.trigger);
        });

        clipboard.on('error', function (e) {
            console.log(e);
        });
        // flexbut.css("display", "none");
    });


    function changeV(){
        
    }
})
// 页面加载完成
// window.onload = function () {
//     // 自动生成一次密码
//     generateBtn.click();
// }

$(document).ready(function () {
    let password = $("#password");
    let createPass = $("#createPass")
    password.click(()=>{
        // alert("密码已复制到剪贴板");
        createPass.css("display","block");
    })
    // password.blur(()=>{
    //     createPass.css("display","none");
    // })

    let eye = $("#showEye");
    eye.click(()=>{
        if(eye[0].children[0].attributes[0].value != "fa fa-eye"){
            eye[0].children[0].attributes[0].value = "fa fa-eye";
            $("#password")[0].type = "text";
        }else{
            eye[0].children[0].attributes[0].value = "fa fa-eye-slash";
            $("#password")[0].type = "password";
        }
    })
})
let length = document.getElementById("slider");
let plength = $("#text-plength");

function changeV() {
    let boxL = parseInt(length.value);
    plength[0].innerText = boxL;
}
