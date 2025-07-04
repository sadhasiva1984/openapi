/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.531 V17.7.0; 5G System; Network Slice Selection Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.531/
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NSSelection

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

type NetworkSliceInformationDocumentApiService service

/*
NetworkSliceInformationDocumentApiService Retrieve the Network Slice Selection Information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param NfType - NF type of the NF service consumer
 * @param NfId - NF Instance ID of the NF service consumer
 * @param SliceInfoRequestForRegistration - Requested network slice information during Registration procedure
 * @param SliceInfoRequestForPduSession - Requested network slice information during PDU session establishment procedure
 * @param SliceInfoRequestForUeCu - Requested network slice information during UE confuguration update procedure
 * @param HomePlmnId - PLMN ID of the HPLMN
 * @param Tai - TAI of the UE
 * @param SupportedFeatures - Features required to be supported by the NFs in the target slice instance

@return NSSelectionGetResponse
*/

// NSSelectionGetRequest
type NSSelectionGetRequest struct {
	NfType                          *models.NrfNfManagementNfType
	NfId                            *string
	SliceInfoRequestForRegistration *models.SliceInfoForRegistration
	SliceInfoRequestForPduSession   *models.SliceInfoForPduSession
	SliceInfoRequestForUeCu         *models.SliceInfoForUeConfigurationUpdate
	HomePlmnId                      *models.PlmnId
	Tai                             *models.Tai
	SupportedFeatures               *string
}

func (r *NSSelectionGetRequest) SetNfType(NfType models.NrfNfManagementNfType) {
	r.NfType = &NfType
}
func (r *NSSelectionGetRequest) SetNfId(NfId string) {
	r.NfId = &NfId
}
func (r *NSSelectionGetRequest) SetSliceInfoRequestForRegistration(SliceInfoRequestForRegistration models.SliceInfoForRegistration) {
	r.SliceInfoRequestForRegistration = &SliceInfoRequestForRegistration
}
func (r *NSSelectionGetRequest) SetSliceInfoRequestForPduSession(SliceInfoRequestForPduSession models.SliceInfoForPduSession) {
	r.SliceInfoRequestForPduSession = &SliceInfoRequestForPduSession
}
func (r *NSSelectionGetRequest) SetSliceInfoRequestForUeCu(SliceInfoRequestForUeCu models.SliceInfoForUeConfigurationUpdate) {
	r.SliceInfoRequestForUeCu = &SliceInfoRequestForUeCu
}
func (r *NSSelectionGetRequest) SetHomePlmnId(HomePlmnId models.PlmnId) {
	r.HomePlmnId = &HomePlmnId
}
func (r *NSSelectionGetRequest) SetTai(Tai models.Tai) {
	r.Tai = &Tai
}
func (r *NSSelectionGetRequest) SetSupportedFeatures(SupportedFeatures string) {
	r.SupportedFeatures = &SupportedFeatures
}

type NSSelectionGetResponse struct {
	AuthorizedNetworkSliceInfo models.AuthorizedNetworkSliceInfo
}

type NSSelectionGetError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

func (a *NetworkSliceInformationDocumentApiService) NSSelectionGet(ctx context.Context, request *NSSelectionGetRequest) (*NSSelectionGetResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NSSelectionGetResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/network-slice-information"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.NfType == nil {
		return nil, openapi.ReportError("NfType must be non nil")
	} else {
		localVarQueryParams.Add("nf-type", openapi.ParameterToString(request.NfType, "multi"))
	}
	if request.NfId == nil {
		return nil, openapi.ReportError("NfId must be non nil")
	} else {
		localVarQueryParams.Add("nf-id", openapi.ParameterToString(request.NfId, "multi"))
	}
	if request.SliceInfoRequestForRegistration != nil {
		localVarQueryParams.Add("slice-info-request-for-registration", openapi.ParameterToString(request.SliceInfoRequestForRegistration, "application/json"))
	}
	if request.SliceInfoRequestForPduSession != nil {
		localVarQueryParams.Add("slice-info-request-for-pdu-session", openapi.ParameterToString(request.SliceInfoRequestForPduSession, "application/json"))
	}
	if request.SliceInfoRequestForUeCu != nil {
		localVarQueryParams.Add("slice-info-request-for-ue-cu", openapi.ParameterToString(request.SliceInfoRequestForUeCu, "application/json"))
	}
	if request.HomePlmnId != nil {
		localVarQueryParams.Add("home-plmn-id", openapi.ParameterToString(request.HomePlmnId, "application/json"))
	}
	if request.Tai != nil {
		localVarQueryParams.Add("tai", openapi.ParameterToString(request.Tai, "application/json"))
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
		err = openapi.Deserialize(&localVarReturnValue.AuthorizedNetworkSliceInfo, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 307:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 308:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v NSSelectionGetError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v NSSelectionGetError
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
