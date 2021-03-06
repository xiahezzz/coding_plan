**<font size=32>目录</font>**

# 算法入门

* `print_binary` 
    * 位运算，循环左移动。
* `cal_factorial` 
    * 计算N阶阶乘之和。
* `select_sort` 
    * 选择排序。
    * 时间复杂度$O(n^2)$
* `bubble_sort` 
    * 冒泡排序。
    * 时间复杂度$O(n^2)$
* `insert_sort` 
    * 插入排序。
* `prefix_sum_arr` 
    * 前缀和数组，为了解决多次频繁查找[L,R]区间内的数值和所设计的数值，在数组规模不大时，还可采用建立L-R矩阵的方式来做。
* `x_to_power2` 
    * math/rand伪随机，crypto/rand真随机
    * [0,1)范围内的随机数中，随机数落在 [0,x) 的概率为 x ，即[0,0.5)出现的概率为0.5
    * 将随机数落在 [0,x) 的概率调整为 $x^2$ 
    * 利用两次相同的独立实验 `max(rand() * rand())` ，即要得到随机数落在 [0,x) 则必须使得两次随机过程得到的随机数均落在 [0,x) 内，即 $x^2$
* `RandChange`
    * $f(x)$ 等概率返回 [1,5] ，构造$g(x)$ 等概率返回 [1,7]。第一步先构造0，1等概发生器$f'(x)$
    * `getRandNum01`
        * $f(x)$ 等概率返回 [1,5] ，将$f(x)$改为 $f'(x)$ 使得其等概率返回 0 和 1
    * 利用$f'(x)$构造三个比特位的二进制，即000～111，遇到000则重新进入循环。每一个比特为0和1由0，1等概发生器提供
* `binary_search`
    * `BinarySearch`
        * 二分查找
        * 时间复杂度$O(logN)$
    * `FindNumBLeft`
        * 利用二分查找，找到大于等于Num最左边的数
    * `FindMinIndex`
        * 利用二分查找，找到局部最小值
        * 数组可以是无序的，但相邻位置必须不相同

# Go语言入门
## `user/trick.go` 

* `ConcatString` Go里面string是最基础的类型，是一个只读类型，针对他的每一个操作都会创建一个新的string。 如果是少量小文本拼接，用 “+” 就好；如果是大量小文本拼接，用 strings.Join；如果是大量大文本拼接，用 bytes.Buffer。
* `RandTest` 随机数生成，Go中不能对不同类型的数值进行任意科学运算！在主函数中先设置随机数种子。
* 注意引用传递时，实参被改变的情况。
  
## `user/utils.go` 
* `nq.PeekQueue()` 注意队列为空时，直接返回nq.Head为nil，但不可调用nq.Head(nil).Data(引用空指针)