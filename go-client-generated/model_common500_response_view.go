/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 请求返回500，出现内部错误
type Common500ResponseView struct {
	// 请求失败时的原因，正常时为空
	Message string `json:"message"`
	// 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff
	RetCode string `json:"ret_code"`
}
