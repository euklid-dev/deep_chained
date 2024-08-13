package langchain

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/tmc/langchaingo/outputparser"
// 	"github.com/tmc/langchaingo/prompts"
// )

// type ConversationParser outputparser.Defined[conversation_model.LLMConversation]

// func GenerateConversationPromptTemplate(tenant *tenant_model.Tenant, conversationParser ConversationParser) (template string, err error) {
// 	parser := outputparser.Defined[conversation_model.LLMConversation](conversationParser)

// 	promptTemplate := prompts.NewPromptTemplate(
// 		conversation_prompts.Prompt_v2,
// 		[]string{"dealer_name", "dealer_info", "response_type"},
// 	)

// 	formatInstructions := parser.GetFormatInstructions()

// 	return promptTemplate.Format(map[string]interface{}{
// 		"dealer_name":   tenant.DisplayName,
// 		"dealer_info":   tenant.Description,
// 		"response_type": formatInstructions,
// 	})

// }

// func NewConversationParser() (ConversationParser, error) {
// 	r, err := outputparser.NewDefined(conversation_model.LLMConversation{})

// 	return ConversationParser(r), err
// }

// func (c ConversationParser) ParseConversation(response string) (conversation_model.LLMConversation, error) {
// 	parser := outputparser.Defined[conversation_model.LLMConversation](c)

// 	var target conversation_model.LLMConversation
// 	var err error

// 	if target, err = parser.Parse(response); err != nil {
// 		return RetryParseConversation(response)
// 	}

// 	return target, nil

// }

// func RetryParseConversation(response string) (conversation_model.LLMConversation, error) {
// 	var target conversation_model.LLMConversation
// 	if err := json.Unmarshal([]byte(response), &target); err != nil {
// 		return target, fmt.Errorf("could not parse generated JSON: %w", err)
// 	}

// 	return target, nil
// }
