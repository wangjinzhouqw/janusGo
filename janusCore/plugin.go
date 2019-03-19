package janusCore

const (
	JANUS_PLUGIN_ERROR = -1
	JANUS_PLUGIN_OK = iota
	JANUS_PLUGIN_OK_WAIT
	)

type JanusPluginResult struct {
	ResultType int
	Text string
	Content map[string]interface{}
}

type JanusPluginSession struct {
	Gatewayhandle interface{} // JanusIceHandle
	PluginHandle interface{} //plugin session instance
	Stopped bool // session stop
}

func NewJanusPluginSession(gatewayhandle interface{}) *JanusPluginSession {
	return &JanusPluginSession{Gatewayhandle: gatewayhandle}
}

type JanusCallbacks interface {
	PushEvent(handle interface{},plugin JanusPlugin,transaction string,message map[string]interface{},jsep map[string]interface{})
	RelayRtp(handle interface{},video int,buf []byte,len int)
	RelayRtcp(handle interface{},video int,buf []byte,len int)
	RelayData(handle interface{},buf []byte,len int)
	ClosePC(handle interface{})
	EndSession(handle interface{})
	EventsIsEnabled() bool
	NotifyEvent(plugin JanusPlugin,handle interface{},event map[string]interface{})
	AuthIsSignatureValid(plugin JanusPlugin,token string)
	AuthSignatureContains(plugin JanusPlugin,token string,descriptor string)
}

type JanusPlugin interface {
	Init(callbacks JanusCallbacks,configPath string) int
	Destroy()

	GetApiCompatibility() int
	GetVersion() int
	GetVersionString() string
	GetDescription() string
	GetName() string
	GetAuthor() string
	GetPackage() string

	CreateSession(handle interface{},err *error)
	HandleMessage(handle interface{},transaction string,message map[string]interface{},jsep map[string]interface{}) JanusPluginResult
	SetupMedia(handle interface{})
	IncomingRtp(handle interface{},video int,buf []byte,len int)
	IncomingRtcp(handle interface{},video int,buf []byte,len int)
	IncomingData(handle interface{},buf []byte,len int)
	SlowLink(handle interface{},uplink int,video int)
	HangupMedia(handle interface{})
	DestroySession(handle interface{},err *error)
	QuerySession(handle interface{})
}
