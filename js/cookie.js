function getCookie(c_name) {
    alert(document.cookie);
    if (document.cookie.length > 0) {
        c_start = document.cookie.indexOf(c_name + "=");//��ȡ�ַ��������
        if (c_start != -1) {
            c_start = c_start + c_name.length + 1;//��ȡֵ�����
            c_end = document.cookie.indexOf(";", c_start);//��ȡ��β��
            if (c_end == -1) c_end = document.cookie.length;//��������һ������β����cookie�ַ����Ľ�β
            return decodeURI(document.cookie.substring(c_start, c_end));//��ȡ�ַ�������
        }
    }
    alert("1111");
    return "";
}

function getCookies(cookieName) {
    var strCookie = document.cookie;
    alert(strCookie);
    var arrCookie = strCookie.split("; ");
    for (var i = 0; i < arrCookie.length; i++) {
        var arr = arrCookie[i].split("=");
        if (cookieName == arr[0]) {
            return arr[1];
        }
    }
    alert("1111");
    return "";
}