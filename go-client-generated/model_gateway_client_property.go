/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 网关关联的客户端的属性信息
type GatewayClientProperty struct {
	// 客户端已关联的网关数量
	IscsiGwCount int32 `json:"iscsi_gw_count,omitempty"`
	// 网关开启CHAP的用户名
	ChapUsername string `json:"chap_username,omitempty"`
	// 资源的ID
	Id string `json:"id"`
	// 资源的最后时间
	UpdateTime float32 `json:"update_time"`
	// 客户端关联的卷数量
	LunCount int32 `json:"lun_count,omitempty"`
	// 网关是否开启CHAP
	IsChap bool `json:"is_chap,omitempty"`
	// 资源的名称
	Name string `json:"name"`
	// 网关开启CHAP的密码
	ChapPassword string `json:"chap_password,omitempty"`
	// 客户端关联的客户端主机数量
	ClientNodeCount int32 `json:"client_node_count,omitempty"`
	// 客户端关联网关的ID
	IscsiGwId string `json:"iscsi_gw_id,omitempty"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 资源的创建时间
	CreateTime float32 `json:"create_time"`
	// 资源的描述
	Description string `json:"description"`
}
