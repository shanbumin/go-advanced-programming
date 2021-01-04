package plugin

//为此我们定义了一个ServiceSpec类型，用于描述服务的元信息：
type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}