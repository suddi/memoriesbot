package photos

import (
	"fmt"
	"io/ioutil"
	"memoriesbot/pkg/config"
	"net/http"
	"strings"
	"time"
)

// MakeRequest - interact with Google Photos API
func MakeRequest() (string, error) {
	c := config.Get()
	t := time.Now()

	requestBody := strings.NewReader(fmt.Sprintf(`
        {
            "filters": {
                "contentFilter": {
                    "includedContentCategories": [
                        "PEOPLE"
                    ]
                },
                "dateFilter": {
                    "dates": [
                        {
                            "month": %d
                        }
                    ],
                    "ranges": [
                        {
                            "startDate": {
                                "month": 5,
                                "year": 2018
                            },
                            "endDate": {
                                "day": %d,
                                "month": %d,
                                "year": %d
                            }
                        }
                    ]
                },
                "mediaTypeFilter": {
                    "mediaTypes": [
                        "PHOTO"
                    ]
                }
            }
        }
    `, t.Month(), t.Day(), t.Month(), t.Year()))

	client := http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://photoslibrary.googleapis.com/v1/mediaItems:search",
		requestBody,
	)

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Auth.GooglePhotos.AccessToken))

	res, err := client.Do(req)

	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
