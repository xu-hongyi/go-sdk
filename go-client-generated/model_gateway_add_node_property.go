/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 网关添加的网关节点属性信息
type GatewayAddNodeProperty struct {
	// 网关节点的网关IP
	GatewayIp string `json:"gateway_ip"`
	// 网关节点的ID
	ServerId string `json:"server_id"`
}
