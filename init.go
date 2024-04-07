package otely

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"io"
)

func init() {

}

// 配置采集源信息
func NewResource() *resource.Resource {
	defaultR, _ := resource.New(context.Background(),
		resource.WithFromEnv(),   // pull attributes from OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables
		resource.WithProcess(),   // 流程信息 This option configures a set of Detectors that discover process information
		resource.WithOS(),        // 操作系统信息 This option configures a set of Detectors that discover OS information
		resource.WithContainer(), // 容器信息 This option configures a set of Detectors that discover container information
		resource.WithHost(),      // 主机信息 This option configures a set of Detectors that discover host information
	)
	// 合并 资源对象属性
	r, _ := resource.Merge(
		defaultR,
		// 更具提供的属性创建资源
		resource.NewWithAttributes(
			semconv.SchemaURL,                       // 版本模式匹配
			semconv.ServiceName("MY FIB"),           // 资源服务名称
			semconv.ServiceVersion("v0.0.1"),        // 资源版本
			attribute.String("environment", "demo"), // 其他属性信息
		),
	)
	return r
}

// 创建导出器
func NewExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),     // 配置写入器 实现 io.Writer 接口
		stdouttrace.WithPrettyPrint(), // 格式化输出内容
	)
}
