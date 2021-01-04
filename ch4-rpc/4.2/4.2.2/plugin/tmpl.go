package plugin

/*
todo 在编写模板之前，我们先查看下我们期望生成的最终代码大概是什么样子：
当Protobuf的插件定制工作完成后，每次hello.proto文件中RPC服务的变化都可以自动生成代码。
也可以通过更新插件的模板，调整或增加生成代码的内容。
在掌握了定制Protobuf插件技术后，你将彻底拥有这个技术。

type HelloServiceInterface interface {
   Hello(in String, out *String) error
}

func RegisterHelloService(srv *rpc.Server, x HelloService) error {
   if err := srv.RegisterName("HelloService", x); err != nil {
      return err
   }
   return nil
}

type HelloServiceClient struct {
   *rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
   c, err := rpc.Dial(network, address)
   if err != nil {
      return nil, err
   }
   return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(in String, out *String) error {
   return p.Client.Call("HelloService.Hello", in, out)
}
*/


const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
   {{- range $_, $m := .MethodList}}
   {{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
   {{- end}}
}

func Register{{.ServiceName}}(
   srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
   if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
      return err
   }
   return nil
}

type {{.ServiceName}}Client struct {
   *rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
   *{{.ServiceName}}Client, error,
) {
   c, err := rpc.Dial(network, address)
   if err != nil {
      return nil, err
   }
   return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
   in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
   return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`




