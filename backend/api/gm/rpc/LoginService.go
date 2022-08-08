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

type ChannelAddr struct {
	SChannel    string `json:"sChannel" form:"sChannel"`
	SAddress    string `json:"sAddress" form:"sAddress"`
	SRes        string `json:"sRes" form:"sRes"`
	SShopVer    string `json:"sShopVer" form:"sShopVer"`
	SFastFollow string `json:"sFastFollow" form:"sFastFollow"`
}

func (st *ChannelAddr) resetDefault() {
}
func (st *ChannelAddr) Copy() *ChannelAddr {
	ret := NewChannelAddr()
	ret.SChannel = st.SChannel
	ret.SAddress = st.SAddress
	ret.SRes = st.SRes
	ret.SShopVer = st.SShopVer
	ret.SFastFollow = st.SFastFollow
	return ret
}
func NewChannelAddr() *ChannelAddr {
	ret := &ChannelAddr{}
	ret.resetDefault()
	return ret
}
func (st *ChannelAddr) Visit(buff *bytes.Buffer, t int) {
	util.Tab(buff, t+1, util.Fieldname("sChannel")+fmt.Sprintf("%v\n", st.SChannel))
	util.Tab(buff, t+1, util.Fieldname("sAddress")+fmt.Sprintf("%v\n", st.SAddress))
	util.Tab(buff, t+1, util.Fieldname("sRes")+fmt.Sprintf("%v\n", st.SRes))
	util.Tab(buff, t+1, util.Fieldname("sShopVer")+fmt.Sprintf("%v\n", st.SShopVer))
	util.Tab(buff, t+1, util.Fieldname("sFastFollow")+fmt.Sprintf("%v\n", st.SFastFollow))
}
func (st *ChannelAddr) ReadStruct(up *codec.UnPacker) error {
	var err error
	var length uint32
	var has bool
	var ty uint32
	st.resetDefault()
	err = up.ReadString(&st.SChannel, 0, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SAddress, 1, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SRes, 2, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SShopVer, 3, false)
	if err != nil {
		return err
	}
	err = up.ReadString(&st.SFastFollow, 4, false)
	if err != nil {
		return err
	}

	_ = length
	_ = has
	_ = ty

	return err
}
func (st *ChannelAddr) ReadStructFromTag(up *codec.UnPacker, tag uint32, require bool) error {
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
func (st *ChannelAddr) WriteStruct(p *codec.Packer) error {
	var err error
	var length uint32
	if false || st.SChannel != "" {
		err = p.WriteString(0, st.SChannel)
		if err != nil {
			return err
		}
	}
	if false || st.SAddress != "" {
		err = p.WriteString(1, st.SAddress)
		if err != nil {
			return err
		}
	}
	if false || st.SRes != "" {
		err = p.WriteString(2, st.SRes)
		if err != nil {
			return err
		}
	}
	if false || st.SShopVer != "" {
		err = p.WriteString(3, st.SShopVer)
		if err != nil {
			return err
		}
	}
	if false || st.SFastFollow != "" {
		err = p.WriteString(4, st.SFastFollow)
		if err != nil {
			return err
		}
	}

	_ = length
	return err
}
func (st *ChannelAddr) WriteStructFromTag(p *codec.Packer, tag uint32, require bool) error {
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

type LoginService struct {
	proxy model.ServicePrxImpl
}

func (s *LoginService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *LoginService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *LoginService) AddNewChannel(sChannel string, sAddress string, sRes string, sShopVer string, sFastFollow string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sChannel != "" {
		err = p.WriteString(1, sChannel)
		if err != nil {
			return ret, err
		}
	}
	if true || sAddress != "" {
		err = p.WriteString(2, sAddress)
		if err != nil {
			return ret, err
		}
	}
	if true || sRes != "" {
		err = p.WriteString(3, sRes)
		if err != nil {
			return ret, err
		}
	}
	if true || sShopVer != "" {
		err = p.WriteString(4, sShopVer)
		if err != nil {
			return ret, err
		}
	}
	if true || sFastFollow != "" {
		err = p.WriteString(5, sFastFollow)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("addNewChannel", p.ToBytes(), &rsp)
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
func (s *LoginService) ModifyChannel(sChannel string, sAddress string, sRes string, sShopVer string, sFastFollow string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sChannel != "" {
		err = p.WriteString(1, sChannel)
		if err != nil {
			return ret, err
		}
	}
	if true || sAddress != "" {
		err = p.WriteString(2, sAddress)
		if err != nil {
			return ret, err
		}
	}
	if true || sRes != "" {
		err = p.WriteString(3, sRes)
		if err != nil {
			return ret, err
		}
	}
	if true || sShopVer != "" {
		err = p.WriteString(4, sShopVer)
		if err != nil {
			return ret, err
		}
	}
	if true || sFastFollow != "" {
		err = p.WriteString(5, sFastFollow)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("modifyChannel", p.ToBytes(), &rsp)
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
func (s *LoginService) DelChannel(sChannel string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sChannel != "" {
		err = p.WriteString(1, sChannel)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("delChannel", p.ToBytes(), &rsp)
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
func (s *LoginService) GetAllChannel(vChannelAddr *[]ChannelAddr) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAllChannel", p.ToBytes(), &rsp)
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
		(*vChannelAddr) = make([]ChannelAddr, length, length)
		for i := uint32(0); i < length; i++ {
			err = (*vChannelAddr)[i].ReadStructFromTag(up, 0, true)
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
func (s *LoginService) GetAddress(sChannel string, sAddress *string, sRes *string, sShopVer *string, sFastFollow *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sChannel != "" {
		err = p.WriteString(1, sChannel)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("getAddress", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sAddress), 2, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sRes), 3, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sShopVer), 4, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sFastFollow), 5, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _LoginServiceImpl interface {
	AddNewChannel(ctx context.Context, sChannel string, sAddress string, sRes string, sShopVer string, sFastFollow string) (int32, error)
	ModifyChannel(ctx context.Context, sChannel string, sAddress string, sRes string, sShopVer string, sFastFollow string) (int32, error)
	DelChannel(ctx context.Context, sChannel string) (int32, error)
	GetAllChannel(ctx context.Context, vChannelAddr *[]ChannelAddr) (int32, error)
	GetAddress(ctx context.Context, sChannel string, sAddress *string, sRes *string, sShopVer *string, sFastFollow *string) (int32, error)
}

func _LoginServiceAddNewChannelImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_LoginServiceImpl)
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
	var p5 string
	err = up.ReadString(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.AddNewChannel(ctx, p1, p2, p3, p4, p5)
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
func _LoginServiceModifyChannelImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_LoginServiceImpl)
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
	var p5 string
	err = up.ReadString(&p5, 5, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.ModifyChannel(ctx, p1, p2, p3, p4, p5)
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
func _LoginServiceDelChannelImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_LoginServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var ret int32
	ret, err = impl.DelChannel(ctx, p1)
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
func _LoginServiceGetAllChannelImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_LoginServiceImpl)
	var p1 []ChannelAddr
	var ret int32
	ret, err = impl.GetAllChannel(ctx, &p1)
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
func _LoginServiceGetAddressImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_LoginServiceImpl)
	var p1 string
	err = up.ReadString(&p1, 1, true)
	if err != nil {
		return err
	}
	var p2 string
	var p3 string
	var p4 string
	var p5 string
	var ret int32
	ret, err = impl.GetAddress(ctx, p1, &p2, &p3, &p4, &p5)
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
	if true || p3 != "" {
		err = p.WriteString(3, p3)
		if err != nil {
			return err
		}
	}
	if true || p4 != "" {
		err = p.WriteString(4, p4)
		if err != nil {
			return err
		}
	}
	if true || p5 != "" {
		err = p.WriteString(5, p5)
		if err != nil {
			return err
		}
	}
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *LoginService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "addNewChannel":
		err = _LoginServiceAddNewChannelImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "modifyChannel":
		err = _LoginServiceModifyChannelImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "delChannel":
		err = _LoginServiceDelChannelImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAllChannel":
		err = _LoginServiceGetAllChannelImpl(ctx, serviceImpl, up, p)
		if err != nil {
			break
		}
		texret = protocol.SDPSERVERSUCCESS
	case "getAddress":
		err = _LoginServiceGetAddressImpl(ctx, serviceImpl, up, p)
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
