package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Printf("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	writen bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.writen = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.writen = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.writen = true
}

func TestServer(t *testing.T) {
	t.Run("non-interupted request", func(t *testing.T) {
		data := "hello, world!"
		spyStore := &SpyStore{data, t}
		svr := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q, want %q", response.Body.String(), data)
		}
	})

	t.Run("tells the store to cancel work, if request is cancelled", func(t *testing.T) {
		data := "hello, world!"
		spyStore := &SpyStore{data, t}
		svr := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.writen {
			t.Error("a response should not have been written")
		}
	})
}
