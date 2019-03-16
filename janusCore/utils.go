package janusCore

const(
	JANUS_OBJECT = iota
	JANUS_ARRAY
	JANUS_STRING
	JANUS_INTERGER
	JANUS_REAL
	JANUS_TRUE
	JANUS_FALSE
	JANUS_NULL
)


type JanusJsonParameter struct {
	Name string
	Type uint
	Flags uint64
}
