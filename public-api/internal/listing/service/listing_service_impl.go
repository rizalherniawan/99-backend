package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/rizalherniawan/99-backend-test/public-api/config"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/exception"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/dto/request"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/dto/response"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/model"
	userModel "github.com/rizalherniawan/99-backend-test/public-api/internal/user/model"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/user/service"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/util"
)

type ListingServiceImpl struct {
	userService service.UserService
}

func New(userService service.UserService) ListingService {
	return &ListingServiceImpl{
		userService: userService,
	}
}

func getServiceInfo() (string, string) {
	return config.GetEnv("LISTING_SERVICE_API_KEY"), config.GetEnv("LISTING_SERVICE_HOST")
}

// CreateListing implements ListingService.
func (l *ListingServiceImpl) CreateListing(req request.CreateListing) (*model.Listing, *exception.ApiError) {
	// validate if userId exists
	_, er := l.userService.GetUserById(strconv.Itoa(req.UserId))
	if er != nil {
		return nil, er
	}

	var res response.CreateListingResponse

	// get necessary API information
	key, host := getServiceInfo()

	// form the url form
	form := url.Values{}
	form.Add("user_id", strconv.Itoa(req.UserId))
	form.Add("listing_type", req.ListingType)
	form.Add("price", strconv.Itoa(int(req.Price)))

	// call the listing service to add listing
	resp, e := util.CallOtherApi("POST", host, key, nil, form)
	if e != nil {
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	// extract response body
	body, er := util.ExtractResponse(resp)
	if er != nil {
		return nil, er
	}
	e = json.Unmarshal(body, &res)
	if e != nil {
		log.Println("failed to unmarshal json due to: ", er.Error())
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}
	return &res.Listing, nil
}

func (l *ListingServiceImpl) GetAllListing(page string, size string, userId string) ([]response.ListingsResponse, *exception.ApiError) {
	// get necessary API information
	key, host := getServiceInfo()

	// form the full path by adding query parameters
	fullPath := host + fmt.Sprintf("?page_num=%s&page_size=%s", page, size)
	if userId != "" {
		fullPath += fmt.Sprintf("&user_id=%s", userId)
	}

	// call the listing service to get all listing data
	resp, e := util.CallOtherApi("GET", fullPath, key, nil, nil)
	if e != nil {
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	// extract the response body
	body, er := util.ExtractResponse(resp)
	if er != nil {
		return nil, er
	}
	var res response.GetAllListingResponse
	e = json.Unmarshal(body, &res)
	if e != nil {
		log.Println("failed to unmarshal json due to: ", er.Error())
		return nil, exception.NewApiError("internal server error", http.StatusInternalServerError)
	}

	// collecting unique userId
	userIds := map[int]interface{}{}
	for _, val := range res.Listings {
		userIds[val.UserId] = nil
	}

	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(len(userIds))

	errCh := make(chan *exception.ApiError, len(userIds)+1)

	// fetch all users based on unique userId
	for key := range userIds {
		go func(userId int) {
			defer wg.Done()
			val, er := l.userService.GetUserById(strconv.Itoa(userId))
			if er != nil {
				errCh <- er
				return
			}
			mu.Lock()
			userIds[userId] = val
			mu.Unlock()
		}(key)
	}

	wg.Wait()

	close(errCh)

	// return an error if system excounters an error
	for val := range errCh {
		return []response.ListingsResponse{}, val
	}

	var listings []response.ListingsResponse

	// map the user data to listing based on user_id
	for _, val := range res.Listings {
		user := userIds[val.UserId].(*userModel.User)
		listing := response.ListingsResponse{
			Listing: val,
			User:    *user,
		}
		listings = append(listings, listing)
	}

	return listings, nil

}
