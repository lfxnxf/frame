package service

import (
	"context"
	api "github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/api/game_center_base"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/manager"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/conf"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/dao"
)

// Service represents several business logic(s)
type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao

	// manager: other client(s), other middleware(s)
	mgr *manager.Manager
}

// New new a service and return.
func New(c *conf.Config) (s *Service) {
	return &Service{
		c:   c,
		dao: dao.New(c),
		mgr: manager.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
	if s.mgr != nil {
		s.mgr.Close()
	}
}


func (s *Service) GetInvitationSenderStatus(ctx context.Context, request *api.UidReq) (*api.InvitationIdResp, error) {
	return nil, nil
}

func (s *Service) SetInvitationInfo(ctx context.Context, request *api.SetInvitationInfoReq) (*api.InvitationIdResp, error) {
	return nil, nil
}

func (s *Service) GetInvitationInfo(ctx context.Context, request *api.InvitationIdReq) (*api.GetInvitationInfoResp, error) {
	return nil, nil
}

func (s *Service) InvitationAccept(ctx context.Context, request *api.InvitationIdReq) (*api.InvitationEmptyResp, error) {
	return nil, nil
}

func (s *Service) InvitationCancel(ctx context.Context, request *api.InvitationIdReq) (*api.InvitationEmptyResp, error) {
	return nil, nil
}

func (s *Service) InvitationRefused(ctx context.Context, request *api.InvitationIdReq) (*api.InvitationEmptyResp, error) {
	return nil, nil
}

func (s *Service) InvitationDontDisturb(ctx context.Context, request *api.UidReq) (*api.InvitationEmptyResp, error) {
	return nil, nil
}

func (s *Service) GetInvitationDontDisturbStatus(ctx context.Context, request *api.UidReq) (*api.DontDisturbStatus, error) {
	return nil, nil
}

func (s *Service) SetContinuousRefused(ctx context.Context, request *api.SetContinuousRefusedReq) (*api.ContinuousRefusedResp, error) {
	return nil, nil
}

func (s *Service) GetContinuousRefused(ctx context.Context, request *api.UidReq) (*api.ContinuousRefusedResp, error) {
	return nil, nil
}

func (s *Service) SetInvitationRoomId(ctx context.Context, request *api.SetInvitationRoomIdReq) (*api.InvitationEmptyResp, error) {
	return nil, nil
}

func (s *Service) InvitationFreqSingle(ctx context.Context, request *api.InvitationFreqSingleReq) (*api.InvitationFreqSingleResp, error) {
	return nil, nil
}

func (s *Service) UpdateUserGameRecord(ctx context.Context, request *api.UpdateUserGameRecordReq) (*api.UpdateUserGameRecordResp, error) {
	return nil, nil
}

func (s *Service) BatchUpdateUserGameRecord(ctx context.Context, request *api.BatchUpdateUserGameRecordReq) (*api.BatchUpdateUserGameRecordResp, error) {
	return nil, nil
}

func (s *Service) GetUserGameRecord(ctx context.Context, request *api.GetUserGameRecordReq) (*api.GetUserGameRecordResp, error) {
	return nil, nil
}

func (s *Service) MGetUserGameRecord(ctx context.Context, request *api.MGetUserGameRecordReq) (*api.MGetUserGameRecordResp, error) {
	return nil, nil
}


