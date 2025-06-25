package model

type Listing struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int64  `json:"price"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
