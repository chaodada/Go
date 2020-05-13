package defs

// 请求结构体
type ShortenReq struct {
	Url                 string `json:"url" validate:"required"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"required"`
}

// 响应结构体
type ShortenResp struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}
