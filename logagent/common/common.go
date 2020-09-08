package common

// CollectEntey 要收集的日志配置结构体
type CollectEntry struct {
	Path  string `json:"path"`  //去哪个路径读取日志文件
	Topic string `json:"topic"` //日志文件发往kafka哪个topic
}
