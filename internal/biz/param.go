package biz

// ReplyReviewParam 商家回复用户评价参数
type ReplyReviewParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// AppealReviewParam 商家申诉用户评价参数
type AppealReviewParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
}
