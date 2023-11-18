package plain

type Headers struct {
	Headers map[string]string
}

type Msg struct {
	Headers Headers
	Body    []byte
}
