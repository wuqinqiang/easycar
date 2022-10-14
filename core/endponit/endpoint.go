package endponit

import "net/url"

// Endpoint is registry endpoint.
type Endpoint interface {
	Endpoint() *url.URL
}

func GetHostByEndpoint(endpoints []string, scheme string) (string, error) {
	for _, e := range endpoints {
		u, err := url.Parse(e)
		if err != nil {
			return "", err
		}
		if u.Scheme == scheme {
			return u.Host, nil
		}
	}
	return "", nil
}
