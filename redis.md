# redis 的特点：

**1）支持数据持久化（可以永久保存到磁盘上），重启时可以再次加载进行使用**

**2）也可以缓存到内存中**

**3）不仅仅支持 key-value 的键值对形式，更支持：**

**4）支持数据备份,既：master-slave(主从模式)的备份**

### list、set、zset、hash 等数据结构的存储

### 字符串形式的设置与获取

### key 类型：

**key _:返回所有的键名**
**?:匹配一个字符**
**_: 匹配任意个(包含 0 个)字符**
**[]:匹配括号中任一个字符，可以使用'-'表示范围**
**randomkey:返回随机的键名,而不是值**

**exists key:检查 key 是否存在**
**exists name:查询结果成功返回 1，失败返回 0**

**expire: 设置过期时间**
**expire 键名 过期时间(单位/秒)**
**expire name 30**

**ttl:查看过期时间**
**ttl 键名**
**ttl name(返回也是一个整数)**

**pexpire:设置过期时间(单位毫秒)**
**pexpire 键名(key) 过期时间**
**pexpire name 20000(20 秒)**

**pttl:以毫秒形式查看过期时间**
**pttl key**
**pttl name**

**type :查看键值对的类型**
**type key**
**type name**

**rename :修改 key 的名称(重命名)**
**rename key newKey**
**rename name newName**

**persist:在 key-value 有效期内设置，让他取消过期时间**
**persist key**
**persist name(如果 key 过期就不能取消)**

**del:删除一个键值对**
**del key**
**del name**

### 字符串类型

**他们都是字符串类型**

**获取**
**set 键名 值**
**set name dandan**

**setnx:设置一个不存在(存在将设置不了)的键值对**
**setnx key value**
**setnx name dandan**

**设置**
**get 键名**
**get name**

**mset:设置多个键值对**
**mset key value key2 value2 key3 value3**
**mset name dan age 18 email 1142170887@qq.com**

**mget:获得多个键值对**
**mget key1 key2 ...**
**mget name age email**

**msetnx :设置不存在的多个键值对**
**msetnx key value key2 value2(获取同上用 mget)**

**setrange:从索引位置替换或覆盖**
**setrange key 偏移位(字符串类型索引起始位置为 0) value(要替换或覆盖的值)**
**setrange name 0 lidan(从第 0 个位置替换覆盖原有的值)**

**getrange :截取字符串**
**getrange key start(开始位置) end(结束位置) end**
**getrange name 0 2(从 0 开始取两个)**

**getset:给旧值赋值一个新的值(删除旧值赋新值)，更改后结果返回旧的值**
**getset oldKey newValue**
**getset name lisi**

**setex:设置指定 key 的值，同时设置过期时间 单位秒**
**setex key seconds value**
**setex name 30 dandan(为 name 键设置 30 秒的过期时间)**

**psetex:设置指定 key 的值，同时设置过期时间 单位毫秒**
**setex key millseconds value**
**setex name 3000 dandan(为 name 键设置 3 秒的过期时间)**

append key value
append:往存在的字符串键值对后面追加值
append name hello(在原有数据上的后面追加)

incr key
incr:自身加一，必须是整型
incr age

incrby key increment
incrby:将 key 所存储的值加上指定增量值(增加多少)
incrby age 50(给 age 加 50)

decr key
decr:自减一
decr age

decrby key decrement
decrby:指定减去多少值
decrby age 20(减去 20)

strlen key
strlen:返回字符串长度
strlen name

incrbyfloat key increment
incrbyfloat:增加浮点数，小数点后面(set 应该不能做到增加浮点数?)
incrbyfloat price 5.99

# list 类型

**Redis 列表是字符串列表，按照插入顺序排序，可以从左(前)、右(后)添加元素**

**连续的列表数据，相当于索引数组，数组可以重复**

lrange key start stop
lrange:查看列表指定区间元素，区间以偏移量 start stop 而定
lrange tea 0 5(从 0 开始查看 5 个)
lrange tea 0 -1(查看所有)

lpush key value [value(代表可以插入多个)]
lpush:从前面插入到列表头部,插入起始位置好像是 0 开始
lpush tea value1 value2 value3

rpush key value [value]
rpush:从后面插入到列表尾部
rpush tea redtea greentea yellowtea

lpushx key value [value(只能添加一个)]
lpushx:将值 value 插入到列表头(只有当 key 为 list 类型并且存在才能完成插入)
lpushx tea redtea

rpushx key value [value(只能添加一个)]
rpushx:将值 value 插入到列表尾(只有当 key 为 list 类型并且存在才能完成插入)
rpushx tea greentea

