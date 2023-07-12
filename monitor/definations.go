package monitor

// 主要Prometheus的命名要求: a-zA-Z0-9 符号只能为包含 _ : 两种

const (
	Namespace = "monitor"
)

const (
	LabelHandler     = "handler"      // 请求路径
	LabelHTTPStatus  = "status_code"  // HTTP状态值
	LabelOrigin      = "origin"       // 访问请求源, 枚举值, 值: API, browser, Unknown
	LabelProcessName = "process_name" // 进程名称: 配置文件中获取
	LabelProcessAddr = "process_addr" // 进程地址: 格式 ip:端口号, 配置文件中获取
	LabelAppCode     = "app_code"     // 请求应用代码
)

// 通过pod定义的valueFrom参数引入. key对应的值在helm中pod声明处, 需要一一对应
const (
	ENVHostIP = "APP_HOST_IP"
	ENVPodIP  = "APP_POD_IP"
)
