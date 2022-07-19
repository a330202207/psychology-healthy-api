// @Project: psychology-healthy-api
// @Author: NedRen
// @Description:
// @Version: 1.0.0
// @Date: 2022/7/18 14:42
// @Package v1

package v1

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `json:"dateRange"` // 日期范围
	PageNum   int      `json:"pageNum"`   // 当前页码
	PageSize  int      `json:"pageSize"`  // 每页数
	OrderBy   string   `json:"orderBy"`   // 排序方式
}

// PageBaseInfo .
type PageBaseInfo struct {
	Page     uint `json:"page"` // 目标页
	PageSize uint `json:"pageSize"`
}

// PageInfo .
type PageInfo struct {
	PageBaseInfo
	Total uint `json:"total"`
}
