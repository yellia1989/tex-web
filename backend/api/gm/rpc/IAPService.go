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
	IAPReceiptType_Apple  = 1
	IAPReceiptType_Google = 2
	IAPReceiptType_Fy     = 3
	IAPReceiptType_AB     = 4
)

func (en IAPReceiptType) String() string {
	ret := ""
	switch en {
	case IAPReceiptType_Apple:
		ret = "IAPReceiptType_Apple"
	case IAPReceiptType_Google:
		ret = "IAPReceiptType_Google"
	case IAPReceiptType_Fy:
		ret = "IAPReceiptType_Fy"
	case IAPReceiptType_AB:
		ret = "IAPReceiptType_AB"
	}
	return ret
}

type IAPReceiptStatus int32

const (
	IAPReceiptStatus_Pending         = 1
	IAPReceiptStatus_Verify_Fail     = 2
	IAPReceiptStatus_Delivering      = 3
	IAPReceiptStatus_Deliver_Success = 4
	IAPReceiptStatus_Deliver_Fail    = 5
)

func (en IAPReceiptStatus) String() string {
	ret := ""
	switch en {
	case IAPReceiptStatus_Pending:
		ret = "IAPReceiptStatus_Pending"
	case IAPReceiptStatus_Verify_Fail:
		ret = "IAPReceiptStatus_Verify_Fail"
	case IAPReceiptStatus_Delivering:
		ret = "IAPReceiptStatus_Delivering"
	case IAPReceiptStatus_Deliver_Success:
		ret = "IAPReceiptStatus_Deliver_Success"
	case IAPReceiptStatus_Deliver_Fail:
		ret = "IAPReceiptStatus_Deliver_Fail"
	}
	return ret
}

