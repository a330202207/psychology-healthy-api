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
	Page     int `json:"page"  d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	PageSize int `json:"pageSize"  d:"20" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}

// PageInfo .
type PageInfo struct {
	PageBaseInfo
	Total int `json:"total" dc:"总数"` // 总数
}
