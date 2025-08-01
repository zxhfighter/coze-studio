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

import * as data_connector_common from './data_connector_common';

export type Int64 = string | number;

export enum ColumnType {
  Unknown = 0,
  /** 文本 */
  Text = 1,
  /** 数字 */
  Number = 2,
  /** 时间 */
  Date = 3,
  /** float */
  Float = 4,
  /** bool */
  Boolean = 5,
  /** 图片 */
  Image = 6,
}

export enum DataSetStatus {
  DataSetprocessing = 0,
  DataSetReady = 1,
  DataSetFailed = 9,
}

export enum DatasetType {
  Coze = 0,
  Volcano = 1,
}

export enum DocumentSourceType {
  /** 本地上传 */
  Document = 0,
  /** url */
  Web = 1,
  /** 自定义类型 */
  Custom = 2,
  /** 三方 */
  ThirdParty = 3,
  /** 前端抓取 */
  FrontCrawl = 4,
  Notion = 101,
  GoogleDrive = 102,
  FeishuWeb = 103,
  LarkWeb = 104,
  WeChat = 109,
}

export enum DocumentStatus {
  /** 上传中 */
  Processing = 0,
  /** 生效 */
  Enable = 1,
  /** 失效 */
  Disable = 2,
  /** 删除 */
  Deleted = 3,
  /** 重新分片中，前端和上游不感知该状态 */
  Resegment = 4,
  /** 刷新中（刷新成功后会删除） */
  Refreshing = 5,
  /** 失败 */
  Failed = 9,
}

export enum DocumentUpdateType {
  NoUpdate = 0,
  Cover = 1,
  Append = 2,
}

/** 文件类型 */
export enum FormatType {
  /** 文本 */
  Text = 0,
  /** 表格 */
  Table = 1,
  /** 图片，暂未支持 */
  Image = 2,
  /** 数据库 */
  Database = 3,
  /** 火山结构化 火山知识库特有 */
  VolcanoStructured = 4,
  /** 火山非结构化 火山知识库特有 */
  VolcanoUnstructured = 5,
}

export enum SliceStatus {
  /** 未向量化 */
  PendingVectoring = 0,
  /** 已向量化 */
  FinishVectoring = 1,
  /** 禁用 */
  Deactive = 9,
  /** 审核不通过 */
  AuditFailed = 1000,
}

export enum VolcanoDatasetServiceStatus {
  DatasetServiceValid = 0,
  DatasetServiceInvalid = 1,
}

export enum VolcanoDatasetStatus {
  DatasetValid = 0,
  DatasetInvalid = 1,
}

export enum WebInfoStatus {
  /** 处理中 */
  Handling = 0,
  /** 已完成 */
  Finish = 1,
  /** 失败 */
  Failed = 2,
}

export interface DataSetInfo {
  data_set_id?: string;
  /** 数据集名称 */
  name?: string;
  /** 文件列表 */
  file_list?: Array<string>;
  /** 所有文件大小 */
  all_file_size?: Int64;
  /** 使用Bot数 */
  bot_used_count?: number;
  status?: DataSetStatus;
  /** 处理中的文件 */
  processing_file_list?: Array<string>;
  /** 更新时间，秒级时间戳 */
  update_time?: number;
  icon_url?: string;
  description?: string;
  icon_uri?: string;
  /** 是否可以编辑 */
  can_edit?: boolean;
  /** 创建时间，秒级时间戳 */
  create_time?: number;
  /** 创建者ID */
  creator_id?: string;
  /** 空间ID */
  space_id?: Int64;
  creator_name?: string;
  avatar_url?: string;
  /** 处理失败的文件 */
  failed_file_list?: Array<string>;
  format_type?: FormatType;
  /** 0=coze知识库 1=火山知识库 */
  dataset_type?: DatasetType;
  /** storage_config详细信息 */
  storage_config?: StorageConfig;
}

/** 表格的列信息 */
export interface DocTableColumn {
  /** 列 id */
  id?: string;
  /** 列名 */
  column_name?: string;
  /** 是否为语义匹配列 */
  is_semantic?: boolean;
  /** 列原本在 excel 的序号 */
  sequence?: string;
  /** 列类型 */
  column_type?: ColumnType;
  contains_empty_value?: boolean;
  /** 描述 */
  desc?: string;
}

export interface DocTableSheet {
  /** sheet 的编号 */
  id?: Int64;
  /** sheet 名 */
  sheet_name?: string;
  /** 总行数 */
  total_row?: Int64;
}

export interface DocumentInfo {
  name?: string;
  document_id?: string;
  /** 文件链接 */
  tos_uri?: string;
  /** 使用的bot数量 */
  bot_used_count?: number;
  /** 创建时间 */
  create_time?: number;
  /** 更新时间 */
  update_time?: number;
  /** 创建人 */
  creator_id?: string;
  /** 包含分段数量 */
  slice_count?: number;
  type?: string;
  /** 文件大小 字节数 */
  size?: number;
  /** 字符数 */
  char_count?: number;
  /** 状态 */
  status?: DocumentStatus;
  /** 命中次数 */
  hit_count?: number;
  /** 枚举 */
  source_type?: DocumentSourceType;
  /** 更新类型 */
  update_type?: DocumentUpdateType;
  /** 更新间隔 */
  update_interval?: number;
  /** 切片规则 */
  rule?: string;
  /** 文件类型 */
  format_type?: FormatType;
  /** 表格类型元数据 */
  table_meta?: Array<DocTableColumn>;
  /** url 地址 */
  web_url?: string;
  /** 状态的详细信息；如果切片失败，返回失败信息 */
  status_descript?: string;
  source_file_id?: string;
  is_disconnect?: boolean;
  /** Deprecated */
  data_source_type?: data_connector_common.DataSourceType;
}

export interface SliceInfo {
  slice_id?: string;
  content?: string;
  status?: SliceStatus;
  /** 命中次数 */
  hit_count?: string;
  /** 字符数 */
  char_count?: string;
  /** token数 */
  token_count?: string;
  /** 序号 */
  sequence?: string;
}

export interface SourceFileInfo {
  /** local: 本地文件上传的 tos 地址 */
  tos_uri?: string;
  /** web url: 如果为第一次 url 上传，传递该值 */
  submit_web_id?: string;
  /** google feishu...: 三方源文件 id */
  source_file_id?: string;
  source_type?: DocumentSourceType;
  /** custom json list<map<string, string>> */
  custom_content?: string;
}

export interface StorageConfig {
  volcano_dataset_config?: VolcanoDataset;
}

export interface VolcanoDataset {
  /** 火山侧知识库id 字符串 */
  id?: string;
  /** 名称 */
  name?: string;
  /** 类型 结构化 or 非结构化知识库 */
  format_type?: FormatType;
  /** 火山知识库详情链接 */
  link?: string;
  /** 火山知识库状态 是否已失效 */
  status?: VolcanoDatasetStatus;
}
/* eslint-enable */
