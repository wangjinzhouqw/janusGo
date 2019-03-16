package main

import (
	"container/list"
	"flag"
	"janusGo/janusCore"
	"time"
)


type JanusParam struct {
	help bool
	version bool
	daemon bool
	pid_file string
	disable_stdout bool
	log_file string
	ip_interface string
	plugins_folder string
	config_file string
	configs_folder string
	cert_pem string
	cert_key string
	cert_pwd string
	stun_server string
	nat_1_1_ip string
	ice_enforce_list string
	ice_ignore_list string
	ipv6_candidates bool
	libnice_debug bool
	full_trickle bool
	ice_lite bool
	ice_tcp bool
	rfc_4588 bool
	max_nack_queue int
	no_media_timer int
	rtp_port_range string
	twcc_period int
	server_name string
	session_timeout int
	reclaim_session_timeout int
	debug_level int
	debug_timestamps bool
	disable_colors bool
	debug_locks bool
	apisecret string
	token_auth bool
	token_auth_secret string
	event_handlers bool
}



type JanusRunVar struct {
	sessions map[uint64]janusCore.JanusSession
	requests list.List
	transports map[string]janusCore.JanusTransport
	transports_so map[string]interface{} //open so file handler
	janusTransportCallbackhandler *JanusTransportCallbackHandler
	websocketTransport *janusCore.WebSoocketsTransport
}

var (
	janusParam JanusParam
	janusRunVar JanusRunVar
)

