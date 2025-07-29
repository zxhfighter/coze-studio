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

package user

import (
	"context"
	"net/mail"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/developer_api"
	"github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/playground"
	"github.com/coze-dev/coze-studio/backend/api/model/passport"
	"github.com/coze-dev/coze-studio/backend/application/base/ctxutil"
	"github.com/coze-dev/coze-studio/backend/domain/user/entity"
	user "github.com/coze-dev/coze-studio/backend/domain/user/service"
	"github.com/coze-dev/coze-studio/backend/infra/contract/storage"
	"github.com/coze-dev/coze-studio/backend/pkg/errorx"
	"github.com/coze-dev/coze-studio/backend/pkg/lang/ptr"
	langSlices "github.com/coze-dev/coze-studio/backend/pkg/lang/slices"
	"github.com/coze-dev/coze-studio/backend/types/consts"
	"github.com/coze-dev/coze-studio/backend/types/errno"
)

var UserApplicationSVC = &UserApplicationService{}

type UserApplicationService struct {
	oss       storage.Storage
	DomainSVC user.User
}

// 添加一个简单的 email 验证函数
func isValidEmail(email string) bool {
	// 如果 email 字符串格式不正确，它会返回一个 error
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (u *UserApplicationService) PassportWebEmailRegisterV2(ctx context.Context, locale string, req *passport.PassportWebEmailRegisterV2PostRequest) (
	resp *passport.PassportWebEmailRegisterV2PostResponse, sessionKey string, err error,
) {
	// 验证 email 格式是否合法
	if !isValidEmail(req.GetEmail()) {
		return nil, "", errorx.New(errno.ErrUserInvalidParamCode, errorx.KV("msg", "Invalid email"))
	}

	// Allow Register Checker
	if !u.allowRegisterChecker(req.GetEmail()) {
		return nil, "", errorx.New(errno.ErrNotAllowedRegisterCode)
	}

	userInfo, err := u.DomainSVC.Create(ctx, &user.CreateUserRequest{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),

		Locale: locale,
	})
	if err != nil {
		return nil, "", err
	}

	userInfo, err = u.DomainSVC.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, "", err
	}

	return &passport.PassportWebEmailRegisterV2PostResponse{
		Data: userDo2PassportTo(userInfo),
		Code: 0,
	}, userInfo.SessionKey, nil
}

func (u *UserApplicationService) allowRegisterChecker(email string) bool {
	disableUserRegistration := os.Getenv(consts.DisableUserRegistration)
	if strings.ToLower(disableUserRegistration) != "true" {
		return true
	}

	allowedEmails := os.Getenv(consts.AllowRegistrationEmail)
	if allowedEmails == "" {
		return false
	}

	return slices.Contains(strings.Split(allowedEmails, ","), strings.ToLower(email))
}

// PassportWebLogoutGet 处理用户登出请求
func (u *UserApplicationService) PassportWebLogoutGet(ctx context.Context, req *passport.PassportWebLogoutGetRequest) (
	resp *passport.PassportWebLogoutGetResponse, err error,
) {
	uid := ctxutil.MustGetUIDFromCtx(ctx)

	err = u.DomainSVC.Logout(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &passport.PassportWebLogoutGetResponse{
		Code: 0,
	}, nil
}

// PassportWebEmailLoginPost 处理用户邮箱登录请求
func (u *UserApplicationService) PassportWebEmailLoginPost(ctx context.Context, req *passport.PassportWebEmailLoginPostRequest) (
	resp *passport.PassportWebEmailLoginPostResponse, sessionKey string, err error,
) {
	userInfo, err := u.DomainSVC.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, "", err
	}

	return &passport.PassportWebEmailLoginPostResponse{
		Data: userDo2PassportTo(userInfo),
		Code: 0,
	}, userInfo.SessionKey, nil
}

