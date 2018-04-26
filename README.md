# lib4go

对常用功能进行封装提供友好的公共库代码，供其它golang程序调用

基本库
* 线程安全的map
* 编码处理（gbk,utf-8编码转换,base64,hex）
* 加解密（md5,des,rsa,sha1等）
* 图片库（绘制图片，文字）
* 系统资源（获取CPU，内存，硬盘使用情况）

组件封装

* 日志记录（文件，MQ）
* 数据库操作（oralce,sqlite）
* influxdb操作（存，取）
* MQ操作（发布，订阅等）
* scheduler（基于cron的任务处理）
* LUA脚本引擎
* HTTP请求(支持证书,cookie等)
* HTTP Server
