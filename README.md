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