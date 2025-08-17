package simulator

type MessageType string

const (
	STRING  MessageType = "string"
	INTEGER MessageType = "int"
	DECIMAL MessageType = "decimal"
	BOOLEAN MessageType = "boolean"
	OBJECT  MessageType = "object"
)

type MessageRate string

const (
	ONCE MessageRate = "once"
	PS   MessageRate = "ps"
	PM   MessageRate = "pm"
	PH   MessageRate = "ph"
	PD   MessageRate = "pd"
)
