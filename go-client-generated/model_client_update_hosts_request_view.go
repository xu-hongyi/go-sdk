/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 更新客户端的客户端主机的请求参数
type ClientUpdateHostsRequestView struct {
	// 网关待添加或者移除的网关节点
	Hosts []string `json:"hosts,omitempty"`
	// 更新客户端主机的操作，添加或者移除
	Action string `json:"action"`
}
