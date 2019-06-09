# go

## 定义变量

* 使用var关键字 
    * var a, b, c bool
    * var s1, s2 string = "hello", "world" 
    * 可放在函数内，或直接放在包内
    * 使用var (aa = 33)集中定义变量
* 让编译器自动决定类型
    * var a, b, i, s1, s2 = true, false, 3, "hello", "world"
* 使用:=定义变量
    * a, b, i, s1, s2 := true, false, 3, "hello", "world"  
    * 只能在函数内使用
    
## 内建变量类型
* bool, string
* (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
* byte, rune
* float32, float64, complex64, complex128
* 强制类型转换
    * 类型转换是强制的
    * var a, b int = 3, 4
    * var c int = int(math.Sqrt(float64(a * a + b * b)))
    
## 常量的定义
* const filename = "abc.txt"
* const 数值可作为各种类型使用
* const a, b = 3, 4
* var c int = int(math.Sqrt(a * a + b * b))

## 变量定义要点回顾
* 变量类型写在变量名之后
* 编译器可推测变量类型
* 没有char，只有rune
* 原生支持复数类型

## 条件判断语句
```
if contents, err := ioutil.ReadFile(filename); err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%s \n", contents)
}
```
* if的条件里可以赋值
* if的条件里赋值的变量作用域就在这个if语句里

## 循环
```
    sun := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
```
* for的条件里不需要括号
* for的条件里可以省略初始条件，结束条件，递增表达式

```
    func convertToBin(n int) string {
    	result := ""
    	for ; n > 0 ; n /= 2 {
    		lsb := n % 2
    		result = strconv.Itoa(lsb) + result
    	}
    	return  result
    }
    
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
    for {
        fmt.Println("abc")
    }
```
* 省略初始条件，相当于while

## 基本语法要点回顾
* for，if后面的条件没有括号
* if条件里也可以定义变量
* 没有while
* switch不需要break，也可以直接switch多个条件

## 函数
```
    func eval(a, b int, op string) int
    
    //函数可返回多个值
    func div(a, b int) (int, int) {
        return a / b, a % b
    }
    
    //函数返回多个值时可以起名字
    //仅用于非常简单的函数
    //对于调用者而言没有区别
    func div(a, b int) (q, r int) {
        q = a / b
        r = a % b
        return
    }
    
    //函数作为参数
    func apple(op func(int, int) int, a, b int) int {
        fmt.Printf("Calling %s whth %d %d",
            runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
        a, b)
        return op(a, b)
    }
    
    //可变参数列表
    func sum(numbers ...int) int{
        s := 0
        for i := range numbers {
            s += numbers[i]
        }
        return s
    }
```
## 函数语法要点回顾
* 返回值类型写在最后面
* 可返回多个值
* 函数作为参数
* 没有默认参数、可选参数

## 指针
```
    //指针不能运算
    var a int = 2
    var pa *int = &a
    *pa = 3
    fmt.Println(a)
    
    //参数传递
    func swap(a, b *int) {
        *b, *a = *a, *b
    }
```

##数组
```
    //数量写在类型前
    var arr1 [5]int
    arr2 := [3]int{1, 3, 5}
    arr3 := [...]int{2, 4, 6, 8, 10}
    var grid [4][5]int
    
    for i := 0; i < len(arr3); i++ {
        fmt.Println(arr3[i])
    }
    
    for i, v := range arr3{
        fmt.Println(i, v)
    }
    
    maxi := -1
    maxValue := -1
    for i, v := range numbers {
        if v > maxValue {
            maxi, maxValue = i, v
        }
    }
    
    //可通过_省略变量
    //不仅range，任何地方都可以通过_省略变量
    //如果只要i，可写成for i := range numbers
    sum := 0
    for _, v := range numbers {
        sum += v
    }
```
* 为什么要用range
    * 意义明确，美观
    * C++没有类似能力
    * Java/Python:只能for each value，不能同时获取i, v
* 数组是值类型
    * [10]int和[20]int是不同类型
    * 调用func f(arr [10]int)会拷贝数组
    * 在go语言中一般不直接使用数组
    
## Slice(切片)
```
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    //s的值为[2 3 4 5]
    s := arr[2:6]
    
    //s1的值为[2 3 4 5], s2的值为[5 6]
    //slice可以向后拓展，不可以向前拓展
    //s[i]不可以超越len(s),向后拓展不可以超越底层数组cap(s)
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s1 := arr[2:6]
    s2 := s1[3:5]
    
    arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    s1 := arr[2:6]
    s2 := s1[3:5]
    s3 := append(s2, 10)
    s4 := append(s3, 11)
    s5 := append(s4, 12)
    //s3, s4, s5 = [5 6 10] [5 6 10 11] [5 6 10 11 12]
    fmt.Println("s3, s4, s5 =", s3, s4, s5)
    //arr = [0 1 2 3 4 5 6 10]
    fmt.Println("arr =", arr)
    
    s1 := []int{2, 4, 6, 8}
    //[2 4 6 8], len=4, cap=4
    printSlice(s1)

    s2 := make([]int, 16)
    //[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
    printSlice(s2)

    s3 := make([]int, 10, 32)
    //[0 0 0 0 0 0 0 0 0 0], len=10, cap=31
    printSlice(s3)


	fmt.Println("Copying slice")
	copy(s2, s1)
	//[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16, cap=16
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	//[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15, cap=16
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	//2
	fmt.Println(front)


	fmt.Println("Popping from back")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	//0
	fmt.Println(tail)

    //[4 6 0 0 0 0 0 0 0 0 0 0 0], len=13, cap=15
	printSlice(s2)
```
* 向Slice添加元素
    * 添加元素时如果超越cap，系统会重新分配更大的底层数组
    * 由于值传递的关系，必须接受append的返回值
    * s = append(s, val)  
    
## map
```
    m := map[string]string {
        "name": "ccmouse",
        "course": "golang",
        "site": "imooc",
        "quality": "notbad",
    }
    
    //m2 == empty map
    m2 := make(map[string]int)

    //m3 == nil
    var m3 map[string]int
    
	fmt.Println("Traversing map")
	for k, v := range m{
		fmt.Println(k,":",v)
	}

	fmt.Println("Getting value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if coursName, ok := m["cours"]; ok{
		fmt.Println(coursName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	delete(m, "name")
	fmt.Println(name, ok)

	name, ok = m["name"]
	fmt.Println(name, ok)
```
* Map的操作
    * 创建：make(map[string]int)
    * 获取元素：m[key]
    * key不存在时，获取value类型的初始值
    * 用value，ok := m[key]来判断是否存在key
    * 用delete删除一个key
* map的遍历
    * 使用range遍历key，或者遍历key，value对
    * 不保证遍历顺序，如需要顺序，需手动对key排序
    * 使用len获取元素个数
* map的key
    * map使用哈希表，必须可以比较相等
    * 除了slice，map，function的内建类型都可以作为key
    * Struct类型不包含上述字段，也可作为key
* rune相等与go的char
    * 使用range遍历pos, rune对
    * 使用utf8.RuneCountInString获得字符数量
    * 使用len获得字节长度
    * 使用[]byte获得字节
* 其他字符串操作
    * fields, Split, Join
    * Contains, Index
    * ToLower, ToUpper
    * Trim, TrimRight, TrimLeft
    
## 面向对象
* go语言仅支持封装，不支持继承和多态
* go语言没有class，只有struct
```
    type tree struct {
        Left, Right *TreeNode
        Value int
    }
    
    //结构的创建
    var root tree
    
    root = tree{value: 3}
    root.left = &tree{}
    root.right = &tree{5, nil, nil}
    root.right.left = new(tree)
    
    func createTree(value int) *tree {
    	return &tree{value: value}
    }
    //使用自定义的工厂函数
    //注意返回了局部变量的地址！
	root.left.right = createTree(2)

    //为结构定义方法
    //显示定义和命名方法接收者
    func (node tree) print() {
        fmt.Print(node.Value)
    }
    
    //只有使用指针才可以改变结构内容
    //nil指针也可以调用方法
    func (node *tree) setValue(value int) {
        node.Value = value
    }
```
* 值接收者vs指针接收者
    * 要改变内容必须使用指针接收者
    * 结构过大也考虑使用指针接收者
    * 一致性：如果指针接收者，最好都是指针接收者
    * 值接收者是go语言特有
    * 值/指针接收者均可接收值/指针
    
* 封装
    * 名字一般使用CamelCase
    * 首字母大写: public
    * 首字母小写: private
    
* 包
    * 每个目录一个包
    * main包包含可执行函数
    * 为结构定义的方法必须放在同一个包内
    * 可以是不同的文件
    * 扩充系统类型或别人的类型
    * 定义别名
    * 使用组合