package janusCore

type WebSoocketTransport struct {

}

func (w *WebSoocketTransport) Init(callback JanusTransportCallbacks, configPath string) int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketTransport) Destroy() {
	//panic("implement me")
}

func (w *WebSoocketTransport) GetApiCompatibility() int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketTransport) GetVersion() int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketTransport) GetVersionString() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketTransport) GetDescription() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketTransport) GetName() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketTransport) GetAuthor() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketTransport) GetPackage() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketTransport) IsJanusApiEnabled() bool {
	//panic("implement me")
	return false
}

func (w *WebSoocketTransport) IsAdminApiEnabled() bool {
	//panic("implement me")
	return false
}

func (w *WebSoocketTransport) SendMessagee(transport *JanusTransportSession, requestId JanusTransport, admin bool, message map[string]interface{}) int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketTransport) SessionCreated(transport *JanusTransportSession, sessionId uint64) {
	//panic("implement me")
}

func (w *WebSoocketTransport) SessionOver(transport *JanusTransportSession, sessionId uint64, isTimeout bool, claimed bool) {
	//panic("implement me")
}

func (w *WebSoocketTransport) SessionClaimed(transport *JanusTransportSession, sessionid uint64) {
	//panic("implement me")
}
