package janusCore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"reflect"
)

const (
	VERSION  = 1
	VERSION_STRING = "0.0.1"
	DESCRIPTION = "This transport plugin adds WebSockets support to the Janus API via libwebsockets."
	NAME = "JANUS WebSockets transport plugin"
	AUTHOR = "jinzhou.wang"
	PACKAGE = "janus.transport.websockets"
	)

type WebSoocketsTransport struct {
	gateway JanusTransportCallbacks
}

func (w *WebSoocketsTransport)handleFunc(conn *websocket.Conn) {
	fmt.Println("conn:",conn)
	ts := NewJanusWebsocketsTransportSession(conn)
	defer conn.Close()
	for {
		msg := make([]byte,2048)
		n,err := conn.Read(msg)
		if err == io.EOF { // conn disconnect
			fmt.Println("close",conn)
			conn.Close()
			break
		}

		fmt.Println(n)

		if err!=nil {
			log.Fatal(err)
		}

		content := msg[:n]
		fmt.Println(content)
		var m map[string]interface{}
		//e := json.Unmarshal(content,&m)
		dec := json.NewDecoder(bytes.NewBuffer(content))
		dec.UseNumber()
		dec.Decode(&m)

		fmt.Println(m)
		//conn.Write(content)
		w.gateway.IncomingRequest(w,ts,nil,false,m,nil)
	}
}

func (w *WebSoocketsTransport)handleAdminFunc(conn *websocket.Conn) {
	fmt.Println("conn admin:",conn)
	ts := NewJanusWebsocketsTransportSession(conn)
	defer conn.Close()
	for {
		msg := make([]byte,2048)
		n,err := conn.Read(msg)
		if err == io.EOF { // conn disconnect
			fmt.Println("close",conn)
			conn.Close()
			break
		}

		if err!=nil {
			log.Fatal(err)
		}

		content := msg[:n]
		fmt.Println(content)
		var m map[string]interface{}
		dec := json.NewDecoder(bytes.NewBuffer(content))
		dec.UseNumber()
		dec.Decode(&m)
		//e := json.Unmarshal(content,&m)
		//if e!= nil{
		//	log.Fatal(e)
		//}
		fmt.Println(m)
		w.gateway.IncomingRequest(w,ts,nil,true,m,nil)
	}
}

func (w *WebSoocketsTransport) Init(callback JanusTransportCallbacks, configPath string) int {
	w.gateway = callback
	muxServer := http.NewServeMux()
	muxServer.Handle("/",websocket.Handler(w.handleFunc))
	go http.ListenAndServe(":8188",muxServer)
	muxServerAdmin := http.NewServeMux()
	muxServerAdmin.Handle("/",websocket.Handler(w.handleAdminFunc))
	go http.ListenAndServe(":7188",muxServerAdmin)
	return 0
}

func (w *WebSoocketsTransport) Destroy() {
	//
}

func (w *WebSoocketsTransport) GetApiCompatibility() int {
	return JANUS_TRANSPORT_API_VERSION
}

func (w *WebSoocketsTransport) GetVersion() int {
	return VERSION
}

func (w *WebSoocketsTransport) GetVersionString() string {
	return VERSION_STRING
}

func (w *WebSoocketsTransport) GetDescription() string {
	return DESCRIPTION
}

func (w *WebSoocketsTransport) GetName() string {
	return NAME
}

func (w *WebSoocketsTransport) GetAuthor() string {
	return AUTHOR
}

func (w *WebSoocketsTransport) GetPackage() string {
	return PACKAGE
}

func (w *WebSoocketsTransport) IsJanusApiEnabled() bool {
	return true
}

func (w *WebSoocketsTransport) IsAdminApiEnabled() bool {
	return true
}

func (w *WebSoocketsTransport) SendMessagee(ts interface{}, requestId JanusTransport, admin bool, message []byte) int {
	jwts,ok := ts.(*JanusWebsocketsTransportSession)
	if !ok {
		fmt.Errorf("%s",reflect.TypeOf(ts).String())
	}
	jwts.conn.Write([]byte(message))
	//fmt.Println(jwts.Destroyed)
	return 0
}

func (w *WebSoocketsTransport) SessionCreated(ts interface{}, sessionId uint64) {
	// don't care
}

func (w *WebSoocketsTransport) SessionOver(ts interface{}, sessionId uint64, isTimeout bool, claimed bool) {
	// don't care
}

func (w *WebSoocketsTransport) SessionClaimed(ts interface{}, sessionid uint64) {
	// don't care
}

func NewWebsocketTransport() *WebSoocketsTransport {
	w := WebSoocketsTransport{}
	return &w
}

type JanusWebsocketsTransportSession struct {
	JanusTransportSession
	IncomingBuf []byte
	OutBuf []byte
	conn *websocket.Conn
}

func NewJanusWebsocketsTransportSession(conn *websocket.Conn) *JanusWebsocketsTransportSession {
	return &JanusWebsocketsTransportSession{conn: conn}
}


