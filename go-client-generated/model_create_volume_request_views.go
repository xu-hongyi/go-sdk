/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 创建单个Volume的请求参数
type CreateVolumeRequestViews struct {
	// 卷的分配容量即卷大小。单位为Bytes, 最小1GB大小
	VolumeSize int64 `json:"volume_size"`
	// 卷名称
	Name string `json:"name"`
	// 是否数据校验。
	VerifyEnabled bool `json:"verify_enabled"`
	// 所属平台
	OwnerPlatform string `json:"owner_platform,omitempty"`
	// IO性能优先级。只能是default或priority
	IoPriority string `json:"io_priority"`
	// Volume的QOS配置
	Qos *QosPropertyView `json:"qos"`
	// 卷所属存储池的uuid
	PoolId string `json:"pool_id"`
}
