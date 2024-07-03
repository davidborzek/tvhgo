package tvheadend_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidborzek/tvhgo/tvheadend"
	"github.com/stretchr/testify/assert"
)

type testResponse struct {
	Test string `json:"test"`
}

func TestQuery(t *testing.T) {
	q := tvheadend.NewQuery()

	q.Limit(10)
	q.Start(5)
	q.SortDir("asc")
	q.SortKey("someKey")
	q.Set("testField", "testValue")

	assert.Equal(t, "10", q.Get("limit"))
	assert.Equal(t, "5", q.Get("start"))
	assert.Equal(t, "asc", q.Get("dir"))
	assert.Equal(t, "someKey", q.Get("sort"))
	assert.Equal(t, "testValue", q.Get("testField"))
}

func TestClientExec(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, _ := r.BasicAuth()
		contentType := r.Header.Get("Content-Type")

		b, _ := io.ReadAll(r.Body)

		assert.Equal(t, "limit=10", string(b))

		assert.Equal(t, "someUsername", username)
		assert.Equal(t, "somePassword", password)
		assert.Equal(t, "application/x-www-form-urlencoded", contentType)
		assert.Equal(t, "/some/path", r.URL.Path)

		json.NewEncoder(w).Encode(&testResponse{
			Test: "testValue",
		})
	}))

	client := tvheadend.New(tvheadend.ClientOpts{
		URL:      srv.URL,
		Username: "someUsername",
		Password: "somePassword",
	})

	q := tvheadend.NewQuery()
	q.Limit(10)

	var model testResponse
	res, err := client.Exec(context.TODO(), "/some/path", &model, q)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "testValue", model.Test)
}

func TestClientExecNoDest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/some/path", r.URL.Path)
	}))

	client := tvheadend.New(tvheadend.ClientOpts{
		URL: srv.URL,
	})
	res, err := client.Exec(context.TODO(), "/some/path", nil)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestClientExecWithoutAuth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.Header.Get("Authorization"))
		assert.Equal(t, "/some/path", r.URL.Path)
	}))

	client := tvheadend.New(tvheadend.ClientOpts{
		URL: srv.URL,
	})
	res, err := client.Exec(context.TODO(), "/some/path", nil)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
