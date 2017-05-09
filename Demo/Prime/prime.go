/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/9 20:07
  */

package Prime

//import "fmt"

var ret []int

func processor(seq chan int, wait chan struct{}, ret *[]int)  {

	go func() {
		prime, ok := <-seq
		*ret = append(*ret, prime)
		if !ok {
			close(wait)
			return
		}
		//fmt.Println(prime)

		out := make(chan int)
		processor(out, wait, ret)
		for num := range seq {
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
	} ()

}

// 获取小于number的所有质数并返回一个slice以及质数个数
func GetPrime(number int) ([]int, int) {
	origin, wait := make(chan int), make(chan struct{})
	processor(origin, wait, &ret)
	for num := 2; num < number; num++ {
		origin <- num
	}
	close(origin)
	<- wait

	ret = append(ret[0:len(ret)-1])
	return ret, len(ret)
}

