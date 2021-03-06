//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
package sdb

import (
	"gopkg.in/amz.v1/aws"
)

func Sign(auth aws.Auth, method, path string, params map[string][]string, headers map[string][]string) {
	sign(auth, method, path, params, headers)
}
