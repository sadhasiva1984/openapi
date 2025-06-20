/*
 * Nudm_EE
 *
 * Nudm Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.10.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package EventExposure

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

type DeleteEESubscriptionApiService service

/*
DeleteEESubscriptionApiService Unsubscribe
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeIdentity - Represents the scope of the UE for which the subscription is applied. Contains the GPSI of the user or the external group ID or any UE.
 * @param SubscriptionId - Id of the EE Subscription

@return DeleteEeSubscriptionResponse
*/

// DeleteEeSubscriptionRequest
type DeleteEeSubscriptionRequest struct {
	UeIdentity     *string
	SubscriptionId *string
}

func (r *DeleteEeSubscriptionRequest) SetUeIdentity(UeIdentity string) {
	r.UeIdentity = &UeIdentity
}
func (r *DeleteEeSubscriptionRequest) SetSubscriptionId(SubscriptionId string) {
	r.SubscriptionId = &SubscriptionId
}

type DeleteEeSubscriptionResponse struct {
}

type DeleteEeSubscriptionError struct {
	ProblemDetails models.ProblemDetails
}

func (a *DeleteEESubscriptionApiService) DeleteEeSubscription(ctx context.Context, request *DeleteEeSubscriptionRequest) (*DeleteEeSubscriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeleteEeSubscriptionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{ueIdentity}/ee-subscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueIdentity"+"}", openapi.StringOfValue(*request.UeIdentity), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", openapi.StringOfValue(*request.SubscriptionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v DeleteEeSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v DeleteEeSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v DeleteEeSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v DeleteEeSubscriptionError
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
