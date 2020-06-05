// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ingrammicro/cio/logger"

	log "github.com/sirupsen/logrus"
)

// checkStandardStatus return error if status is not OK
func checkStandardStatus(status int, response []byte) error {
	if status < 300 {
		return nil
	}

	// default, raw, not parsing
	message := string(response[:])

	var responseContent map[string]interface{}
	err := json.Unmarshal(response, &responseContent)
	if err == nil {
		if responseContent["errors"] != nil {
			message = ""
			for key, value := range responseContent["errors"].(map[string]interface{}) {
				subMessages := make([]string, len(value.([]interface{})))
				for i, v := range value.([]interface{}) {
					subMessages[i] = fmt.Sprint(v)
				}
				composedMsg := strings.Join(subMessages, ",")
				message = fmt.Sprintf("%s#%s:%s", message, key, composedMsg)
			}
		}
		if responseContent["errors"] == nil && responseContent["error"] != nil {
			message = responseContent["error"].(string)
		}
	}
	return fmt.Errorf("HTTP request failed: (%d) [%s]", status, message)
}

// UploadFile uploads a file to target url
func (imco *IMCOClient) UploadFile(sourceFilePath, targetURL string) error {
	logger.DebugFuncInfo()

	data, status, err := imco.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}
	if err = checkStandardStatus(status, data); err != nil {
		return err
	}
	return nil
}

// DownloadFile gets a file from given url saving file into given file path
func (imco *IMCOClient) DownloadFile(url string, filepath string, discoveryFileName bool,
) (realFileName string, status int, err error) {
	logger.DebugFuncInfo()

	realFileName, status, err = imco.GetFile(url, filepath, discoveryFileName)
	if err != nil {
		return realFileName, status, err
	}
	return realFileName, status, nil
}

func (imco *IMCOClient) getAndCheck(path string, bCheck bool, i interface{}) (int, error) {
	data, status, err := imco.Get(path)
	if err != nil {
		return status, err
	}
	if bCheck {
		if err = checkStandardStatus(status, data); err != nil {
			return status, err
		}
	}
	if i != nil {
		if err = json.Unmarshal(data, &i); err != nil {
			return status, err
		}
	}
	return status, nil
}

func (imco *IMCOClient) putAndCheck(
	path string,
	payload *map[string]interface{},
	bCheck bool,
	i interface{},
) (int, error) {
	data, status, err := imco.Put(path, payload)
	if err != nil {
		return status, err
	}
	if bCheck {
		if err = checkStandardStatus(status, data); err != nil {
			return status, err
		}
	}
	if i != nil {
		if err = json.Unmarshal(data, &i); err != nil {
			return status, err
		}
	}
	return status, nil
}

func (imco *IMCOClient) postAndCheck(
	path string,
	payload *map[string]interface{},
	bCheck bool,
	i interface{},
) (int, error) {
	data, status, err := imco.Post(path, payload)
	if err != nil {
		return status, err
	}
	if bCheck {
		if err = checkStandardStatus(status, data); err != nil {
			return status, err
		}
	}
	if i != nil {
		if err = json.Unmarshal(data, &i); err != nil {
			return status, err
		}
	}
	return status, nil
}

func (imco *IMCOClient) deleteAndCheck(path string, bCheck bool, i interface{}) (int, error) {
	data, status, err := imco.Delete(path)
	if err != nil {
		return status, err
	}
	if bCheck {
		if err = checkStandardStatus(status, data); err != nil {
			return status, err
		}
	}
	if i != nil {
		if err = json.Unmarshal(data, &i); err != nil {
			return status, err
		}
	}
	return status, nil
}

//func (imco *IMCOClient) getAndCheck(path string, bCheck bool) ([]byte, int, error) {
//	data, status, err := imco.Get(path)
//	if err != nil {
//		return nil, status, err
//	}
//	if bCheck {
//		if err = checkStandardStatus(status, data); err != nil {
//			return nil, status, err
//		}
//	}
//	return data, status, nil
//}
//
//func (imco *IMCOClient) putAndCheck(path string, payload *map[string]interface{}, bCheck bool) ([]byte, int, error) {
//	data, status, err := imco.Put(path, payload)
//	if err != nil {
//		return nil, status, err
//	}
//	if bCheck {
//		if err = checkStandardStatus(status, data); err != nil {
//			return nil, status, err
//		}
//	}
//	return data, status, nil
//}
//
//func (imco *IMCOClient) postAndCheck(path string, payload *map[string]interface{}, bCheck bool) ([]byte, int, error) {
//	data, status, err := imco.Post(path, payload)
//	if err != nil {
//		return nil, status, err
//	}
//	if bCheck {
//		if err = checkStandardStatus(status, data); err != nil {
//			return nil, status, err
//		}
//	}
//	return data, status, nil
//}
//
//func (imco *IMCOClient) deleteAndCheck(path string, bCheck bool) ([]byte, int, error) {
//	data, status, err := imco.Delete(path)
//	if err != nil {
//		return nil, status, err
//	}
//	if bCheck {
//		if err = checkStandardStatus(status, data); err != nil {
//			return nil, status, err
//		}
//	}
//	return data, status, nil
//}

