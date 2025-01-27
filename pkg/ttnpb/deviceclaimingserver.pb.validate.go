// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _deviceclaimingserver_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on ClaimEndDeviceRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClaimEndDeviceRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ClaimEndDeviceRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "target_application_ids":

			if v, ok := interface{}(&m.TargetApplicationIDs).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ClaimEndDeviceRequestValidationError{
						field:  "target_application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "target_device_id":

			if utf8.RuneCountInString(m.GetTargetDeviceID()) > 36 {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_device_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ClaimEndDeviceRequest_TargetDeviceID_Pattern.MatchString(m.GetTargetDeviceID()) {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_device_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "target_network_server_address":

			if !_ClaimEndDeviceRequest_TargetNetworkServerAddress_Pattern.MatchString(m.GetTargetNetworkServerAddress()) {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_network_server_address",
					reason: "value does not match regex pattern \"^(?:(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\\\-]*[a-zA-Z0-9])\\\\.)*(?:[A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\\\-]*[A-Za-z0-9])(?::[0-9]{1,5})?$|^$\"",
				}
			}

		case "target_network_server_kek_label":

			if utf8.RuneCountInString(m.GetTargetNetworkServerKEKLabel()) > 2048 {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_network_server_kek_label",
					reason: "value length must be at most 2048 runes",
				}
			}

		case "target_application_server_address":

			if !_ClaimEndDeviceRequest_TargetApplicationServerAddress_Pattern.MatchString(m.GetTargetApplicationServerAddress()) {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_application_server_address",
					reason: "value does not match regex pattern \"^(?:(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\\\-]*[a-zA-Z0-9])\\\\.)*(?:[A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\\\-]*[A-Za-z0-9])(?::[0-9]{1,5})?$|^$\"",
				}
			}

		case "target_application_server_kek_label":

			if utf8.RuneCountInString(m.GetTargetApplicationServerKEKLabel()) > 2048 {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_application_server_kek_label",
					reason: "value length must be at most 2048 runes",
				}
			}

		case "target_application_server_id":

			if utf8.RuneCountInString(m.GetTargetApplicationServerID()) > 100 {
				return ClaimEndDeviceRequestValidationError{
					field:  "target_application_server_id",
					reason: "value length must be at most 100 runes",
				}
			}

		case "target_net_id":
			// no validation rules for TargetNetID
		case "invalidate_authentication_code":
			// no validation rules for InvalidateAuthenticationCode
		case "source_device":
			if len(subs) == 0 {
				subs = []string{
					"authenticated_identifiers", "qr_code",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "authenticated_identifiers":

					if v, ok := interface{}(m.GetAuthenticatedIdentifiers()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ClaimEndDeviceRequestValidationError{
								field:  "authenticated_identifiers",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "qr_code":

					if l := len(m.GetQRCode()); l < 0 || l > 1024 {
						return ClaimEndDeviceRequestValidationError{
							field:  "qr_code",
							reason: "value length must be between 0 and 1024 bytes, inclusive",
						}
					}

				default:
					return ClaimEndDeviceRequestValidationError{
						field:  "source_device",
						reason: "value is required",
					}
				}
			}
		default:
			return ClaimEndDeviceRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ClaimEndDeviceRequestValidationError is the validation error returned by
// ClaimEndDeviceRequest.ValidateFields if the designated constraints aren't met.
type ClaimEndDeviceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClaimEndDeviceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClaimEndDeviceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClaimEndDeviceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClaimEndDeviceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClaimEndDeviceRequestValidationError) ErrorName() string {
	return "ClaimEndDeviceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ClaimEndDeviceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClaimEndDeviceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClaimEndDeviceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClaimEndDeviceRequestValidationError{}

var _ClaimEndDeviceRequest_TargetDeviceID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

var _ClaimEndDeviceRequest_TargetNetworkServerAddress_Pattern = regexp.MustCompile("^(?:(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*(?:[A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])(?::[0-9]{1,5})?$|^$")

var _ClaimEndDeviceRequest_TargetApplicationServerAddress_Pattern = regexp.MustCompile("^(?:(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*(?:[A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])(?::[0-9]{1,5})?$|^$")

// ValidateFields checks the field values on AuthorizeApplicationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *AuthorizeApplicationRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = AuthorizeApplicationRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_ids":

			if v, ok := interface{}(&m.ApplicationIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return AuthorizeApplicationRequestValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "api_key":

			if utf8.RuneCountInString(m.GetAPIKey()) < 1 {
				return AuthorizeApplicationRequestValidationError{
					field:  "api_key",
					reason: "value length must be at least 1 runes",
				}
			}

		default:
			return AuthorizeApplicationRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// AuthorizeApplicationRequestValidationError is the validation error returned
// by AuthorizeApplicationRequest.ValidateFields if the designated constraints
// aren't met.
type AuthorizeApplicationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizeApplicationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizeApplicationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizeApplicationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizeApplicationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizeApplicationRequestValidationError) ErrorName() string {
	return "AuthorizeApplicationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizeApplicationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizeApplicationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizeApplicationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizeApplicationRequestValidationError{}

// ValidateFields checks the field values on
// ClaimEndDeviceRequest_AuthenticatedIdentifiers with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClaimEndDeviceRequest_AuthenticatedIdentifiers) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ClaimEndDeviceRequest_AuthenticatedIdentifiersFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":
			// no validation rules for JoinEUI
		case "dev_eui":
			// no validation rules for DevEUI
		case "authentication_code":

			if !_ClaimEndDeviceRequest_AuthenticatedIdentifiers_AuthenticationCode_Pattern.MatchString(m.GetAuthenticationCode()) {
				return ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError{
					field:  "authentication_code",
					reason: "value does not match regex pattern \"^[A-Z0-9]{1,32}$\"",
				}
			}

		default:
			return ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError is the
// validation error returned by
// ClaimEndDeviceRequest_AuthenticatedIdentifiers.ValidateFields if the
// designated constraints aren't met.
type ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) ErrorName() string {
	return "ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError"
}

// Error satisfies the builtin error interface
func (e ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClaimEndDeviceRequest_AuthenticatedIdentifiers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClaimEndDeviceRequest_AuthenticatedIdentifiersValidationError{}

var _ClaimEndDeviceRequest_AuthenticatedIdentifiers_AuthenticationCode_Pattern = regexp.MustCompile("^[A-Z0-9]{1,32}$")
