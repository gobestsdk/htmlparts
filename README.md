#htmlparts

golang的web库

##一句话概括htmlparts
html=part1+part2+...partn，一种拼装html模块的模版引擎 
![](/img/htmlparts.png)
##基本思想:
在MVC中，htmlparts充当着呈现view的角色

在htmlparts的设计思想中，有2种不同类别的文件：
1.  后缀是.part的part文件
2.  后缀是.html的html文件

part文件，顾名思义，是不可以直接呈现给用户的局部html内容，web的control部分应当对part文件的路径请求做个屏蔽。

而html，则是既可以包含其它part，又可以含有html内容。web的control部分，直接每一个网页的url地址，都会对应一个html文件。

part集合和html集合，是m对n的满射。
 
## 关于part中可变化的部分
比如nav条，对不同的页面，nav条有的li元素是active，而其它li，应当是非active状态，这些工作，前端根据url.path完全可以自动适应，交给前端去判断就可以了。

## htmlparts语法示例

.html文件

index.html
``` html
<!--{{LoadTemplate(header.part)}}-->

<div class="content">
 <h1>hello word</h1>
</div>
<!--{{LoadTemplate(footer.part)}}-->

```

.part文件

header.part
``` html
<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="//cdn.muicss.com/mui-latest/css/mui.min.css" rel="stylesheet" type="text/css" />
  <link href="assets/css/style.css" rel="stylesheet" type="text/css" />
  <script src="//cdn.muicss.com/mui-latest/js/mui.min.js"></script>
  <script src="//code.jquery.com/jquery-2.1.4.min.js"></script>
  <script src="assets/js/script.js"></script>
</head>
<body>
```
## htmlpart关键词

### 1.LoadTemplate(part文件名.part)

```
<!--{{LoadTemplate(footer.part)}}-->
```
### 2.=Value

```
<!--{{=title}}-->
```
其中title通过
 htmlparts.Render(mainpage, hps, cs)传入contentstring=map[string]string
