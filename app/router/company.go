package router

import (
	"xm/internal/entity/request"
	"xm/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (c *Capsule) CompanyRoutes(r *gin.RouterGroup) {
	companyHandler := c.Handler.CompanyHandler
	mdl := c.Middleware

	companyV1 := r.Group("/company", mdl.Authorize())

	companyV1.GET("/:company_id",
		companyHandler.HandleGetCompany)
	companyV1.POST("",
		middleware.Validate[request.CreateCompanyRequest]("CreateCompanyRequest"),
		companyHandler.HandleCreateCompany)
	companyV1.PATCH("/:company_id",
		middleware.Validate[request.PatchCompanyRequest]("PatchCompanyRequest"),
		companyHandler.HandlePatchCompany)
	companyV1.DELETE("/:company_id",
		companyHandler.HandleDeleteCompany)
}
