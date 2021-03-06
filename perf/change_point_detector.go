package perf

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/evergreen-ci/cedar/model"
	"github.com/evergreen-ci/cedar/util"
	"github.com/evergreen-ci/gimlet"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
)

type ChangeDetector interface {
	DetectChanges(context.Context, []float64) ([]model.ChangePoint, error)
}
type jsonChangePoint struct {
	Index     int
	Algorithm jsonAlgorithm
}

type jsonAlgorithm struct {
	Name          string
	Version       int
	Configuration map[string]interface{}
}

type signalProcessingClient struct {
	user    string
	token   string
	baseURL string
}

func NewMicroServiceChangeDetector(baseURL, user string, token string) ChangeDetector {
	return &signalProcessingClient{user: user, token: token, baseURL: baseURL}
}

func (spc *signalProcessingClient) DetectChanges(ctx context.Context, series []float64) ([]model.ChangePoint, error) {
	startAt := time.Now()

	jsonChangePoints := &struct {
		ChangePoints []jsonChangePoint `json:"changePoints"`
	}{}

	request := &struct {
		Series []float64 `json:"series"`
	}{
		Series: series,
	}

	if err := spc.doRequest(http.MethodPost, spc.baseURL+"/change_points/detect", ctx, request, jsonChangePoints); err != nil {
		return nil, errors.WithStack(err)
	}

	var result []model.ChangePoint
	for _, point := range jsonChangePoints.ChangePoints {
		mapped := model.ChangePoint{
			Index: point.Index,
			Algorithm: model.AlgorithmInfo{
				Name:    point.Algorithm.Name,
				Version: point.Algorithm.Version,
			},
		}

		for k, v := range point.Algorithm.Configuration {
			additionalOption := model.AlgorithmOption{
				Name:  k,
				Value: v,
			}
			mapped.Algorithm.Options = append(mapped.Algorithm.Options, additionalOption)
		}

		result = append(result, mapped)
	}

	grip.Debug(message.Fields{
		"message":        "change point detection completed",
		"num_points":     len(series),
		"cp_detected":    len(result),
		"duration_secs":  time.Since(startAt).Seconds(),
		"implementation": "MicroServiceChangePointDetector",
	})

	return result, nil
}

func (spc *signalProcessingClient) doRequest(method, route string, ctx context.Context, in, out interface{}) error {
	body, err := json.Marshal(in)
	if err != nil {
		return errors.WithStack(err)
	}

	conf := util.HTTPRetryConfiguration{
		MaxRetries:      50,
		TemporaryErrors: true,
		MaxDelay:        30 * time.Second,
		BaseDelay:       50 * time.Millisecond,
		Methods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
		},
		Statuses: []int{
			// status code for timeouts from ELB in AWS (Kanopy infrastructure)
			499,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
			http.StatusInsufficientStorage,
			http.StatusConflict,
			http.StatusRequestTimeout,
			http.StatusPreconditionFailed,
			http.StatusExpectationFailed,
		},
		RetryableErrors: []error{
			// If a connection gets cut by the ELB, sometimes the client doesn't get an actual error
			// The client only receives a nil body leading to an EOF
			io.EOF,
		},
	}
	client := util.GetHTTPRetryableClient(conf)
	defer util.PutHTTPClient(client)

	req, err := http.NewRequest(method, route, bytes.NewBuffer(body))
	if err != nil {
		return errors.WithStack(err)
	}
	req = req.WithContext(ctx)
	req.Header.Add("Cookie", fmt.Sprintf("auth_user=%v;auth_token=%v", spc.user, spc.token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		grip.Warning(message.Fields{
			"message":   "Failed to detect changes in metric data",
			"status":    http.StatusText(resp.StatusCode),
			"url":       route,
			"auth_user": spc.user,
		})
		return errors.Errorf("Failed to detect changes in metric data, status: %q, url: %q, auth_user: %q", http.StatusText(resp.StatusCode), route, spc.user)
	}

	if err = gimlet.GetJSON(resp.Body, out); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