type IAPStatus struct {
	IReceiptId      uint32 `json:"iReceiptId" form:"iReceiptId"`
	IReceiptStatus  uint32 `json:"iReceiptStatus" form:"iReceiptStatus"`
	IDeliverRoleId  uint64 `json:"iDeliverRoleId" form:"iDeliverRoleId"`
	IDeliverZoneId  uint32 `json:"iDeliverZoneId" form:"iDeliverZoneId"`
	IProxyRoleId    uint64 `json:"iProxyRoleId" form:"iProxyRoleId"`
	IProxyZoneId    uint32 `json:"iProxyZoneId" form:"iProxyZoneId"`
	IAddTime        uint32 `json:"iAddTime" form:"iAddTime"`
	IVerifyTime     uint32 `json:"iVerifyTime" form:"iVerifyTime"`
	IDeliverTime    uint32 `json:"iDeliverTime" form:"iDeliverTime"`
	IRetryNum       uint32 `json:"iRetryNum" form:"iRetryNum"`
	INextTryTime    uint32 `json:"iNextTryTime" form:"iNextTryTime"`
	SFailReason     string `json:"sFailReason" form:"sFailReason"`
	SDeliverItem    string `json:"sDeliverItem" form:"sDeliverItem"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
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
	SFlowId      string `json:"sFlowId" form:"sFlowId"`
	IProductId   uint32 `json:"iProductId" form:"iProductId"`
	IReceiptType uint32 `json:"iReceiptType" form:"iReceiptType"`
	IRoleId      uint64 `json:"iRoleId" form:"iRoleId"`
	IZoneId      uint32 `json:"iZoneId" form:"iZoneId"`
	ICreateTime  uint32 `json:"iCreateTime" form:"iCreateTime"`
	SPayload     string `json:"sPayload" form:"sPayload"`
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
	SReceiptData    string `json:"sReceiptData" form:"sReceiptData"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
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
	IQuantity          uint32 `json:"iQuantity" form:"iQuantity"`
	SProductId         string `json:"sProductId" form:"sProductId"`
	STransactionId     string `json:"sTransactionId" form:"sTransactionId"`
	IPurchaseDate      uint32 `json:"iPurchaseDate" form:"iPurchaseDate"`
	SBId               string `json:"sBId" form:"sBId"`
	SBVrs              string `json:"sBVrs" form:"sBVrs"`
	SOriTransactionId  string `json:"sOriTransactionId" form:"sOriTransactionId"`
	IOriPurchaseDate   uint32 `json:"iOriPurchaseDate" form:"iOriPurchaseDate"`
	SAppItemId         string `json:"sAppItemId" form:"sAppItemId"`
	SVersionExternalId string `json:"sVersionExternalId" form:"sVersionExternalId"`
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
	IResponseCode   int32  `json:"iResponseCode" form:"iResponseCode"`
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
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
	SOrderId          string `json:"sOrderId" form:"sOrderId"`
	SProductId        string `json:"sProductId" form:"sProductId"`
	SPackageName      string `json:"sPackageName" form:"sPackageName"`
	IPurchaseTime     uint32 `json:"iPurchaseTime" form:"iPurchaseTime"`
	IPurchaseState    uint32 `json:"iPurchaseState" form:"iPurchaseState"`
	SDeveloperPayload string `json:"sDeveloperPayload" form:"sDeveloperPayload"`
	SPurchaseToken    string `json:"sPurchaseToken" form:"sPurchaseToken"`
	BAutoRenewing     bool   `json:"bAutoRenewing" form:"bAutoRenewing"`
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
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	IConnId         uint32 `json:"iConnId" form:"iConnId"`
	SChannel        string `json:"sChannel" form:"sChannel"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
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
	STradeStatus string `json:"sTradeStatus" form:"sTradeStatus"`
	STradeNo     string `json:"sTradeNo" form:"sTradeNo"`
	STradeTime   string `json:"sTradeTime" form:"sTradeTime"`
	SOutTradeNo  string `json:"sOutTradeNo" form:"sOutTradeNo"`
	ITotalAmount uint32 `json:"iTotalAmount" form:"iTotalAmount"`
	SGoodsId     string `json:"sGoodsId" form:"sGoodsId"`
	SAppId       string `json:"sAppId" form:"sAppId"`
	SPlayerId    string `json:"sPlayerId" form:"sPlayerId"`
	SOpenId      string `json:"sOpenId" form:"sOpenId"`
	IServerId    uint32 `json:"iServerId" form:"iServerId"`
	SChannelId   string `json:"sChannelId" form:"sChannelId"`
	ISandBox     uint32 `json:"iSandBox" form:"iSandBox"`
	ITimeStamp   uint32 `json:"iTimeStamp" form:"iTimeStamp"`
	SNotifyExt   string `json:"sNotifyExt" form:"sNotifyExt"`
}

func (st *FyReceipt) resetDefault() {
}
func (st *FyReceipt) Copy() *FyReceipt {
	ret := NewFyReceipt()
	ret.STradeStatus = st.STradeStatus
	ret.STradeNo = st.STradeNo
	ret.STradeTime = st.STradeTime
	ret.SOutTradeNo = st.SOutTradeNo
	ret.ITotalAmount = st.ITotalAmount
	ret.SGoodsId = st.SGoodsId
	ret.SAppId = st.SAppId
	ret.SPlayerId = st.SPlayerId
	ret.SOpenId = st.SOpenId
	ret.IServerId = st.IServerId
	ret.SChannelId = st.SChannelId
	ret.ISandBox = st.ISandBox
	ret.ITimeStamp = st.ITimeStamp
	ret.SNotifyExt = st.SNotifyExt
	return ret
}
func NewFyReceipt() *FyReceipt {
	ret := &FyReceipt{}
	ret.resetDefault()
	return ret
}
func (st *FyReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sTradeStatus")+fmt.Sprintf("%v\n", st.STradeStatus))
	util.Tab(buff, t+1, util.Fieldname("sTradeNo")+fmt.Sprintf("%v\n", st.STradeNo))
	util.Tab(buff, t+1, util.Fieldname("sTradeTime")+fmt.Sprintf("%v\n", st.STradeTime))
	util.Tab(buff, t+1, util.Fieldname("sOutTradeNo")+fmt.Sprintf("%v\n", st.SOutTradeNo))
	util.Tab(buff, t+1, util.Fieldname("iTotalAmount")+fmt.Sprintf("%v\n", st.ITotalAmount))
	util.Tab(buff, t+1, util.Fieldname("sGoodsId")+fmt.Sprintf("%v\n", st.SGoodsId))
	util.Tab(buff, t+1, util.Fieldname("sAppId")+fmt.Sprintf("%v\n", st.SAppId))
	util.Tab(buff, t+1, util.Fieldname("sPlayerId")+fmt.Sprintf("%v\n", st.SPlayerId))
	util.Tab(buff, t+1, util.Fieldname("sOpenId")+fmt.Sprintf("%v\n", st.SOpenId))
	util.Tab(buff, t+1, util.Fieldname("iServerId")+fmt.Sprintf("%v\n", st.IServerId))
	util.Tab(buff, t+1, util.Fieldname("sChannelId")+fmt.Sprintf("%v\n", st.SChannelId))
	util.Tab(buff, t+1, util.Fieldname("iSandBox")+fmt.Sprintf("%v\n", st.ISandBox))
	util.Tab(buff, t+1, util.Fieldname("iTimeStamp")+fmt.Sprintf("%v\n", st.ITimeStamp))
	util.Tab(buff, t+1, util.Fieldname("sNotifyExt")+fmt.Sprintf("%v\n", st.SNotifyExt))
}
func (st *FyReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.STradeStatus, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STradeNo, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STradeTime, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOutTradeNo, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITotalAmount, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SGoodsId, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppId, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPlayerId, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOpenId, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IServerId, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SChannelId, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ISandBox, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITimeStamp, 12, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SNotifyExt, 13, false)
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
	if false || st.STradeStatus != "" {
		err = p.WriteString(0, st.STradeStatus)
		if err != nil {
			return err
		}
	}
	if false || st.STradeNo != "" {
		err = p.WriteString(1, st.STradeNo)
		if err != nil {
			return err
		}
	}
	if false || st.STradeTime != "" {
		err = p.WriteString(2, st.STradeTime)
		if err != nil {
			return err
		}
	}
	if false || st.SOutTradeNo != "" {
		err = p.WriteString(3, st.SOutTradeNo)
		if err != nil {
			return err
		}
	}
	if false || st.ITotalAmount != 0 {
		err = p.WriteUint32(4, st.ITotalAmount)
		if err != nil {
			return err
		}
	}
	if false || st.SGoodsId != "" {
		err = p.WriteString(5, st.SGoodsId)
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
	if false || st.SPlayerId != "" {
		err = p.WriteString(7, st.SPlayerId)
		if err != nil {
			return err
		}
	}
	if false || st.SOpenId != "" {
		err = p.WriteString(8, st.SOpenId)
		if err != nil {
			return err
		}
	}
	if false || st.IServerId != 0 {
		err = p.WriteUint32(9, st.IServerId)
		if err != nil {
			return err
		}
	}
	if false || st.SChannelId != "" {
		err = p.WriteString(10, st.SChannelId)
		if err != nil {
			return err
		}
	}
	if false || st.ISandBox != 0 {
		err = p.WriteUint32(11, st.ISandBox)
		if err != nil {
			return err
		}
	}
	if false || st.ITimeStamp != 0 {
		err = p.WriteUint32(12, st.ITimeStamp)
		if err != nil {
			return err
		}
	}
	if false || st.SNotifyExt != "" {
		err = p.WriteString(13, st.SNotifyExt)
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

type ABPurchase struct {
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	IConnId         uint32 `json:"iConnId" form:"iConnId"`
	SChannel        string `json:"sChannel" form:"sChannel"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
}

