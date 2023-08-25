package main

import (
	"flag"
	"fmt"

	"demo4/internal/config"
	"demo4/internal/handler"
	"demo4/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/fileservice.yaml", "the config file") //用于声明一个字符串类型的命令行参数。该函数接受三个参数：参数名称、默认值和参数描述。在这里，参数名称是 "f"，默认值是 "etc/fileservice.yaml"，参数描述是 "the config file"。

func main() {
	flag.Parse() //用于解析命令行参数并将结果赋值给相应的变量。在这里，该函数会解析命令行参数，并将结果存储在 configFile 变量所指向的内存地址中。通过这段代码，你可以在命令行中通过 -f 参数指定配置文件的路径，如果不指定，则会使用默认的配置文件路径 "etc/fileservice.yaml"。解析后的配置文件路径将存储在 configFile 变量中供后续使用

	var c config.Config
	conf.MustLoad(*configFile, &c) //加载配置文件到c

	server := rest.MustNewServer(c.RestConf) //用于创建一个 REST 服务器的实例。该函数接受一个参数 c.RestConf，表示 REST 服务器的配置参数。c.RestConf 可能是一个自定义的结构体，用于存储 REST 服务器的配置信息，例如监听地址、证书、路由配置等。
	//是将创建的 REST 服务器实例赋值给变量 server。通过这个变量，可以对服务器进行进一步的配置和操作。
	defer server.Stop()

	ctx := svc.NewServiceContext(c)       //创建一个新的服务上下文
	handler.RegisterHandlers(server, ctx) //用于注册处理程序到指定的服务器。该函数接受两个参数：server 表示服务器实例，ctx 表示服务上下文。可以将处理程序注册到服务器中，以便服务器可以处理相关的请求

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
