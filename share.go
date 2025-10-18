package share

type Logger interface {
	Log(LogLevel, ...any)
}

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

type Sender interface {
	Send([]byte)
}

type Builder interface {
	SendGroupMessage(group_id string, contents ...Content) []byte
}

type Message interface {
	OfGroup() MessageGroup
}

type MessageGroup interface {
	GetContents() []Content
	GetBotID() string
	GetTimestamp() int64
	GetMessageId() string
	GetSender() MessageSender
	GetGroupID() string
}

type MessageSender struct {
	ID   string
	Name string
}

type Content struct {
	OfText  *TText
	OfAt    *TAt
	OfReply *TReply
}

type TText string

type TAt string

type TReply string

func Text(text string) Content {
	return Content{OfText: (*TText)(&text)}
}

func At(user_id string) Content {
	return Content{OfAt: (*TAt)(&user_id)}
}

func Reply(message_id string) Content {
	return Content{OfReply: (*TReply)(&message_id)}
}
