/*
 * Nudsf_DataRepository
 *
 * Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.598 UDSF Services, V17.6.0.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.598/
 *
 * API version: 1.1.0
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

type NotificationSubscriptionsCRUDApiService service

/*
NotificationSubscriptionsCRUDApiService Notification subscription retrieval
retrieve all notification subscriptions of the storage
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param RealmId - Identifier of the Realm
 * @param StorageId - Identifier of the Storage
 * @param LimitRange - The maximum number of NotificationSubscriptions to fetch
 * @param SupportedFeatures - Features required to be supported by the target NF

@return GetNotificationSubscriptionsResponse
*/

// GetNotificationSubscriptionsRequest
type GetNotificationSubscriptionsRequest struct {
	RealmId           *string
	StorageId         *string
	LimitRange        *int32
	SupportedFeatures *string
}

func (r *GetNotificationSubscriptionsRequest) SetRealmId(RealmId string) {
	r.RealmId = &RealmId
}
func (r *GetNotificationSubscriptionsRequest) SetStorageId(StorageId string) {
	r.StorageId = &StorageId
}
func (r *GetNotificationSubscriptionsRequest) SetLimitRange(LimitRange int32) {
	r.LimitRange = &LimitRange
}
func (r *GetNotificationSubscriptionsRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type GetNotificationSubscriptionsResponse struct {
	NotificationSubscription []models.NotificationSubscription
}

type GetNotificationSubscriptionsError struct {
	CacheControl   string
	ETag           string
	RetryAfter     interface{}
	ProblemDetails models.ProblemDetails
}

func (a *NotificationSubscriptionsCRUDApiService) GetNotificationSubscriptions(ctx context.Context, request *GetNotificationSubscriptionsRequest) (*GetNotificationSubscriptionsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetNotificationSubscriptionsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/{realmId}/{storageId}/subs-to-notify"
	localVarPath = strings.Replace(localVarPath, "{"+"realmId"+"}", openapi.StringOfValue(*request.RealmId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"storageId"+"}", openapi.StringOfValue(*request.StorageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.LimitRange != nil {
		localVarQueryParams.Add("limit-range", openapi.ParameterToString(request.LimitRange, "multi"))
	}
	if request.SupportedFeatures != nil {
		localVarQueryParams.Add("supported-features", openapi.ParameterToString(request.SupportedFeatures, "multi"))
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
		err = openapi.Deserialize(&localVarReturnValue.NotificationSubscription, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 304:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.CacheControl = localVarHTTPResponse.Header.Get("Cache-Control")
		v.ETag = localVarHTTPResponse.Header.Get("ETag")
		v.RetryAfter = localVarHTTPResponse.Header.Get("Retry-After")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetNotificationSubscriptionsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetNotificationSubscriptionsError
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
