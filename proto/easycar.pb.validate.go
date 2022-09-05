// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/easycar.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReqMultiError, or
// nil if none found.
func (m *RegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := RegisterReqValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetBranches() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterReqValidationError{
						field:  fmt.Sprintf("Branches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterReqValidationError{
						field:  fmt.Sprintf("Branches[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterReqValidationError{
					field:  fmt.Sprintf("Branches[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegisterReqMultiError(errors)
	}

	return nil
}

// RegisterReqMultiError is an error wrapping multiple validation errors
// returned by RegisterReq.ValidateAll() if the designated constraints aren't met.
type RegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReqMultiError) AllErrors() []error { return m }

// RegisterReqValidationError is the validation error returned by
// RegisterReq.Validate if the designated constraints aren't met.
type RegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReqValidationError) ErrorName() string { return "RegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReqValidationError{}

// Validate checks the field values on RegisterResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterRespMultiError, or
// nil if none found.
func (m *RegisterResp) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterRespMultiError(errors)
	}

	return nil
}

// RegisterRespMultiError is an error wrapping multiple validation errors
// returned by RegisterResp.ValidateAll() if the designated constraints aren't met.
type RegisterRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRespMultiError) AllErrors() []error { return m }

// RegisterRespValidationError is the validation error returned by
// RegisterResp.Validate if the designated constraints aren't met.
type RegisterRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRespValidationError) ErrorName() string { return "RegisterRespValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRespValidationError{}

// Validate checks the field values on BeginResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BeginResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BeginResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BeginRespMultiError, or nil
// if none found.
func (m *BeginResp) ValidateAll() error {
	return m.validate(true)
}

func (m *BeginResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GId

	if len(errors) > 0 {
		return BeginRespMultiError(errors)
	}

	return nil
}

// BeginRespMultiError is an error wrapping multiple validation errors returned
// by BeginResp.ValidateAll() if the designated constraints aren't met.
type BeginRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BeginRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BeginRespMultiError) AllErrors() []error { return m }

// BeginRespValidationError is the validation error returned by
// BeginResp.Validate if the designated constraints aren't met.
type BeginRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BeginRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BeginRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BeginRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BeginRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BeginRespValidationError) ErrorName() string { return "BeginRespValidationError" }

// Error satisfies the builtin error interface
func (e BeginRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBeginResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BeginRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BeginRespValidationError{}

// Validate checks the field values on StartReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartReqMultiError, or nil
// if none found.
func (m *StartReq) ValidateAll() error {
	return m.validate(true)
}

func (m *StartReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := StartReqValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StartReqMultiError(errors)
	}

	return nil
}

// StartReqMultiError is an error wrapping multiple validation errors returned
// by StartReq.ValidateAll() if the designated constraints aren't met.
type StartReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartReqMultiError) AllErrors() []error { return m }

// StartReqValidationError is the validation error returned by
// StartReq.Validate if the designated constraints aren't met.
type StartReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartReqValidationError) ErrorName() string { return "StartReqValidationError" }

// Error satisfies the builtin error interface
func (e StartReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartReqValidationError{}

// Validate checks the field values on StartResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StartResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StartResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StartRespMultiError, or nil
// if none found.
func (m *StartResp) ValidateAll() error {
	return m.validate(true)
}

func (m *StartResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for State

	if len(errors) > 0 {
		return StartRespMultiError(errors)
	}

	return nil
}

// StartRespMultiError is an error wrapping multiple validation errors returned
// by StartResp.ValidateAll() if the designated constraints aren't met.
type StartRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StartRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StartRespMultiError) AllErrors() []error { return m }

// StartRespValidationError is the validation error returned by
// StartResp.Validate if the designated constraints aren't met.
type StartRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StartRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StartRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StartRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StartRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StartRespValidationError) ErrorName() string { return "StartRespValidationError" }

// Error satisfies the builtin error interface
func (e StartRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStartResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StartRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StartRespValidationError{}

// Validate checks the field values on CommitReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommitReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommitReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommitReqMultiError, or nil
// if none found.
func (m *CommitReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := CommitReqValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommitReqMultiError(errors)
	}

	return nil
}

// CommitReqMultiError is an error wrapping multiple validation errors returned
// by CommitReq.ValidateAll() if the designated constraints aren't met.
type CommitReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitReqMultiError) AllErrors() []error { return m }

// CommitReqValidationError is the validation error returned by
// CommitReq.Validate if the designated constraints aren't met.
type CommitReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitReqValidationError) ErrorName() string { return "CommitReqValidationError" }

// Error satisfies the builtin error interface
func (e CommitReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitReqValidationError{}

// Validate checks the field values on CommitResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommitResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommitResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommitRespMultiError, or
// nil if none found.
func (m *CommitResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := CommitRespValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CommitRespMultiError(errors)
	}

	return nil
}

// CommitRespMultiError is an error wrapping multiple validation errors
// returned by CommitResp.ValidateAll() if the designated constraints aren't met.
type CommitRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitRespMultiError) AllErrors() []error { return m }

// CommitRespValidationError is the validation error returned by
// CommitResp.Validate if the designated constraints aren't met.
type CommitRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitRespValidationError) ErrorName() string { return "CommitRespValidationError" }

// Error satisfies the builtin error interface
func (e CommitRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitRespValidationError{}

// Validate checks the field values on RollBckReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RollBckReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RollBckReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RollBckReqMultiError, or
// nil if none found.
func (m *RollBckReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RollBckReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := RollBckReqValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RollBckReqMultiError(errors)
	}

	return nil
}

