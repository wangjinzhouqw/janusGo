package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
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
)

func init(){
	//-h, --help                    Print help and exit
	flag.BoolVar(&help,"help",false,"Print help and exit")
	//-V, --version                 Print version and exit
	flag.BoolVar(&version,"version",false,"Print version and exit")
	//-b, --daemon                  Launch Janus in background as a daemon (default=off)
	flag.BoolVar(&daemon,"daemon",false,"Launch Janus in background as a daemon\n(default=off)")
	//-p, --pid-file=path           Open the specified PID file when starting Janus (default=none)
	flag.StringVar(&pid_file,"pid-file","","Open the specified PID file when starting Janus\n(default=none)")
	//-N, --disable-stdout          Disable stdout based logging  (default=off)
	flag.BoolVar(&disable_stdout, "disable-stdout",false,"Disable stdout based logging  (default=off)")
	//-L, --log-file=path           Log to the specified file (default=stdout only)
	flag.StringVar(&log_file,"log-file","","Log to the specified file (default=stdout only)")
	//-i, --interface=ipaddress     Interface to use (will be the public IP)
	flag.StringVar(&ip_interface,"interface","","Interface to use (will be the public IP)")
	//--plugins-folder=path     Plugins folder (default=./plugins)
	flag.StringVar(&plugins_folder,"plugins-folder","./plugins","Plugins folder (default=./plugins")
	//-C, --config=filename         Configuration file to use
	flag.StringVar(&config_file,"config","","Configuration file to use")
	//-F, --configs-folder=path     Configuration files folder (default=./conf)
	flag.StringVar(&configs_folder,"configs-folder","./conf","Configuration files folder (default=./conf)")
	//-c, --cert-pem=filename       DTLS certificate
	flag.StringVar(&cert_pem,"cert-pem","","DTLS certificate")
	//-k, --cert-key=filename       DTLS certificate key
	flag.StringVar(&cert_key,"cert-key","","DTLS certificate key")
	//-K, --cert-pwd=text           DTLS certificate key passphrase (if needed)
	flag.StringVar(&cert_pwd,"cert-pwd","","DTLS certificate key passphrase (if needed)")
	//-S, --stun-server=ip:port     STUN server(:port) to use, if needed (e.g., Janus behind NAT, default=none)
	flag.StringVar(&stun_server,"stun-server","","STUN server(:port) to use, if needed (e.g., Janus behind NAT, default=none)")
	//-1, --nat-1-1=ip              Public IP to put in all host candidates, assuming a 1:1 NAT is in place (e.g., Amazon EC2 instances, default=none)
	flag.StringVar(&nat_1_1_ip,"nat-1-1","","Public IP to put in all host candidates, assuming a 1:1 NAT is in place (e.g., Amazon EC2 instances, default=none)")
	//-E, --ice-enforce-list=list   Comma-separated list of the only interfaces to use for ICE gathering; partial strings are supported (e.g., eth0 or eno1,wlan0, default=none)
	flag.StringVar(&ice_enforce_list,"ice-enforce-list","","Comma-separated list of the only interfaces to use for ICE gathering; partial strings are supported (e.g., eth0 or eno1,wlan0, default=none")
	//-X, --ice-ignore-list=list    Comma-separated list of interfaces or IP addresses to ignore for ICE gathering; partial strings are supported (e.g., vmnet8,192.168.0.1,10.0.0.1 or vmnet,192.168., default=vmnet)
	flag.StringVar(&ice_ignore_list,"ice-ignore-list","","Comma-separated list of interfaces or IP addresses to ignore for ICE gathering; partial strings are supported (e.g., vmnet8,192.168.0.1,10.0.0.1 or vmnet,192.168., default=vmnet)")
	//-6, --ipv6-candidates         Whether to enable IPv6 candidates or not (experimental)  (default=off)
	flag.BoolVar(&ipv6_candidates,"ipv6-candidates",false,"Whether to enable IPv6 candidates or not (experimental)  (default=off)")
	//-l, --libnice-debug           Whether to enable libnice debugging or not (default=off)
	flag.BoolVar(&libnice_debug,"libnice-debug",false,"Whether to enable libnice debugging or not (default=off)")
	//-f, --full-trickle            Do full-trickle instead of half-trickle (default=off)
	flag.BoolVar(&full_trickle,"full-trickle",false,"Do full-trickle instead of half-trickle (default=off)")
	//-I, --ice-lite                Whether to enable the ICE Lite mode or not (default=off)
	flag.BoolVar(&ice_lite,"ice-lite",false,"Whether to enable the ICE Lite mode or not (default=off)")
	//-T, --ice-tcp                 Whether to enable ICE-TCP or not (warning: only works with ICE Lite)  (default=off)
	flag.BoolVar(&ice_tcp,"ice-tcp",false,"Whether to enable ICE-TCP or not (warning: only works with ICE Lite)  (default=off)")
	//-R, --rfc-4588                Whether to enable RFC4588 retransmissions support or not  (default=off)
	flag.BoolVar(&rfc_4588,"rfc-4588 ",false,"Whether to enable RFC4588 retransmissions support or not  (default=off)")
	//-q, --max-nack-queue=number   Maximum size of the NACK queue (in ms) per user for retransmissions
	flag.IntVar(&max_nack_queue,"max-nack-queue",0,"Maximum size of the NACK queue (in ms) per user for retransmissions")
	//-t, --no-media-timer=number   Time (in s) that should pass with no media (audio or video) being received before Janus notifies you about this
	flag.IntVar(&no_media_timer,"no-media-timer",0,"Time (in s) that should pass with no media (audio or video) being received before Janus notifies you about this")
	//-r, --rtp-port-range=min-max  Port range to use for RTP/RTCP
	flag.StringVar(&rtp_port_range,"rtp-port-range","","Port range to use for RTP/RTCP")
	//-B, --twcc-period=number      How often (in ms) to send TWCC feedback back to senders, if negotiated (default=1s)
	flag.IntVar(&twcc_period,"twcc-period",0,"How often (in ms) to send TWCC feedback back to senders, if negotiated (default=1s)")
	//-n, --server-name=name        Public name of this Janus instance (default=MyJanusInstance)
	flag.StringVar(&server_name,"server-name","","Public name of this Janus instance (default=MyJanusInstance)")
	//-s, --session-timeout=number  Session timeout value, in seconds (default=60)
	flag.IntVar(&session_timeout,"session-timeout",60,"Session timeout value, in seconds (default=60)")
	//-m, --reclaim-session-timeout=number Reclaim session timeout value, in seconds (default=0)
	flag.IntVar(&reclaim_session_timeout,"reclaim-session-timeout",0,"Reclaim session timeout value, in seconds (default=0)")
	//-d, --debug-level=1-7         Debug/logging level (0=disable debugging, 7=maximum debug level; default=4)
	flag.IntVar(&debug_level,"debug-level",4,"Debug/logging level (0=disable debugging, 7=maximum debug level; default=4)")
	//-D, --debug-timestamps        Enable debug/logging timestamps  (default=off)
	flag.BoolVar(&debug_timestamps,"debug-timestamps",false,"Enable debug/logging timestamps  (default=off)")
	//-o, --disable-colors          Disable color in the logging  (default=off)
	flag.BoolVar(&disable_colors,"disable-colors",false,"Disable color in the logging  (default=off)")
	//-M, --debug-locks             Enable debugging of locks/mutexes (very verbose!)  (default=off)
	flag.BoolVar(&debug_locks,"debug-locks",false,"Enable debugging of locks/mutexes (very verbose!)  (default=off)")
	//-a, --apisecret=randomstring  API secret all requests need to pass in order to be accepted by Janus (useful when wrapping Janus API requests in a server, none by default)
	flag.StringVar(&apisecret,"apisecret","","API secret all requests need to pass in order to be accepted by Janus (useful when wrapping Janus API requests in a server, none by default)")
	//-A, --token-auth              Enable token-based authentication for all requests  (default=off)
	flag.BoolVar(&token_auth,"token-auth",false,"Enable token-based authentication for all requests  (default=off)")
	//--token-auth-secret=randomstring Secret to verify HMAC-signed tokens with, to be used with -A
	flag.StringVar(&token_auth_secret,"token-auth-secret","","Secret to verify HMAC-signed tokens with, to be used with -A")
	//-e, --event-handlers          Enable event handlers  (default=off)
	flag.BoolVar(&event_handlers,"event-handlers",false,"Enable event handlers  (default=off)")
}

func main()  {
	funcName,fileName,line,ok := runtime.Caller(0)
	if ok{
		fmt.Println(fileName,":",line,",",runtime.FuncForPC(funcName).Name(),"begin running")
	}

	s,sep := "",""
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	flag.Parse()
	if help {
		flag.PrintDefaults()
	}

}
