package rtp_rtcp

type Configuration struct {
	Audio bool
	ReceiverOnly bool
}

type RtpRtcp struct {
	Config *Configuration
}

// receiver function
//@begin
func (r *RtpRtcp)IncomingRtcpPacket(incPacket []byte){

}

func (r *RtpRtcp)SetRemoteSSRC(ssrc uint32){

}
//@end

// sender function
//@begin
func (r *RtpRtcp)SetMaxTransferUnit(size uint16){

}

func (r *RtpRtcp)SetTransportOverhead(tcp bool,ipv6 bool,authenticationOverhead uint8){

}

func (r *RtpRtcp)MaxPayloadLength() uint16{
	return 1200
}

func (r *RtpRtcp)MaxDataPayloadLength() uint16{
	return 1200
}

func (r *RtpRtcp)RegisterSendPayload(){
}

func (r *RtpRtcp)DeregisterSendPayload(){
}

func (r *RtpRtcp)RegisterSendRtpHeaderExtension(){
}

func (r *RtpRtcp)DeregisterSendRtpHeaderExtension(){
}

func (r *RtpRtcp)StartTimestamp(){
}
func (r *RtpRtcp)SetStartTimestamp(){
}

func (r *RtpRtcp)SequenceNumber(){
}
func (r *RtpRtcp)SetSequenceNumber(){
}

func (r *RtpRtcp)SetRtpState(){
}
func (r *RtpRtcp)SetRtxState(){
}

func (r *RtpRtcp)Ssrc(){
}
func (r *RtpRtcp)SetSsrc(){
}

func (r *RtpRtcp)CNAME(){
}
func (r *RtpRtcp)SetCNAME(){
}
func (r *RtpRtcp)RemoteCNAME(){
}

func (r *RtpRtcp)Rtt(){
}

//@end


func NewRtpRtcp(config *Configuration) *RtpRtcp {
	return &RtpRtcp{Config: config}
}