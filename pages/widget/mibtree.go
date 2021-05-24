//MIT License
//
//Copyright (c) 2021 kiloson
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

package widget

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	MibTree *tview.TreeView
)

type mibNode struct {
	oid      string
	name     string
	parent   *mibNode
	children []*mibNode
}

func (n *mibNode) String() string {
	//return fmt.Sprintf("%s %s", n.oid, n.name)
	return fmt.Sprintf(" %s", n.name)
}

func init() {
	root := &mibNode{
		oid:      ".1.3.6.1.2.1",
		name:     "mib-2",
		children: []*mibNode{{}},
	}

	setSystem(root)
	setIf(root)
	setAt(root)
	setIp(root)
	setIcmp(root)
	setTcp(root)
	setUdp(root)
	setEgp(root)

	node := tview.NewTreeNode(root.String()).SetReference(root.oid).SetColor(tcell.ColorGreen)
	var addNodes func(*tview.TreeNode, *mibNode, string)
	addNodes = func(node *tview.TreeNode, mib *mibNode, ref string) {
		if mib.children == nil {
			return
		}
		for _, n := range mib.children[1:] {
			ref := fmt.Sprintf("%s.%s", ref, n.oid)
			t := tview.NewTreeNode(n.String()).SetReference(ref).SetExpanded(false)
			if len(n.children) > 1 {
				t.SetColor(tcell.ColorGreen)
			}
			addNodes(t, n, ref)
			node.AddChild(t)
		}
	}

	addNodes(node, root, root.oid)
	MibTree = tview.NewTreeView()
	MibTree.SetRoot(node).SetCurrentNode(node)
}

func setSystem(root *mibNode) {
	system := &mibNode{
		oid:    "1",
		name:   "system",
		parent: root,
	}
	children := []*mibNode{
		{},
		{
			oid:    "1",
			name:   "sysDescr",
			parent: system,
		},
		{
			oid:    "2",
			name:   "sysObjectID",
			parent: system,
		},
		{
			oid:    "3",
			name:   "sysUpTime",
			parent: system,
		},
		{
			oid:    "4",
			name:   "sysContact",
			parent: system,
		},
		{
			oid:    "5",
			name:   "sysName",
			parent: system,
		},
		{
			oid:    "6",
			name:   "sysLocation",
			parent: system,
		},
		{
			oid:    "7",
			name:   "sysServices",
			parent: system,
		},
	}
	system.children = children
	root.children = append(root.children, system)
}

func setIf(root *mibNode) {
	interfaces := &mibNode{
		name: "interfaces",
		oid:  "2",
		children: []*mibNode{
			{},
			{
				name: "ifNumber",
				oid:  "1",
			},
			{
				name: "ifTable",
				oid:  "2",
				children: []*mibNode{
					{},
					{
						name: "ifEntry",
						oid:  "1",
						children: []*mibNode{
							{},
							{
								name: "ifIndex",
								oid:  "1",
							},
							{
								name: "ifDescr",
								oid:  "2",
							},
							{
								name: "ifType",
								oid:  "3",
							},
							{
								name: "ifMtu",
								oid:  "4",
							},
							{
								name: "ifSpeed",
								oid:  "5",
							},
							{
								name: "ifPhysAddress",
								oid:  "6",
							},
							{
								name: "ifAdminStatus",
								oid:  "7",
							},
							{
								name: "ifOperStatus",
								oid:  "8",
							},
							{
								name: "ifLastChange",
								oid:  "9",
							},
							{
								name: "ifInOctets",
								oid:  "10",
							},
							{
								name: "ifInUcastPkts",
								oid:  "11",
							},
							{
								name: "ifInNUcastPkts",
								oid:  "12",
							},
							{
								name: "ifInDiscards",
								oid:  "13",
							},
							{
								name: "ifInErrors",
								oid:  "14",
							},
							{
								name: "ifInUnknownPorotos",
								oid:  "15",
							},
							{
								name: "ifOutOctets",
								oid:  "16",
							},
							{
								name: "ifOutUcastPkts",
								oid:  "17",
							},
							{
								name: "ifOutUcastPkts",
								oid:  "18",
							},
							{
								name: "ifOutDiscards",
								oid:  "19",
							},
							{
								name: "ifOutError",
								oid:  "20",
							},
							{
								name: "ifOutQLen",
								oid:  "21",
							},
							{
								name: "ifSpecfic",
								oid:  "22",
							},
						},
					},
				},
			},
		},
	}
	root.children = append(root.children, interfaces)
}

