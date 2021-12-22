// 此文件为sdp2go工具自动生成,请不要手动编辑

package rpc

type ErrorCode int32

const (
	Error_Unknown              = -1
	Error_InvalidParam         = -2
	Error_InvalidZoneId        = -100
	Error_ZoneNotExist         = -101
	Error_ZoneHasExist         = -102
	Error_DirDatabase          = -103
	Error_DirServer            = -104
	Error_VersionSmall         = -105
	Error_ChannelExist         = -150
	Error_ChannelAddrError     = -151
	Error_ChannelNotExist      = -152
	Error_LoginDatabase        = -153
	Error_InvalidVersion       = -154
	Error_Bulletin_Database    = -201
	Error_Bulletin_NotExist    = -202
	Error_Bulletin_InvalidTime = -203
	Error_Gm_Param             = 1000
	Error_Gm_NoParser          = 1001
	Error_Actor_Null           = 1002
)

func (en ErrorCode) String() string {
	ret := ""
	switch en {
	case Error_Unknown:
		ret = "Error_Unknown"
	case Error_InvalidParam:
		ret = "Error_InvalidParam"
	case Error_InvalidZoneId:
		ret = "Error_InvalidZoneId"
	case Error_ZoneNotExist:
		ret = "Error_ZoneNotExist"
	case Error_ZoneHasExist:
		ret = "Error_ZoneHasExist"
	case Error_DirDatabase:
		ret = "Error_DirDatabase"
	case Error_DirServer:
		ret = "Error_DirServer"
	case Error_VersionSmall:
		ret = "Error_VersionSmall"
	case Error_ChannelExist:
		ret = "Error_ChannelExist"
	case Error_ChannelAddrError:
		ret = "Error_ChannelAddrError"
	case Error_ChannelNotExist:
		ret = "Error_ChannelNotExist"
	case Error_LoginDatabase:
		ret = "Error_LoginDatabase"
	case Error_InvalidVersion:
		ret = "Error_InvalidVersion"
	case Error_Bulletin_Database:
		ret = "Error_Bulletin_Database"
	case Error_Bulletin_NotExist:
		ret = "Error_Bulletin_NotExist"
	case Error_Bulletin_InvalidTime:
		ret = "Error_Bulletin_InvalidTime"
	case Error_Gm_Param:
		ret = "Error_Gm_Param"
	case Error_Gm_NoParser:
		ret = "Error_Gm_NoParser"
	case Error_Actor_Null:
		ret = "Error_Actor_Null"
	}
	return ret
}
