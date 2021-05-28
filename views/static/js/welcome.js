(function(){
    // 保证这个方法只在支持loadTimes的chrome浏览器下执行
    if(window.chrome && typeof chrome.loadTimes === 'function') {
        var loadTimes = window.chrome.loadTimes();
        var spdy = loadTimes.wasFetchedViaSpdy;
        var info = loadTimes.npnNegotiatedProtocol || loadTimes.connectionInfo;
        // 就以 「h2」作为判断标识
        if(spdy && /^h2/i.test(info)) {
            return console.info('本站点使用了HTTP/2');
        }
    }
    console.warn('本站点没有使用HTTP/2');
})();

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

