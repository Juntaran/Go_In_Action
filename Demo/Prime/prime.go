/**
  * Author: Juntaran
  * Email:  Jacinthmail@gmail.com
  * Date:   2017/5/9 20:07
  */

package Prime

import "fmt"

func Processor(seq chan int, wait chan struct{})  {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
	} ()
}
