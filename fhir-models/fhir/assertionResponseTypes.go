// Copyright 2019 - 2025 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/csotherden/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AssertionResponseTypes is documented here http://hl7.org/fhir/ValueSet/assert-response-code-types
type AssertionResponseTypes int

const (
	AssertionResponseTypesContinue AssertionResponseTypes = iota
	AssertionResponseTypesSwitchingProtocols
	AssertionResponseTypesOkay
	AssertionResponseTypesCreated
	AssertionResponseTypesAccepted
	AssertionResponseTypesNonAuthoritativeInformation
	AssertionResponseTypesNoContent
	AssertionResponseTypesResetContent
	AssertionResponseTypesPartialContent
	AssertionResponseTypesMultipleChoices
	AssertionResponseTypesMovedPermanently
	AssertionResponseTypesFound
	AssertionResponseTypesSeeOther
	AssertionResponseTypesNotModified
	AssertionResponseTypesUseProxy
	AssertionResponseTypesTemporaryRedirect
	AssertionResponseTypesPermanentRedirect
	AssertionResponseTypesBadRequest
	AssertionResponseTypesUnauthorized
	AssertionResponseTypesPaymentRequired
	AssertionResponseTypesForbidden
	AssertionResponseTypesNotFound
	AssertionResponseTypesMethodNotAllowed
	AssertionResponseTypesNotAcceptable
	AssertionResponseTypesProxyAuthenticationRequired
	AssertionResponseTypesRequestTimeout
	AssertionResponseTypesConflict
	AssertionResponseTypesGone
	AssertionResponseTypesLengthRequired
	AssertionResponseTypesPreconditionFailed
	AssertionResponseTypesContentTooLarge
	AssertionResponseTypesUriTooLong
	AssertionResponseTypesUnsupportedMediaType
	AssertionResponseTypesRangeNotSatisfiable
	AssertionResponseTypesExpectationFailed
	AssertionResponseTypesMisdirectedRequest
	AssertionResponseTypesUnprocessableContent
	AssertionResponseTypesUpgradeRequired
	AssertionResponseTypesInternalServerError
	AssertionResponseTypesNotImplemented
	AssertionResponseTypesBadGateway
	AssertionResponseTypesServiceUnavailable
	AssertionResponseTypesGatewayTimeout
	AssertionResponseTypesHttpVersionNotSupported
)

