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

type JanusCallbacks interface {
	PushEvent(handle JanusPluginSession,plugin JanusPlugin,transaction string,message map[string]interface{},jsep map[string]interface{})
	RelayRtp(handle JanusPluginSession,video int,buf []byte,len int)
	RelayRtcp(handle JanusPluginSession,video int,buf []byte,len int)
	RelayData(handle JanusPluginSession,buf []byte,len int)
	ClosePC(handle JanusPluginSession)
	EndSession(handle JanusPluginSession)
	EventsIsEnabled() bool
	NotifyEvent(plugin JanusPlugin,handle JanusPluginSession,event map[string]interface{})
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

	CreateSession(handle JanusPluginSession,err *error)
	HandleMessage(handle JanusPluginSession,transaction string,message map[string]interface{},jsep map[string]interface{}) JanusPluginResult
	SetupMedia(handle JanusPluginSession)
	IncomingRtp(handle JanusPluginSession,video int,buf []byte,len int)
	IncomingRtcp(handle JanusPluginSession,video int,buf []byte,len int)
	IncomingData(handle JanusPluginSession,buf []byte,len int)
	SlowLink(handle JanusPluginSession,uplink int,video int)
	HangupMedia(handle JanusPluginSession)
	DestroySession(handle JanusPluginSession,err *error)
	QuerySession(handle JanusPluginSession)
}
