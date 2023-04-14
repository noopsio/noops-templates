package handler

// #include "handler.h"
import "C"

import "unsafe"

type HandlerResponse struct {
	Status uint16
	Body   string
}

type HandlerRequest struct {
	QueryParams []HandlerTuple2StringStringT
}

type HandlerTuple2StringStringT struct {
	F0 string
	F1 string
}

var handler Handler = nil

func SetHandler(i Handler) {
	handler = i
}

type Handler interface {
	Handle(req HandlerRequest) HandlerResponse
}

//export handler_handle
func HandlerHandle(req *C.handler_request_t, ret *C.handler_response_t) {
	defer C.handler_request_free(req)
	var lift_req HandlerRequest
	var lift_req_QueryParams []HandlerTuple2StringStringT
	lift_req_QueryParams = make([]HandlerTuple2StringStringT, req.query_params.len)
	if req.query_params.len > 0 {
		for lift_req_QueryParams_i := 0; lift_req_QueryParams_i < int(req.query_params.len); lift_req_QueryParams_i++ {
			var empty_lift_req_QueryParams C.handler_tuple2_string_string_t
			lift_req_QueryParams_ptr := *(*C.handler_tuple2_string_string_t)(unsafe.Pointer(uintptr(unsafe.Pointer(req.query_params.ptr)) +
				uintptr(lift_req_QueryParams_i)*unsafe.Sizeof(empty_lift_req_QueryParams)))
			var list_lift_req_QueryParams HandlerTuple2StringStringT
			var list_lift_req_QueryParams_F0 string
			list_lift_req_QueryParams_F0 = C.GoStringN(lift_req_QueryParams_ptr.f0.ptr, C.int(lift_req_QueryParams_ptr.f0.len))
			list_lift_req_QueryParams.F0 = list_lift_req_QueryParams_F0
			var list_lift_req_QueryParams_F1 string
			list_lift_req_QueryParams_F1 = C.GoStringN(lift_req_QueryParams_ptr.f1.ptr, C.int(lift_req_QueryParams_ptr.f1.len))
			list_lift_req_QueryParams.F1 = list_lift_req_QueryParams_F1
			lift_req_QueryParams[lift_req_QueryParams_i] = list_lift_req_QueryParams
		}
	}
	lift_req.QueryParams = lift_req_QueryParams
	result := handler.Handle(lift_req)
	var lower_result C.handler_response_t
	lower_result_status := C.uint16_t(result.Status)
	lower_result.status = lower_result_status
	var lower_result_body C.handler_string_t

	lower_result_body.ptr = C.CString(result.Body)
	lower_result_body.len = C.size_t(len(result.Body))
	lower_result.body = lower_result_body
	*ret = lower_result

}
