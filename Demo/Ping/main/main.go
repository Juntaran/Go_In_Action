/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/23 18:58
 */

package main

import (
	"Go_In_Action/Demo/Ping"
	"flag"
)

var (
	host = flag.String("host", "127.0.0.1", "Where are u want to Ping")
	count = flag.Int("n", 0, "The times of Ping")
)

func main() {
	flag.Parse()
	Ping.Ping(*host, *count)
}
