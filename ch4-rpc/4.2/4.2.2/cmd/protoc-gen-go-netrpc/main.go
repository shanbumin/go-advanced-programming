package main

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)
//因为Go语言的包只能静态导入，我们无法向已经安装的protoc-gen-go添加我们新编写的插件。我们将重新克隆protoc-gen-go对应的main函数：
//为了避免对protoc-gen-go插件造成干扰，我们将我们的可执行程序命名为protoc-gen-go-netrpc，表示包含了netrpc插件。
func main() {
	g := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}


//然后用以下命令重新编译hello.proto文件：
//protoc --go-netrpc_out=plugins=netrpc:. hello.proto
//其中`--go-netrpc_out`参数告知protoc编译器加载名为protoc-gen-go-netrpc的插件，插件中的`plugins=netrpc`指示启用内部唯一的名为netrpc的netrpcPlugin插件。
//在新生成的hello.pb.go文件中将包含增加的注释代码。
//至此，手工定制的Protobuf代码生成插件终于可以工作了。
