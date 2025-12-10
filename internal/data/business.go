package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-b/api/review/v1"
	"review-b/internal/biz"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// ReplyReview 商家回复评价
func (r *businessRepo) ReplyReview(ctx context.Context, param *biz.ReplyReviewParam) (int64, error) {
	reply, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		return 0, err
	}
	return reply.ReplyID, nil
}

// AppealReview 商家申诉评价
func (r *businessRepo) AppealReview(ctx context.Context, param *biz.AppealReviewParam) (int64, error) {
	appeal, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Reason:    param.Reason,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		return 0, err
	}
	return appeal.AppealID, nil
}
