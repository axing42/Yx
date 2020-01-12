```php
/***********实例化接口助手函数*************/
$request = Request::instance();

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
```