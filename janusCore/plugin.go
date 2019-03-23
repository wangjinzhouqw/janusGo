package janusCore

const (
	JANUS_PLUGIN_ERROR = -1
	JANUS_PLUGIN_OK = iota
	JANUS_PLUGIN_OK_WAIT
	)
const JANUS_PLUGIN_API_VERSION  = 10

type JanusPluginResult struct {
	ResultType int
	Text string
	Content map[string]interface{}
}

type JanusPluginSession struct {
	JanusIcehandler interface{} // JanusIceHandle
	SpecPluginSessionHandler interface{} //plugin session instance
	Stopped bool // session stop
}

func NewJanusPluginSession(janusIceHandler interface{}) *JanusPluginSession {
	return &JanusPluginSession{JanusIcehandler: janusIceHandler}
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

	CreateSession(janusPluginSession interface{},err *error)
	HandleMessage(janusPluginSession interface{},transaction string,body map[string]interface{},jsep map[string]interface{}) JanusPluginResult
	SetupMedia(janusPluginSession interface{})
	IncomingRtp(janusPluginSession interface{},video int,buf []byte,len int)
	IncomingRtcp(janusPluginSession interface{},video int,buf []byte,len int)
	IncomingData(janusPluginSession interface{},buf []byte,len int)
	SlowLink(janusPluginSession interface{},uplink int,video int)
	HangupMedia(janusPluginSession interface{})
	DestroySession(janusPluginSession interface{},err *error)
	QuerySession(janusPluginSession interface{})
}
