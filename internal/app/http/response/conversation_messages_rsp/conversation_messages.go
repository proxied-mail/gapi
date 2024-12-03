package conversation_messages_rsp

type ConversationMessagesResponse struct {
	Messages []ConversationMessagesResponseMessage `json:"messages"`
	LastId   string
}

type ConversationMessagesResponseMessage struct {
}
