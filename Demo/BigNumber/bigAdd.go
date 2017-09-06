/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/6 17:01 
  */

package BigNumber

import (
	"math"
	"fmt"
	"strconv"
	"errors"
)

// 大数相加
func BigAdd(num1 string, num2 string) (string, error) {

	len1 := len(num1)
	len2 := len(num2)
	lenMax := int(math.Max(float64(len1), float64(len2)))
	var ret []uint8
	var result string

	var a = make([]uint8, lenMax + 1)
	var b = make([]uint8, lenMax + 1)
	var c = make([]uint8, lenMax + 1)

	// 检验 num1 和 num2
	var temp1 = make([]uint8, lenMax + 1)
	var temp2 = make([]uint8, lenMax + 1)

	if (temp1[0] > 9 || temp1[0] < 0) && (temp1[0] != '-' || temp1[0] != '+') {
		return "", errors.New("Error: Input")
	}
	if (temp2[0] > 9 || temp2[0] < 0) && (temp2[0] != '-' || temp2[0] != '+') {
		return "", errors.New("Error: Input")
	}

	for i := 0; i < len1; i++ {
		temp1[i] = num1[i]
		if i >= 1 {
			if temp1[i] > '9' || temp1[i] < '0' {
				return "", errors.New("Error: Input")
			}
		}
	}
	for j := 0; j < len2; j++ {
		temp2[j] = num2[j]
		if j >= 1 {
			if temp2[j] > '9' || temp2[j] < '0' {
				return "", errors.New("Error: Input")
			}
		}
	}

	var str1 = make([]uint8, lenMax + 1)
	var str2 = make([]uint8, lenMax + 1)

	// 正负判断
	if temp1[0] >= '0' && temp1[0] <= '9' {
		str1 = temp1
	} else {
		str1 = temp1[1:]
		len1 --
		lenMax = int(math.Max(float64(len1), float64(len2)))
	}
	if temp2[0] >= '0' && temp2[0] <= '9' {
		str2 = temp2
	} else {
		str2 = temp2[1:]
		len2 --
		lenMax = int(math.Max(float64(len1), float64(len2)))
	}

	symbol := ""

	if temp1[0] == '+' && temp2[0] == '+' {

	}
	if temp1[0] == '-' && temp2[0] == '-' {
		symbol = "-"
	}
	if (temp1[0] == '+' && temp2[0] == '-') || (temp1[0] >= '0' && temp1[0] <= '9' && temp2[0] == '-') {
		return BigSubtract(num1, num2[1:])
	}
	if (temp1[0] == '-' && temp2[0] == '+') || (temp2[0] >= '0' && temp2[0] <= '9' && temp1[0] == '-') {
		return BigSubtract(num2, num1[1:])
	}


	for i := 0; i < lenMax; i++ {
		if len1 - i - 1 < 0 {
			b[len2-i-1] = str2[i] - '0'
			continue
		}
		if len2 - i - 1 < 0 {
			a[len1-i-1] = str1[i] - '0'
			continue
		}
		a[len1-i-1] = str1[i] - '0'
		b[len2-i-1] = str2[i] - '0'
	}

	for i := 0; i < lenMax; i++ {
		c[i] = a[i] + b[i]
	}

	for i := 0; i < lenMax; i++ {
		c[i+1] = c[i] / 10 + c[i+1]
		c[i] = c[i] % 10
	}
	if c[lenMax] != 0 {
		for i := lenMax; i >= 0; i-- {
			ret = append(ret, c[i])
		}
	} else {
		for i := lenMax - 1; i >= 0; i-- {
			ret = append(ret, c[i])
		}
	}

	for _, v := range ret {
		result += strconv.Itoa(int(v))
	}
	result = symbol + result
	fmt.Println(result)

	return result, nil
}

