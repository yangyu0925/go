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
