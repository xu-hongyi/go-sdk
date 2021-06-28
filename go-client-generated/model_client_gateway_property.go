/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 客户端关联网关的详情信息
type ClientGatewayProperty struct {
	// 网关的IQN
	Iqn string `json:"iqn"`
	// 网关开启CHAP的用户名
	ChapUsername string `json:"chap_username,omitempty"`
	// 网关的ID
	Id string `json:"id"`
	// 网关使用的端口号
	Port int32 `json:"port"`
	// 网关的最后更新时间
	UpdateTime int64 `json:"update_time,omitempty"`
	// 网关关联的卷数量
	LunCount int32 `json:"lun_count,omitempty"`
	// 网关是否开启CHAP
	IsChap bool `json:"is_chap,omitempty"`
	// 网关的名称
	Name string `json:"name"`
	// 网关添加的网关节点数量
	NodeCount int32 `json:"node_count,omitempty"`
	// 网关是否启用
	IsEnabled bool `json:"is_enabled,omitempty"`
	// 网关开启CHAP的密码
	ChapPassword string `json:"chap_password,omitempty"`
	// 资源的运行状态
	RunStatus string `json:"run_status"`
	// 网关的创建时间
	CreateTime int64 `json:"create_time,omitempty"`
	// 网关关联的客户端数量
	ClientCount int32 `json:"client_count,omitempty"`
	// 网关的描述信息
	Description string `json:"description"`
}
