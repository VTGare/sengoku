package sengoku

type responce struct {
	Header  *header   `json:"header"`
	Results []*Result `json:"results"`
}

type header struct {
	AccountType       string  `json:"account_type,omitempty"`
	LongLimit         string  `json:"long_limit,omitempty"`
	LongRemaining     int     `json:"long_remaining,omitempty"`
	Message           string  `json:"message,omitempty"`
	MinimumSimilarity float64 `json:"minimum_similarity,omitempty"`
	QueryImage        string  `json:"query_image,omitempty"`
	ResultsRequested  string  `json:"results_requested,omitempty"`
	ResultsReturned   int     `json:"results_returned,omitempty"`
	SearchDepth       string  `json:"search_depth,omitempty"`
	ShortLimit        string  `json:"short_limit,omitempty"`
	ShortRemaining    int     `json:"short_remaining,omitempty"`
}

type Result struct {
	Data   *ResultData   `json:"data,omitempty"`
	Header *ResultHeader `json:"header,omitempty"`
}

type ResultData struct {
	ExternalURLs []string    `json:"ext_urls,omitempty"`
	Title        string      `json:"title,omitempty"`
	DaID         interface{} `json:"da_id,omitempty"`
	Author       string      `json:"author,omitempty"`
	AuthorName   string      `json:"author_name,omitempty"`
	AuthorURL    string      `json:"author_url,omitempty"`
	AnidbAid     interface{} `json:"anidb_aid,omitempty"`
	Artist       string      `json:"artist,omitempty"`
	BcyID        interface{} `json:"bcy_id,omitempty"`
	BcyType      string      `json:"bcy_type,omitempty"`
	DanbooruID   interface{} `json:"danbooru_id,omitempty"`
	DdbID        interface{} `json:"ddb_id,omitempty"`
	DrawrID      interface{} `json:"drawr_id,omitempty"`
	Creator      interface{} `json:"creator,omitempty"`
	EngName      string      `json:"eng_name,omitempty"`
	E621ID       interface{} `json:"e621_id,omitempty"`
	File         string      `json:"file,omitempty"`
	GelbooruID   interface{} `json:"gelbooru_id,omitempty"`
	IdolID       interface{} `json:"idol_id,omitempty"`
	ImdbID       interface{} `json:"imdb_id,omitempty"`
	JpName       string      `json:"jp_name,omitempty"`
	KonachanID   interface{} `json:"konachan_id,omitempty"`
	Material     string      `json:"material,omitempty"`
	MemberLinkID interface{} `json:"member_link_id,omitempty"`
	MuID         interface{} `json:"mu_id,omitempty"`
	NijieID      interface{} `json:"nijie_id,omitempty"`
	Part         string      `json:"part,omitempty"`
	PawooID      interface{} `json:"pawoo_id,omitempty"`
	PgID         interface{} `json:"pg_id,omitempty"`
	PixivID      interface{} `json:"pixiv_id,omitempty"`
	SankakuID    interface{} `json:"sankaku_id,omitempty"`
	SeigaID      interface{} `json:"seiga_id,omitempty"`
	Source       string      `json:"source,omitempty"`
	URL          string      `json:"url,omitempty"`
	UserAcct     string      `json:"user_acct,omitempty"`
	YandereID    interface{} `json:"yandere_id,omitempty"`
	MemberID     interface{} `json:"member_id,omitempty"`
	MemberName   string      `json:"member_name,omitempty"`
}

type ResultHeader struct {
	IndexID    int    `json:"index_id,omitempty"`
	IndexName  string `json:"index_name,omitempty"`
	Similarity string `json:"similarity,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
}
