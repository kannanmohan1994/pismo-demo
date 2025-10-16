// xm
//
// xm API Documentation !!
//
//	 Schemes: http, https
//	 BasePath: /
//	 Version: 1.0.0
//	 Host: localhost:9000
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//
//	SecurityDefinitions:
//	bearer:
//	  type: apiKey
//	  name: Authorization
//	  in: header
//
// swagger:meta
package docs

import "xm/utils"

// swagger:response CustomErrorWrapper
type CustomErrorWrapper struct {
	// in:body
	Body utils.CustomError
}
