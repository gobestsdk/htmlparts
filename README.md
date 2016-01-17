#htmlparts

golang的web库

##一句话概括htmlparts
html=part1+part2+...partn，一种拼装html模块的模版引擎 

##基本思想:
在MVC中，htmlparts充当着呈现view的角色

在htmlparts的设计思想中，有2种不同类别的文件：
1.  后缀是.part的part文件
2.  后缀是.html的html文件

part文件，顾名思义，是不可以直接呈现给用户的局部html内容，web的control部分应当对part文件的路径请求做个屏蔽。

而html，则是既可以包含其它part，又可以含有html内容。web的control部分，直接每一个网页的url地址，都会对应一个html文件。

part集合和html集合，是n对1的单射（映射模型）。
