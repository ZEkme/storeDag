package dagestan

type Place struct {
	PlaceId     uint     `json:"placeId"`
	CoverUrl    []string `json:"coverUrl"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	X           float64  `json:"x"`
	Y           float64  `json:"y"`
	MarshrutIds []string `json:"marshrutIDs"`
}
