/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 快照响应的基本视图
type SnapResponseView struct {
	// 数据量
	DataSize float32 `json:"data_size"`
	// 资源的ID
	Id string `json:"id"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 资源的创建方式
	CreateFrom string `json:"create_from"`
	// 资源的名称
	Name string `json:"name"`
	// 快照大小
	SnapSize float32 `json:"snap_size"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 所属卷的uuid
	VolumeId string `json:"volume_id"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 存储集群中快照的名称
	SnapName string `json:"snap_name"`
	// 资源的同步状态
	SyncStatus string `json:"sync_status"`
	// 资源的描述
	Description string `json:"description"`
}