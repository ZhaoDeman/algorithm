package main

import "fmt"

func Combine(n int, k int) [][]int {
	ans = make([][]int, 0)
	combine(1, n, k, []int{})
	fmt.Println("ans:", ans)
	return ans
}

var ans [][]int

func combine(cur, n int, k int, item []int) {
	if (n-cur+1)+len(item) < k {
		return
	}
	if k == len(item) {
		dst := make([]int, k)
		copy(dst, item)
		fmt.Println("结果：", dst)
		ans = append(ans, dst)
		return
	}
	item = append(item, cur)
	fmt.Println("递归前：", item)
	combine(cur+1, n, k, item)
	item = item[:len(item)-1]
	fmt.Println("递归后：", item)
	combine(cur+1, n, k, item)
}

func LetterCombinations(digits string) []string {
	str := ""
	for i := 0;i<len(digits);i++ {
		str += phoneMap[string(digits[i])]
	}
	fmt.Println("str-->",str)
	dfs2(len(str),len(digits),0,str,"")
	fmt.Println(numsAns)
	return numsAns
}

var numsAns []string
var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqr",
	"8": "stu",
	"9": "wxyz",
}

func dfs2(m,n,index int,str string,item string) {
	/*if (m - index + 1) + len(item) < n {
		return
	}*/
	if len(item) == n {
		newStr := item
		fmt.Println("结果-->",newStr)
		numsAns = append(numsAns, newStr)
		return
	}
	item += string(str[index])
	fmt.Println("递归前-->",item)
	dfs2(m, n, index+1, str, item)
	item = item[:len(item)-1]
	fmt.Println("递归后-->",item)
	dfs2(m , n , index +1, str , item )
}