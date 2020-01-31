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
