/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 用户角色信息 响应结果
type AuthRoleResponseView struct {
	// 资源的ID
	Id string `json:"id"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 资源的名称
	Name string `json:"name"`
	// 权限信息列表
	AuthList []AuthInfoResponseView `json:"auth_list"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 资源的描述
	Description string `json:"description"`
}
