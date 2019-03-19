package janusCore

import "sync"

type JanusIceHandle struct {
	HandleId uint64
	App interface{}
	AppHandle interface{}
}

func NewJanusIceHandle() *JanusIceHandle {
	return &JanusIceHandle{}
}

type JanusReuest struct {
	Transport JanusTransport
	Instance interface{}
	RequestId *JanusReuest
	Admin bool
	Message map[string]interface{}
}

func NewJanusReuest(transport JanusTransport, instance interface{}, requestId *JanusReuest, admin bool, message map[string]interface{}) *JanusReuest {
	return &JanusReuest{Transport: transport, Instance: instance, RequestId: requestId, Admin: admin, Message: message}
}



type JanusSession struct {
	SessionId uint64
	IceHandlers map[uint64]interface{}
	LastActivity int64
	Source JanusReuest
	timeout int64
	TransportGone int64
	Mu sync.Mutex
	Destroyed int64
}

func NewJanusSession(sessionId uint64) *JanusSession {
	return &JanusSession{SessionId: sessionId}
}

