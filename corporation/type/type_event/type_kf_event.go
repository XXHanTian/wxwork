package type_event

const (
	EventTypeKfMsgOrEvent = "kf_msg_or_event"
)

type KfEvent struct {
	Event
	Token    string `xml:"Token" json:"token"`
	OpenKfId string `xml:"OpenKfId" json:"open_kf_id"`
}
