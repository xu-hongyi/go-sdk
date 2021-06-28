/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 修改快照名称请求参数
type RenameSnapRequestView struct {
	// 快照名称
	Name string `json:"name,omitempty"`
	// 快照的描述信息
	Description string `json:"description,omitempty"`
}
