package testinglocal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/client"
)

type poolPayload struct {
	Data string `json:"data"`
}

type poolResult struct {
	want string
	got  string
	err  error
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var p poolPayload

	_ = json.NewDecoder(r.Body).Decode(&p)
	_ = json.NewEncoder(w).Encode(p)
}

func TestBufferPoolUnderMemoryPressure(t *testing.T) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		echoHandler(w, r)
	}))
	defer srv.Close()

	c := client.Client{Client: srv.Client(), BaseUrl: srv.URL}

	payloads := make([]string, 50)
	for i := range payloads {
		payloads[i] = strings.Repeat(string(rune('a'+i%26)), 100+i*10)
	}

	for i := range 5 {
		for _, data := range payloads {
			payload := poolPayload{Data: data}
			resp, err := client.Post[poolPayload, poolPayload](
				context.Background(), c, "test", nil, payload,
			)

			if err != nil {
				t.Fatalf("iteration %d payload %q: %v", i, data[:10], err)
			}

			if resp.Data != data {
				t.Fatalf("data corruption: got %q, want %q (len: got %d, want %d)",
					resp.Data, data, len(resp.Data), len(data))
			}
		}
	}

	runtime.GC()
	runtime.GC()
}

func TestConcurrentPoolRequests(t *testing.T) {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		echoHandler(w, r)
	}))
	defer srv.Close()

	c := client.Client{Client: srv.Client(), BaseUrl: srv.URL}

	ch := make(chan poolResult, 50)

	var wg sync.WaitGroup

	for i := range 50 {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			data := strings.Repeat(string(rune('A'+idx%26)), 200+idx)
			payload := poolPayload{Data: data}

			resp, err := client.Post[poolPayload, poolPayload](
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
