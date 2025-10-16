package docs

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	"xm/utils"
)

// swagger:route POST /v1/company Company idCreateCompany
// Creates a Company
//
// security:
//	- bearer:
//
// responses:
// 		200: CreateCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idCreateCompany
type CreateCompanyRequestWrapper struct {
	// in:body
	Body request.CreateCompanyRequest
}

// swagger:response CreateCompanyResponseWrapper
type CreateCompanyResponseWrapper struct {
	// in:body
	Body models.Company
}

// swagger:route GET /v1/company/{company_id} Company idGetCompany
// Fetch company data
//
// security:
//	- bearer:
//
// responses:
// 		200: GetCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idGetCompany
type GetCompanyRequestWrapper struct {
	// In: path
	CompanyID string `json:"company_id"`
}

// swagger:response GetCompanyResponseWrapper
type GetCompanyResponseWrapper struct {
	// in:body
	Body models.Company
}

// swagger:route PATCH /v1/company/{company_id} Company idPatchCompany
// Patch company data
//
// security:
//	- bearer:
//
// responses:
// 		200: PatchCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idPatchCompany
type PatchCompanyRequestWrapper struct {
	// In: path
	CompanyID string `json:"company_id"`
	// in: body
	Body request.PatchCompanyRequest
}

// swagger:response PatchCompanyResponseWrapper
type PatchCompanyResponseWrapper struct {
	// in:body
	Body utils.Response
}

// swagger:route DELETE /v1/company/{company_id} Company idDeleteCompany
// Delete company data
//
// security:
//	- bearer:
//
// responses:
// 		200: DeleteCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idDeleteCompany
type DeleteCompanyRequestWrapper struct {
	// In: path
	CompanyID string `json:"company_id"`
}

// swagger:response DeleteCompanyResponseWrapper
type DeleteCompanyResponseWrapper struct {
	// in:body
	Body utils.Response
}
