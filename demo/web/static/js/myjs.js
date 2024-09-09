

$("#upadduced").click(function () {

    var dic = {
        mgdphone: $("#xgmgdphone").val(),
        passwd: $("#xgpasswd").val(),
        address: $("#xgaddress").val(),
        uinfo: $("#xguinfo").val(),
        xgtype: $("#xgyou_type").val(),
        hour: $("#xghour").val(),
        minute: $("#xgminute").val(),
        week_start: $("#xgweek_start").val(),
        week_end: $("#xgweek_end").val(),
    };
    // console.log(dic);

    // console.log(dic);
    // 将获取的字符串转换为 json数据
    var msg = JSON.stringify(dic);
    // var $test = $(".test");
    // ajax发送

    $.ajax({
        // 发送到的地址
        url: "/../IndexViews/UpdateCrontab/",
        //请求方式
        type: "post",
        //指定请求的数据格式为json
        contentType: "application/json",
        // 数据
        data: msg,
        //指定响应的格式为json,注意如果后台没有放回json类型的数据格式下方的success不会执行
        dataType: "json",
        success: function (res) {
            console.log(res);
            var data = res;
            if (data.status == 200) {
                console.log(res);
                alert(data.data[0].info);
                // location.href = "./admin/index.html";
                location.href = "./table-list.html";
            } else if (data.status == 400) {
                // console.log(res);
                alert(data.data[0].info);
                location.href = "./table-list.html";
            } else {
                // alert("111");
            }
        },
        error: function () {
            console.log("请求出错！");
        }
    });

});


$("#del").click(function () {
    var dic = {
        type: $("#type_msg1").text()
    };
    // console.log(dic);
    // console.log(dic);
    // 将获取的字符串转换为 json数据
    var msg = JSON.stringify(dic);
    // var $test = $(".test");
    // ajax发送

    $.ajax({
        // 发送到的地址
        url: "/../IndexViews/DeleteCrontab/",
        //请求方式
        type: "post",
        //指定请求的数据格式为json
        contentType: "application/json",
        // 数据
        data: msg,
        //指定响应的格式为json,注意如果后台没有放回json类型的数据格式下方的success不会执行
        dataType: "json",
        success: function (res) {
            console.log(res);
            var data = res;
            if (data.status == 200) {
                console.log(res);
                alert(data.data[0].info);
                location.href = "./table-list.html";
                // location.href = "./admin/index.html";
            } else if (data.status == 400) {
                // console.log(res);
                alert(data.data[0].info);
                // location.href = "./table-list.html";
            } else {
                // alert("111");
            }
        },
        error: function () {
            console.log("请求出错！");
        }
    });

});

function xgfz1() {
    // alert("加载完成");
    var xgp = $("#uphne_msg1").text()//账号
    var xgpass = $("#upass_msg1").text()//mima
    var xgmsg = $("#rmsg_msg1").text()//msg
    var xgadd = $("#radd_msg1").text()//打卡地点
    var xgtype = $("#type_msg1").text()//类型
    var xgtime = $("#rtime_msg1").text()//添加时间
    var xgrtime = $("#rdtime_msg1").text()//打卡时间周期
    // console.log(xgrtime);
    var xgrminute = xgrtime.slice(0, 2)
    var xgrhour = xgrtime.slice(3, 4)
    var xgrweek_start = xgrtime.slice(9, 10)
    var xgrweek_end = xgrtime.slice(11, 12)
    // console.log(xgrweek_end);



    $("#xgmgdphone").val(xgp), //提交修改手机号
        $("#xgpasswd").val(xgpass),//提交修改密码
        $("#xgaddress").val(xgadd),//提交修改打卡地址
        $("#xguinfo").val(xgmsg),//提交修改打卡信息
        $("#xgyou_type").val(xgtype),//提交修改类型
        $("#xghour").val(xgrhour),     //提交修改小时
        $("#xgminute").val(xgrminute),    //提交修改分钟
        $("#xgweek_start").val(xgrweek_start),   //提交修改开始时间
        $("#xgweek_end").val(xgrweek_end)  //提交修改结束时间

}
function xgfz2() {
    // alert("加载完成");
    var xgp = $("#uphne_msg2").text()
    var xgpass = $("#upass_msg2").text()
    var xgmsg = $("#rmsg_msg2").text()
    var xgadd = $("#radd_msg2").text()
    var xgtype = $("#type_msg2").text()
    var xgtime = $("#rtime_msg2").text()
    var xgrtime = $("#rdtime_msg2").text()

    var xgrminute = xgrtime.slice(0, 2)
    var xgrhour = xgrtime.slice(3, 4)
    var xgrweek_start = xgrtime.slice(9, 10)
    var xgrweek_end = xgrtime.slice(11, 12)


    // console.log(lll);
    // alert(xgp);
    $("#xgmgdphone").val(xgp), //提交修改手机号
        $("#xgpasswd").val(xgpass),//提交修改密码
        $("#xgaddress").val(xgadd),//提交修改打卡地址
        $("#xguinfo").val(xgmsg),//提交修改打卡信息
        $("#xgyou_type").val(xgtype),//提交修改类型
        $("#xghour").val(xgrhour),     //提交修改小时
        $("#xgminute").val(xgrminute),    //提交修改分钟
        $("#xgweek_start").val(xgrweek_start),   //提交修改开始时间
        $("#xgweek_end").val(xgrweek_end)  //提交修改结束时间

}