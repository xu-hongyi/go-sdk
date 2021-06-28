/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Volume的Qos配置信息
type QosPropertyView struct {
	// 最大带宽，默认为0，为表示无限制
	BpsLimit int64 `json:"bps_limit,omitempty"`
	// 突发IOPS，默认为0，为表示无限制
	IopsBurst int64 `json:"iops_burst,omitempty"`
	// 最大IOPS，默认为0，为表示无限制
	IopsLimit int64 `json:"iops_limit,omitempty"`
	// 突发带宽，默认为0，为表示无限制
	BpsBurst int64 `json:"bps_burst,omitempty"`
}