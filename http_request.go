package goutils

import (
    "time"
    "github.com/mozillazg/request"
    "fmt"
    "errors"
    "net/http"
    "encoding/json"
)

const DEFAULT_REQUEST_TIMEOUT = 30

func MakePostRequest(url string, data map[string]string, headers map[string]string) (result map[string]interface{}, err error) {
    req := request.NewRequest(new(http.Client))
    req.Client.Timeout = time.Duration(DEFAULT_REQUEST_TIMEOUT * time.Second)
    req.Data = data
    req.Headers = headers

    resp, err := req.Post(url)
    if err != nil {
        err = errors.New(fmt.Sprintf("Error POST request (req.Post phrase) to %s, message: %s", url, err.Error()))
        return
    }

    var content []byte
    content, err = resp.Content()
    if err != nil {
        err = errors.New(fmt.Sprintf("Error POST request (resp.Content phrase) to %s, message: %s, response: %s", url, err.Error(), string(content)))
        return
    }

    err = json.Unmarshal(content, &result)
    if err != nil {
        err = errors.New(fmt.Sprintf("Error POST request (json.Unmarshal phrase) to %s, message: %s, response: %s", url, err.Error(), string(content)))
        return
    }

    return
}

func MakeGetRequest(url string, data, headers map[string]string) (result map[string]interface{}, err error) {
    req := request.NewRequest(new(http.Client))
    req.Client.Timeout = time.Duration(DEFAULT_REQUEST_TIMEOUT * time.Second)
    req.Data = data
    req.Headers = headers

    resp, err := req.Get(url)
    if err != nil {
        err = errors.New(fmt.Sprintf("Error GET request (req.Post phrase) to %s, message: %s", url, err.Error()))
        return
    }

    var content []byte
    content, err = resp.Content()
    if err != nil {
        err = errors.New(fmt.Sprintf("Error GET request (resp.Content phrase) to %s, message: %s, response: %s", url, err.Error(), string(content)))
        return
    }

    err = json.Unmarshal(content, &result)
    if err != nil {
        err = errors.New(fmt.Sprintf("Error GET request (json.Unmarshal phrase) to %s, message: %s, response: %s", url, err.Error(), string(content)))
        return
    }

    return
}
