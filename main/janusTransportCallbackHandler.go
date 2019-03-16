package main

import "janusGo/janusCore"

type JanusTransportCallbackHandler struct {

}

func (h *JanusTransportCallbackHandler) IncomingRequest(plugin *janusCore.JanusTransport, transport *janusCore.JanusTransportSession, requestId interface{},admin bool,message map[string]interface{},err interface{}) {

}

func (h *JanusTransportCallbackHandler) TransportGone(plugin *janusCore.JanusTransport, transport *janusCore.JanusTransportSession) {

}

func (h *JanusTransportCallbackHandler) IsApiSecretNeeded(plugin *janusCore.JanusTransport) bool {
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