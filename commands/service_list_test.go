package commands

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/mesg-foundation/core/protobuf/coreapi"
	"github.com/mesg-foundation/core/protobuf/definitions"
	"github.com/stretchr/testify/require"
)

func TestServiceList(t *testing.T) {
	var (
		services = []*coreapi.Service{
			{Definition: &definitions.Service{Hash: "1", Name: "a"}, Status: coreapi.Service_RUNNING},
			{Definition: &definitions.Service{Hash: "2", Name: "b"}, Status: coreapi.Service_PARTIAL},
		}
		m = newMockExecutor()
		c = newServiceListCmd(m)
	)

	m.On("ServiceList").Return(services, nil)

	closeStd := captureStd(t)
	c.cmd.Execute()
	stdout, _ := closeStd()
	r := bufio.NewReader(strings.NewReader(stdout))

	for _, s := range []string{
		`Listing services\.\.\.`,
		`HASH\s+SID\s+NAME\s+STATUS`,
	} {
		matched, err := regexp.Match(fmt.Sprintf(`^\s*%s\s*$`, s), readLine(t, r))
		require.NoError(t, err)
		require.True(t, matched)
	}

	for _, s := range services {
		status := strings.ToLower(s.Status.String())
		pattern := fmt.Sprintf(`^\s*%s\s+%s\s+%s\s+%s\s*$`, s.Definition.Hash, s.Definition.Sid, s.Definition.Name, status)
		matched, err := regexp.Match(pattern, readLine(t, r))
		require.NoError(t, err)
		require.True(t, matched)
	}
}
