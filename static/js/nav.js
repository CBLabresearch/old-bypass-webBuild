$(function(){
    console.log("开始");
    // 在键盘按下并释放及提交后验证提交表单
    $('#reg_form').validate({
        errorPlacement:function(error, element){
            console.log(error);
            console.log(element);
            console.log(element.attr( "id" ));
            $( element )
                .closest(".form-group")
                .find( "label[for='" + element.attr( "id" ) + "']" )
                .append( error );
        },
        rules:{
            name:{
                required:true,
                minlength:6,
                maxlength:15
            },
            realname:{
                required:true
            },
            password:{
                required:true,
                minlength:5
            },
            password2:{
                required:true,
                minlength:5,
                equalTo:"#reg_password"

            }


        },
        messages:{
            name:{
                required:"(必须字段)",
                minlength:"(必须由6-15个字符组成)",
                maxlength:"(必须由6-15个字符组成)"

            },
            realname:{
                required:"(必须字段)"
            },
            password:{
                required:"(必须字段)",
                minlength:"(密码长度不能小于五个字符)"
            },
            password2:{
                required:"(必须字段)",
                minlength:"(密码长度不能小于五个字符)",
                equalTo:"两次密码输入不一致"

            }

        },
        submitHandler:function(form){
            alert("提交事件!");
            form.submit();
        }

    });
});
$("#reg_user").blur(function(){
    console.log("验证用户名是否被注册");
    var name=$(this).val();
    console.log(name);
    $("label[for='ok']").remove();
    $("label[for='error']").remove();
    //ajax
    if(name.length>=6){
        $.ajax({
            type:"get",
            url:"validuser.do?name="+name,
            dataType:"json",
            success:function(result){
                console.log(result);
                var pos=$("label[for='reg_user']");
                if(result["code"]=="100"){
                    console.log("可以被注册");
                    var info="<label for='ok' class='ok'>可以注册</label>";
                    pos.append(info);

                }else if(result["code"]=="200"){
                    console.log("已被注册");
                    var info="<label for='error' class='error'>已被注册</label>";
                    pos.append(info);
                }
            },
            error:function(error){
                console.log("请求失败！");
            }

        });
    }

})
