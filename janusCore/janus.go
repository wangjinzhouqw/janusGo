package janusCore

import (
	"math/rand"
	"strconv"
	"sync"
)

type JanusIceHandle struct {
	HandleId string
	JanusSessionHander interface{}
	JanusPluginHander interface{}
	JanusPluginSessionHandler interface{}

	JanusIceStreamHandler JanusIceHandle
	LocalSdp string
	RemoteSdp string
	RtpProfile string
}

func NewJanusIceHandle(janusSessionHander interface{}) *JanusIceHandle {
	jih := JanusIceHandle{
		JanusSessionHander:janusSessionHander,
		HandleId:strconv.FormatUint(rand.Uint64(),16),
	}
	return &jih
}

type JanusIceStream struct {

}

type JanusReuest struct {
	Transport JanusTransport
	TransportSessionHandler interface{}
	RequestId *JanusReuest
	Admin bool
	Message map[string]interface{}
}

func NewJanusReuest(transport JanusTransport, instance interface{}, requestId *JanusReuest, admin bool, message map[string]interface{}) *JanusReuest {
	return &JanusReuest{Transport: transport, TransportSessionHandler: instance, RequestId: requestId, Admin: admin, Message: message}
}



type JanusSession struct {
	SessionId string
	IceHandlers map[string]interface{}
	LastActivity int64
	Source JanusReuest
	timeout int64
	TransportGone int64
	Mu sync.Mutex
	Destroyed int64
}

func NewJanusSession(sessionId string) *JanusSession {
	return &JanusSession{SessionId: sessionId}
}

