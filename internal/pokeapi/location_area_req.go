package pokeapi


import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {

	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	// method, URL, body
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	// executes the http request
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResp{}, err
	}


	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResp{}, err
	}

	// parse json into struct

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil

}