# 前端示例
这里的例子来自于 [w3school.howto](https://www.w3schools.com/howto/default.asp),这个网站是学习前端的最佳平台，可在线实验，需要翻墙访问;

有些demo演示需要用到file server, 可以直接编译并运行file_server.go;

## [html:控制布局](https://www.w3schools.com/html/default.asp)
- 最新语法[html5](https://www.w3schools.com/html/html5_intro.asp),
- h5新增的控件:< audio>,< video>,< svg>,< canvas>,
- h5新增语义标签：< section>,< header>, < footer>,< article>，
- h5表单新增属性:date,time,calendar,range;
- h5新增API: Drag&Drop, LocalStorage, Geolocation,ApplicationCache;
- html file path[解释](https://www.w3schools.com/html/html_filepaths.asp),推荐使用相对路径,举例："/images/picture.jpg"表示domain/images,而 "images/picture.jpg"表示当前目录下的images子目录；

## [JavaScript:负责交互](https://www.w3schools.com/js/default.asp)
- 最新语法[ES6](https://www.w3schools.com/js/js_es6.asp)


## [CSS:样式](https://www.w3schools.com/css/default.asp)
 - CSS = color and styles 
 - 最新语法css3
 - 添加css有[3种方式](https://www.w3schools.com/css/css_howto.asp):   1)external link 引入单独的css文件,比如:< link rel="stylesheet" href=absolute url or relative path>;
  2)internal: < style>xxx< /style>; 
  3)Inline CSS, 作为节点属性，比如< p style="display:none">;


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
- 最常用的是通过 .class,   #id 设置style;
- float:left  横铺
- display:block 竖铺, display:none 隐藏， display="" 显示;
- display:flex 列布局
- < br>换行
- < b>加粗< /b>
- 常用style：font-size,border,padding,margin,color,background-color;


## [常用标签tag](https://www.w3schools.com/tags/default.asp)
- [所有tag 列表参考官网](https://www.w3schools.com/tags/default.asp),tag推荐用小写，虽然大小写不敏感；
- [div和span区别](https://www.w3schools.com/html/html_blocks.asp),这2个都是group tag, div用来包装其它elements,span用来包装text文本，div会换行，span不会，它们的常用属性是:class,id,style;
- element是否会换行[参考分类](https://www.w3schools.com/html/html_blocks.asp)，比如< p>会换行，而< a>不会换行;
- < div id="xx" class="xx">xx< /div> 用div来包装若干个节点
- < h1>xx< /h1>  一级header
- < p>xx< /p>  段落，带换行
- < a href=url target=xx>xx< /a> 链接,target指定在哪个窗口打开，可以在指定的iframe打开[target取值参考](https://www.w3schools.com/html/html_links.asp);
- < form> 表单，由多个input组成的;
- < button onclick="myfunc">xx< /button> 按钮，可点击；
- < input type="text" name="username" value="">
- < input type="submit" value="Submit"> 用在表单里面
- < img src=url alt="alternative text" width="104" height="142">
- < ul>< li>xx< /li>< /ul> 列表

## 技巧总结
- node从无到有3种方式：1)新建节点，2）本来就存在，display="none",修改为"",3)先放一个占位的node，然后给innerHTML赋值;
- < a href=url,target="fr1">标签和< iframe name="fr1">配合使用，在指定区域打开连接；

## [Bootstrap:CSS framework](https://www.w3schools.com/bootstrap/bootstrap_ver.asp)


## [jQuery,deprecated](https://www.w3schools.com/js/js_jquery_selectors.asp)
- 现在没人使用jquery了，因为代码太难读了，都被js替换了，
[完成同样的功能，jquery对比javascript](https://www.w3schools.com/js/js_jquery_selectors.asp)