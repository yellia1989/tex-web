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

type PatchTaskStatus int32

const (
	PatchTaskStatus_NOT_START = 0
	PatchTaskStatus_RUNNING   = 1
	PatchTaskStatus_SUCCESS   = 2
	PatchTaskStatus_FAILED    = 3
	PatchTaskStatus_CANCEL    = 4
)

func (en PatchTaskStatus) String() string {
	ret := ""
	switch en {
	case PatchTaskStatus_NOT_START:
		ret = "PatchTaskStatus_NOT_START"
	case PatchTaskStatus_RUNNING:
		ret = "PatchTaskStatus_RUNNING"
	case PatchTaskStatus_SUCCESS:
		ret = "PatchTaskStatus_SUCCESS"
	case PatchTaskStatus_FAILED:
		ret = "PatchTaskStatus_FAILED"
	case PatchTaskStatus_CANCEL:
		ret = "PatchTaskStatus_CANCEL"
	}
	return ret
}

type PatchTaskItemReq struct {
	STaskNo   string            `json:"sTaskNo" form:"sTaskNo"`
	SApp      string            `json:"sApp" form:"sApp"`
	SServer   string            `json:"sServer" form:"sServer"`
	SDivision string            `json:"sDivision" form:"sDivision"`
	SNodeName string            `json:"sNodeName" form:"sNodeName"`
	SCommand  string            `json:"sCommand" form:"sCommand"`
	MParam    map[string]string `json:"mParam" form:"mParam"`
}

func (st *PatchTaskItemReq) resetDefault() {
}
func (st *PatchTaskItemReq) Copy() *PatchTaskItemReq {
	ret := NewPatchTaskItemReq()
	ret.STaskNo = st.STaskNo
	ret.SApp = st.SApp
	ret.SServer = st.SServer
	ret.SDivision = st.SDivision
	ret.SNodeName = st.SNodeName
	ret.SCommand = st.SCommand
	ret.MParam = make(map[string]string)
	for k, v := range st.MParam {
		ret.MParam[k] = v
	}
	return ret
}
func NewPatchTaskItemReq() *PatchTaskItemReq {
	ret := &PatchTaskItemReq{}
	ret.resetDefault()
	return ret
}
func (st *PatchTaskItemReq) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sTaskNo")+fmt.Sprintf("%v\n", st.STaskNo))
	util.Tab(buff, t+1, util.Fieldname("sApp")+fmt.Sprintf("%v\n", st.SApp))
	util.Tab(buff, t+1, util.Fieldname("sServer")+fmt.Sprintf("%v\n", st.SServer))
	util.Tab(buff, t+1, util.Fieldname("sDivision")+fmt.Sprintf("%v\n", st.SDivision))
	util.Tab(buff, t+1, util.Fieldname("sNodeName")+fmt.Sprintf("%v\n", st.SNodeName))
	util.Tab(buff, t+1, util.Fieldname("sCommand")+fmt.Sprintf("%v\n", st.SCommand))
	util.Tab(buff, t+1, util.Fieldname("mParam")+strconv.Itoa(len(st.MParam)))
	if len(st.MParam) == 0 {
		buff.WriteString(", {}\n")
	} else {
		buff.WriteString(", {\n")
	}
	for k, v := range st.MParam {
		util.Tab(buff, t+1+1, "(\n")
		util.Tab(buff, t+1+2, util.Fieldname("")+fmt.Sprintf("%v\n", k))
		util.Tab(buff, t+1+2, util.Fieldname("")+fmt.Sprintf("%v\n", v))
		util.Tab(buff, t+1+1, ")\n")
	}
	if len(st.MParam) != 0 {
		util.Tab(buff, t+1, "}\n")
	}
}
func (st *PatchTaskItemReq) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.STaskNo, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SApp, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SServer, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDivision, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SNodeName, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCommand, 6, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(7, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Map {
			return fmt.Errorf("tag:%d got wrong type %d", 7, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.MParam = make(map[string]string)
		for i := uint32(0); i < length; i++ {
			var k string
			err = up.ReadString(&k, 0, true)
			if err != nil {
				return err
			}
			var v string
			err = up.ReadString(&v, 0, true)
			if err != nil {
				return err
			}
			st.MParam[k] = v
		}
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *PatchTaskItemReq) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PatchTaskItemReq) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.STaskNo != "" {
		err = p.WriteString(0, st.STaskNo)
		if err != nil {
			return err
		}
	}
	if false || st.SApp != "" {
		err = p.WriteString(2, st.SApp)
		if err != nil {
			return err
		}
	}
	if false || st.SServer != "" {
		err = p.WriteString(3, st.SServer)
		if err != nil {
			return err
		}
	}
	if false || st.SDivision != "" {
		err = p.WriteString(4, st.SDivision)
		if err != nil {
			return err
		}
	}
	if false || st.SNodeName != "" {
		err = p.WriteString(5, st.SNodeName)
		if err != nil {
			return err
		}
	}
	if false || st.SCommand != "" {
		err = p.WriteString(6, st.SCommand)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.MParam))
	if false || length != 0 {
		err = p.WriteHeader(7, codec.SdpType_Map)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _k, _v := range st.MParam {
			if true || _k != "" {
				err = p.WriteString(0, _k)
				if err != nil {
					return err
				}
			}
			if true || _v != "" {
				err = p.WriteString(0, _v)
				if err != nil {
					return err
				}
			}
		}
	}

	_ = length
	return err
}
func (st *PatchTaskItemReq) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PatchTaskItemRsp struct {
	Req        PatchTaskItemReq `json:"req" form:"req"`
	IStatus    uint32           `json:"iStatus" form:"iStatus"`
	IStartTime uint32           `json:"iStartTime" form:"iStartTime"`
	IEndTime   uint32           `json:"iEndTime" form:"iEndTime"`
	SResult    string           `json:"sResult" form:"sResult"`
	IPercent   uint32           `json:"iPercent" form:"iPercent"`
}

