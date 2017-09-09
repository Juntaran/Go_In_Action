/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/9/9 16:40
 */

package BigNumber

import (
	"errors"
	"strconv"
)

// 大数相乘
func BigMulti(num1, num2 string) (string, error) {
	
	// 正负号分类
	if num1[0] >= '0' && num1[0] <= '9' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigMulti(num1, num2)
		} else if num2[0] == '+' {
			return bigMulti(num1, num2[1:])
		} else if num2[0] == '-' {
			ret, err := bigMulti(num1, num2[1:])
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '+' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigMulti(num1[1:], num2)
		} else if num2[0] == '+' {
			return bigMulti(num1[1:], num2[1:])
		} else if num2[0] == '-' {
			ret, err := bigMulti(num1[1:], num2[1:])
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '-' {
		if num2[0] >= '0' && num2[0] <= '9' {
			ret, err := bigMulti(num1[1:], num2)
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else if num2[0] == '+' {
			ret, err := bigMulti(num1[1:], num2[1:])
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else if num2[0] == '-' {
			return bigMulti(num1[1:], num2[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else {
		return "", errors.New("Error: Input")
	}
}

// 两个正整数相乘
func bigMulti(num1, num2 string) (string, error) {
	len1 := len(num1)
	len2 := len(num2)
	lenMax := len1 + len2
	var ret []uint8
	var result string

	var a = make([]uint8, lenMax + 1)
	var b = make([]uint8, lenMax + 1)
	var c = make([]uint8, lenMax + 1)

	// 检验 num1 和 num2
	var str1 = make([]uint8, lenMax + 1)
	var str2 = make([]uint8, lenMax + 1)

	for i := 0; i < len1; i++ {
		str1[i] = num1[i]
		if str1[i] > '9' || str1[i] < '0' {
			return "", errors.New("Error: Input")
		}
	}
	for j := 0; j < len2; j++ {
		str2[j] = num2[j]
		if str2[j] > '9' || str2[j] < '0' {
			return "", errors.New("Error: Input")
		}
	}

	for i := 0; i < len1; i++ {
		a[len1-i-1] = str1[i] - '0'
	}
	for i := 0; i < len2; i++ {
		b[len2-i-1] = str2[i] - '0'
	}

	for i := 0; i < len2; i++ {
		for j := 0; j < len1; j++ {
			c[i+j] += a[j] * b[i]
		}
	}
	for i := 0; i < lenMax; i++ {
		if c[i] >= 10 {
			c[i+1] = c[i+1] + c[i]/10
			c[i] = c[i] % 10
		}
	}
	var mark int
	for i := lenMax - 1; i >= 0; i-- {
		if c[i] != 0 {
			mark = i
			break
		}
	}
	for j := mark; j >= 0; j-- {
		ret = append(ret, c[j])
	}
	for _, v := range ret {
		result += strconv.Itoa(int(v))
	}

	return result, nil
}