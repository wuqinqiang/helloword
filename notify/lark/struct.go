package lark

type body struct {
	MsgType string  `json:"msg_type"`
	Context context `json:"content"`
}

type context struct {
	Text string `json:"text"`
}
