package serializers

type Resp struct {
	Result interface{} `json:"Result"`
	Error  interface{} `json:"Error"`
}

type Eerror struct {
	Code    int
	Message string
}