func (st *PatchTaskItemRsp) resetDefault() {
	st.Req.resetDefault()
}
func (st *PatchTaskItemRsp) Copy() *PatchTaskItemRsp {
	ret := NewPatchTaskItemRsp()
	ret.Req = *(st.Req.Copy())
	ret.IStatus = st.IStatus
	ret.IStartTime = st.IStartTime
	ret.IEndTime = st.IEndTime
	ret.SResult = st.SResult
	ret.IPercent = st.IPercent
	return ret
}
func NewPatchTaskItemRsp() *PatchTaskItemRsp {
	ret := &PatchTaskItemRsp{}
	ret.resetDefault()
	return ret
}
func (st *PatchTaskItemRsp) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("req")+"{\n")
	st.Req.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("iStatus")+fmt.Sprintf("%v\n", st.IStatus))
	util.Tab(buff, t+1, util.Fieldname("iStartTime")+fmt.Sprintf("%v\n", st.IStartTime))
	util.Tab(buff, t+1, util.Fieldname("iEndTime")+fmt.Sprintf("%v\n", st.IEndTime))
	util.Tab(buff, t+1, util.Fieldname("sResult")+fmt.Sprintf("%v\n", st.SResult))
	util.Tab(buff, t+1, util.Fieldname("iPercent")+fmt.Sprintf("%v\n", st.IPercent))
}
func (st *PatchTaskItemRsp) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = st.Req.ReadStructFromTag(up, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IStatus, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IStartTime, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IEndTime, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SResult, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPercent, 5, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *PatchTaskItemRsp) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PatchTaskItemRsp) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	err = st.Req.WriteStructFromTag(p, 0, false)
	if err != nil {
		return err
	}
	if false || st.IStatus != 0 {
		err = p.WriteUint32(1, st.IStatus)
		if err != nil {
			return err
		}
	}
	if false || st.IStartTime != 0 {
		err = p.WriteUint32(2, st.IStartTime)
		if err != nil {
			return err
		}
	}
	if false || st.IEndTime != 0 {
		err = p.WriteUint32(3, st.IEndTime)
		if err != nil {
			return err
		}
	}
	if false || st.SResult != "" {
		err = p.WriteString(4, st.SResult)
		if err != nil {
			return err
		}
	}
	if false || st.IPercent != 0 {
		err = p.WriteUint32(5, st.IPercent)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *PatchTaskItemRsp) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PatchTaskReq struct {
	STaskNo string             `json:"sTaskNo" form:"sTaskNo"`
	VItem   []PatchTaskItemReq `json:"vItem" form:"vItem"`
}

