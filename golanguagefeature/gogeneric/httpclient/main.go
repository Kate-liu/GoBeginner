package httpclient

type HttpClient struct {
	endpoint string
}

type GetBuilder[T any] struct {
	client *HttpClient
	path   string
}

func (g *GetBuilder[T]) Path(path string) *GetBuilder[T] {
	g.path = path
	return g
}

func (g *GetBuilder[T]) Do() T {
	// 真实发出 HTTP 请求
	url := g.client.endpoint + g.path
}

func NewGetRequest[T any](client *HttpClient) *GetBuilder {
	return &GetBuilder[T]{client: client}
}
