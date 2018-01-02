package dialogflow

import "time"

type FulfillmentResponse struct {
	Speech      string              `json:"speech,omitempty"`
	DisplayText string              `json:"displayText,omitempty"`
	Source      string              `json:"source,omitempty"`
	ContextOut  []map[string]string `json:"contextOut"`
}

type FulfillmentRequest struct {
	Lang   string `json:"lang"`
	Status struct {
		ErrorType string `json:"errorType"`
		Code      int    `json:"code"`
	} `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	SessionID string    `json:"sessionId"`
	Result    struct {
		Parameters    map[string]interface{} `json:"parameters"`
		Contexts      []interface{}          `json:"contexts"`
		ResolvedQuery string                 `json:"resolvedQuery"`
		Source        string                 `json:"source"`
		Score         float64                `json:"score"`
		Speech        string                 `json:"speech"`
		Fulfillment   struct {
			Messages []struct {
				Speech string      `json:"speech"`
				Type   interface{} `json:"type"`
			} `json:"messages"`
			Speech string `json:"speech"`
		} `json:"fulfillment"`
		ActionIncomplete bool   `json:"actionIncomplete"`
		Action           string `json:"action"`
		Metadata         struct {
			IntentID                  string `json:"intentId"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
			WebhookUsed               string `json:"webhookUsed"`
		} `json:"metadata"`
	} `json:"result"`
	ID              string `json:"id"`
	OriginalRequest struct {
		Source string `json:"source"`
		Data   struct {
			Inputs []struct {
				RawInputs []struct {
					Query     string      `json:"query"`
					InputType interface{} `json:"input_type"`
				} `json:"raw_inputs"`
				Intent    string `json:"intent"`
				Arguments []struct {
					TextValue string `json:"text_value"`
					RawText   string `json:"raw_text"`
					Name      string `json:"name"`
				} `json:"arguments"`
			} `json:"inputs"`
			User struct {
				UserID string `json:"user_id"`
			} `json:"user"`
			Conversation struct {
				ConversationID    string      `json:"conversation_id"`
				Type              interface{} `json:"type"`
				ConversationToken string      `json:"conversation_token"`
			} `json:"conversation"`
		} `json:"data"`
	} `json:"originalRequest"`
}
