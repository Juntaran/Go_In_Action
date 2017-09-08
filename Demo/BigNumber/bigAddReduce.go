/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/9/6 17:01 
  */

package BigNumber

import (
	"math"
	"strconv"
	"errors"
)

// 大数相加
func BigAdd(num1 string, num2 string) (string, error) {

	// 正负号分类，只需要写基本的 bigAdd 和 bigReduce 即可
	if num1[0] >= '0' && num1[0] <= '9' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigAdd(num1, num2)
		} else if num2[0] == '+' {
			return bigAdd(num1, num2[1:])
		} else if num2[0] == '-' {
			return bigReduce(num1, num2[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '+' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigAdd(num1[1:], num2)
		} else if num2[0] == '+' {
			return bigAdd(num1[1:], num2[1:])
		} else if num2[0] == '-' {
			return bigReduce(num1[1:], num2[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '-' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigReduce(num2, num1[1:])
		} else if num2[0] == '+' {
			return bigReduce(num2[1:], num1[1:])
		} else if num2[0] == '-' {
			ret, err :=  bigAdd(num1[1:], num2[1:])
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else {
			return "", errors.New("Error: Input")
		}
	} else {
		return "", errors.New("Error: Input")
	}
}

// 大数相减
func BigReduce(num1 string, num2 string) (string, error) {

	// 正负号分类，只需要写基本的 bigAdd 和 bigReduce 即可
	if num1[0] >= '0' && num1[0] <= '9' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigReduce(num1, num2)
		} else if num2[0] == '+' {
			return bigReduce(num1, num2[1:])
		} else if num2[0] == '-' {
			return bigAdd(num1, num2[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '+' {
		if num2[0] >= '0' && num2[0] <= '9' {
			return bigReduce(num1[1:], num2)
		} else if num2[0] == '+' {
			return bigReduce(num1[1:], num2[1:])
		} else if num2[0] == '-' {
			return bigAdd(num1[1:], num2[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else if num1[0] == '-' {
		if num2[0] >= '0' && num2[0] <= '9' {
			ret, err :=  bigAdd(num1[1:], num2)
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else if num2[0] == '+' {
			ret, err :=  bigAdd(num1[1:], num2[1:])
			if err != nil {
				return "", err
			} else {
				return "-" + ret, nil
			}
		} else if num2[0] == '-' {
			return bigReduce(num2[1:], num1[1:])
		} else {
			return "", errors.New("Error: Input")
		}
	} else {
		return "", errors.New("Error: Input")
	}
}

// 两个正整数相加
func bigAdd(num1, num2 string) (string, error) {
	len1 := len(num1)
	len2 := len(num2)
	lenMax := int(math.Max(float64(len1), float64(len2)))
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

	return result, nil
}

// 两个正整数相减，num1 - num2  如果 num1 < num2 会先算 num2 - num1 再加负号
func bigReduce(num1, num2 string) (string, error) {
	len1 := len(num1)
	len2 := len(num2)
	lenMax := int(math.Max(float64(len1), float64(len2)))
	var ret []int
	var result string
	
	// 判断 num1 和 num2 大小
	var judge int  = 0
	if len1 > len2 {
		judge = 1
	} else if len1 < len2 {
		judge = -1
	} else {
		// len1 == len2
		for i := 0; i < len1; i++ {
			// 从左向右逐位判断 num1 和 num2 谁大
			if num1[i] == num2[i] {
				continue
			} else if num1[i] > num2[i] {
				judge = 1
				break
			} else {
				judge = -1
				break
			}
		}
	}
	switch judge {
	case 0:			// num1 == num2
		return "0", nil
	case 1:			// num1 > num2
		// do nothing
		break
	case -1:		// num1 < num2
		num1, num2 = num2, num1
		len1, len2 = len2, len1
		result = "-"
	}

	var a = make([]uint8, lenMax+1)
	var b = make([]uint8, lenMax+1)
	var c = make([]int, lenMax+1)

	// 检验 num1 和 num2
	var str1 = make([]uint8, lenMax+1)
	var str2 = make([]uint8, lenMax+1)

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
		c[i] = int(a[i]) - int(b[i]) + c[i]
		if c[i] < 0 {
			c[i] += 10
			c[i+1] --
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
		result += strconv.Itoa(v)
	}

	return result, nil
}