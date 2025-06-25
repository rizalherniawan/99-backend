package listing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizalherniawan/99-backend-test/public-api/internal/listing/dto/request"
	listingService "github.com/rizalherniawan/99-backend-test/public-api/internal/listing/service"
	userService "github.com/rizalherniawan/99-backend-test/public-api/internal/user/service"
)

type ListingHandler struct {
	userService    userService.UserService
	listingService listingService.ListingService
}

func New(userService userService.UserService, listingService listingService.ListingService) *ListingHandler {
	return &ListingHandler{
		userService:    userService,
		listingService: listingService,
	}
}

func (l *ListingHandler) CreateListing(r *gin.Context) {
	req := request.CreateListing{}
	er := r.ShouldBindJSON(&req)
	if er != nil {
		r.Status(http.StatusNotAcceptable)
		r.Error(er)
		r.Abort()
		return
	}
	res, e := l.listingService.CreateListing(req)
	if e != nil {
		r.Status(e.StatusCode)
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusOK, map[string]interface{}{
		"result":  true,
		"listing": res,
	})
}

func (l *ListingHandler) GetAllListing(r *gin.Context) {
	userId := r.DefaultQuery("user_id", "")
	page := r.DefaultQuery("page_num", "1")
	size := r.DefaultQuery("page_size", "10")

	res, e := l.listingService.GetAllListing(page, size, userId)
	if e != nil {
		r.Status(e.StatusCode)
		r.Error(e)
		r.Abort()
		return
	}

	r.JSON(http.StatusOK, map[string]interface{}{
		"result":   true,
		"listings": res,
	})
}
