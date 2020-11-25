// 此文件为sdp2go工具自动生成,请不要手动编辑

package rpc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/yellia1989/tex-go/sdp/protocol"
	"github.com/yellia1989/tex-go/service/model"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-go/tools/net"
	"github.com/yellia1989/tex-go/tools/sdp/codec"
	"github.com/yellia1989/tex-go/tools/sdp/util"
	"strconv"
	"time"
)

type ZoneFlagType int32

const (
	ZoneFlagType_Normal = 1
	ZoneFlagType_New    = 2
	ZoneFlagType_Close  = 3
	ZoneFlagType_Audit  = 4
)

func (en ZoneFlagType) String() string {
	ret := ""
	switch en {
	case ZoneFlagType_Normal:
		ret = "ZoneFlagType_Normal"
	case ZoneFlagType_New:
		ret = "ZoneFlagType_New"
	case ZoneFlagType_Close:
		ret = "ZoneFlagType_Close"
	case ZoneFlagType_Audit:
		ret = "ZoneFlagType_Audit"
	}
	return ret
}

type ZoneStatusType int32

const (
	ZoneStatusType_Smooth   = 1
	ZoneStatusType_Crowd    = 2
	ZoneStatusType_Busy     = 3
	ZoneStatusType_Maintain = 4
)

func (en ZoneStatusType) String() string {
	ret := ""
	switch en {
	case ZoneStatusType_Smooth:
		ret = "ZoneStatusType_Smooth"
	case ZoneStatusType_Crowd:
		ret = "ZoneStatusType_Crowd"
	case ZoneStatusType_Busy:
		ret = "ZoneStatusType_Busy"
	case ZoneStatusType_Maintain:
		ret = "ZoneStatusType_Maintain"
	}
	return ret
}

type ZoneVersion struct {
	SRes string `json:"sRes"`
	SExe string `json:"sExe"`
}

