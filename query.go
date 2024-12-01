package meander

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var (
	GOOGLE_API_KEY         = os.Getenv("GOOGLE_API_KEY")
	URL_GOOGLE_API_JOURNEY = os.Getenv("URL_GOOGLE_API_JOURNEY")
	URL_GOOGLE_API_PHOTO   = os.Getenv("URL_GOOGLE_API_PHOTO")
)

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

func (p *Place) Public() any {
	return map[string]any{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (q *Query) find(types string) (*googleResponse, error) {
	//GoogleのAPIサーバーに対してリクエストを行う
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("types", types)
	vals.Set("key", GOOGLE_API_KEY)
	if len(q.CostRangeStr) > 0 {
		r := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}

	res, err := http.Get(URL_GOOGLE_API_JOURNEY + "?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var gr googleResponse
	if err := json.NewDecoder(res.Body).Decode(&gr); err != nil {
		return nil, err
	}

	return &gr, nil
}

func (q *Query) Run() any {
	randgen := rand.New(rand.NewSource(time.Now().UnixNano()))
	var w sync.WaitGroup
	var l sync.Mutex
	places := make([]any, len(q.Journey))

	for i, r := range q.Journey {
		w.Add(1)
		go func(types string, i int) {
			defer w.Done()

			res, err := q.find(types)
			if err != nil {
				log.Println("施設の検索に失敗しました. error:", err)
				return
			}

			if len(res.Results) == 0 {
				log.Println("施設が見つかりませんでした. types:", types)
				return
			}

			//別途、写真用のGoogleAPIサーバにリクエストをする
			for _, result := range res.Results {
				for _, photo := range result.Photos {
					photo.URL = URL_GOOGLE_API_PHOTO +
						"maxwidth=1000&photoreference" + photo.PhotoRef +
						"&key" + GOOGLE_API_KEY
				}
			}

			randI := randgen.Intn(len(res.Results))
			l.Lock()
			places[i] = res.Results[randI]
			l.Unlock()
		}(r, i)
	}

	w.Wait()
	return places
}