// Post sends POST request to Concerto API
func (imco *IMCOClient) Post(path string, payload *map[string]interface{}) ([]byte, int, error) {
	url, jsPayload, err := imco.prepareCall(path, payload)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending POST request to %s with payload %v ", url, jsPayload)
	req, err := http.NewRequest("POST", url, jsPayload)
	req.Header.Add("Content-Type", ContentTypeApplicationJson)
	if imco.config.BrownfieldToken != "" {
		log.Debugf(
			"Including brownfield token %s in POST request as X-Concerto-Brownfield-Token header ",
			imco.config.BrownfieldToken,
		)
		req.Header.Add("X-Concerto-Brownfield-Token", imco.config.BrownfieldToken)
	}
	if imco.config.CommandPollingToken != "" && imco.config.ServerID != "" {
		log.Debugf(
			"Including command polling token %s in POST request as X-IMCO-Command-Polling-Token header ",
			imco.config.CommandPollingToken,
		)
		req.Header.Add("X-IMCO-Command-Polling-Token", imco.config.CommandPollingToken)
		log.Debugf("Including server id %s in POST request as X-IMCO-server-ID header ", imco.config.ServerID)
		req.Header.Add("X-IMCO-server-ID", imco.config.ServerID)
	}
	response, err := imco.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	return imco.receiveResponse(response)
}

// Put sends PUT request to Concerto API
func (imco *IMCOClient) Put(path string, payload *map[string]interface{}) ([]byte, int, error) {
	url, jsPayload, err := imco.prepareCall(path, payload)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending PUT request to %s with payload %v ", url, jsPayload)
	request, err := http.NewRequest("PUT", url, jsPayload)
	if err != nil {
		return nil, 0, err
	}
	request.Header = map[string][]string{"Content-type": {ContentTypeApplicationJson}}
	response, err := imco.client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	return imco.receiveResponse(response)
}

// Delete sends DELETE request to Concerto API
func (imco *IMCOClient) Delete(path string) ([]byte, int, error) {
	url, _, err := imco.prepareCall(path, nil)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending DELETE request to %s", url)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, 0, err
	}
	request.Header = map[string][]string{"Content-type": {ContentTypeApplicationJson}}
	response, err := imco.client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	return imco.receiveResponse(response)
}

// Get sends GET request to Concerto API
func (imco *IMCOClient) Get(path string) ([]byte, int, error) {
	url, _, err := imco.prepareCall(path, nil)
	if err != nil {
		return nil, 0, err
	}

	log.Debugf("Sending GET request to %s", url)
	response, err := imco.client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	return imco.receiveResponse(response)
}

// GetFile sends GET request to Concerto API and receives a file
func (imco *IMCOClient) GetFile(url string, filePath string, discoveryFileName bool) (string, int, error) {
	log.Debugf("Sending GET request to %s", url)
	response, err := imco.client.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer response.Body.Close()
	log.Debugf("Status code:%d message:%s", response.StatusCode, response.Status)

	realFileName := filePath
	if discoveryFileName {
		r, err := regexp.Compile("filename=\\\"([^\\\"]*){1}\\\"")
		if err != nil {
			return "", 0, err
		}
		realFileName = fmt.Sprintf(
			"%s/%s",
			filePath,
			r.FindStringSubmatch(response.Header.Get("Content-Disposition"))[1],
		)
	}

	output, err := os.Create(realFileName)
	if err != nil {
		return "", response.StatusCode, err
	}
	defer output.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		return "", response.StatusCode, err
	}

	log.Debugf("%#v bytes downloaded", n)
	return realFileName, response.StatusCode, nil
}

// PutFile sends PUT request to send a file
func (imco *IMCOClient) PutFile(sourceFilePath string, targetURL string) ([]byte, int, error) {
	data, err := os.Open(sourceFilePath)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PUT", targetURL, data)
	if err != nil {
		return nil, 0, err
	}

	res, err := imco.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	status := res.StatusCode
	return buf, status, nil
}

func (imco *IMCOClient) prepareCall(path string, payload *map[string]interface{},
) (url string, jsPayload *strings.Reader, err error) {
	if imco.config == nil || imco.client == nil {
		return "", nil, fmt.Errorf("cannot call web service without loading configuration")
	}

	url = fmt.Sprintf("%s%s", imco.config.APIEndpoint, path)

	if payload == nil {
		return url, nil, nil
	}

	// payload to json
	json, err := json.Marshal(payload)
	if err != nil {
		return "", nil, err
	}
	return url, strings.NewReader(string(json)), err
}

func (imco *IMCOClient) receiveResponse(response *http.Response) (body []byte, status int, err error) {
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	log.Debugf("Response : %s", body)
	log.Debugf("Status code: (%d) %s", response.StatusCode, response.Status)
	return body, response.StatusCode, nil
}
