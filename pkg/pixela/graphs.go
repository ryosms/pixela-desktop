package pixela

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type GraphDefinition struct {
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	Unit                string   `json:"unit"`
	Type                string   `json:"type"`
	Color               string   `json:"color"`
	Timezone            string   `json:"timezone"`
	PurgeCacheUrls      []string `json:"purgeCacheURLs"`
	SelfSufficient      string   `json:"selfSufficient"`
	IsSecret            bool     `json:"isSecret"`
	PublishOptionalData bool     `json:"publishOptionalData"`
}

type listGraphs struct {
	Graphs []GraphDefinition `json:"graphs"`
}

func GetGraphDefinitions(username string, token string) (*[]GraphDefinition, error) {
	url := GenerateUrl("users", username, "graphs")
	req, err := generateRequest("GET", url, &token, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, string(b))
	}

	var graphs listGraphs
	err = json.Unmarshal(b, &graphs)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &graphs.Graphs, nil
}

type GraphStats struct {
	TotalPixelsCount int     `json:"totalPixelsCount"`
	MaxQuantity      float64 `json:"maxQuantity"`
	MinQuantity      float64 `json:"minQuantity"`
	TotalQuantity    float64 `json:"totalQuantity"`
	AvgQuantity      float64 `json:"AvgQuantity"`
	TodaysQuantity   float64 `json:"TodaysQuantity"`
}

func GetGraphStats(username string, graphId string) (*GraphStats, error) {
	path := fmt.Sprintf("users/%s/graphs/%s/stats", username, graphId)
	req, err := generateRequest("GET", path, nil, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", res.StatusCode, string(b))
	}

	var stats GraphStats
	err = json.Unmarshal(b, &stats)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &stats, nil
}
