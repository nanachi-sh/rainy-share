package share

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
	OfText *TText
	OfAt   *TAt
}

type TText string

type TAt string

func Text(text string) Content {
	return Content{OfText: (*TText)(&text)}
}

func At(user_id string) Content {
	return Content{OfAt: (*TAt)(&user_id)}
}
