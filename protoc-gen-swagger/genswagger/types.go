package genswagger

import (
	"github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
)

type param struct {
	*descriptor.File
	reg *descriptor.Registry
}

type binding struct {
	*descriptor.Binding
}

// http://swagger.io/specification/#infoObject
type swaggerInfoObject struct {
	Version string `json:"version"`
	Title   string `json:"title"`
}

// http://swagger.io/specification/#swaggerObject
type swaggerObject struct {
	Swagger     string                   `json:"swagger"`
	Info        swaggerInfoObject        `json:"info"`
	Host        string                   `json:"host,omitempty"`
	BasePath    string                   `json:"basePath,omitempty"`
	Schemes     []string                 `json:"schemes"`
	Consumes    []string                 `json:"consumes"`
	Produces    []string                 `json:"produces"`
	Paths       swaggerPathsObject       `json:"paths"`
	Definitions swaggerDefinitionsObject `json:"definitions"`
}

// http://swagger.io/specification/#pathsObject
type swaggerPathsObject map[string]swaggerPathItemObject

// http://swagger.io/specification/#pathItemObject
type swaggerPathItemObject struct {
	Get    *swaggerOperationObject `json:"get,omitempty"`
	Delete *swaggerOperationObject `json:"delete,omitempty"`
	Post   *swaggerOperationObject `json:"post,omitempty"`
	Put    *swaggerOperationObject `json:"put,omitempty"`
}

// http://swagger.io/specification/#operationObject
type swaggerOperationObject struct {
	Summary    string                  `json:"summary"`
	Responses  swaggerResponsesObject  `json:"responses"`
	Parameters swaggerParametersObject `json:"parameters,omitempty"`
}

type swaggerParametersObject []swaggerParameterObject

// http://swagger.io/specification/#parameterObject
type swaggerParameterObject struct {
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	In          string              `json:"in,omitempty"`
	Required    bool                `json:"required"`
	Type        string              `json:"type,omitempty"`
	Format      string              `json:"format,omitempty"`
	Items       *swaggerItemsObject `json:"items,omitempty"`

	// Or you can explicitly refer to another type. If this is defined all
	// other fields should be empty
	Schema *swaggerSchemaObject `json:"schema,omitempty"`
}

// http://swagger.io/specification/#itemsObject
type swaggerItemsObject struct {
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Ref    string `json:"$ref,omitempty"`
}

// http://swagger.io/specification/#responsesObject
type swaggerResponsesObject map[string]swaggerResponseObject

// http://swagger.io/specification/#responseObject
type swaggerResponseObject struct {
	Description string              `json:"description"`
	Schema      swaggerSchemaObject `json:"schema"`
}

// http://swagger.io/specification/#schemaObject
type swaggerSchemaObject struct {
	Ref    string `json:"$ref,omitempty"`
	Type   string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	// Properties can be recursively defined
	Properties map[string]swaggerSchemaObject `json:"properties,omitempty"`
	Items      *swaggerItemsObject            `json:"items,omitempty"`
}

// http://swagger.io/specification/#referenceObject
type swaggerReferenceObject struct {
	Ref string `json:"$ref"`
}

// http://swagger.io/specification/#definitionsObject
type swaggerDefinitionsObject map[string]swaggerSchemaObject

type messageMap map[string]*descriptor.Message
