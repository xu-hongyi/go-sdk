/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 创建Qos响应结果
type CreateQosResponseView struct {
	// qos策略的uuid
	QosPolicyId string `json:"qos_policy_id,omitempty"`
	Message     string `json:"message"`
}
