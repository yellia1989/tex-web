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

type BulletinFlag int32

const (
	BulletinFlag_Normal   = 0
	BulletinFlag_NotBegin = 1
	BulletinFlag_HasOver  = 2
)

func (en BulletinFlag) String() string {
	ret := ""
	switch en {
	case BulletinFlag_Normal:
		ret = "BulletinFlag_Normal"
	case BulletinFlag_NotBegin:
		ret = "BulletinFlag_NotBegin"
	case BulletinFlag_HasOver:
		ret = "BulletinFlag_HasOver"
	}
	return ret
}

type BulletinDataInfo struct {
	IBulletinId       uint32 `json:"iBulletinId"`
	STitle            string `json:"sTitle"`
	SContent          string `json:"sContent"`
	IFlag             uint32 `json:"iFlag"`
	SBeginTime        string `json:"sBeginTime"`
	SEndTime          string `json:"sEndTime"`
	IDisplay          uint32 `json:"iDisplay"`
	IType             uint32 `json:"iType"`
	IPopWindow        uint32 `json:"iPopWindow"`
	SPopWindowEndTime string `json:"sPopWindowEndTime"`
}

func (st *BulletinDataInfo) resetDefault() {
}
func (st *BulletinDataInfo) Copy() *BulletinDataInfo {
	ret := NewBulletinDataInfo()
	ret.IBulletinId = st.IBulletinId
	ret.STitle = st.STitle
	ret.SContent = st.SContent
	ret.IFlag = st.IFlag
	ret.SBeginTime = st.SBeginTime
	ret.SEndTime = st.SEndTime
	ret.IDisplay = st.IDisplay
	ret.IType = st.IType
	ret.IPopWindow = st.IPopWindow
	ret.SPopWindowEndTime = st.SPopWindowEndTime
	return ret
}
func NewBulletinDataInfo() *BulletinDataInfo {
	ret := &BulletinDataInfo{}
	ret.resetDefault()
	return ret
}
func (st *BulletinDataInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iBulletinId")+fmt.Sprintf("%v\n", st.IBulletinId))
	util.Tab(buff, t+1, util.Fieldname("sTitle")+fmt.Sprintf("%v\n", st.STitle))
	util.Tab(buff, t+1, util.Fieldname("sContent")+fmt.Sprintf("%v\n", st.SContent))
	util.Tab(buff, t+1, util.Fieldname("iFlag")+fmt.Sprintf("%v\n", st.IFlag))
	util.Tab(buff, t+1, util.Fieldname("sBeginTime")+fmt.Sprintf("%v\n", st.SBeginTime))
	util.Tab(buff, t+1, util.Fieldname("sEndTime")+fmt.Sprintf("%v\n", st.SEndTime))
	util.Tab(buff, t+1, util.Fieldname("iDisplay")+fmt.Sprintf("%v\n", st.IDisplay))
	util.Tab(buff, t+1, util.Fieldname("iType")+fmt.Sprintf("%v\n", st.IType))
	util.Tab(buff, t+1, util.Fieldname("iPopWindow")+fmt.Sprintf("%v\n", st.IPopWindow))
	util.Tab(buff, t+1, util.Fieldname("sPopWindowEndTime")+fmt.Sprintf("%v\n", st.SPopWindowEndTime))
}
func (st *BulletinDataInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IBulletinId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STitle, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SContent, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IFlag, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBeginTime, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SEndTime, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDisplay, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IType, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPopWindow, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPopWindowEndTime, 10, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *BulletinDataInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *BulletinDataInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IBulletinId != 0 {
		err = p.WriteUint32(0, st.IBulletinId)
		if err != nil {
			return err
		}
	}
	if false || st.STitle != "" {
		err = p.WriteString(1, st.STitle)
		if err != nil {
			return err
		}
	}
	if false || st.SContent != "" {
		err = p.WriteString(2, st.SContent)
		if err != nil {
			return err
		}
	}
	if false || st.IFlag != 0 {
		err = p.WriteUint32(3, st.IFlag)
		if err != nil {
			return err
		}
	}
	if false || st.SBeginTime != "" {
		err = p.WriteString(5, st.SBeginTime)
		if err != nil {
			return err
		}
	}
	if false || st.SEndTime != "" {
		err = p.WriteString(6, st.SEndTime)
		if err != nil {
			return err
		}
	}
	if false || st.IDisplay != 0 {
		err = p.WriteUint32(7, st.IDisplay)
		if err != nil {
			return err
		}
	}
	if false || st.IType != 0 {
		err = p.WriteUint32(8, st.IType)
		if err != nil {
			return err
		}
	}
	if false || st.IPopWindow != 0 {
		err = p.WriteUint32(9, st.IPopWindow)
		if err != nil {
			return err
		}
	}
	if false || st.SPopWindowEndTime != "" {
		err = p.WriteString(10, st.SPopWindowEndTime)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *BulletinDataInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type NoticeDataInfo struct {
	INoticeId        uint32   `json:"iNoticeId"`
	IType            uint32   `json:"iType"`
	SContent         string   `json:"sContent"`
	SBeginTime       string   `json:"sBeginTime"`
	SEndTime         string   `json:"sEndTime"`
	IDisplayInterval uint32   `json:"iDisplayInterval"`
	IDisplayType     uint32   `json:"iDisplayType"`
	IDisplayNum      uint32   `json:"iDisplayNum"`
	IPause           uint32   `json:"iPause"`
	VZoneId          []uint32 `json:"vZoneId"`
	IMaintenanceTime uint32   `json:"iMaintenanceTime"`
}

func (st *NoticeDataInfo) resetDefault() {
}
func (st *NoticeDataInfo) Copy() *NoticeDataInfo {
	ret := NewNoticeDataInfo()
	ret.INoticeId = st.INoticeId
	ret.IType = st.IType
	ret.SContent = st.SContent
	ret.SBeginTime = st.SBeginTime
	ret.SEndTime = st.SEndTime
	ret.IDisplayInterval = st.IDisplayInterval
	ret.IDisplayType = st.IDisplayType
	ret.IDisplayNum = st.IDisplayNum
	ret.IPause = st.IPause
	ret.VZoneId = make([]uint32, len(st.VZoneId))
	for i, v := range st.VZoneId {
		ret.VZoneId[i] = v
	}
	ret.IMaintenanceTime = st.IMaintenanceTime
	return ret
}
func NewNoticeDataInfo() *NoticeDataInfo {
	ret := &NoticeDataInfo{}
	ret.resetDefault()
	return ret
}
func (st *NoticeDataInfo) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iNoticeId")+fmt.Sprintf("%v\n", st.INoticeId))
	util.Tab(buff, t+1, util.Fieldname("iType")+fmt.Sprintf("%v\n", st.IType))
	util.Tab(buff, t+1, util.Fieldname("sContent")+fmt.Sprintf("%v\n", st.SContent))
	util.Tab(buff, t+1, util.Fieldname("sBeginTime")+fmt.Sprintf("%v\n", st.SBeginTime))
	util.Tab(buff, t+1, util.Fieldname("sEndTime")+fmt.Sprintf("%v\n", st.SEndTime))
	util.Tab(buff, t+1, util.Fieldname("iDisplayInterval")+fmt.Sprintf("%v\n", st.IDisplayInterval))
	util.Tab(buff, t+1, util.Fieldname("iDisplayType")+fmt.Sprintf("%v\n", st.IDisplayType))
	util.Tab(buff, t+1, util.Fieldname("iDisplayNum")+fmt.Sprintf("%v\n", st.IDisplayNum))
	util.Tab(buff, t+1, util.Fieldname("iPause")+fmt.Sprintf("%v\n", st.IPause))
	util.Tab(buff, t+1, util.Fieldname("vZoneId")+strconv.Itoa(len(st.VZoneId)))
	if len(st.VZoneId) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VZoneId {
		util.Tab(buff, t+1+1, util.Fieldname("")+fmt.Sprintf("%v\n", v))
	}
	if len(st.VZoneId) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
	util.Tab(buff, t+1, util.Fieldname("iMaintenanceTime")+fmt.Sprintf("%v\n", st.IMaintenanceTime))
}
func (st *NoticeDataInfo) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.INoticeId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IType, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SContent, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBeginTime, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SEndTime, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDisplayInterval, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDisplayType, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDisplayNum, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPause, 8, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(9, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 9, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.VZoneId = make([]uint32, length, length)
		for i := uint32(0); i < length; i++ {
			err = up.ReadUint32(&st.VZoneId[i], 0, true)
			if err != nil {
				return err
			}
		}
	}
	err = up.ReadUint32(&st.IMaintenanceTime, 10, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *NoticeDataInfo) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *NoticeDataInfo) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.INoticeId != 0 {
		err = p.WriteUint32(0, st.INoticeId)
		if err != nil {
			return err
		}
	}
	if false || st.IType != 0 {
		err = p.WriteUint32(1, st.IType)
		if err != nil {
			return err
		}
	}
	if false || st.SContent != "" {
		err = p.WriteString(2, st.SContent)
		if err != nil {
			return err
		}
	}
	if false || st.SBeginTime != "" {
		err = p.WriteString(3, st.SBeginTime)
		if err != nil {
			return err
		}
	}
	if false || st.SEndTime != "" {
		err = p.WriteString(4, st.SEndTime)
		if err != nil {
			return err
		}
	}
	if false || st.IDisplayInterval != 0 {
		err = p.WriteUint32(5, st.IDisplayInterval)
		if err != nil {
			return err
		}
	}
	if false || st.IDisplayType != 0 {
		err = p.WriteUint32(6, st.IDisplayType)
		if err != nil {
			return err
		}
	}
	if false || st.IDisplayNum != 0 {
		err = p.WriteUint32(7, st.IDisplayNum)
		if err != nil {
			return err
		}
	}
	if false || st.IPause != 0 {
		err = p.WriteUint32(8, st.IPause)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.VZoneId))
	if false || length != 0 {
		err = p.WriteHeader(9, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VZoneId {
			if true || v != 0 {
				err = p.WriteUint32(0, v)
				if err != nil {
					return err
				}
			}
		}
	}
	if false || st.IMaintenanceTime != 0 {
		err = p.WriteUint32(10, st.IMaintenanceTime)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *NoticeDataInfo) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type BulletinService struct {
	proxy model.ServicePrxImpl
}

func (s *BulletinService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *BulletinService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *BulletinService) AddBulletin(info BulletinDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = info.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("addBulletin", p.ToBytes(), &rsp)
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
func (s *BulletinService) ModifyBulletin(info BulletinDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = info.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyBulletin", p.ToBytes(), &rsp)
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
func (s *BulletinService) GetAllBulletin(vInfo *[]BulletinDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAllBulletin", p.ToBytes(), &rsp)
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
		(*vInfo) = make([]BulletinDataInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vInfo)[i].ReadStructFromTag(up, 0, true)
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
func (s *BulletinService) DelBulletin(iBulletinId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iBulletinId != 0 {
		err = p.WriteUint32(1, iBulletinId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("delBulletin", p.ToBytes(), &rsp)
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
func (s *BulletinService) GetLatestInsertBulletin(stInfo *BulletinDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getLatestInsertBulletin", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stInfo).ReadStructFromTag(up, 1, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *BulletinService) AddNotice(info NoticeDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = info.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("addNotice", p.ToBytes(), &rsp)
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
func (s *BulletinService) ModifyNotice(info NoticeDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = info.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyNotice", p.ToBytes(), &rsp)
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
func (s *BulletinService) GetAllNotice(vInfo *[]NoticeDataInfo) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAllNotice", p.ToBytes(), &rsp)
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
		(*vInfo) = make([]NoticeDataInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vInfo)[i].ReadStructFromTag(up, 0, true)
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
func (s *BulletinService) DelNotice(iNoticeId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iNoticeId != 0 {
		err = p.WriteUint32(1, iNoticeId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("delNotice", p.ToBytes(), &rsp)
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
func (s *BulletinService) GetLatestBulletin(vInfo *[]BulletinDataInfo, bDisplay bool) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || bDisplay != false {
		err = p.WriteBool(2, bDisplay)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getLatestBulletin", p.ToBytes(), &rsp)
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
		(*vInfo) = make([]BulletinDataInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vInfo)[i].ReadStructFromTag(up, 0, true)
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
func (s *BulletinService) GetNotice(iZoneId uint32, iLastNoticeId uint32, vInfo *[]NoticeDataInfo, vDel *[]uint32) (int32, error) {
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
	if true || iLastNoticeId != 0 {
		err = p.WriteUint32(2, iLastNoticeId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getNotice", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	has, ty, err = up.SkipToTag(3, true)
	if err != nil {
		return ret, err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return ret, fmt.Errorf("tag:%d got wrong type %d", 3, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return ret, err
		}
		(*vInfo) = make([]NoticeDataInfo, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vInfo)[i].ReadStructFromTag(up, 0, true)
			if err != nil {
				return ret, err
			}
		}
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
		(*vDel) = make([]uint32, length, length)
		for i := uint32(0); i < length; i++ {
			err = up.ReadUint32(&(*vDel)[i], 0, true)
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

type _BulletinServiceImpl interface {
	AddBulletin(ctx context.Context, info BulletinDataInfo) (int32, error)
	ModifyBulletin(ctx context.Context, info BulletinDataInfo) (int32, error)
	GetAllBulletin(ctx context.Context, vInfo *[]BulletinDataInfo) (int32, error)
	DelBulletin(ctx context.Context, iBulletinId uint32) (int32, error)
	GetLatestInsertBulletin(ctx context.Context, stInfo *BulletinDataInfo) (int32, error)
	AddNotice(ctx context.Context, info NoticeDataInfo) (int32, error)
	ModifyNotice(ctx context.Context, info NoticeDataInfo) (int32, error)
	GetAllNotice(ctx context.Context, vInfo *[]NoticeDataInfo) (int32, error)
	DelNotice(ctx context.Context, iNoticeId uint32) (int32, error)
	GetLatestBulletin(ctx context.Context, vInfo *[]BulletinDataInfo, bDisplay bool) (int32, error)
	GetNotice(ctx context.Context, iZoneId uint32, iLastNoticeId uint32, vInfo *[]NoticeDataInfo, vDel *[]uint32) (int32, error)
}

func _BulletinServiceAddBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 BulletinDataInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.AddBulletin(ctx, p1)
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
func _BulletinServiceModifyBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 BulletinDataInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyBulletin(ctx, p1)
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
func _BulletinServiceGetAllBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 []BulletinDataInfo
	var ret int32
	ret, err = impl.GetAllBulletin(ctx, &p1)
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
func _BulletinServiceDelBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DelBulletin(ctx, p1)
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
func _BulletinServiceGetLatestInsertBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 BulletinDataInfo
	var ret int32
	ret, err = impl.GetLatestInsertBulletin(ctx, &p1)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	err = p1.WriteStructFromTag(p, 1, true)
	if err != nil {
		return err
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _BulletinServiceAddNoticeImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 NoticeDataInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.AddNotice(ctx, p1)
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
func _BulletinServiceModifyNoticeImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 NoticeDataInfo
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyNotice(ctx, p1)
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
func _BulletinServiceGetAllNoticeImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 []NoticeDataInfo
	var ret int32
	ret, err = impl.GetAllNotice(ctx, &p1)
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
func _BulletinServiceDelNoticeImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DelNotice(ctx, p1)
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
func _BulletinServiceGetLatestBulletinImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
	var p2 bool
	err = up.ReadBool(&p2, 2, true)
	if err != nil {
		return err
	}
	var p1 []BulletinDataInfo
	var ret int32
	ret, err = impl.GetLatestBulletin(ctx, &p1, p2)
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
func _BulletinServiceGetNoticeImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_BulletinServiceImpl)
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
	var p3 []NoticeDataInfo
	var p4 []uint32
	var ret int32
	ret, err = impl.GetNotice(ctx, p1, p2, &p3, &p4)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}

	length = uint32(len(p3))
	if true || length != 0 {
		err = p.WriteHeader(3, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range p3 {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
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
			if true || v != 0 {
				err = p.WriteUint32(0, v)
				if err != nil {
					return err
				}
			}
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *BulletinService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "addBulletin":
		err = _BulletinServiceAddBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyBulletin":
		err = _BulletinServiceModifyBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAllBulletin":
		err = _BulletinServiceGetAllBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "delBulletin":
		err = _BulletinServiceDelBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getLatestInsertBulletin":
		err = _BulletinServiceGetLatestInsertBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "addNotice":
		err = _BulletinServiceAddNoticeImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyNotice":
		err = _BulletinServiceModifyNoticeImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAllNotice":
		err = _BulletinServiceGetAllNoticeImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "delNotice":
		err = _BulletinServiceDelNoticeImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getLatestBulletin":
		err = _BulletinServiceGetLatestBulletinImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getNotice":
		err = _BulletinServiceGetNoticeImpl(ctx, serviceImpl, up, p)
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
