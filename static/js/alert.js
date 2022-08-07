//自定义弹框
function alert(context, title) {
    //创建弹框div
    var alertFram = document.createElement("div");
    alertFram.id = "alertFram";
    alertFram.style = "position: absolute; width: 280px; height: 150px; left: 50%; top: 50%; margin-left: -140px; margin-top: -110px; text-align: center; line-height: 150px; z-index: 300;";
    var strHtml = '';
    strHtml += '<div style="list-style:none;margin:0px;padding:0px;width:100%">';
    strHtml += '	 <div id="alertFramTitle" style="background:#626262;text-align:left;padding-left:20px;font-size:14px;font-weight:bold;height:25px;line-height:25px;border:1px solid #F9CADE;color:white">[默认值]</div>';
    strHtml += '	 <div id="alertFramContext" style="background:#787878;text-align:center;font-size:12px;height:95px;line-height:95px;border-left:1px solid #F9CADE;border-right:1px solid #F9CADE;color:#fff"></div>';
    strHtml += '	 <div style="background:#626262;text-align:center;font-weight:bold;height:30px;line-height:25px; border:1px solid #F9CADE;"><input type="button" value="确 定" onclick="doOK()" style="width:80px;height:20px;background:#626262;color:white;border:1px solid white;font-size:14px;line-height:20px;outline:none;margin-top: 4px"></div>';
    strHtml += '	</div>';
    alertFram.innerHTML = strHtml;
    //将弹框添加到页面末尾
    document.body.appendChild(alertFram);
    //title
    var alertFramTitle = document.getElementById("alertFramTitle");
    alertFramTitle.innerHTML = title || "[温馨提示]";//默认值
    //context
    var alertFramContext = document.getElementById("alertFramContext");
    alertFramContext.innerHTML = context || "";//默认值
}

//确定按钮事件
function doOK() {
    //移除弹框
    var x = document.getElementById("alertFram");
    x.remove();
}
