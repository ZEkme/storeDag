package dagestan

type User struct {
	UserId               uint   `json:"userID"`
	Name                 string `json:"name"`
	Email                string `json:"e-mail"`
	Password             string `json:"password"`
	FavoritePlaceIds     []int  `json:"favoritePlaceIds"`
	FavoriteMarshrutsIds []int  `json:"favoriteMarshrutsIds"`
}
