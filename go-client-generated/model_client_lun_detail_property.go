/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 客户端卷设备详情的属性信息
type ClientLunDetailProperty struct {
	// 卷是否开启数据保护
	DataProtection bool `json:"data_protection,omitempty"`
	// 资源的性能健康状态非健康的原因
	PerfHealthStatusReason string `json:"perf_health_status_reason"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 卷所在存储池名称
	PoolName string `json:"pool_name"`
	// 卷的使用状态
	UseStatus string `json:"use_status"`
	// 资源的创建方式
	CreateFrom string `json:"create_from"`
	// 资源的容量健康状态
	CapHealthStatus string `json:"cap_health_status"`
	// 卷是否为只读
	IsReadonly bool `json:"is_readonly"`
	// 卷是否开启校验
	VerifyEnabled bool `json:"verify_enabled,omitempty"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 资源的同步状态
	SyncStatus string `json:"sync_status"`
	// 资源的性能健康状态
	PerfHealthStatus string `json:"perf_health_status"`
	// 资源的容量健康状态非健康的原因
	CapHealthStatusReason string `json:"cap_health_status_reason"`
	// 卷的数据量
	DataSize int64 `json:"data_size"`
	// 卷是否设置QOS
	QosStatus bool `json:"qos_status,omitempty"`
	// 卷所属平台
	OwnerPlatform string `json:"owner_platform,omitempty"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 卷的大小
	VolumeSize int64 `json:"volume_size"`
	// 卷的类型
	VolumeType string `json:"volume_type,omitempty"`
	// 卷关联网关的ID
	IscsiGwId string `json:"iscsi_gw_id"`
	// 卷的IO优先级
	IoPriority string `json:"io_priority,omitempty"`
	// 资源的描述
	Description string `json:"description"`
	// 卷所在客户端的ID
	ClientId string `json:"client_id"`
	// 卷的底层名称
	VolumeName string `json:"volume_name,omitempty"`
}