func (u *UserApplicationService) PassportWebEmailPasswordResetGet(ctx context.Context, req *passport.PassportWebEmailPasswordResetGetRequest) (
	resp *passport.PassportWebEmailPasswordResetGetResponse, err error,
) {
	err = u.DomainSVC.ResetPassword(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &passport.PassportWebEmailPasswordResetGetResponse{
		Code: 0,
	}, nil
}

func (u *UserApplicationService) PassportAccountInfoV2(ctx context.Context, req *passport.PassportAccountInfoV2Request) (
	resp *passport.PassportAccountInfoV2Response, err error,
) {
	userID := ctxutil.MustGetUIDFromCtx(ctx)

	userInfo, err := u.DomainSVC.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &passport.PassportAccountInfoV2Response{
		Data: userDo2PassportTo(userInfo),
		Code: 0,
	}, nil
}

// UserUpdateAvatar 更新用户头像
func (u *UserApplicationService) UserUpdateAvatar(ctx context.Context, mimeType string, req *passport.UserUpdateAvatarRequest) (
	resp *passport.UserUpdateAvatarResponse, err error,
) {
	// 根据 MIME type 获取文件后缀
	var ext string
	switch mimeType {
	case "image/jpeg", "image/jpg":
		ext = "jpg"
	case "image/png":
		ext = "png"
	case "image/gif":
		ext = "gif"
	case "image/webp":
		ext = "webp"
	default:
		return nil, errorx.WrapByCode(err, errno.ErrUserInvalidParamCode,
			errorx.KV("msg", "unsupported image type"))
	}

	uid := ctxutil.MustGetUIDFromCtx(ctx)

	url, err := u.DomainSVC.UpdateAvatar(ctx, uid, ext, req.GetAvatar())
	if err != nil {
		return nil, err
	}

	return &passport.UserUpdateAvatarResponse{
		Data: &passport.UserUpdateAvatarResponseData{
			WebURI: url,
		},
		Code: 0,
	}, nil
}

// UserUpdateProfile 更新用户资料
func (u *UserApplicationService) UserUpdateProfile(ctx context.Context, req *passport.UserUpdateProfileRequest) (
	resp *passport.UserUpdateProfileResponse, err error,
) {
	userID := ctxutil.MustGetUIDFromCtx(ctx)

	err = u.DomainSVC.UpdateProfile(ctx, &user.UpdateProfileRequest{
		UserID:      userID,
		Name:        req.Name,
		UniqueName:  req.UserUniqueName,
		Description: req.Description,
		Locale:      req.Locale,
	})
	if err != nil {
		return nil, err
	}

	return &passport.UserUpdateProfileResponse{
		Code: 0,
	}, nil
}

func (u *UserApplicationService) GetSpaceListV2(ctx context.Context, req *playground.GetSpaceListV2Request) (
	resp *playground.GetSpaceListV2Response, err error,
) {
	uid := ctxutil.MustGetUIDFromCtx(ctx)

	spaces, err := u.DomainSVC.GetUserSpaceList(ctx, uid)
	if err != nil {
		return nil, err
	}

	botSpaces := langSlices.Transform(spaces, func(space *entity.Space) *playground.BotSpaceV2 {
		return &playground.BotSpaceV2{
			ID:          space.ID,
			Name:        space.Name,
			Description: space.Description,
			SpaceType:   playground.SpaceType(space.SpaceType),
			IconURL:     space.IconURL,
		}
	})

	return &playground.GetSpaceListV2Response{
		Data: &playground.SpaceInfo{
			BotSpaceList:          botSpaces,
			HasPersonalSpace:      true,
			TeamSpaceNum:          0,
			RecentlyUsedSpaceList: botSpaces,
			Total:                 ptr.Of(int32(len(botSpaces))),
			HasMore:               ptr.Of(false),
		},
		Code: 0,
	}, nil
}

func (u *UserApplicationService) MGetUserBasicInfo(ctx context.Context, req *playground.MGetUserBasicInfoRequest) (
	resp *playground.MGetUserBasicInfoResponse, err error,
) {
	userIDs, err := langSlices.TransformWithErrorCheck(req.GetUserIds(), func(s string) (int64, error) {
		return strconv.ParseInt(s, 10, 64)
	})
	if err != nil {
		return nil, errorx.WrapByCode(err, errno.ErrUserInvalidParamCode, errorx.KV("msg", "invalid user id"))
	}

	userInfos, err := u.DomainSVC.MGetUserProfiles(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return &playground.MGetUserBasicInfoResponse{
		UserBasicInfoMap: langSlices.ToMap(userInfos, func(userInfo *entity.User) (string, *playground.UserBasicInfo) {
			return strconv.FormatInt(userInfo.UserID, 10), userDo2PlaygroundTo(userInfo)
		}),
		Code: 0,
	}, nil
}

func (u *UserApplicationService) UpdateUserProfileCheck(ctx context.Context, req *developer_api.UpdateUserProfileCheckRequest) (resp *developer_api.UpdateUserProfileCheckResponse, err error) {
	if req.GetUserUniqueName() == "" {
		return &developer_api.UpdateUserProfileCheckResponse{
			Code: 0,
			Msg:  "no content to update",
		}, nil
	}

	validateResp, err := u.DomainSVC.ValidateProfileUpdate(ctx, &user.ValidateProfileUpdateRequest{
		UniqueName: req.UserUniqueName,
	})
	if err != nil {
		return nil, err
	}

	return &developer_api.UpdateUserProfileCheckResponse{
		Code: int64(validateResp.Code),
		Msg:  validateResp.Msg,
	}, nil
}

func (u *UserApplicationService) ValidateSession(ctx context.Context, sessionKey string) (*entity.Session, error) {
	session, exist, err := u.DomainSVC.ValidateSession(ctx, sessionKey)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errorx.New(errno.ErrUserAuthenticationFailed, errorx.KV("reason", "session not exist"))
	}

	return session, nil
}

func userDo2PassportTo(userDo *entity.User) *passport.User {
	var locale *string
	if userDo.Locale != "" {
		locale = ptr.Of(userDo.Locale)
	}

	return &passport.User{
		UserIDStr:      userDo.UserID,
		Name:           userDo.Name,
		ScreenName:     ptr.Of(userDo.Name),
		UserUniqueName: userDo.UniqueName,
		Email:          userDo.Email,
		Description:    userDo.Description,
		AvatarURL:      userDo.IconURL,
		AppUserInfo: &passport.AppUserInfo{
			UserUniqueName: userDo.UniqueName,
		},
		Locale: locale,

		UserCreateTime: userDo.CreatedAt / 1000,
	}
}

func userDo2PlaygroundTo(userDo *entity.User) *playground.UserBasicInfo {
	return &playground.UserBasicInfo{
		UserId:         userDo.UserID,
		Username:       userDo.Name,
		UserUniqueName: ptr.Of(userDo.UniqueName),
		UserAvatar:     userDo.IconURL,
		CreateTime:     ptr.Of(userDo.CreatedAt / 1000),
	}
}
