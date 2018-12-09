/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/23 18:20
 */

package main

import (
	"net/http"
	"github.com/Juntaran/Go_In_Action/Demo/Upload"
)

func main() {
	http.HandleFunc("/", Upload.Index)
	http.HandleFunc("/upload", Upload.Upload)
	http.ListenAndServe(":1789", nil)
}
