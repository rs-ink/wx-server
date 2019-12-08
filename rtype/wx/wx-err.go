package wx

import "fmt"

type ErrCode int

type Error struct {
	ErrCode ErrCode `json:"errcode"`
	ErrMsg  string  `json:"errmsg"`
}

func (err Error) IsError() bool {
	return err.ErrCode != 0
}

func (err ErrCode) String() string {
	switch int(err) {
	case 40001:
		return "不合法的调用凭证"
	case 40002:
		return "不合法的 grant_type"
	case 40003:
		return "不合法的 OpenID"
	case 40004:
		return "不合法的媒体文件类型"
	case 40007:
		return "不合法的 media_id"
	case 40008:
		return "不合法的 message_type"
	case 40009:
		return "不合法的图片大小"
	case 40010:
		return "不合法的语音大小"
	case 40011:
		return "不合法的视频大小"
	case 40012:
		return "不合法的缩略图大小"
	case 40013:
		return "不合法的 AppID"
	case 40014:
		return "不合法的 access_token"
	case 40015:
		return "不合法的菜单类型"
	case 40016:
		return "不合法的菜单按钮个数"
	case 40017:
		return "不合法的按钮类型"
	case 40018:
		return "不合法的按钮名称长度"
	case 40019:
		return "不合法的按钮 KEY 长度"
	case 40020:
		return "不合法的 url 长度"
	case 40023:
		return "不合法的子菜单按钮个数"
	case 40024:
		return "不合法的子菜单类型"
	case 40025:
		return "不合法的子菜单按钮名称长度"
	case 40026:
		return "不合法的子菜单按钮 KEY 长度"
	case 40027:
		return "不合法的子菜单按钮 url 长度"
	case 40029:
		return "不合法或已过期的 code"
	case 40030:
		return "不合法的 refresh_token"
	case 40036:
		return "不合法的 template_id 长度"
	case 40037:
		return "不合法的 template_id"
	case 40039:
		return "不合法的 url 长度"
	case 40048:
		return "不合法的 url 域名"
	case 40054:
		return "不合法的子菜单按钮 url 域名"
	case 40055:
		return "不合法的菜单按钮 url 域名"
	case 40066:
		return "不合法的 url"
	case 41001:
		return "缺失 access_token 参数"
	case 41002:
		return "缺失 appid 参数"
	case 41003:
		return "缺失 refresh_token 参数"
	case 41004:
		return "缺失 secret 参数"
	case 41005:
		return "缺失二进制媒体文件"
	case 41006:
		return "缺失 media_id 参数"
	case 41007:
		return "缺失子菜单数据"
	case 41008:
		return "缺失 code 参数"
	case 41009:
		return "缺失 openid 参数"
	case 41010:
		return "缺失 url 参数"
	case 42001:
		return "access_token 超时"
	case 42002:
		return "refresh_token 超时"
	case 42003:
		return "code 超时"
	case 43001:
		return "需要使用 GET 方法请求"
	case 43002:
		return "需要使用 POST 方法请求"
	case 43003:
		return "需要使用 HTTPS"
	case 43004:
		return "需要订阅关系"
	case 44001:
		return "空白的二进制数据"
	case 44002:
		return "空白的 POST 数据"
	case 44003:
		return "空白的 news 数据"
	case 44004:
		return "空白的内容"
	case 44005:
		return "空白的列表"
	case 45001:
		return "二进制文件超过限制"
	case 45002:
		return "content 参数超过限制"
	case 45003:
		return "title 参数超过限制"
	case 45004:
		return "description 参数超过限制"
	case 45005:
		return "url 参数长度超过限制"
	case 45006:
		return "picurl 参数超过限制"
	case 45007:
		return "播放时间超过限制（语音为 60s 最大）"
	case 45008:
		return "article 参数超过限制"
	case 45009:
		return "接口调动频率超过限制"
	case 45010:
		return "建立菜单被限制"
	case 45011:
		return "频率限制"
	case 45012:
		return "模板大小超过限制"
	case 45016:
		return "不能修改默认组"
	case 45017:
		return "修改组名过长"
	case 45018:
		return "组数量过多"
	case 50001:
		return "接口未授权"
	default:
		return fmt.Sprintf("%d", int(err))
	}
}

const (
	ErrInvalidCredential ErrCode = 40001
	ErrInvalidGrantType          = 40002
	ErrInvalidOpenid             = 40003
	ErrInvalidMedia              = 40004
)

const (
	_ ErrCode = 40006 + iota
	ErrInvalidMediaId
	ErrInvalidMessageType
	ErrInvalidImageSize
	ErrInvalidVoiceSize
	ErrInvalidVideoSize
	ErrInvalidThumbSize
	ErrInvalidAppId
	ErrInvalidAccessToken
	ErrInvalidMenuType
	ErrInvalidButtonSize
	ErrInvalidButtonType
	ErrInvalidButtonNameSize
	ErrInvalidButtonKeySize
	ErrInvalidButtonUrlSize
)
const (
	_ ErrCode = 40022 + iota
	ErrInvalidSubButtonSize
	ErrInvalidSubButtonType
	ErrInvalidSubButtonNameSize
	ErrInvalidSubButtonKeySize
	ErrInvalidSubButtonUrlSize
)
const (
	_ ErrCode = 40028 + iota
	ErrInvalidCode
	ErrInvalidRefreshToken
)
const (
	_ ErrCode = 40035 + iota
	ErrInvalidTemplateIdSize
	ErrInvalidTemplateId
)
const (
	ErrInvalidUrlSize   ErrCode = 40039
	ErrInvalidUrlDomain         = 40048
)
const (
	_ ErrCode = 40053 + iota
	ErrInvalidSubButtonUrlDomain
	ErrInvalidButtonUrlDomain
)
const (
	ErrInvalidUrl ErrCode = 40066
)
const (
	_ ErrCode = 41000 + iota
	ErrAccessTokenMissing
	ErrAppIdMissing
	ErrRefreshTokenMissing
	ErrAppSecretMissing
	ErrMediaDataMissing
	ErrMediaIdMissing
	ErrSubMenuDataMissing
	ErrMissingCode
	ErrMissingOpenId
	ErrMissingUrl
)
const (
	_ErrCode = 42000 + iota
	ErrAccessTokenExpired
	ErrRefreshTokenExpired
	ErrCodeExpired
)
const (
	_ ErrCode = 43000 + iota
	ErrRequireGetMethod
	ErrRequirePostMethod
	ErrRequireHttps
	ErrSubscribe
)
const (
	_ ErrCode = 44000 + iota
	ErrEmptyMed000iaData
	ErrEmptyPostData
	ErrEmptyNewsData
	ErrEmptyContent
	ErrEmptyListSize
)
const (
	_ ErrCode = 45000 + iota
	ErrMediaSizeOutOfLimit
	ErrContentSizeOutOfLimit
	ErrTitleSizeOutOfLimit
	ErrDescriptionSizeOutOfLimit
	ErrUrlSizeOutOfLimit
	ErrPicurlSizeOutOfLimit
	ErrPlayTimeOutOfLimit
	ErrArticleSizeOutOfLimit
	ErrApiFreqOutOfLimit
	ErrCreateMenuLimit
	ErrApiLimit
	ErrTemplateSizeOutOfLimit
)
const (
	_ ErrCode = 45015 + iota
	ErrCanotModifySysGroup
	ErrCanotSetGroupNameTooLongSysGroup
	ErrTooManyGroupNowNoNeedToAddNew
)
const (
	ErrApiUnauthorized ErrCode = 50001
)
