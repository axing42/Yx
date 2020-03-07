```vb
# 查询已安装的软件包
rpm -qa | grep 软件包名

# 卸载软件包
rpm -e 卸载软件的完整名称

# 安装软件包:安装必须在当前软件包路径
rpm -ivh 软件名称
i:表示安装 install
v:显示进度条
h:进度条以‘#’号显示



# yum(w自动解决依赖关系)安装软件
yum -y install 安装软件名称
-y 确认安装不提示

repo文件是Fedora中yum源(软件仓库)的配置文件
配置目录：/etc/yum.repos.d/CentOS-Media.repo


# 关闭防火墙
他叫：iptables
service iptables stop|start|restart|status

# 解压缩包
.gz gzip tar zxvf
.bz bzip tar jxvf

gzip压缩文件
gzip 文件名:压缩后文件后缀为gz
gunzip 文件名:解压

# 查看命令帮助
tar --help | grep z

语法：tar -zxvf 文件名
-c：产生.tar打包文件
-v:显示详细信息
-f:指定压缩后的文件名
-z:打包同事压缩
-x:解压文件
tar 压缩后新生成文件不会覆盖原文件
tar 可选项 压缩后的文件名(它的后缀为.tar.gz??) 被压缩文件
tar -zcvf new.php.tar.gz test.php

tar: 解压
tar -zxvf new.php.tar.gz

tar -t 查看压缩文件
tar -ztvf new.php.tar.gz

zip: 压缩文件
zip 压缩的文件名 被压缩的文件
zip test.php.zip test.php

zip -r :从目录压缩文件
zip -r 压缩的文件名 被压缩的目录名

unzip: 解压缩
unzip 要解压的文件名

bzip2 :压缩后源文件不保留
bzip2 要压缩文件名
bzip2 -k 文件名 //压缩文件并保留

bunzip2 :解压
bunzip2 文件名

```
# 进程管理

```vb
ls /etc/init.d:查看守护进程

service iptables start/stop/restart/status: 开启/关闭/重启/状态.防火墙

chkconfig --list 进程名:查看当前进程是否开启

mysqld  0:off   1:off   2:on    3:on    4:on    5:on    6:off
0表示关机模式，1表示单用户模式，2无网络多用户模式
3有网络多用户模式 4不可用模式 5带图像界面的多用户模式
6重启模式

查看所有进程
ps -aux:详细进程
ps -e：简易的进程信息
top：实时进程

ls | grep t.zip:查找包含t.zip字符的文件

kill:杀死进程
kill 进程pid
kill 999
kill -9:强制杀死进程

ps -A | grep 软件进程

crond:定时任务，守护进程
分 时 日 月 周
命令格式：分minute 时hour 日day 月month 周week command
周：0或者7代表星期日
command:要执行的命令
(*):所有可能的值，例如month字段如果是*,则表示在满足其他字段条件后每月都执行该命令
(,):逗号隔开的值代表一个列表范围，如'1,5,7,9';
(-):中杠表示一个整数范围，例如'1-6'表示1,2,3,4,5,6;
(/):斜线指频率，例如'0-23/2'表示每两小时执行一次，同时星号可以和/使用
例如'*/10',如果用在minte字段表示每十分钟执行一次

配置文件：/etc/crontab
crontab -l：查看定时任务及进程
创建编辑定时任务
crontab -e(要写绝对路径)

写一个定时任务：每分钟自动写入当前时间
*/1 * * * * date >> /root/time.txt

/etc/profile: 环境变量的配置文件目录

du:查看文件目录磁盘使用情况
du -s 文件
du -sh 文件

# 文件修改防火墙开启端口
vim /etc/syscomfig/iptables

```
安装tp5指定版本
composer create-project topthink/think 要放的目录名称 5.0.* --prefer-dist

## 关于php在linux设置操作快捷命令

cd ~/.bash_profile
加入以下命令 yy(复制当前行) p(粘贴到下一行)
alias php=php的结对路径








