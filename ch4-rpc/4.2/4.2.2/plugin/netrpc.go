package plugin

import (
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"log"
	"text/template"
)

//要使用该插件需要先通过generator.RegisterPlugin函数注册插件，可以在init函数中完成：

func init() {
   generator.RegisterPlugin(new(netrpcPlugin))
}



type netrpcPlugin struct{
	*generator.Generator
}

func (p *netrpcPlugin) Name() string{
	return "netrpc"
}

func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
		if len(file.Service) > 0 {
			 p.genImportCode(file)
		}
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
		for _, svc := range file.Service {
			 p.genServiceCode(svc)
		}
}
//生成导入代码
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	//p.P("// TODO: import code")
	p.P(`import "net/rpc"`)
}


//生成每个服务的代码
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	//p.P("// TODO: service code, Name = " + svc.GetName())
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService)) //其中tmplService是服务的模板。
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}


//然后我们新建一个buildServiceSpec方法用来解析每个服务的ServiceSpec元信息：
//todo 就是生成服务
//其中输入参数是`*descriptor.ServiceDescriptorProto`类型，完整描述了一个服务的所有信息。todo 就是proto文件中描述的那些啦
func (p *netrpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto, ) *ServiceSpec {
	//`svc.GetName()`就可以获取Protobuf文件中定义的服务的名字。
	//Protobuf文件中的名字转为Go语言的名字后，需要通过`generator.CamelCase`函数进行一次转换。
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	//在for循环中我们通过`m.GetName()`获取方法的名字，然后再转为Go语言中对应的名字。
	//todo 比较复杂的是对输入和输出参数名字的解析：首先需要通过`m.GetInputType()`获取输入参数的类型，
	//     然后通过`p.ObjectNamed`获取类型对应的类对象信息，最后获取类对象的名字。
	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}

	return spec
}


