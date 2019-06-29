

nums1 := []int{1, 3, 9, 8, 7}

对代码块起一个名字 => 函数名
定义一个函数
参数
返回值

for j := 0; j < len(nums1)-1; j++ {
    for i := 0; i < len(nums1)-1-j; i++ {
        if nums1[i] > nums1[i+1] {
            tmp := nums1[i]
            nums1[i] = nums1[i+1]
            nums1[i+1] = tmp
        }
    }
}


nums2 := []int{2, 3, 9, 8, 7}

for j := 0; j < len(nums2)-1; j++ {
    for i := 0; i < len(nums2)-1-j; i++ {
        if nums2[i] > nums2[i+1] {
            tmp := nums2[i]
            nums2[i] = nums2[i+1]
            nums2[i+1] = tmp
        }
    }
}

f(n) 1...n

0 - 100 f(100)
0 - 99 + 100  f(99) + 100
0 - 98 + 99 f(98) + 99

f(n) = n + f(n-1)
f(1) = 1


n阶乘:
n! = 1 * 2 * 3 * 4 * 5 * ... * n
n = 0 => n! = 1


命令行的用户管理

用户信息存储
    => 内存
    => 结构 map => []
    => 用户 id name age tel addr
            map
            值类型使用string

用户添加
用户的查询

用户修改
    // 请输入需要修改的用户ID:
    // users[id] => 在 不在(你输入的用户ID不正确)
    // 打印用户信息, 提示用户是否确认修改(Y/N)
    // Y 提示用户输入修改后内容
    // name, age ,tel, addr

用户删除
    // 请输入需要删除的用户ID:
    // users[id] => 在 不在(你输入的用户ID不正确)
    // 打印用户信息, 提示用户是否确认删除(Y/N)
    // Y delete()