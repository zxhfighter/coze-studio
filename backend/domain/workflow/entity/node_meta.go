/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package entity

import (
	"fmt"
	"strconv"

	"github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/workflow"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/entity/vo"
	"github.com/coze-dev/coze-studio/backend/types/errno"
)

type NodeType string

type NodeTypeMeta struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	Type            NodeType `json:"type"`
	Category        string   `json:"category"`
	Color           string   `json:"color"`
	Desc            string   `json:"desc"`
	IconURL         string   `json:"icon_url"`
	SupportBatch    bool     `json:"support_batch"`
	Disabled        bool     `json:"disabled,omitempty"`
	EnUSName        string   `json:"en_us_name,omitempty"`
	EnUSDescription string   `json:"en_us_description,omitempty"`

	ExecutableMeta
}

type Category struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	EnUSName string `json:"en_us_name"`
}

type StreamingParadigm string

const (
	Invoke    StreamingParadigm = "invoke"
	Stream    StreamingParadigm = "stream"
	Collect   StreamingParadigm = "collect"
	Transform StreamingParadigm = "transform"
)

type ExecutableMeta struct {
	IsComposite          bool                       `json:"is_composite,omitempty"`
	DefaultTimeoutMS     int64                      `json:"default_timeout_ms,omitempty"` // default timeout in milliseconds, 0 means no timeout
	PreFillZero          bool                       `json:"pre_fill_zero,omitempty"`
	PostFillNil          bool                       `json:"post_fill_nil,omitempty"`
	CallbackEnabled      bool                       `json:"callback_enabled,omitempty"` // is false, Eino framework will inject callbacks for this node
	MayUseChatModel      bool                       `json:"may_use_chat_model,omitempty"`
	InputSourceAware     bool                       `json:"input_source_aware,omitempty"` // whether this node needs to know the runtime status of its input sources
	StreamingParadigms   map[StreamingParadigm]bool `json:"streaming_paradigms,omitempty"`
	StreamSourceEOFAware bool                       `json:"needs_stream_source_eof,omitempty"` // whether this node needs to be aware stream sources' SourceEOF error
	/*
	 IncrementalOutput indicates that the node's output is intended for progressive, user-facing streaming.
	 This distinguishes nodes that actually stream text to the user (e.g., 'Exit', 'Output')
	 from those that are merely capable of streaming internally (defined by StreamingParadigms),
	 whose output is consumed by other nodes.
	 In essence, nodes with IncrementalOutput are a subset of those defined in StreamingParadigms.
	 When set to true, stream chunks from the node are persisted in real-time and can be fetched by get_process.
	*/
	IncrementalOutput bool `json:"incremental_output,omitempty"`
}

type PluginNodeMeta struct {
	PluginID int64    `json:"plugin_id"`
	NodeType NodeType `json:"node_type"`
	Category string   `json:"category"`
	ApiID    int64    `json:"api_id"`
	ApiName  string   `json:"api_name"`
	Name     string   `json:"name"`
	Desc     string   `json:"desc"`
	IconURL  string   `json:"icon_url"`
}

type PluginCategoryMeta struct {
	PluginCategoryMeta int64    `json:"plugin_category_meta"`
	NodeType           NodeType `json:"node_type"`
	Category           string   `json:"category"`
	Name               string   `json:"name"`
	OnlyOfficial       bool     `json:"only_official"`
	IconURL            string   `json:"icon_url"`
}

const (
	NodeTypeVariableAggregator         NodeType = "VariableAggregator"
	NodeTypeIntentDetector             NodeType = "IntentDetector"
	NodeTypeTextProcessor              NodeType = "TextProcessor"
	NodeTypeHTTPRequester              NodeType = "HTTPRequester"
	NodeTypeLoop                       NodeType = "Loop"
	NodeTypeContinue                   NodeType = "Continue"
	NodeTypeBreak                      NodeType = "Break"
	NodeTypeVariableAssigner           NodeType = "VariableAssigner"
	NodeTypeVariableAssignerWithinLoop NodeType = "VariableAssignerWithinLoop"
	NodeTypeQuestionAnswer             NodeType = "QuestionAnswer"
	NodeTypeInputReceiver              NodeType = "InputReceiver"
	NodeTypeOutputEmitter              NodeType = "OutputEmitter"
	NodeTypeDatabaseCustomSQL          NodeType = "DatabaseCustomSQL"
	NodeTypeDatabaseQuery              NodeType = "DatabaseQuery"
	NodeTypeDatabaseInsert             NodeType = "DatabaseInsert"
	NodeTypeDatabaseDelete             NodeType = "DatabaseDelete"
	NodeTypeDatabaseUpdate             NodeType = "DatabaseUpdate"
	NodeTypeKnowledgeIndexer           NodeType = "KnowledgeIndexer"
	NodeTypeKnowledgeRetriever         NodeType = "KnowledgeRetriever"
	NodeTypeKnowledgeDeleter           NodeType = "KnowledgeDeleter"
	NodeTypeEntry                      NodeType = "Entry"
	NodeTypeExit                       NodeType = "Exit"
	NodeTypeCodeRunner                 NodeType = "CodeRunner"
	NodeTypePlugin                     NodeType = "Plugin"
	NodeTypeCreateConversation         NodeType = "CreateConversation"
	NodeTypeConversationList           NodeType = "ConversationList"
	NodeTypeMessageList                NodeType = "MessageList"
	NodeTypeCreateMessage              NodeType = "CreateMessage"
	NodeTypeEditMessage                NodeType = "EditMessage"
	NodeTypeDeleteMessage              NodeType = "DeleteMessage"
	NodeTypeLambda                     NodeType = "Lambda"
	NodeTypeLLM                        NodeType = "LLM"
	NodeTypeSelector                   NodeType = "Selector"
	NodeTypeBatch                      NodeType = "Batch"
	NodeTypeSubWorkflow                NodeType = "SubWorkflow"
	NodeTypeJsonSerialization          NodeType = "JsonSerialization"
	NodeTypeJsonDeserialization        NodeType = "JsonDeserialization"
	NodeTypeConversationUpdate         NodeType = "ConversationUpdate"
	NodeTypeConversationDelete         NodeType = "ConversationDelete"
	NodeTypeClearConversationHistory   NodeType = "ClearConversationHistory"
	NodeTypeConversationHistory        NodeType = "ConversationHistory"
	NodeTypeComment                    NodeType = "Comment"
)

