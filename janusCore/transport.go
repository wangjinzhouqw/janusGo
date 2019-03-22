package janusCore

import "sync"

const JANUS_TRANSPORT_API_VERSION = 7

type JanusTransportCallbacks interface {
	IncomingRequest(plugin JanusTransport, ts interface{}, requestId interface{},admin bool,message map[string]interface{},err interface{})
	TransportGone(plugin JanusTransport, ts interface{})
	IsApiSecretNeeded(plugin JanusTransport) bool
	IsApiSecretValid(plugin JanusTransport,apisecret string) bool
	IsAuthTokenNeeded(plugin JanusTransport) bool
	IaAuthTokenValid(plugin JanusTransport, token string) bool
	EventIsEnabled() bool
	NotifyEvent(plugin JanusTransport, transport interface{},event map[string]interface{})
}

type JanusTransport interface {
	Init(callback JanusTransportCallbacks,configPath string) int
	Destroy()

	GetApiCompatibility() int
	GetVersion() int
	GetVersionString() string
	GetDescription() string
	GetName() string
	GetAuthor() string
	GetPackage() string

	IsJanusApiEnabled() bool
	IsAdminApiEnabled() bool

	SendMessagee(ts interface{},requestId JanusTransport,admin bool,message []byte) int
	SessionCreated(ts interface{},sessionId uint64)
	SessionOver(ts interface{},sessionId uint64,isTimeout bool,claimed bool)
	SessionClaimed(ts interface{},sessionid uint64)
}

type JanusTransportSession struct {
	Transport interface{}
	Destroyed bool
	Mu sync.Mutex
}

func NewJanusTransportSession(transport *JanusTransportSession) *JanusTransportSession {
	return &JanusTransportSession{Transport: transport}
}