func (st *ZoneVersion) resetDefault() {
}
func (st *ZoneVersion) Copy() *ZoneVersion {
	ret := NewZoneVersion()
	ret.SRes = st.SRes
	ret.SExe = st.SExe
	return ret
}
func NewZoneVersion() *ZoneVersion {
	ret := &ZoneVersion{}
	ret.resetDefault()
	return ret
}
func (st *ZoneVersion) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sRes")+fmt.Sprintf("%v\n", st.SRes))
	util.Tab(buff, t+1, util.Fieldname("sExe")+fmt.Sprintf("%v\n", st.SExe))
}
func (st *ZoneVersion) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SRes, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SExe, 1, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ZoneVersion) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
	var err error
	var has bool
	var ty uint32

	has, ty, err = up.SkipToTag(tag, require)
	if !has || err != nil {
		return err
	}

	if ty != codec.SdpType_StructBegin {
		return fmt.Errorf("tag:%d got wrong type %d", tag, ty)
	}

	err = st.ReadStruct(up)
	if err != nil {
		return err
	}
	err = up.SkipStruct()
	if err != nil {
		return err
	}

	_ = has
	_ = ty
	return nil
}
func (st *ZoneVersion) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SRes != "" {
		err = p.WriteString(0, st.SRes)
		if err != nil {
			return err
		}
	}
	if false || st.SExe != "" {
		err = p.WriteString(1, st.SExe)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ZoneVersion) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
	var err error

	if require {
		err = p.WriteHeader(tag, codec.SdpType_StructBegin)
		if err != nil {
			return err
		}
		err = st.WriteStruct(p)
		if err != nil {
			return err
		}
		err = p.WriteHeader(0, codec.SdpType_StructEnd)
		if err != nil {
			return err
		}
	} else {
		p2 := codec.NewPacker()
		err = st.WriteStruct(p2)
		if err != nil {
			return err
		}
		if p2.Len() != 0 {
			err = p.WriteHeader(tag, codec.SdpType_StructBegin)
			if err != nil {
				return err
			}
			err = p.WriteData(p2.ToBytes())
			if err != nil {
				return err
			}
			err = p.WriteHeader(0, codec.SdpType_StructEnd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type ZoneInfo struct {
	IZoneId           uint32                 `json:"iZoneId"`
	SZoneName         string                 `json:"sZoneName"`
	SConnServer       string                 `json:"sConnServer"`
	SGameServer       string                 `json:"sGameServer"`
	IZoneFlag         uint32                 `json:"iZoneFlag"`
	IIsManual         uint32                 `json:"iIsManual"`
	IManualZoneStatus uint32                 `json:"iManualZoneStatus"`
	IMaxNum           uint32                 `json:"iMaxNum"`
	IPublishTime      uint32                 `json:"iPublishTime"`
	IIsKick           uint32                 `json:"iIsKick"`
	MVersion          map[string]ZoneVersion `json:"mVersion"`
	IMaxOnline        uint32                 `json:"iMaxOnline"`
	ICurNum           uint32                 `json:"iCurNum"`
	ILastReportTime   uint32                 `json:"iLastReportTime"`
	ICurZoneStatus    uint32                 `json:"iCurZoneStatus"`
	ICurOnline        uint32                 `json:"iCurOnline"`
}

func (st *ZoneInfo) resetDefault() {
}
func (st *ZoneInfo) Copy() *ZoneInfo {
	ret := NewZoneInfo()
	ret.IZoneId = st.IZoneId
	ret.SZoneName = st.SZoneName
	ret.SConnServer = st.SConnServer
	ret.SGameServer = st.SGameServer
	ret.IZoneFlag = st.IZoneFlag
	ret.IIsManual = st.IIsManual
	ret.IManualZoneStatus = st.IManualZoneStatus
	ret.IMaxNum = st.IMaxNum
	ret.IPublishTime = st.IPublishTime
	ret.IIsKick = st.IIsKick
	ret.MVersion = make(map[string]ZoneVersion)
	for k, v := range st.MVersion {
		ret.MVersion[k] = *(v.Copy())
	}
	ret.IMaxOnline = st.IMaxOnline
	ret.ICurNum = st.ICurNum
	ret.ILastReportTime = st.ILastReportTime
	ret.ICurZoneStatus = st.ICurZoneStatus
	ret.ICurOnline = st.ICurOnline
	return ret
}
func NewZoneInfo() *ZoneInfo {
	ret := &ZoneInfo{}
	ret.resetDefault()
	return ret
}
func (st *ZoneInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iZoneId")+fmt.Sprintf("%v\n", st.IZoneId))
	util.Tab(buff, t+1, util.Fieldname("sZoneName")+fmt.Sprintf("%v\n", st.SZoneName))
	util.Tab(buff, t+1, util.Fieldname("sConnServer")+fmt.Sprintf("%v\n", st.SConnServer))
	util.Tab(buff, t+1, util.Fieldname("sGameServer")+fmt.Sprintf("%v\n", st.SGameServer))
	util.Tab(buff, t+1, util.Fieldname("iZoneFlag")+fmt.Sprintf("%v\n", st.IZoneFlag))
	util.Tab(buff, t+1, util.Fieldname("iIsManual")+fmt.Sprintf("%v\n", st.IIsManual))
	util.Tab(buff, t+1, util.Fieldname("iManualZoneStatus")+fmt.Sprintf("%v\n", st.IManualZoneStatus))
	util.Tab(buff, t+1, util.Fieldname("iMaxNum")+fmt.Sprintf("%v\n", st.IMaxNum))
	util.Tab(buff, t+1, util.Fieldname("iPublishTime")+fmt.Sprintf("%v\n", st.IPublishTime))
	util.Tab(buff, t+1, util.Fieldname("iIsKick")+fmt.Sprintf("%v\n", st.IIsKick))
	util.Tab(buff, t+1, util.Fieldname("mVersion")+strconv.Itoa(len(st.MVersion)))
	if len(st.MVersion) == 0 {
		buff.WriteString(", {}\n")
	} else {
		buff.WriteString(", {\n")
	}
	for k, v := range st.MVersion {
		util.Tab(buff, t+1+1, "(\n")
		util.Tab(buff, t+1+2, util.Fieldname("")+fmt.Sprintf("%v\n", k))
		util.Tab(buff, t+1+2, util.Fieldname("")+"{\n")
		v.Visit(buff, t+1+2+1)
		util.Tab(buff, t+1+2, "}\n")
		util.Tab(buff, t+1+1, ")\n")
	}
	if len(st.MVersion) != 0 {
		util.Tab(buff, t+1, "}\n")
	}
	util.Tab(buff, t+1, util.Fieldname("iMaxOnline")+fmt.Sprintf("%v\n", st.IMaxOnline))
	util.Tab(buff, t+1, util.Fieldname("iCurNum")+fmt.Sprintf("%v\n", st.ICurNum))
	util.Tab(buff, t+1, util.Fieldname("iLastReportTime")+fmt.Sprintf("%v\n", st.ILastReportTime))
	util.Tab(buff, t+1, util.Fieldname("iCurZoneStatus")+fmt.Sprintf("%v\n", st.ICurZoneStatus))
	util.Tab(buff, t+1, util.Fieldname("iCurOnline")+fmt.Sprintf("%v\n", st.ICurOnline))
}
func (st *ZoneInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IZoneId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SZoneName, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SConnServer, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SGameServer, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IZoneFlag, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IIsManual, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IManualZoneStatus, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IMaxNum, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPublishTime, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IIsKick, 14, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(15, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Map {
			return fmt.Errorf("tag:%d got wrong type %d", 15, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.MVersion = make(map[string]ZoneVersion)
		for i := uint32(0); i < length; i++ {
			var k string
			err = up.ReadString(&k, 0, true)
			if err != nil {
				return err
			}
			var v ZoneVersion
			err = v.ReadStructFromTag(up, 0, true)
			if err != nil {
				return err
			}
			st.MVersion[k] = v
		}
	}
	err = up.ReadUint32(&st.IMaxOnline, 16, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICurNum, 20, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ILastReportTime, 21, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICurZoneStatus, 22, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICurOnline, 23, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ZoneInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
	var err error
	var has bool
	var ty uint32

	has, ty, err = up.SkipToTag(tag, require)
	if !has || err != nil {
		return err
	}

	if ty != codec.SdpType_StructBegin {
		return fmt.Errorf("tag:%d got wrong type %d", tag, ty)
	}

	err = st.ReadStruct(up)
	if err != nil {
		return err
	}
	err = up.SkipStruct()
	if err != nil {
		return err
	}

	_ = has
	_ = ty
	return nil
}
func (st *ZoneInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IZoneId != 0 {
		err = p.WriteUint32(0, st.IZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.SZoneName != "" {
		err = p.WriteString(1, st.SZoneName)
		if err != nil {
			return err
		}
	}
	if false || st.SConnServer != "" {
		err = p.WriteString(2, st.SConnServer)
		if err != nil {
			return err
		}
	}
	if false || st.SGameServer != "" {
		err = p.WriteString(3, st.SGameServer)
		if err != nil {
			return err
		}
	}
	if false || st.IZoneFlag != 0 {
		err = p.WriteUint32(4, st.IZoneFlag)
		if err != nil {
			return err
		}
	}
	if false || st.IIsManual != 0 {
		err = p.WriteUint32(5, st.IIsManual)
		if err != nil {
			return err
		}
	}
	if false || st.IManualZoneStatus != 0 {
		err = p.WriteUint32(6, st.IManualZoneStatus)
		if err != nil {
			return err
		}
	}
	if false || st.IMaxNum != 0 {
		err = p.WriteUint32(7, st.IMaxNum)
		if err != nil {
			return err
		}
	}
	if false || st.IPublishTime != 0 {
		err = p.WriteUint32(8, st.IPublishTime)
		if err != nil {
			return err
		}
	}
	if false || st.IIsKick != 0 {
		err = p.WriteUint32(14, st.IIsKick)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.MVersion))
	if false || length != 0 {
		err = p.WriteHeader(15, codec.SdpType_Map)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _k, _v := range st.MVersion {
			if true || _k != "" {
				err = p.WriteString(0, _k)
				if err != nil {
					return err
				}
			}
			err = _v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}
	if false || st.IMaxOnline != 0 {
		err = p.WriteUint32(16, st.IMaxOnline)
		if err != nil {
			return err
		}
	}
	if false || st.ICurNum != 0 {
		err = p.WriteUint32(20, st.ICurNum)
		if err != nil {
			return err
		}
	}
	if false || st.ILastReportTime != 0 {
		err = p.WriteUint32(21, st.ILastReportTime)
		if err != nil {
			return err
		}
	}
	if false || st.ICurZoneStatus != 0 {
		err = p.WriteUint32(22, st.ICurZoneStatus)
		if err != nil {
			return err
		}
	}
	if false || st.ICurOnline != 0 {
		err = p.WriteUint32(23, st.ICurOnline)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ZoneInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
	var err error

	if require {
		err = p.WriteHeader(tag, codec.SdpType_StructBegin)
		if err != nil {
			return err
		}
		err = st.WriteStruct(p)
		if err != nil {
			return err
		}
		err = p.WriteHeader(0, codec.SdpType_StructEnd)
		if err != nil {
			return err
		}
	} else {
		p2 := codec.NewPacker()
		err = st.WriteStruct(p2)
		if err != nil {
			return err
		}
		if p2.Len() != 0 {
			err = p.WriteHeader(tag, codec.SdpType_StructBegin)
			if err != nil {
				return err
			}
			err = p.WriteData(p2.ToBytes())
			if err != nil {
				return err
			}
			err = p.WriteHeader(0, codec.SdpType_StructEnd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type ZoneModifyInfo struct {
	BModifyZoneName              bool `json:"bModifyZoneName"`
	BModifyConnServer            bool `json:"bModifyConnServer"`
	BModifyGameServer            bool `json:"bModifyGameServer"`
	BModifyZoneFlag              bool `json:"bModifyZoneFlag"`
	BModifyIsManual              bool `json:"bModifyIsManual"`
	BModifyManualZoneStatus      bool `json:"bModifyManualZoneStatus"`
	BModifyMaxNum                bool `json:"bModifyMaxNum"`
	BModifyPublishTime           bool `json:"bModifyPublishTime"`
	BModifyClientVersion         bool `json:"bModifyClientVersion"`
	BModifyForceUpdateVersion    bool `json:"bModifyForceUpdateVersion"`
	BModifyAndClientVersion      bool `json:"bModifyAndClientVersion"`
	BModifyAndForceUpdateVersion bool `json:"bModifyAndForceUpdateVersion"`
	BKick                        bool `json:"bKick"`
}

func (st *ZoneModifyInfo) resetDefault() {
}
func (st *ZoneModifyInfo) Copy() *ZoneModifyInfo {
	ret := NewZoneModifyInfo()
	ret.BModifyZoneName = st.BModifyZoneName
	ret.BModifyConnServer = st.BModifyConnServer
	ret.BModifyGameServer = st.BModifyGameServer
	ret.BModifyZoneFlag = st.BModifyZoneFlag
	ret.BModifyIsManual = st.BModifyIsManual
	ret.BModifyManualZoneStatus = st.BModifyManualZoneStatus
	ret.BModifyMaxNum = st.BModifyMaxNum
	ret.BModifyPublishTime = st.BModifyPublishTime
	ret.BModifyClientVersion = st.BModifyClientVersion
	ret.BModifyForceUpdateVersion = st.BModifyForceUpdateVersion
	ret.BModifyAndClientVersion = st.BModifyAndClientVersion
	ret.BModifyAndForceUpdateVersion = st.BModifyAndForceUpdateVersion
	ret.BKick = st.BKick
	return ret
}
func NewZoneModifyInfo() *ZoneModifyInfo {
	ret := &ZoneModifyInfo{}
	ret.resetDefault()
	return ret
}
func (st *ZoneModifyInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("bModifyZoneName")+fmt.Sprintf("%v\n", st.BModifyZoneName))
	util.Tab(buff, t+1, util.Fieldname("bModifyConnServer")+fmt.Sprintf("%v\n", st.BModifyConnServer))
	util.Tab(buff, t+1, util.Fieldname("bModifyGameServer")+fmt.Sprintf("%v\n", st.BModifyGameServer))
	util.Tab(buff, t+1, util.Fieldname("bModifyZoneFlag")+fmt.Sprintf("%v\n", st.BModifyZoneFlag))
	util.Tab(buff, t+1, util.Fieldname("bModifyIsManual")+fmt.Sprintf("%v\n", st.BModifyIsManual))
	util.Tab(buff, t+1, util.Fieldname("bModifyManualZoneStatus")+fmt.Sprintf("%v\n", st.BModifyManualZoneStatus))
	util.Tab(buff, t+1, util.Fieldname("bModifyMaxNum")+fmt.Sprintf("%v\n", st.BModifyMaxNum))
	util.Tab(buff, t+1, util.Fieldname("bModifyPublishTime")+fmt.Sprintf("%v\n", st.BModifyPublishTime))
	util.Tab(buff, t+1, util.Fieldname("bModifyClientVersion")+fmt.Sprintf("%v\n", st.BModifyClientVersion))
	util.Tab(buff, t+1, util.Fieldname("bModifyForceUpdateVersion")+fmt.Sprintf("%v\n", st.BModifyForceUpdateVersion))
	util.Tab(buff, t+1, util.Fieldname("bModifyAndClientVersion")+fmt.Sprintf("%v\n", st.BModifyAndClientVersion))
	util.Tab(buff, t+1, util.Fieldname("bModifyAndForceUpdateVersion")+fmt.Sprintf("%v\n", st.BModifyAndForceUpdateVersion))
	util.Tab(buff, t+1, util.Fieldname("bKick")+fmt.Sprintf("%v\n", st.BKick))
}
func (st *ZoneModifyInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadBool(&st.BModifyZoneName, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyConnServer, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyGameServer, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyZoneFlag, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyIsManual, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyManualZoneStatus, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyMaxNum, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyPublishTime, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyClientVersion, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyForceUpdateVersion, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyAndClientVersion, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BModifyAndForceUpdateVersion, 12, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BKick, 13, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ZoneModifyInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
	var err error
	var has bool
	var ty uint32

	has, ty, err = up.SkipToTag(tag, require)
	if !has || err != nil {
		return err
	}

	if ty != codec.SdpType_StructBegin {
		return fmt.Errorf("tag:%d got wrong type %d", tag, ty)
	}

	err = st.ReadStruct(up)
	if err != nil {
		return err
	}
	err = up.SkipStruct()
	if err != nil {
		return err
	}

	_ = has
	_ = ty
	return nil
}
func (st *ZoneModifyInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.BModifyZoneName != false {
		err = p.WriteBool(1, st.BModifyZoneName)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyConnServer != false {
		err = p.WriteBool(2, st.BModifyConnServer)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyGameServer != false {
		err = p.WriteBool(3, st.BModifyGameServer)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyZoneFlag != false {
		err = p.WriteBool(4, st.BModifyZoneFlag)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyIsManual != false {
		err = p.WriteBool(5, st.BModifyIsManual)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyManualZoneStatus != false {
		err = p.WriteBool(6, st.BModifyManualZoneStatus)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyMaxNum != false {
		err = p.WriteBool(7, st.BModifyMaxNum)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyPublishTime != false {
		err = p.WriteBool(8, st.BModifyPublishTime)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyClientVersion != false {
		err = p.WriteBool(9, st.BModifyClientVersion)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyForceUpdateVersion != false {
		err = p.WriteBool(10, st.BModifyForceUpdateVersion)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyAndClientVersion != false {
		err = p.WriteBool(11, st.BModifyAndClientVersion)
		if err != nil {
			return err
		}
	}
	if false || st.BModifyAndForceUpdateVersion != false {
		err = p.WriteBool(12, st.BModifyAndForceUpdateVersion)
		if err != nil {
			return err
		}
	}
	if false || st.BKick != false {
		err = p.WriteBool(13, st.BKick)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ZoneModifyInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
	var err error

	if require {
		err = p.WriteHeader(tag, codec.SdpType_StructBegin)
		if err != nil {
			return err
		}
		err = st.WriteStruct(p)
		if err != nil {
			return err
		}
		err = p.WriteHeader(0, codec.SdpType_StructEnd)
		if err != nil {
			return err
		}
	} else {
		p2 := codec.NewPacker()
		err = st.WriteStruct(p2)
		if err != nil {
			return err
		}
		if p2.Len() != 0 {
			err = p.WriteHeader(tag, codec.SdpType_StructBegin)
			if err != nil {
				return err
			}
			err = p.WriteData(p2.ToBytes())
			if err != nil {
				return err
			}
			err = p.WriteHeader(0, codec.SdpType_StructEnd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type DirService struct {
	proxy model.ServicePrxImpl
}

func (s *DirService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *DirService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *DirService) CreateZone(stZoneInfo ZoneInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stZoneInfo.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("createZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) ModifyZone(stZoneInfo ZoneInfo, stModify ZoneModifyInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stZoneInfo.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	err = stModify.WriteStructFromTag(p, 2, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) DeleteZone(iZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iZoneId != 0 {
		err = p.WriteUint32(1, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deleteZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) ReportZone(iZoneId uint32, iCurOnline uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iZoneId != 0 {
		err = p.WriteUint32(1, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	if true || iCurOnline != 0 {
		err = p.WriteUint32(2, iCurOnline)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("reportZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) ReportZone2(iZoneId uint32, iCurNum uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iZoneId != 0 {
		err = p.WriteUint32(1, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	if true || iCurNum != 0 {
		err = p.WriteUint32(2, iCurNum)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("reportZone2", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) GetZone(iZoneId uint32, stZoneInfo *ZoneInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iZoneId != 0 {
		err = p.WriteUint32(1, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stZoneInfo).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *DirService) GetAllZone(vZoneInfo *[]ZoneInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAllZone", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	has, ty, err = up.SkipToTag(1, true)
	if err != nil {
		return ret, err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return ret, fmt.Errorf("tag:%d got wrong type %d", 1, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return ret, err
		}
		(*vZoneInfo) = make([]ZoneInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vZoneInfo)[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return ret, err
			}
		}
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _DirServiceImpl interface {
	CreateZone(ctx context.Context, stZoneInfo ZoneInfo) (int32, error)
	ModifyZone(ctx context.Context, stZoneInfo ZoneInfo, stModify ZoneModifyInfo) (int32, error)
	DeleteZone(ctx context.Context, iZoneId uint32) (int32, error)
	ReportZone(ctx context.Context, iZoneId uint32, iCurOnline uint32) (int32, error)
	ReportZone2(ctx context.Context, iZoneId uint32, iCurNum uint32) (int32, error)
	GetZone(ctx context.Context, iZoneId uint32, stZoneInfo *ZoneInfo) (int32, error)
	GetAllZone(ctx context.Context, vZoneInfo *[]ZoneInfo) (int32, error)
}

func _DirServiceCreateZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 ZoneInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.CreateZone(ctx, p1)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceModifyZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 ZoneInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 ZoneModifyInfo
	err = p2.ReadStructFromTag(up, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyZone(ctx, p1, p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceDeleteZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeleteZone(ctx, p1)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceReportZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ReportZone(ctx, p1, p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceReportZone2Impl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ReportZone2(ctx, p1, p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceGetZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 ZoneInfo
	var ret int32
	ret, err = impl.GetZone(ctx, p1, &p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	err = p2.WriteStructFromTag(p, 2, true)
	if err != nil {
		return err
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _DirServiceGetAllZoneImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_DirServiceImpl)
	var p1 []ZoneInfo
	var ret int32
	ret, err = impl.GetAllZone(ctx, &p1)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}

	length = uint32(len(p1))
	if true || length != 0 {
		err = p.WriteHeader(1, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range p1 {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *DirService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "createZone":
		err = _DirServiceCreateZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyZone":
		err = _DirServiceModifyZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deleteZone":
		err = _DirServiceDeleteZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "reportZone":
		err = _DirServiceReportZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "reportZone2":
		err = _DirServiceReportZone2Impl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getZone":
		err = _DirServiceGetZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAllZone":
		err = _DirServiceGetAllZoneImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	default:
		texret = protocol.SDPSERVERNOFUNCERR
	}

	if err != nil {
		log.FErrorf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d, err: %s", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId, err.Error())
	}

	if current.Rsp() {
		current.SendTexResponse(int32(texret), p.ToBytes())
	}
}
