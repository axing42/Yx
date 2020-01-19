```vb
关机:
shutdown -h now,shutdown -h 1:1:一分钟后关机
halt
poweroff
init 0;
重启：
shutdown -r now
reboot
init 6;

more /proc/cpuinfo：查看cpu信息
more /proc/meminfo:查看内存信息
more /proc/version；查看linux版本

mv:改文件文件名,它具有移动、剪切，改名等功能
mv 被移动更改的文件名 新文件名

ls -lh:查看文件并显示大小，以kb、mb、gb形式

cat -n: 从前往后查看文件并显示行号
cat -n /etc/issue
tac：从后往前查看一个文件
tac /etc/issue

nl:直接查看文件并显示行号
nl 1.php

more:分页显示文件,上一页快捷键：b,下一页f,退出：q
显示当前行号：':f'
more 文件名

find:文件查找
find 要查找的目录 按文件名查找 查找的文件名,例子
find /al -name 1.php 

which:查找命令是否存在，找到的命令包含所在目录
which ls


```

```vb
快捷键：
ctrl+a:光标移动到命令行首，开头
ctrl+e:移动到命令行命令行
ctrl+c:结束命令执行,终止当前进程
ctrl+d:结束命令执行,退出

```
# 远程链接
```vb
客户端链接选择22端口，服务器也要开这个端口，地址填URL、服务器ip
用户名一般为root 密码。。。
ssh 用户名@ip 回车，输入密码链接
```