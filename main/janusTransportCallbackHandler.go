package main

import "janusGo/janusCore"

type JanusTransportCallbackHandler struct {

}

func (h *JanusTransportCallbackHandler) IncomingRequest(plugin janusCore.JanusTransport, ts interface{}, requestId interface{}, admin bool, message map[string]interface{}, err interface{}) {
	request := janusCore.NewJanusReuest(plugin,ts,nil,admin,message)
	janusRunVar.requests.PushBack(request)
	janusRunVar.requestChan <- 1
}

func (h *JanusTransportCallbackHandler) TransportGone(plugin janusCore.JanusTransport, ts interface{}) {

}

func (h *JanusTransportCallbackHandler) IsApiSecretNeeded(plugin janusCore.JanusTransport) bool {
	return false
}

func (h *JanusTransportCallbackHandler) IsApiSecretValid(plugin janusCore.JanusTransport,apisecret string) bool {
	return false
}

func (h *JanusTransportCallbackHandler) IsAuthTokenNeeded(plugin janusCore.JanusTransport) bool {
	return false
}

func (h *JanusTransportCallbackHandler) IaAuthTokenValid(plugin janusCore.JanusTransport, token string) bool {
	return false
}

func (h *JanusTransportCallbackHandler) EventIsEnabled() bool {
	return false
}

func (h *JanusTransportCallbackHandler) NotifyEvent(plugin janusCore.JanusTransport, transport interface{},event map[string]interface{}) {

}

func NewJanusTransportCallbackHandler() *JanusTransportCallbackHandler{
	handler := JanusTransportCallbackHandler{}
	return &handler
}
