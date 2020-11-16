package sengoku

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
)

var (
	nhentaiRegex = regexp.MustCompile(`\/nhentai(\d+)`)
	sites        = map[int]resultFunc{
		5: func(r *Result) (*Sauce, error) {
			urls := &SauceURLs{}
			if r.Data.PixivID != 0 {
				urls.Source = fmt.Sprintf("https://pixiv.net/en/artworks/%v", r.Data.PixivID)
			} else if len(r.Data.ExternalURLs) != 0 {
				urls.Source = r.Data.ExternalURLs[0]
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}

			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.MemberName,
					URL:  fmt.Sprintf("https://www.pixiv.net/en/users/%v", r.Data.MemberID),
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
		6: func(r *Result) (*Sauce, error) {
			urls := &SauceURLs{}
			if r.Data.PixivID != 0 {
				urls.Source = fmt.Sprintf("https://pixiv.net/en/artworks/%v", r.Data.PixivID)
			} else if len(r.Data.ExternalURLs) != 0 {
				urls.Source = r.Data.ExternalURLs[0]
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.MemberName,
					URL:  fmt.Sprintf("https://www.pixiv.net/en/users/%v", r.Data.MemberID),
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
		8: func(r *Result) (*Sauce, error) {
			var urls = &SauceURLs{}
			if len(r.Data.ExternalURLs) != 0 {
				urls.Source = r.Data.ExternalURLs[0]
			} else {
				urls.Source = fmt.Sprintf("https://seiga.nicovideo.jp/seiga/im%v", r.Data.SeigaID)
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.MemberName,
					URL:  fmt.Sprintf("https://seiga.nicovideo.jp/user/illust/%v", r.Data.MemberID),
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
		9: func(r *Result) (*Sauce, error) {
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      "Material: " + r.Data.Material,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.Creator.(string),
				},
				URLs: &SauceURLs{
					Source:       r.Data.Source,
					ExternalURLs: r.Data.ExternalURLs,
				},
				Raw: r,
			}
			return sauce, nil
		},
		12: func(r *Result) (*Sauce, error) {
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      "Material: " + r.Data.Material,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.Creator.(string),
				},
				URLs: &SauceURLs{
					Source:       r.Data.Source,
					ExternalURLs: r.Data.ExternalURLs,
				},
				Raw: r,
			}
			return sauce, nil
		},
		18: func(r *Result) (*Sauce, error) {
			creator := ""
			if creators, ok := r.Data.Creator.([]interface{}); ok {
				if len(creators) != 0 {
					creator = creators[0].(string)
				}
			}

			title := ""
			switch {
			case r.Data.EngName != "":
				title = r.Data.EngName
			case r.Data.Source != "":
				title = r.Data.Source
			case r.Data.JpName != "":
				title = r.Data.JpName
			}

			source := ""
			if matches := nhentaiRegex.FindAllString(r.Header.Thumbnail, -1); matches != nil {
				source = fmt.Sprintf("https://nhentai.net/g/%v", matches[1])
			} else {
				source = fmt.Sprintf("https://nhentai.net/search/?q=%v", url.QueryEscape(r.Data.Source))
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: creator,
				},
				URLs: &SauceURLs{
					Source: source,
				},
				Raw: r,
			}
			return sauce, nil
		},
		25: func(r *Result) (*Sauce, error) {
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      "Material: " + r.Data.Material,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.Creator.(string),
				},
				URLs: &SauceURLs{
					Source:       r.Data.Source,
					ExternalURLs: r.Data.ExternalURLs,
				},
				Raw: r,
			}
			return sauce, nil
		},
		31: func(r *Result) (*Sauce, error) {
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.MemberName,
					URL:  fmt.Sprintf("https://bcy.net/u/%v", r.Data.MemberID),
				},
				URLs: &SauceURLs{
					Source:       fmt.Sprintf("https://bcy.net/%v/detail/%v/%v", r.Data.BcyType, r.Data.MemberLinkID, r.Data.BcyID),
					ExternalURLs: r.Data.ExternalURLs,
				},
				Raw: r,
			}
			return sauce, nil
		},
		34: func(r *Result) (*Sauce, error) {
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.AuthorName,
					URL:  r.Data.AuthorURL,
				},
				URLs: &SauceURLs{
					Source: fmt.Sprintf("https://deviantart.com/view/%v", r.Data.DaID),
				},
				Raw: r,
			}
			return sauce, nil
		},
		37: func(r *Result) (*Sauce, error) {
			urls := &SauceURLs{}
			if l := len(r.Data.ExternalURLs); l != 0 {
				urls.Source = r.Data.ExternalURLs[0]
				if l > 1 {
					urls.ExternalURLs = r.Data.ExternalURLs[1:]
				}
			}
			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Source + r.Data.Part,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.Author,
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
		38: func(r *Result) (*Sauce, error) {
			creator := ""
			if creators, ok := r.Data.Creator.([]interface{}); ok {
				if len(creators) != 0 {
					creator = creators[0].(string)
				}
			}

			title := ""
			switch {
			case r.Data.EngName != "":
				title = r.Data.EngName
			case r.Data.Source != "":
				title = r.Data.Source
			case r.Data.JpName != "":
				title = r.Data.JpName
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: creator,
				},
				URLs: &SauceURLs{
					Source: fmt.Sprintf("https://e-hentai.org/?f_search=%v", url.QueryEscape(r.Data.Source)),
				},
				Raw: r,
			}
			return sauce, nil
		},
		39: func(r *Result) (*Sauce, error) {
			urls := &SauceURLs{}
			if l := len(r.Data.ExternalURLs); l != 0 {
				urls.Source = r.Data.ExternalURLs[0]
				if l > 1 {
					urls.ExternalURLs = r.Data.ExternalURLs[1:]
				}
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.AuthorName,
					URL:  r.Data.AuthorURL,
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
		40: func(r *Result) (*Sauce, error) {
			urls := &SauceURLs{}
			if l := len(r.Data.ExternalURLs); l != 0 {
				urls.Source = r.Data.ExternalURLs[0]
				if l > 1 {
					urls.ExternalURLs = r.Data.ExternalURLs[1:]
				}
			}

			sim, err := strconv.ParseFloat(r.Header.Similarity, 64)
			if err != nil {
				return nil, err
			}
			sauce := &Sauce{
				Title:      r.Data.Title,
				Thumbnail:  r.Header.Thumbnail,
				Similarity: sim,
				Author: &SauceAuthor{
					Name: r.Data.AuthorName,
					URL:  r.Data.AuthorURL,
				},
				URLs: urls,
				Raw:  r,
			}
			return sauce, nil
		},
	}
)

type resultFunc func(*Result) (*Sauce, error)

//Sauce is an abstaction over SauceNAO response. If Pretty is false, then all fields except Raw are default.
type Sauce struct {
	Title      string
	Thumbnail  string
	Similarity float64
	Author     *SauceAuthor
	URLs       *SauceURLs
	Pretty     bool
	Raw        *Result
}

//SauceAuthor contains information about art's author
type SauceAuthor struct {
	Name string
	URL  string
}

//SauceURLs contains all possible URLs from a response.
type SauceURLs struct {
	Source       string
	ExternalURLs []string
}
