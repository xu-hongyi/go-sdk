/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 更新卷名称的请求参数
type UpdateVolumeNameRequestView struct {
	// 更新卷的名称
	Name string `json:"name"`
}