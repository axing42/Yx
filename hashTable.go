package main

import "fmt"

// 定义存放哈希表值得结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

type EmpLink struct {
	Head *Emp
	lens int
}

func (e *EmpLink) Find(id int) *Emp {
	cur := e.Head
	for cur != nil {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
	return nil
}
func (e *EmpLink) Insert(emp *Emp) {
	// 辅助变量帮助添加节点
	cur := e.Head
	//辅助指针
	var pre *Emp
	if cur == nil {
		e.Head = emp
		return
	}
	// 如果不是空链表
	for cur !=nil{
		// 判断进行顺序插入
			if cur.Id > emp.Id{
				break
			}
			pre = cur
			cur = cur.Next
	}
	// 空指针,插入时只能顺序插入(从小到大插入,从大到小插入就不行 待解决!!!)
	//fmt.Println("pre : ->", pre)
	//fmt.Println("cur : ->", cur)
	// 表示代表插入的数要往两个数中间插
	if pre !=nil{
		pre.Next = emp
		emp.Next = cur
		e.lens++
		return
	}
	// 当只有一个数时,要插入到第一个位置,把要替换的位置往下挪
	e.Head = emp
	emp.Next = cur
	e.lens++
	return
}

func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	cur := e.Head
	fmt.Printf("*****************************\n\n")
	for cur != nil {
		fmt.Printf("链表%d 成员编号:%d\t 成员名字:%v -->", no, cur.Id, cur.Name)
		cur = cur.Next
	}
	fmt.Printf("\n\n*****************************")
	println()
}

func (e *EmpLink) remove(id int) {
	// 判断是否为空
	if e.Head == nil {
		fmt.Printf("删除失败,ID为 %d 为空\n", id)
		return
	}
	cur := e.Head
	// 前一个指针
	var cur2 *Emp
	for cur != nil {
		if cur.Id == id {
			// 等于说明当前链表只有一个
			if e.lens == 0 {
				e.Head = nil
				fmt.Println("唯一一个链表值被删除了", cur)
				return
			} else if cur.Next == nil {
				// 判断是否删除倒数最后一个
				//cur2.Next = nil
				cur2 = nil
				e.lens--
				fmt.Println("倒数最后一个删除成功!", cur)
				return
			}else if cur2==nil {
				// 当有多个值 并且要删除第一个时:
				fmt.Println("删除了第一个且他后面还有值:",cur.Next)
				e.Head = cur.Next
				cur = nil
				e.lens--
				return
			} else {
				// 当删除中间任意一个时
				cur.Next = cur.Next.Next
				e.lens--
				fmt.Println("cur2 :",cur2)
				fmt.Println("删除成功", cur)
				return
			}

		}
		cur2 = cur
		cur = cur.Next
	}
	fmt.Printf("链表ID为:%d 没有找到", id)
	return
}

func (e *EmpLink) len() {
	fmt.Println(e.lens + 1)
}

// 哈希表
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给哈希表添加一个数据
func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数确定将数据添加到那个链表
	LinkNo := h.HashFunc(emp.Id)
	// 向存放数据的结构体添加数据
	h.LinkArr[LinkNo].Insert(emp)
}

// 显示所有哈希值
func (h *HashTable) show() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

// 查询
func (h *HashTable) FindById(id int) *Emp {
	// 将id通过散列函数来找对应数据
	LinkNo := h.HashFunc(id)
	return h.LinkArr[LinkNo].Find(id)
}

// 删除
func (h *HashTable) remove(id int) {
	LinkId := h.HashFunc(id)
	h.LinkArr[LinkId].remove(id)
}
func (h *HashTable) len(id int) {
	LinkId := h.HashFunc(id)
	h.LinkArr[LinkId].len()
	return
}

// 根据id进行散列
func (h *HashTable) HashFunc(id int) int {
	return id % 7
}

func main() {
	var key string
	var err error
	var id int
	var name string
	hashtable := HashTable{}
	// 输出提示开关
	flag := false
	for {
		if flag == false {
			fmt.Println("************测试哈希table****************")
			fmt.Println("help h 输入h查看帮助")
			fmt.Println("input i 表示添加")
			fmt.Println("show s 表示查看")
			fmt.Println("find f 表示查找")
			fmt.Println("remove r 表示删除")
			fmt.Println("len l 查看长度")
			flag = true
		}
		fmt.Println("请输入你的选择!")
		_, err = fmt.Scanln(&key)
		if err != nil {
			goto exit
		}
		switch key {
		case "i":
			fmt.Println("请输入Id,Name")
			_, err = fmt.Scanln(&id,&name)
			// 创建一个哈希表需要的结构体并存放
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "s":
			hashtable.show()
		case "f":
			fmt.Println("请输入要查找的ID")
			_, err = fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp != nil {
				fmt.Println("----------当前查找内容为:-------------")
				fmt.Printf("ID:%v\t Name:%v\n", emp.Id, emp.Name)
				fmt.Println("----------------end------------------")
			} else {
				fmt.Printf("当前ID为%v的数据,没有找到\n", id)
			}
		case "r":
			fmt.Println("请输入删除的ID编号~")
			_, err = fmt.Scanln(&id)
			hashtable.remove(id)
		case "l":
			fmt.Println("请输入查看的链表ID编号~")
			_, err = fmt.Scanln(&id)
			hashtable.len(id)
		case "h":
			flag = false
		default:
			fmt.Println("输入有误,请重新输入!")
		}
	}
exit:
	fmt.Println(err)
	return
}
