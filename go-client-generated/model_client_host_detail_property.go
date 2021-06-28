/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 客户端主机详情的属性信息
type ClientHostDetailProperty struct {
	// 客户端主机的ID
	Id string `json:"id"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 客户端的ID
	ClientId string `json:"client_id"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 客户端主机的IQN
	Node string `json:"node"`
	// 客户端主机的登录状态
	UseStatus string `json:"use_status"`
}
