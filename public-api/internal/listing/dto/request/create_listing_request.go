package request

type CreateListing struct {
	UserId      int    `json:"user_id" binding:"required"`
	ListingType string `json:"listing_type" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
}
