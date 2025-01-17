package ovs

import (
	"fmt"
	"os/exec"
	"strings"
)

type OvsDumpSourceCLI struct{}

func ovsCtlRun(params ...string) ([]string, error) {
	cmd := exec.Command("ovs-ofctl", params...)
	out, err := cmd.Output()
	outString := string(out)
	//if error was occured we return
	fmt.Println(outString)
	if err != nil {
		return nil, err
	}
	//if command was succesfull we further parse the output

	lines := strings.Split(outString, "\n")
	//skip the first and last lines, since it is just a response header and an empty line
	lines = lines[1:(len(lines) - 1)]
	return lines, nil
}

func (o OvsDumpSourceCLI) DumpFlows(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-flows", "br-int")
}

func (o OvsDumpSourceCLI) DumpPorts(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-ports", "br-int")
}

func (o OvsDumpSourceCLI) DumpGroups(ip string, port int) ([]string, error) {
	return ovsCtlRun("-O", "openflow13", "dump-groups", "br-int")
}

func (o OvsDumpSourceCLI) DumpGroupStats(ip string, port int) ([]string, error) {
	return ovsCtlRun("-O", "openflow13", "dump-group-stats", "br-int")
}