func setAt(root *mibNode) {
	at := &mibNode{
		name: "at",
		oid:  "3",
		children: []*mibNode{
			{},
			{
				name: "atTable",
				oid:  "1",
				children: []*mibNode{
					{},
					{
						name: "atEntry",
						oid:  "1",
						children: []*mibNode{
							{},
							{
								name: "atIfIndex",
								oid:  "1",
							},
							{
								name: "atPhysAddress",
								oid:  "2",
							},
							{
								name: "atNetAddress",
								oid:  "3",
							},
						},
					},
				},
			},
		},
	}
	root.children = append(root.children, at)
}

func setIp(root *mibNode) {
	ip := &mibNode{
		name: "ip",
		oid:  "4",
		children: []*mibNode{
			{},
			{
				name: "ipForwarding",
				oid:  "1",
			},
			{
				name: "ipDefaultTTL",
				oid:  "2",
			},
			{
				name: "ipInReceives",
				oid:  "3",
			},
			{
				name: "ipInHdrErrors",
				oid:  "4",
			},
			{
				name: "ipInAddrErrors",
				oid:  "5",
			},
			{
				name: "ipForwDatagrams",
				oid:  "6",
			},
			{
				name: "ipInUnknownProtos",
				oid:  "7",
			},
			{
				name: "ipInDiscards",
				oid:  "8",
			},
			{
				name: "ipInDelivers",
				oid:  "9",
			},
			{
				name: "ipOutRequests",
				oid:  "10",
			},
			{
				name: "ipOutDircards",
				oid:  "11",
			},
			{
				name: "ipOutNoRoutes",
				oid:  "12",
			},
			{
				name: "ipOutReasmTimeout",
				oid:  "13",
			},
			{
				name: "ipOutReasmReqds",
				oid:  "14",
			},
			{
				name: "ipOutReasmOKs",
				oid:  "15",
			},
			{
				name: "ipOutReasmFails",
				oid:  "16",
			},
			{
				name: "ipFragOKs",
				oid:  "17",
			},
			{
				name: "ipFragFails",
				oid:  "18",
			},
			{
				name: "ipFragCreates",
				oid:  "19",
			},
			{
				name: "ipAddrTable",
				oid:  "20",
				children: []*mibNode{
					{},
					{
						name: "ipAddrEntry",
						oid:  "1",
						children: []*mibNode{
							{},
							{
								name: "ipAdEtnAddr",
								oid:  "1",
							},
							{
								name: "ipAdEtnIfIndex",
								oid:  "2",
							},
							{
								name: "ipAdEtnNetMask",
								oid:  "3",
							},
							{
								name: "ipAdEtnBcastAddr",
								oid:  "4",
							},
							{
								name: "ipAdEtnResamMaskSize",
								oid:  "5",
							},
						},
					},
				},
			},
			{
				name: "ipRouteTable",
				oid:  "21",
				children: []*mibNode{
					{},

					{
						name: "ipRouteEntry",
						oid:  "1",
						children: []*mibNode{
							{},
							{
								name: "ipRouteDest",
								oid:  "1",
							},
							{
								name: "ipRouteIfIndex",
								oid:  "2",
							},
							{
								name: "ipRouteMetric1",
								oid:  "3",
							},
							{
								name: "ipRouteMetric2",
								oid:  "4",
							},
							{
								name: "ipRouteMetric3",
								oid:  "5",
							},
							{
								name: "ipRouteMetric4",
								oid:  "6",
							},
							{
								name: "ipRouteNextHop",
								oid:  "7",
							},
							{
								name: "ipRouteType",
								oid:  "8",
							},
							{
								name: "ipRoutePoroto",
								oid:  "9",
							},
							{
								name: "ipRouteAge",
								oid:  "10",
							},
							{
								name: "ipRouteMask",
								oid:  "11",
							},
							{
								name: "ipRouteMetric5",
								oid:  "12",
							},
							{
								name: "ipRouteInfo",
								oid:  "13",
							},
						},
					},
				},
			},
			{
				name: "ipNetToMediaTable",
				oid:  "22",
				children: []*mibNode{
					{},
					{
						name: "ipNetToMediaEntry",
						oid:  "1",
						children: []*mibNode{
							{},
							{
								name: "ipNetToMediaIfIndex",
								oid:  "1",
							},
							{
								name: "ipNetToMediaPhysAddress",
								oid:  "2",
							},
							{
								name: "ipNetToMediaNetAddress",
								oid:  "3",
							},
							{
								name: "ipNetToMediaType",
								oid:  "4",
							},
						},
					},
				},
			},
			{
				name: "ipRoutingDiscards",
				oid:  "23",
			},
		},
	}
	root.children = append(root.children, ip)
}

