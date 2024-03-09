{
"title": "现代Javascript",
"top": true,
"toc": true,
"category": "Javascript",
"tags":["浏览器前端","Javascript"],
"preview": "Javascript的语法梳理"
}
---

# 将Javascript引入HTML中
## 通过 `<script>` 标签
所有属性的目的是按照约定“告诉”浏览器如何处理脚本文件，即什么时候下载，解析执行。
### 标签位置
过去放在`<head>`中，但是这样等到所有 JavaScript 代码都下载、解析和解释完成后，才能开始渲染页面。会出现明显白屏。
### 推迟执行脚本defer
这个属性表示脚本在执行的时候不会改变页面的结构,浏览器可以立刻下载脚本。浏览器可以先下载源代码，等页面渲染完成后再执行。
### 异步执行脚本async
浏览器要立刻下载脚本，渲染后续DOM结构。
### 动态加载脚本
即通用Javascript代码动态添加`<script>`标签。在将`<script>`元素添加DOM时浏览器不会立即发出请求或执行其内部的代码，相当于async模式。\
动态加载影响页面渲染速度。要想让预加载器知道这些动态请求文件的存在，可以在文档头部显式声明它们：\
`<link rel="preload" href="dynamic.js">`
## 行内代码和外部文件
# 语言基础
## 语法
类似C,但比C宽松。
### 区分大小写
### 严格模式
### 语句
分号，多条语句块。
## 关键字和保留字
## 变量
ECMAScript 变量是松散类型的，意思是变量可以用于保存任何类型的数据。\
有 3 个关键字可以声明变量： var 、 const 和 let 。
### 声明风格及最佳实践
1. 不使用 var
2. const 优先， let 次之
## 数据类型
 6 种简单数据类型：Undefined, Null, Boolean, Number, String, Symbol。\
 1 种复杂数据类型：Object。\
 因为 ECMAScript 的类型系统是松散的，所以需要一种手段来确定任意变量的数据类型。
 ### typeof操作符
 对一个值使用 typeof 操作符会返回下列字符串之一：
-  "undefined" 表示值未定义；
-  "boolean" 表示值为布尔值；
-  "string" 表示值为字符串；
-  "number" 表示值为数值；
-  "object" 表示值为对象（而不是函数）或 null ；
-  "function" 表示值为函数；
-  "symbol" 表示值为符号
### Boolean类型
其余5种数据(不包含Symbol)类型转换成false值的情况：
- string  ""
- Number 0，NaN
- Object null
- Undefined undefined
- Null  null
### Number类型
NaN值的产生，判断方法。\
数值转换\
其余5种（不包括Symbol）转换成数值类型的函数。\
可能出现的情况。\
### String类型
字符串是不可变的。\
其它4种类型转换成字符串的方式。\
模板字符串。\
### Symbol类型
符号是原始值，且符号实例是唯一、不可变的。符号的用途是确保对象属性使用唯一标识符，不会发生属性冲突的危险。\
符号没有字面量语法，这也是它们发挥作用的关键。按照规范，你只要创建 Symbol() 实例并将其用作对象的新属性，就可以保证它不会覆盖已有的对象属性，无论是符号属性还是字符串属性。\
凡是可以使用字符串或数值作为属性的地方，都可以使用符号。\
### Object类型
每个Oject实例都有1个属性6个方法。
## 操作符
一般操作符的对象是数值类型，在ES6以后可以操作字符串和对象类型。\
## 语句
### if语句
### do-while语句
###