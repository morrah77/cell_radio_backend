package common

type CellNetworkParams struct {
	Arfcn int
	Ulsc int
	Dlsc int
}

type Config struct {
	RxGain int
	Network int
	A_2g_params CellNetworkParams `json:"2_g_params"`
	A_3g_params CellNetworkParams `json:"3_g_params"`
	A_4g_params CellNetworkParams `json:"4_g_params"`
	Alpha float64
	Slot string
	Celltrack int
	Watcher int
	Antenna int
	GpsSrc int
	Version int
	App int
	GsmMode int
	Url string
	Sound float64
}
