// 此文件为sdp2go工具自动生成,请不要手动编辑

package rpc

import (
	"context"
	"github.com/yellia1989/tex-go/sdp/protocol"
	"github.com/yellia1989/tex-go/service/model"
	"github.com/yellia1989/tex-go/tools/log"
	"github.com/yellia1989/tex-go/tools/net"
	"github.com/yellia1989/tex-go/tools/sdp/codec"
	"time"
)

type MapService struct {
	proxy model.ServicePrxImpl
}

func (s *MapService) SetPrxImpl(impl model.ServicePrxImpl) {
	s.proxy = impl
}
func (s *MapService) SetTimeout(timeout time.Duration) {
	s.proxy.SetTimeout(timeout)
}
func (s *MapService) DoGmCmd(sOperatorName string, sReqCmd string, sResp *string) (int32, error) {
	p := codec.NewPacker()
	var ret int32
	var err error
	var has bool
	var ty uint32
	var length uint32
	if true || sOperatorName != "" {
		err = p.WriteString(1, sOperatorName)
		if err != nil {
			return ret, err
		}
	}
	if true || sReqCmd != "" {
		err = p.WriteString(2, sReqCmd)
		if err != nil {
			return ret, err
		}
	}
	var rsp *protocol.ResponsePacket
	err = s.proxy.Invoke("doGmCmd", p.ToBytes(), &rsp)
	if err != nil {
		return ret, err
	}
	up := codec.NewUnPacker([]byte(rsp.SRspPayload))
	err = up.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}
	err = up.ReadString(&(*sResp), 3, true)
	if err != nil {
		return ret, err
	}
	_ = has
	_ = ty
	_ = length
	return ret, nil
}

type _MapServiceImpl interface {
	DoGmCmd(ctx context.Context, sOperatorName string, sReqCmd string, sResp *string) (int32, error)
}

func _MapServiceDoGmCmdImpl(ctx context.Context, serviceImpl interface{}, up *codec.UnPacker, p *codec.Packer) error {
	var err error
	var length uint32
	var ty uint32
	var has bool
	impl := serviceImpl.(_MapServiceImpl)
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
	var ret int32
	ret, err = impl.DoGmCmd(ctx, p1, p2, &p3)
	if err != nil {
		return err
	}
	if true || ret != 0 {
		err = p.WriteInt32(0, ret)
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
	_ = length
	_ = ty
	_ = has
	return nil
}

func (s *MapService) Dispatch(ctx context.Context, serviceImpl interface{}, req *protocol.RequestPacket) {
	current := net.ContextGetCurrent(ctx)

	log.FDebugf("handle tex request, peer: %s:%d, obj: %s, func: %s, reqid: %d", current.IP, current.Port, req.SServiceName, req.SFuncName, req.IRequestId)

	texret := protocol.SDPSERVERUNKNOWNERR
	up := codec.NewUnPacker([]byte(req.SReqPayload))
	p := codec.NewPacker()

	var err error
	switch req.SFuncName {
	case "doGmCmd":
		err = _MapServiceDoGmCmdImpl(ctx, serviceImpl, up, p)
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
