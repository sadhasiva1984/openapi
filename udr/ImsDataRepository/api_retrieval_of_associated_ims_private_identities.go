/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.504 V18.3.0; 5G System; Unified Data Repository Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.504/
 *
 * API version: 2.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ImsDataRepository

import (
	"github.com/sadhasiva1984/openapi"
	"github.com/sadhasiva1984/openapi/models"

	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type RetrievalOfAssociatedIMSPrivateIdentitiesApiService service

/*
RetrievalOfAssociatedIMSPrivateIdentitiesApiService Retrieve the associated private identities to the IMS public identity included in the service request
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ImsUeId - IMS Public Identity
 * @param SupportedFeatures - Supported Features
 * @param Impi - Private Identity of type IMPI

@return GetImsPrivateIdsResponse
*/

// GetImsPrivateIdsRequest
type GetImsPrivateIdsRequest struct {
	ImsUeId           *string
	SupportedFeatures *string
	Impi              *string
}

func (r *GetImsPrivateIdsRequest) SetImsUeId(ImsUeId string) {
	r.ImsUeId = &ImsUeId
}
func (r *GetImsPrivateIdsRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}
func (r *GetImsPrivateIdsRequest) SetImpi(Impi string) {
	r.Impi = &Impi
}

func (a *RetrievalOfAssociatedIMSPrivateIdentitiesApiService) GetImsPrivateIds(ctx context.Context, request *GetImsPrivateIdsRequest) (*GetImsPrivateIdsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetImsPrivateIdsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{imsUeId}/identities/private-identities"
	localVarPath = strings.Replace(localVarPath, "{"+"imsUeId"+"}", fmt.Sprintf("%v", request.ImsUeId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, ""))
	}
	if request.Impi != nil {
		localVarQueryParams.Add("impi", openapi.ParameterToString(request.Impi, ""))
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
		err = openapi.Deserialize(&localVarReturnValue.PrivateIdentities, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 307:
		var v GetImsPrivateIdsError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v GetImsPrivateIdsError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetImsPrivateIdsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetImsPrivateIdsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return &localVarReturnValue, apiError
	}
}

type GetImsPrivateIdsResponse struct {
	PrivateIdentities models.PrivateIdentities
}

type GetImsPrivateIdsError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}
