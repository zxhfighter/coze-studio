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
import * as annotation_job from './annotation_job';

export type Int64 = string | number;

export interface OpenCloneAnnotationJobRequest {
  datasetID: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  jobID: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  targetDatasetID: string;
  jobName: string;
  base?: base.Base;
}

export interface OpenCloneAnnotationJobResponse {
  jobID?: string;
  baseResp?: base.BaseResp;
}

export interface OpenDeleteAnnotationJobRequest {
  datasetID?: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  jobID?: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  base?: base.Base;
}

export interface OpenDeleteAnnotationJobResponse {
  baseResp?: base.BaseResp;
}

export interface OpenGetAnnotationJobInstanceRequest {
  datasetID: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  jobID: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  jobInstanceID?: string;
  base?: base.Base;
}

export interface OpenGetAnnotationJobInstanceResponse {
  data?: annotation_job.AnnotationJobRunInstance;
  baseResp?: base.BaseResp;
}

export interface OpenListAnnotationJobsRequest {
  datasetID: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  cursor?: string;
  base?: base.Base;
}

export interface OpenListAnnotationJobsResponse {
  data?: Array<annotation_job.AnnotationJob>;
  nextCursor?: string;
  baseResp?: base.BaseResp;
}

export interface OpenRunAnnotationJobRequest {
  datasetID: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  jobID: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  /** 离线推理需要 */
  jwtToken?: string;
  base?: base.Base;
}

export interface OpenRunAnnotationJobResponse {
  jobInstanceID?: string;
  baseResp?: base.BaseResp;
}

export interface OpenTerminateAnnotationJobInstanceRequest {
  datasetID: string;
  'FlowDevops-Agw-OpenAPI-AppId'?: string;
  'FlowDevops-Agw-OpenAPI-SpaceId'?: string;
  jobID: string;
  'FlowDevops-Agw-OpenAPI-AccountId'?: string;
  instanceID: string;
  jwtToken?: string;
  base?: base.Base;
}

export interface OpenTerminateAnnotationJobInstanceResponse {
  baseResp?: base.BaseResp;
}
/* eslint-enable */
