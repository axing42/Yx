源码安装PHP时：
必备gcc autoconfig等、
在解压完php压缩包，会有一个php.ini-development文件
将它cp到php运行目录
cp ./php.ini-development /usr/bin/php/lib/(解压到php的lib目录下)
cd 刚刚复制php.ini目录下
mv php.ini-development php.ini(改名为php.ini);

//查看当前PHP配置文件路径
php -i | grep php.ini

安装swoole扩展时：
下载下来的swoole没有config文件，需要在当前swoole解压目录执行：
/usr/bin/php/bin/phpize 回车就会生成config 文件
