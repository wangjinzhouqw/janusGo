package janusCore

import (
	"container/list"
	"fmt"
)

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
	SpecPluginSession list.List
}

func (e *EchoPlugin) Destroy() {
	panic("implement me")
}

func (e *EchoPlugin) GetApiCompatibility() int {
	return JANUS_PLUGIN_API_VERSION
}

func (e *EchoPlugin) GetVersion() int {
	return ECHO_PLUGIN_VERSION
}

func (e *EchoPlugin) GetVersionString() string {
	return ECHO_PLUGIN_VERSION_STRING
}

func (e *EchoPlugin) GetDescription() string {
	return ECHO_PLUGIN_DESCRIPTION
}

func (e *EchoPlugin) GetName() string {
	return ECHO_PLUGIN_NAME
}

func (e *EchoPlugin) GetAuthor() string {
	return ECHO_PLUGIN_AUTHOR
}

func (e *EchoPlugin) GetPackage() string {
	return ECHO_PLUGIN_PACKAGE
}

func (e *EchoPlugin) CreateSession(janusPluginSession interface{}, err *error) {
	jps,ok := janusPluginSession.(*JanusPluginSession)
	if !ok {
		fmt.Println("januspluginSession is error")
	}

	jps.SpecPluginSessionHandler = NewJanusEchoPluginSession(jps)
	e.SpecPluginSession.PushBack(jps.SpecPluginSessionHandler)
}

func (e *EchoPlugin) HandleMessage(janusPluginSession interface{}, transaction string, body map[string]interface{}, jsep map[string]interface{}) JanusPluginResult {
	jps,ok := janusPluginSession.(*JanusPluginSession)
	if !ok {
		fmt.Println("!ok")
	}
	sjps := jps.SpecPluginSessionHandler.(*JanusEchoPluginSession)
	fmt.Println(sjps)

	if jsep!=nil{
		sdpType := jsep["type"].(string)
		sdp := jsep["sdp"]
		simulcat := jsep["simulcast"]
		fmt.Println("sdpType:",sdpType,sdp,simulcat)
	}

	audioEnable := body["audio"].(bool)
	videoEnable := body["video"].(bool)
	sjps.AudioActive = audioEnable
	sjps.VideoActive = videoEnable

	return JanusPluginResult{}
}

func (e *EchoPlugin) SetupMedia(janusPluginSession interface{}) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingRtp(janusPluginSession interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingRtcp(janusPluginSession interface{}, video int, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) IncomingData(janusPluginSession interface{}, buf []byte, len int) {
	panic("implement me")
}

func (e *EchoPlugin) SlowLink(janusPluginSession interface{}, uplink int, video int) {
	panic("implement me")
}

func (e *EchoPlugin) HangupMedia(janusPluginSession interface{}) {
	panic("implement me")
}

func (e *EchoPlugin) DestroySession(janusPluginSession interface{}, err *error) {
	panic("implement me")
}

func (e *EchoPlugin) QuerySession(janusPluginSession interface{}) {
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
	*JanusPluginSession
	HasAudio bool
	HasVideo bool
	HasData bool
	AudioActive bool
	VideoActive bool
}

func NewJanusEchoPluginSession(janusPluginSession *JanusPluginSession) *JanusEchoPluginSession {
	return &JanusEchoPluginSession{JanusPluginSession: janusPluginSession}
}


