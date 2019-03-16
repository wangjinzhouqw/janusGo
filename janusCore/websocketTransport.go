package janusCore

const (
	VERSION  = 1
	VERSION_STRING = "0.0.1"
	DESCRIPTION = "This transport plugin adds WebSockets support to the Janus API via libwebsockets."
	NAME = "JANUS WebSockets transport plugin"
	AUTHOR = "jinzhou.wang"
	PACKAGE = "janus.transport.websockets"
	)

type WebSoocketsTransport struct {
	gateway JanusTransportCallbacks
}

func (w *WebSoocketsTransport) Init(callback JanusTransportCallbacks, configPath string) int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketsTransport) Destroy() {
	//panic("implement me")
}

func (w *WebSoocketsTransport) GetApiCompatibility() int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketsTransport) GetVersion() int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketsTransport) GetVersionString() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketsTransport) GetDescription() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketsTransport) GetName() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketsTransport) GetAuthor() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketsTransport) GetPackage() string {
	//panic("implement me")
	return ""
}

func (w *WebSoocketsTransport) IsJanusApiEnabled() bool {
	//panic("implement me")
	return false
}

func (w *WebSoocketsTransport) IsAdminApiEnabled() bool {
	//panic("implement me")
	return false
}

func (w *WebSoocketsTransport) SendMessagee(transport *JanusTransportSession, requestId JanusTransport, admin bool, message map[string]interface{}) int {
	//panic("implement me")
	return 0
}

func (w *WebSoocketsTransport) SessionCreated(transport *JanusTransportSession, sessionId uint64) {
	//panic("implement me")
}

func (w *WebSoocketsTransport) SessionOver(transport *JanusTransportSession, sessionId uint64, isTimeout bool, claimed bool) {
	//panic("implement me")
}

func (w *WebSoocketsTransport) SessionClaimed(transport *JanusTransportSession, sessionid uint64) {
	//panic("implement me")
}

func NewWebsocketTransport() *WebSoocketsTransport {
	w := WebSoocketsTransport{}
	return &w
}

type JanusWebsocketsTransportSession struct {
	JanusTransportSession
	IncomingBuf []byte
	OutBuf []byte
}
