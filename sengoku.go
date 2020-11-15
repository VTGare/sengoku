package sengoku

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

type DB int

var (
	client      = &http.Client{}
	defaultPath = "/search.php"
)

const (
	HMagazines        DB = 0
	HGameCG           DB = 2
	DoujinshiDB       DB = 3
	Pixiv             DB = 5
	PixivHistorical   DB = 6
	NicoNicoSeiga     DB = 8
	Danbooru          DB = 9
	Drawr             DB = 10
	Nijie             DB = 11
	Yandere           DB = 12
	Openingsmoe       DB = 13
	Shutterstock      DB = 15
	FAKKU             DB = 16
	HMisc             DB = 18
	TwoDMarket        DB = 19
	MediBang          DB = 20
	Anime             DB = 21
	HAnime            DB = 22
	Movies            DB = 23
	Shows             DB = 24
	Gelbooru          DB = 25
	Konachan          DB = 26
	SankakuChannel    DB = 27
	AnimePicturesnet  DB = 28
	E621net           DB = 29
	IdolComplex       DB = 30
	BcynetIllust      DB = 31
	BcynetCosplay     DB = 32
	PortalGraphicsnet DB = 33
	DeviantArt        DB = 34
	Pawoonet          DB = 35
	Madokami          DB = 36
	MangaDex          DB = 37
	HMiscEHentai      DB = 38
	Artstation        DB = 39
	FurAffinity       DB = 40
	All               DB = 999
)

//Sengoku is a SauceNAO API wrapper.
type Sengoku struct {
	DefaultConfig *Config
	client        *http.Client
	baseURL       string
}

//Config is a SauceNAO API call configuration structure.
//
//APIKEy is an API token used on requests.
//
//DB searches a specific database index without having to generate a mask.
//
//TestMode causes each index that has a match for the given image to output at most 1 result. Useful for testing some things.
//
//Results controls how many results are returned from SauceNAO.
//
//IncludeMask is an array of the only database indices you want returned from SauceNAO. A mask of [5] would only return results from Pixiv.
//
//ExcludeMask is an opposite of IncludeMask, a mask of [5] would return results from anywhere but Pixiv.
type Config struct {
	APIKey   string
	DB       DB
	TestMode bool
	Results  int
	//IncludeMask []int
	//ExcludeMask []int
}

//DefaultConfig returns default Config configuration.
func DefaultConfig() *Config {
	return &Config{DB: 999, TestMode: false, Results: 0}
}

//NewSengoku creates a new Sengoku application instance.
func NewSengoku(config ...Config) *Sengoku {
	sengoku := &Sengoku{DefaultConfig(), &http.Client{}, "https://saucenao.com"}
	if len(config) != 0 {
		sengoku.DefaultConfig = &config[0]
	}

	return sengoku
}

//Search performs a SauceNAO API call with a given image URL.
func (s *Sengoku) Search(imageURL string) ([]*Sauce, error) {
	uri, _ := url.Parse(s.baseURL + defaultPath)
	uri.RawQuery = s.DefaultConfig.query(uri, imageURL)

	resp, err := get(uri.String())
	if err != nil {
		logrus.Warnln("Search(): ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &responce{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	sauces := make([]*Sauce, len(res.Results))
	for ind, r := range res.Results {
		sauce := &Sauce{}
		if toSauce, ok := sites[r.Header.IndexID]; ok {
			sauce = toSauce(r)
			sauce.Pretty = true
		} else {
			sauce = &Sauce{Raw: r}
		}
		sauces[ind] = sauce
	}

	return sauces, nil
}

//SearchWithConfig performs a SauceNAO API call with a given URL and custom configuration. If API key is empty in given configuration it will try to use default API key.
func (s *Sengoku) SearchWithConfig(imageURL string, config *Config) ([]*Sauce, error) {
	if config.APIKey == "" {
		config.APIKey = s.DefaultConfig.APIKey
	}
	uri, _ := url.Parse(s.baseURL + defaultPath)
	uri.RawQuery = config.query(uri, imageURL)

	resp, err := get(uri.String())
	if err != nil {
		logrus.Warnln("SearchWithConfig(): ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &responce{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	sauces := make([]*Sauce, len(res.Results))
	for ind, r := range res.Results {
		sauce := &Sauce{}
		if toSauce, ok := sites[r.Header.IndexID]; ok {
			sauce = toSauce(r)
			sauce.Pretty = true
		} else {
			sauce = &Sauce{Raw: r}
		}
		sauces[ind] = sauce
	}

	return sauces, nil
}

func get(uri string) (*http.Response, error) {
	logrus.Infof("Making an SauceNAO API request. URL: %v", uri)
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		switch resp.StatusCode {
		case 403:
			return nil, ErrInvalidAPIKey
		case 413:
			return nil, ErrFileTooLarge
		case 429:
			return nil, ErrRateLimitReached
		default:
			return nil, errUnknown(resp.StatusCode)
		}
	}

	return resp, nil
}

func (c *Config) query(uri *url.URL, imageURL string) string {
	q := uri.Query()
	if c.APIKey != "" {
		q.Set("api_key", c.APIKey)
	}

	if c.TestMode == true {
		q.Set("test_mode", "1")
	}

	if c.Results != 0 {
		q.Set("num_res", strconv.Itoa(c.Results))
	}

	if c.DB == 0 {
		c.DB = 999
	}

	q.Set("db", strconv.Itoa(int(c.DB)))
	q.Set("output_type", "2")
	q.Set("url", imageURL)

	return q.Encode()
}
