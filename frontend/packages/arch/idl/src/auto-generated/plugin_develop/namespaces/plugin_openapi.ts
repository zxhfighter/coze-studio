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

// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
/* eslint-disable */
/* tslint:disable */
// @ts-nocheck

import * as base from './base';
import * as plugin_develop_common from './plugin_develop_common';

export type Int64 = string | number;

export interface GetOAuthPluginListData {
  items?: Array<OpenAPIOAuthPluginInfo>;
}

export interface OpenAPIGetOAuthPluginListRequest {
  entity_id?: string;
  /** '授权上下文, 0-agent, 1-app, 2-workflow' */
  entity_type?: string;
  connector_id?: string;
  /** connector_uid */
  user_id?: string;
  Base?: base.Base;
}

export interface OpenAPIGetOAuthPluginListResponse {
  data?: GetOAuthPluginListData;
  /** 调用结果 */
  code: Int64;
  /** 成功为success, 失败为简单的错误信息 */
  msg?: string;
  BaseResp: base.BaseResp;
}

export interface OpenAPIOAuthPluginInfo {
  plugin_id?: string;
  /** 用户授权状态 */
  status?: plugin_develop_common.OAuthStatus;
  /** 插件name */
  plugin_name?: string;
  /** 插件头像 */
  plugin_icon?: string;
}

export interface OpenAPIRevokeAuthTokenRequest {
  /** 如果为空，该实体下的plugin全部取消授权 */
  plugin_id?: string;
  entity_id?: string;
  /** '授权上下文, 0-agent, 1-app, 2-workflow' */
  entity_type?: string;
  connector_id?: string;
  /** connector_uid */
  user_id?: string;
  Base?: base.Base;
}

export interface OpenAPIRevokeAuthTokenResponse {
  /** 调用结果 */
  code: Int64;
  /** 成功为success, 失败为简单的错误信息 */
  msg?: string;
  BaseResp: base.BaseResp;
}
/* eslint-enable */
