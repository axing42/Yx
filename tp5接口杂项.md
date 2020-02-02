```js
/***********实例化接口助手函数*************/
// tp5.1模型更新数据注意事项：更新条件的where一定要写在更新条件的前面!!!
// 在使用layer弹窗时，如果使用ajax异步请求数据，再使用弹窗将不能同步接受到后端返回的信息
// 应该是异步的结果，他们是同时进行的，而弹窗确实在等待执行所以只能手动设置提示信息 如：

if (data.status == 1) {
                layer.msg('注册成功！正在跳转...', {
                icon: 6,
                time: 1000
                });
                window.location.href = "{:url('index/index/index')}";
                } else {
                layer.msg('注册失败！请检查...', {
                icon: 5,
                time: 1000
        });
}
```
```php
$request = Request::instance();
echo '请求方法是：'. $request->method().'<br/>';
echo '请求的地址是：' .$request->ip().'<br/>';
echo '请求参数:';
dump($request->param());
echo '请求参数：仅包含name';
dump($request->only(['name']));
echo '请求参数：排除name';
dump($request->except(['name']));

/*************判断请求方法******************/
        if ($request->isGet()) echo '您用get请求了接口';
        if ($request->isPost()) echo '您用post请求了接口';
        if ($request->isPatch()) echo '您用patch请求了接口';
        if ($request->isPut()) echo '您用Put请求了接口';
        if ($request->isDelete()) echo '您用delete请求了接口';

/***********验证接口数据***********/
use think\Validate;

$guize = [
                'name'  => 'require|max:5|chs',
                'age'   => 'number|between:1,100',
                'email' => 'email',
        ];
        $msg = [
                'name.require' => '名称必填！',
                'name.chs'     => '只能为中文！',
                'name.max'     => '最大长度不能超过5个字符',
                'age.number'   => '年龄只能为数字！',
                'age.between'  => '年龄不能超过100岁！',
                'email'        => '邮箱格式错误',
        ];
// 获取表单提交过来的数据
        $data = input('post.');
// 实例化验证规则
        $validate = new Validate($guize, $msg);
// 验证
        $res = $validate->check($data);
        if (!$res) {
            dump($validate->getError());
        }
    }

/***********路由笔记！！！*******/
//Route::rule('test','user/index/user');//模块/控制器/操作方法:记死
//Route::rule('test','user/index/user','GET|POST');
//如果路由指定了要获取的参数那么在地址栏必传不传报错
Route::get('user/:id','user/index/user');

$all = $request->param();
        获取指定参数 param如果不填参数默认接受所有
        只获取地址栏参数将param改为get() 获取路由参数更改为route
        对以上获取参数的方式简写为input
        $all = input('param.id');
        $id = Request::instance()->param('id');
        $all = Request::instance()->get();
        print_r($all);

原生查询
            $data = Db::query('select * from banner where id=?',[$id]);
            链式查询
            $data = Db::table('banner')->where('id', '=', $id)->select();
            闭包
            $data = Db::table('banner')->where(function ($res) use ($id) {
                $res->where('id', '=', $id);
            })->select();

对象笔记：在没有继承使用的对象前，使用继承类的方法不能$this->，只能实例化new出来!!!!

临时隐藏某个字段
$collection = collection(要隐藏的模型并且已执行方法调用数据库) 返回是一组模型对象;
$product = $collection->hidden([要隐藏的字段]);

filed：第二个参数为true代表过滤或隐藏当前字段

$recount = Db::table('product')
->limit($cc)
->order('create_time desc')->field(['id','from','img_id'],true)
->select();


<!--  需要继承时使用：  -->
<!--  {extend name="public/base"} name填当前公共模板的文件路径  -->
<!--  使用占位符内容:  -->
<!--  {block name='test'} 写你要替换的内容,注意name里面的值引号要对应不能一个单一个双 {/block}  -->
<!--  并且使用某个块标签必须使name的值对应  -->
<!--  块级标签也可以使用标签语法include加载文件  -->
<!--  如果某些块级元素的内容想要在继承里面使用：  -->
{block}
{__block__}
{/block}

```
