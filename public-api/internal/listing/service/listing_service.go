package service

import (
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/dto/request"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/dto/response"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/model"
)

type ListingService interface {
	CreateListing(req request.CreateListing) (*model.Listing, *exception.ApiError)
	GetAllListing(page string, size string, userId string) ([]response.ListingsResponse, *exception.ApiError)
}
