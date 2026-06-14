package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"go.uber.org/mock/gomock"
)

type poolPayload struct {
	Data string `json:"data"`
}

type poolResult struct {
	want string
	got  string
	err  error
}

func TestBufferPoolUnderMemoryPressure(t *testing.T) {
	// Disable GC during the test to guarantee that dirty buffers are retained
	// in the sync.Pool and reused. This allows us to verify that buffers
	// are properly Reset() and do not suffer from data corruption.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDoer := NewMockDoer(ctrl)
	c := Client{Client: mockDoer, BaseUrl: "http://test"}

	mockDoer.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		body, _ := io.ReadAll(req.Body)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}, nil
	}).AnyTimes()

	for i := range 500 {
		data := strings.Repeat("1234567890", i%100)
		payload := poolPayload{Data: data}
		resp, _, err := Post[poolPayload, poolPayload](
			context.Background(), c, "test", nil, payload,
		)

		if err != nil {
			t.Fatalf("iteration %d payload %q: %v", i, data[:min(10, len(data))], err)
		}

		if resp.Data != data {
			t.Fatalf("data corruption: got %q, want %q (len: got %d, want %d)",
				resp.Data, data, len(resp.Data), len(data))
		}
	}

	runtime.GC()
	runtime.GC()
}

func TestConcurrentPoolRequests(t *testing.T) {
	// Disable GC during the test to guarantee that dirty buffers are retained
	// in the sync.Pool and reused. This allows us to verify that buffers
	// are properly Reset() and do not suffer from data corruption.
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDoer := NewMockDoer(ctrl)
	c := Client{Client: mockDoer, BaseUrl: "http://test"}

	mockDoer.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		body, _ := io.ReadAll(req.Body)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(body)),
		}, nil
	}).AnyTimes()

	ch := make(chan poolResult, 500)

	var wg sync.WaitGroup

	for i := range 500 {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			data := strings.Repeat("1234567890", idx%100)
			payload := poolPayload{Data: data}

			resp, _, err := Post[poolPayload, poolPayload](
				context.Background(), c, "test", nil, payload,
			)

			ch <- poolResult{want: data, got: resp.Data, err: err}
		}(i)
	}

	wg.Wait()
	close(ch)

	for r := range ch {
		if r.err != nil {
			t.Errorf("request failed: %v", r.err)

			continue
		}

		if r.got != r.want {
			t.Errorf("data corruption: got %q, want %q (len: got %d, want %d)",
				r.got, r.want, len(r.got), len(r.want))
		}
	}
}
