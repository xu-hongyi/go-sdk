/*
 * ET-GDS
 *
 * ET-GDS API doc.
 *
 * API version: ET-GDS API 7.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// 可用卷设备详情信息
type AvailLunDetailProperty struct {
	// 卷设备的数据量
	DataSize int64 `json:"data_size"`
	// 卷设备的ID
	Id string `json:"id"`
	// 卷设备的大小
	VolumeSize int64 `json:"volume_size"`
	// 卷设备的名称
	Name string `json:"name"`
	// 卷设备所在存储池的名称
	PoolName string `json:"pool_name"`
	// 卷设备的名称
	LunType string `json:"lun_type"`
}
