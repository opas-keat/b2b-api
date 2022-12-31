package shareutils

import (
	"bytes"
	"context"
	"io"
	"models"
	"net/http"
	"time"
)

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 120,
	}
}

func Post[T any, R any](ctx context.Context, url string, body *T) (*models.Response[R], error) {
	jsonBytes, err := ToJson(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json")
	//setHttpHeader(ctx, request)
	response, err := newHttpClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		r1, err := FromJson[models.Response[interface{}]](responseBody)
		if err != nil {
			return nil, err
		}
		err2 := models.Response[R]{
			StatusCode: r1.StatusCode,
			Code:       r1.Code,
			Message:    r1.Message,
		}
		return &err2, nil
	}
	r2, err := FromJson[models.Response[R]](responseBody)
	if err != nil {
		return nil, err
	}
	return r2, nil
}

func Get[R any](ctx context.Context, url string) (*models.Response[R], error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	//setHttpHeader(ctx, request)
	response, err := newHttpClient().Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		r1, err := FromJson[models.Response[interface{}]](responseBody)
		if err != nil {
			return nil, err
		}
		err2 := models.Response[R]{
			StatusCode: r1.StatusCode,
			Code:       r1.Code,
			Message:    r1.Message,
		}
		return &err2, nil
	}
	r2, err := FromJson[models.Response[R]](responseBody)
	if err != nil {
		return nil, err
	}
	return r2, nil
}