func setIcmp(root *mibNode) {
	icmp := &mibNode{
		name: "icmp",
		oid:  "5",
		children: []*mibNode{
			{},
			{
				name: "icmpInMsgs",
				oid:  "1",
			},
			{
				name: "icmpInErrors",
				oid:  "2",
			},
			{
				name: "icmpInDestUnreachs",
				oid:  "3",
			},
			{
				name: "icnmpInTimeExcds",
				oid:  "4",
			},
			{
				name: "icmpInPramProbe",
				oid:  "5",
			},
			{
				name: "icmpInSrcQuenchs",
				oid:  "6",
			},
			{
				name: "icmpInRedirects",
				oid:  "7",
			},
			{
				name: "icmpInEchos",
				oid:  "8",
			},
			{
				name: "icmpInEchoReps",
				oid:  "9",
			},
			{
				name: "icmpInTimestamps",
				oid:  "10",
			},
			{
				name: "icmpInTimestampReps",
				oid:  "11",
			},
			{
				name: "icmpInAddrMasks",
				oid:  "12",
			},
			{
				name: "icmpInAddrMaskReps",
				oid:  "13",
			},
			{
				name: "icmpOutMsgs",
				oid:  "14",
			},
			{
				name: "icmpOutErrors",
				oid:  "15",
			},
			{
				name: "icmpOutDestUnreachs",
				oid:  "16",
			},
			{
				name: "icmpOutTimeExcds",
				oid:  "17",
			},
			{
				name: "icmpOutParamProbe",
				oid:  "18",
			},
			{
				name: "icmpOutSrcQuenchs",
				oid:  "19",
			},
			{
				name: "icmpOutRedirects",
				oid:  "20",
			},
			{
				name: "icmpOutEchos",
				oid:  "21",
			},
			{
				name: "icmpOutEchoReps",
				oid:  "22",
			},
			{
				name: "icmpOutTimestamps",
				oid:  "23",
			},
			{
				name: "icmpOutTimestampReps",
				oid:  "24",
			},
			{
				name: "icmpOutAddrMasks",
				oid:  "25",
			},
			{
				name: "icmpOutAddrMaskReps",
				oid:  "26",
			},
		},
	}
	root.children = append(root.children, icmp)
}

