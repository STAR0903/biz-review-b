package service

import (
	"context"
	v1 "review-b/api/business/v1"
	"review-b/internal/biz"
)

type BusinessService struct {
	v1.UnimplementedBusinessServer

	uc *biz.BusinessUsecase
}

// NewBusinessService 创建商家服务
func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

// ReplyReview 商家回复评价
func (s *BusinessService) ReplyReview(ctx context.Context, req *v1.ReplyReviewRequest) (*v1.ReplyReviewReply, error) {
	// 参数转换
	param := &biz.ReplyReviewParam{
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	}
	// 调用biz层
	replyID, err := s.uc.ReplyReview(ctx, param)
	if err != nil {
		return nil, err
	}
	// 返回响应
	return &v1.ReplyReviewReply{
		ReplyID: replyID,
	}, nil
}

// AppealReview 商家申诉用户评价
func (s *BusinessService) AppealReview(ctx context.Context, req *v1.AppealReviewRequest) (*v1.AppealReviewReply, error) {
	// 参数转换
	param := &biz.AppealReviewParam{
		ReviewID:  req.ReviewID,
		StoreID:   req.StoreID,
		Reason:    req.Reason,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	}
	// 调用biz层
	appealID, err := s.uc.AppealReview(ctx, param)
	if err != nil {
		return nil, err
	}
	// 返回响应
	return &v1.AppealReviewReply{
		AppealID: appealID,
	}, nil
}
