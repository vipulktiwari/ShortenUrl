package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
	"vipulktiwari/ShortenUrl/api"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Our basic service, which will have all the api as interface funtions
type BasicService interface {
	CreateURL(context.Context, api.CreateURLReq) (string, error)
	AccessURL(context.Context, api.AccessURLReq) (string, error)
}

type basicService struct{}

var URLMap map[string]string

func (b basicService) CreateURL(ctx context.Context, req api.CreateURLReq) (string, error) {
	//implement
	fmt.Println("url : ", req.URL)
	s := StringWithCharset(12, charset)
	_, ok := URLMap[s]
	if !ok {
		URLMap[s] = req.URL
	} else {
		return "", errors.New("shortend url already exist")
	}
	return s, nil
}

func (b basicService) AccessURL(ctx context.Context, req api.AccessURLReq) (string, error) {
	//implement
	fmt.Println("url : ", req.URL)
	s, ok := URLMap[req.URL]
	if !ok {
		return "", errors.New("Url not found")
	}

	return s, nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeCreateUrlEndpoint(svc BasicService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		r := req.(api.CreateURLReq)
		resp, err = svc.CreateURL(ctx, r)
		return resp, err
	}
}
func makeAccessUrlEndpoint(svc BasicService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		r := req.(api.AccessURLReq)
		resp, err = svc.AccessURL(ctx, r)
		return resp, err
	}
}

func decodeCreateUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// as ping is a get request we will do nothing here
	var req api.CreateURLReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Kindly provide request body")
		return nil, err
	}

	json.Unmarshal(reqBody, &req)
	return req, nil
}

func decodeAccessUrlRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// as ping is a get request we will do nothing here
	var req api.AccessURLReq
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Kindly provide request body")
		return nil, err
	}

	json.Unmarshal(reqBody, &req)
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := basicService{}
	accessUrlhandler := httptransport.NewServer(
		makeAccessUrlEndpoint(svc),
		decodeAccessUrlRequest,
		encodeResponse,
	)
	createUrlhandler := httptransport.NewServer(
		makeCreateUrlEndpoint(svc),
		decodeCreateUrlRequest,
		encodeResponse,
	)

	URLMap = make(map[string]string)

	log.Println("server started, listening on port 8080")
	http.Handle("/createurl", createUrlhandler)
	http.Handle("/accessurl", accessUrlhandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
