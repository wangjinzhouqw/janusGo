package janusCore

type JanusIceTrickle struct {
	IceHandler JanusIceHandle
	Received uint64
	Transaction string
	Candidate map[string]interface{}
}

