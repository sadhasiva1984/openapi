/*
 * Nudm_ReportSMDeliveryStatus
 *
 * UDM Report SM Delivery Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.503 Unified Data Management Services, version 17.8.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ReportSMDeliveryStatus

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

type ReportSMDeliveryStatusApiService service

/*
ReportSMDeliveryStatusApiService Report the SM Delivery Status
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeIdentity - Represents the scope of the UE for which the Service Specific Parameters are authorized. Contains the GPSI of the user or the external group ID.
 * @param SmDeliveryStatus -

@return ReportSMDeliveryStatusResponse
*/

// ReportSMDeliveryStatusRequest
type ReportSMDeliveryStatusRequest struct {
	UeIdentity       *string
	SmDeliveryStatus *models.SmDeliveryStatus
}

func (r *ReportSMDeliveryStatusRequest) SetUeIdentity(UeIdentity string) {
	r.UeIdentity = &UeIdentity
}
func (r *ReportSMDeliveryStatusRequest) SetSmDeliveryStatus(SmDeliveryStatus models.SmDeliveryStatus) {
	r.SmDeliveryStatus = &SmDeliveryStatus
}

type ReportSMDeliveryStatusResponse struct {
}

type ReportSMDeliveryStatusError struct {
	ProblemDetails models.ProblemDetails
}

func (a *ReportSMDeliveryStatusApiService) ReportSMDeliveryStatus(ctx context.Context, request *ReportSMDeliveryStatusRequest) (*ReportSMDeliveryStatusResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReportSMDeliveryStatusResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{ueIdentity}/sm-delivery-status"
	localVarPath = strings.Replace(localVarPath, "{"+"ueIdentity"+"}", openapi.StringOfValue(*request.UeIdentity), -1)

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

	// body params
	localVarPostBody = request.SmDeliveryStatus

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
		var v ReportSMDeliveryStatusError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v ReportSMDeliveryStatusError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v ReportSMDeliveryStatusError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v ReportSMDeliveryStatusError
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
