/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V17.12.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/sadhasiva1984/openapi"
	"github.com/sadhasiva1984/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AMFNon3GPPAccessRegistrationDocumentApiService service

/*
AMFNon3GPPAccessRegistrationDocumentApiService To modify the AMF context data of a UE using non 3gpp access in the UDR
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return AmfContextNon3gppResponse
*/

// AmfContextNon3gppRequest
type AmfContextNon3gppRequest struct {
	UeId              *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *AmfContextNon3gppRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *AmfContextNon3gppRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *AmfContextNon3gppRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type AmfContextNon3gppResponse struct {
	PatchResult models.PatchResult
}

type AmfContextNon3gppError struct {
	ProblemDetails models.ProblemDetails
}

func (a *AMFNon3GPPAccessRegistrationDocumentApiService) AmfContextNon3gpp(ctx context.Context, request *AmfContextNon3gppRequest) (*AmfContextNon3gppResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AmfContextNon3gppResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/amf-non-3gpp-access"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.PatchItem

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PatchResult, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 403:
		var v AmfContextNon3gppError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}

/*
AMFNon3GPPAccessRegistrationDocumentApiService To store the AMF context data of a UE using non-3gpp access in the UDR
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param AmfNon3GppAccessRegistration -

@return CreateAmfContextNon3gppResponse
*/

// CreateAmfContextNon3gppRequest
type CreateAmfContextNon3gppRequest struct {
	UeId                         *string
	AmfNon3GppAccessRegistration *models.AmfNon3GppAccessRegistration
}

func (r *CreateAmfContextNon3gppRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateAmfContextNon3gppRequest) SetAmfNon3GppAccessRegistration(AmfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration) {
	r.AmfNon3GppAccessRegistration = &AmfNon3GppAccessRegistration
}

type CreateAmfContextNon3gppResponse struct {
	Location                  string
	Amf3GppAccessRegistration models.Amf3GppAccessRegistration
}

type CreateAmfContextNon3gppError struct {
}

func (a *AMFNon3GPPAccessRegistrationDocumentApiService) CreateAmfContextNon3gpp(ctx context.Context, request *CreateAmfContextNon3gppRequest) (*CreateAmfContextNon3gppResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateAmfContextNon3gppResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/amf-non-3gpp-access"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.AmfNon3GppAccessRegistration

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return &localVarReturnValue, nil
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.Amf3GppAccessRegistration, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
AMFNon3GPPAccessRegistrationDocumentApiService Retrieves the AMF context data of a UE using non-3gpp access
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param Fields - attributes to be retrieved
 * @param SupportedFeatures - Supported Features

@return QueryAmfContextNon3gppResponse
*/

// QueryAmfContextNon3gppRequest
type QueryAmfContextNon3gppRequest struct {
	UeId              *string
	Fields            []string
	SupportedFeatures *string
}

func (r *QueryAmfContextNon3gppRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *QueryAmfContextNon3gppRequest) SetFields(Fields []string) {
	r.Fields = Fields
}
func (r *QueryAmfContextNon3gppRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type QueryAmfContextNon3gppResponse struct {
	AmfNon3GppAccessRegistration models.AmfNon3GppAccessRegistration
}

type QueryAmfContextNon3gppError struct {
}

func (a *AMFNon3GPPAccessRegistrationDocumentApiService) QueryAmfContextNon3gpp(ctx context.Context, request *QueryAmfContextNon3gppRequest) (*QueryAmfContextNon3gppResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QueryAmfContextNon3gppResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/amf-non-3gpp-access"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Fields != nil {
		if len(request.Fields) < 1 {
			return &localVarReturnValue, openapi.ReportError("Fields must have at least 1 elements")
		}
		localVarQueryParams.Add("fields", openapi.ParameterToString(request.Fields, "csv"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.AmfNon3GppAccessRegistration, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
