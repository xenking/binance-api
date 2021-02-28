package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/segmentio/encoding/json"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

func newRestClient(apikey, secret string) *restClient {
	return &restClient{
		window: 5000,
		apikey: apikey,
		hmac: hmac.New(sha256.New, s2b(secret)),
		client: newHTTPClient(),
	}
}

// restClient represents the actual HTTP restClient, that is being used to interact with binance API server
type restClient struct {
	apikey string
	hmac   hash.Hash
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
	encoded := values.Encode()
	var pb []byte
	if sign {
		pb = make([]byte, len(encoded), len(encoded)+116)
	} else {
		pb = make([]byte, len(encoded))
	}
	copy(pb, encoded)
	// Signed requests require the additional timestamp, window size and signature of the payload
	// Remark: This is done only to routes with actual data
	if sign {
		buf := bytebufferpool.Get()
		pb = append(pb, "&timestamp="...)
		pb = append(pb, strconv.AppendInt(buf.B, time.Now().UnixNano()/(1000*1000), 10)...)
		buf.Reset()
		pb = append(pb, "&recvWindow="...)
		pb = append(pb, strconv.AppendInt(buf.B, int64(c.window), 10)...)

		_, err = c.hmac.Write(pb)
		if err != nil {
			return nil, err
		}
		pb = append(pb, "&signature="...)
		sum := c.hmac.Sum(nil)
		enc := make([]byte, len(sum)*2)
		hex.Encode(enc, sum)
		pb = append(pb, enc...)

		c.hmac.Reset()
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

	if method == fasthttp.MethodGet {
		b.Grow(len(pb)+1)
		b.WriteByte('?')
		b.Write(pb)
	} else {
		req.Header.Add("Content-Type", defaultHeaderForm)
		req.SetBody(pb)
	}
	req.SetRequestURI(b.String())

	if sign || stream {
		req.Header.Add("X-MBX-APIKEY", c.apikey)
	}

	req.Header.Add("Accept", defaultHeaderJson)
	resp := fasthttp.AcquireResponse()

	if err = c.client.Do(req, resp); err != nil {
		return nil, err
	}
	fasthttp.ReleaseRequest(req)

	body := append(resp.Body())
	status := resp.StatusCode()

	fasthttp.ReleaseResponse(resp)

	if status != fasthttp.StatusOK {
		apiErr := &APIError{}
		if err = json.Unmarshal(body, apiErr); err != nil {
			return nil, err
		}
		return nil, apiErr
	}
	return body, err
}
