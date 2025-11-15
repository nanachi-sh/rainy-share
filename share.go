package share

type Config struct {
	Version   string
	Websocket Config_websocket
	Minecraft Config_minecraft
	WhiteList []string
	FilterQQ  []string
}

type Config_minecraft struct {
	Addr       string
	Port       int
	HexModChan string
	HexModList string
}

type Config_websocket struct {
	Addr  string
	Port  int
	Token string
}

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
	SendGroupMessage(echo string, group_id string, contents ...Content) []byte
}

type Message interface {
	OfGroup() MessageGroup
	OfEvent() MessageEvent
}

type MessageGroup interface {
	GetContents() []Content
	GetBotID() string
	GetTimestamp() int64
	GetMessageId() string
	GetSender() MessageSender
	GetGroupID() string
}

type MessageEvent interface {
	Response() (echo string, data map[string]any)
}

type MessageSender struct {
	ID   string
	Name string
}

type Content struct {
	OfText  *TText
	OfAt    *TAt
	OfReply *TReply
	OfImage *TImage
	OfVoice *TVoice
}

type TText string

type TAt string

type TReply string

type TImage string

type TVoice string

func Text(text string) Content {
	return Content{OfText: (*TText)(&text)}
}

func At(user_id string) Content {
	return Content{OfAt: (*TAt)(&user_id)}
}

func Reply(message_id string) Content {
	return Content{OfReply: (*TReply)(&message_id)}
}

func Image(url string) Content {
	return Content{OfImage: (*TImage)(&url)}
}

func Voice(url string) Content {
	return Content{OfVoice: (*TVoice)(&url)}
}
