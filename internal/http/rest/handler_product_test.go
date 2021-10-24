package rest

import (
	"Gymondo/internal/subscription"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var dummyTime = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)

func TestHandler_Product(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       subscription.Product
		initService        func() subscription.Service
	}{
		{
			description:        "not found must return 404 status code",
			expectedStatusCode: http.StatusNotFound,
			expectedData:       subscription.Product{},
			initService: func() subscription.Service {
				mockRep := new(MockSubscriptionService)
				mockRep.On("GetProductById", mock.Anything).Return(&subscription.Product{
					ID:        1,
					Name:      "",
					CreatedAt: dummyTime,
					UpdatedAt: dummyTime,
					DeletedAt: gorm.DeletedAt{},
				}, nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			handler := CreateHandler(service)
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			router := gin.Default()
			router.GET("/product/:id", handler.Product)

			req, _ := http.NewRequest("GET", "/product", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}

func TestHandler_Products(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       []subscription.Product
		initService        func() subscription.Service
	}{
		{
			description:        "200 ok",
			expectedStatusCode: http.StatusOK,
			expectedData:       []subscription.Product{},
			initService: func() subscription.Service {
				mockRep := new(MockSubscriptionService)
				mockRep.On("GetProductsList").Return([]subscription.Product{}, nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			handler := CreateHandler(service)
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			router := gin.Default()
			router.GET("/products", handler.Products)

			req, _ := http.NewRequest("GET", "/products", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}

func TestHandler_BuyProduct(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       subscription.UserPlan
		initService        func() subscription.Service
	}{
		{
			description:        "400 request without params",
			expectedStatusCode: http.StatusBadRequest,
			expectedData:       subscription.UserPlan{},
			initService: func() subscription.Service {
				mockRep := new(MockSubscriptionService)
				mockRep.On("BuyProduct", mock.Anything).Return(subscription.UserPlan{}, nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			handler := CreateHandler(service)
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			router := gin.Default()
			router.GET("/buy_product", handler.BuyProduct)

			req, _ := http.NewRequest("GET", "/buy_product", nil)
			q := req.URL.Query()
			q.Add("userId", "7")
			req.URL.RawQuery = q.Encode()
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}

func TestHandler_FetchPlansByUserId(t *testing.T) {
	tests := []struct {
		description        string
		expectedStatusCode int
		expectedData       subscription.Status
		initService        func() subscription.Service
	}{
		{
			description:        "404 not found",
			expectedStatusCode: http.StatusNotFound,
			expectedData:       subscription.Status{},
			initService: func() subscription.Service {
				mockRep := new(MockSubscriptionService)
				mockRep.On("FetchPlansByUserId", mock.Anything).Return(subscription.Status{}, nil)
				return mockRep
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			service := tt.initService()
			handler := CreateHandler(service)
			gin.SetMode(gin.TestMode)
			gin.DefaultWriter = ioutil.Discard
			router := gin.Default()
			router.GET("/my_plan", handler.FetchPlansByUserId)

			req, _ := http.NewRequest("GET", "/my_plan", nil)
			q := req.URL.Query()
			q.Set("user_id", "1")
			req.URL.RawQuery = q.Encode()

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, w.Code, tt.expectedStatusCode)
		})
	}
}
