// Package advice_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package advice_kit_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

// Defines values for DescribingEndpointReferenceMethod.
const (
	GET DescribingEndpointReferenceMethod = "GET"
)

// AdviceDefinition Provides details about a advice
type AdviceDefinition struct {
	// AssessmentQueryApplicable A Assessment Target Base Query that is used identifies targets that could have this advice.
	AssessmentQueryApplicable string `json:"assessmentQueryApplicable"`

	// Icon A svg of an icon that represents the advice.
	Icon string `json:"icon"`

	// Id A technical ID that is used to uniquely identify this type of advice. You will typically want to use something like `org.example.extension.my-fancy-advice`.
	Id string `json:"id"`

	// Label A human-readable label for the advice.
	Label string `json:"label"`

	// Status Provides details about a advice
	Status AdviceDefinitionStatus `json:"status"`

	// Tags A list of tags that describe the advice.
	Tags *[]string `json:"tags,omitempty"`

	// Version The version of the advice. This is used to identify the version of the advice and is used for compatibility checks.
	Version string `json:"version"`
}

// AdviceDefinitionStatus Provides details about a advice
type AdviceDefinitionStatus struct {
	// ActionNeeded Provides details about a advice lifecycle status actions needed
	ActionNeeded AdviceDefinitionStatusActionNeeded `json:"actionNeeded"`

	// Implemented Provides details about a advice lifecycle status implemented
	Implemented AdviceDefinitionStatusImplemented `json:"implemented"`

	// ValidationNeeded Provides details about a advice lifecycle status validation needed
	ValidationNeeded AdviceDefinitionStatusValidationNeeded `json:"validationNeeded"`
}

// AdviceDefinitionStatusActionNeeded Provides details about a advice lifecycle status actions needed
type AdviceDefinitionStatusActionNeeded struct {
	// AssessmentQuery A Assessment Target Query Addon that is used to identify targets with this advice in the target list of the assessmentQueryApplicable
	AssessmentQuery string `json:"assessmentQuery"`

	// Description Provides details description about a advice lifecycle status actions needed
	Description AdviceDefinitionStatusActionNeededDescription `json:"description"`
}

// AdviceDefinitionStatusActionNeededDescription Provides details description about a advice lifecycle status actions needed
type AdviceDefinitionStatusActionNeededDescription struct {
	// Instruction A human-readable instructions of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Instruction string `json:"instruction"`

	// Motivation A human-readable motivation of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Motivation string `json:"motivation"`

	// Summary A human-readable summary of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionStatusImplemented Provides details about a advice lifecycle status implemented
type AdviceDefinitionStatusImplemented struct {
	// Description Provides details description about a advice lifecycle status implemented
	Description AdviceDefinitionStatusImplementedDescription `json:"description"`
}

// AdviceDefinitionStatusImplementedDescription Provides details description about a advice lifecycle status implemented
type AdviceDefinitionStatusImplementedDescription struct {
	// Summary A human-readable summary of the implemented in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionStatusValidationNeeded Provides details about a advice lifecycle status validation needed
type AdviceDefinitionStatusValidationNeeded struct {
	// Description Provides details description about a advice lifecycle status validation needed
	Description AdviceDefinitionStatusValidationNeededDescription `json:"description"`

	// Validation A list of validations that are available for this advice.
	Validation *[]Validation `json:"validation,omitempty"`
}

// AdviceDefinitionStatusValidationNeededDescription Provides details description about a advice lifecycle status validation needed
type AdviceDefinitionStatusValidationNeededDescription struct {
	// Summary A human-readable summary of the validation needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceKitError RFC 7807 Problem Details for HTTP APIs compliant response body for error scenarios
type AdviceKitError struct {
	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`

	// Title A short, human-readable summary of the problem type.
	Title string `json:"title"`

	// Type A URI reference that identifies the problem type.
	Type *string `json:"type,omitempty"`
}

// AdviceList Lists all advice that the platform/agent could consume.
type AdviceList struct {
	Advice []DescribingEndpointReference `json:"advice"`
}

// DescribingEndpointReference HTTP endpoint which the Steadybit platform/agent could communicate with.
type DescribingEndpointReference struct {
	// Method HTTP method to use when calling the HTTP endpoint.
	Method DescribingEndpointReferenceMethod `json:"method"`

	// Path Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// DescribingEndpointReferenceMethod HTTP method to use when calling the HTTP endpoint.
type DescribingEndpointReferenceMethod string

// Experiment Provides a experiment json exported from the ui
type Experiment = interface{}

// Validation Provides a either a template about a advice experiment or a textual validation like a checklist item
type Validation struct {
	// Description A human-readable description for the validation or for the experiment template.
	Description string `json:"description"`

	// Experiment Provides a experiment json exported from the ui
	Experiment *Experiment `json:"experiment,omitempty"`

	// Id A technical ID that is used to uniquely identify this validation. You will typically want to use something like `org.example.extension.my-fancy-advice-validation.1`.
	Id string `json:"id"`

	// Name A human-readable name for the validation.
	Name string `json:"name"`

	// Type The type of the validation. Either `experiment` or `text`.
	Type string `json:"type"`
}

// AdviceDefinitionResponse defines model for AdviceDefinitionResponse.
type AdviceDefinitionResponse struct {
	union json.RawMessage
}

// AdviceListResponse defines model for AdviceListResponse.
type AdviceListResponse struct {
	union json.RawMessage
}

// AsAdviceDefinition returns the union data inside the AdviceDefinitionResponse as a AdviceDefinition
func (t AdviceDefinitionResponse) AsAdviceDefinition() (AdviceDefinition, error) {
	var body AdviceDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceDefinition overwrites any union data inside the AdviceDefinitionResponse as the provided AdviceDefinition
func (t *AdviceDefinitionResponse) FromAdviceDefinition(v AdviceDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceDefinition performs a merge with any union data inside the AdviceDefinitionResponse, using the provided AdviceDefinition
func (t *AdviceDefinitionResponse) MergeAdviceDefinition(v AdviceDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAdviceKitError returns the union data inside the AdviceDefinitionResponse as a AdviceKitError
func (t AdviceDefinitionResponse) AsAdviceKitError() (AdviceKitError, error) {
	var body AdviceKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceKitError overwrites any union data inside the AdviceDefinitionResponse as the provided AdviceKitError
func (t *AdviceDefinitionResponse) FromAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceKitError performs a merge with any union data inside the AdviceDefinitionResponse, using the provided AdviceKitError
func (t *AdviceDefinitionResponse) MergeAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AdviceDefinitionResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AdviceDefinitionResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAdviceList returns the union data inside the AdviceListResponse as a AdviceList
func (t AdviceListResponse) AsAdviceList() (AdviceList, error) {
	var body AdviceList
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceList overwrites any union data inside the AdviceListResponse as the provided AdviceList
func (t *AdviceListResponse) FromAdviceList(v AdviceList) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceList performs a merge with any union data inside the AdviceListResponse, using the provided AdviceList
func (t *AdviceListResponse) MergeAdviceList(v AdviceList) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAdviceKitError returns the union data inside the AdviceListResponse as a AdviceKitError
func (t AdviceListResponse) AsAdviceKitError() (AdviceKitError, error) {
	var body AdviceKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceKitError overwrites any union data inside the AdviceListResponse as the provided AdviceKitError
func (t *AdviceListResponse) FromAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceKitError performs a merge with any union data inside the AdviceListResponse, using the provided AdviceKitError
func (t *AdviceListResponse) MergeAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AdviceListResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AdviceListResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
