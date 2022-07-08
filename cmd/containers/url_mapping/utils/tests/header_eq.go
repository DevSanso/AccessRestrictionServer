package tests


import "core/proto"


func HeaderEq(h,eq *proto.MSAHeader) bool {
	return h.HttpHeader ==eq.HttpHeader &&
	h.IsNext == eq.IsNext &&
	h.MessageIndex == eq.MessageIndex &&
	h.MessageUuid == eq.MessageUuid &&
	h.RequesterId == eq.RequesterId &&
	h.State == eq.State &&
	h.Type == eq.Type &&
	h.Method == eq.Method
}