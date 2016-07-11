package request

type (
	RequestType  string // 请求格式
	ResponseType string // 返回格式
)

const (
	POST   RequestType = "POST"   // post请求
	GET    RequestType = "GET"    // get请求
	PUT    RequestType = "PUT"    // put请求
	DELETE RequestType = "DELETE" // delete请求
)

// RequestData 接口请求格式
type RequestData struct {
	requestT RequestType       `json:"request_type"` // 接口请求格式
	Data     string            `json:"data"`         // 接口请求数据
	Header   map[string]string `json:"header"`       // 接口请求携带的header
}