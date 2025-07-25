/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.13.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 2.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement

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

type GroupIdentifiersApiService service

/*
GroupIdentifiersApiService Mapping of Group Identifiers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExtGroupId - External Group Identifier
 * @param IntGroupId - Internal Group Identifier
 * @param UeIdInd - Indication whether UE identifiers are required or not
 * @param SupportedFeatures - Supported Features
 * @param AfId - AF identifier
 * @param IfNoneMatch - Validator for conditional requests, as described in RFC 7232, 3.2
 * @param IfModifiedSince - Validator for conditional requests, as described in RFC 7232, 3.3

@return GetGroupIdentifiersResponse
*/

// GetGroupIdentifiersRequest
type GetGroupIdentifiersRequest struct {
	ExtGroupId        *string
	IntGroupId        *string
	UeIdInd           *bool
	SupportedFeatures *string
	AfId              *string
	IfNoneMatch       *string
	IfModifiedSince   *string
}

func (r *GetGroupIdentifiersRequest) SetExtGroupId(ExtGroupId string) {
	r.ExtGroupId = &ExtGroupId
}
func (r *GetGroupIdentifiersRequest) SetIntGroupId(IntGroupId string) {
	r.IntGroupId = &IntGroupId
}
func (r *GetGroupIdentifiersRequest) SetUeIdInd(UeIdInd bool) {
	r.UeIdInd = &UeIdInd
}
func (r *GetGroupIdentifiersRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *GetGroupIdentifiersRequest) SetAfId(AfId string) {
	r.AfId = &AfId
}
func (r *GetGroupIdentifiersRequest) SetIfNoneMatch(IfNoneMatch string) {
	r.IfNoneMatch = &IfNoneMatch
}
func (r *GetGroupIdentifiersRequest) SetIfModifiedSince(IfModifiedSince string) {
	r.IfModifiedSince = &IfModifiedSince
}

type GetGroupIdentifiersResponse struct {
	CacheControl           string
	ETag                   string
	LastModified           string
	UdmSdmGroupIdentifiers models.UdmSdmGroupIdentifiers
}

type GetGroupIdentifiersError struct {
	ProblemDetails models.ProblemDetails
}

func (a *GroupIdentifiersApiService) GetGroupIdentifiers(ctx context.Context, request *GetGroupIdentifiersRequest) (*GetGroupIdentifiersResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetGroupIdentifiersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/group-data/group-identifiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.ExtGroupId != nil {
		localVarQueryParams.Add("ext-group-id", openapi.ParameterToString(request.ExtGroupId, "multi"))
	}
	if request.IntGroupId != nil {
		localVarQueryParams.Add("int-group-id", openapi.ParameterToString(request.IntGroupId, "multi"))
	}
	if request.UeIdInd != nil {
		localVarQueryParams.Add("ue-id-ind", openapi.ParameterToString(request.UeIdInd, "multi"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
	}
	if request.AfId != nil {
		localVarQueryParams.Add("af-id", openapi.ParameterToString(request.AfId, "multi"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	if request.IfNoneMatch != nil {
		localVarHeaderParams["If-None-Match"] = openapi.ParameterToString(request.IfNoneMatch, "csv")
	}

	if request.IfModifiedSince != nil {
		localVarHeaderParams["If-Modified-Since"] = openapi.ParameterToString(request.IfModifiedSince, "csv")
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
		err = openapi.Deserialize(&localVarReturnValue.UdmSdmGroupIdentifiers, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		localVarReturnValue.ETag = localVarHTTPResponse.Header.Get("ETag")
		localVarReturnValue.LastModified = localVarHTTPResponse.Header.Get("Last-Modified")
		return &localVarReturnValue, nil
	case 400:
		var v GetGroupIdentifiersError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetGroupIdentifiersError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetGroupIdentifiersError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetGroupIdentifiersError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetGroupIdentifiersError
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
