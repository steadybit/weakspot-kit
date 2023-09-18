// Package weakspot_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package weakspot_kit_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for DescribingEndpointReferenceMethod.
const (
	GET DescribingEndpointReferenceMethod = "GET"
)

// DescribingEndpointReference HTTP endpoint which the Steadybit platform/agent could communicate with.
type DescribingEndpointReference struct {
	// Method HTTP method to use when calling the HTTP endpoint.
	Method DescribingEndpointReferenceMethod `json:"method"`

	// Path Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// DescribingEndpointReferenceMethod HTTP method to use when calling the HTTP endpoint.
type DescribingEndpointReferenceMethod string

// Experiment Provides a template about a weakspot experiment
type Experiment = interface{}

// WeakspotDescription Provides details about a weakspot
type WeakspotDescription struct {
	// AssesmentBaseQuery A Assessment Target Base Query that is used identifies targets that could have this weakspot.
	AssesmentBaseQuery string `json:"assesmentBaseQuery"`

	// AssesmentQueryAddon A Assessment Target Query Addon that is used to identify targets with this weakspot.
	AssesmentQueryAddon string `json:"assesmentQueryAddon"`

	// Experiments A list of experiments that are available for this weakspot.
	Experiments *[]Experiment `json:"experiments,omitempty"`

	// Icon A svg of an icon that represents the weakspot.
	Icon string `json:"icon"`

	// Id A technical ID that is used to uniquely identify this type of weakspot. You will typically want to use something like `org.example.extension.my-fancy-weakspot`.
	Id string `json:"id"`

	// Label A human-readable label for the weakspot.
	Label string `json:"label"`

	// Tags A list of tags that describe the weakspot.
	Tags *[]string `json:"tags,omitempty"`

	// Version The version of the weakspot. This is used to identify the version of the weakspot and is used for compatibility checks.
	Version string `json:"version"`
}

// WeakspotKitError RFC 7807 Problem Details for HTTP APIs compliant response body for error scenarios
type WeakspotKitError struct {
	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`

	// Title A short, human-readable summary of the problem type.
	Title string `json:"title"`

	// Type A URI reference that identifies the problem type.
	Type *string `json:"type,omitempty"`
}

// WeakspotList Lists all weakspots that the platform/agent could consume.
type WeakspotList struct {
	Weakspots []DescribingEndpointReference `json:"weakspots"`
}

// WeakspotDescriptionResponse defines model for WeakspotDescriptionResponse.
type WeakspotDescriptionResponse struct {
	union json.RawMessage
}

// WeakspotListResponse defines model for WeakspotListResponse.
type WeakspotListResponse struct {
	union json.RawMessage
}

// AsWeakspotDescription returns the union data inside the WeakspotDescriptionResponse as a WeakspotDescription
func (t WeakspotDescriptionResponse) AsWeakspotDescription() (WeakspotDescription, error) {
	var body WeakspotDescription
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWeakspotDescription overwrites any union data inside the WeakspotDescriptionResponse as the provided WeakspotDescription
func (t *WeakspotDescriptionResponse) FromWeakspotDescription(v WeakspotDescription) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWeakspotDescription performs a merge with any union data inside the WeakspotDescriptionResponse, using the provided WeakspotDescription
func (t *WeakspotDescriptionResponse) MergeWeakspotDescription(v WeakspotDescription) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsWeakspotKitError returns the union data inside the WeakspotDescriptionResponse as a WeakspotKitError
func (t WeakspotDescriptionResponse) AsWeakspotKitError() (WeakspotKitError, error) {
	var body WeakspotKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWeakspotKitError overwrites any union data inside the WeakspotDescriptionResponse as the provided WeakspotKitError
func (t *WeakspotDescriptionResponse) FromWeakspotKitError(v WeakspotKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWeakspotKitError performs a merge with any union data inside the WeakspotDescriptionResponse, using the provided WeakspotKitError
func (t *WeakspotDescriptionResponse) MergeWeakspotKitError(v WeakspotKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t WeakspotDescriptionResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WeakspotDescriptionResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsWeakspotList returns the union data inside the WeakspotListResponse as a WeakspotList
func (t WeakspotListResponse) AsWeakspotList() (WeakspotList, error) {
	var body WeakspotList
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWeakspotList overwrites any union data inside the WeakspotListResponse as the provided WeakspotList
func (t *WeakspotListResponse) FromWeakspotList(v WeakspotList) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWeakspotList performs a merge with any union data inside the WeakspotListResponse, using the provided WeakspotList
func (t *WeakspotListResponse) MergeWeakspotList(v WeakspotList) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsWeakspotKitError returns the union data inside the WeakspotListResponse as a WeakspotKitError
func (t WeakspotListResponse) AsWeakspotKitError() (WeakspotKitError, error) {
	var body WeakspotKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromWeakspotKitError overwrites any union data inside the WeakspotListResponse as the provided WeakspotKitError
func (t *WeakspotListResponse) FromWeakspotKitError(v WeakspotKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeWeakspotKitError performs a merge with any union data inside the WeakspotListResponse, using the provided WeakspotKitError
func (t *WeakspotListResponse) MergeWeakspotKitError(v WeakspotKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t WeakspotListResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *WeakspotListResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xW348jNQz+V6xwT2iu3RMPh/q2sAusALEsRQidii7NuJ3cZZK5xLPd0an/O3LmZzvZ",
	"HyDBWztx7M+f7S/+LJQrK2fRUhCrz8JjqJwNGP/8gfJjqBxdYVBeV6SdvevO+Vg5S2iJf8qqMlpJtlh+",
	"CM7yt6AKLGU8tc0vO7F691m88rgTK/HFcoy6bO3CMhFNHLOX3flR07X3zovj5ng8ZiKf+FiJHjTsnAcq",
	"ENrjLcKhuw9o88ppSyITpMmgWImr3qgPAkPyx2zg5icd6H8ghcP8V2wMJBgdkkwMBDCKCQvHrMsndkvH",
	"l7b7687FHe7Qo1WRmFMQP6zXt0MoOBRaFRHLb4Qyb7aaoDKSds6XS7lHS6BcbXJQrixry6QiHDQVC5GJ",
	"yrsKPem2Z0ukwuWPBGwPgRzUAeFQoAUljdF2H4OfgGLXaOtSrN6J76/XYpMJaiqmI5DXds9NUEkq5qEu",
	"t8GZmhD4GNwu7buSROjZ/q/l4stXYub+mAmPn2rtMWcMMVbW57cZqxM995TDyPkxE9cPFXpddv14ivLW",
	"u3udYwAJhCWzjSC3riaQk7EYHTAa7jxpm1RrTEJNpmM6y48jyJGkNmEWflZcGQIGjvGNDPhrjb5JsA+X",
	"bBXNYC39HgnYHKI9UCEJdOAGyEHnaEnvNAagaBna87bZCnmPQIUOA6CFSDTBACoGuMzzVK4pVC2geOEU",
	"FrkeWTPg4m5/AZixYCEFIk64203q2mUsPYK8l9rIremV4SyWJiyj06ck6LQNOnjSe8lNU1v9qcab1g/5",
	"Go+Z0CpNV7jfM1BpgS1akB4rj6EDjU8TofOUV0JVsHwYuLmacd7CM82EfOaAXTOUIRz86Wo4aGP4iJ2Z",
	"Bg7SUq8rwfGQsqgY/RHhvfP7BT7IsjK4wAdCG7Szi7J5vZNWNa97v++TeRi5RZNKpahLaV97lHksWbSb",
	"SXrSJcn9k83B5y05wzt57nPohbnz50t+jz4kBWFdIHSHvWyOpK+5FskJefwWSJsPd5gb7lhJequNpgZU",
	"gepjWDwrvToXI+q+JF3rZilRSovCJiGbU4EcYLjtB1QnOjq86DPO7r77Ft5+ffEWbr3bGizhqhNTTjc+",
	"DZe3NyEmbjQ3ab/hwdblTbRC9gxBoZVeuzBT3VaeX9CD+FAZaePCA6FCpXdacaniGDmlah9fpr5KVYs4",
	"Pb42kExuDpfw+90N+P6Z66Z4IuQFjsH/WdCuOgkxKpyn7DzdUJel9M2Z56gXaffxw7/J5xnXZ+3apjHp",
	"tm4ZPF1a50j4awBpzDA/nQxECOl1zIa6xPkWNjjgPy96OJ7aHl8iK2ccjAA2j+yx83FjH11GgX3M9vcN",
	"9+XOcSpGK+w2/qFWXqzEzzdrkQkrS+z+jG0lxt12MtY8nhN1WYk3i4vFBV9zFVpZabESXy3eLC7albEI",
	"YmVrYyLhea0eQ3r8OwAA//96THtJ2Q0AAA==",
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
