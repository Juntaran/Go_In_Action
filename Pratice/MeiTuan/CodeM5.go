/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/17 8:22
 */

/*

给定两个整数 l 和 r ，对于所有满足1 ≤ l ≤ x ≤ r ≤ 10^9 的 x ，
把 x 的所有约数全部写下来。对于每个写下来的数，只保留最高位的那个数码。
求1～9每个数码出现的次数。

输入描述:

	一行，两个整数 l 和 r (1 ≤ l ≤ r ≤ 10^9)。

输出描述:

	输出9行。
	第 i 行，输出数码 i 出现的次数。

输入例子:

	1 4

输出例子:

	4
	2
	1
	1
	0
	0
	0
	0
	0

 */

package main

import "fmt"

func getCount(x, l, r int) int {
	n := r - l + 1
	var count = n / x
	if count > 0 {
		m := n % x
		if (l + m - 1) >= x && (l + m - 1) % x == 0 {
			count ++
		}
	} else {
		for i := 0; l + i <= r; i++ {
			if l + i == x {
				count = 1
			}
		}

	}
	return count
}

func main() {
	var l, r int
	fmt.Scanf("%d %d", &l, &r)

	if l < 0 || r < l || r > (10 << 8) {
		return
	}
	// 输入结束

	count := 0
	n := r - l + 1

	// 第一行
	count = n
	fmt.Println(count)

	// 第二行
	count = n / 2
	if l % 2 == 0 {
		count ++
	}
	fmt.Println(count)

	// 第三行到第9行
	for i := 3; i <= 9; i++ {
		fmt.Println(getCount(i, l, r))
	}
}