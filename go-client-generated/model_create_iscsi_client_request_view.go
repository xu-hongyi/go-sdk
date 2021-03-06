/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 创建客户端的请求参数
type CreateIscsiClientRequestView struct {
	// iSCSI网关chap用户名(约束：8-64个字符，支持大小写字母数字、下划线、连字符、’@‘、冒号、句点，用户名和密码必须同时存在)
	ChapUsername string `json:"chap_username,omitempty"`
	// 客户端包含的客户端主机列表
	Hosts []string `json:"hosts"`
	// 客户端是否开启CHAP
	IsChap bool `json:"is_chap"`
	// 客户端的名称
	Name string `json:"name"`
	// iSCSI网关chap密码(约束:12-16个字符，支持大小写字母、数字、下划线、连字符、‘/’、’@‘)
	ChapPassword string `json:"chap_password,omitempty"`
	// iSCSI客户端描述信息（支持utf-8字符,长度256）
	Description string `json:"description,omitempty"`
}
