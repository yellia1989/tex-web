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
	"time"
)

type IAPReceiptType int32

const (
	IAPReceiptType_Apple    = 1
	IAPReceiptType_Google   = 2
	IAPReceiptType_Fy       = 3
	IAPReceiptType_HeePay   = 4
	IAPReceiptType_HeePayH5 = 5
)

type IAPReceiptStatus int32

const (
	IAPReceiptStatus_Pending         = 1
	IAPReceiptStatus_Verify_Fail     = 2
	IAPReceiptStatus_Delivering      = 3
	IAPReceiptStatus_Deliver_Success = 4
	IAPReceiptStatus_Deliver_Fail    = 5
)

type IAPStatus struct {
	IReceiptId      uint32 `json:"iReceiptId"`
	IReceiptStatus  uint32 `json:"iReceiptStatus"`
	IDeliverRoleId  uint64 `json:"iDeliverRoleId"`
	IDeliverZoneId  uint32 `json:"iDeliverZoneId"`
	IProxyRoleId    uint64 `json:"iProxyRoleId"`
	IProxyZoneId    uint32 `json:"iProxyZoneId"`
	IAddTime        uint32 `json:"iAddTime"`
	IVerifyTime     uint32 `json:"iVerifyTime"`
	IDeliverTime    uint32 `json:"iDeliverTime"`
	IRetryNum       uint32 `json:"iRetryNum"`
	INextTryTime    uint32 `json:"iNextTryTime"`
	SFailReason     string `json:"sFailReason"`
	SDeliverItem    string `json:"sDeliverItem"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
}

func (st *IAPStatus) resetDefault() {
}
func (st *IAPStatus) Copy() *IAPStatus {
	ret := NewIAPStatus()
	ret.IReceiptId = st.IReceiptId
	ret.IReceiptStatus = st.IReceiptStatus
	ret.IDeliverRoleId = st.IDeliverRoleId
	ret.IDeliverZoneId = st.IDeliverZoneId
	ret.IProxyRoleId = st.IProxyRoleId
	ret.IProxyZoneId = st.IProxyZoneId
	ret.IAddTime = st.IAddTime
	ret.IVerifyTime = st.IVerifyTime
	ret.IDeliverTime = st.IDeliverTime
	ret.IRetryNum = st.IRetryNum
	ret.INextTryTime = st.INextTryTime
	ret.SFailReason = st.SFailReason
	ret.SDeliverItem = st.SDeliverItem
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewIAPStatus() *IAPStatus {
	ret := &IAPStatus{}
	ret.resetDefault()
	return ret
}
func (st *IAPStatus) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iReceiptId")+fmt.Sprintf("%v\n", st.IReceiptId))
	util.Tab(buff, t+1, util.Fieldname("iReceiptStatus")+fmt.Sprintf("%v\n", st.IReceiptStatus))
	util.Tab(buff, t+1, util.Fieldname("iDeliverRoleId")+fmt.Sprintf("%v\n", st.IDeliverRoleId))
	util.Tab(buff, t+1, util.Fieldname("iDeliverZoneId")+fmt.Sprintf("%v\n", st.IDeliverZoneId))
	util.Tab(buff, t+1, util.Fieldname("iProxyRoleId")+fmt.Sprintf("%v\n", st.IProxyRoleId))
	util.Tab(buff, t+1, util.Fieldname("iProxyZoneId")+fmt.Sprintf("%v\n", st.IProxyZoneId))
	util.Tab(buff, t+1, util.Fieldname("iAddTime")+fmt.Sprintf("%v\n", st.IAddTime))
	util.Tab(buff, t+1, util.Fieldname("iVerifyTime")+fmt.Sprintf("%v\n", st.IVerifyTime))
	util.Tab(buff, t+1, util.Fieldname("iDeliverTime")+fmt.Sprintf("%v\n", st.IDeliverTime))
	util.Tab(buff, t+1, util.Fieldname("iRetryNum")+fmt.Sprintf("%v\n", st.IRetryNum))
	util.Tab(buff, t+1, util.Fieldname("iNextTryTime")+fmt.Sprintf("%v\n", st.INextTryTime))
	util.Tab(buff, t+1, util.Fieldname("sFailReason")+fmt.Sprintf("%v\n", st.SFailReason))
	util.Tab(buff, t+1, util.Fieldname("sDeliverItem")+fmt.Sprintf("%v\n", st.SDeliverItem))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *IAPStatus) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IReceiptId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IReceiptStatus, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint64(&st.IDeliverRoleId, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDeliverZoneId, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint64(&st.IProxyRoleId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IProxyZoneId, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IAddTime, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IVerifyTime, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IDeliverTime, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IRetryNum, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.INextTryTime, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SFailReason, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDeliverItem, 12, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 13, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 14, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *IAPStatus) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *IAPStatus) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IReceiptId != 0 {
		err = p.WriteUint32(0, st.IReceiptId)
		if err != nil {
			return err
		}
	}
	if false || st.IReceiptStatus != 0 {
		err = p.WriteUint32(1, st.IReceiptStatus)
		if err != nil {
			return err
		}
	}
	if false || st.IDeliverRoleId != 0 {
		err = p.WriteUint64(2, st.IDeliverRoleId)
		if err != nil {
			return err
		}
	}
	if false || st.IDeliverZoneId != 0 {
		err = p.WriteUint32(3, st.IDeliverZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.IProxyRoleId != 0 {
		err = p.WriteUint64(4, st.IProxyRoleId)
		if err != nil {
			return err
		}
	}
	if false || st.IProxyZoneId != 0 {
		err = p.WriteUint32(5, st.IProxyZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.IAddTime != 0 {
		err = p.WriteUint32(6, st.IAddTime)
		if err != nil {
			return err
		}
	}
	if false || st.IVerifyTime != 0 {
		err = p.WriteUint32(7, st.IVerifyTime)
		if err != nil {
			return err
		}
	}
	if false || st.IDeliverTime != 0 {
		err = p.WriteUint32(8, st.IDeliverTime)
		if err != nil {
			return err
		}
	}
	if false || st.IRetryNum != 0 {
		err = p.WriteUint32(9, st.IRetryNum)
		if err != nil {
			return err
		}
	}
	if false || st.INextTryTime != 0 {
		err = p.WriteUint32(10, st.INextTryTime)
		if err != nil {
			return err
		}
	}
	if false || st.SFailReason != "" {
		err = p.WriteString(11, st.SFailReason)
		if err != nil {
			return err
		}
	}
	if false || st.SDeliverItem != "" {
		err = p.WriteString(12, st.SDeliverItem)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(13, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(14, st.STraceFlowId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *IAPStatus) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type IAPTmpOrder struct {
	SFlowId      string `json:"sFlowId"`
	IProductId   uint32 `json:"iProductId"`
	IReceiptType uint32 `json:"iReceiptType"`
	IRoleId      uint64 `json:"iRoleId"`
	IZoneId      uint32 `json:"iZoneId"`
	ICreateTime  uint32 `json:"iCreateTime"`
	SPayload     string `json:"sPayload"`
}

func (st *IAPTmpOrder) resetDefault() {
}
func (st *IAPTmpOrder) Copy() *IAPTmpOrder {
	ret := NewIAPTmpOrder()
	ret.SFlowId = st.SFlowId
	ret.IProductId = st.IProductId
	ret.IReceiptType = st.IReceiptType
	ret.IRoleId = st.IRoleId
	ret.IZoneId = st.IZoneId
	ret.ICreateTime = st.ICreateTime
	ret.SPayload = st.SPayload
	return ret
}
func NewIAPTmpOrder() *IAPTmpOrder {
	ret := &IAPTmpOrder{}
	ret.resetDefault()
	return ret
}
func (st *IAPTmpOrder) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sFlowId")+fmt.Sprintf("%v\n", st.SFlowId))
	util.Tab(buff, t+1, util.Fieldname("iProductId")+fmt.Sprintf("%v\n", st.IProductId))
	util.Tab(buff, t+1, util.Fieldname("iReceiptType")+fmt.Sprintf("%v\n", st.IReceiptType))
	util.Tab(buff, t+1, util.Fieldname("iRoleId")+fmt.Sprintf("%v\n", st.IRoleId))
	util.Tab(buff, t+1, util.Fieldname("iZoneId")+fmt.Sprintf("%v\n", st.IZoneId))
	util.Tab(buff, t+1, util.Fieldname("iCreateTime")+fmt.Sprintf("%v\n", st.ICreateTime))
	util.Tab(buff, t+1, util.Fieldname("sPayload")+fmt.Sprintf("%v\n", st.SPayload))
}
func (st *IAPTmpOrder) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SFlowId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IProductId, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IReceiptType, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint64(&st.IRoleId, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IZoneId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICreateTime, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPayload, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *IAPTmpOrder) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *IAPTmpOrder) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SFlowId != "" {
		err = p.WriteString(0, st.SFlowId)
		if err != nil {
			return err
		}
	}
	if false || st.IProductId != 0 {
		err = p.WriteUint32(1, st.IProductId)
		if err != nil {
			return err
		}
	}
	if false || st.IReceiptType != 0 {
		err = p.WriteUint32(2, st.IReceiptType)
		if err != nil {
			return err
		}
	}
	if false || st.IRoleId != 0 {
		err = p.WriteUint64(3, st.IRoleId)
		if err != nil {
			return err
		}
	}
	if false || st.IZoneId != 0 {
		err = p.WriteUint32(4, st.IZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.ICreateTime != 0 {
		err = p.WriteUint32(5, st.ICreateTime)
		if err != nil {
			return err
		}
	}
	if false || st.SPayload != "" {
		err = p.WriteString(6, st.SPayload)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *IAPTmpOrder) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type ApplePurchase struct {
	SReceiptData    string `json:"sReceiptData"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
}

