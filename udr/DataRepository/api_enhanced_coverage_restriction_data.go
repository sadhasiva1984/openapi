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

type EnhancedCoverageRestrictionDataApiService service

/*
EnhancedCoverageRestrictionDataApiService Retrieves the subscribed enhanced Coverage Restriction Data of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId - UE id
 * @param SupportedFeatures - Supported Features
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return QueryCoverageRestrictionDataResponse
*/

// QueryCoverageRestrictionDataRequest
type QueryCoverageRestrictionDataRequest struct {
	UeId              *string
	SupportedFeatures *string
	IfNoneMatch       *string
	IfModifiedSince   *string
}

func (r *QueryCoverageRestrictionDataRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *QueryCoverageRestrictionDataRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *QueryCoverageRestrictionDataRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *QueryCoverageRestrictionDataRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type QueryCoverageRestrictionDataResponse struct {
	CacheControl                    string
	ETag                            string
	LastModified                    string
	EnhancedCoverageRestrictionData models.EnhancedCoverageRestrictionData
}

type QueryCoverageRestrictionDataError struct {
}

func (a *EnhancedCoverageRestrictionDataApiService) QueryCoverageRestrictionData(ctx context.Context, request *QueryCoverageRestrictionDataRequest) (*QueryCoverageRestrictionDataResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QueryCoverageRestrictionDataResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/coverage-restriction-data"
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

	if request.IfNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "multi")
	}

	if request.IfModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "multi")
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
		err = openapi.Deserialize(&localVarReturnValue.EnhancedCoverageRestrictionData, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}
