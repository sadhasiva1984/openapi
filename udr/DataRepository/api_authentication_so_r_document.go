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

type AuthenticationSoRDocumentApiService service

/*
AuthenticationSoRDocumentApiService To store the SoR acknowledgement information of a UE and ME support of SOR CMCI
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param SorData -
 * @param SupportedFeatures - Supported Features

@return CreateAuthenticationSoRResponse
*/

// CreateAuthenticationSoRRequest
type CreateAuthenticationSoRRequest struct {
	UeId              *string
	SorData           *models.SorData
	SupportedFeatures *string
}

func (r *CreateAuthenticationSoRRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateAuthenticationSoRRequest) SetSorData(SorData models.SorData) {
	r.SorData = &SorData
}
func (r *CreateAuthenticationSoRRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type CreateAuthenticationSoRResponse struct {
}

type CreateAuthenticationSoRError struct {
}

func (a *AuthenticationSoRDocumentApiService) CreateAuthenticationSoR(ctx context.Context, request *CreateAuthenticationSoRRequest) (*CreateAuthenticationSoRResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateAuthenticationSoRResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.SorData

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
	default:
		return nil, apiError
	}
}

/*
AuthenticationSoRDocumentApiService Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param SupportedFeatures - Supported Features

@return QueryAuthSoRResponse
*/

// QueryAuthSoRRequest
type QueryAuthSoRRequest struct {
	UeId              *string
	SupportedFeatures *string
}

func (r *QueryAuthSoRRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *QueryAuthSoRRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type QueryAuthSoRResponse struct {
	SorData models.SorData
}

type QueryAuthSoRError struct {
}

func (a *AuthenticationSoRDocumentApiService) QueryAuthSoR(ctx context.Context, request *QueryAuthSoRRequest) (*QueryAuthSoRResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QueryAuthSoRResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		err = openapi.Deserialize(&localVarReturnValue.SorData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
AuthenticationSoRDocumentApiService Updates the ME support of SOR CMCI information of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return UpdateAuthenticationSoRResponse
*/

// UpdateAuthenticationSoRRequest
type UpdateAuthenticationSoRRequest struct {
	UeId              *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *UpdateAuthenticationSoRRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *UpdateAuthenticationSoRRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *UpdateAuthenticationSoRRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type UpdateAuthenticationSoRResponse struct {
	PatchResult models.PatchResult
}

type UpdateAuthenticationSoRError struct {
	ProblemDetails models.ProblemDetails
}

func (a *AuthenticationSoRDocumentApiService) UpdateAuthenticationSoR(ctx context.Context, request *UpdateAuthenticationSoRRequest) (*UpdateAuthenticationSoRResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateAuthenticationSoRResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
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
		var v UpdateAuthenticationSoRError
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