// RollBckReqMultiError is an error wrapping multiple validation errors
// returned by RollBckReq.ValidateAll() if the designated constraints aren't met.
type RollBckReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RollBckReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RollBckReqMultiError) AllErrors() []error { return m }

// RollBckReqValidationError is the validation error returned by
// RollBckReq.Validate if the designated constraints aren't met.
type RollBckReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RollBckReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RollBckReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RollBckReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RollBckReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RollBckReqValidationError) ErrorName() string { return "RollBckReqValidationError" }

// Error satisfies the builtin error interface
func (e RollBckReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRollBckReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RollBckReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RollBckReqValidationError{}

// Validate checks the field values on RollBckResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RollBckResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RollBckResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RollBckRespMultiError, or
// nil if none found.
func (m *RollBckResp) ValidateAll() error {
	return m.validate(true)
}

func (m *RollBckResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := RollBckRespValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RollBckRespMultiError(errors)
	}

	return nil
}

// RollBckRespMultiError is an error wrapping multiple validation errors
// returned by RollBckResp.ValidateAll() if the designated constraints aren't met.
type RollBckRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RollBckRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RollBckRespMultiError) AllErrors() []error { return m }

// RollBckRespValidationError is the validation error returned by
// RollBckResp.Validate if the designated constraints aren't met.
type RollBckRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RollBckRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RollBckRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RollBckRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RollBckRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RollBckRespValidationError) ErrorName() string { return "RollBckRespValidationError" }

// Error satisfies the builtin error interface
func (e RollBckRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRollBckResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RollBckRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RollBckRespValidationError{}

// Validate checks the field values on GetStateReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetStateReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStateReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetStateReqMultiError, or
// nil if none found.
func (m *GetStateReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStateReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetGId()); l < 1 || l > 50 {
		err := GetStateReqValidationError{
			field:  "GId",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetStateReqMultiError(errors)
	}

	return nil
}

// GetStateReqMultiError is an error wrapping multiple validation errors
// returned by GetStateReq.ValidateAll() if the designated constraints aren't met.
type GetStateReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStateReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStateReqMultiError) AllErrors() []error { return m }

// GetStateReqValidationError is the validation error returned by
// GetStateReq.Validate if the designated constraints aren't met.
type GetStateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStateReqValidationError) ErrorName() string { return "GetStateReqValidationError" }

// Error satisfies the builtin error interface
func (e GetStateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStateReqValidationError{}

// Validate checks the field values on GetStateResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetStateResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetStateResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetStateRespMultiError, or
// nil if none found.
func (m *GetStateResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetStateResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for State

	if len(errors) > 0 {
		return GetStateRespMultiError(errors)
	}

	return nil
}

// GetStateRespMultiError is an error wrapping multiple validation errors
// returned by GetStateResp.ValidateAll() if the designated constraints aren't met.
type GetStateRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetStateRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetStateRespMultiError) AllErrors() []error { return m }

// GetStateRespValidationError is the validation error returned by
// GetStateResp.Validate if the designated constraints aren't met.
type GetStateRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStateRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStateRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStateRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStateRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStateRespValidationError) ErrorName() string { return "GetStateRespValidationError" }

// Error satisfies the builtin error interface
func (e GetStateRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStateResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStateRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStateRespValidationError{}

// Validate checks the field values on RegisterReq_Branch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterReq_Branch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq_Branch with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterReq_BranchMultiError, or nil if none found.
func (m *RegisterReq_Branch) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq_Branch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUri()); l < 1 || l > 299 {
		err := RegisterReq_BranchValidationError{
			field:  "Uri",
			reason: "value length must be between 1 and 299 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ReqData

	// no validation rules for ReqHeader

	if _, ok := _RegisterReq_Branch_TranType_InLookup[m.GetTranType()]; !ok {
		err := RegisterReq_BranchValidationError{
			field:  "TranType",
			reason: "value must be in list [1 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _RegisterReq_Branch_Protocol_InLookup[m.GetProtocol()]; !ok {
		err := RegisterReq_BranchValidationError{
			field:  "Protocol",
			reason: "value must be in list [http https grpc]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _RegisterReq_Branch_Action_InLookup[m.GetAction()]; !ok {
		err := RegisterReq_BranchValidationError{
			field:  "Action",
			reason: "value must be in list [1 2 3 4 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLevel(); val < 1 || val > 99999 {
		err := RegisterReq_BranchValidationError{
			field:  "Level",
			reason: "value must be inside range [1, 99999]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Timeout

	if len(errors) > 0 {
		return RegisterReq_BranchMultiError(errors)
	}

	return nil
}

// RegisterReq_BranchMultiError is an error wrapping multiple validation errors
// returned by RegisterReq_Branch.ValidateAll() if the designated constraints
// aren't met.
type RegisterReq_BranchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReq_BranchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReq_BranchMultiError) AllErrors() []error { return m }

// RegisterReq_BranchValidationError is the validation error returned by
// RegisterReq_Branch.Validate if the designated constraints aren't met.
type RegisterReq_BranchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReq_BranchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReq_BranchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReq_BranchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReq_BranchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReq_BranchValidationError) ErrorName() string {
	return "RegisterReq_BranchValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterReq_BranchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq_Branch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReq_BranchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReq_BranchValidationError{}

var _RegisterReq_Branch_TranType_InLookup = map[TranType]struct{}{
	1: {},
	2: {},
}

var _RegisterReq_Branch_Protocol_InLookup = map[string]struct{}{
	"http":  {},
	"https": {},
	"grpc":  {},
}

var _RegisterReq_Branch_Action_InLookup = map[Action]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}
