package testLee

import (
	"fmt"
	"math"
)

type Obj struct {
}

//测试代码有没有走到这里
func (this *Obj) Test0(str string) string {
	//fmt.Println(str)
	//log.Println("test0")
	//go fmt.Println("test0")
	return str + " is ok"
}

/*
题目twoSum
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

要点  该情况 是  a+b =c  则要先想到 c-a = b 这种方式来求解
*/
func (this *Obj) Test1(nums []int, target int) []int {
	result := make([]int, 2)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			result[1] = i
			result[0] = m[target-nums[i]]
			break
		}
		m[nums[i]] = i
	}
	return result
}

/*
	reverse(x int)
	给定一个 32 位有符号整数，将整数中的数字进行反转。

	要点
	1  注意溢出的情况
	2  负数也当正数处理 在返回的时候再把正负号+1
	3  用64位来计算32位的数
*/

func (this *Obj) Test2(x int) int {
	a := int64(x)
	rt := int64(0)
	pos := int64(1)
	if x < 0 {
		pos = -1
		a = -a
	}
	for a > 0 {
		n := a % 10
		a = a / 10
		rt = rt*10 + n
	}
	if rt > math.MaxInt32 {
		return 0
	}
	rt = pos * rt
	return int(rt)
}

/*
	判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
    10 以下的数字也是回文数
	要点为
	1 数字对称 所有 只要翻转后一半的数字来跟前一半的数字进行对比
	2 负数都不是回文数
	3 正数都不是以0 开头的  则可以进行尾数为0时 进行特殊判断
*/
func (this *Obj) Test3(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

/*
 罗马数字包含以下七种字符：I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，
 所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。
 这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

 要点  用特定设置反而比较快
*/
func (this *Obj) Test4(s string) int {
	var sum = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'V':
					sum += 4
					i++
					continue
				case 'X':
					sum += 9
					i++
					continue
				}
			}
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'L':
					sum += 40
					i++
					continue
				case 'C':
					sum += 90
					i++
					continue
				}
			}
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			if i+1 < len(s) {
				switch s[i+1] {
				case 'D':
					sum += 400
					i++
					continue
				case 'M':
					sum += 900
					i++
					continue
				}
			}
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		default:
			fmt.Println("Unknow char:", s[i])
		}
	}
	return sum
}

/*
  	编写一个函数来查找字符串数组中的最长公共前缀。
  ["flower","flow","flight"] 输出: "fl"
  输入: ["dog","racecar","car"] 输出: ""
Input: ["dog","racecar","car"]
Output: ""
*/
func (this *Obj) Test5(strs []string) string {
	result := ""
	if strs == nil || len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][0:i]
			}
		}
	}

	result = strs[0]
	return result
}

/*
   给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
   输入: "([)]"
输出: false
   输入: "()"
输出: true
   Input: "{[]}"
Output: true
   "()[]{}" true
    要点
   1 这种情况符合栈的特征后进先出
   2
*/
func (this *Obj) Test6(s string) bool {
	endTags := []string{"]", "}", ")"}
	beginTags := []string{"[", "{", "("}
	findInEend := func(str string) int {
		for index, endTag := range endTags {
			if str == endTag {
				return index
			}
		}
		return -1
	}

	findInBegin := func(str string) bool {
		findResult := false
		for _, beginTag := range beginTags {
			if str == beginTag {
				findResult = true
				break
			}
		}
		return findResult
	}
	finStr := make([]string, 0)
	for i := 0; i < len(s); i++ {
		curStr := string(s[i])
		if findInBegin(curStr) == true { //入栈
			finStr = append(finStr, curStr)
		} else { //出栈
			if len(finStr) == 0 { //无可出栈的元素
				return false
			} else if finStr[len(finStr)-1] != beginTags[findInEend(curStr)] { //出栈失败  栈顶的元素与该元素不符合
				a := beginTags[findInEend(curStr)]
				_ = a
				return false
			} else {
				finStr = finStr[:len(finStr)-1]
			}
		}
	}

	return len(finStr) == 0
}