func (st *PatchTaskReq) resetDefault() {
}
func (st *PatchTaskReq) Copy() *PatchTaskReq {
	ret := NewPatchTaskReq()
	ret.STaskNo = st.STaskNo
	ret.VItem = make([]PatchTaskItemReq, len(st.VItem))
	for i, v := range st.VItem {
		ret.VItem[i] = *(v.Copy())
	}
	return ret
}
func NewPatchTaskReq() *PatchTaskReq {
	ret := &PatchTaskReq{}
	ret.resetDefault()
	return ret
}
func (st *PatchTaskReq) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sTaskNo")+fmt.Sprintf("%v\n", st.STaskNo))
	util.Tab(buff, t+1, util.Fieldname("vItem")+strconv.Itoa(len(st.VItem)))
	if len(st.VItem) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VItem {
		util.Tab(buff, t+1+1, util.Fieldname("")+"{\n")
		v.Visit(buff, t+1+1+1)
		util.Tab(buff, t+1+1, "}\n")
	}
	if len(st.VItem) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
}
func (st *PatchTaskReq) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.STaskNo, 0, false)
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
		st.VItem = make([]PatchTaskItemReq, length, length)
		for i := uint32(0); i < length; i++ {
			err = st.VItem[i].ReadStructFromTag(up, 0, true)
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
func (st *PatchTaskReq) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PatchTaskReq) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.STaskNo != "" {
		err = p.WriteString(0, st.STaskNo)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.VItem))
	if false || length != 0 {
		err = p.WriteHeader(2, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VItem {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}

	_ = length
	return err
}
func (st *PatchTaskReq) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type PatchTaskRsp struct {
	STaskNo     string             `json:"sTaskNo" form:"sTaskNo"`
	IStatus     uint32             `json:"iStatus" form:"iStatus"`
	ICreateTime uint32             `json:"iCreateTime" form:"iCreateTime"`
	VItem       []PatchTaskItemRsp `json:"vItem" form:"vItem"`
}

func (st *PatchTaskRsp) resetDefault() {
}
func (st *PatchTaskRsp) Copy() *PatchTaskRsp {
	ret := NewPatchTaskRsp()
	ret.STaskNo = st.STaskNo
	ret.IStatus = st.IStatus
	ret.ICreateTime = st.ICreateTime
	ret.VItem = make([]PatchTaskItemRsp, len(st.VItem))
	for i, v := range st.VItem {
		ret.VItem[i] = *(v.Copy())
	}
	return ret
}
func NewPatchTaskRsp() *PatchTaskRsp {
	ret := &PatchTaskRsp{}
	ret.resetDefault()
	return ret
}
func (st *PatchTaskRsp) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sTaskNo")+fmt.Sprintf("%v\n", st.STaskNo))
	util.Tab(buff, t+1, util.Fieldname("iStatus")+fmt.Sprintf("%v\n", st.IStatus))
	util.Tab(buff, t+1, util.Fieldname("iCreateTime")+fmt.Sprintf("%v\n", st.ICreateTime))
	util.Tab(buff, t+1, util.Fieldname("vItem")+strconv.Itoa(len(st.VItem)))
	if len(st.VItem) == 0 {
		buff.WriteString(", []\n")
	} else {
		buff.WriteString(", [\n")
	}
	for _, v := range st.VItem {
		util.Tab(buff, t+1+1, util.Fieldname("")+"{\n")
		v.Visit(buff, t+1+1+1)
		util.Tab(buff, t+1+1, "}\n")
	}
	if len(st.VItem) != 0 {
		util.Tab(buff, t+1, "]\n")
	}
}
func (st *PatchTaskRsp) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.STaskNo, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IStatus, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICreateTime, 2, false)
	if err != nil {
		return err
	}

	has, ty, err = up.SkipToTag(3, false)
	if err != nil {
		return err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return fmt.Errorf("tag:%d got wrong type %d", 3, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return err
		}
		st.VItem = make([]PatchTaskItemRsp, length, length)
		for i := uint32(0); i < length; i++ {
			err = st.VItem[i].ReadStructFromTag(up, 0, true)
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
func (st *PatchTaskRsp) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *PatchTaskRsp) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.STaskNo != "" {
		err = p.WriteString(0, st.STaskNo)
		if err != nil {
			return err
		}
	}
	if false || st.IStatus != 0 {
		err = p.WriteUint32(1, st.IStatus)
		if err != nil {
			return err
		}
	}
	if false || st.ICreateTime != 0 {
		err = p.WriteUint32(2, st.ICreateTime)
		if err != nil {
			return err
		}
	}

	length = uint32(len(st.VItem))
	if false || length != 0 {
		err = p.WriteHeader(3, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range st.VItem {
			err = v.WriteStructFromTag(p, 0, true)
			if err != nil {
				return err
			}
		}
	}

	_ = length
	return err
}
func (st *PatchTaskRsp) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type Patch struct {
	proxy model.ServicePrxImpl
}

func (s *Patch) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *Patch) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *Patch) DownloadPatch(sFilePath string, iPos uint64, sBuff *[]byte, bEnd *bool) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sFilePath != "" {
		err = p.WriteString(1, sFilePath)
		if err != nil {
			return ret, err
		}
	}
	if true || iPos != 0 {
		err = p.WriteUint64(2, iPos)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("downloadPatch", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	var ssBuff string
	err = up.ReadString(&ssBuff, 3, true)
	if err != nil {
		return ret, err
	}
	(*sBuff) = []byte(ssBuff)
	err = up.ReadBool(&(*bEnd), 4, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *Patch) AddTask(stTaskReq PatchTaskReq) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stTaskReq.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("addTask", p.ToBytes(), &rsp)
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
func (s *Patch) GetTask(sTaskNo string, stTaskRsp *PatchTaskRsp) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sTaskNo != "" {
		err = p.WriteString(1, sTaskNo)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getTask", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stTaskRsp).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *Patch) CancelTask(sTaskNo string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sTaskNo != "" {
		err = p.WriteString(1, sTaskNo)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("cancelTask", p.ToBytes(), &rsp)
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
func (s *Patch) GetTemplate(sTemplateName string, sContent *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sTemplateName != "" {
		err = p.WriteString(1, sTemplateName)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getTemplate", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sContent), 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _PatchImpl interface {
	DownloadPatch(ctx context.Context, sFilePath string, iPos uint64, sBuff *[]byte, bEnd *bool) (int32, error)
	AddTask(ctx context.Context, stTaskReq PatchTaskReq) (int32, error)
	GetTask(ctx context.Context, sTaskNo string, stTaskRsp *PatchTaskRsp) (int32, error)
	CancelTask(ctx context.Context, sTaskNo string) (int32, error)
	GetTemplate(ctx context.Context, sTemplateName string, sContent *string) (int32, error)
}

func _PatchDownloadPatchImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PatchImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint64
	err = up.ReadUint64(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 []byte
	var p4 bool
	var ret int32
	ret, err = impl.DownloadPatch(ctx, p1, p2, &p3, &p4)
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
		stmp := string(p3)
		err = p.WriteString(3, stmp)
		if err != nil {
			return err
		}
	}
	if true || p4 != false {
		err = p.WriteBool(4, p4)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _PatchAddTaskImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PatchImpl)
	var p1 PatchTaskReq
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.AddTask(ctx, p1)
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
func _PatchGetTaskImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PatchImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 PatchTaskRsp
	var ret int32
	ret, err = impl.GetTask(ctx, p1, &p2)
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
func _PatchCancelTaskImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PatchImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.CancelTask(ctx, p1)
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
func _PatchGetTemplateImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_PatchImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 string
	var ret int32
	ret, err = impl.GetTemplate(ctx, p1, &p2)
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

func (s *Patch) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "downloadPatch":
		err = _PatchDownloadPatchImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "addTask":
		err = _PatchAddTaskImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getTask":
		err = _PatchGetTaskImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "cancelTask":
		err = _PatchCancelTaskImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getTemplate":
		err = _PatchGetTemplateImpl(ctx, serviceImpl, up, p)
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