func setTcp(root *mibNode) {
	tcp := &mibNode{
		oid:  "6",
		name: "tcp",
		children: []*mibNode{
			{
				name: "tcpRtoAlgorithm",
				oid:  "1",
			},
			{
				name: "tcpRtoMin",
				oid:  "2",
			},
			{
				name: "tcpRtoMax",
				oid:  "3",
			},
			{
				name: "tcpMaxConn",
				oid:  "4",
			},
			{
				name: "tcpActiveOpens",
				oid:  "5",
			},
			{
				name: "tcpPassiveOpens",
				oid:  "6",
			},
			{
				name: "tcpAttemptFails",
				oid:  "7",
			},
			{
				name: "tcpEstabResets",
				oid:  "8",
			},
			{
				name: "tcpCurrEstab",
				oid:  "9",
			},
			{
				name: "tcpInSegs",
				oid:  "10",
			},
			{
				name: "tcpOutSegs",
				oid:  "11",
			},
			{
				name: "tcpRetrsnsSegs",
				oid:  "12",
			},
			{
				name: "tcpConnTable",
				oid:  "13",
				children: []*mibNode{
					{
						name: "tcpConnEntry",
						oid:  "1",
						children: []*mibNode{
							{
								oid:  "1",
								name: "tcpConnState",
							},
							{
								oid:  "2",
								name: "tcpConnLocalAddress",
							},
							{
								oid:  "3",
								name: "tcpConnLocalPort",
							},
							{
								oid:  "4",
								name: "tcpConnRemAddress",
							},
							{
								oid:  "5",
								name: "tcpConnRemPort",
							},
						},
					},
				},
			},
			{
				name: "tcpInErrors",
				oid:  "14",
			},
			{
				name: "tcpOutRests",
				oid:  "15",
			},
		},
	}
	root.children = append(root.children, tcp)
}

func setUdp(root *mibNode) {
	udp := &mibNode{
		name: "udp",
		oid:  "7",
		children: []*mibNode{
			{
				oid:  "1",
				name: "udpInDatagrams",
			},
			{
				oid:  "2",
				name: "udpNoPorts",
			},
			{
				oid:  "3",
				name: "udpInErrors",
			},
			{
				oid:  "4",
				name: "udpOutDatagrams",
			},
			{
				oid:  "5",
				name: "udpTable",
				children: []*mibNode{
					{
						name: "udpEntry",
						oid:  "1",
						children: []*mibNode{
							{
								oid:  "1",
								name: "udpLocalAddress",
							},
							{
								oid:  "2",
								name: "udpLocalPort",
							},
						},
					},
				},
			},
		},
	}
	root.children = append(root.children, udp)
}

func setEgp(root *mibNode) {
	egp := &mibNode{
		name: "egp",
		oid:  "8",
		children: []*mibNode{
			{
				oid:  "1",
				name: "egpInMegs",
			},
			{
				oid:  "2",
				name: "egpInErrors",
			},
			{
				oid:  "3",
				name: "egpOutMegs",
			},
			{
				oid:  "4",
				name: "egpOutErrors",
			},
			{
				oid:  "5",
				name: "egpNeighTable",
				children: []*mibNode{
					{
						name: "egpNeighEntry",
						oid:  "1",
						children: []*mibNode{
							{
								oid:  "1",
								name: "egpNeighState",
							},
							{
								oid:  "2",
								name: "egpNeighAddr",
							},
							{
								oid:  "3",
								name: "egpNeighAs",
							},
							{
								oid:  "4",
								name: "egpNeighMegs",
							},
							{
								oid:  "5",
								name: "egpNeighInErrs",
							},
							{
								oid:  "6",
								name: "egpNeighOutMegs",
							},
							{
								oid:  "7",
								name: "egpNeighOutErr",
							},
							{
								oid:  "8",
								name: "egpNeighInErrMegs",
							},
							{
								oid:  "9",
								name: "egpNeighOutErrMegs",
							},
							{
								oid:  "10",
								name: "egpNeighStateUps",
							},
							{
								oid:  "11",
								name: "egpNeighStateDowns",
							},
							{
								oid:  "12",
								name: "egpNeighIntervalHello",
							},
							{
								oid:  "13",
								name: "egpNeighIntervalPoll",
							},
							{
								oid:  "14",
								name: "egpNeighMode",
							},
							{
								oid:  "15",
								name: "egpNeighEventTrigger",
							},
						},
					},
				},
			},
			{
				oid:  "6",
				name: "egpAs",
			},
		},
	}
	root.children = append(root.children, egp)
}
