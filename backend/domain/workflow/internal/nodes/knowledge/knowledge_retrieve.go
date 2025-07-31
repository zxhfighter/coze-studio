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

package knowledge

import (
	"context"
	"errors"

	"github.com/cloudwego/eino/schema"
	oceanworkflow "github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/workflow"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/crossdomain/conversation"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/crossdomain/knowledge"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/entity/vo"
	"github.com/coze-dev/coze-studio/backend/domain/workflow/internal/execute"
	nodesconversation "github.com/coze-dev/coze-studio/backend/domain/workflow/internal/nodes/conversation"
	"github.com/coze-dev/coze-studio/backend/pkg/ctxcache"
	"github.com/coze-dev/coze-studio/backend/pkg/lang/slices"
	"github.com/coze-dev/coze-studio/backend/pkg/logs"
)

const outputList = "outputList"

type contextKey string

const chatHistoryKey contextKey = "chatHistory"

type RetrieveConfig struct {
	KnowledgeIDs       []int64
	RetrievalStrategy  *knowledge.RetrievalStrategy
	Retriever          knowledge.KnowledgeOperator
	ChatHistorySetting *vo.ChatHistorySetting
}

type KnowledgeRetrieve struct {
	config *RetrieveConfig
}

func NewKnowledgeRetrieve(_ context.Context, cfg *RetrieveConfig) (*KnowledgeRetrieve, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}

	if cfg.Retriever == nil {
		return nil, errors.New("retriever is required")
	}

	if len(cfg.KnowledgeIDs) == 0 {
		return nil, errors.New("knowledgeI ids is required")
	}

	if cfg.RetrievalStrategy == nil {
		return nil, errors.New("retrieval strategy is required")
	}

	return &KnowledgeRetrieve{
		config: cfg,
	}, nil
}

func (kr *KnowledgeRetrieve) Retrieve(ctx context.Context, input map[string]any) (map[string]any, error) {

	query, ok := input["Query"].(string)
	if !ok {
		return nil, errors.New("capital query key is required")
	}

	req := &knowledge.RetrieveRequest{
		Query:             query,
		KnowledgeIDs:      kr.config.KnowledgeIDs,
		RetrievalStrategy: kr.config.RetrievalStrategy,
		ChatHistory:       kr.GetChatHistoryOrNil(ctx, kr.config),
	}

	response, err := kr.config.Retriever.Retrieve(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	result[outputList] = slices.Transform(response.Slices, func(m *knowledge.Slice) any {
		return map[string]any{
			"documentId": m.DocumentID,
			"output":     m.Output,
		}
	})

	return result, nil
}

func (kr *KnowledgeRetrieve) GetChatHistoryOrNil(ctx context.Context, cfg *RetrieveConfig) []*schema.Message {
	if cfg.ChatHistorySetting == nil || !cfg.ChatHistorySetting.EnableChatHistory {
		return nil
	}

	exeCtx := execute.GetExeCtx(ctx)
	if exeCtx == nil {
		logs.CtxWarnf(ctx, "execute context is nil, skipping chat history")
		return nil
	}
	if exeCtx.ExeCfg.WorkflowMode != oceanworkflow.WorkflowMode_ChatFlow {
		return nil
	}

	historyFromCtx, ok := ctxcache.Get[[]*conversation.Message](ctx, chatHistoryKey)
	var messages []*conversation.Message
	if ok {
		messages = historyFromCtx
	}

	if len(messages) == 0 {
		logs.CtxWarnf(ctx, "conversation history is empty")
		return nil
	}

	historyMessages := make([]*schema.Message, 0, len(messages))
	for _, msg := range messages {
		schemaMsg, err := nodesconversation.ConvertMessageToSchema(ctx, msg)
		if err != nil {
			logs.CtxWarnf(ctx, "failed to convert history message, skipping: %v", err)
			continue
		}
		historyMessages = append(historyMessages, schemaMsg)
	}
	return historyMessages
}

func (kr *KnowledgeRetrieve) ToCallbackInput(ctx context.Context, in map[string]any) (map[string]any, error) {
	if kr.config.ChatHistorySetting == nil || !kr.config.ChatHistorySetting.EnableChatHistory {
		return in, nil
	}

	messageList, err := nodesconversation.GetConversationHistoryFromCtx(ctx, kr.config.ChatHistorySetting.ChatHistoryRound)
	if err != nil {
		logs.CtxErrorf(ctx, "failed to get conversation history: %v", err)
		return in, nil
	}

	ret := map[string]any{
		"chatHistory": messageList,
		"Query":       in["Query"],
	}
	return ret, nil
}
