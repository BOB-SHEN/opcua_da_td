package nodesparse

import (
	"time"

	"opcua_da_td/nodesparse/nodelist"
	"opcua_da_td/nodesparse/nodevfdlist"

	"github.com/gopcua/opcua/ua"
)

type Node struct {
	ID       *ua.NodeID
	Interval time.Duration
	Ticker   *time.Ticker
	Name     string
	Stable   string
	Subtable string
}

var Nodes []Node

func InitNodes() {
	for i := 0; i < len(nodelist.Idlist); i++ {
		n := Node{
			ID:       ua.MustParseNodeID(nodelist.Idlist[i]),
			Interval: time.Duration(nodelist.Intervallist[i]) * time.Second,
			Name:     nodelist.Namelist[i],
			Stable:   nodelist.Stablelist[i],
			Subtable: nodelist.Subtablelist[i],
		}
		Nodes = append(Nodes, n)
	}
}

type NodeVfd struct {
	ID_S     *ua.NodeID
	ID_T     *ua.NodeID
	ID_P     *ua.NodeID
	ID_C     *ua.NodeID
	Interval time.Duration
	Ticker   *time.Ticker
	Name     string
	Stable   string
	Subtable string
}

var NodesVfd []NodeVfd

func InitNodesVfd() {
	for i := 0; i < len(nodevfdlist.NameVfdlist); i++ {
		n := NodeVfd{
			ID_S:     ua.MustParseNodeID(nodevfdlist.IdVfdSlist[i]),
			ID_T:     ua.MustParseNodeID(nodevfdlist.IdVfdTlist[i]),
			ID_P:     ua.MustParseNodeID(nodevfdlist.IdVfdPlist[i]),
			ID_C:     ua.MustParseNodeID(nodevfdlist.IdVfdClist[i]),
			Interval: time.Duration(nodevfdlist.IntervalVfdlist[i]) * time.Second,
			Name:     nodevfdlist.NameVfdlist[i],
			Stable:   nodevfdlist.StableVfdlist[i],
			Subtable: nodevfdlist.SubtableVfdlist[i],
		}
		NodesVfd = append(NodesVfd, n)
	}
}

// {
// 	//------------------Temperature------------------
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21104/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-104.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21105/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-105.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1002",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21110/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-110.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1003",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21210/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-210.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1004",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21300/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-300.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1005",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21302/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-302.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1006",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|TT21404/PV.PV_Out#Value"),
// 		Interval: 30 * time.Second,
// 		Name:     "DEC21-TT-404.PV",
// 		Stable:   "DEC21.Temperature",
// 		Sbutable: "DEC21.t1007",
// 	},
// 	//------------------Pressure------------------
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// 	{
// 		ID:       ua.MustParseNodeID("ns=1;s=t|PT21101/PV.PV_Out#Value"),
// 		Interval: 5 * time.Second,
// 		Name:     "DEC21-PT-101",
// 		Stable:   "DEC21.Preasure",
// 		Sbutable: "DEC21.p1001",
// 	},
// }
