/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 更新卷的请求参数
type UpdateVolumeRequestView struct {
	// 更新卷的数据校验
	VerifyEnabled bool `json:"verify_enabled,omitempty"`
	// 更新卷的名称
	Name string `json:"name,omitempty"`
	// 更新卷的IO性能优先级
	IoPriority string `json:"io_priority,omitempty"`
}
