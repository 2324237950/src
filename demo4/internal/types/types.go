// Code generated by goctl. DO NOT EDIT.
package types

type UploadResponse struct {
	Code int `json:"code"`
}

type DownloadRequest struct {
	File string `path:"file"`
	//path:"file" 是一个结构体标签（struct tag），
	// 用于为 File 字段添加元数据信息。
	// 在这里，path:"file" 表示将 File 字段与路径字符串 “file” 关联起来，
	// 可能用于标识该字段在 HTTP 请求中的位置或名称。
	//不知道会不会赋值
}