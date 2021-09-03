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

type CDKeyCreateMode int32

const (
	CDKeyCreateMode_PreAllocate = 1
)

func (en CDKeyCreateMode) String() string {
	ret := ""
	switch en {
	case CDKeyCreateMode_PreAllocate:
		ret = "CDKeyCreateMode_PreAllocate"
	}
	return ret
}

type CDKeyDeliveryMode int32

const (
	CDKeyDeliveryMode_Automatic = 1
)

func (en CDKeyDeliveryMode) String() string {
	ret := ""
	switch en {
	case CDKeyDeliveryMode_Automatic:
		ret = "CDKeyDeliveryMode_Automatic"
	}
	return ret
}

type MPProjectConfig struct {
	SProjectId      string `json:"sProjectId"`
	SProjectName    string `json:"sProjectName"`
	SDeliveryServer string `json:"sDeliveryServer"`
}

func (st *MPProjectConfig) resetDefault() {
}
func (st *MPProjectConfig) Copy() *MPProjectConfig {
	ret := NewMPProjectConfig()
	ret.SProjectId = st.SProjectId
	ret.SProjectName = st.SProjectName
	ret.SDeliveryServer = st.SDeliveryServer
	return ret
}
func NewMPProjectConfig() *MPProjectConfig {
	ret := &MPProjectConfig{}
	ret.resetDefault()
	return ret
}
func (st *MPProjectConfig) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sProjectId")+fmt.Sprintf("%v\n", st.SProjectId))
	util.Tab(buff, t+1, util.Fieldname("sProjectName")+fmt.Sprintf("%v\n", st.SProjectName))
	util.Tab(buff, t+1, util.Fieldname("sDeliveryServer")+fmt.Sprintf("%v\n", st.SDeliveryServer))
}
func (st *MPProjectConfig) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SProjectId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProjectName, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDeliveryServer, 2, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *MPProjectConfig) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *MPProjectConfig) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SProjectId != "" {
		err = p.WriteString(0, st.SProjectId)
		if err != nil {
			return err
		}
	}
	if false || st.SProjectName != "" {
		err = p.WriteString(1, st.SProjectName)
		if err != nil {
			return err
		}
	}
	if false || st.SDeliveryServer != "" {
		err = p.WriteString(2, st.SDeliveryServer)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *MPProjectConfig) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type CDKeyConfig struct {
	SProjectId     string `json:"sProjectId"`
	ICDKeyId       uint32 `json:"iCDKeyId"`
	SCDKeyName     string `json:"sCDKeyName"`
	ICDKeyNum      uint32 `json:"iCDKeyNum"`
	ICreateMode    uint32 `json:"iCreateMode"`
	IDeliveryMode  uint32 `json:"iDeliveryMode"`
	IBeginTime     uint32 `json:"iBeginTime"`
	IEndTime       uint32 `json:"iEndTime"`
	SRewardInfo    string `json:"sRewardInfo"`
	IExchangeLimit uint32 `json:"iExchangeLimit"`
	SZoneLimit     string `json:"sZoneLimit"`
	SCustomLimit   string `json:"sCustomLimit"`
	IGeneratedNum  uint32 `json:"iGeneratedNum"`
	IExchangedNum  uint32 `json:"iExchangedNum"`
	ICommon        uint32 `json:"iCommon"`
	SCommonCdk     string `json:"sCommonCdk"`
	IActive        uint32 `json:"iActive"`
}

