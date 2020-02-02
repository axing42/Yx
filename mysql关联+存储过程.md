```sql
-- 当数据更新时，指定字段自动更新时间

alter table 表名 add 字段名 字段类型(timestamp) default 添加字段自动记录当前时间(current_timestamp) on update
on update ：自动更新时间

SELECT TIMESTAMPDIFF(DAY,brithday,NOW()) AS '从出生到现在经过了多少天' FROM sta WHERE brithday !='';

-- 查询两个表指定字段
SELECT bj.sname,bj.sex,bj2.sname FROM bj,bj2 WHERE bj2.class_id = bj.bj_id;

-- 两表关联
SELECT * FROM bj as b INNER JOIN bj2 as b2
ON b.id=b2.class_id;

-- 与左边的表关联：bj
SELECT * FROM bj as b
LEFT JOIN article AS a
ON b.id=a.id
WHERE a.`status` =''
;

-- 查询同一张表 all:不过滤重复的
(SELECT * FROM bj)
UNION ALL
(SELECT * FROM bj)
;

(SELECT * FROM class LIMIT 2)
UNION ALL
(SELECT * FROM class LIMIT 1)
;

(SELECT sname FROM bj where sex=2)
UNION ALL
(SELECT sname FROM class)
limit 2; //总的查询数取两条

-- 删除没有班级的学生
DELETE b FROM bj as b
LEFT JOIN class as c
ON c.class_id=b.id
WHERE c.class_id is NULL;

-- 关联更新
a表(id,sex,par,c1,c2)
b表(id,sex,par,c1,c2)
-- 第一种
-- 将b表的c1,2更新到a表中并且b表的年龄大于50岁
update a,b set a.c1=b.c1,a.c2=b.c2 where a.id=b.id and b.age>50; 
-- 第二种
update a inner join b on a.id=b.id
set a.c1=b.c1,a.c2=b.c2 where b.age>50;
```

# mysql 高级

```php
    //    验证规则并过滤数据
       $vali = [
               'id' => $id,
       ];
       $validate = new Common();
       $te = $validate->check($vali);
       if ($te) {
           $data = BannerModel::with(['items', 'items.img'])->find($id);

           return $data;
       } else {
           $this->return_msg(400, '参数错误!或没有数据');
       }
```

```sql
-- 给触发器取名字
-- 监听动作 after before
-- 监听事件 insert update dalete
-- new：每insert一下都会出现新的一行，new的就能监听新增的数据

CREATE TRIGGER tt
AFTER INSERT
ON ord FOR EACH ROW
BEGIN
UPDATE goods set num=num-new.much WHERE gid=new.gid;
END;

create trigger t2
after delete
on ord fro each row
begin
update goods set num=num+old.much where gid=old.gid;
end;

-- MySQL存储过程：把若干sql语句存储起来，类似于函数的封装，但是函数有返回值，存储过程没有返回值；
-- 查看存储过程：show procedure status;
-- 删除：drop procedure 存储过程的Name;

create procedure procedureName()
begin
sql语句
end;

-- 声明变量：declare 变量名 变量类型 default 可选的默认值
CREATE PROCEDURE p2()
BEGIN
DECLARE age TINYINT DEFAULT 18;
SELECT CONCAT('我的年龄是：',age);
END;

-- 运算
CREATE PROCEDURE p3()
BEGIN
DECLARE age TINYINT DEFAULT 18;
SET age := age+20;
SELECT CONCAT('我的年龄是：',age);
END;

-- 流程控制
CREATE PROCEDURE p1()
BEGIN
DECLARE ss TINYINT DEFAULT 18;
IF ss >= 18 THEN
SELECT '你成年了！';
ELSE
SELECT '你还是个孩子！';
END IF;
END;

-- 存储过程传参：
-- 传参语法：存储过程的名字(in/out/inout 变量名 变量类型)
-- in:接收型 接收调用时的参数，out：输出型，在调用存储过程的时候定义一个变量@var接收，然后可以
-- select @var; 打印查看
-- inout：既可以输出又可以接收, 在调用的时候必须先定义一个变量，然后再把变量传到储存过程里

CREATE PROCEDURE p1(n1 INT,n2 TINYINT)
BEGIN
SELECT CONCAT('传的参数乘积为：',n1 * n2) AS '乘积';
IF n1 > n2 THEN
SELECT 'N1大于10!';
ELSE
SELECT 'N1小于10！';
END IF;
END;


-- 函数：存储过程可以有也可以没有返回值，而函数必须只有一个返回值
-- 参数格式：参数名 参数类型
-- 函数体结束一定要有return，没有就报错
-- 函数的查看：show create function 函数名
-- 删除：drop function 函数名
-- 语法：
create function functionName(参数1,参数2) returns 返回类型
begin
    sql语句
end
-- 调用：select 函数名（参数）

-- 无参数的函数例子：查询表中数据总和(多少条数据)并用变量sum返回
create function func() returns int
begin
    declare sum int default 0; -- 定义变量
    select count(*) into sum   -- into: 将查询结果保存到变量sum
    from banji;
    return sum;
end;
select func();

-- 带参数函数:输入用户名并查询tp.user表中学生的零花钱,有就返回零花钱没有返回str提示信息
create function funcPrice(uname varchar(10)) returns int
begin
    set @var:=0;
    set @str:='没有查询到数据！';
    select price into @var 
    from user where `name` = uname;
    if @var = 0 then
    return @str;
    else return @var;
    end if;
end;
select funcPrice('阿鑫');

-- 触发器案例： 当文章表被删除，下面的评论也跟着删除
CREATE TRIGGER dele
AFTER DELETE
ON zh_article for EACH ROW
BEGIN
DELETE FROM zh_user_comments WHERE article_id = old.id;
END;
```
