package goutils

import (
    "time"
    "github.com/mozillazg/request"
    "fmt"
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
        err = fmt.Errorf("Error POST request (req.Post phrase) to %s, message: %s", url, err.Error())
        return
    }

    var content []byte
    content, err = resp.Content()
    if err != nil {
        err = fmt.Errorf("Error POST request (resp.Content phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
        return
    }

    err = json.Unmarshal(content, &result)
    if err != nil {
        err = fmt.Errorf("Error POST request (json.Unmarsha phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
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
        err = fmt.Errorf("Error GET request (req.Get phrase) to %s, message: %s", url, err.Error())
        return
    }

    var content []byte
    content, err = resp.Content()
    if err != nil {
        err = fmt.Errorf("Error GET request (resp.Content phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
        return
    }

    err = json.Unmarshal(content, &result)
    if err != nil {
        err = fmt.Errorf("Error GET request (json.Unmarshal phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
        return
    }

    return
}

func MakeGetRequestToTarget(url string, dest *interface{}) error {
    req := request.NewRequest(new(http.Client))
    req.Client.Timeout = time.Duration(DEFAULT_REQUEST_TIMEOUT * time.Second)

    resp, err := req.Get(url)
    if err != nil {
        return fmt.Errorf("Error GET request (req.Get phrase) to %s, message: %s", url, err.Error())
    }

    var content []byte
    content, err = resp.Content()
    if err != nil {
        return fmt.Errorf("Error GET request (resp.Content phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
    }

    err = json.Unmarshal(content, dest)
    if err != nil {
        return fmt.Errorf("Error GET request (json.Unmarshal phrase) to %s, message: %s, response: %s", url, err.Error(), string(content))
    }

    return nil
}