func (st *CDKeyConfig) resetDefault() {
}
func (st *CDKeyConfig) Copy() *CDKeyConfig {
	ret := NewCDKeyConfig()
	ret.SProjectId = st.SProjectId
	ret.ICDKeyId = st.ICDKeyId
	ret.SCDKeyName = st.SCDKeyName
	ret.ICDKeyNum = st.ICDKeyNum
	ret.ICreateMode = st.ICreateMode
	ret.IDeliveryMode = st.IDeliveryMode
	ret.IBeginTime = st.IBeginTime
	ret.IEndTime = st.IEndTime
	ret.SRewardInfo = st.SRewardInfo
	ret.IExchangeLimit = st.IExchangeLimit
	ret.SZoneLimit = st.SZoneLimit
	ret.SCustomLimit = st.SCustomLimit
	ret.IGeneratedNum = st.IGeneratedNum
	ret.IExchangedNum = st.IExchangedNum
	ret.ICommon = st.ICommon
	ret.SCommonCdk = st.SCommonCdk
	ret.IActive = st.IActive
	return ret
}
func NewCDKeyConfig() *CDKeyConfig {
	ret := &CDKeyConfig{}
	ret.resetDefault()
	return ret
}
func (st *CDKeyConfig) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sProjectId")+fmt.Sprintf("%v\n", st.SProjectId))
	util.Tab(buff, t+1, util.Fieldname("iCDKeyId")+fmt.Sprintf("%v\n", st.ICDKeyId))
	util.Tab(buff, t+1, util.Fieldname("sCDKeyName")+fmt.Sprintf("%v\n", st.SCDKeyName))
	util.Tab(buff, t+1, util.Fieldname("iCDKeyNum")+fmt.Sprintf("%v\n", st.ICDKeyNum))
	util.Tab(buff, t+1, util.Fieldname("iCreateMode")+fmt.Sprintf("%v\n", st.ICreateMode))
	util.Tab(buff, t+1, util.Fieldname("iDeliveryMode")+fmt.Sprintf("%v\n", st.IDeliveryMode))
	util.Tab(buff, t+1, util.Fieldname("iBeginTime")+fmt.Sprintf("%v\n", st.IBeginTime))
	util.Tab(buff, t+1, util.Fieldname("iEndTime")+fmt.Sprintf("%v\n", st.IEndTime))
	util.Tab(buff, t+1, util.Fieldname("sRewardInfo")+fmt.Sprintf("%v\n", st.SRewardInfo))
	util.Tab(buff, t+1, util.Fieldname("iExchangeLimit")+fmt.Sprintf("%v\n", st.IExchangeLimit))
	util.Tab(buff, t+1, util.Fieldname("sZoneLimit")+fmt.Sprintf("%v\n", st.SZoneLimit))
	util.Tab(buff, t+1, util.Fieldname("sCustomLimit")+fmt.Sprintf("%v\n", st.SCustomLimit))
	util.Tab(buff, t+1, util.Fieldname("iGeneratedNum")+fmt.Sprintf("%v\n", st.IGeneratedNum))
	util.Tab(buff, t+1, util.Fieldname("iExchangedNum")+fmt.Sprintf("%v\n", st.IExchangedNum))
	util.Tab(buff, t+1, util.Fieldname("iCommon")+fmt.Sprintf("%v\n", st.ICommon))
	util.Tab(buff, t+1, util.Fieldname("sCommonCdk")+fmt.Sprintf("%v\n", st.SCommonCdk))
	util.Tab(buff, t+1, util.Fieldname("iActive")+fmt.Sprintf("%v\n", st.IActive))
}
func (st *CDKeyConfig) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SProjectId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICDKeyId, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCDKeyName, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICDKeyNum, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICreateMode, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDeliveryMode, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IBeginTime, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IEndTime, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRewardInfo, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IExchangeLimit, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SZoneLimit, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCustomLimit, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IGeneratedNum, 20, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IExchangedNum, 21, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICommon, 22, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCommonCdk, 23, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IActive, 24, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *CDKeyConfig) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *CDKeyConfig) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SProjectId != "" {
		err = p.WriteString(0, st.SProjectId)
		if err != nil {
			return err
		}
	}
	if false || st.ICDKeyId != 0 {
		err = p.WriteUint32(1, st.ICDKeyId)
		if err != nil {
			return err
		}
	}
	if false || st.SCDKeyName != "" {
		err = p.WriteString(2, st.SCDKeyName)
		if err != nil {
			return err
		}
	}
	if false || st.ICDKeyNum != 0 {
		err = p.WriteUint32(3, st.ICDKeyNum)
		if err != nil {
			return err
		}
	}
	if false || st.ICreateMode != 0 {
		err = p.WriteUint32(4, st.ICreateMode)
		if err != nil {
			return err
		}
	}
	if false || st.IDeliveryMode != 0 {
		err = p.WriteUint32(5, st.IDeliveryMode)
		if err != nil {
			return err
		}
	}
	if false || st.IBeginTime != 0 {
		err = p.WriteUint32(6, st.IBeginTime)
		if err != nil {
			return err
		}
	}
	if false || st.IEndTime != 0 {
		err = p.WriteUint32(7, st.IEndTime)
		if err != nil {
			return err
		}
	}
	if false || st.SRewardInfo != "" {
		err = p.WriteString(8, st.SRewardInfo)
		if err != nil {
			return err
		}
	}
	if false || st.IExchangeLimit != 0 {
		err = p.WriteUint32(9, st.IExchangeLimit)
		if err != nil {
			return err
		}
	}
	if false || st.SZoneLimit != "" {
		err = p.WriteString(10, st.SZoneLimit)
		if err != nil {
			return err
		}
	}
	if false || st.SCustomLimit != "" {
		err = p.WriteString(11, st.SCustomLimit)
		if err != nil {
			return err
		}
	}
	if false || st.IGeneratedNum != 0 {
		err = p.WriteUint32(20, st.IGeneratedNum)
		if err != nil {
			return err
		}
	}
	if false || st.IExchangedNum != 0 {
		err = p.WriteUint32(21, st.IExchangedNum)
		if err != nil {
			return err
		}
	}
	if false || st.ICommon != 0 {
		err = p.WriteUint32(22, st.ICommon)
		if err != nil {
			return err
		}
	}
	if false || st.SCommonCdk != "" {
		err = p.WriteString(23, st.SCommonCdk)
		if err != nil {
			return err
		}
	}
	if false || st.IActive != 0 {
		err = p.WriteUint32(24, st.IActive)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *CDKeyConfig) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type CDKeyInfo struct {
	SCDKey           string `json:"sCDKey"`
	SBindAccount     string `json:"sBindAccount"`
	IExchangeTime    uint32 `json:"iExchangeTime"`
	SExchangeAccount string `json:"sExchangeAccount"`
}

