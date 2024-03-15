{
"title": "Typescript急速入门",
"top": true,
"toc": true,
"category": "Typescript",
"tags":["编程语言","Typescript"],
"preview": "快速搭建Typescript编程环境，能使用Typescript开发一些功能。"
}
---
# 类型
## 强类型和弱类型
通过以下代码体会Javascript的弱类型
```
let num = 1;
console.log(typeof num);
num = "1";
console.log(typeof num);
num = true;
console.log(typeof num);
```
## 动态类型和静态类型
运行期间做数据类型检查。在编译期间做数据类型检查。（通过数据类型检查，可以判断一种操作是否能被执行，非常关键）。