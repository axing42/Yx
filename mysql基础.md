```sql
-- is null: 检测字段为null
select * from 表名 where id = is null;
select * from 表名 where id = is not null;

-- 值为null的时候，给它设为一个默认值
select ifnull('字段名','给定的默认值') from 表名;
SELECT IFNULL(`name`,'无知') FROM goods;

-- distinct:去掉查询结果的重复行(两行数据完全一样的)
-- 去重必须指定字段不然无效
select distinct 字段1 字段2 from tableName;

SELECT DISTINCT class_name FROM `class`;

select 的or与and
or:是或者(||)的关系，满足一个条件就出结果
and:是并且(&&)关系,条件必须满足才出结果

-- 通过命令行导入sql文件
source 然后拖动文件到命令行 回车


-- 创建表并复制表tableNametaget结构
create table 新表名 like 要copy的表
-- 创建表并复制数据与结构
create table NewTableName select * from oldTableName;
-- 创建表结构并复制数据
insert into 表名 select * from 其他表名;
-- 插入别的指定表字段
insert into 表名(字段名) select 要取的字段 from 表名;
-- 建表时插入其他表的指定列,并且新表取旧表的字段要与被取表字段的字段名对应，可以取别名
create table tableName(id int primary key auto_increment aname varchar(30))
select cname as aname from tabnametaget

-- where:条件筛选+查询，as:取别名
select * from tableName;
select id,name as cname from tableName;
select * from TableName where id = 1;
-- like:模糊查询,查询包含蛋蛋的字段，效率低下不建议使用
-- 一个百分号%开始表示以蛋蛋开始，两个表示任何位置出现蛋蛋
-- 一个百分号结尾表示以蛋蛋结束
select * from TableName where name like '%蛋蛋%';
-- and:并且
select * from TableName where name like '%蛋蛋%' and id = 1 and num=95;
-- NOT like取反，不包含
select * from TableName where name not like '%猫%';
-- CONCAT:字段连接,将两个同表字段组合在一起
select CONCAT(字段一，字段二) from tableName;
SELECT concat(sname,leirong) FROM `class` ;
-- OR:满足一个条件就成功,满足的条件都返回
select * from 表名 where id=5 or `name` like '%蛋%';
-- AND: 都满足才成功
select * from 表名 where id=5 and name like '%蛋%';
-- DISTINCT:搜索结果去重
 SELECT DISTINCT class_id FROM class;
-- between:范围查询
select * from 表名 where age>=20 and age<=30;
select * from class where between 20 and 30;
-- NOT:查询取反
select * from class where not between 20 and 30;
-- IN:查询两个指定条件，id=1与id=3的字段
select * from 表名 where id in (1,3);
-- NOT IN:结果取反
select * from 表名 where id NOT IN (1,3);
-- order BY 字段名 ASC:对字段升序排列desc降序从大到小
select * from 表名 order by id desc;
-- limit:从查询字段里取指定数据条数
-- limit 0,1:第一个代表从哪开始取，第二个参数表示取多少个
select * from 表名 order by id asc limit 1;

-- 更新
-- update
select 表名 set 字段名=值 where id=10 or id is null;
-- 更新字段里没有生日的，值为当前时间
UPDATE sta SET brithday=NOW() WHERE brithday is null;

-- delete：删除
delete from 表名 where id=1;
DELETE FROM sta WHERE sname like '%a%';

-- 插入数据,字段名要与值对应
insert into 表名 set 字段名 = 值,字段2 = 值2 ;
insert into 表名 (字段名1字段名2) values(值1，值二),(值1，值2);

--
```

# mysql 字符集的 utf8_general_ci:

### 它不区分大小写字母，对查询排序没有影响

## msyql 字符串函数

```sql
-- left 1:从左取一个字符
SELECT LEFT(sname,1) FROM sta;
-- right 1:从右取一个字符
SELECT RIGHT(sname,1) FROM sta;
-- mid:从第一个字符取俩字符
SELECT MID(sname,1,2) FROM sta;
-- 从第2个字符开始取所有
SELECT MID(sname,2) FROM sta;

-- substr:从第一个取两个字符
SELECT SUBSTR(sname ,1,2) FROM sta;
-- CHAR_LENGTH:获取字段字符串的长度
SELECT CHAR_LENGTH(sname) FROM sta;
-- concat: 字符串拼接
SELECT concat('编号:',id) as id , concat('姓名:',sname) FROM sta;

-- if:判断字段最大值是否为400大于了就显示文字
select sname,if(click > 400,'我最大',click) FROM sta;

-- set:添加某个字段时，set相当于多选项
ALTER TABLE sta ADD flag SET('推荐','热门','置顶','图文');
-- set更新时，类似枚举类型，更新1就表示推荐，2就代表热门
UPDATE sta SET flag = 1 WHERE id = 20;
-- set插入多条也只能为设置好的四个选项字符不能为数字
insert INTO sta(sname,flag) VALUES('dandan','热门,推荐');
-- 查询是只能FIND_IN_SET,查询值不能为数字只能为字符
SELECT * FROM sta WHERE FIND_IN_SET('推荐',flag);
```
