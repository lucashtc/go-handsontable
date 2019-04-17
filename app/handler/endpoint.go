package handler

// Endpoint routers
type Endpoint struct {
	Ping string
	GetAll string
	GetFind string
	Delete string
}

// SetEndpoint return link routers
func SetEndpoint() *Endpoint {
	point := &Endpoint{
		Ping : "/ping",
		GetAll: "/get",
		GetFind: "/find",
	}
	return point

}