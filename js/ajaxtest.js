﻿// ajax 对象
function ajaxObject() {
    var xmlHttp;
    try {
        // Firefox, Opera 8.0+, Safari
        xmlHttp = new XMLHttpRequest();
    }
    catch (e) {
        // Internet Explorer
        try {
            xmlHttp = new ActiveXObject("Msxml2.XMLHTTP");
        } catch (e) {
            try {
                xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
            } catch (e) {
                alert("您的浏览器不支持AJAX！");
                return false;
            }
        }
    }
    return xmlHttp;
}

// ajax post请求：
function ajaxPost(url, data, fnSucceed, fnFail, fnLoading) {
    var ajax = ajaxObject();
    ajax.open("post", url, true);
    ajax.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    ajax.onreadystatechange = function () {
        if (ajax.readyState == 4) {
            if (ajax.status == 200) {
                return fnSucceed(ajax.responseText);
            }
            else {
                fnFail("HTTP请求错误！错误码：" + ajax.status);
            }
        }
        else {
            fnLoading();
        }
    }
    ajax.send(data);

}

function asuccess(responseText) {
    //alert(responseText);
    if (responseText == "1") {
        alert("已有该用户名！");
    }
    return responseText;
}
function bsuccess(responseText) {
    //alert(responseText);
    if (responseText == "1") {
        alert("已有该课程！");
    }
}
function afailure() {
    //alert("failure");
}

function apost(url, data) {
    //alert(JSON.stringify(data));
    ajaxPost(url,JSON.stringify(data), asuccess, afailure, null);
}
function bpost(url, data) {
    //alert(JSON.stringify(data));
    ajaxPost(url, JSON.stringify(data), bsuccess, afailure, null);
}

function apostRaw(url, data) {
    ajaxPost(url, data, asuccess, afailure, null);
}