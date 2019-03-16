package janusCore

import "sync"

type JanusTransportSession struct {
	Transport interface{}
	destroyed bool
	mu sync.Mutex
}
type JanusIceHandle struct {

}

type JanusReuest struct {
	Transport *JanusTransport
	Instance *JanusTransportSession
	RequestId *JanusReuest
	Admin bool
	Message map[string]interface{}
}

type JanusSession struct {
	SessionId uint64
	IceHandlers map[uint64]JanusIceHandle
	LastActivity int64
	Source JanusReuest
	timeout int64
	TransportGone int64
	Mu sync.Mutex
	Destroyed int64
}

