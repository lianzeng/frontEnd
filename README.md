# 前端示例
这里的例子来自于 [w3school.howto](https://www.w3schools.com/howto/default.asp)

有些demo演示需要用到file server, 可以直接编译并运行file_server.go;

## [获取dom element的几种方式](https://www.w3schools.com/js/js_htmldom_elements.asp)
- document.getElementsByClassName() 选多个，根据class=xx
- document.getElementsByTagName() ，根据标签,比如"button","li","a",
- document.getElementById() 选某个元素,根据id=xx, 最常用，

- document.querySelectorAll("p.ab") 通过css选择,< p class="ab>;

- document.forms["form_id"],获取表单，再根据x.elements[i]获取填写的内容 

- 也可以逐层获取节点,根节点是document；

- innerHTML:元素内容，很重要的属性；

## [一些最常用的DOM操作](https://www.w3schools.com/js/js_htmldom_document.asp)
- document.getElementById(id)
- element.innerHTML =  new content
- element.style.property = new style, eg: element.style.display="none;
- document.appendChild(element)
- document.removeChild(element);
- element.addEventListener("click",func(){ });
- document.documentElement:	Returns the < html> element;
- document.forms:	Returns all < form> elements;
- document.readyState:	Returns the loading status of the document;
- document.referrer:	Returns the URI of the referrer (the linking document);
- document.inputEncoding:	Returns the document's encoding (character set)
- document.baseURI/URL/title/body/cookie/documentURI/domain/head/images/links/scripts;


## [定时器](https://www.w3schools.com/js/js_timing.asp)
- t = setTimeout(function, milliseconds)， 1次性定时器
- clearTimeout(t);
- t = setInterval(function, milliseconds) 周期定时器
- clearInterval(t);

## [Ajax ](https://www.w3schools.com/js/js_ajax_intro.asp)
- 让浏览器发送XMLHttpRequest到webserver来获取data,并通过设置好的回调函数（异步）处理response data,这些都是通过javascript完成的；
- AJAX = Asynchronous JavaScript.


## 常见的style
- float:left  横铺
- display:block 竖铺, display:none 隐藏， display="" 显示;
- display:flex 列布局
- .class 通过class 设置
- #id  通过id设置；
- < br>换行
- < b>加粗< /b>

## [jQuery,deprecated](https://www.w3schools.com/js/js_jquery_selectors.asp)
- 现在没人使用jquery了，因为代码太难读了，都被js替换了，
[完成同样的功能，jquery对比javascript](https://www.w3schools.com/js/js_jquery_selectors.asp)