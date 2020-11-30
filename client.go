package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

// restClient represents the actual HTTP restClient, that is being used to interact with binance API server
type restClient struct {
	apikey string
	secret string
	client *fasthttp.HostClient
	window int
}

const (
	defaultHeaderJson = "application/json"
	defaultHeaderForm = "application/x-www-form-urlencoded"
	defaultSchema     = "https"
)

// newHTTPClient create fasthttp.HostClient with default settings
func newHTTPClient() *fasthttp.HostClient {
	return &fasthttp.HostClient{
		NoDefaultUserAgentHeader:      true, // Don't send: User-Agent: fasthttp
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        false,
		IsTLS:                         true,
		Name:                          DefaultUserAgent,
		Addr:                          BaseHost,
	}
}

// do invokes the given API command with the given data
// sign indicates whether the api call should be done with signed payload
// stream indicates if the request is stream related
func (c *restClient) do(method, endpoint string, data interface{}, sign bool, stream bool) ([]byte, error) {
	// Convert the given data to urlencoded format
	values, err := query.Values(data)
	if err != nil {
		return nil, err
	}

	var pb strings.Builder
	pb.WriteString(values.Encode())
	// Signed requests require the additional timestamp, window size and signature of the payload
	// Remark: This is done only to routes with actual data
	if sign {
		buf := bytebufferpool.Get()
		pb.WriteString("&timestamp=")
		pb.Write(strconv.AppendInt(buf.B, time.Now().UnixNano()/(1000*1000), 10))
		pb.WriteString("&recvWindow=")
		pb.Write(strconv.AppendInt(buf.B, int64(c.window), 10))

		mac := hmac.New(sha256.New, []byte(c.secret))
		_, err = mac.Write([]byte(pb.String()))
		if err != nil {
			return nil, err
		}
		pb.WriteString("&signature=")
		pb.WriteString(hex.EncodeToString(mac.Sum(nil)))
		bytebufferpool.Put(buf)
	}

	var b strings.Builder
	b.WriteString(endpoint)

	// Construct the http request
	// Remark: GET requests payload is as a query parameters
	// POST requests payload is given as a body
	req := fasthttp.AcquireRequest()
	req.Header.SetHost(BaseHost)
	req.URI().SetScheme(defaultSchema)
	req.Header.SetMethod(method)

	if method == http.MethodGet {
		b.WriteByte('?')
		b.WriteString(pb.String())
	} else {
		req.Header.Add("Content-Type", defaultHeaderForm)
		req.SetBodyString(pb.String())
	}
	req.SetRequestURI(b.String())

	if sign || stream {
		req.Header.Add("X-MBX-APIKEY", c.apikey)
	}

	req.Header.Add("Accept", defaultHeaderJson)
	resp := fasthttp.AcquireResponse()

	if err := c.client.Do(req, resp); err != nil {
		return nil, err
	}
	fasthttp.ReleaseRequest(req)

	body := append(resp.Body())
	status := resp.StatusCode()

	fasthttp.ReleaseResponse(resp)

	if status != fasthttp.StatusOK {
		return nil, APIError{
			Code:    status,
			Message: body,
		}
	}
	return body, err
}
