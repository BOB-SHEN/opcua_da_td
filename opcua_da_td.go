package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"database/sql"
	"time"

	//	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosSql"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/debug"

	//"github.com/gopcua/opcua/mycode/nodesparse"
	//"github.com/BOB-SHEN/opcua_td/opcua_da_td/nodesparse"
	"opcua_da_td/nodesparse"

	//"github.com/gopcua/opcua/mycode/utils"
	//"github.com/BOB-SHEN/opcua_td/opcua_da_td/utils"
	"opcua_da_td/utils"

	"github.com/gopcua/opcua/ua"
)

func main() {

	//-----------------------------------ConnectTaosdb----------------------------------
	var taosUri = "root:taosdata@http(DC21TD01:6030)/"
	taos, err := sql.Open("taosSql", taosUri)
	if err != nil {
		log.Fatalln("failed to connect TDengine, err:", err)
	}
	defer taos.Close()

	//-----------------------------------CreateStable----------------------------------
	utils.CreateDatabase(taos)
	utils.CreateStableT(taos)
	utils.CreateStableP(taos)
	utils.CreateStableL(taos)
	utils.CreateStableF(taos)
	utils.CreateStableVFD(taos)

	//-----------------------------------ConnectOpcuaServer----------------------------------
	var (
		endpoint = flag.String("endpoint", "opc.tcp://192.168.21.4:4862", "OPC UA Endpoint URL")
	)
	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
	flag.Parse()
	log.SetFlags(0)

	ctx := context.Background()

	c := opcua.NewClient(*endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.CloseWithContext(ctx)

	nodesparse.InitNodes()
	nodesparse.InitNodesVfd()

	//-----------------------------------goroutine loop----------------------------------
	for i := range nodesparse.Nodes {
		j := i
		node := &nodesparse.Nodes[i]
		node.Ticker = time.NewTicker(node.Interval)
		go func() {
			for range node.Ticker.C {

				if j >= 7 && j < 15 {
					time.Sleep(2 * time.Second)
				}

				if j >= 15 && j < 22 {
					time.Sleep(4 * time.Second)
				}

				timeStr := time.Now().Format("2006-01-02 15:04:05")

				req := &ua.ReadRequest{
					MaxAge: 2000,
					NodesToRead: []*ua.ReadValueID{
						{NodeID: node.ID},
					},
					TimestampsToReturn: ua.TimestampsToReturnBoth,
				}

				//------------------ReadOPCUA------------------
				resp, err := c.ReadWithContext(ctx, req)
				if err != nil {
					log.Fatalf("Read failed: %s", err)
				}
				if resp.Results[0].Status != ua.StatusOK {
					log.Fatalf("Status not OK: %v", resp.Results[0].Status)
				}
				// log.Printf("%#v", resp.Results[0].Value.Value())

				//------------------WriteTaosdb------------------
				intermevar := fmt.Sprint(resp.Results[0].Value.Value())
				sql := "INSERT INTO " + node.Subtable + " USING " + node.Stable + " TAGS('China.Shanghai', 4) VALUES ('" + timeStr + "','" + node.Name + "'," + intermevar + ")"
				result, err := taos.Exec(sql)
				if err != nil {
					log.Fatalln("failed to insert, err:", err)
				}

				_ = result

				// rowsAffected, err := result.RowsAffected()
				// if err != nil {
				// 	log.Fatalln("failed to get affected rows, err:", err)
				// }
				// fmt.Println("RowsAffected", rowsAffected)

			}

		}()

	}

	//-----------------------------------goroutine loop----------------------------------
	for i := range nodesparse.NodesVfd {
		node := &nodesparse.NodesVfd[i]
		node.Ticker = time.NewTicker(node.Interval)
		go func() {
			for range node.Ticker.C {

				timeStr := time.Now().Format("2006-01-02 15:04:05")

				req_s := &ua.ReadRequest{
					MaxAge: 2000,
					NodesToRead: []*ua.ReadValueID{
						{NodeID: node.ID_S},
					},
					TimestampsToReturn: ua.TimestampsToReturnBoth,
				}

				req_t := &ua.ReadRequest{
					MaxAge: 2000,
					NodesToRead: []*ua.ReadValueID{
						{NodeID: node.ID_T},
					},
					TimestampsToReturn: ua.TimestampsToReturnBoth,
				}

				req_p := &ua.ReadRequest{
					MaxAge: 2000,
					NodesToRead: []*ua.ReadValueID{
						{NodeID: node.ID_P},
					},
					TimestampsToReturn: ua.TimestampsToReturnBoth,
				}

				req_c := &ua.ReadRequest{
					MaxAge: 2000,
					NodesToRead: []*ua.ReadValueID{
						{NodeID: node.ID_C},
					},
					TimestampsToReturn: ua.TimestampsToReturnBoth,
				}

				//------------------ReadOPCUA------------------
				resp_s, err := c.ReadWithContext(ctx, req_s)
				if err != nil {
					log.Fatalf("Read failed: %s", err)
				}
				if resp_s.Results[0].Status != ua.StatusOK {
					log.Fatalf("Status not OK: %v", resp_s.Results[0].Status)
				}
				// log.Printf("%#v", resp_s.Results[0].Value.Value())

				resp_t, err := c.ReadWithContext(ctx, req_t)
				if err != nil {
					log.Fatalf("Read failed: %s", err)
				}
				if resp_t.Results[0].Status != ua.StatusOK {
					log.Fatalf("Status not OK: %v", resp_t.Results[0].Status)
				}
				// log.Printf("%#v", resp_t.Results[0].Value.Value())

				resp_p, err := c.ReadWithContext(ctx, req_p)
				if err != nil {
					log.Fatalf("Read failed: %s", err)
				}
				if resp_p.Results[0].Status != ua.StatusOK {
					log.Fatalf("Status not OK: %v", resp_p.Results[0].Status)
				}
				// log.Printf("%#v", resp_p.Results[0].Value.Value())

				resp_c, err := c.ReadWithContext(ctx, req_c)
				if err != nil {
					log.Fatalf("Read failed: %s", err)
				}
				if resp_c.Results[0].Status != ua.StatusOK {
					log.Fatalf("Status not OK: %v", resp_c.Results[0].Status)
				}
				// log.Printf("%#v", resp_c.Results[0].Value.Value())
				//------------------WriteTaosdb------------------
				intermevar_s := fmt.Sprint(resp_s.Results[0].Value.Value())
				intermevar_t := fmt.Sprint(resp_t.Results[0].Value.Value())
				intermevar_p := fmt.Sprint(resp_p.Results[0].Value.Value())
				intermevar_c := fmt.Sprint(resp_c.Results[0].Value.Value())

				sql := "INSERT INTO " + node.Subtable + " USING " + node.Stable + " TAGS('China.Shanghai', 4) VALUES ('" + timeStr + "','" + node.Name + "'," + intermevar_s + "," + intermevar_t + "," + intermevar_p + "," + intermevar_c + ")"
				result, err := taos.Exec(sql)
				if err != nil {
					log.Fatalln("failed to insert, err:", err)
				}

				_ = result

				// rowsAffected, err := result.RowsAffected()
				// if err != nil {
				// 	log.Fatalln("failed to get affected rows, err:", err)
				// }
				// fmt.Println("RowsAffected", rowsAffected)

			}

		}()

	}

	// time.Sleep(10000000 * time.Second)

	for {

	}
}