func (st *CDKeyInfo) resetDefault() {
}
func (st *CDKeyInfo) Copy() *CDKeyInfo {
	ret := NewCDKeyInfo()
	ret.SCDKey = st.SCDKey
	ret.SBindAccount = st.SBindAccount
	ret.IExchangeTime = st.IExchangeTime
	ret.SExchangeAccount = st.SExchangeAccount
	return ret
}
func NewCDKeyInfo() *CDKeyInfo {
	ret := &CDKeyInfo{}
	ret.resetDefault()
	return ret
}
func (st *CDKeyInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sCDKey")+fmt.Sprintf("%v\n", st.SCDKey))
	util.Tab(buff, t+1, util.Fieldname("sBindAccount")+fmt.Sprintf("%v\n", st.SBindAccount))
	util.Tab(buff, t+1, util.Fieldname("iExchangeTime")+fmt.Sprintf("%v\n", st.IExchangeTime))
	util.Tab(buff, t+1, util.Fieldname("sExchangeAccount")+fmt.Sprintf("%v\n", st.SExchangeAccount))
}
func (st *CDKeyInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SCDKey, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBindAccount, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IExchangeTime, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SExchangeAccount, 3, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *CDKeyInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *CDKeyInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SCDKey != "" {
		err = p.WriteString(0, st.SCDKey)
		if err != nil {
			return err
		}
	}
	if false || st.SBindAccount != "" {
		err = p.WriteString(1, st.SBindAccount)
		if err != nil {
			return err
		}
	}
	if false || st.IExchangeTime != 0 {
		err = p.WriteUint32(2, st.IExchangeTime)
		if err != nil {
			return err
		}
	}
	if false || st.SExchangeAccount != "" {
		err = p.WriteString(3, st.SExchangeAccount)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *CDKeyInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type ExchangedCDKeyInfo struct {
	SCDKey        string `json:"sCDKey"`
	IExchangeTime uint32 `json:"iExchangeTime"`
}

func (st *ExchangedCDKeyInfo) resetDefault() {
}
func (st *ExchangedCDKeyInfo) Copy() *ExchangedCDKeyInfo {
	ret := NewExchangedCDKeyInfo()
	ret.SCDKey = st.SCDKey
	ret.IExchangeTime = st.IExchangeTime
	return ret
}
func NewExchangedCDKeyInfo() *ExchangedCDKeyInfo {
	ret := &ExchangedCDKeyInfo{}
	ret.resetDefault()
	return ret
}
func (st *ExchangedCDKeyInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sCDKey")+fmt.Sprintf("%v\n", st.SCDKey))
	util.Tab(buff, t+1, util.Fieldname("iExchangeTime")+fmt.Sprintf("%v\n", st.IExchangeTime))
}
func (st *ExchangedCDKeyInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SCDKey, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IExchangeTime, 1, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ExchangedCDKeyInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ExchangedCDKeyInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SCDKey != "" {
		err = p.WriteString(0, st.SCDKey)
		if err != nil {
			return err
		}
	}
	if false || st.IExchangeTime != 0 {
		err = p.WriteUint32(1, st.IExchangeTime)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ExchangedCDKeyInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type AccountExchangeInfo struct {
	SAccount        string               `json:"sAccount"`
	ICDKeyId        uint32               `json:"iCDKeyId"`
	VExchangedCDKey []ExchangedCDKeyInfo `json:"vExchangedCDKey"`
}

func (st *AccountExchangeInfo) resetDefault() {
}
func (st *AccountExchangeInfo) Copy() *AccountExchangeInfo {
	ret := NewAccountExchangeInfo()
	ret.SAccount = st.SAccount
	ret.ICDKeyId = st.ICDKeyId
	ret.VExchangedCDKey = make([]ExchangedCDKeyInfo, len(st.VExchangedCDKey))
	for i, v := range st.VExchangedCDKey {
		ret.VExchangedCDKey[i] = *(v.Copy())
	}
	return ret
}
func NewAccountExchangeInfo() *AccountExchangeInfo {
	ret := &AccountExchangeInfo{}
	ret.resetDefault()
	return ret
}
func (st *AccountExchangeInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sAccount")+fmt.Sprintf("%v\n", st.SAccount))
	util.Tab(buff, t+1, util.Fieldname("iCDKeyId")+fmt.Sprintf("%v\n", st.ICDKeyId))
	util.Tab(buff, t+1, util.Fieldname("vExchangedCDKey")+strconv.Itoa(len(st.VExchangedCDKey)))
	if len(st.VExchangedCDKey) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VExchangedCDKey {
		util.Tab(buff, t+1+1, util.Fieldname("")+"{\n")
		v.Visit(buff, t+1+1+1)
		util.Tab(buff, t+1+1, "}\n")
	}
	if len(st.VExchangedCDKey) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
}
func (st *AccountExchangeInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SAccount, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICDKeyId, 1, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(2, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 2, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.VExchangedCDKey = make([]ExchangedCDKeyInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = st.VExchangedCDKey[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return err
			}
		}
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *AccountExchangeInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *AccountExchangeInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SAccount != "" {
		err = p.WriteString(0, st.SAccount)
		if err != nil {
			return err
		}
	}
	if false || st.ICDKeyId != 0 {
		err = p.WriteUint32(1, st.ICDKeyId)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.VExchangedCDKey))
	if false || length != 0 {
		err = p.WriteHeader(2, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VExchangedCDKey {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}

	_ = length
	return err
}
func (st *AccountExchangeInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type MPService struct {
	proxy model.ServicePrxImpl
}

func (s *MPService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *MPService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *MPService) CreateProject(stProjectConfig MPProjectConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stProjectConfig.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("createProject", p.ToBytes(), &rsp)
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
func (s *MPService) ModifyProject(stProjectConfig MPProjectConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stProjectConfig.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyProject", p.ToBytes(), &rsp)
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
func (s *MPService) GetProject(sProjectId string, stProjectConfig *MPProjectConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sProjectId != "" {
		err = p.WriteString(1, sProjectId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getProject", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stProjectConfig).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) GetProjectList(vProjectConfig *[]MPProjectConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getProjectList", p.ToBytes(), &rsp)
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
		(*vProjectConfig) = make([]MPProjectConfig, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vProjectConfig)[i].ReadStructFromTag(up, 0, true)
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
func (s *MPService) CreateCDKey(stCDKeyConfig CDKeyConfig, iCDKeyId *uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stCDKeyConfig.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("createCDKey", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadUint32(&(*iCDKeyId), 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) ModifyCDKey(stCDKeyConfig CDKeyConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stCDKeyConfig.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyCDKey", p.ToBytes(), &rsp)
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
func (s *MPService) GetCDKey(iCDKeyId uint32, stCDKeyConfig *CDKeyConfig) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iCDKeyId != 0 {
		err = p.WriteUint32(1, iCDKeyId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getCDKey", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stCDKeyConfig).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) GetCDKeyList(sProjectId string, iOffset uint32, iNum uint32, vCDKeyConfig *[]CDKeyConfig, iTotalNum *uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sProjectId != "" {
		err = p.WriteString(1, sProjectId)
		if err != nil {
			return ret, err
		}
	}
	if true || iOffset != 0 {
		err = p.WriteUint32(2, iOffset)
		if err != nil {
			return ret, err
		}
	}
	if true || iNum != 0 {
		err = p.WriteUint32(3, iNum)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getCDKeyList", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	has, ty, err = up.SkipToTag(4, true)
	if err != nil {
		return ret, err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return ret, fmt.Errorf("tag:%d got wrong type %d", 4, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return ret, err
		}
		(*vCDKeyConfig) = make([]CDKeyConfig, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vCDKeyConfig)[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return ret, err
			}
		}
	}
	err = up.ReadUint32(&(*iTotalNum), 5, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) GetCDKeyInfo(sCDKey string, stCDKeyConfig *CDKeyConfig, stCDKeyInfo *CDKeyInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sCDKey != "" {
		err = p.WriteString(1, sCDKey)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getCDKeyInfo", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stCDKeyConfig).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = (*stCDKeyInfo).ReadStructFromTag(up, 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) ExchangeCDKey(sAccount string, sApp string, sDivision string, sCDKey string, bActive bool) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sAccount != "" {
		err = p.WriteString(1, sAccount)
		if err != nil {
			return ret, err
		}
	}
	if true || sApp != "" {
		err = p.WriteString(2, sApp)
		if err != nil {
			return ret, err
		}
	}
	if true || sDivision != "" {
		err = p.WriteString(3, sDivision)
		if err != nil {
			return ret, err
		}
	}
	if true || sCDKey != "" {
		err = p.WriteString(4, sCDKey)
		if err != nil {
			return ret, err
		}
	}
	if true || bActive != false {
		err = p.WriteBool(5, bActive)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("exchangeCDKey", p.ToBytes(), &rsp)
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
func (s *MPService) GetAccountExchangeInfo(sAccount string, iCDKeyId uint32, stAccountExchangeInfo *AccountExchangeInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sAccount != "" {
		err = p.WriteString(1, sAccount)
		if err != nil {
			return ret, err
		}
	}
	if true || iCDKeyId != 0 {
		err = p.WriteUint32(2, iCDKeyId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAccountExchangeInfo", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stAccountExchangeInfo).ReadStructFromTag(up, 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *MPService) ExportCDKey(iCDKeyId uint32, sAllCDKeys *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iCDKeyId != 0 {
		err = p.WriteUint32(1, iCDKeyId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("exportCDKey", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sAllCDKeys), 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _MPServiceImpl interface {
	CreateProject(ctx context.Context, stProjectConfig MPProjectConfig) (int32, error)
	ModifyProject(ctx context.Context, stProjectConfig MPProjectConfig) (int32, error)
	GetProject(ctx context.Context, sProjectId string, stProjectConfig *MPProjectConfig) (int32, error)
	GetProjectList(ctx context.Context, vProjectConfig *[]MPProjectConfig) (int32, error)
	CreateCDKey(ctx context.Context, stCDKeyConfig CDKeyConfig, iCDKeyId *uint32) (int32, error)
	ModifyCDKey(ctx context.Context, stCDKeyConfig CDKeyConfig) (int32, error)
	GetCDKey(ctx context.Context, iCDKeyId uint32, stCDKeyConfig *CDKeyConfig) (int32, error)
	GetCDKeyList(ctx context.Context, sProjectId string, iOffset uint32, iNum uint32, vCDKeyConfig *[]CDKeyConfig, iTotalNum *uint32) (int32, error)
	GetCDKeyInfo(ctx context.Context, sCDKey string, stCDKeyConfig *CDKeyConfig, stCDKeyInfo *CDKeyInfo) (int32, error)
	ExchangeCDKey(ctx context.Context, sAccount string, sApp string, sDivision string, sCDKey string, bActive bool) (int32, error)
	GetAccountExchangeInfo(ctx context.Context, sAccount string, iCDKeyId uint32, stAccountExchangeInfo *AccountExchangeInfo) (int32, error)
	ExportCDKey(ctx context.Context, iCDKeyId uint32, sAllCDKeys *string) (int32, error)
}

func _MPServiceCreateProjectImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 MPProjectConfig
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.CreateProject(ctx, p1)
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
func _MPServiceModifyProjectImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 MPProjectConfig
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyProject(ctx, p1)
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
func _MPServiceGetProjectImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 MPProjectConfig
	var ret int32
	ret, err = impl.GetProject(ctx, p1, &p2)
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
func _MPServiceGetProjectListImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 []MPProjectConfig
	var ret int32
	ret, err = impl.GetProjectList(ctx, &p1)
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
func _MPServiceCreateCDKeyImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 CDKeyConfig
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	var ret int32
	ret, err = impl.CreateCDKey(ctx, p1, &p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	if true || p2 != 0 {
		err = p.WriteUint32(2, p2)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _MPServiceModifyCDKeyImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 CDKeyConfig
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyCDKey(ctx, p1)
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
func _MPServiceGetCDKeyImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 CDKeyConfig
	var ret int32
	ret, err = impl.GetCDKey(ctx, p1, &p2)
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
func _MPServiceGetCDKeyListImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 uint32
	err = up.ReadUint32(&p3, 3, true)
	if err != nil {
		return err
	}
	var p4 []CDKeyConfig
	var p5 uint32
	var ret int32
	ret, err = impl.GetCDKeyList(ctx, p1, p2, p3, &p4, &p5)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}

	length = uint32(len(p4))
	if true || length != 0 {
		err = p.WriteHeader(4, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range p4 {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}
	if true || p5 != 0 {
		err = p.WriteUint32(5, p5)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _MPServiceGetCDKeyInfoImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 CDKeyConfig
	var p3 CDKeyInfo
	var ret int32
	ret, err = impl.GetCDKeyInfo(ctx, p1, &p2, &p3)
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
	err = p3.WriteStructFromTag(p, 3, true)
	if err != nil {
		return err
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _MPServiceExchangeCDKeyImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 string
	err = up.ReadString(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 string
	err = up.ReadString(&p3, 3, true)
	if err != nil {
		return err
	}
	var p4 string
	err = up.ReadString(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 bool
	err = up.ReadBool(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ExchangeCDKey(ctx, p1, p2, p3, p4, p5)
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
func _MPServiceGetAccountExchangeInfoImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 AccountExchangeInfo
	var ret int32
	ret, err = impl.GetAccountExchangeInfo(ctx, p1, p2, &p3)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	err = p3.WriteStructFromTag(p, 3, true)
	if err != nil {
		return err
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _MPServiceExportCDKeyImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MPServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 string
	var ret int32
	ret, err = impl.ExportCDKey(ctx, p1, &p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	if true || p2 != "" {
		err = p.WriteString(2, p2)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *MPService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "createProject":
		err = _MPServiceCreateProjectImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyProject":
		err = _MPServiceModifyProjectImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getProject":
		err = _MPServiceGetProjectImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getProjectList":
		err = _MPServiceGetProjectListImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "createCDKey":
		err = _MPServiceCreateCDKeyImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyCDKey":
		err = _MPServiceModifyCDKeyImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getCDKey":
		err = _MPServiceGetCDKeyImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getCDKeyList":
		err = _MPServiceGetCDKeyListImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getCDKeyInfo":
		err = _MPServiceGetCDKeyInfoImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "exchangeCDKey":
		err = _MPServiceExchangeCDKeyImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAccountExchangeInfo":
		err = _MPServiceGetAccountExchangeInfoImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "exportCDKey":
		err = _MPServiceExportCDKeyImpl(ctx, serviceImpl, up, p)
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
