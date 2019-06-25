package main

import "fmt"

// 第一种

//func main()  {
//
//	test_slice := []int{1, 0, 4, 3, -2, -1, 10}
//	for i:=0;i<=len(test_slice)-1;i++ {
//		for j:=i;j<len(test_slice);j++{
//
//			if test_slice[i] < test_slice[j] {
//				test_slice[i], test_slice[j] = test_slice[j], test_slice[j]
//			}
//		}
//	}
//	fmt.Printf("第二大值: %d\n", test_slice[1])
//}


// 第二种

func main()  {
	max_num, second_num := 0, 0
	test_slice := []int{1, 0, 4, 3, -2, -1, 10}

	for i := range test_slice {
		if test_slice[i] > max_num {
			max_num = test_slice[i]
			second_num = max_num
		}else if test_slice[i] > second_num {
			second_num = test_slice[i]
		}

	}

	fmt.Printf("第二大值是： %d\n", second_num)
}
