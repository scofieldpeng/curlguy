package curl

import (
	"net/http"

)

// Request 请求
type Request struct {
	header http.Header // 请求携带的头部被信息
	host string // 请求携带的host
	query map[string]string // 请求携带的query参数
	body *http.Request // 请求携带的数据
}

// Response 返回的结构体
type Response struct {
	
}

type Curl struct {
	header          http.Header // 请求携带头部信息
	requestType     string      // 请求方式
	responseContext *http.Response // 返回的结
}
