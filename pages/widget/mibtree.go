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
	_ = &mibNode{
		name: "icmp",
		oid:  "5",
		children: []*mibNode{
			{},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
			{
				name: "",
				oid:  "",
			},
		},
	}
}
