/*
https://golang.cafe/blog/golang-httptest-example.html
https://pkg.go.dev/net/http/httptest

*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

const HAND_OUTPUT = `# HELP flowAge The number of seconds have passed since the given OpenFlow entry was created
# TYPE flowAge gauge
flowAge{action="LOCAL",match="ip",priority="10",table="0"} 588.591
flowAge{action="LOCAL",match="ip,nw_dst=10.96.0.1",priority="70",table="0"} 588.59
flowAge{action="NORMAL",match="arp",priority="41000",table="0"} 588.593
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.140.235,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.13,tp_src=80",priority="100",table="2"} 63.654
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.235.60,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.12,tp_src=80",priority="100",table="2"} 69.049
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.17,tp_src=8079",priority="100",table="2"} 38.034
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.18,tp_src=8079",priority="100",table="2"} 36.956
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.7,tp_src=8079",priority="100",table="2"} 92.215
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.104.140.235,mod_tp_src:5672,resubmit(,4)",match="tcp,nw_src=10.244.7.14,tp_src=5672",priority="100",table="2"} 44.216
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.105.230.145,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.11,tp_src=80",priority="100",table="2"} 73.383
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.108.104.82,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.8,tp_src=80",priority="100",table="2"} 82.834
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.119.135,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.10,tp_src=80",priority="100",table="2"} 76.29
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.186.133,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.4,tp_src=27017",priority="100",table="2"} 103.409
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.110.235.87,mod_tp_src:3306,resubmit(,4)",match="tcp,nw_src=10.244.7.5,tp_src=3306",priority="100",table="2"} 97.361
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.111.233.21,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.9,tp_src=27017",priority="100",table="2"} 81.776
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 588.59
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 588.59
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 588.591
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 588.591
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.97.137.63,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.15,tp_src=80",priority="100",table="2"} 37.079
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.153.149,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.6,tp_src=80",priority="100",table="2"} 86.171
flowAge{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.168.141,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.16,tp_src=27017",priority="100",table="2"} 6.549
flowAge{action="mod_dl_dst:0a:58:0a:f4:07:01,LOCAL",match="ip,nw_dst=10.244.7.1",priority="90",table="0"} 588.592
flowAge{action="resubmit(,2)",match="ip,nw_src=10.244.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 588.592
flowAge{action="resubmit(,2)",match="ip,nw_src=172.16.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 588.592
flowAge{action="resubmit(,3)",match="ip,nw_src=10.244.0.0/16,nw_dst=10.96.0.0/12",priority="33",table="0"} 588.592
flowAge{action="resubmit(,4)",match="ip",priority="1",table="2"} 588.59
flowAge{action="resubmit(,4)",match="ip,nw_dst=10.244.0.0/16",priority="30",table="0"} 588.591
flowAge{action="resubmit(,5)",match="in_port=1",priority="50",table="0"} 588.592
# HELP flowBytes The number of bytes matched for the given OpenFlow entry
# TYPE flowBytes counter
flowBytes{action="LOCAL",match="ip",priority="10",table="0"} 45580
flowBytes{action="LOCAL",match="ip,nw_dst=10.96.0.1",priority="70",table="0"} 46620
flowBytes{action="NORMAL",match="arp",priority="41000",table="0"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.140.235,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.13,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.235.60,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.12,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.17,tp_src=8079",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.18,tp_src=8079",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.7,tp_src=8079",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.104.140.235,mod_tp_src:5672,resubmit(,4)",match="tcp,nw_src=10.244.7.14,tp_src=5672",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.105.230.145,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.11,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.108.104.82,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.8,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.119.135,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.10,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.186.133,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.4,tp_src=27017",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.110.235.87,mod_tp_src:3306,resubmit(,4)",match="tcp,nw_src=10.244.7.5,tp_src=3306",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.111.233.21,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.9,tp_src=27017",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 1290
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.97.137.63,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.15,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.153.149,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.6,tp_src=80",priority="100",table="2"} 0
flowBytes{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.168.141,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.16,tp_src=27017",priority="100",table="2"} 0
flowBytes{action="mod_dl_dst:0a:58:0a:f4:07:01,LOCAL",match="ip,nw_dst=10.244.7.1",priority="90",table="0"} 0
flowBytes{action="resubmit(,2)",match="ip,nw_src=10.244.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 1290
flowBytes{action="resubmit(,2)",match="ip,nw_src=172.16.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 0
flowBytes{action="resubmit(,3)",match="ip,nw_src=10.244.0.0/16,nw_dst=10.96.0.0/12",priority="33",table="0"} 11136
flowBytes{action="resubmit(,4)",match="ip",priority="1",table="2"} 0
flowBytes{action="resubmit(,4)",match="ip,nw_dst=10.244.0.0/16",priority="30",table="0"} 0
flowBytes{action="resubmit(,5)",match="in_port=1",priority="50",table="0"} 0
# HELP flowIdleTime The number of seconds have passed since the last packet has seen for the given OpenFlow entry
# TYPE flowIdleTime gauge
flowIdleTime{action="LOCAL",match="ip",priority="10",table="0"} 3
flowIdleTime{action="LOCAL",match="ip,nw_dst=10.96.0.1",priority="70",table="0"} 0
flowIdleTime{action="NORMAL",match="arp",priority="41000",table="0"} 588
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.140.235,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.13,tp_src=80",priority="100",table="2"} 63
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.235.60,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.12,tp_src=80",priority="100",table="2"} 69
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.17,tp_src=8079",priority="100",table="2"} 38
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.18,tp_src=8079",priority="100",table="2"} 36
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.7,tp_src=8079",priority="100",table="2"} 92
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.104.140.235,mod_tp_src:5672,resubmit(,4)",match="tcp,nw_src=10.244.7.14,tp_src=5672",priority="100",table="2"} 44
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.105.230.145,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.11,tp_src=80",priority="100",table="2"} 73
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.108.104.82,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.8,tp_src=80",priority="100",table="2"} 82
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.119.135,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.10,tp_src=80",priority="100",table="2"} 76
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.186.133,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.4,tp_src=27017",priority="100",table="2"} 103
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.110.235.87,mod_tp_src:3306,resubmit(,4)",match="tcp,nw_src=10.244.7.5,tp_src=3306",priority="100",table="2"} 97
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.111.233.21,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.9,tp_src=27017",priority="100",table="2"} 81
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 588
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 588
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 588
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 49
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.97.137.63,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.15,tp_src=80",priority="100",table="2"} 37
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.153.149,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.6,tp_src=80",priority="100",table="2"} 86
flowIdleTime{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.168.141,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.16,tp_src=27017",priority="100",table="2"} 6
flowIdleTime{action="mod_dl_dst:0a:58:0a:f4:07:01,LOCAL",match="ip,nw_dst=10.244.7.1",priority="90",table="0"} 588
flowIdleTime{action="resubmit(,2)",match="ip,nw_src=10.244.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 49
flowIdleTime{action="resubmit(,2)",match="ip,nw_src=172.16.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 588
flowIdleTime{action="resubmit(,3)",match="ip,nw_src=10.244.0.0/16,nw_dst=10.96.0.0/12",priority="33",table="0"} 0
flowIdleTime{action="resubmit(,4)",match="ip",priority="1",table="2"} 588
flowIdleTime{action="resubmit(,4)",match="ip,nw_dst=10.244.0.0/16",priority="30",table="0"} 588
flowIdleTime{action="resubmit(,5)",match="in_port=1",priority="50",table="0"} 588
# HELP flowPackets The number of packets matched for the given OpenFlow entry.
# TYPE flowPackets counter
flowPackets{action="LOCAL",match="ip",priority="10",table="0"} 460
flowPackets{action="LOCAL",match="ip,nw_dst=10.96.0.1",priority="70",table="0"} 630
flowPackets{action="NORMAL",match="arp",priority="41000",table="0"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.140.235,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.13,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.101.235.60,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.12,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.17,tp_src=8079",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.18,tp_src=8079",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.103.85.83,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.7,tp_src=8079",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.104.140.235,mod_tp_src:5672,resubmit(,4)",match="tcp,nw_src=10.244.7.14,tp_src=5672",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.105.230.145,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.11,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.108.104.82,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.8,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.119.135,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.10,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.109.186.133,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.4,tp_src=27017",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.110.235.87,mod_tp_src:3306,resubmit(,4)",match="tcp,nw_src=10.244.7.5,tp_src=3306",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.111.233.21,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.9,tp_src=27017",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="tcp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.2,tp_src=53",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.96.0.10,mod_tp_src:53,resubmit(,4)",match="udp,nw_src=10.244.7.3,tp_src=53",priority="100",table="2"} 8
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.97.137.63,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.15,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.153.149,mod_tp_src:80,resubmit(,4)",match="tcp,nw_src=10.244.7.6,tp_src=80",priority="100",table="2"} 0
flowPackets{action="load:0xaf4->NXM_OF_IP_DST[16..31],mod_nw_src:10.98.168.141,mod_tp_src:27017,resubmit(,4)",match="tcp,nw_src=10.244.7.16,tp_src=27017",priority="100",table="2"} 0
flowPackets{action="mod_dl_dst:0a:58:0a:f4:07:01,LOCAL",match="ip,nw_dst=10.244.7.1",priority="90",table="0"} 0
flowPackets{action="resubmit(,2)",match="ip,nw_src=10.244.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 8
flowPackets{action="resubmit(,2)",match="ip,nw_src=172.16.0.0/16,nw_dst=192.168.0.0/16",priority="33",table="0"} 0
flowPackets{action="resubmit(,3)",match="ip,nw_src=10.244.0.0/16,nw_dst=10.96.0.0/12",priority="33",table="0"} 118
flowPackets{action="resubmit(,4)",match="ip",priority="1",table="2"} 0
flowPackets{action="resubmit(,4)",match="ip,nw_dst=10.244.0.0/16",priority="30",table="0"} 0
flowPackets{action="resubmit(,5)",match="in_port=1",priority="50",table="0"} 0
# HELP groupBucketBytes The number of bytes that was sent by a given group bucket
# TYPE groupBucketBytes counter
groupBucketBytes{bucketActions="set_field:10.244.7.10->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1003",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.11->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1010",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.12->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1000",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.13->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1014",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.14->ip_dst,set_field:5672->tcp_dst,resubmit(,4)
",groupId="1001",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.15->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1005",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.16->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1004",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.17->ip_dst,set_field:8079->tcp_dst,resubmit(,4)
",groupId="1011",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.18->ip_dst,set_field:8079->tcp_dst,resubmit(,4)",groupId="1011",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.2->ip_dst,set_field:53->tcp_dst,resubmit(,4)
",groupId="1008",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.2->ip_dst,set_field:53->udp_dst,resubmit(,4)
",groupId="1007",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.3->ip_dst,set_field:53->tcp_dst,resubmit(,4)",groupId="1008",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.3->ip_dst,set_field:53->udp_dst,resubmit(,4)",groupId="1007",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.4->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1002",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.5->ip_dst,set_field:3306->tcp_dst,resubmit(,4)
",groupId="1012",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.6->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1009",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.7->ip_dst,set_field:8079->tcp_dst,resubmit(,4)",groupId="1011",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.8->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1006",groupType="select"} 0
groupBucketBytes{bucketActions="set_field:10.244.7.9->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1013",groupType="select"} 0
# HELP groupBucketPackets The number of packet that was sent by a given group bucket
# TYPE groupBucketPackets counter
groupBucketPackets{bucketActions="set_field:10.244.7.10->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1003",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.11->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1010",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.12->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1000",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.13->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1014",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.14->ip_dst,set_field:5672->tcp_dst,resubmit(,4)
",groupId="1001",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.15->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1005",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.16->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1004",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.17->ip_dst,set_field:8079->tcp_dst,resubmit(,4)
",groupId="1011",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.18->ip_dst,set_field:8079->tcp_dst,resubmit(,4)",groupId="1011",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.2->ip_dst,set_field:53->tcp_dst,resubmit(,4)
",groupId="1008",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.2->ip_dst,set_field:53->udp_dst,resubmit(,4)
",groupId="1007",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.3->ip_dst,set_field:53->tcp_dst,resubmit(,4)",groupId="1008",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.3->ip_dst,set_field:53->udp_dst,resubmit(,4)",groupId="1007",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.4->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1002",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.5->ip_dst,set_field:3306->tcp_dst,resubmit(,4)
",groupId="1012",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.6->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1009",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.7->ip_dst,set_field:8079->tcp_dst,resubmit(,4)",groupId="1011",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.8->ip_dst,set_field:80->tcp_dst,resubmit(,4)
",groupId="1006",groupType="select"} 0
groupBucketPackets{bucketActions="set_field:10.244.7.9->ip_dst,set_field:27017->tcp_dst,resubmit(,4)
",groupId="1013",groupType="select"} 0
# HELP groupBytes The number of bytes that was sent by a given group
# TYPE groupBytes counter
groupBytes{groupId="1000",groupType="select"} 0
groupBytes{groupId="1001",groupType="select"} 0
groupBytes{groupId="1002",groupType="select"} 0
groupBytes{groupId="1003",groupType="select"} 0
groupBytes{groupId="1004",groupType="select"} 0
groupBytes{groupId="1005",groupType="select"} 0
groupBytes{groupId="1006",groupType="select"} 0
groupBytes{groupId="1007",groupType="select"} 0
groupBytes{groupId="1008",groupType="select"} 0
groupBytes{groupId="1009",groupType="select"} 0
groupBytes{groupId="1010",groupType="select"} 0
groupBytes{groupId="1011",groupType="select"} 0
groupBytes{groupId="1012",groupType="select"} 0
groupBytes{groupId="1013",groupType="select"} 0
groupBytes{groupId="1014",groupType="select"} 0
# HELP groupDuration The number of seconds passed since the group entry was added
# TYPE groupDuration counter
groupDuration{groupId="1000",groupType="select"} 0
groupDuration{groupId="1001",groupType="select"} 0
groupDuration{groupId="1002",groupType="select"} 0
groupDuration{groupId="1003",groupType="select"} 0
groupDuration{groupId="1004",groupType="select"} 0
groupDuration{groupId="1005",groupType="select"} 0
groupDuration{groupId="1006",groupType="select"} 0
groupDuration{groupId="1007",groupType="select"} 0
groupDuration{groupId="1008",groupType="select"} 0
groupDuration{groupId="1009",groupType="select"} 0
groupDuration{groupId="1010",groupType="select"} 0
groupDuration{groupId="1011",groupType="select"} 0
groupDuration{groupId="1012",groupType="select"} 0
groupDuration{groupId="1013",groupType="select"} 0
groupDuration{groupId="1014",groupType="select"} 0
# HELP groupPackets The number of packet that was sent by a given group
# TYPE groupPackets counter
groupPackets{groupId="1000",groupType="select"} 0
groupPackets{groupId="1001",groupType="select"} 0
groupPackets{groupId="1002",groupType="select"} 0
groupPackets{groupId="1003",groupType="select"} 0
groupPackets{groupId="1004",groupType="select"} 0
groupPackets{groupId="1005",groupType="select"} 0
groupPackets{groupId="1006",groupType="select"} 0
groupPackets{groupId="1007",groupType="select"} 0
groupPackets{groupId="1008",groupType="select"} 0
groupPackets{groupId="1009",groupType="select"} 0
groupPackets{groupId="1010",groupType="select"} 0
groupPackets{groupId="1011",groupType="select"} 0
groupPackets{groupId="1012",groupType="select"} 0
groupPackets{groupId="1013",groupType="select"} 0
groupPackets{groupId="1014",groupType="select"} 0
# HELP portRxBytes The number of bytes that was recieved by a given port
# TYPE portRxBytes counter
portRxBytes{portNumber="1"} 0
portRxBytes{portNumber="10"} 760
portRxBytes{portNumber="11"} 0
portRxBytes{portNumber="12"} 752
portRxBytes{portNumber="13"} 3162
portRxBytes{portNumber="14"} 0
portRxBytes{portNumber="15"} 0
portRxBytes{portNumber="16"} 0
portRxBytes{portNumber="17"} 1536
portRxBytes{portNumber="18"} 1536
portRxBytes{portNumber="2"} 49290
portRxBytes{portNumber="3"} 50968
portRxBytes{portNumber="4"} 0
portRxBytes{portNumber="5"} 0
portRxBytes{portNumber="6"} 1128
portRxBytes{portNumber="7"} 2112
portRxBytes{portNumber="8"} 2416
portRxBytes{portNumber="9"} 0
portRxBytes{portNumber="LOCAL"} 0
# HELP portRxDrops The number of packets that was dropped on receive side by a given port
# TYPE portRxDrops counter
portRxDrops{portNumber="1"} 0
portRxDrops{portNumber="10"} 0
portRxDrops{portNumber="11"} 0
portRxDrops{portNumber="12"} 0
portRxDrops{portNumber="13"} 0
portRxDrops{portNumber="14"} 0
portRxDrops{portNumber="15"} 0
portRxDrops{portNumber="16"} 0
portRxDrops{portNumber="17"} 0
portRxDrops{portNumber="18"} 0
portRxDrops{portNumber="2"} 0
portRxDrops{portNumber="3"} 0
portRxDrops{portNumber="4"} 0
portRxDrops{portNumber="5"} 0
portRxDrops{portNumber="6"} 0
portRxDrops{portNumber="7"} 0
portRxDrops{portNumber="8"} 0
portRxDrops{portNumber="9"} 0
portRxDrops{portNumber="LOCAL"} 1157
# HELP portRxPackets The number of packet that was recieved by a given port
# TYPE portRxPackets counter
portRxPackets{portNumber="1"} 0
portRxPackets{portNumber="10"} 8
portRxPackets{portNumber="11"} 0
portRxPackets{portNumber="12"} 8
portRxPackets{portNumber="13"} 34
portRxPackets{portNumber="14"} 0
portRxPackets{portNumber="15"} 0
portRxPackets{portNumber="16"} 0
portRxPackets{portNumber="17"} 16
portRxPackets{portNumber="18"} 16
portRxPackets{portNumber="2"} 585
portRxPackets{portNumber="3"} 595
portRxPackets{portNumber="4"} 0
portRxPackets{portNumber="5"} 0
portRxPackets{portNumber="6"} 12
portRxPackets{portNumber="7"} 22
portRxPackets{portNumber="8"} 26
portRxPackets{portNumber="9"} 0
portRxPackets{portNumber="LOCAL"} 0
# HELP portTxBytes The number of bytes that was sent by a given port
# TYPE portTxBytes counter
portTxBytes{portNumber="1"} 0
portTxBytes{portNumber="10"} 866
portTxBytes{portNumber="11"} 796
portTxBytes{portNumber="12"} 866
portTxBytes{portNumber="13"} 866
portTxBytes{portNumber="14"} 866
portTxBytes{portNumber="15"} 866
portTxBytes{portNumber="16"} 866
portTxBytes{portNumber="17"} 726
portTxBytes{portNumber="18"} 726
portTxBytes{portNumber="2"} 3403
portTxBytes{portNumber="3"} 4135
portTxBytes{portNumber="4"} 866
portTxBytes{portNumber="5"} 866
portTxBytes{portNumber="6"} 866
portTxBytes{portNumber="7"} 866
portTxBytes{portNumber="8"} 2156
portTxBytes{portNumber="9"} 866
portTxBytes{portNumber="LOCAL"} 0
# HELP portTxDrops The number of packets that was dropped on sending side by a given port
# TYPE portTxDrops counter
portTxDrops{portNumber="1"} 0
portTxDrops{portNumber="10"} 0
portTxDrops{portNumber="11"} 0
portTxDrops{portNumber="12"} 0
portTxDrops{portNumber="13"} 0
portTxDrops{portNumber="14"} 0
portTxDrops{portNumber="15"} 0
portTxDrops{portNumber="16"} 0
portTxDrops{portNumber="17"} 0
portTxDrops{portNumber="18"} 0
portTxDrops{portNumber="2"} 0
portTxDrops{portNumber="3"} 0
portTxDrops{portNumber="4"} 0
portTxDrops{portNumber="5"} 0
portTxDrops{portNumber="6"} 0
portTxDrops{portNumber="7"} 0
portTxDrops{portNumber="8"} 0
portTxDrops{portNumber="9"} 0
portTxDrops{portNumber="LOCAL"} 0
# HELP portTxPackets The number of packet that was sent by a given port
# TYPE portTxPackets counter
portTxPackets{portNumber="1"} 0
portTxPackets{portNumber="10"} 11
portTxPackets{portNumber="11"} 10
portTxPackets{portNumber="12"} 11
portTxPackets{portNumber="13"} 11
portTxPackets{portNumber="14"} 11
portTxPackets{portNumber="15"} 11
portTxPackets{portNumber="16"} 11
portTxPackets{portNumber="17"} 9
portTxPackets{portNumber="18"} 9
portTxPackets{portNumber="2"} 41
portTxPackets{portNumber="3"} 49
portTxPackets{portNumber="4"} 11
portTxPackets{portNumber="5"} 11
portTxPackets{portNumber="6"} 11
portTxPackets{portNumber="7"} 11
portTxPackets{portNumber="8"} 19
portTxPackets{portNumber="9"} 11
portTxPackets{portNumber="LOCAL"} 0
`

func TestGetMetricsHandoutput(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://127.0.0.1:8081/metrics?target=127.0.0.1", nil)
	handler(w, r)
	resp := w.Result()
	defer r.Body.Close()
	raw_body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(raw_body))
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	body := string(raw_body)
	if body != HAND_OUTPUT {
		t.Errorf("Handler output mismatch, assumed length: %d, got length %d", len(HANDLER_OUTPUT), len(body))
	}
}
