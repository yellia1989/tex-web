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

type PushTaskStatus int32

const (
	PushTaskStatus_Resolving = 1
	PushTaskStatus_Sorting   = 2
	PushTaskStatus_Pushing   = 3
	PushTaskStatus_Finished  = 4
	PushTaskStatus_Paused    = 5
)

func (en PushTaskStatus) String() string {
	ret := ""
	switch en {
	case PushTaskStatus_Resolving:
		ret = "PushTaskStatus_Resolving"
	case PushTaskStatus_Sorting:
		ret = "PushTaskStatus_Sorting"
	case PushTaskStatus_Pushing:
		ret = "PushTaskStatus_Pushing"
	case PushTaskStatus_Finished:
		ret = "PushTaskStatus_Finished"
	case PushTaskStatus_Paused:
		ret = "PushTaskStatus_Paused"
	}
	return ret
}

type PushTargetAccountInfo struct {
	IAccountId      uint64   `json:"iAccountId" form:"iAccountId"`
	VDeviceToken    []string `json:"vDeviceToken" form:"vDeviceToken"`
	VRegistrationId []string `json:"vRegistrationId" form:"vRegistrationId"`
}

func (st *PushTargetAccountInfo) resetDefault() {
}
func (st *PushTargetAccountInfo) Copy() *PushTargetAccountInfo {
	ret := NewPushTargetAccountInfo()
	ret.IAccountId = st.IAccountId
	ret.VDeviceToken = make([]string, len(st.VDeviceToken))
	for i, v := range st.VDeviceToken {
		ret.VDeviceToken[i] = v
	}
	ret.VRegistrationId = make([]string, len(st.VRegistrationId))
	for i, v := range st.VRegistrationId {
		ret.VRegistrationId[i] = v
	}
	return ret
}
func NewPushTargetAccountInfo() *PushTargetAccountInfo {
	ret := &PushTargetAccountInfo{}
	ret.resetDefault()
	return ret
}
func (st *PushTargetAccountInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iAccountId")+fmt.Sprintf("%v\n", st.IAccountId))
	util.Tab(buff, t+1, util.Fieldname("vDeviceToken")+strconv.Itoa(len(st.VDeviceToken)))
	if len(st.VDeviceToken) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VDeviceToken {
		util.Tab(buff, t+1+1, util.Fieldname("")+fmt.Sprintf("%v\n", v))
	}
	if len(st.VDeviceToken) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
	util.Tab(buff, t+1, util.Fieldname("vRegistrationId")+strconv.Itoa(len(st.VRegistrationId)))
	if len(st.VRegistrationId) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VRegistrationId {
		util.Tab(buff, t+1+1, util.Fieldname("")+fmt.Sprintf("%v\n", v))
	}
	if len(st.VRegistrationId) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
}
func (st *PushTargetAccountInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint64(&st.IAccountId, 0, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(1, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 1, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.VDeviceToken = make([]string, length, length)
		for i := uint32(0); i < length; i++ {
			err = up.ReadString(&st.VDeviceToken[i], 0, true)
			if err != nil {
				return err
			}
		}
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
		st.VRegistrationId = make([]string, length, length)
		for i := uint32(0); i < length; i++ {
			err = up.ReadString(&st.VRegistrationId[i], 0, true)
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
func (st *PushTargetAccountInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PushTargetAccountInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IAccountId != 0 {
		err = p.WriteUint64(0, st.IAccountId)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.VDeviceToken))
	if false || length != 0 {
		err = p.WriteHeader(1, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VDeviceToken {
			if true || v != "" {
				err = p.WriteString(0, v)
				if err != nil {
					return err
				}
			}
		}
	}

	length = uint32(len(st.VRegistrationId))
	if false || length != 0 {
		err = p.WriteHeader(2, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VRegistrationId {
			if true || v != "" {
				err = p.WriteString(0, v)
				if err != nil {
					return err
				}
			}
		}
	}

	_ = length
	return err
}
func (st *PushTargetAccountInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PushTargetDeviceInfo struct {
	IAccountId      uint64 `json:"iAccountId" form:"iAccountId"`
	SDeviceToken    string `json:"sDeviceToken" form:"sDeviceToken"`
	SRegistrationId string `json:"sRegistrationId" form:"sRegistrationId"`
}

func (st *PushTargetDeviceInfo) resetDefault() {
}
func (st *PushTargetDeviceInfo) Copy() *PushTargetDeviceInfo {
	ret := NewPushTargetDeviceInfo()
	ret.IAccountId = st.IAccountId
	ret.SDeviceToken = st.SDeviceToken
	ret.SRegistrationId = st.SRegistrationId
	return ret
}
func NewPushTargetDeviceInfo() *PushTargetDeviceInfo {
	ret := &PushTargetDeviceInfo{}
	ret.resetDefault()
	return ret
}
func (st *PushTargetDeviceInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iAccountId")+fmt.Sprintf("%v\n", st.IAccountId))
	util.Tab(buff, t+1, util.Fieldname("sDeviceToken")+fmt.Sprintf("%v\n", st.SDeviceToken))
	util.Tab(buff, t+1, util.Fieldname("sRegistrationId")+fmt.Sprintf("%v\n", st.SRegistrationId))
}
func (st *PushTargetDeviceInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint64(&st.IAccountId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDeviceToken, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRegistrationId, 2, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *PushTargetDeviceInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PushTargetDeviceInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IAccountId != 0 {
		err = p.WriteUint64(0, st.IAccountId)
		if err != nil {
			return err
		}
	}
	if false || st.SDeviceToken != "" {
		err = p.WriteString(1, st.SDeviceToken)
		if err != nil {
			return err
		}
	}
	if false || st.SRegistrationId != "" {
		err = p.WriteString(2, st.SRegistrationId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *PushTargetDeviceInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PushPayloadInfo struct {
	SApplePayload  string `json:"sApplePayload" form:"sApplePayload"`
	SGooglePayload string `json:"sGooglePayload" form:"sGooglePayload"`
	SUPushPayload  string `json:"sUPushPayload" form:"sUPushPayload"`
}

func (st *PushPayloadInfo) resetDefault() {
}
func (st *PushPayloadInfo) Copy() *PushPayloadInfo {
	ret := NewPushPayloadInfo()
	ret.SApplePayload = st.SApplePayload
	ret.SGooglePayload = st.SGooglePayload
	ret.SUPushPayload = st.SUPushPayload
	return ret
}
func NewPushPayloadInfo() *PushPayloadInfo {
	ret := &PushPayloadInfo{}
	ret.resetDefault()
	return ret
}
func (st *PushPayloadInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sApplePayload")+fmt.Sprintf("%v\n", st.SApplePayload))
	util.Tab(buff, t+1, util.Fieldname("sGooglePayload")+fmt.Sprintf("%v\n", st.SGooglePayload))
	util.Tab(buff, t+1, util.Fieldname("sUPushPayload")+fmt.Sprintf("%v\n", st.SUPushPayload))
}
func (st *PushPayloadInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SApplePayload, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SGooglePayload, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SUPushPayload, 2, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *PushPayloadInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PushPayloadInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SApplePayload != "" {
		err = p.WriteString(0, st.SApplePayload)
		if err != nil {
			return err
		}
	}
	if false || st.SGooglePayload != "" {
		err = p.WriteString(1, st.SGooglePayload)
		if err != nil {
			return err
		}
	}
	if false || st.SUPushPayload != "" {
		err = p.WriteString(2, st.SUPushPayload)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *PushPayloadInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PushTaskInfo struct {
	ITaskId          uint32          `json:"iTaskId" form:"iTaskId"`
	STaskName        string          `json:"sTaskName" form:"sTaskName"`
	StPayload        PushPayloadInfo `json:"stPayload" form:"stPayload"`
	IAddTime         uint32          `json:"iAddTime" form:"iAddTime"`
	IFinishTime      uint32          `json:"iFinishTime" form:"iFinishTime"`
	IStatus          uint32          `json:"iStatus" form:"iStatus"`
	IResolveTotalNum uint32          `json:"iResolveTotalNum" form:"iResolveTotalNum"`
	IResolveDoneNum  uint32          `json:"iResolveDoneNum" form:"iResolveDoneNum"`
	IResolveFailNum  uint32          `json:"iResolveFailNum" form:"iResolveFailNum"`
	IPushTotalNum    uint32          `json:"iPushTotalNum" form:"iPushTotalNum"`
	IPushDoneNum     uint32          `json:"iPushDoneNum" form:"iPushDoneNum"`
	IPushFailNum     uint32          `json:"iPushFailNum" form:"iPushFailNum"`
}

func (st *PushTaskInfo) resetDefault() {
	st.StPayload.resetDefault()
}
func (st *PushTaskInfo) Copy() *PushTaskInfo {
	ret := NewPushTaskInfo()
	ret.ITaskId = st.ITaskId
	ret.STaskName = st.STaskName
	ret.StPayload = *(st.StPayload.Copy())
	ret.IAddTime = st.IAddTime
	ret.IFinishTime = st.IFinishTime
	ret.IStatus = st.IStatus
	ret.IResolveTotalNum = st.IResolveTotalNum
	ret.IResolveDoneNum = st.IResolveDoneNum
	ret.IResolveFailNum = st.IResolveFailNum
	ret.IPushTotalNum = st.IPushTotalNum
	ret.IPushDoneNum = st.IPushDoneNum
	ret.IPushFailNum = st.IPushFailNum
	return ret
}
func NewPushTaskInfo() *PushTaskInfo {
	ret := &PushTaskInfo{}
	ret.resetDefault()
	return ret
}
func (st *PushTaskInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iTaskId")+fmt.Sprintf("%v\n", st.ITaskId))
	util.Tab(buff, t+1, util.Fieldname("sTaskName")+fmt.Sprintf("%v\n", st.STaskName))
	util.Tab(buff, t+1, util.Fieldname("stPayload")+"{\n")
	st.StPayload.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("iAddTime")+fmt.Sprintf("%v\n", st.IAddTime))
	util.Tab(buff, t+1, util.Fieldname("iFinishTime")+fmt.Sprintf("%v\n", st.IFinishTime))
	util.Tab(buff, t+1, util.Fieldname("iStatus")+fmt.Sprintf("%v\n", st.IStatus))
	util.Tab(buff, t+1, util.Fieldname("iResolveTotalNum")+fmt.Sprintf("%v\n", st.IResolveTotalNum))
	util.Tab(buff, t+1, util.Fieldname("iResolveDoneNum")+fmt.Sprintf("%v\n", st.IResolveDoneNum))
	util.Tab(buff, t+1, util.Fieldname("iResolveFailNum")+fmt.Sprintf("%v\n", st.IResolveFailNum))
	util.Tab(buff, t+1, util.Fieldname("iPushTotalNum")+fmt.Sprintf("%v\n", st.IPushTotalNum))
	util.Tab(buff, t+1, util.Fieldname("iPushDoneNum")+fmt.Sprintf("%v\n", st.IPushDoneNum))
	util.Tab(buff, t+1, util.Fieldname("iPushFailNum")+fmt.Sprintf("%v\n", st.IPushFailNum))
}
func (st *PushTaskInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.ITaskId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STaskName, 1, false)
	if err != nil {
		return err
	}
	err = st.StPayload.ReadStructFromTag(up, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IAddTime, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IFinishTime, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IStatus, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IResolveTotalNum, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IResolveDoneNum, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IResolveFailNum, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPushTotalNum, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPushDoneNum, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPushFailNum, 11, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *PushTaskInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PushTaskInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.ITaskId != 0 {
		err = p.WriteUint32(0, st.ITaskId)
		if err != nil {
			return err
		}
	}
	if false || st.STaskName != "" {
		err = p.WriteString(1, st.STaskName)
		if err != nil {
			return err
		}
	}
	err = st.StPayload.WriteStructFromTag(p, 2, false)
	if err != nil {
		return err
	}
	if false || st.IAddTime != 0 {
		err = p.WriteUint32(3, st.IAddTime)
		if err != nil {
			return err
		}
	}
	if false || st.IFinishTime != 0 {
		err = p.WriteUint32(4, st.IFinishTime)
		if err != nil {
			return err
		}
	}
	if false || st.IStatus != 0 {
		err = p.WriteUint32(5, st.IStatus)
		if err != nil {
			return err
		}
	}
	if false || st.IResolveTotalNum != 0 {
		err = p.WriteUint32(6, st.IResolveTotalNum)
		if err != nil {
			return err
		}
	}
	if false || st.IResolveDoneNum != 0 {
		err = p.WriteUint32(7, st.IResolveDoneNum)
		if err != nil {
			return err
		}
	}
	if false || st.IResolveFailNum != 0 {
		err = p.WriteUint32(8, st.IResolveFailNum)
		if err != nil {
			return err
		}
	}
	if false || st.IPushTotalNum != 0 {
		err = p.WriteUint32(9, st.IPushTotalNum)
		if err != nil {
			return err
		}
	}
	if false || st.IPushDoneNum != 0 {
		err = p.WriteUint32(10, st.IPushDoneNum)
		if err != nil {
			return err
		}
	}
	if false || st.IPushFailNum != 0 {
		err = p.WriteUint32(11, st.IPushFailNum)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *PushTaskInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PushService struct {
	proxy model.ServicePrxImpl
}

func (s *PushService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *PushService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *PushService) Push(stTargetAccount PushTargetAccountInfo, stPayload PushPayloadInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stTargetAccount.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	err = stPayload.WriteStructFromTag(p, 2, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("push", p.ToBytes(), &rsp)
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
func (s *PushService) PushBatch(vTargetAccount []PushTargetAccountInfo, stPayload PushPayloadInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32

	length = uint32(len(vTargetAccount))
	if true || length != 0 {
		err = p.WriteHeader(1, codec.SdpType_Vector)
		if err != nil {
			return ret, err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return ret, err
		}
		for _, v := range vTargetAccount {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return ret, err
			}
		}
	}
	err = stPayload.WriteStructFromTag(p, 2, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("pushBatch", p.ToBytes(), &rsp)
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
func (s *PushService) AddPushTask(vPushTargetAccount []PushTargetAccountInfo, sTaskName string, stPayload PushPayloadInfo, iTaskId *uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32

	length = uint32(len(vPushTargetAccount))
	if true || length != 0 {
		err = p.WriteHeader(1, codec.SdpType_Vector)
		if err != nil {
			return ret, err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return ret, err
		}
		for _, v := range vPushTargetAccount {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return ret, err
			}
		}
	}
	if true || sTaskName != "" {
		err = p.WriteString(2, sTaskName)
		if err != nil {
			return ret, err
		}
	}
	err = stPayload.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("addPushTask", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadUint32(&(*iTaskId), 4, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *PushService) GetPushTaskInfo(iTaskId uint32, stPushTaskInfo *PushTaskInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iTaskId != 0 {
		err = p.WriteUint32(1, iTaskId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getPushTaskInfo", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stPushTaskInfo).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *PushService) GetAllPushTaskInfo(vPushTaskInfo *[]PushTaskInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAllPushTaskInfo", p.ToBytes(), &rsp)
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
		(*vPushTaskInfo) = make([]PushTaskInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vPushTaskInfo)[i].ReadStructFromTag(up, 0, true)
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
func (s *PushService) PausePushTask(iTaskId uint32, bPause bool) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iTaskId != 0 {
		err = p.WriteUint32(1, iTaskId)
		if err != nil {
			return ret, err
		}
	}
	if true || bPause != false {
		err = p.WriteBool(2, bPause)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("pausePushTask", p.ToBytes(), &rsp)
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

type _PushServiceImpl interface {
	Push(ctx context.Context, stTargetAccount PushTargetAccountInfo, stPayload PushPayloadInfo) (int32, error)
	PushBatch(ctx context.Context, vTargetAccount []PushTargetAccountInfo, stPayload PushPayloadInfo) (int32, error)
	AddPushTask(ctx context.Context, vPushTargetAccount []PushTargetAccountInfo, sTaskName string, stPayload PushPayloadInfo, iTaskId *uint32) (int32, error)
	GetPushTaskInfo(ctx context.Context, iTaskId uint32, stPushTaskInfo *PushTaskInfo) (int32, error)
	GetAllPushTaskInfo(ctx context.Context, vPushTaskInfo *[]PushTaskInfo) (int32, error)
	PausePushTask(ctx context.Context, iTaskId uint32, bPause bool) (int32, error)
}

func _PushServicePushImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 PushTargetAccountInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 PushPayloadInfo
	err = p2.ReadStructFromTag(up, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.Push(ctx, p1, p2)
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
func _PushServicePushBatchImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 []PushTargetAccountInfo

	has, ty, err = up.SkipToTag(1, true)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 1, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		p1 = make([]PushTargetAccountInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = p1[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return err
			}
		}
	}
	var p2 PushPayloadInfo
	err = p2.ReadStructFromTag(up, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.PushBatch(ctx, p1, p2)
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
func _PushServiceAddPushTaskImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 []PushTargetAccountInfo

	has, ty, err = up.SkipToTag(1, true)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 1, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		p1 = make([]PushTargetAccountInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = p1[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return err
			}
		}
	}
	var p2 string
	err = up.ReadString(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 PushPayloadInfo
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint32
	var ret int32
	ret, err = impl.AddPushTask(ctx, p1, p2, p3, &p4)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	if true || p4 != 0 {
		err = p.WriteUint32(4, p4)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _PushServiceGetPushTaskInfoImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 PushTaskInfo
	var ret int32
	ret, err = impl.GetPushTaskInfo(ctx, p1, &p2)
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
func _PushServiceGetAllPushTaskInfoImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 []PushTaskInfo
	var ret int32
	ret, err = impl.GetAllPushTaskInfo(ctx, &p1)
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
func _PushServicePausePushTaskImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PushServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 bool
	err = up.ReadBool(&p2, 2, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.PausePushTask(ctx, p1, p2)
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

func (s *PushService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "push":
		err = _PushServicePushImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "pushBatch":
		err = _PushServicePushBatchImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "addPushTask":
		err = _PushServiceAddPushTaskImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getPushTaskInfo":
		err = _PushServiceGetPushTaskInfoImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAllPushTaskInfo":
		err = _PushServiceGetAllPushTaskInfoImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "pausePushTask":
		err = _PushServicePausePushTaskImpl(ctx, serviceImpl, up, p)
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
