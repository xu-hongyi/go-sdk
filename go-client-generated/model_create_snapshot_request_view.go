/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 创建快照请求参数
type CreateSnapshotRequestView struct {
	// 快照的描述信息
	Description string `json:"description,omitempty"`
	// 快照名称
	Name string `json:"name"`
	// 卷的uuid
	VolumeId string `json:"volume_id"`
}