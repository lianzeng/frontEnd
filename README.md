# 前端示例
这里的例子来自于 [w3school.howto](https://www.w3schools.com/howto/default.asp)

有些demo演示需要用到file server, 可以直接编译并运行file_server.go;

## [获取dom element的几种方式](https://www.w3schools.com/js/js_htmldom_elements.asp)
- document.getElementsByClassName() 选多个，根据class=xx
- document.getElementsByTagName() ，根据标签,比如"button","li","a",
- document.getElementById() 选某个元素,根据id=xx, 最常用，

- document.querySelectorAll("p.ab") 通过css选择,< p class="ab>;

- document.forms["form_id"],获取表单，再根据x.elements[i]获取填写的内容 

- 通过[dom节点之间的关系](https://www.w3schools.com/js/js_htmldom_navigation.asp)来查找,根节点是document,  常见的关系：parentNode/childNodes[i]/firstChild/lastChild/nextSibling/previousSibling;



## [一些最常用的DOM操作](https://www.w3schools.com/js/js_htmldom_document.asp)
- document.getElementById(id) , 根据id定位node;
- element.innerHTML 元素内容，非常重要的属性；
- element.style.property = new style, 比如: element.style.display="none;
- element.addEventListener("click",func(){ }) 添加事件回调;
- document.documentElement:	表示整个 < html> element;
- document.forms:获取表单内容;
- document.readyState == 4 加载完成;
- document.referrer:the URI of the referrer (the linking document);
- document.inputEncoding: 字符编码(character set)；
- document.createElement("p") 创建一个< p>
- parent.appendChild(node), parent.insertBefore(newnode,child) 添加节点;
- 常用属性document.head/title/body/cookie/images/links/scripts/documentURI/domain/baseURI/URL;


## [定时器](https://www.w3schools.com/js/js_timing.asp)
- t = setTimeout(function, milliseconds)， 1次性定时器
- clearTimeout(t);
- t = setInterval(function, milliseconds) 周期定时器
- clearInterval(t);

## [Ajax：和server交互 ](https://www.w3schools.com/js/js_ajax_intro.asp)
- 让浏览器发送XMLHttpRequest到webserver来获取data,并通过设置好的回调函数（异步）处理response data,这些都是通过javascript完成的；
- AJAX = Asynchronous JavaScript.


## [CSS常见的style](https://www.w3schools.com/css/css_intro.asp)
添加css有[3种方式](https://www.w3schools.com/css/css_howto.asp):1)external link 引入单独的css文件;2)internal < style>; 3)Inline CSS, 作为节点属性;
- float:left  横铺
- display:block 竖铺, display:none 隐藏， display="" 显示;
- display:flex 列布局
- .class 通过class 设置
- #id  通过id设置；
- < br>换行
- < b>加粗< /b>


## 常用标签tag
- < div id="xx" class="xx">xx< /div> 用div来包装若干个节点
- < h1>xx< /h1>  一级header
- < p>xx< /p>  段落，带换行
- < a href=url>xx< /a> 超链接
- < form> 表单，由多个input组成的;
- < button onclick="myfunc">xx< /button> 按钮，可点击；
- < input type="text" name="username" value="">
- < input type="submit" value="Submit"> 用在表单里面

## [jQuery,deprecated](https://www.w3schools.com/js/js_jquery_selectors.asp)
- 现在没人使用jquery了，因为代码太难读了，都被js替换了，
[完成同样的功能，jquery对比javascript](https://www.w3schools.com/js/js_jquery_selectors.asp)