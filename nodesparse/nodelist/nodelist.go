package nodelist

var Idlist = []string{
	//Temperature 30s
	"ns=1;s=t|TT21104/PV.PV_Out#Value",
	"ns=1;s=t|TT21105/PV.PV_Out#Value",
	"ns=1;s=t|TT21110/PV.PV_Out#Value",
	"ns=1;s=t|TT21210/PV.PV_Out#Value",
	"ns=1;s=t|TT21300/PV.PV_Out#Value",
	"ns=1;s=t|TT21302/PV.PV_Out#Value",
	"ns=1;s=t|TT21404/PV.PV_Out#Value",
	//Pressure 5s
	"ns=1;s=t|PT21101/PV.PV_Out#Value",
	"ns=1;s=t|PT21106/PV.PV_Out#Value",
	"ns=1;s=t|PT21200/PV.PV_Out#Value",
	"ns=1;s=t|PT21205/PV.PV_Out#Value",
	"ns=1;s=t|PT21212/PV.PV_Out#Value",
	"ns=1;s=t|PT21301/PV.PV_Out#Value",
	"ns=1;s=t|PT21400/PV.PV_Out#Value",
	"ns=1;s=t|PT21405/PV.PV_Out#Value",
	//Level 10s
	"ns=1;s=t|LT21107/PV.PV_Out#Value",
	"ns=1;s=t|LT21108/PV.PV_Out#Value",
	"ns=1;s=t|LT21221/PV.PV_Out#Value",
	"ns=1;s=t|LT21416/PV.PV_Out#Value",
	"ns=1;s=t|LIC21416/PID.PV_Out#Value",
	"ns=1;s=t|LIC21416/PID.SP#Value",
	"ns=1;s=t|LIC21416/PID.MV#Value",
	//Flow 1s
	"ns=1;s=t|FIC21208/PID.PV_Out#Value",
	"ns=1;s=t|FIC21208/PID.SP#Value",
	"ns=1;s=t|FIC21208/PID.MV#Value",
	"ns=1;s=t|FIC21407/PID.PV_Out#Value",
	"ns=1;s=t|FIC21407/PID.SP#Value",
	"ns=1;s=t|FIC21407/PID.MV#Value",
	"ns=1;s=t|FIC21410/PID.PV_Out#Value",
	"ns=1;s=t|FIC21410/PID.SP#Value",
	"ns=1;s=t|FIC21410/PID.MV#Value",
}

var Namelist = []string{
	//1 Temperature 30s
	"DEC21-TT-104.PV",
	"DEC21-TT-105.PV",
	"DEC21-TT-110.PV",
	"DEC21-TT-210.PV",
	"DEC21-TT-300.PV",
	"DEC21-TT-302.PV",
	"DEC21-TT-404.PV",
	//2 Pressure 5s
	"DEC21-PT-101.PV",
	"DEC21-PT-106.PV",
	"DEC21-PT-200.PV",
	"DEC21-PT-205.PV",
	"DEC21-PT-212.PV",
	"DEC21-PT-301.PV",
	"DEC21-PT-400.PV",
	"DEC21-PT-405.PV",
	//3 Level 10s
	"DEC21-LT-107.PV",
	"DEC21-LT-108.PV",
	"DEC21-LT-221.PV",
	"DEC21-LT-416.PV",
	"DEC21-LIC-416.PV",
	"DEC21-LIC-416.SP",
	"DEC21-LIC-416.MV",
	//4 Flow PID 1s
	"DEC21-FIC-208.PV",
	"DEC21-FIC-208.SP",
	"DEC21-FIC-208.MV",
	"DEC21-FIC-407.PV",
	"DEC21-FIC-407.SP",
	"DEC21-FIC-407.MV",
	"DEC21-FIC-410.PV",
	"DEC21-FIC-410.SP",
	"DEC21-FIC-410.MV",
}

var Intervallist = []int{
	//1 Temperature 30s
	30, 30, 30, 30, 30, 30, 30,
	//2 Pressure 5s
	5, 5, 5, 5, 5, 5, 5, 5,
	//3 Level 10s
	10, 10, 10, 10, 10, 10, 10,
	//4 Flow PID 1s
	1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var Stablelist = []string{
	//1 Temperature 30s
	"DEC21.Temperature", "DEC21.Temperature", "DEC21.Temperature", "DEC21.Temperature", "DEC21.Temperature", "DEC21.Temperature", "DEC21.Temperature",
	//2 Pressure 5s
	"DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure", "DEC21.Pressure",
	//3 Level 10s
	"DEC21.Level", "DEC21.Level", "DEC21.Level", "DEC21.Level", "DEC21.Level", "DEC21.Level", "DEC21.Level",
	//4 Flow PID 1s
	"DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow", "DEC21.Flow",
}

var Subtablelist = []string{
	//1 Temperature 30s
	"DEC21.t1001", "DEC21.t1002", "DEC21.t1003", "DEC21.t1004", "DEC21.t1005", "DEC21.t1006", "DEC21.t1007",
	//2 Pressure 5s
	"DEC21.p1001", "DEC21.p1002", "DEC21.p1003", "DEC21.p1004", "DEC21.p1005", "DEC21.p1006", "DEC21.p1007", "DEC21.p1008",
	//3 Level 10s
	"DEC21.l1001", "DEC21.l1002", "DEC21.l1003", "DEC21.l1004", "DEC21.l1005", "DEC21.l1006", "DEC21.l1007",
	//4 Flow PID 1s
	"DEC21.f1001", "DEC21.f1002", "DEC21.f1003", "DEC21.f1004", "DEC21.f1005", "DEC21.f1006", "DEC21.f1007", "DEC21.f1008", "DEC21.f1009",
}
