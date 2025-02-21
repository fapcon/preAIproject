// Package classification infoblog.
//
// Documentation of auto trade API.
//
//	Schemes:
//	- http
//	- https
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	Security:
//	- basic
//
//
//	SecurityDefinitions:
//	  Bearer:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//
// swagger:meta
package docs

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models