lpop key
lpop :删除列表头的第一个元素并返回删除的值
lpop tea

rpop key
rpop :删除列表尾的第一个元素并返回删除的值
rpop tea

llen key
llen:返回列表 key 的长度,元素个数
llen tea(假入元素有 4 个就返回 integer(4))

lset key index value
lset:将列表下标为 index(数字，列表起始位置为 0)的 oldValue 改为新的 value
lset tea 4 dan(将 tea 列表的第 5 个改为 dan)

lindex key index
lindex:返回列表 key 中指定的下标的值
lindex tea 4

ltrim key start stop
ltrim:让列表只保留指定区间的元素，不在这个区间将会被删除
ltrim tea 2 4(保留第三个到五个的值)

linsert key before|after pivot(要设置的值) value
linsert:插入操作，位置指定值 pivot 之前或之后
linsert tea after name libai(在键值对中，在 name 之后添加一个 libai)

rem :remove
lrem key count value
根据参数 count 的值，移除列表与参数 value 的值

count>0:从表头开始向表尾搜索，移除与 value 相等的元素,数量为 count

count<0:从表尾开始向表头搜索，移除与 value 相等的元素,数量为 count

count=0:移除列表中与 value 相等的值，根据值移除
lrem tea 0 manggo(删除所有包含 manggo 得值)

# hash 类型

**Redis hash 是一个 string 类型 field 和 value 的隐射表,hash 特别适合储存对象，类型关联数组**

**使用场景：储存部分变更数据，如用户信息等**

hset key(通常为关联数组的下标索引) field value
hset:将哈希表 kry 中的字段 field 的值设为 value
hset userinfo username dandan
hset userinfo phpversion 7.0
key 类似于:
['Userinfo'=>[
    'name'=>'dandan'
]];

hget key field value
hget:查看某个field字段的值
hget userinfo phpversion

hmset key field value field2 value2
hmset:插入多个数据
hmset userinfo phpversion 7.0 MysqlVersion 5.6

hmget key field field2 field3
hmget:获取多个字段
hmget useringo phpversion MysqlVersion

hsetnx key field value
hsetnx :字段不存在才可以设置
hsetnx userinfo fruit(水果) banana

hkeys key
hkeys:返回指定哈希表所有字段(key键)
hkeys userinfo

hvals key
hvals:返回指定哈希表所有值
hvals userinfo

hgetall key
hgetall :获取指定key的所有字段和值
hgetall userinfo

hlen key
hlen :获取表中字段数量
hlen userinfo

hexists key field
hexists:查看指定哈希表的指定字段是否存在
hexists userinfo dan(查看蛋蛋这个字段是否存在)

hincrby key field increment
hincrby: 为哈希表key中的field字段增加整数

hincrbyfloat key field increment
hincrbyfloat :为哈希表key中指定字段field的浮点数增加指定浮点数量
hincrbufloat userinfo age 20

hdel key field field2
hdel :删除一个或多个哈希表字段
hdel userinfo name age phpversion

# set 类型

**Redis的Set是String类型的无序集合，集合成员是唯一的，意味着不能重复**

sadd key  member1 member2 member3
sadd:添加，将member添加到集合key中,一组可以添加多个但不能添加多个组
sadd php1 xinxin1 xinxin2 xinxin3

smembers key
smember:查看
smembers php1

srem key member
srem:移除集合中的member元素
srem php1 xinxin1

smembers key
smembers:返回集合中的所有成员
smembers php1

sismember key member
sismember :判断member元素是否存在于key中
sismember php1 xinxin2

scard key
scard；返回集合key的基数(集合元素中的数量,个数？)
scard php1

smove soure destination member
smove；将member元素从soure集合移动到destination集合中
smove php1 php2 xinxin1

spop key
spop:随机移除一个集合中的元素并返回被删除的
spop php1

sinter key key2
sinter:返回两个集合中相同的数据(交集)
sinter php1 php2

sinterstore destination key1 key2
sinterstore:将俩个交集保存到desination中
sinterstore php3 php1 php2

sunion key1 key2
sunion:返回并集(待理解)
sunion php1 php2

sunionstore destination key1 key2
sunionstore:将并集保存到 destination中
sunionstore test php1 php2

sdiff key1 key2
sdiff: 计算差集,确定前后关系，保留的是前面集合的元素
sdiff php1 php2

sdiffstore destination key1 key2
sdiffstore；将差集保存到destination中
sdiffstore test php1 php2

srandmember key
srandmember :返回集合中随机一个元素
srandmember php1

# sorted set String不可重复类型

**不同的每个成员，都会关联一个double类型的分数**

**通过分数进行成员排序，有序集合的成员是唯一的，但分数score可以重复**

****