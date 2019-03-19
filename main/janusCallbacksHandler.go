package main

import "janusGo/janusCore"

type JanusCallbacksHandler struct {
	
}

func (j *JanusCallbacksHandler) PushEvent(handle interface{}, plugin janusCore.JanusPlugin, transaction string, message map[string]interface{}, jsep map[string]interface{}) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) RelayRtp(handle interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) RelayRtcp(handle interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) RelayData(handle interface{}, buf []byte, len int) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) ClosePC(handle interface{}) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) EndSession(handle interface{}) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) EventsIsEnabled() bool {
	panic("implement me")
}

func (j *JanusCallbacksHandler) NotifyEvent(plugin janusCore.JanusPlugin, handle interface{}, event map[string]interface{}) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) AuthIsSignatureValid(plugin janusCore.JanusPlugin, token string) {
	panic("implement me")
}

func (j *JanusCallbacksHandler) AuthSignatureContains(plugin janusCore.JanusPlugin, token string, descriptor string) {
	panic("implement me")
}

func NewJanusCallbacksHandler() *JanusCallbacksHandler {
	return &JanusCallbacksHandler{}
}




