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
	IAPReceiptType_GameHub  = 6
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
	case IAPReceiptType_HeePay:
		ret = "IAPReceiptType_HeePay"
	case IAPReceiptType_HeePayH5:
		ret = "IAPReceiptType_HeePayH5"
	case IAPReceiptType_GameHub:
		ret = "IAPReceiptType_GameHub"
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
	SReceiptJson   string `json:"sReceiptJson" form:"sReceiptJson"`
	SEnvironment   string `json:"sEnvironment" form:"sEnvironment"`
	IStatus        uint32 `json:"iStatus" form:"iStatus"`
	SBundleId      string `json:"sBundleId" form:"sBundleId"`
	SProductId     string `json:"sProductId" form:"sProductId"`
	IPurchaseDate  uint32 `json:"iPurchaseDate" form:"iPurchaseDate"`
	IQuantity      uint32 `json:"iQuantity" form:"iQuantity"`
	STransactionId string `json:"sTransactionId" form:"sTransactionId"`
}

func (st *AppleReceipt) resetDefault() {
}
func (st *AppleReceipt) Copy() *AppleReceipt {
	ret := NewAppleReceipt()
	ret.SReceiptJson = st.SReceiptJson
	ret.SEnvironment = st.SEnvironment
	ret.IStatus = st.IStatus
	ret.SBundleId = st.SBundleId
	ret.SProductId = st.SProductId
	ret.IPurchaseDate = st.IPurchaseDate
	ret.IQuantity = st.IQuantity
	ret.STransactionId = st.STransactionId
	return ret
}
func NewAppleReceipt() *AppleReceipt {
	ret := &AppleReceipt{}
	ret.resetDefault()
	return ret
}
func (st *AppleReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sReceiptJson")+fmt.Sprintf("%v\n", st.SReceiptJson))
	util.Tab(buff, t+1, util.Fieldname("sEnvironment")+fmt.Sprintf("%v\n", st.SEnvironment))
	util.Tab(buff, t+1, util.Fieldname("iStatus")+fmt.Sprintf("%v\n", st.IStatus))
	util.Tab(buff, t+1, util.Fieldname("sBundleId")+fmt.Sprintf("%v\n", st.SBundleId))
	util.Tab(buff, t+1, util.Fieldname("sProductId")+fmt.Sprintf("%v\n", st.SProductId))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseDate")+fmt.Sprintf("%v\n", st.IPurchaseDate))
	util.Tab(buff, t+1, util.Fieldname("iQuantity")+fmt.Sprintf("%v\n", st.IQuantity))
	util.Tab(buff, t+1, util.Fieldname("sTransactionId")+fmt.Sprintf("%v\n", st.STransactionId))
}
func (st *AppleReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SReceiptJson, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SEnvironment, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IStatus, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SBundleId, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProductId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseDate, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IQuantity, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.STransactionId, 7, false)
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
	if false || st.SReceiptJson != "" {
		err = p.WriteString(0, st.SReceiptJson)
		if err != nil {
			return err
		}
	}
	if false || st.SEnvironment != "" {
		err = p.WriteString(1, st.SEnvironment)
		if err != nil {
			return err
		}
	}
	if false || st.IStatus != 0 {
		err = p.WriteUint32(2, st.IStatus)
		if err != nil {
			return err
		}
	}
	if false || st.SBundleId != "" {
		err = p.WriteString(3, st.SBundleId)
		if err != nil {
			return err
		}
	}
	if false || st.SProductId != "" {
		err = p.WriteString(4, st.SProductId)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseDate != 0 {
		err = p.WriteUint32(5, st.IPurchaseDate)
		if err != nil {
			return err
		}
	}
	if false || st.IQuantity != 0 {
		err = p.WriteUint32(6, st.IQuantity)
		if err != nil {
			return err
		}
	}
	if false || st.STransactionId != "" {
		err = p.WriteString(7, st.STransactionId)
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
	SPurchaseToken  string `json:"sPurchaseToken" form:"sPurchaseToken"`
	SProductId      string `json:"sProductId" form:"sProductId"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
}

func (st *GooglePurchase) resetDefault() {
}
func (st *GooglePurchase) Copy() *GooglePurchase {
	ret := NewGooglePurchase()
	ret.SPurchaseToken = st.SPurchaseToken
	ret.SProductId = st.SProductId
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
	util.Tab(buff, t+1, util.Fieldname("sPurchaseToken")+fmt.Sprintf("%v\n", st.SPurchaseToken))
	util.Tab(buff, t+1, util.Fieldname("sProductId")+fmt.Sprintf("%v\n", st.SProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *GooglePurchase) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SPurchaseToken, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProductId, 2, false)
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
	if false || st.SPurchaseToken != "" {
		err = p.WriteString(1, st.SPurchaseToken)
		if err != nil {
			return err
		}
	}
	if false || st.SProductId != "" {
		err = p.WriteString(2, st.SProductId)
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
	SKind                        string `json:"sKind" form:"sKind"`
	SPurchaseTimeMillis          string `json:"sPurchaseTimeMillis" form:"sPurchaseTimeMillis"`
	IPurchaseState               uint32 `json:"iPurchaseState" form:"iPurchaseState"`
	IConsumptionState            uint32 `json:"iConsumptionState" form:"iConsumptionState"`
	SDeveloperPayload            string `json:"sDeveloperPayload" form:"sDeveloperPayload"`
	SOrderId                     string `json:"sOrderId" form:"sOrderId"`
	IPurchaseType                uint32 `json:"iPurchaseType" form:"iPurchaseType"`
	IAcknowledgementState        uint32 `json:"iAcknowledgementState" form:"iAcknowledgementState"`
	SPurchaseToken               string `json:"sPurchaseToken" form:"sPurchaseToken"`
	SProductId                   string `json:"sProductId" form:"sProductId"`
	IQuantity                    uint32 `json:"iQuantity" form:"iQuantity"`
	SObfuscatedExternalAccountId string `json:"sObfuscatedExternalAccountId" form:"sObfuscatedExternalAccountId"`
	SObfuscatedExternalProfileId string `json:"sObfuscatedExternalProfileId" form:"sObfuscatedExternalProfileId"`
	SRegionCode                  string `json:"sRegionCode" form:"sRegionCode"`
}

func (st *GoogleReceipt) resetDefault() {
}
func (st *GoogleReceipt) Copy() *GoogleReceipt {
	ret := NewGoogleReceipt()
	ret.SKind = st.SKind
	ret.SPurchaseTimeMillis = st.SPurchaseTimeMillis
	ret.IPurchaseState = st.IPurchaseState
	ret.IConsumptionState = st.IConsumptionState
	ret.SDeveloperPayload = st.SDeveloperPayload
	ret.SOrderId = st.SOrderId
	ret.IPurchaseType = st.IPurchaseType
	ret.IAcknowledgementState = st.IAcknowledgementState
	ret.SPurchaseToken = st.SPurchaseToken
	ret.SProductId = st.SProductId
	ret.IQuantity = st.IQuantity
	ret.SObfuscatedExternalAccountId = st.SObfuscatedExternalAccountId
	ret.SObfuscatedExternalProfileId = st.SObfuscatedExternalProfileId
	ret.SRegionCode = st.SRegionCode
	return ret
}
func NewGoogleReceipt() *GoogleReceipt {
	ret := &GoogleReceipt{}
	ret.resetDefault()
	return ret
}
func (st *GoogleReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sKind")+fmt.Sprintf("%v\n", st.SKind))
	util.Tab(buff, t+1, util.Fieldname("sPurchaseTimeMillis")+fmt.Sprintf("%v\n", st.SPurchaseTimeMillis))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseState")+fmt.Sprintf("%v\n", st.IPurchaseState))
	util.Tab(buff, t+1, util.Fieldname("iConsumptionState")+fmt.Sprintf("%v\n", st.IConsumptionState))
	util.Tab(buff, t+1, util.Fieldname("sDeveloperPayload")+fmt.Sprintf("%v\n", st.SDeveloperPayload))
	util.Tab(buff, t+1, util.Fieldname("sOrderId")+fmt.Sprintf("%v\n", st.SOrderId))
	util.Tab(buff, t+1, util.Fieldname("iPurchaseType")+fmt.Sprintf("%v\n", st.IPurchaseType))
	util.Tab(buff, t+1, util.Fieldname("iAcknowledgementState")+fmt.Sprintf("%v\n", st.IAcknowledgementState))
	util.Tab(buff, t+1, util.Fieldname("sPurchaseToken")+fmt.Sprintf("%v\n", st.SPurchaseToken))
	util.Tab(buff, t+1, util.Fieldname("sProductId")+fmt.Sprintf("%v\n", st.SProductId))
	util.Tab(buff, t+1, util.Fieldname("iQuantity")+fmt.Sprintf("%v\n", st.IQuantity))
	util.Tab(buff, t+1, util.Fieldname("sObfuscatedExternalAccountId")+fmt.Sprintf("%v\n", st.SObfuscatedExternalAccountId))
	util.Tab(buff, t+1, util.Fieldname("sObfuscatedExternalProfileId")+fmt.Sprintf("%v\n", st.SObfuscatedExternalProfileId))
	util.Tab(buff, t+1, util.Fieldname("sRegionCode")+fmt.Sprintf("%v\n", st.SRegionCode))
}
func (st *GoogleReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SKind, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPurchaseTimeMillis, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseState, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IConsumptionState, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SDeveloperPayload, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOrderId, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPurchaseType, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IAcknowledgementState, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPurchaseToken, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SProductId, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IQuantity, 10, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SObfuscatedExternalAccountId, 11, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SObfuscatedExternalProfileId, 12, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRegionCode, 13, false)
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
	if false || st.SKind != "" {
		err = p.WriteString(0, st.SKind)
		if err != nil {
			return err
		}
	}
	if false || st.SPurchaseTimeMillis != "" {
		err = p.WriteString(1, st.SPurchaseTimeMillis)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseState != 0 {
		err = p.WriteUint32(2, st.IPurchaseState)
		if err != nil {
			return err
		}
	}
	if false || st.IConsumptionState != 0 {
		err = p.WriteUint32(3, st.IConsumptionState)
		if err != nil {
			return err
		}
	}
	if false || st.SDeveloperPayload != "" {
		err = p.WriteString(4, st.SDeveloperPayload)
		if err != nil {
			return err
		}
	}
	if false || st.SOrderId != "" {
		err = p.WriteString(5, st.SOrderId)
		if err != nil {
			return err
		}
	}
	if false || st.IPurchaseType != 0 {
		err = p.WriteUint32(6, st.IPurchaseType)
		if err != nil {
			return err
		}
	}
	if false || st.IAcknowledgementState != 0 {
		err = p.WriteUint32(7, st.IAcknowledgementState)
		if err != nil {
			return err
		}
	}
	if false || st.SPurchaseToken != "" {
		err = p.WriteString(8, st.SPurchaseToken)
		if err != nil {
			return err
		}
	}
	if false || st.SProductId != "" {
		err = p.WriteString(9, st.SProductId)
		if err != nil {
			return err
		}
	}
	if false || st.IQuantity != 0 {
		err = p.WriteUint32(10, st.IQuantity)
		if err != nil {
			return err
		}
	}
	if false || st.SObfuscatedExternalAccountId != "" {
		err = p.WriteString(11, st.SObfuscatedExternalAccountId)
		if err != nil {
			return err
		}
	}
	if false || st.SObfuscatedExternalProfileId != "" {
		err = p.WriteString(12, st.SObfuscatedExternalProfileId)
		if err != nil {
			return err
		}
	}
	if false || st.SRegionCode != "" {
		err = p.WriteString(13, st.SRegionCode)
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
	SOrderId        string `json:"sOrderId" form:"sOrderId"`
	SUuid           string `json:"sUuid" form:"sUuid"`
	SAppCallbackExt string `json:"sAppCallbackExt" form:"sAppCallbackExt"`
	IPayAmount      uint32 `json:"iPayAmount" form:"iPayAmount"`
	ISandBox        uint32 `json:"iSandBox" form:"iSandBox"`
	IPayTime        uint32 `json:"iPayTime" form:"iPayTime"`
	ITime           uint32 `json:"iTime" form:"iTime"`
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
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	IConnId         uint32 `json:"iConnId" form:"iConnId"`
	SChannel        string `json:"sChannel" form:"sChannel"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
	ICreateTime     uint32 `json:"iCreateTime" form:"iCreateTime"`
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
	IResult      int32   `json:"iResult" form:"iResult"`
	SPayMessage  string  `json:"sPayMessage" form:"sPayMessage"`
	SAgentId     string  `json:"sAgentId" form:"sAgentId"`
	SJnetBillNo  string  `json:"sJnetBillNo" form:"sJnetBillNo"`
	SAgentBillId string  `json:"sAgentBillId" form:"sAgentBillId"`
	IPayType     uint32  `json:"iPayType" form:"iPayType"`
	FPayAmt      float32 `json:"fPayAmt" form:"fPayAmt"`
	SRemark      string  `json:"sRemark" form:"sRemark"`
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
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	IConnId         uint32 `json:"iConnId" form:"iConnId"`
	SChannel        string `json:"sChannel" form:"sChannel"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
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
	SMethod        string `json:"sMethod" form:"sMethod"`
	SVersion       string `json:"sVersion" form:"sVersion"`
	SCharset       string `json:"sCharset" form:"sCharset"`
	SSignType      string `json:"sSignType" form:"sSignType"`
	SReturnCode    string `json:"sReturnCode" form:"sReturnCode"`
	SReturnMsg     string `json:"sReturnMsg" form:"sReturnMsg"`
	SAppId         string `json:"sAppId" form:"sAppId"`
	SMchId         string `json:"sMchId" form:"sMchId"`
	SNonceStr      string `json:"sNonceStr" form:"sNonceStr"`
	SResultCode    string `json:"sResultCode" form:"sResultCode"`
	SErrCode       string `json:"sErrCode" form:"sErrCode"`
	SErrCodeDesc   string `json:"sErrCodeDesc" form:"sErrCodeDesc"`
	SOpenId        string `json:"sOpenId" form:"sOpenId"`
	SFeeType       string `json:"sFeeType" form:"sFeeType"`
	ITotalFee      uint32 `json:"iTotalFee" form:"iTotalFee"`
	ICouponFee     uint32 `json:"iCouponFee" form:"iCouponFee"`
	STransactionId string `json:"sTransactionId" form:"sTransactionId"`
	SOutTradeNo    string `json:"sOutTradeNo" form:"sOutTradeNo"`
	STimeEnd       string `json:"sTimeEnd" form:"sTimeEnd"`
	SBuyerLogonId  string `json:"sBuyerLogonId" form:"sBuyerLogonId"`
	SFundBillList  string `json:"sFundBillList" form:"sFundBillList"`
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

type GameHubPurchase struct {
	SPurchaseData   string `json:"sPurchaseData" form:"sPurchaseData"`
	SSignature      string `json:"sSignature" form:"sSignature"`
	IConnId         uint32 `json:"iConnId" form:"iConnId"`
	SChannel        string `json:"sChannel" form:"sChannel"`
	STraceProductId string `json:"sTraceProductId" form:"sTraceProductId"`
	STraceFlowId    string `json:"sTraceFlowId" form:"sTraceFlowId"`
}

func (st *GameHubPurchase) resetDefault() {
}
func (st *GameHubPurchase) Copy() *GameHubPurchase {
	ret := NewGameHubPurchase()
	ret.SPurchaseData = st.SPurchaseData
	ret.SSignature = st.SSignature
	ret.IConnId = st.IConnId
	ret.SChannel = st.SChannel
	ret.STraceProductId = st.STraceProductId
	ret.STraceFlowId = st.STraceFlowId
	return ret
}
func NewGameHubPurchase() *GameHubPurchase {
	ret := &GameHubPurchase{}
	ret.resetDefault()
	return ret
}
func (st *GameHubPurchase) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sPurchaseData")+fmt.Sprintf("%v\n", st.SPurchaseData))
	util.Tab(buff, t+1, util.Fieldname("sSignature")+fmt.Sprintf("%v\n", st.SSignature))
	util.Tab(buff, t+1, util.Fieldname("iConnId")+fmt.Sprintf("%v\n", st.IConnId))
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sTraceProductId")+fmt.Sprintf("%v\n", st.STraceProductId))
	util.Tab(buff, t+1, util.Fieldname("sTraceFlowId")+fmt.Sprintf("%v\n", st.STraceFlowId))
}
func (st *GameHubPurchase) ReadStruct(up *codec.UnPacker) error {
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
func (st *GameHubPurchase) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *GameHubPurchase) WriteStruct(p *codec.Packer) error {
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
func (st *GameHubPurchase) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type GameHubReceipt struct {
	SChannelId      string  `json:"sChannelId" form:"sChannelId"`
	SChannelUid     string  `json:"sChannelUid" form:"sChannelUid"`
	SOrderNo        string  `json:"sOrderNo" form:"sOrderNo"`
	SCpOrderNo      string  `json:"sCpOrderNo" form:"sCpOrderNo"`
	SZoneId         string  `json:"sZoneId" form:"sZoneId"`
	SRoleId         string  `json:"sRoleId" form:"sRoleId"`
	SGoodsId        string  `json:"sGoodsId" form:"sGoodsId"`
	FAmount         float32 `json:"fAmount" form:"fAmount"`
	IPayTime        uint32  `json:"iPayTime" form:"iPayTime"`
	SPassbackParams string  `json:"sPassbackParams" form:"sPassbackParams"`
	IIsTest         uint32  `json:"iIsTest" form:"iIsTest"`
}

func (st *GameHubReceipt) resetDefault() {
}
func (st *GameHubReceipt) Copy() *GameHubReceipt {
	ret := NewGameHubReceipt()
	ret.SChannelId = st.SChannelId
	ret.SChannelUid = st.SChannelUid
	ret.SOrderNo = st.SOrderNo
	ret.SCpOrderNo = st.SCpOrderNo
	ret.SZoneId = st.SZoneId
	ret.SRoleId = st.SRoleId
	ret.SGoodsId = st.SGoodsId
	ret.FAmount = st.FAmount
	ret.IPayTime = st.IPayTime
	ret.SPassbackParams = st.SPassbackParams
	ret.IIsTest = st.IIsTest
	return ret
}
func NewGameHubReceipt() *GameHubReceipt {
	ret := &GameHubReceipt{}
	ret.resetDefault()
	return ret
}
func (st *GameHubReceipt) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sChannelId")+fmt.Sprintf("%v\n", st.SChannelId))
	util.Tab(buff, t+1, util.Fieldname("sChannelUid")+fmt.Sprintf("%v\n", st.SChannelUid))
	util.Tab(buff, t+1, util.Fieldname("sOrderNo")+fmt.Sprintf("%v\n", st.SOrderNo))
	util.Tab(buff, t+1, util.Fieldname("sCpOrderNo")+fmt.Sprintf("%v\n", st.SCpOrderNo))
	util.Tab(buff, t+1, util.Fieldname("sZoneId")+fmt.Sprintf("%v\n", st.SZoneId))
	util.Tab(buff, t+1, util.Fieldname("sRoleId")+fmt.Sprintf("%v\n", st.SRoleId))
	util.Tab(buff, t+1, util.Fieldname("sGoodsId")+fmt.Sprintf("%v\n", st.SGoodsId))
	util.Tab(buff, t+1, util.Fieldname("fAmount")+fmt.Sprintf("%v\n", st.FAmount))
	util.Tab(buff, t+1, util.Fieldname("iPayTime")+fmt.Sprintf("%v\n", st.IPayTime))
	util.Tab(buff, t+1, util.Fieldname("sPassbackParams")+fmt.Sprintf("%v\n", st.SPassbackParams))
	util.Tab(buff, t+1, util.Fieldname("iIsTest")+fmt.Sprintf("%v\n", st.IIsTest))
}
func (st *GameHubReceipt) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SChannelId, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SChannelUid, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SOrderNo, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SCpOrderNo, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SZoneId, 4, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRoleId, 5, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SGoodsId, 6, false)
	if err != nil {
		return err
	}
	err = up.ReadFloat32(&st.FAmount, 7, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IPayTime, 8, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SPassbackParams, 9, false)
	if err != nil {
		return err
	}
	err = up.ReadUint32(&st.IIsTest, 10, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *GameHubReceipt) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *GameHubReceipt) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SChannelId != "" {
		err = p.WriteString(0, st.SChannelId)
		if err != nil {
			return err
		}
	}
	if false || st.SChannelUid != "" {
		err = p.WriteString(1, st.SChannelUid)
		if err != nil {
			return err
		}
	}
	if false || st.SOrderNo != "" {
		err = p.WriteString(2, st.SOrderNo)
		if err != nil {
			return err
		}
	}
	if false || st.SCpOrderNo != "" {
		err = p.WriteString(3, st.SCpOrderNo)
		if err != nil {
			return err
		}
	}
	if false || st.SZoneId != "" {
		err = p.WriteString(4, st.SZoneId)
		if err != nil {
			return err
		}
	}
	if false || st.SRoleId != "" {
		err = p.WriteString(5, st.SRoleId)
		if err != nil {
			return err
		}
	}
	if false || st.SGoodsId != "" {
		err = p.WriteString(6, st.SGoodsId)
		if err != nil {
			return err
		}
	}
	if false || st.FAmount != 0 {
		err = p.WriteFloat32(7, st.FAmount)
		if err != nil {
			return err
		}
	}
	if false || st.IPayTime != 0 {
		err = p.WriteUint32(8, st.IPayTime)
		if err != nil {
			return err
		}
	}
	if false || st.SPassbackParams != "" {
		err = p.WriteString(9, st.SPassbackParams)
		if err != nil {
			return err
		}
	}
	if false || st.IIsTest != 0 {
		err = p.WriteUint32(10, st.IIsTest)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *GameHubReceipt) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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
	IReceiptType      uint32          `json:"iReceiptType" form:"iReceiptType"`
	StStatus          IAPStatus       `json:"stStatus" form:"stStatus"`
	StAppleReceipt    AppleReceipt    `json:"stAppleReceipt" form:"stAppleReceipt"`
	StGoogleReceipt   GoogleReceipt   `json:"stGoogleReceipt" form:"stGoogleReceipt"`
	StFyReceipt       FyReceipt       `json:"stFyReceipt" form:"stFyReceipt"`
	StHeePayReceipt   HeePayReceipt   `json:"stHeePayReceipt" form:"stHeePayReceipt"`
	StHeePayH5Receipt HeePayH5Receipt `json:"stHeePayH5Receipt" form:"stHeePayH5Receipt"`
	StGameHubReceipt  GameHubReceipt  `json:"stGameHubReceipt" form:"stGameHubReceipt"`
}

func (st *IAPReceiptInAll) resetDefault() {
	st.StStatus.resetDefault()
	st.StAppleReceipt.resetDefault()
	st.StGoogleReceipt.resetDefault()
	st.StFyReceipt.resetDefault()
	st.StHeePayReceipt.resetDefault()
	st.StHeePayH5Receipt.resetDefault()
	st.StGameHubReceipt.resetDefault()
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
	ret.StGameHubReceipt = *(st.StGameHubReceipt.Copy())
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
	util.Tab(buff, t+1, util.Fieldname("stGameHubReceipt")+"{\n")
	st.StGameHubReceipt.Visit(buff, t+1+1)
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
	err = st.StGameHubReceipt.ReadStructFromTag(up, 7, false)
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
	err = st.StGameHubReceipt.WriteStructFromTag(p, 7, false)
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
func (s *IAPService) DeliverGameHubReceipt(iRoleId uint64, iZoneId uint32, stPurchase GameHubPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error) {
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
	err = s.proxy.Invoke("deliverGameHubReceipt", p.ToBytes(), &rsp)
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
	DeliverGameHubReceipt(ctx context.Context, iRoleId uint64, iZoneId uint32, stPurchase GameHubPurchase, iProxyRoleId uint64, iProxyZoneId uint32) (int32, error)
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
func _IAPServiceDeliverGameHubReceiptImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
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
	var p3 GameHubPurchase
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
	ret, err = impl.DeliverGameHubReceipt(ctx, p1, p2, p3, p4, p5)
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
	case "deliverGameHubReceipt":
		err = _IAPServiceDeliverGameHubReceiptImpl(ctx, serviceImpl, up, p)
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
