package response

import (
	listingModel "github.com/rizalherniawan/99-backend-test/public-api/internal/listing/model"
	userModel "github.com/rizalherniawan/99-backend-test/public-api/internal/user/model"
)

type GetAllListingResponse struct {
	Result   bool                   `json:"result"`
	Listings []listingModel.Listing `json:"listings"`
}

type ListingsResponse struct {
	listingModel.Listing
	User userModel.User `json:"user"`
}
