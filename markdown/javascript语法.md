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
#