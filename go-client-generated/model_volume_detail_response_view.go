/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Volume的详情信息
type VolumeDetailResponseView struct {
	// 卷是否开启数据保护
	DataProtection bool `json:"data_protection"`
	// 父快照的ID，仅用于克隆卷
	ParentSnapId string `json:"parent_snap_id"`
	// 资源的性能健康状态非健康的原因
	PerfHealthStatusReason string `json:"perf_health_status_reason"`
	// 卷所属一致性组的ID
	CgId string `json:"cg_id"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 卷所在存储池的名称
	PoolName string `json:"pool_name"`
	// 卷关联的终端数量
	ClientNum int32 `json:"client_num"`
	// 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff
	RetCode string `json:"ret_code"`
	// 资源的创建方式
	CreateFrom string `json:"create_from"`
	// 资源的容量健康状态
	CapHealthStatus string `json:"cap_health_status"`
	// 卷是否开启数据校验
	VerifyEnabled bool `json:"verify_enabled"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 卷所在存储池的ID
	PoolId string `json:"pool_id"`
	// 资源的同步状态
	SyncStatus string `json:"sync_status"`
	// 资源的性能健康状态
	PerfHealthStatus string `json:"perf_health_status"`
	// 资源的容量健康状态非健康的原因
	CapHealthStatusReason string `json:"cap_health_status_reason"`
	// 卷已使用数据量，单位为字节
	DataSize int64 `json:"data_size"`
	// 请求失败时的原因，正常时为空
	Message string `json:"message"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 卷所属对接平台
	OwnerPlatform string `json:"owner_platform"`
	// 卷的QOS状态
	QosStatus bool `json:"qos_status"`
	// 卷创建的快照数量
	SnapNum int32 `json:"snap_num"`
	// 卷的大小，单位为字节
	VolumeSize int64 `json:"volume_size"`
	// 资源的名称
	Name string `json:"name"`
	// Volume的类型
	VolumeType string `json:"volume_type"`
	// 卷的IO性能优先级
	IoPriority string `json:"io_priority"`
	// 父快照的名称，仅用于克隆卷
	ParentSnapName string `json:"parent_snap_name"`
	// 卷所属一致性组的名称
	CgName string `json:"cg_name"`
	// 资源的描述
	Description string `json:"description"`
	// 资源的ID
	Id string `json:"id"`
	// 卷的QOS配置
	Qos *QosPropertyView `json:"qos,omitempty"`
}