func (st *ABPurchase) resetDefault() {
}
func (st *ABPurchase) Copy() *ABPurchase {
	ret := NewABPurchase()
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.IConnId = st.IConnId
	ret.SChannel = st.SChannel
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewABPurchase() *ABPurchase {
	ret := &ABPurchase{}
	ret.resetDefault()
	return ret
}
func (st *ABPurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("iConnId")+fmt.Sprintf("%v\n", st.IConnId))
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *ABPurchase) ReadStruct(up *codec.UnPacker) error {
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
func (st *ABPurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ABPurchase) WriteStruct(p *codec.Packer) error {
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
func (st *ABPurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type ABReceipt struct {
	STransid   string  `json:"sTransid" form:"sTransid"`
	ITranstype uint32  `json:"iTranstype" form:"iTranstype"`
	SCporderid string  `json:"sCporderid" form:"sCporderid"`
	SAppuserid string  `json:"sAppuserid" form:"sAppuserid"`
	SAppid     string  `json:"sAppid" form:"sAppid"`
	IWaresid   uint32  `json:"iWaresid" form:"iWaresid"`
	IFeetype   uint32  `json:"iFeetype" form:"iFeetype"`
	FMoney     float32 `json:"fMoney" form:"fMoney"`
	SCurrency  string  `json:"sCurrency" form:"sCurrency"`
	IResult    uint32  `json:"iResult" form:"iResult"`
	ITranstime uint32  `json:"iTranstime" form:"iTranstime"`
	SCpprivate string  `json:"sCpprivate" form:"sCpprivate"`
	IPaytype   uint32  `json:"iPaytype" form:"iPaytype"`
}

func (st *ABReceipt) resetDefault() {
}
func (st *ABReceipt) Copy() *ABReceipt {
	ret := NewABReceipt()
	ret.STransid = st.STransid
	ret.ITranstype = st.ITranstype
	ret.SCporderid = st.SCporderid
	ret.SAppuserid = st.SAppuserid
	ret.SAppid = st.SAppid
	ret.IWaresid = st.IWaresid
	ret.IFeetype = st.IFeetype
	ret.FMoney = st.FMoney
	ret.SCurrency = st.SCurrency
	ret.IResult = st.IResult
	ret.ITranstime = st.ITranstime
	ret.SCpprivate = st.SCpprivate
	ret.IPaytype = st.IPaytype
	return ret
}
func NewABReceipt() *ABReceipt {
	ret := &ABReceipt{}
	ret.resetDefault()
	return ret
}
func (st *ABReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sTransid")+fmt.Sprintf("%v\n", st.STransid))
	util.Tab(buff, t+1, util.Fieldname("iTranstype")+fmt.Sprintf("%v\n", st.ITranstype))
	util.Tab(buff, t+1, util.Fieldname("sCporderid")+fmt.Sprintf("%v\n", st.SCporderid))
	util.Tab(buff, t+1, util.Fieldname("sAppuserid")+fmt.Sprintf("%v\n", st.SAppuserid))
	util.Tab(buff, t+1, util.Fieldname("sAppid")+fmt.Sprintf("%v\n", st.SAppid))
	util.Tab(buff, t+1, util.Fieldname("iWaresid")+fmt.Sprintf("%v\n", st.IWaresid))
	util.Tab(buff, t+1, util.Fieldname("iFeetype")+fmt.Sprintf("%v\n", st.IFeetype))
	util.Tab(buff, t+1, util.Fieldname("fMoney")+fmt.Sprintf("%v\n", st.FMoney))
	util.Tab(buff, t+1, util.Fieldname("sCurrency")+fmt.Sprintf("%v\n", st.SCurrency))
	util.Tab(buff, t+1, util.Fieldname("iResult")+fmt.Sprintf("%v\n", st.IResult))
	util.Tab(buff, t+1, util.Fieldname("iTranstime")+fmt.Sprintf("%v\n", st.ITranstime))
	util.Tab(buff, t+1, util.Fieldname("sCpprivate")+fmt.Sprintf("%v\n", st.SCpprivate))
	util.Tab(buff, t+1, util.Fieldname("iPaytype")+fmt.Sprintf("%v\n", st.IPaytype))
}
func (st *ABReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.STransid, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITranstype, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCporderid, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppuserid, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAppid, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IWaresid, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IFeetype, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadFloat32(&st.FMoney, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCurrency, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IResult, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.ITranstime, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCpprivate, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPaytype, 12, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ABReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ABReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.STransid != "" {
		err = p.WriteString(0, st.STransid)
		if err != nil {
			return err
		}
	}
	if false || st.ITranstype != 0 {
		err = p.WriteUint32(1, st.ITranstype)
		if err != nil {
			return err
		}
	}
	if false || st.SCporderid != "" {
		err = p.WriteString(2, st.SCporderid)
		if err != nil {
			return err
		}
	}
	if false || st.SAppuserid != "" {
		err = p.WriteString(3, st.SAppuserid)
		if err != nil {
			return err
		}
	}
	if false || st.SAppid != "" {
		err = p.WriteString(4, st.SAppid)
		if err != nil {
			return err
		}
	}
	if false || st.IWaresid != 0 {
		err = p.WriteUint32(5, st.IWaresid)
		if err != nil {
			return err
		}
	}
	if false || st.IFeetype != 0 {
		err = p.WriteUint32(6, st.IFeetype)
		if err != nil {
			return err
		}
	}
	if false || st.FMoney != 0 {
		err = p.WriteFloat32(7, st.FMoney)
		if err != nil {
			return err
		}
	}
	if false || st.SCurrency != "" {
		err = p.WriteString(8, st.SCurrency)
		if err != nil {
			return err
		}
	}
	if false || st.IResult != 0 {
		err = p.WriteUint32(9, st.IResult)
		if err != nil {
			return err
		}
	}
	if false || st.ITranstime != 0 {
		err = p.WriteUint32(10, st.ITranstime)
		if err != nil {
			return err
		}
	}
	if false || st.SCpprivate != "" {
		err = p.WriteString(11, st.SCpprivate)
		if err != nil {
			return err
		}
	}
	if false || st.IPaytype != 0 {
		err = p.WriteUint32(12, st.IPaytype)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ABReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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
	IReceiptType    uint32        `json:"iReceiptType" form:"iReceiptType"`
	StStatus        IAPStatus     `json:"stStatus" form:"stStatus"`
	StAppleReceipt  AppleReceipt  `json:"stAppleReceipt" form:"stAppleReceipt"`
	StGoogleReceipt GoogleReceipt `json:"stGoogleReceipt" form:"stGoogleReceipt"`
	StFyReceipt     FyReceipt     `json:"stFyReceipt" form:"stFyReceipt"`
	StABReceipt     ABReceipt     `json:"stABReceipt" form:"stABReceipt"`
}

func (st *IAPReceiptInAll) resetDefault() {
	st.StStatus.resetDefault()
	st.StAppleReceipt.resetDefault()
	st.StGoogleReceipt.resetDefault()
	st.StFyReceipt.resetDefault()
	st.StABReceipt.resetDefault()
}
func (st *IAPReceiptInAll) Copy() *IAPReceiptInAll {
	ret := NewIAPReceiptInAll()
	ret.IReceiptType = st.IReceiptType
	ret.StStatus = *(st.StStatus.Copy())
	ret.StAppleReceipt = *(st.StAppleReceipt.Copy())
	ret.StGoogleReceipt = *(st.StGoogleReceipt.Copy())
	ret.StFyReceipt = *(st.StFyReceipt.Copy())
	ret.StABReceipt = *(st.StABReceipt.Copy())
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
	util.Tab(buff, t+1, util.Fieldname("stABReceipt")+"{\n")
	st.StABReceipt.Visit(buff, t+1+1)
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
	err = st.StABReceipt.ReadStructFromTag(up, 5, false)
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
	err = st.StABReceipt.WriteStructFromTag(p, 5, false)
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
	IDeliverZoneId uint32 `json:"iDeliverZoneId" form:"iDeliverZoneId"`
	IDeliverRoleId uint64 `json:"iDeliverRoleId" form:"iDeliverRoleId"`
	IReceiptType   uint32 `json:"iReceiptType" form:"iReceiptType"`
	IReceiptStatus uint32 `json:"iReceiptStatus" form:"iReceiptStatus"`
	IAddTimeBegin  uint32 `json:"iAddTimeBegin" form:"iAddTimeBegin"`
	IAddTimeEnd    uint32 `json:"iAddTimeEnd" form:"iAddTimeEnd"`
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
func (s *IAPService) DeliverABReceipt(iRoleId uint64, iZoneId uint32, stPurchase ABPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
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
	err = s.proxy.Invoke("deliverABReceipt", p.ToBytes(), &rsp)
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
	DeliverABReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase ABPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
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
func _IAPServiceDeliverABReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
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
	var p3 ABPurchase
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
	ret, err = impl.DeliverABReceipt(ctx, p1, p2, p3, p4, p5)
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
	case "deliverABReceipt":
		err = _IAPServiceDeliverABReceiptImpl(ctx, serviceImpl, up, p)
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
