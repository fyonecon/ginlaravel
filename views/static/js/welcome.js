
(function (){
    let welcome_num = 0;
    let welcome = setInterval(function (){
        try {
            view.notice_txt("静态文件引用成功！", 3500);
            clearInterval(welcome)
        }catch (e){
            if (welcome_num > 10){
                clearInterval(welcome)
                alert("静态文件引用成功！")
            }
        }finally {
            welcome_num ++
        }
    }, 4000);
})();