func init(){
	//-h, --help                    Print help and exit
	flag.BoolVar(&janusParam.help,"help",false,"Print help and exit")
	//-V, --version                 Print version and exit
	flag.BoolVar(&janusParam.version,"version",false,"Print version and exit")
	//-b, --daemon                  Launch Janus in background as a daemon (default=off)
	flag.BoolVar(&janusParam.daemon,"daemon",false,"Launch Janus in background as a daemon\n(default=off)")
	//-p, --pid-file=path           Open the specified PID file when starting Janus (default=none)
	flag.StringVar(&janusParam.pid_file,"pid-file","","Open the specified PID file when starting Janus\n(default=none)")
	//-N, --disable-stdout          Disable stdout based logging  (default=off)
	flag.BoolVar(&janusParam.disable_stdout, "disable-stdout",false,"Disable stdout based logging  (default=off)")
	//-L, --log-file=path           Log to the specified file (default=stdout only)
	flag.StringVar(&janusParam.log_file,"log-file","","Log to the specified file (default=stdout only)")
	//-i, --interface=ipaddress     Interface to use (will be the public IP)
	flag.StringVar(&janusParam.ip_interface,"interface","","Interface to use (will be the public IP)")
	//--plugins-folder=path     Plugins folder (default=./plugins)
	flag.StringVar(&janusParam.plugins_folder,"plugins-folder","./plugins","Plugins folder (default=./plugins")
	//-C, --config=filename         Configuration file to use
	flag.StringVar(&janusParam.config_file,"config","","Configuration file to use")
	//-F, --configs-folder=path     Configuration files folder (default=./conf)
	flag.StringVar(&janusParam.configs_folder,"configs-folder","./conf","Configuration files folder (default=./conf)")
	//-c, --cert-pem=filename       DTLS certificate
	flag.StringVar(&janusParam.cert_pem,"cert-pem","","DTLS certificate")
	//-k, --cert-key=filename       DTLS certificate key
	flag.StringVar(&janusParam.cert_key,"cert-key","","DTLS certificate key")
	//-K, --cert-pwd=text           DTLS certificate key passphrase (if needed)
	flag.StringVar(&janusParam.cert_pwd,"cert-pwd","","DTLS certificate key passphrase (if needed)")
	//-S, --stun-server=ip:port     STUN server(:port) to use, if needed (e.g., Janus behind NAT, default=none)
	flag.StringVar(&janusParam.stun_server,"stun-server","","STUN server(:port) to use, if needed (e.g., Janus behind NAT, default=none)")
	//-1, --nat-1-1=ip              Public IP to put in all host candidates, assuming a 1:1 NAT is in place (e.g., Amazon EC2 instances, default=none)
	flag.StringVar(&janusParam.nat_1_1_ip,"nat-1-1","","Public IP to put in all host candidates, assuming a 1:1 NAT is in place (e.g., Amazon EC2 instances, default=none)")
	//-E, --ice-enforce-list=list   Comma-separated list of the only interfaces to use for ICE gathering; partial strings are supported (e.g., eth0 or eno1,wlan0, default=none)
	flag.StringVar(&janusParam.ice_enforce_list,"ice-enforce-list","","Comma-separated list of the only interfaces to use for ICE gathering; partial strings are supported (e.g., eth0 or eno1,wlan0, default=none")
	//-X, --ice-ignore-list=list    Comma-separated list of interfaces or IP addresses to ignore for ICE gathering; partial strings are supported (e.g., vmnet8,192.168.0.1,10.0.0.1 or vmnet,192.168., default=vmnet)
	flag.StringVar(&janusParam.ice_ignore_list,"ice-ignore-list","","Comma-separated list of interfaces or IP addresses to ignore for ICE gathering; partial strings are supported (e.g., vmnet8,192.168.0.1,10.0.0.1 or vmnet,192.168., default=vmnet)")
	//-6, --ipv6-candidates         Whether to enable IPv6 candidates or not (experimental)  (default=off)
	flag.BoolVar(&janusParam.ipv6_candidates,"ipv6-candidates",false,"Whether to enable IPv6 candidates or not (experimental)  (default=off)")
	//-l, --libnice-debug           Whether to enable libnice debugging or not (default=off)
	flag.BoolVar(&janusParam.libnice_debug,"libnice-debug",false,"Whether to enable libnice debugging or not (default=off)")
	//-f, --full-trickle            Do full-trickle instead of half-trickle (default=off)
	flag.BoolVar(&janusParam.full_trickle,"full-trickle",false,"Do full-trickle instead of half-trickle (default=off)")
	//-I, --ice-lite                Whether to enable the ICE Lite mode or not (default=off)
	flag.BoolVar(&janusParam.ice_lite,"ice-lite",false,"Whether to enable the ICE Lite mode or not (default=off)")
	//-T, --ice-tcp                 Whether to enable ICE-TCP or not (warning: only works with ICE Lite)  (default=off)
	flag.BoolVar(&janusParam.ice_tcp,"ice-tcp",false,"Whether to enable ICE-TCP or not (warning: only works with ICE Lite)  (default=off)")
	//-R, --rfc-4588                Whether to enable RFC4588 retransmissions support or not  (default=off)
	flag.BoolVar(&janusParam.rfc_4588,"rfc-4588 ",false,"Whether to enable RFC4588 retransmissions support or not  (default=off)")
	//-q, --max-nack-queue=number   Maximum size of the NACK queue (in ms) per user for retransmissions
	flag.IntVar(&janusParam.max_nack_queue,"max-nack-queue",0,"Maximum size of the NACK queue (in ms) per user for retransmissions")
	//-t, --no-media-timer=number   Time (in s) that should pass with no media (audio or video) being received before Janus notifies you about this
	flag.IntVar(&janusParam.no_media_timer,"no-media-timer",0,"Time (in s) that should pass with no media (audio or video) being received before Janus notifies you about this")
	//-r, --rtp-port-range=min-max  Port range to use for RTP/RTCP
	flag.StringVar(&janusParam.rtp_port_range,"rtp-port-range","","Port range to use for RTP/RTCP")
	//-B, --twcc-period=number      How often (in ms) to send TWCC feedback back to senders, if negotiated (default=1s)
	flag.IntVar(&janusParam.twcc_period,"twcc-period",0,"How often (in ms) to send TWCC feedback back to senders, if negotiated (default=1s)")
	//-n, --server-name=name        Public name of this Janus instance (default=MyJanusInstance)
	flag.StringVar(&janusParam.server_name,"server-name","","Public name of this Janus instance (default=MyJanusInstance)")
	//-s, --session-timeout=number  Session timeout value, in seconds (default=60)
	flag.IntVar(&janusParam.session_timeout,"session-timeout",60,"Session timeout value, in seconds (default=60)")
	//-m, --reclaim-session-timeout=number Reclaim session timeout value, in seconds (default=0)
	flag.IntVar(&janusParam.reclaim_session_timeout,"reclaim-session-timeout",0,"Reclaim session timeout value, in seconds (default=0)")
	//-d, --debug-level=1-7         Debug/logging level (0=disable debugging, 7=maximum debug level; default=4)
	flag.IntVar(&janusParam.debug_level,"debug-level",4,"Debug/logging level (0=disable debugging, 7=maximum debug level; default=4)")
	//-D, --debug-timestamps        Enable debug/logging timestamps  (default=off)
	flag.BoolVar(&janusParam.debug_timestamps,"debug-timestamps",false,"Enable debug/logging timestamps  (default=off)")
	//-o, --disable-colors          Disable color in the logging  (default=off)
	flag.BoolVar(&janusParam.disable_colors,"disable-colors",false,"Disable color in the logging  (default=off)")
	//-M, --debug-locks             Enable debugging of locks/mutexes (very verbose!)  (default=off)
	flag.BoolVar(&janusParam.debug_locks,"debug-locks",false,"Enable debugging of locks/mutexes (very verbose!)  (default=off)")
	//-a, --apisecret=randomstring  API secret all requests need to pass in order to be accepted by Janus (useful when wrapping Janus API requests in a server, none by default)
	flag.StringVar(&janusParam.apisecret,"apisecret","","API secret all requests need to pass in order to be accepted by Janus (useful when wrapping Janus API requests in a server, none by default)")
	//-A, --token-auth              Enable token-based authentication for all requests  (default=off)
	flag.BoolVar(&janusParam.token_auth,"token-auth",false,"Enable token-based authentication for all requests  (default=off)")
	//--token-auth-secret=randomstring Secret to verify HMAC-signed tokens with, to be used with -A
	flag.StringVar(&janusParam.token_auth_secret,"token-auth-secret","","Secret to verify HMAC-signed tokens with, to be used with -A")
	//-e, --event-handlers          Enable event handlers  (default=off)
	flag.BoolVar(&janusParam.event_handlers,"event-handlers",false,"Enable event handlers  (default=off)")
}

func janusCheckSession(){

}
func WatchDogCheck()  {
	ticker := time.NewTicker(time.Second*2) // 2s
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			janusCheckSession()
		}
	}
}

func DaemonizeRun(){

}

//gorontine will dispatch incoming requests
func JanusTransportRequests(){
	
}

func JanusTransportTask(){

}

func LoadJanusTransport() int{
	janusRunVar.janusTransportCallbackhandler = NewJanusTransportCallbackHandler()
	janusRunVar.websocketTransport = janusCore.NewWebsocketTransport()
	janusRunVar.websocketTransport.Init(janusRunVar.janusTransportCallbackhandler,janusParam.config_file)
	janusRunVar.transports[janusRunVar.websocketTransport.GetPackage()] = janusRunVar.websocketTransport
	return 0
}

func main()  {
	flag.Parse()
	if janusParam.help {
		flag.PrintDefaults()
	}

	if(janusParam.daemon){
		DaemonizeRun()
	}

	JanusTransportRequests()
	JanusTransportTask()
	LoadJanusTransport()

	go WatchDogCheck()
	ach := make(chan int,1)
	<-ach
}
