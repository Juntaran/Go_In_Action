# Upload

## 基于HTTP的文件上传

    func (r *Request) ParseForm() error
    func (r *Request) ParseMultipartForm(maxMemory int64) error
    
`ParseForm`解析 URL 中的查询字符串，并将解析结果更新到`r.Form`字段  
`ParseMultipartForm`会自动调用`ParseForm`。重复调用本方法是无意义的  
`ParseMultipartForm`将请求的主体作为`multipart/form-data`解析  
请求的整个主体都会被解析，得到的文件记录最多`maxMemery`字节保存在内存，其余部分保存在硬盘的`temp`文件里  
如果必要，`ParseMultipartForm`会自行调用`ParseForm`，重复调用本方法是无意义的  

    func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)
    
`FormFile`返回以`key`为键查询`r.MultipartForm`字段得到结果中的第一个文件和它的信息  
如果必要，本函数会隐式调用`ParseMultipartForm`和`ParseForm`，查询失败会返回`ErrMissingFile`错误  

## Reference:
* [JieLinDee](http://blog.csdn.net/fyxichen/article/details/51165214)
