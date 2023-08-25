package logic

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"demo4/internal/svc"
	"demo4/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 //10MB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(), //ctx 是请求的上下文，包含了请求相关的上下文信息
		r:      r,
		svcCtx: svcCtx,//svcCtx 表示服务上下文（service context），它是一个自定义的上下文对象，用于传递服务相关的上下文信息。在这里，*svc.ServiceContext 是一个自定义的类型，表示包含了服务所需的所有依赖和配置的上下文对象。通过将 svcCtx 作为参数传递给逻辑处理器，可以使逻辑处理器能够访问和使用这些服务所需的上下文信息，比如数据库连接、缓存实例等。
	}
}

func (l *UploadLogic) Upload() (resp *types.UploadResponse, err error) {
	// todo: add your logic here and delete this line
	l.r.ParseMultipartForm(maxFileSize) //将 HTTP 请求中的多部分表单数据解析为可用的表单值。它会解析请求体中的表单字段，包括普通字段和文件字段，并将它们存储在 r.MultipartForm 字段中，供后续代码使用,同时限制上传的大小
	file, handler, err := l.r.FormFile("myFile")
	//这行代码的作用是从 HTTP 请求的表单数据中获取名为 “myFile” 的文件字段。它返回三个值：
	//file 是一个实现了 io.Reader 和 io.Closer 接口的对象，可以通过读取该对象来获取上传文件的内容。
	//handler 是一个包含了上传文件的元数据信息的对象，包括文件名、文件大小等。
	//err 表示从表单中获取文件的过程中是否发生了错误，如果没有错误，err 的值为 nil。
	if err != nil {
		return nil, err

	}
	defer file.Close()
	logx.Infof("upload file: %+v, file size: %d,MIME header: %+v",
		handler.Filename, handler.Size, handler.Header) //%+v用于将结构体类型的变量的字段名和字段值一并输出。
	tempfile, err := os.Create(path.Join(l.svcCtx.Config.Path, handler.Filename))
	//用于创建一个文件并返回文件对象。
	//os.Create() 是一个用于创建文件的函数，它接受一个文件路径作为参数，并返回创建的文件对象和一个错误对象。
	//path.Join(l.svcCtx.Config.Path, handler.Filename) 是一个用于拼接文件路径的函数，它将 l.svcCtx.Config.Path 和 handler.Filename 这两个路径合并为一个完整的文件路径。
	//这行代码的作用是在指定的路径下创建一个文件，文件名由 handler.Filename 决定。它返回两个值:
	//tempfile 是一个 *os.File 指针类型的变量，代表创建的文件对象。你可以通过对 tempfile 进行读取、写入和其他操作来操作这个文件。
	if err != nil {
		return nil, err
	}
	defer tempfile.Close()
	io.Copy(tempfile, file)
	return &types.UploadResponse{
		Code: 0,
	}, nil
}
