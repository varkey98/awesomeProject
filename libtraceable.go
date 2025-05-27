// traceable.go
package main

/*
#cgo windows CFLAGS: -I.
#cgo windows LDFLAGS: -L. -ltraceable
#cgo LDFLAGS: -ltraceable

#include "libtraceable.h"
*/
import "C"
import (
	"errors"
)

// Wraps the traceable_libtraceable type
type LibTraceable struct {
	ptr C.traceable_libtraceable
}

// Wraps traceable_process_request_result
type ProcessResult struct {
	result C.traceable_process_request_result
}

// Initialize config
func InitConfig() C.traceable_libtraceable_config {
	return C.init_libtraceable_config()
}

// Create a new instance
func New(libCfg C.traceable_libtraceable_config) (*LibTraceable, error) {
	libCfg.log_config.mode = C.TRACEABLE_LOG_STDOUT
	libCfg.log_config.level = C.TRACEABLE_LOG_LEVEL_TRACE

	var lib C.traceable_libtraceable
	ret := C.traceable_new_libtraceable(libCfg, &lib)
	if ret != 0 {
		return nil, errors.New("failed to create libtraceable")
	}
	return &LibTraceable{ptr: lib}, nil
}

// Start instance
func (lt *LibTraceable) Start() error {
	ret := C.traceable_start_libtraceable(lt.ptr)
	if ret != 0 {
		return errors.New("failed to start libtraceable")
	}
	return nil
}

// Delete instance
func (lt *LibTraceable) Delete() error {
	ret := C.traceable_delete_libtraceable(lt.ptr)
	if ret != 0 {
		return errors.New("failed to delete libtraceable")
	}
	return nil
}

// Process request
func (lt *LibTraceable) Process(attrs C.traceable_attributes) (*ProcessResult, error) {
	var out C.traceable_process_request_result
	ret := C.traceable_process_request(lt.ptr, attrs, &out)
	if ret != 0 {
		return nil, errors.New("failed to process request")
	}
	return &ProcessResult{result: out}, nil
}

// Clean up process result
func (pr *ProcessResult) Delete() error {
	ret := C.traceable_delete_process_request_result_data(pr.result)
	if ret != 0 {
		return errors.New("failed to delete process request result")
	}
	return nil
}
