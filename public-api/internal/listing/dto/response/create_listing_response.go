package response

import "github.com/rizalherniawan/99-backend-test/public-api/internal/listing/model"

type CreateListingResponse struct {
	Result  bool          `json:"result"`
	Listing model.Listing `json:"listing"`
}
