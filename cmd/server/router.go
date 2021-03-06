package main

import (
	"Gymondo/internal/http/rest"
	"Gymondo/internal/metrics"

	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
)

func SetupRouter(handler *rest.Handler, cfg *MainConfig, p *metrics.Prometheus) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, rest.NotFound)
	})

	r.GET("/health", handler.Health)

	v1 := r.Group("/v1")
	if cfg.Server.AuthEnabled {
		v1.Use(gin.BasicAuth(gin.Accounts{
			cfg.Server.User: cfg.Server.Pass,
		}))
	}

	{
		v1.GET("/browse", handler.Health)
		v1.GET("/dummy", handler.DummyDataGenerator)
		v1.GET("/products", handler.Products)
		v1.GET("/product/:id", handler.Product)
		v1.GET("/products_by_voucher/:id", handler.ProductByVoucher)
		v1.GET("/buy_product", handler.BuyProduct)
		v1.GET("/change_plan_status", handler.ChangePlanStatus)
		v1.GET("/my_plan/:user_id", handler.FetchPlansByUserId)
	}

	p.MetricsPath = fmt.Sprintf("/%s", "metrics")
	if cfg.Server.AuthEnabled {
		p.UseWithAuth(r, gin.Accounts{
			cfg.Server.User: cfg.Server.Pass,
		})
	} else {
		p.Use(r)
	}

	var AllowedRoutes = make(map[string]bool, 0)
	routes := r.Routes()
	for _, i := range routes {
		AllowedRoutes[i.Path] = true
	}

	paramStripMap := make(map[string]bool, 0)

	for _, sp := range []string{"q", "username"} {
		paramStripMap[sp] = true
	}

	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if _, ok := paramStripMap[p.Key]; ok {
				// Found param
				url = strings.Replace(url, p.Value, fmt.Sprintf(":%s", p.Key), 1)
			}
		}
		if _, ok := AllowedRoutes[url]; ok {
			return url
		} else {
			return ""
		}
	}
	return r
}