func (code AssertionResponseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AssertionResponseTypes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "continue":
		*code = AssertionResponseTypesContinue
	case "switchingProtocols":
		*code = AssertionResponseTypesSwitchingProtocols
	case "okay":
		*code = AssertionResponseTypesOkay
	case "created":
		*code = AssertionResponseTypesCreated
	case "accepted":
		*code = AssertionResponseTypesAccepted
	case "nonAuthoritativeInformation":
		*code = AssertionResponseTypesNonAuthoritativeInformation
	case "noContent":
		*code = AssertionResponseTypesNoContent
	case "resetContent":
		*code = AssertionResponseTypesResetContent
	case "partialContent":
		*code = AssertionResponseTypesPartialContent
	case "multipleChoices":
		*code = AssertionResponseTypesMultipleChoices
	case "movedPermanently":
		*code = AssertionResponseTypesMovedPermanently
	case "found":
		*code = AssertionResponseTypesFound
	case "seeOther":
		*code = AssertionResponseTypesSeeOther
	case "notModified":
		*code = AssertionResponseTypesNotModified
	case "useProxy":
		*code = AssertionResponseTypesUseProxy
	case "temporaryRedirect":
		*code = AssertionResponseTypesTemporaryRedirect
	case "permanentRedirect":
		*code = AssertionResponseTypesPermanentRedirect
	case "badRequest":
		*code = AssertionResponseTypesBadRequest
	case "unauthorized":
		*code = AssertionResponseTypesUnauthorized
	case "paymentRequired":
		*code = AssertionResponseTypesPaymentRequired
	case "forbidden":
		*code = AssertionResponseTypesForbidden
	case "notFound":
		*code = AssertionResponseTypesNotFound
	case "methodNotAllowed":
		*code = AssertionResponseTypesMethodNotAllowed
	case "notAcceptable":
		*code = AssertionResponseTypesNotAcceptable
	case "proxyAuthenticationRequired":
		*code = AssertionResponseTypesProxyAuthenticationRequired
	case "requestTimeout":
		*code = AssertionResponseTypesRequestTimeout
	case "conflict":
		*code = AssertionResponseTypesConflict
	case "gone":
		*code = AssertionResponseTypesGone
	case "lengthRequired":
		*code = AssertionResponseTypesLengthRequired
	case "preconditionFailed":
		*code = AssertionResponseTypesPreconditionFailed
	case "contentTooLarge":
		*code = AssertionResponseTypesContentTooLarge
	case "uriTooLong":
		*code = AssertionResponseTypesUriTooLong
	case "unsupportedMediaType":
		*code = AssertionResponseTypesUnsupportedMediaType
	case "rangeNotSatisfiable":
		*code = AssertionResponseTypesRangeNotSatisfiable
	case "expectationFailed":
		*code = AssertionResponseTypesExpectationFailed
	case "misdirectedRequest":
		*code = AssertionResponseTypesMisdirectedRequest
	case "unprocessableContent":
		*code = AssertionResponseTypesUnprocessableContent
	case "upgradeRequired":
		*code = AssertionResponseTypesUpgradeRequired
	case "internalServerError":
		*code = AssertionResponseTypesInternalServerError
	case "notImplemented":
		*code = AssertionResponseTypesNotImplemented
	case "badGateway":
		*code = AssertionResponseTypesBadGateway
	case "serviceUnavailable":
		*code = AssertionResponseTypesServiceUnavailable
	case "gatewayTimeout":
		*code = AssertionResponseTypesGatewayTimeout
	case "httpVersionNotSupported":
		*code = AssertionResponseTypesHttpVersionNotSupported
	default:
		return fmt.Errorf("unknown AssertionResponseTypes code `%s`", s)
	}
	return nil
}
func (code AssertionResponseTypes) String() string {
	return code.Code()
}
func (code AssertionResponseTypes) Code() string {
	switch code {
	case AssertionResponseTypesContinue:
		return "continue"
	case AssertionResponseTypesSwitchingProtocols:
		return "switchingProtocols"
	case AssertionResponseTypesOkay:
		return "okay"
	case AssertionResponseTypesCreated:
		return "created"
	case AssertionResponseTypesAccepted:
		return "accepted"
	case AssertionResponseTypesNonAuthoritativeInformation:
		return "nonAuthoritativeInformation"
	case AssertionResponseTypesNoContent:
		return "noContent"
	case AssertionResponseTypesResetContent:
		return "resetContent"
	case AssertionResponseTypesPartialContent:
		return "partialContent"
	case AssertionResponseTypesMultipleChoices:
		return "multipleChoices"
	case AssertionResponseTypesMovedPermanently:
		return "movedPermanently"
	case AssertionResponseTypesFound:
		return "found"
	case AssertionResponseTypesSeeOther:
		return "seeOther"
	case AssertionResponseTypesNotModified:
		return "notModified"
	case AssertionResponseTypesUseProxy:
		return "useProxy"
	case AssertionResponseTypesTemporaryRedirect:
		return "temporaryRedirect"
	case AssertionResponseTypesPermanentRedirect:
		return "permanentRedirect"
	case AssertionResponseTypesBadRequest:
		return "badRequest"
	case AssertionResponseTypesUnauthorized:
		return "unauthorized"
	case AssertionResponseTypesPaymentRequired:
		return "paymentRequired"
	case AssertionResponseTypesForbidden:
		return "forbidden"
	case AssertionResponseTypesNotFound:
		return "notFound"
	case AssertionResponseTypesMethodNotAllowed:
		return "methodNotAllowed"
	case AssertionResponseTypesNotAcceptable:
		return "notAcceptable"
	case AssertionResponseTypesProxyAuthenticationRequired:
		return "proxyAuthenticationRequired"
	case AssertionResponseTypesRequestTimeout:
		return "requestTimeout"
	case AssertionResponseTypesConflict:
		return "conflict"
	case AssertionResponseTypesGone:
		return "gone"
	case AssertionResponseTypesLengthRequired:
		return "lengthRequired"
	case AssertionResponseTypesPreconditionFailed:
		return "preconditionFailed"
	case AssertionResponseTypesContentTooLarge:
		return "contentTooLarge"
	case AssertionResponseTypesUriTooLong:
		return "uriTooLong"
	case AssertionResponseTypesUnsupportedMediaType:
		return "unsupportedMediaType"
	case AssertionResponseTypesRangeNotSatisfiable:
		return "rangeNotSatisfiable"
	case AssertionResponseTypesExpectationFailed:
		return "expectationFailed"
	case AssertionResponseTypesMisdirectedRequest:
		return "misdirectedRequest"
	case AssertionResponseTypesUnprocessableContent:
		return "unprocessableContent"
	case AssertionResponseTypesUpgradeRequired:
		return "upgradeRequired"
	case AssertionResponseTypesInternalServerError:
		return "internalServerError"
	case AssertionResponseTypesNotImplemented:
		return "notImplemented"
	case AssertionResponseTypesBadGateway:
		return "badGateway"
	case AssertionResponseTypesServiceUnavailable:
		return "serviceUnavailable"
	case AssertionResponseTypesGatewayTimeout:
		return "gatewayTimeout"
	case AssertionResponseTypesHttpVersionNotSupported:
		return "httpVersionNotSupported"
	}
	return "<unknown>"
}
func (code AssertionResponseTypes) Display() string {
	switch code {
	case AssertionResponseTypesContinue:
		return "Continue"
	case AssertionResponseTypesSwitchingProtocols:
		return "Switching Protocols"
	case AssertionResponseTypesOkay:
		return "OK"
	case AssertionResponseTypesCreated:
		return "Created"
	case AssertionResponseTypesAccepted:
		return "Accepted"
	case AssertionResponseTypesNonAuthoritativeInformation:
		return "Non-Authoritative Information"
	case AssertionResponseTypesNoContent:
		return "No Content"
	case AssertionResponseTypesResetContent:
		return "Reset Content"
	case AssertionResponseTypesPartialContent:
		return "Partial Content"
	case AssertionResponseTypesMultipleChoices:
		return "Multiple Choices"
	case AssertionResponseTypesMovedPermanently:
		return "Moved Permanently"
	case AssertionResponseTypesFound:
		return "Found"
	case AssertionResponseTypesSeeOther:
		return "See Other"
	case AssertionResponseTypesNotModified:
		return "Not Modified"
	case AssertionResponseTypesUseProxy:
		return "Use Proxy"
	case AssertionResponseTypesTemporaryRedirect:
		return "Temporary Redirect"
	case AssertionResponseTypesPermanentRedirect:
		return "Permanent Redirect"
	case AssertionResponseTypesBadRequest:
		return "Bad Request"
	case AssertionResponseTypesUnauthorized:
		return "Unauthorized"
	case AssertionResponseTypesPaymentRequired:
		return "Payment Required"
	case AssertionResponseTypesForbidden:
		return "Forbidden"
	case AssertionResponseTypesNotFound:
		return "Not Found"
	case AssertionResponseTypesMethodNotAllowed:
		return "Method Not Allowed"
	case AssertionResponseTypesNotAcceptable:
		return "Not Acceptable"
	case AssertionResponseTypesProxyAuthenticationRequired:
		return "Proxy Authentication Required"
	case AssertionResponseTypesRequestTimeout:
		return "Request Timeout"
	case AssertionResponseTypesConflict:
		return "Conflict"
	case AssertionResponseTypesGone:
		return "Gone"
	case AssertionResponseTypesLengthRequired:
		return "Length Required"
	case AssertionResponseTypesPreconditionFailed:
		return "Precondition Failed"
	case AssertionResponseTypesContentTooLarge:
		return "Content Too Large"
	case AssertionResponseTypesUriTooLong:
		return "URI Too Long"
	case AssertionResponseTypesUnsupportedMediaType:
		return "Unsupported Media Type"
	case AssertionResponseTypesRangeNotSatisfiable:
		return "Range Not Satisfiable"
	case AssertionResponseTypesExpectationFailed:
		return "Expectation Failed"
	case AssertionResponseTypesMisdirectedRequest:
		return "Misdirected Request"
	case AssertionResponseTypesUnprocessableContent:
		return "Unprocessable Content"
	case AssertionResponseTypesUpgradeRequired:
		return "Upgrade Required"
	case AssertionResponseTypesInternalServerError:
		return "Internal Server Error"
	case AssertionResponseTypesNotImplemented:
		return "Not Implemented"
	case AssertionResponseTypesBadGateway:
		return "Bad Gateway"
	case AssertionResponseTypesServiceUnavailable:
		return "Service Unavailable"
	case AssertionResponseTypesGatewayTimeout:
		return "Gateway Timeout"
	case AssertionResponseTypesHttpVersionNotSupported:
		return "HTTP Version Not Supported"
	}
	return "<unknown>"
}
func (code AssertionResponseTypes) Definition() string {
	switch code {
	case AssertionResponseTypesContinue:
		return "Response code is 100."
	case AssertionResponseTypesSwitchingProtocols:
		return "Response code is 101."
	case AssertionResponseTypesOkay:
		return "Response code is 200."
	case AssertionResponseTypesCreated:
		return "Response code is 201."
	case AssertionResponseTypesAccepted:
		return "Response code is 202."
	case AssertionResponseTypesNonAuthoritativeInformation:
		return "Response code is 203."
	case AssertionResponseTypesNoContent:
		return "Response code is 204."
	case AssertionResponseTypesResetContent:
		return "Response code is 205."
	case AssertionResponseTypesPartialContent:
		return "Response code is 206."
	case AssertionResponseTypesMultipleChoices:
		return "Response code is 300."
	case AssertionResponseTypesMovedPermanently:
		return "Response code is 301."
	case AssertionResponseTypesFound:
		return "Response code is 302."
	case AssertionResponseTypesSeeOther:
		return "Response code is 303."
	case AssertionResponseTypesNotModified:
		return "Response code is 304."
	case AssertionResponseTypesUseProxy:
		return "Response code is 305."
	case AssertionResponseTypesTemporaryRedirect:
		return "Response code is 307."
	case AssertionResponseTypesPermanentRedirect:
		return "Response code is 308."
	case AssertionResponseTypesBadRequest:
		return "Response code is 400."
	case AssertionResponseTypesUnauthorized:
		return "Response code is 401."
	case AssertionResponseTypesPaymentRequired:
		return "Response code is 402."
	case AssertionResponseTypesForbidden:
		return "Response code is 403."
	case AssertionResponseTypesNotFound:
		return "Response code is 404."
	case AssertionResponseTypesMethodNotAllowed:
		return "Response code is 405."
	case AssertionResponseTypesNotAcceptable:
		return "Response code is 406."
	case AssertionResponseTypesProxyAuthenticationRequired:
		return "Response code is 407."
	case AssertionResponseTypesRequestTimeout:
		return "Response code is 408."
	case AssertionResponseTypesConflict:
		return "Response code is 409."
	case AssertionResponseTypesGone:
		return "Response code is 410."
	case AssertionResponseTypesLengthRequired:
		return "Response code is 411."
	case AssertionResponseTypesPreconditionFailed:
		return "Response code is 412."
	case AssertionResponseTypesContentTooLarge:
		return "Response code is 413."
	case AssertionResponseTypesUriTooLong:
		return "Response code is 414."
	case AssertionResponseTypesUnsupportedMediaType:
		return "Response code is 415."
	case AssertionResponseTypesRangeNotSatisfiable:
		return "Response code is 416."
	case AssertionResponseTypesExpectationFailed:
		return "Response code is 417."
	case AssertionResponseTypesMisdirectedRequest:
		return "Response code is 421."
	case AssertionResponseTypesUnprocessableContent:
		return "Response code is 422."
	case AssertionResponseTypesUpgradeRequired:
		return "Response code is 426."
	case AssertionResponseTypesInternalServerError:
		return "Response code is 500."
	case AssertionResponseTypesNotImplemented:
		return "Response code is 501."
	case AssertionResponseTypesBadGateway:
		return "Response code is 502."
	case AssertionResponseTypesServiceUnavailable:
		return "Response code is 503."
	case AssertionResponseTypesGatewayTimeout:
		return "Response code is 504."
	case AssertionResponseTypesHttpVersionNotSupported:
		return "Response code is 505."
	}
	return "<unknown>"
}