func (st *ApplePurchase) resetDefault() {
}
func (st *ApplePurchase) Copy() *ApplePurchase {
	ret := NewApplePurchase()
	ret.SReceiptData = st.SReceiptData
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewApplePurchase() *ApplePurchase {
	ret := &ApplePurchase{}
	ret.resetDefault()
	return ret
}
func (st *ApplePurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sReceiptData")+fmt.Sprintf("%v\n", st.SReceiptData))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *ApplePurchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SReceiptData, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 11, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ApplePurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ApplePurchase) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SReceiptData != "" {
		err = p.WriteString(0, st.SReceiptData)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(10, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(11, st.STraceFlowId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ApplePurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type AppleReceipt struct {
	IQuantity          uint32 `json:"iQuantity"`
	SProductId         string `json:"sProductId"`
	STransactionId     string `json:"sTransactionId"`
	IPurchaseDate      uint32 `json:"iPurchaseDate"`
	SBId               string `json:"sBId"`
	SBVrs              string `json:"sBVrs"`
	SOriTransactionId  string `json:"sOriTransactionId"`
	IOriPurchaseDate   uint32 `json:"iOriPurchaseDate"`
	SAppItemId         string `json:"sAppItemId"`
	SVersionExternalId string `json:"sVersionExternalId"`
}

func (st *AppleReceipt) resetDefault() {
}
func (st *AppleReceipt) Copy() *AppleReceipt {
	ret := NewAppleReceipt()
	ret.IQuantity = st.IQuantity
	ret.SProductId = st.SProductId
	ret.STransactionId = st.STransactionId
	ret.IPurchaseDate = st.IPurchaseDate
	ret.SBId = st.SBId
	ret.SBVrs = st.SBVrs
	ret.SOriTransactionId = st.SOriTransactionId
	ret.IOriPurchaseDate = st.IOriPurchaseDate
	ret.SAppItemId = st.SAppItemId
	ret.SVersionExternalId = st.SVersionExternalId
	return ret
}
func NewAppleReceipt() *AppleReceipt {
	ret := &AppleReceipt{}
	ret.resetDefault()
	return ret
}
func (st *AppleReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iQuantity")+fmt.Sprintf("%v\n", st.IQuantity))
	util.Tab(buff, t+1, util.Fieldname("sProductId")+fmt.Sprintf("%v\n", st.SProductId))
	util.Tab(buff, t+1, util.Fieldname("sTransactionId")+fmt.Sprintf("%v\n", st.STransactionId))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseDate")+fmt.Sprintf("%v\n", st.IPurchaseDate))
	util.Tab(buff, t+1, util.Fieldname("sBId")+fmt.Sprintf("%v\n", st.SBId))
	util.Tab(buff, t+1, util.Fieldname("sBVrs")+fmt.Sprintf("%v\n", st.SBVrs))
	util.Tab(buff, t+1, util.Fieldname("sOriTransactionId")+fmt.Sprintf("%v\n", st.SOriTransactionId))
	util.Tab(buff, t+1, util.Fieldname("iOriPurchaseDate")+fmt.Sprintf("%v\n", st.IOriPurchaseDate))
	util.Tab(buff, t+1, util.Fieldname("sAppItemId")+fmt.Sprintf("%v\n", st.SAppItemId))
	util.Tab(buff, t+1, util.Fieldname("sVersionExternalId")+fmt.Sprintf("%v\n", st.SVersionExternalId))
}
func (st *AppleReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IQuantity, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProductId, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STransactionId, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseDate, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBVrs, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOriTransactionId, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IOriPurchaseDate, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppItemId, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SVersionExternalId, 9, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *AppleReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *AppleReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IQuantity != 0 {
		err = p.WriteUint32(0, st.IQuantity)
		if err != nil {
			return err
		}
	}
	if false || st.SProductId != "" {
		err = p.WriteString(1, st.SProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STransactionId != "" {
		err = p.WriteString(2, st.STransactionId)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseDate != 0 {
		err = p.WriteUint32(3, st.IPurchaseDate)
		if err != nil {
			return err
		}
	}
	if false || st.SBId != "" {
		err = p.WriteString(4, st.SBId)
		if err != nil {
			return err
		}
	}
	if false || st.SBVrs != "" {
		err = p.WriteString(5, st.SBVrs)
		if err != nil {
			return err
		}
	}
	if false || st.SOriTransactionId != "" {
		err = p.WriteString(6, st.SOriTransactionId)
		if err != nil {
			return err
		}
	}
	if false || st.IOriPurchaseDate != 0 {
		err = p.WriteUint32(7, st.IOriPurchaseDate)
		if err != nil {
			return err
		}
	}
	if false || st.SAppItemId != "" {
		err = p.WriteString(8, st.SAppItemId)
		if err != nil {
			return err
		}
	}
	if false || st.SVersionExternalId != "" {
		err = p.WriteString(9, st.SVersionExternalId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *AppleReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type GooglePurchase struct {
	IResponseCode   int32  `json:"iResponseCode"`
	SPurchaseData   string `json:"sPurchaseData"`
	SSignature      string `json:"sSignature"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
}

func (st *GooglePurchase) resetDefault() {
}
func (st *GooglePurchase) Copy() *GooglePurchase {
	ret := NewGooglePurchase()
	ret.IResponseCode = st.IResponseCode
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewGooglePurchase() *GooglePurchase {
	ret := &GooglePurchase{}
	ret.resetDefault()
	return ret
}
func (st *GooglePurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iResponseCode")+fmt.Sprintf("%v\n", st.IResponseCode))
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *GooglePurchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadInt32(&st.IResponseCode, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPurchaseData, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SSignature, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 11, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *GooglePurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *GooglePurchase) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IResponseCode != 0 {
		err = p.WriteInt32(0, st.IResponseCode)
		if err != nil {
			return err
		}
	}
	if false || st.SPurchaseData != "" {
		err = p.WriteString(1, st.SPurchaseData)
		if err != nil {
			return err
		}
	}
	if false || st.SSignature != "" {
		err = p.WriteString(2, st.SSignature)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(10, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(11, st.STraceFlowId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *GooglePurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type GoogleReceipt struct {
	SOrderId          string `json:"sOrderId"`
	SProductId        string `json:"sProductId"`
	SPackageName      string `json:"sPackageName"`
	IPurchaseTime     uint32 `json:"iPurchaseTime"`
	IPurchaseState    uint32 `json:"iPurchaseState"`
	SDeveloperPayload string `json:"sDeveloperPayload"`
	SPurchaseToken    string `json:"sPurchaseToken"`
	BAutoRenewing     bool   `json:"bAutoRenewing"`
}

func (st *GoogleReceipt) resetDefault() {
}
func (st *GoogleReceipt) Copy() *GoogleReceipt {
	ret := NewGoogleReceipt()
	ret.SOrderId = st.SOrderId
	ret.SProductId = st.SProductId
	ret.SPackageName = st.SPackageName
	ret.IPurchaseTime = st.IPurchaseTime
	ret.IPurchaseState = st.IPurchaseState
	ret.SDeveloperPayload = st.SDeveloperPayload
	ret.SPurchaseToken = st.SPurchaseToken
	ret.BAutoRenewing = st.BAutoRenewing
	return ret
}
func NewGoogleReceipt() *GoogleReceipt {
	ret := &GoogleReceipt{}
	ret.resetDefault()
	return ret
}
func (st *GoogleReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sOrderId")+fmt.Sprintf("%v\n", st.SOrderId))
	util.Tab(buff, t+1, util.Fieldname("sProductId")+fmt.Sprintf("%v\n", st.SProductId))
	util.Tab(buff, t+1, util.Fieldname("sPackageName")+fmt.Sprintf("%v\n", st.SPackageName))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseTime")+fmt.Sprintf("%v\n", st.IPurchaseTime))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseState")+fmt.Sprintf("%v\n", st.IPurchaseState))
	util.Tab(buff, t+1, util.Fieldname("sDeveloperPayload")+fmt.Sprintf("%v\n", st.SDeveloperPayload))
	util.Tab(buff, t+1, util.Fieldname("sPurchaseToken")+fmt.Sprintf("%v\n", st.SPurchaseToken))
	util.Tab(buff, t+1, util.Fieldname("bAutoRenewing")+fmt.Sprintf("%v\n", st.BAutoRenewing))
}
func (st *GoogleReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SOrderId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProductId, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPackageName, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseTime, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseState, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDeveloperPayload, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPurchaseToken, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadBool(&st.BAutoRenewing, 7, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *GoogleReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *GoogleReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SOrderId != "" {
		err = p.WriteString(0, st.SOrderId)
		if err != nil {
			return err
		}
	}
	if false || st.SProductId != "" {
		err = p.WriteString(1, st.SProductId)
		if err != nil {
			return err
		}
	}
	if false || st.SPackageName != "" {
		err = p.WriteString(2, st.SPackageName)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseTime != 0 {
		err = p.WriteUint32(3, st.IPurchaseTime)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseState != 0 {
		err = p.WriteUint32(4, st.IPurchaseState)
		if err != nil {
			return err
		}
	}
	if false || st.SDeveloperPayload != "" {
		err = p.WriteString(5, st.SDeveloperPayload)
		if err != nil {
			return err
		}
	}
	if false || st.SPurchaseToken != "" {
		err = p.WriteString(6, st.SPurchaseToken)
		if err != nil {
			return err
		}
	}
	if false || st.BAutoRenewing != false {
		err = p.WriteBool(7, st.BAutoRenewing)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *GoogleReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type FyPurchase struct {
	SPurchaseData   string `json:"sPurchaseData"`
	SSignature      string `json:"sSignature"`
	IConnId         uint32 `json:"iConnId"`
	SChannel        string `json:"sChannel"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
}

func (st *FyPurchase) resetDefault() {
}
func (st *FyPurchase) Copy() *FyPurchase {
	ret := NewFyPurchase()
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.IConnId = st.IConnId
	ret.SChannel = st.SChannel
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewFyPurchase() *FyPurchase {
	ret := &FyPurchase{}
	ret.resetDefault()
	return ret
}
func (st *FyPurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("iConnId")+fmt.Sprintf("%v\n", st.IConnId))
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *FyPurchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SPurchaseData, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SSignature, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IConnId, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SChannel, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 11, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *FyPurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *FyPurchase) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SPurchaseData != "" {
		err = p.WriteString(1, st.SPurchaseData)
		if err != nil {
			return err
		}
	}
	if false || st.SSignature != "" {
		err = p.WriteString(2, st.SSignature)
		if err != nil {
			return err
		}
	}
	if false || st.IConnId != 0 {
		err = p.WriteUint32(3, st.IConnId)
		if err != nil {
			return err
		}
	}
	if false || st.SChannel != "" {
		err = p.WriteString(4, st.SChannel)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(10, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(11, st.STraceFlowId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *FyPurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type FyReceipt struct {
	SOrderId        string `json:"sOrderId"`
	SUuid           string `json:"sUuid"`
	SAppCallbackExt string `json:"sAppCallbackExt"`
	IPayAmount      uint32 `json:"iPayAmount"`
	ISandBox        uint32 `json:"iSandBox"`
	IPayTime        uint32 `json:"iPayTime"`
	ITime           uint32 `json:"iTime"`
}

func (st *FyReceipt) resetDefault() {
}
func (st *FyReceipt) Copy() *FyReceipt {
	ret := NewFyReceipt()
	ret.SOrderId = st.SOrderId
	ret.SUuid = st.SUuid
	ret.SAppCallbackExt = st.SAppCallbackExt
	ret.IPayAmount = st.IPayAmount
	ret.ISandBox = st.ISandBox
	ret.IPayTime = st.IPayTime
	ret.ITime = st.ITime
	return ret
}
func NewFyReceipt() *FyReceipt {
	ret := &FyReceipt{}
	ret.resetDefault()
	return ret
}
func (st *FyReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sOrderId")+fmt.Sprintf("%v\n", st.SOrderId))
	util.Tab(buff, t+1, util.Fieldname("sUuid")+fmt.Sprintf("%v\n", st.SUuid))
	util.Tab(buff, t+1, util.Fieldname("sAppCallbackExt")+fmt.Sprintf("%v\n", st.SAppCallbackExt))
	util.Tab(buff, t+1, util.Fieldname("iPayAmount")+fmt.Sprintf("%v\n", st.IPayAmount))
	util.Tab(buff, t+1, util.Fieldname("iSandBox")+fmt.Sprintf("%v\n", st.ISandBox))
	util.Tab(buff, t+1, util.Fieldname("iPayTime")+fmt.Sprintf("%v\n", st.IPayTime))
	util.Tab(buff, t+1, util.Fieldname("iTime")+fmt.Sprintf("%v\n", st.ITime))
}
func (st *FyReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SOrderId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SUuid, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppCallbackExt, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPayAmount, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ISandBox, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPayTime, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITime, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *FyReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *FyReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SOrderId != "" {
		err = p.WriteString(0, st.SOrderId)
		if err != nil {
			return err
		}
	}
	if false || st.SUuid != "" {
		err = p.WriteString(1, st.SUuid)
		if err != nil {
			return err
		}
	}
	if false || st.SAppCallbackExt != "" {
		err = p.WriteString(2, st.SAppCallbackExt)
		if err != nil {
			return err
		}
	}
	if false || st.IPayAmount != 0 {
		err = p.WriteUint32(3, st.IPayAmount)
		if err != nil {
			return err
		}
	}
	if false || st.ISandBox != 0 {
		err = p.WriteUint32(4, st.ISandBox)
		if err != nil {
			return err
		}
	}
	if false || st.IPayTime != 0 {
		err = p.WriteUint32(5, st.IPayTime)
		if err != nil {
			return err
		}
	}
	if false || st.ITime != 0 {
		err = p.WriteUint32(6, st.ITime)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *FyReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type HeePayPurchase struct {
	SPurchaseData   string `json:"sPurchaseData"`
	SSignature      string `json:"sSignature"`
	IConnId         uint32 `json:"iConnId"`
	SChannel        string `json:"sChannel"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
	ICreateTime     uint32 `json:"iCreateTime"`
}

func (st *HeePayPurchase) resetDefault() {
}
func (st *HeePayPurchase) Copy() *HeePayPurchase {
	ret := NewHeePayPurchase()
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.IConnId = st.IConnId
	ret.SChannel = st.SChannel
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	ret.ICreateTime = st.ICreateTime
	return ret
}
func NewHeePayPurchase() *HeePayPurchase {
	ret := &HeePayPurchase{}
	ret.resetDefault()
	return ret
}
func (st *HeePayPurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("iConnId")+fmt.Sprintf("%v\n", st.IConnId))
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
	util.Tab(buff, t+1, util.Fieldname("iCreateTime")+fmt.Sprintf("%v\n", st.ICreateTime))
}
func (st *HeePayPurchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SPurchaseData, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SSignature, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IConnId, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SChannel, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICreateTime, 12, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *HeePayPurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *HeePayPurchase) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SPurchaseData != "" {
		err = p.WriteString(0, st.SPurchaseData)
		if err != nil {
			return err
		}
	}
	if false || st.SSignature != "" {
		err = p.WriteString(1, st.SSignature)
		if err != nil {
			return err
		}
	}
	if false || st.IConnId != 0 {
		err = p.WriteUint32(2, st.IConnId)
		if err != nil {
			return err
		}
	}
	if false || st.SChannel != "" {
		err = p.WriteString(3, st.SChannel)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(10, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(11, st.STraceFlowId)
		if err != nil {
			return err
		}
	}
	if false || st.ICreateTime != 0 {
		err = p.WriteUint32(12, st.ICreateTime)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *HeePayPurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type HeePayReceipt struct {
	IResult      int32   `json:"iResult"`
	SPayMessage  string  `json:"sPayMessage"`
	SAgentId     string  `json:"sAgentId"`
	SJnetBillNo  string  `json:"sJnetBillNo"`
	SAgentBillId string  `json:"sAgentBillId"`
	IPayType     uint32  `json:"iPayType"`
	FPayAmt      float32 `json:"fPayAmt"`
	SRemark      string  `json:"sRemark"`
}

func (st *HeePayReceipt) resetDefault() {
}
func (st *HeePayReceipt) Copy() *HeePayReceipt {
	ret := NewHeePayReceipt()
	ret.IResult = st.IResult
	ret.SPayMessage = st.SPayMessage
	ret.SAgentId = st.SAgentId
	ret.SJnetBillNo = st.SJnetBillNo
	ret.SAgentBillId = st.SAgentBillId
	ret.IPayType = st.IPayType
	ret.FPayAmt = st.FPayAmt
	ret.SRemark = st.SRemark
	return ret
}
func NewHeePayReceipt() *HeePayReceipt {
	ret := &HeePayReceipt{}
	ret.resetDefault()
	return ret
}
func (st *HeePayReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iResult")+fmt.Sprintf("%v\n", st.IResult))
	util.Tab(buff, t+1, util.Fieldname("sPayMessage")+fmt.Sprintf("%v\n", st.SPayMessage))
	util.Tab(buff, t+1, util.Fieldname("sAgentId")+fmt.Sprintf("%v\n", st.SAgentId))
	util.Tab(buff, t+1, util.Fieldname("sJnetBillNo")+fmt.Sprintf("%v\n", st.SJnetBillNo))
	util.Tab(buff, t+1, util.Fieldname("sAgentBillId")+fmt.Sprintf("%v\n", st.SAgentBillId))
	util.Tab(buff, t+1, util.Fieldname("iPayType")+fmt.Sprintf("%v\n", st.IPayType))
	util.Tab(buff, t+1, util.Fieldname("fPayAmt")+fmt.Sprintf("%v\n", st.FPayAmt))
	util.Tab(buff, t+1, util.Fieldname("sRemark")+fmt.Sprintf("%v\n", st.SRemark))
}
func (st *HeePayReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadInt32(&st.IResult, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPayMessage, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAgentId, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SJnetBillNo, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAgentBillId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPayType, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadFloat32(&st.FPayAmt, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRemark, 7, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *HeePayReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *HeePayReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IResult != 0 {
		err = p.WriteInt32(0, st.IResult)
		if err != nil {
			return err
		}
	}
	if false || st.SPayMessage != "" {
		err = p.WriteString(1, st.SPayMessage)
		if err != nil {
			return err
		}
	}
	if false || st.SAgentId != "" {
		err = p.WriteString(2, st.SAgentId)
		if err != nil {
			return err
		}
	}
	if false || st.SJnetBillNo != "" {
		err = p.WriteString(3, st.SJnetBillNo)
		if err != nil {
			return err
		}
	}
	if false || st.SAgentBillId != "" {
		err = p.WriteString(4, st.SAgentBillId)
		if err != nil {
			return err
		}
	}
	if false || st.IPayType != 0 {
		err = p.WriteUint32(5, st.IPayType)
		if err != nil {
			return err
		}
	}
	if false || st.FPayAmt != 0 {
		err = p.WriteFloat32(6, st.FPayAmt)
		if err != nil {
			return err
		}
	}
	if false || st.SRemark != "" {
		err = p.WriteString(7, st.SRemark)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *HeePayReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type HeePayH5Purchase struct {
	SPurchaseData   string `json:"sPurchaseData"`
	SSignature      string `json:"sSignature"`
	IConnId         uint32 `json:"iConnId"`
	SChannel        string `json:"sChannel"`
	STraceProductId string `json:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId"`
}

func (st *HeePayH5Purchase) resetDefault() {
}
func (st *HeePayH5Purchase) Copy() *HeePayH5Purchase {
	ret := NewHeePayH5Purchase()
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.IConnId = st.IConnId
	ret.SChannel = st.SChannel
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewHeePayH5Purchase() *HeePayH5Purchase {
	ret := &HeePayH5Purchase{}
	ret.resetDefault()
	return ret
}
func (st *HeePayH5Purchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("iConnId")+fmt.Sprintf("%v\n", st.IConnId))
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *HeePayH5Purchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SPurchaseData, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SSignature, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IConnId, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SChannel, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceProductId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STraceFlowId, 11, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *HeePayH5Purchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *HeePayH5Purchase) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SPurchaseData != "" {
		err = p.WriteString(0, st.SPurchaseData)
		if err != nil {
			return err
		}
	}
	if false || st.SSignature != "" {
		err = p.WriteString(1, st.SSignature)
		if err != nil {
			return err
		}
	}
	if false || st.IConnId != 0 {
		err = p.WriteUint32(2, st.IConnId)
		if err != nil {
			return err
		}
	}
	if false || st.SChannel != "" {
		err = p.WriteString(3, st.SChannel)
		if err != nil {
			return err
		}
	}
	if false || st.STraceProductId != "" {
		err = p.WriteString(10, st.STraceProductId)
		if err != nil {
			return err
		}
	}
	if false || st.STraceFlowId != "" {
		err = p.WriteString(11, st.STraceFlowId)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *HeePayH5Purchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type HeePayH5Receipt struct {
	SMethod        string `json:"sMethod"`
	SVersion       string `json:"sVersion"`
	SCharset       string `json:"sCharset"`
	SSignType      string `json:"sSignType"`
	SReturnCode    string `json:"sReturnCode"`
	SReturnMsg     string `json:"sReturnMsg"`
	SAppId         string `json:"sAppId"`
	SMchId         string `json:"sMchId"`
	SNonceStr      string `json:"sNonceStr"`
	SResultCode    string `json:"sResultCode"`
	SErrCode       string `json:"sErrCode"`
	SErrCodeDesc   string `json:"sErrCodeDesc"`
	SOpenId        string `json:"sOpenId"`
	SFeeType       string `json:"sFeeType"`
	ITotalFee      uint32 `json:"iTotalFee"`
	ICouponFee     uint32 `json:"iCouponFee"`
	STransactionId string `json:"sTransactionId"`
	SOutTradeNo    string `json:"sOutTradeNo"`
	STimeEnd       string `json:"sTimeEnd"`
	SBuyerLogonId  string `json:"sBuyerLogonId"`
	SFundBillList  string `json:"sFundBillList"`
}

func (st *HeePayH5Receipt) resetDefault() {
}
func (st *HeePayH5Receipt) Copy() *HeePayH5Receipt {
	ret := NewHeePayH5Receipt()
	ret.SMethod = st.SMethod
	ret.SVersion = st.SVersion
	ret.SCharset = st.SCharset
	ret.SSignType = st.SSignType
	ret.SReturnCode = st.SReturnCode
	ret.SReturnMsg = st.SReturnMsg
	ret.SAppId = st.SAppId
	ret.SMchId = st.SMchId
	ret.SNonceStr = st.SNonceStr
	ret.SResultCode = st.SResultCode
	ret.SErrCode = st.SErrCode
	ret.SErrCodeDesc = st.SErrCodeDesc
	ret.SOpenId = st.SOpenId
	ret.SFeeType = st.SFeeType
	ret.ITotalFee = st.ITotalFee
	ret.ICouponFee = st.ICouponFee
	ret.STransactionId = st.STransactionId
	ret.SOutTradeNo = st.SOutTradeNo
	ret.STimeEnd = st.STimeEnd
	ret.SBuyerLogonId = st.SBuyerLogonId
	ret.SFundBillList = st.SFundBillList
	return ret
}
func NewHeePayH5Receipt() *HeePayH5Receipt {
	ret := &HeePayH5Receipt{}
	ret.resetDefault()
	return ret
}
func (st *HeePayH5Receipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sMethod")+fmt.Sprintf("%v\n", st.SMethod))
	util.Tab(buff, t+1, util.Fieldname("sVersion")+fmt.Sprintf("%v\n", st.SVersion))
	util.Tab(buff, t+1, util.Fieldname("sCharset")+fmt.Sprintf("%v\n", st.SCharset))
	util.Tab(buff, t+1, util.Fieldname("sSignType")+fmt.Sprintf("%v\n", st.SSignType))
	util.Tab(buff, t+1, util.Fieldname("sReturnCode")+fmt.Sprintf("%v\n", st.SReturnCode))
	util.Tab(buff, t+1, util.Fieldname("sReturnMsg")+fmt.Sprintf("%v\n", st.SReturnMsg))
	util.Tab(buff, t+1, util.Fieldname("sAppId")+fmt.Sprintf("%v\n", st.SAppId))
	util.Tab(buff, t+1, util.Fieldname("sMchId")+fmt.Sprintf("%v\n", st.SMchId))
	util.Tab(buff, t+1, util.Fieldname("sNonceStr")+fmt.Sprintf("%v\n", st.SNonceStr))
	util.Tab(buff, t+1, util.Fieldname("sResultCode")+fmt.Sprintf("%v\n", st.SResultCode))
	util.Tab(buff, t+1, util.Fieldname("sErrCode")+fmt.Sprintf("%v\n", st.SErrCode))
	util.Tab(buff, t+1, util.Fieldname("sErrCodeDesc")+fmt.Sprintf("%v\n", st.SErrCodeDesc))
	util.Tab(buff, t+1, util.Fieldname("sOpenId")+fmt.Sprintf("%v\n", st.SOpenId))
	util.Tab(buff, t+1, util.Fieldname("sFeeType")+fmt.Sprintf("%v\n", st.SFeeType))
	util.Tab(buff, t+1, util.Fieldname("iTotalFee")+fmt.Sprintf("%v\n", st.ITotalFee))
	util.Tab(buff, t+1, util.Fieldname("iCouponFee")+fmt.Sprintf("%v\n", st.ICouponFee))
	util.Tab(buff, t+1, util.Fieldname("sTransactionId")+fmt.Sprintf("%v\n", st.STransactionId))
	util.Tab(buff, t+1, util.Fieldname("sOutTradeNo")+fmt.Sprintf("%v\n", st.SOutTradeNo))
	util.Tab(buff, t+1, util.Fieldname("sTimeEnd")+fmt.Sprintf("%v\n", st.STimeEnd))
	util.Tab(buff, t+1, util.Fieldname("sBuyerLogonId")+fmt.Sprintf("%v\n", st.SBuyerLogonId))
	util.Tab(buff, t+1, util.Fieldname("sFundBillList")+fmt.Sprintf("%v\n", st.SFundBillList))
}
func (st *HeePayH5Receipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SMethod, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SVersion, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCharset, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SSignType, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SReturnCode, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SReturnMsg, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppId, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SMchId, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SNonceStr, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SResultCode, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SErrCode, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SErrCodeDesc, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOpenId, 12, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SFeeType, 13, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITotalFee, 14, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ICouponFee, 15, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STransactionId, 16, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOutTradeNo, 17, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STimeEnd, 18, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBuyerLogonId, 19, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SFundBillList, 20, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *HeePayH5Receipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *HeePayH5Receipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SMethod != "" {
		err = p.WriteString(0, st.SMethod)
		if err != nil {
			return err
		}
	}
	if false || st.SVersion != "" {
		err = p.WriteString(1, st.SVersion)
		if err != nil {
			return err
		}
	}
	if false || st.SCharset != "" {
		err = p.WriteString(2, st.SCharset)
		if err != nil {
			return err
		}
	}
	if false || st.SSignType != "" {
		err = p.WriteString(3, st.SSignType)
		if err != nil {
			return err
		}
	}
	if false || st.SReturnCode != "" {
		err = p.WriteString(4, st.SReturnCode)
		if err != nil {
			return err
		}
	}
	if false || st.SReturnMsg != "" {
		err = p.WriteString(5, st.SReturnMsg)
		if err != nil {
			return err
		}
	}
	if false || st.SAppId != "" {
		err = p.WriteString(6, st.SAppId)
		if err != nil {
			return err
		}
	}
	if false || st.SMchId != "" {
		err = p.WriteString(7, st.SMchId)
		if err != nil {
			return err
		}
	}
	if false || st.SNonceStr != "" {
		err = p.WriteString(8, st.SNonceStr)
		if err != nil {
			return err
		}
	}
	if false || st.SResultCode != "" {
		err = p.WriteString(9, st.SResultCode)
		if err != nil {
			return err
		}
	}
	if false || st.SErrCode != "" {
		err = p.WriteString(10, st.SErrCode)
		if err != nil {
			return err
		}
	}
	if false || st.SErrCodeDesc != "" {
		err = p.WriteString(11, st.SErrCodeDesc)
		if err != nil {
			return err
		}
	}
	if false || st.SOpenId != "" {
		err = p.WriteString(12, st.SOpenId)
		if err != nil {
			return err
		}
	}
	if false || st.SFeeType != "" {
		err = p.WriteString(13, st.SFeeType)
		if err != nil {
			return err
		}
	}
	if false || st.ITotalFee != 0 {
		err = p.WriteUint32(14, st.ITotalFee)
		if err != nil {
			return err
		}
	}
	if false || st.ICouponFee != 0 {
		err = p.WriteUint32(15, st.ICouponFee)
		if err != nil {
			return err
		}
	}
	if false || st.STransactionId != "" {
		err = p.WriteString(16, st.STransactionId)
		if err != nil {
			return err
		}
	}
	if false || st.SOutTradeNo != "" {
		err = p.WriteString(17, st.SOutTradeNo)
		if err != nil {
			return err
		}
	}
	if false || st.STimeEnd != "" {
		err = p.WriteString(18, st.STimeEnd)
		if err != nil {
			return err
		}
	}
	if false || st.SBuyerLogonId != "" {
		err = p.WriteString(19, st.SBuyerLogonId)
		if err != nil {
			return err
		}
	}
	if false || st.SFundBillList != "" {
		err = p.WriteString(20, st.SFundBillList)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *HeePayH5Receipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type IAPReceiptInAll struct {
	IReceiptType      uint32          `json:"iReceiptType"`
	StStatus          IAPStatus       `json:"stStatus"`
	StAppleReceipt    AppleReceipt    `json:"stAppleReceipt"`
	StGoogleReceipt   GoogleReceipt   `json:"stGoogleReceipt"`
	StFyReceipt       FyReceipt       `json:"stFyReceipt"`
	StHeePayReceipt   HeePayReceipt   `json:"stHeePayReceipt"`
	StHeePayH5Receipt HeePayH5Receipt `json:"stHeePayH5Receipt"`
}

func (st *IAPReceiptInAll) resetDefault() {
	st.StStatus.resetDefault()
	st.StAppleReceipt.resetDefault()
	st.StGoogleReceipt.resetDefault()
	st.StFyReceipt.resetDefault()
	st.StHeePayReceipt.resetDefault()
	st.StHeePayH5Receipt.resetDefault()
}
func (st *IAPReceiptInAll) Copy() *IAPReceiptInAll {
	ret := NewIAPReceiptInAll()
	ret.IReceiptType = st.IReceiptType
	ret.StStatus = *(st.StStatus.Copy())
	ret.StAppleReceipt = *(st.StAppleReceipt.Copy())
	ret.StGoogleReceipt = *(st.StGoogleReceipt.Copy())
	ret.StFyReceipt = *(st.StFyReceipt.Copy())
	ret.StHeePayReceipt = *(st.StHeePayReceipt.Copy())
	ret.StHeePayH5Receipt = *(st.StHeePayH5Receipt.Copy())
	return ret
}
func NewIAPReceiptInAll() *IAPReceiptInAll {
	ret := &IAPReceiptInAll{}
	ret.resetDefault()
	return ret
}
func (st *IAPReceiptInAll) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iReceiptType")+fmt.Sprintf("%v\n", st.IReceiptType))
	util.Tab(buff, t+1, util.Fieldname("stStatus")+"{\n")
	st.StStatus.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("stAppleReceipt")+"{\n")
	st.StAppleReceipt.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("stGoogleReceipt")+"{\n")
	st.StGoogleReceipt.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("stFyReceipt")+"{\n")
	st.StFyReceipt.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("stHeePayReceipt")+"{\n")
	st.StHeePayReceipt.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
	util.Tab(buff, t+1, util.Fieldname("stHeePayH5Receipt")+"{\n")
	st.StHeePayH5Receipt.Visit(buff, t+1+1)
	util.Tab(buff, t+1, "}\n")
}
func (st *IAPReceiptInAll) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IReceiptType, 0, false)
	if err != nil {
		return err
	}
	err = st.StStatus.ReadStructFromTag(up, 1, false)
	if err != nil {
		return err
	}
	err = st.StAppleReceipt.ReadStructFromTag(up, 2, false)
	if err != nil {
		return err
	}
	err = st.StGoogleReceipt.ReadStructFromTag(up, 3, false)
	if err != nil {
		return err
	}
	err = st.StFyReceipt.ReadStructFromTag(up, 4, false)
	if err != nil {
		return err
	}
	err = st.StHeePayReceipt.ReadStructFromTag(up, 5, false)
	if err != nil {
		return err
	}
	err = st.StHeePayH5Receipt.ReadStructFromTag(up, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *IAPReceiptInAll) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *IAPReceiptInAll) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IReceiptType != 0 {
		err = p.WriteUint32(0, st.IReceiptType)
		if err != nil {
			return err
		}
	}
	err = st.StStatus.WriteStructFromTag(p, 1, false)
	if err != nil {
		return err
	}
	err = st.StAppleReceipt.WriteStructFromTag(p, 2, false)
	if err != nil {
		return err
	}
	err = st.StGoogleReceipt.WriteStructFromTag(p, 3, false)
	if err != nil {
		return err
	}
	err = st.StFyReceipt.WriteStructFromTag(p, 4, false)
	if err != nil {
		return err
	}
	err = st.StHeePayReceipt.WriteStructFromTag(p, 5, false)
	if err != nil {
		return err
	}
	err = st.StHeePayH5Receipt.WriteStructFromTag(p, 6, false)
	if err != nil {
		return err
	}

	_ = length
	return err
}
func (st *IAPReceiptInAll) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type ReceiptQueryParam struct {
	IDeliverZoneId uint32 `json:"iDeliverZoneId"`
	IDeliverRoleId uint64 `json:"iDeliverRoleId"`
	IReceiptType   uint32 `json:"iReceiptType"`
	IReceiptStatus uint32 `json:"iReceiptStatus"`
	IAddTimeBegin  uint32 `json:"iAddTimeBegin"`
	IAddTimeEnd    uint32 `json:"iAddTimeEnd"`
}

func (st *ReceiptQueryParam) resetDefault() {
}
func (st *ReceiptQueryParam) Copy() *ReceiptQueryParam {
	ret := NewReceiptQueryParam()
	ret.IDeliverZoneId = st.IDeliverZoneId
	ret.IDeliverRoleId = st.IDeliverRoleId
	ret.IReceiptType = st.IReceiptType
	ret.IReceiptStatus = st.IReceiptStatus
	ret.IAddTimeBegin = st.IAddTimeBegin
	ret.IAddTimeEnd = st.IAddTimeEnd
	return ret
}
func NewReceiptQueryParam() *ReceiptQueryParam {
	ret := &ReceiptQueryParam{}
	ret.resetDefault()
	return ret
}
func (st *ReceiptQueryParam) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("iDeliverZoneId")+fmt.Sprintf("%v\n", st.IDeliverZoneId))
	util.Tab(buff, t+1, util.Fieldname("iDeliverRoleId")+fmt.Sprintf("%v\n", st.IDeliverRoleId))
	util.Tab(buff, t+1, util.Fieldname("iReceiptType")+fmt.Sprintf("%v\n", st.IReceiptType))
	util.Tab(buff, t+1, util.Fieldname("iReceiptStatus")+fmt.Sprintf("%v\n", st.IReceiptStatus))
	util.Tab(buff, t+1, util.Fieldname("iAddTimeBegin")+fmt.Sprintf("%v\n", st.IAddTimeBegin))
	util.Tab(buff, t+1, util.Fieldname("iAddTimeEnd")+fmt.Sprintf("%v\n", st.IAddTimeEnd))
}
func (st *ReceiptQueryParam) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadUint32(&st.IDeliverZoneId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint64(&st.IDeliverRoleId, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IReceiptType, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IReceiptStatus, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IAddTimeBegin, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IAddTimeEnd, 5, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ReceiptQueryParam) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ReceiptQueryParam) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.IDeliverZoneId != 0 {
		err = p.WriteUint32(0, st.IDeliverZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.IDeliverRoleId != 0 {
		err = p.WriteUint64(1, st.IDeliverRoleId)
		if err != nil {
			return err
		}
	}
	if false || st.IReceiptType != 0 {
		err = p.WriteUint32(2, st.IReceiptType)
		if err != nil {
			return err
		}
	}
	if false || st.IReceiptStatus != 0 {
		err = p.WriteUint32(3, st.IReceiptStatus)
		if err != nil {
			return err
		}
	}
	if false || st.IAddTimeBegin != 0 {
		err = p.WriteUint32(4, st.IAddTimeBegin)
		if err != nil {
			return err
		}
	}
	if false || st.IAddTimeEnd != 0 {
		err = p.WriteUint32(5, st.IAddTimeEnd)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ReceiptQueryParam) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type IAPService struct {
	proxy model.ServicePrxImpl
}

func (s *IAPService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *IAPService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *IAPService) VerifyAppleReceipt(stPurchase ApplePurchase, stReceipt *AppleReceipt, sErrorInfo *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stPurchase.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("verifyAppleReceipt", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stReceipt).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sErrorInfo), 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) DeliverAppleReceipt(iRoleId uint64, iZoneId uint32, stPurchase ApplePurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iRoleId != 0 {
		err = p.WriteUint64(1, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(2, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	err = stPurchase.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	if true || iProxyRoleId != 0 {
		err = p.WriteUint64(4, iProxyRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProxyZoneId != 0 {
		err = p.WriteUint32(5, iProxyZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deliverAppleReceipt", p.ToBytes(), &rsp)
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
func (s *IAPService) GetAppleReceiptStatus(stPurchase ApplePurchase, stReceipt *AppleReceipt, stStatus *IAPStatus) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stPurchase.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAppleReceiptStatus", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stReceipt).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = (*stStatus).ReadStructFromTag(up, 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) GetTransactionStatus(sTransactionId string, stReceipt *AppleReceipt, stStatus *IAPStatus) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sTransactionId != "" {
		err = p.WriteString(1, sTransactionId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getTransactionStatus", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stReceipt).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = (*stStatus).ReadStructFromTag(up, 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) VerifyGoogleReceipt(stPurchase GooglePurchase, stReceipt *GoogleReceipt, sErrorInfo *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stPurchase.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("verifyGoogleReceipt", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stReceipt).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sErrorInfo), 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) DeliverGoogleReceipt(iRoleId uint64, iZoneId uint32, stPurchase GooglePurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iRoleId != 0 {
		err = p.WriteUint64(1, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(2, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	err = stPurchase.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	if true || iProxyRoleId != 0 {
		err = p.WriteUint64(4, iProxyRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProxyZoneId != 0 {
		err = p.WriteUint32(5, iProxyZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deliverGoogleReceipt", p.ToBytes(), &rsp)
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
func (s *IAPService) GetGoogleReceiptStatus(sOrderId string, stReceipt *GoogleReceipt, stStatus *IAPStatus) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sOrderId != "" {
		err = p.WriteString(1, sOrderId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getGoogleReceiptStatus", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stReceipt).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	err = (*stStatus).ReadStructFromTag(up, 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) DeliverFyReceipt(iRoleId uint64, iZoneId uint32, stPurchase FyPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iRoleId != 0 {
		err = p.WriteUint64(1, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(2, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	err = stPurchase.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	if true || iProxyRoleId != 0 {
		err = p.WriteUint64(4, iProxyRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProxyZoneId != 0 {
		err = p.WriteUint32(5, iProxyZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deliverFyReceipt", p.ToBytes(), &rsp)
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
func (s *IAPService) DeliverHeePayReceipt(iRoleId uint64, iZoneId uint32, stPurchase HeePayPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iRoleId != 0 {
		err = p.WriteUint64(1, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(2, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	err = stPurchase.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	if true || iProxyRoleId != 0 {
		err = p.WriteUint64(4, iProxyRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProxyZoneId != 0 {
		err = p.WriteUint32(5, iProxyZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deliverHeePayReceipt", p.ToBytes(), &rsp)
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
func (s *IAPService) DeliverHeePayH5Receipt(iRoleId uint64, iZoneId uint32, stPurchase HeePayH5Purchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iRoleId != 0 {
		err = p.WriteUint64(1, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(2, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	err = stPurchase.WriteStructFromTag(p, 3, true)
	if err != nil {
		return ret, err
	}
	if true || iProxyRoleId != 0 {
		err = p.WriteUint64(4, iProxyRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProxyZoneId != 0 {
		err = p.WriteUint32(5, iProxyZoneId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("deliverHeePayH5Receipt", p.ToBytes(), &rsp)
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
func (s *IAPService) GetReceiptStatusByFlow(sFlowId string, stIAPReceiptInAll *IAPReceiptInAll) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sFlowId != "" {
		err = p.WriteString(1, sFlowId)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getReceiptStatusByFlow", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stIAPReceiptInAll).ReadStructFromTag(up, 2, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}
func (s *IAPService) GetReceiptStatusList(stQueryParam ReceiptQueryParam, vIAPReceiptInAll *[]IAPReceiptInAll) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	err = stQueryParam.WriteStructFromTag(p, 1, true)
	if err != nil {
		return ret, err
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getReceiptStatusList", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	has, ty, err = up.SkipToTag(2, true)
	if err != nil {
		return ret, err
	}
	if has {
		if ty != codec.SdpType_Vector {
			return ret, fmt.Errorf("tag:%d got wrong type %d", 2, ty)
		}

		_, length, err = up.ReadNumber32()
		if err != nil {
			return ret, err
		}
		(*vIAPReceiptInAll) = make([]IAPReceiptInAll, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vIAPReceiptInAll)[i].ReadStructFromTag(up, 0, true)
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
func (s *IAPService) CreateOrder(iReceiptType uint32, sFlowId string, iProductId uint32, iRoleId uint64, iZoneId uint32, sPayload string, stOrder *IAPTmpOrder) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || iReceiptType != 0 {
		err = p.WriteUint32(1, iReceiptType)
		if err != nil {
			return ret, err
		}
	}
	if true || sFlowId != "" {
		err = p.WriteString(2, sFlowId)
		if err != nil {
			return ret, err
		}
	}
	if true || iProductId != 0 {
		err = p.WriteUint32(3, iProductId)
		if err != nil {
			return ret, err
		}
	}
	if true || iRoleId != 0 {
		err = p.WriteUint64(4, iRoleId)
		if err != nil {
			return ret, err
		}
	}
	if true || iZoneId != 0 {
		err = p.WriteUint32(5, iZoneId)
		if err != nil {
			return ret, err
		}
	}
	if true || sPayload != "" {
		err = p.WriteString(6, sPayload)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("createOrder", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = (*stOrder).ReadStructFromTag(up, 7, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _IAPServiceImpl interface {
	VerifyAppleReceipt(ctx context.Context, stPurchase ApplePurchase, stReceipt *AppleReceipt, sErrorInfo *string) (int32, error)
	DeliverAppleReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase ApplePurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
	GetAppleReceiptStatus(ctx context.Context, stPurchase ApplePurchase, stReceipt *AppleReceipt, stStatus *IAPStatus) (int32, error)
	GetTransactionStatus(ctx context.Context, sTransactionId string, stReceipt *AppleReceipt, stStatus *IAPStatus) (int32, error)
	VerifyGoogleReceipt(ctx context.Context, stPurchase GooglePurchase, stReceipt *GoogleReceipt, sErrorInfo *string) (int32, error)
	DeliverGoogleReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase GooglePurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
	GetGoogleReceiptStatus(ctx context.Context, sOrderId string, stReceipt *GoogleReceipt, stStatus *IAPStatus) (int32, error)
	DeliverFyReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase FyPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
	DeliverHeePayReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase HeePayPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
	DeliverHeePayH5Receipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase HeePayH5Purchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
	GetReceiptStatusByFlow(ctx context.Context, sFlowId string, stIAPReceiptInAll *IAPReceiptInAll) (int32, error)
	GetReceiptStatusList(ctx context.Context, stQueryParam ReceiptQueryParam, vIAPReceiptInAll *[]IAPReceiptInAll) (int32, error)
	CreateOrder(ctx context.Context, iReceiptType uint32, sFlowId string, iProductId uint32, iRoleId uint64, iZoneId uint32, sPayload string, stOrder *IAPTmpOrder) (int32, error)
}

func _IAPServiceVerifyAppleReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 ApplePurchase
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 AppleReceipt
	var p3 string
	var ret int32
	ret, err = impl.VerifyAppleReceipt(ctx, p1, &p2, &p3)
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
	if true || p3 != "" {
		err = p.WriteString(3, p3)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _IAPServiceDeliverAppleReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint64
	err = up.ReadUint64(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 ApplePurchase
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeliverAppleReceipt(ctx, p1, p2, p3, p4, p5)
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
func _IAPServiceGetAppleReceiptStatusImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 ApplePurchase
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 AppleReceipt
	var p3 IAPStatus
	var ret int32
	ret, err = impl.GetAppleReceiptStatus(ctx, p1, &p2, &p3)
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
func _IAPServiceGetTransactionStatusImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 AppleReceipt
	var p3 IAPStatus
	var ret int32
	ret, err = impl.GetTransactionStatus(ctx, p1, &p2, &p3)
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
func _IAPServiceVerifyGoogleReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 GooglePurchase
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 GoogleReceipt
	var p3 string
	var ret int32
	ret, err = impl.VerifyGoogleReceipt(ctx, p1, &p2, &p3)
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
	if true || p3 != "" {
		err = p.WriteString(3, p3)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}
func _IAPServiceDeliverGoogleReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint64
	err = up.ReadUint64(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 GooglePurchase
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeliverGoogleReceipt(ctx, p1, p2, p3, p4, p5)
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
func _IAPServiceGetGoogleReceiptStatusImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 GoogleReceipt
	var p3 IAPStatus
	var ret int32
	ret, err = impl.GetGoogleReceiptStatus(ctx, p1, &p2, &p3)
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
func _IAPServiceDeliverFyReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint64
	err = up.ReadUint64(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 FyPurchase
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeliverFyReceipt(ctx, p1, p2, p3, p4, p5)
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
func _IAPServiceDeliverHeePayReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint64
	err = up.ReadUint64(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 HeePayPurchase
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeliverHeePayReceipt(ctx, p1, p2, p3, p4, p5)
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
func _IAPServiceDeliverHeePayH5ReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint64
	err = up.ReadUint64(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 uint32
	err = up.ReadUint32(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 HeePayH5Purchase
	err = p3.ReadStructFromTag(up, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DeliverHeePayH5Receipt(ctx, p1, p2, p3, p4, p5)
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
func _IAPServiceGetReceiptStatusByFlowImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 IAPReceiptInAll
	var ret int32
	ret, err = impl.GetReceiptStatusByFlow(ctx, p1, &p2)
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
func _IAPServiceGetReceiptStatusListImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 ReceiptQueryParam
	err = p1.ReadStructFromTag(up, 1, true)
	if err != nil {
		return err
	}
	var p2 []IAPReceiptInAll
	var ret int32
	ret, err = impl.GetReceiptStatusList(ctx, p1, &p2)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}

	length = uint32(len(p2))
	if true || length != 0 {
		err = p.WriteHeader(2, codec.SdpType_Vector)
		if err != nil {
			return err
		}
		err = p.WriteNumber32(length)
		if err != nil {
			return err
		}
		for _, v := range p2 {
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
func _IAPServiceCreateOrderImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_IAPServiceImpl)
	var p1 uint32
	err = up.ReadUint32(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 string
	err = up.ReadString(&p2, 2, true)
	if err != nil {
		return err
	}
	var p3 uint32
	err = up.ReadUint32(&p3, 3, true)
	if err != nil {
		return err
	}
	var p4 uint64
	err = up.ReadUint64(&p4, 4, true)
	if err != nil {
		return err
	}
	var p5 uint32
	err = up.ReadUint32(&p5, 5, true)
	if err != nil {
		return err
	}
	var p6 string
	err = up.ReadString(&p6, 6, true)
	if err != nil {
		return err
	}
	var p7 IAPTmpOrder
	var ret int32
	ret, err = impl.CreateOrder(ctx, p1, p2, p3, p4, p5, p6, &p7)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
		if err != nil {
			return err
		}
	}
	err = p7.WriteStructFromTag(p, 7, true)
	if err != nil {
		return err
	}
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *IAPService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "verifyAppleReceipt":
		err = _IAPServiceVerifyAppleReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deliverAppleReceipt":
		err = _IAPServiceDeliverAppleReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAppleReceiptStatus":
		err = _IAPServiceGetAppleReceiptStatusImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getTransactionStatus":
		err = _IAPServiceGetTransactionStatusImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "verifyGoogleReceipt":
		err = _IAPServiceVerifyGoogleReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deliverGoogleReceipt":
		err = _IAPServiceDeliverGoogleReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getGoogleReceiptStatus":
		err = _IAPServiceGetGoogleReceiptStatusImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deliverFyReceipt":
		err = _IAPServiceDeliverFyReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deliverHeePayReceipt":
		err = _IAPServiceDeliverHeePayReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "deliverHeePayH5Receipt":
		err = _IAPServiceDeliverHeePayH5ReceiptImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getReceiptStatusByFlow":
		err = _IAPServiceGetReceiptStatusByFlowImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getReceiptStatusList":
		err = _IAPServiceGetReceiptStatusListImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "createOrder":
		err = _IAPServiceCreateOrderImpl(ctx, serviceImpl, up, p)
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
