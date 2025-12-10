package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo 商家仓库接口
type BusinessRepo interface {
	// ReplyReview 商家回复评价
	ReplyReview(ctx context.Context, param *ReplyReviewParam) (int64, error)
	// AppealReview 商家申诉评价
	AppealReview(ctx context.Context, param *AppealReviewParam) (int64, error)
}

// BusinessUsecase 商家用例
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewBusinessUsecase 创建商家用例
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

// ReplyReview 商家回复评价
func (uc *BusinessUsecase) ReplyReview(ctx context.Context, param *ReplyReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Debugf("ReplyReview param: %v", param)
	return uc.repo.ReplyReview(ctx, param)
}

// AppealReview 商家申诉评价
func (uc *BusinessUsecase) AppealReview(ctx context.Context, param *AppealReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Debugf("AppealReview param: %v", param)
	return uc.repo.AppealReview(ctx, param)
}
