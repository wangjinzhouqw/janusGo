package janusCore

const (
	ECHO_PLUGIN_VERSION  = 7
	ECHO_PLUGIN_VERSION_STRING = "0.0.7"
	ECHO_PLUGIN_DESCRIPTION = "This is a trivial EchoTest plugin for Janus, just used to showcase the plugin interface."
	ECHO_PLUGIN_NAME = "JANUS EchoTest plugin"
	ECHO_PLUGIN_AUTHOR = "jinzhou.wang"
	ECHO_PLUGIN_PACKAGE = "janus.plugin.echotest"
)

type EchoPlugin struct {
	janusCallback JanusCallbacks
	configPath string
}

func (e *EchoPlugin) Destroy() {
	panic("implement me")
}

func (e *EchoPlugin) GetApiCompatibility() int {
	panic("implement me")
}

func (e *EchoPlugin) GetVersion() int {
	panic("implement me")
}

func (e *EchoPlugin) GetVersionString() string {
	panic("implement me")
}

func (e *EchoPlugin) GetDescription() string {
	panic("implement me")
}

func (e *EchoPlugin) GetName() string {
	panic("implement me")
}

func (e *EchoPlugin) GetAuthor() string {
	panic("implement me")
}

func (e *EchoPlugin) GetPackage() string {
	panic("implement me")
}

func (e *EchoPlugin) CreateSession(handle interface{}, err *error) {
	panic("implement me")
}

func (e *EchoPlugin) HandleMessage(handle interface{}, transaction string, message map[string]interface{}, jsep map[string]interface{}) JanusPluginResult {
	panic("implement me")
}

func (e *EchoPlugin) SetupMedia(handle interface{}) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingRtp(handle interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingRtcp(handle interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingData(handle interface{}, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) SlowLink(handle interface{}, uplink int, video int) {
	panic("implement me")
}

func (e *EchoPlugin) HangupMedia(handle interface{}) {
	panic("implement me")
}

func (e *EchoPlugin) DestroySession(handle interface{}, err *error) {
	panic("implement me")
}

func (e *EchoPlugin) QuerySession(handle interface{}) {
	panic("implement me")
}

func NewEchoPlugin() *EchoPlugin {
	return &EchoPlugin{}
}

func (e *EchoPlugin) Init(callbacks JanusCallbacks, configPath string) int {
	e.janusCallback = callbacks
	e.configPath = configPath
	return 0
}

type JanusEchoPluginSession struct {
	JanusPluginSession

}


