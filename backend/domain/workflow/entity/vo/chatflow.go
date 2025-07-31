package vo

type ChatFlowEvent string

const (
	ChatFlowCreated          ChatFlowEvent = "conversation.chat.created"
	ChatFlowInProgress       ChatFlowEvent = "conversation.chat.in_progress"
	ChatFlowCompleted        ChatFlowEvent = "conversation.chat.completed"
	ChatFlowFailed           ChatFlowEvent = "conversation.chat.failed"
	ChatFlowRequiresAction   ChatFlowEvent = "conversation.chat.requires_action"
	ChatFlowError            ChatFlowEvent = "error"
	ChatFlowDone             ChatFlowEvent = "done"
	ChatFlowMessageDelta     ChatFlowEvent = "conversation.message.delta"
	ChatFlowMessageCompleted ChatFlowEvent = "conversation.message.completed"
)

type LastError struct {
	Code int32  `form:"code,required" json:"code,required"`
	Msg  string `form:"msg,required" json:"msg,required"`
}
type RequiredAction struct {
	Type string `form:"type" json:"type"`
	//SubmitToolOutputs *SubmitToolOutputs `form:"submit_tool_outputs" json:"submit_tool_outputs"`
}
type Usage struct {
	TokenCount   *int32 `form:"token_count" json:"token_count,omitempty"`
	OutputTokens *int32 `form:"output_count" json:"output_count,omitempty"`
	InputTokens  *int32 `form:"input_count" json:"input_count,omitempty"`
}

type Status string

const (
	Created        Status = "created"
	InProgress     Status = "in_progress"
	Completed      Status = "completed"
	Failed         Status = "failed"
	RequiresAction Status = "requires_action"
	Canceled       Status = "canceled"
)

type ChatFlowDetail struct {
	ID             string          `json:"id,omitempty"`
	ConversationID string          `json:"conversation_id,omitempty"`
	BotID          string          `json:"bot_id,omitempty"`
	Status         Status          `json:"status,omitempty"`
	LastError      *LastError      `json:"last_error,omitempty"`
	Usage          *Usage          `json:"usage,omitempty"`
	RequiredAction *RequiredAction `json:"required_action,omitempty"`
	ExecuteID      string          `json:"execute_id,omitempty"`
}

type MessageDetail struct {
	ID             string `json:"id"`
	ChatID         string `json:"chat_id"`
	ConversationID string `json:"conversation_id"`
	BotID          string `json:"bot_id"`
	Role           string `json:"role"`
	Type           string `json:"type"`
	Content        string `json:"content"`
	ContentType    string `json:"content_type"`
}

type ErrorDetail struct {
	Code     string `form:"code,required" json:"code,required"`
	Msg      string `form:"msg,required" json:"msg,required"`
	DebugUrl string `form:"debug_url" json:"debug_url,omitempty"`
}
