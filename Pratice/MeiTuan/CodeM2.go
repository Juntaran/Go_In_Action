/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/17 6:37
 */

/*

组委会正在为美团点评CodeM大赛的决赛设计新赛制。
比赛有 n 个人参加（其中 n 为2的幂），
每个参赛者根据资格赛和预赛、复赛的成绩，会有不同的积分。
比赛采取锦标赛赛制，分轮次进行，设某一轮有 m 个人参加，
那么参赛者会被分为 m/2 组，每组恰好 2 人，m/2 组的人分别厮杀。
我们假定积分高的人肯定获胜，若积分一样，则随机产生获胜者。
获胜者获得参加下一轮的资格，输的人被淘汰。重复这个过程，直至决出冠军。
现在请问，参赛者小美最多可以活到第几轮（初始为第0轮）？

输入描述:

	第一行一个整数 n (1≤n≤ 2^20)，表示参加比赛的总人数。
	接下来 n 个数字（数字范围：-1000000…1000000），表示每个参赛者的积分。
	小美是第一个参赛者。

输出描述:

	小美最多参赛的轮次。

输入例子:

	4
	4 1 2 3

输出例子:

	2

 */

/*

模拟一下比赛过程即可：

尽量让小美和比她分数低（包括相同，因为相同时是随机比赛结果，且要求能通过的最大比赛场次）的选手去比赛，
那么首先想到排序，再求出小美分数在排序数组中的上界，计算出比小美分数高的选手数量r，然后就是模拟比赛的过程。

将选手划分为两个阵营（l：分数<=小美的选手，包括小美; r: 分数>小美的选手）。

规则是：在一轮比赛中，如果r为奇数，需要在l个中，抽一个（比小美分数小的）给r才能比赛，

　　　　　　如果l==1，即只剩小美一个人了，那这局肯定是输的，不计算在内，这时可能r中可能还有选手，但小美只能到这了。
　　　　　　如果l>1，那这一轮比赛分别在l-1个选手， 与r+1个选手， 两个阵营内进行。小美在l阵营中，肯定会赢。胜利次数+1。

在一轮比赛中，如果r 为0，小美一直比下去。

*/

package main

import (
	"fmt"
	"sort"
)

// 二分查找首先要数组有序
func Binary_Search(nums []int, key int) (int, int, int) {
	left, right := 0, len(nums)-1
	same := 0
	for left <= right {
		middle := (right - left) >> 1 + left
		if key < nums[middle] {
			right = middle - 1
		} else if key > nums[middle] {
			left = middle + 1
		} else {
			same = 1
			var lsame, rsame = 0, 0
			for i := 1; i+middle < len(nums); i++ {
				if nums[i+middle] == key {
					rsame ++
				}
			}
			for i := 1; middle-i >= 0; i++ {
				if nums[middle-i] == key {
					lsame ++
				}
			}
			same = same + lsame + rsame
			return middle - lsame, middle - lsame + same - 1, same	// 返回数组左右下标和key的个数
		}
	}
	return -1, -1, 0
}

func getSum(m int) int {
	sum := 1
	for i := 0; i < m; i++ {
		sum *= 2
	}
	return sum
}

func main() {
	var n int		// 参加比赛的总人数
	fmt.Scanf("%d", &n)
	if n < 1 || n > (2 << 19) {
		return
	}
	var points = make([]int, n)		// 每个参赛者的积分
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &points[i])
	}
	// 输入结束

	var count = 0		// 小美的轮次
	k := points[0]		// 小美的分数

	sort.Ints(points)	// 排序分数

	// 让小美和分数低的比
	left, right, same := Binary_Search(points, k)		// 获取同分个数以及左右下标
	// left 即为低于小美分数的个数
	// n - right - 1 即为高于小美分数的个数
	// left - 1 为低于小美分数的最右下标
	// right + 1 为高于小美分数的最左下标

	//fmt.Println(left, right, same)

	l := left + same
	r := n - right - 1

	if r == 0 {
		count = l / 2
	}

	for r > 0 && l > 1 {
		if r & 1 != 0 {
			r ++
			l --
		}
		count ++
		r /= 2
		l /= 2
	}

	fmt.Printf("%d", count)
	return
}

/*

package main

import (
"fmt"
"sort"
)

func getSum(m int) int {
	sum := 1
	for i := 0; i < m; i++ {
		sum *= 2
	}
	return sum
}

func main() {
	var n int		// 参加比赛的总人数
	fmt.Scanf("%d", &n)
	if n < 1 || n > (2 << 19) {
		return
	}
	var points = make([]int, n+1)		// 每个参赛者的积分
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &points[i])
	}
	// 输入结束

	var count = 0		// 小美的轮次
	k := points[0]		// 小美的分数

	sort.Ints(points)	// 排序分数

	// 让小美和分数低的比
	for i := 1; ; i++ {

		sum := getSum(i)
		if sum > n {
			break
		}
		if k >= points[sum] {
			count ++
		} else {
			break
		}
	}
	fmt.Printf("%d", count)
	return
}
*/