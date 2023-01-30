package host

import "fmt"

var _host string

func Init(host string) error {
	_host = host
	return nil
}

func URL(path string) string {
	return fmt.Sprintf("%s%s", _host, path)
}
