/* 密码重复输入严重 */
let upwd = document.getElementById("user_password");
let cpwd = document.getElementById("confirm_password");
cpwd.addEventListener("blur", function (event) {
    if (cpwd.value !==  upwd.value){
        cpwd.innerHTML = "密码不一致"
    }else{
        cpwd.innerHTML = "ok"
    }
}, true);