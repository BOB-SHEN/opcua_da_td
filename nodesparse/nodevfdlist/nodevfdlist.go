package nodevfdlist

var IdVfdSlist = []string{
	"ns=1;s=t|J21200/MOT.Rbk#Value",
	"ns=1;s=t|J21201/MOT.Rbk#Value",
	"ns=1;s=t|J21400/MOT.Rbk#Value",
}

var IdVfdTlist = []string{
	"ns=1;s=t|J21200/MOT.UserAna1#Value",
	"ns=1;s=t|J21201/MOT.UserAna1#Value",
	"ns=1;s=t|J21400/MOT.UserAna1#Value",
}

var IdVfdPlist = []string{
	"ns=1;s=t|J21200/MOT.UserAna2#Value",
	"ns=1;s=t|J21201/MOT.UserAna2#Value",
	"ns=1;s=t|J21400/MOT.UserAna2#Value",
}

var IdVfdClist = []string{
	"ns=1;s=t|J21200/current_AV.PV_Out#Value",
	"ns=1;s=t|J21201/current_AV.PV_Out#Value",
	"ns=1;s=t|J21400/current_AV.PV_Out#Value",
}

var NameVfdlist = []string{
	"DEC21-J-200",
	"DEC21-J-201",
	"DEC21-J-400",
}

var IntervalVfdlist = []int{
	//VFD 1s
	1, 1, 1,
}

var StableVfdlist = []string{
	"DEC21.VFD", "DEC21.VFD", "DEC21.VFD",
}

var SubtableVfdlist = []string{
	"DEC21.v1001", "DEC21.v1002", "DEC21.v1003",
}
