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

type HSSEventSubscriptionInfoDocumentApiService service

/*
HSSEventSubscriptionInfoDocumentApiService Create HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -
 * @param HssSubscriptionInfo -

@return CreateHSSSubscriptionsResponse
*/

// CreateHSSSubscriptionsRequest
type CreateHSSSubscriptionsRequest struct {
	UeId                *string
	SubsId              *string
	HssSubscriptionInfo *models.HssSubscriptionInfo
}

func (r *CreateHSSSubscriptionsRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *CreateHSSSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *CreateHSSSubscriptionsRequest) SetHssSubscriptionInfo(HssSubscriptionInfo models.HssSubscriptionInfo) {
	r.HssSubscriptionInfo = &HssSubscriptionInfo
}

type CreateHSSSubscriptionsResponse struct {
	HssSubscriptionInfo models.HssSubscriptionInfo
}

type CreateHSSSubscriptionsError struct {
}

func (a *HSSEventSubscriptionInfoDocumentApiService) CreateHSSSubscriptions(ctx context.Context, request *CreateHSSSubscriptionsRequest) (*CreateHSSSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateHSSSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
	localVarPostBody = request.HssSubscriptionInfo

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
		err = openapi.Deserialize(&localVarReturnValue.HssSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
HSSEventSubscriptionInfoDocumentApiService Retrieve HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExternalGroupId -
 * @param SubsId -

@return GetHssGroupSubscriptionsResponse
*/

// GetHssGroupSubscriptionsRequest
type GetHssGroupSubscriptionsRequest struct {
	ExternalGroupId *string
	SubsId          *string
}

func (r *GetHssGroupSubscriptionsRequest) SetExternalGroupId(ExternalGroupId string) {
	r.ExternalGroupId = &ExternalGroupId
}
func (r *GetHssGroupSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type GetHssGroupSubscriptionsResponse struct {
	HssSubscriptionInfo models.HssSubscriptionInfo
}

type GetHssGroupSubscriptionsError struct {
}

func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssGroupSubscriptions(ctx context.Context, request *GetHssGroupSubscriptionsRequest) (*GetHssGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetHssGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", openapi.StringOfValue(*request.ExternalGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		err = openapi.Deserialize(&localVarReturnValue.HssSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
HSSEventSubscriptionInfoDocumentApiService Retrieve HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -

@return GetHssSubscriptionInfoResponse
*/

// GetHssSubscriptionInfoRequest
type GetHssSubscriptionInfoRequest struct {
	UeId   *string
	SubsId *string
}

func (r *GetHssSubscriptionInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *GetHssSubscriptionInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type GetHssSubscriptionInfoResponse struct {
	SmfSubscriptionInfo models.SmfSubscriptionInfo
}

type GetHssSubscriptionInfoError struct {
}

func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssSubscriptionInfo(ctx context.Context, request *GetHssSubscriptionInfoRequest) (*GetHssSubscriptionInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetHssSubscriptionInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		err = openapi.Deserialize(&localVarReturnValue.SmfSubscriptionInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	default:
		return nil, apiError
	}
}

/*
HSSEventSubscriptionInfoDocumentApiService Modify HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExternalGroupId -
 * @param SubsId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return ModifyHssGroupSubscriptionsResponse
*/

// ModifyHssGroupSubscriptionsRequest
type ModifyHssGroupSubscriptionsRequest struct {
	ExternalGroupId   *string
	SubsId            *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *ModifyHssGroupSubscriptionsRequest) SetExternalGroupId(ExternalGroupId string) {
	r.ExternalGroupId = &ExternalGroupId
}
func (r *ModifyHssGroupSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *ModifyHssGroupSubscriptionsRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *ModifyHssGroupSubscriptionsRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type ModifyHssGroupSubscriptionsResponse struct {
	PatchResult models.PatchResult
}

type ModifyHssGroupSubscriptionsError struct {
	ProblemDetails models.ProblemDetails
}

func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssGroupSubscriptions(ctx context.Context, request *ModifyHssGroupSubscriptionsRequest) (*ModifyHssGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModifyHssGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", openapi.StringOfValue(*request.ExternalGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		var v ModifyHssGroupSubscriptionsError
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
HSSEventSubscriptionInfoDocumentApiService Modify HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -
 * @param PatchItem -
 * @param SupportedFeatures - Features required to be supported by the target NF

@return ModifyHssSubscriptionInfoResponse
*/

// ModifyHssSubscriptionInfoRequest
type ModifyHssSubscriptionInfoRequest struct {
	UeId              *string
	SubsId            *string
	PatchItem         []models.PatchItem
	SupportedFeatures *string
}

func (r *ModifyHssSubscriptionInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *ModifyHssSubscriptionInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}
func (r *ModifyHssSubscriptionInfoRequest) SetPatchItem(PatchItem []models.PatchItem) {
	r.PatchItem = PatchItem
}
func (r *ModifyHssSubscriptionInfoRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type ModifyHssSubscriptionInfoResponse struct {
	PatchResult models.PatchResult
}

type ModifyHssSubscriptionInfoError struct {
	ProblemDetails models.ProblemDetails
}

func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssSubscriptionInfo(ctx context.Context, request *ModifyHssSubscriptionInfoRequest) (*ModifyHssSubscriptionInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModifyHssSubscriptionInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

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
		var v ModifyHssSubscriptionInfoError
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
HSSEventSubscriptionInfoDocumentApiService Delete HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ExternalGroupId -
 * @param SubsId -

@return RemoveHssGroupSubscriptionsResponse
*/

// RemoveHssGroupSubscriptionsRequest
type RemoveHssGroupSubscriptionsRequest struct {
	ExternalGroupId *string
	SubsId          *string
}

func (r *RemoveHssGroupSubscriptionsRequest) SetExternalGroupId(ExternalGroupId string) {
	r.ExternalGroupId = &ExternalGroupId
}
func (r *RemoveHssGroupSubscriptionsRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type RemoveHssGroupSubscriptionsResponse struct {
}

type RemoveHssGroupSubscriptionsError struct {
}

func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssGroupSubscriptions(ctx context.Context, request *RemoveHssGroupSubscriptionsRequest) (*RemoveHssGroupSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoveHssGroupSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", openapi.StringOfValue(*request.ExternalGroupId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
	default:
		return nil, apiError
	}
}

/*
HSSEventSubscriptionInfoDocumentApiService Delete HSS Subscription Info
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param UeId -
 * @param SubsId -

@return RemoveHssSubscriptionsInfoResponse
*/

// RemoveHssSubscriptionsInfoRequest
type RemoveHssSubscriptionsInfoRequest struct {
	UeId   *string
	SubsId *string
}

func (r *RemoveHssSubscriptionsInfoRequest) SetUeId(UeId string) {
	r.UeId = &UeId
}
func (r *RemoveHssSubscriptionsInfoRequest) SetSubsId(SubsId string) {
	r.SubsId = &SubsId
}

type RemoveHssSubscriptionsInfoResponse struct {
}

type RemoveHssSubscriptionsInfoError struct {
}

func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssSubscriptionsInfo(ctx context.Context, request *RemoveHssSubscriptionsInfoRequest) (*RemoveHssSubscriptionsInfoResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RemoveHssSubscriptionsInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", openapi.StringOfValue(*request.UeId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", openapi.StringOfValue(*request.SubsId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
	default:
		return nil, apiError
	}
}
