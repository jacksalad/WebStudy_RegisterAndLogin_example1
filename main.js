$(document).ready(function () {
    // 注册表单提交事件
    $("#registerForm").submit(function (e) {
        e.preventDefault();
        var username = $("#username").val();
        var password = $("#password").val();

        $.post("http://127.0.0.1:8080/register", { username: username, password: password })
            .done(function (response) {
                // 注册成功的处理逻辑
                var data = JSON.parse(response)
                alert("注册成功！用户名: " + data.username);
            })
            .fail(function (xhr) {
                // 登录失败的处理逻辑
                if (xhr.responseText === "00\n") {
                    alert("注册失败，用户名已存在");
                } else {
                    alert("注册失败，请重试");
                }
            });
    });

    // 登录表单提交事件
    $("#loginForm").submit(function (e) {
        e.preventDefault();
        var loginUsername = $("#loginUsername").val();
        var loginPassword = $("#loginPassword").val();

        $.post("http://127.0.0.1:8080/login", { username: loginUsername, password: loginPassword })
            .done(function (response) {
                // 登录成功的处理逻辑
                var data = JSON.parse(response)
                alert("登录成功！用户名: " + data.username);
            })
            .fail(function (xhr) {
                // 登录失败的处理逻辑
                if (xhr.responseText === "01\n") {
                    alert("用户名不存在");
                } else if (xhr.responseText === "02\n") {
                    alert("密码错误");
                } else {
                    alert("登录失败，请重试");
                }
            });
    });
});
