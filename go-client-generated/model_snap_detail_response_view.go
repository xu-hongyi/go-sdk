/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 快照详情 返回的响应信息
type SnapDetailResponseView struct {
	// 快照大小
	SnapSize int64 `json:"snap_size"`
	// 所属卷的uuid
	VolumeId string `json:"volume_id"`
	// 存储池的名称
	PoolName string `json:"pool_name"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 资源的创建方式
	CreateFrom string `json:"create_from"`
	// 一致性组快照的名称
	CgSnapName string `json:"cg_snap_name"`
	// 资源的容量健康状态
	CapHealthStatus string `json:"cap_health_status"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 存储集群中快照的名称
	SnapName string `json:"snap_name"`
	// 存储池的uuid
	PoolId string `json:"pool_id"`
	// 资源的同步状态
	SyncStatus string `json:"sync_status"`
	// 资源的容量健康状态非健康的原因
	CapHealthStatusReason string `json:"cap_health_status_reason"`
	// 克隆卷数量
	CloneVolNum float32 `json:"clone_vol_num"`
	// 数据量
	DataSize int64 `json:"data_size"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 资源的名称
	Name string `json:"name"`
	// 资源的描述
	Description string `json:"description"`
	// 资源的ID
	Id string `json:"id"`
	// 一致性组快照的uuid
	CgSnapId string `json:"cg_snap_id"`
	// 所属卷再存储集群中的名称
	VolumeName string `json:"volume_name"`
	// 所属卷名称
	VolumeDispName string `json:"volume_disp_name"`
	Message        string `json:"message"`
}
