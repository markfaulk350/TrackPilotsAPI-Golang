package entity

type JsonResponse struct {
	Success bool   `json:"success"`
	Payload string `json:"payload"`
}

type User struct {
	ID               int    `json:"id"`
	Fname            string `json:"fName"`
	Lname            string `json:"lName"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Country          string `json:"country"`
	Trklink          string `json:"trkLink"`
	Trktype          string `json:"trkType"`
	GliderBrand      string `json:"gldBrand"`
	GliderMake       string `json:"gldMake"`
	GliderColor      string `json:"gldColor"`
	LastLocationPing int64  `json:"lastLocationPing"`
	LastApiCall      int64  `json:"lastApiCall"`
	Created          string `json:"created"`
}

type Group struct {
	ID        int    `json:"id"`
	Groupname string `json:"groupName"`
	Creatorid string `json:"creatorId"`
	Region    string `json:"region"`
	Info      string `json:"info"`
	Radio     string `json:"radioFrq"`
	Created   string `json:"created"`
}

type Roster struct {
	ID      int `json:"id"`
	Groupid int `json:"group_id"`
	Pilotid int `json:"pilot_id"`
}

type UserAndPings struct {
	User
	Pings []Ping `json:"pings"`
}

type Ping struct {
	ID          int     `json:"id"`
	UnixTime    int64   `json:"unixTime"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Alt         float64 `json:"alt"`
	AGL         string  `json:"agl"`
	Velocity    string  `json:"velocity"`
	Heading     string  `json:"heading"`
	TxtMsg      string  `json:"txtMsg"`
	IsEmergency string  `json:"isEmergency"`
}
