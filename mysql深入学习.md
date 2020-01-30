```sql
-- 代金券
union 语法 : sql语句 union sql语句2
-- union作用：
-- 可以将两次查询的结果集合并为一条sql ,查询出来后可以使用order by 排序
-- 两表查询字段的字段类型可以不一样
select * FROM zh_user_comments WHERE id=1
UNION select * from zh_user_comments WHERE id=3;
-- 可同时连接两张表，连接的是结果集
SELECT * FROM `zh_article` where comment_id = 1
UNION
SELECT * FROM `zh_user_comments` WHERE article_id = 1;
-- 使用UNION注意事项：
-- 查询的两张表字段数量要一直，不然报错
-- 查询出来的结果默认去重，想要查看重复的使用:
union all


-- select 使用流程控制的switch case
SELECT id,content,
CASE status 
WHEN 1
THEN '正常'
WHEN 0
THEN '异常'
END
FROM zh_user_comments;
-- if(expr1,expr2,expr3) :expr1表达式结果为真是取expr2的值
-- 表达式判断时要使用一个<!等于符号>
SELECT id,content,IF(`status` =1, '正常','异常') FROM zh_user_comments;

-- ifnull:判断第一个表达式是否为null，为null走第二个表达式，否则输出自身
SELECT id,IFNULL(reply_id,'空')from zh_user_comments;


-- View视图：查询里常常把查询结果当成临时表来看,创建之后是一张新的表,但看不见
-- 它可以当做一张虚拟表，是通过某种运算得到的一个投影
-- 视图的创建：建视图时，不需要指定视图列名与列类型，只是继承了表数据？
-- 创建之后，表的数据改变视图也会随之变化
-- 好处：把复杂的sql语句先分解成视图再查询
-- 语法：
create view 视图名 As select 语句

-- view的algorithm：指定合并还是建立临时表,在创建了简单视图，不需要创建临时表就
-- 可以选择使用algorithm来选择建立合并
algorithm:merge : 合并查询语句(将视图的查询结果保存起来)
algorithm:temptable 临时表
algorithm:undefined 未定义，有系统决定


```