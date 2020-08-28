package vo

type ArticleRequest struct {
	CategoryId uint   `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required,max=40"`
	HeadImg    string `json:"head_img"` // 存储文章头图地址
	Content    string `json:"content" binding:"required"`
}
