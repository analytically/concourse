//go:build linux

package runtime_test

import (
	"os"
	"path"

	"code.cloudfoundry.org/localip"
	"github.com/concourse/concourse/worker/runtime"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ResolveconfParserSuite struct {
	suite.Suite
	*require.Assertions
}

func (s *ResolveconfParserSuite) TestParseHostResolvConf() {
	file := `
nameserver 8.8.8.8
nameserver 127.0.0.16
nameserver something 9.9.9.9
search something
`

	tmpDir, _ := os.MkdirTemp("", "test-resolv")
	defer os.RemoveAll(tmpDir)
	os.WriteFile(path.Join(tmpDir, "resolv.conf"), []byte(file), 0644)

	entries, err := runtime.ParseHostResolveConf(path.Join(tmpDir, "resolv.conf"))
	s.NoError(err)

	s.Equal([]string{"nameserver 8.8.8.8", "search something"}, entries)
}

func (s *ResolveconfParserSuite) TestParseHostResolvConfWithLoopback() {
	file := `nameserver 127.0.0.1`

	tmpDir, _ := os.MkdirTemp("", "test-resolv-noloopback")
	defer os.RemoveAll(tmpDir)
	os.WriteFile(path.Join(tmpDir, "resolv.conf"), []byte(file), 0644)

	entries, err := runtime.ParseHostResolveConf(path.Join(tmpDir, "resolv.conf"))
	s.NoError(err)

	ip, _ := localip.LocalIP()
	s.Equal([]string{"nameserver " + ip}, entries)
}