const (
	EntryNodeKey = "100001"
	ExitNodeKey  = "900001"
)

var blockType2NodeType = map[vo.BlockType]NodeType{
	vo.BlockTypeBotStart:                 NodeTypeEntry,
	vo.BlockTypeBotEnd:                   NodeTypeExit,
	vo.BlockTypeBotLLM:                   NodeTypeLLM,
	vo.BlockTypeBotAPI:                   NodeTypePlugin,
	vo.BlockTypeBotCode:                  NodeTypeCodeRunner,
	vo.BlockTypeBotDataset:               NodeTypeKnowledgeRetriever,
	vo.BlockTypeCondition:                NodeTypeSelector,
	vo.BlockTypeBotSubWorkflow:           NodeTypeSubWorkflow,
	vo.BlockTypeDatabase:                 NodeTypeDatabaseCustomSQL,
	vo.BlockTypeBotMessage:               NodeTypeOutputEmitter,
	vo.BlockTypeBotText:                  NodeTypeTextProcessor,
	vo.BlockTypeQuestion:                 NodeTypeQuestionAnswer,
	vo.BlockTypeBotBreak:                 NodeTypeBreak,
	vo.BlockTypeBotLoopSetVariable:       NodeTypeVariableAssignerWithinLoop,
	vo.BlockTypeBotLoop:                  NodeTypeLoop,
	vo.BlockTypeBotIntent:                NodeTypeIntentDetector,
	vo.BlockTypeBotDatasetWrite:          NodeTypeKnowledgeIndexer,
	vo.BlockTypeBotInput:                 NodeTypeInputReceiver,
	vo.BlockTypeBotBatch:                 NodeTypeBatch,
	vo.BlockTypeBotContinue:              NodeTypeContinue,
	vo.BlockTypeBotComment:               NodeTypeComment,
	vo.BlockTypeBotVariableMerge:         NodeTypeVariableAggregator,
	vo.BlockTypeCreateConversation:       NodeTypeCreateConversation,
	vo.BlockTypeBotAssignVariable:        NodeTypeVariableAssigner,
	vo.BlockTypeDatabaseUpdate:           NodeTypeDatabaseUpdate,
	vo.BlockTypeDatabaseSelect:           NodeTypeDatabaseQuery,
	vo.BlockTypeDatabaseDelete:           NodeTypeDatabaseDelete,
	vo.BlockTypeDatabaseInsert:           NodeTypeDatabaseInsert,
	vo.BlockTypeBotHttp:                  NodeTypeHTTPRequester,
	vo.BlockTypeConversationUpdate:       NodeTypeConversationUpdate,
	vo.BlockTypeConversationDelete:       NodeTypeConversationDelete,
	vo.BlockTypeJsonSerialization:        NodeTypeJsonSerialization,
	vo.BlockTypeJsonDeserialization:      NodeTypeJsonDeserialization,
	vo.BlockTypeBotDatasetDelete:         NodeTypeKnowledgeDeleter,
	vo.BlockTypeConversationList:         NodeTypeConversationList,
	vo.BlockTypeClearConversationHistory: NodeTypeClearConversationHistory,
	vo.BlockTypeConversationHistory:      NodeTypeConversationHistory,
	vo.BlockTypeBotMessageList:           NodeTypeMessageList,
	vo.BlockTypeCreateMessage:            NodeTypeCreateMessage,
	vo.BlockTypeEditeMessage:             NodeTypeEditMessage,
	vo.BlockTypeDeleteMessage:            NodeTypeDeleteMessage,
}
var nodeType2BlockType = func() map[NodeType]vo.BlockType {
	nodeType2BlockType := make(map[NodeType]vo.BlockType, len(blockType2NodeType))
	for k, v := range blockType2NodeType {
		nodeType2BlockType[v] = k
	}
	return nodeType2BlockType
}()

func BlockType2EntityNodeType(t string) (NodeType, error) {
	blockType := vo.BlockType(t)
	if nodeType, ok := blockType2NodeType[blockType]; ok {
		return nodeType, nil
	}
	return "", fmt.Errorf("cannot map block type'%s' to a node type", t)
}

func NodeTypeToAPINodeTemplateType(nodeType NodeType) (workflow.NodeTemplateType, error) {
	if blockType, ok := nodeType2BlockType[nodeType]; ok {
		blockTypeInt, err := strconv.ParseInt(string(blockType), 10, 64)
		if err != nil {
			return 0, err
		}
		return workflow.NodeTemplateType(blockTypeInt), nil
	}
	return workflow.NodeTemplateType(0), fmt.Errorf("cannot map entity node type '%s' to a workflow.NodeTemplateType", nodeType)
}

func NodeTypeToBlockType(nodeType NodeType) (vo.BlockType, error) {
	if t, ok := nodeType2BlockType[nodeType]; ok {
		return t, nil
	}
	return "", vo.WrapError(errno.ErrSchemaConversionFail,
		fmt.Errorf("cannot map entity node type '%s' to a block type", nodeType))

}
