/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/6/23 18:19
 */

package Upload

import (
	"fmt"
	"os"
	"io"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	// 获取文件名
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	fmt.Println(handler.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打开文件
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 写文件
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{...{.}...}"/>
	<input type="submit" value="upload" />
</form>
</body>
</html>`
