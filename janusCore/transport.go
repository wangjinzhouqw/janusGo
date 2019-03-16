package janusCore

import "sync"

type JanusTransportCallbacks interface {
	IncomingRequest(plugin *JanusTransport, transport *JanusTransportSession, requestId interface{},admin bool,message map[string]interface{},err interface{})
	TransportGone(plugin *JanusTransport, transport *JanusTransportSession)
	IsApiSecretNeeded(plugin *JanusTransport) bool
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

	SendMessagee(transport *JanusTransportSession,requestId JanusTransport,admin bool,message map[string]interface{}) int
	SessionCreated(transport *JanusTransportSession,sessionId uint64)
	SessionOver(transport *JanusTransportSession,sessionId uint64,isTimeout bool,claimed bool)
	SessionClaimed(transport *JanusTransportSession,sessionid uint64)
}

type JanusTransportSession struct {
	Transport *JanusTransportSession
	destroyed bool
	mu sync.Mutex
}

func NewJanusTransportSession(transport *JanusTransportSession) *JanusTransportSession {
	return &JanusTransportSession{Transport: transport}
}

