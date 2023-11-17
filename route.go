package dagestan

type Route struct {
	RouteId     uint     `json:"routeId"`
	Title       string   `json:"title"`
	Places      []string `json:"places"`
	Description string   `json:"description"`
}
