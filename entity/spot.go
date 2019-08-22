package entity

type SpotDataStructure struct {
	Response *Response `json:"response"`
}

type Response struct {
	FeedMessageResponse *FeedMessageResponse `json:"feedMessageResponse"`
	Errors              *Errors              `json:"errors"`
}

type FeedMessageResponse struct {
	Count         int       `json:"count"`
	TotalCount    int       `json:"totalCount"`
	ActivityCount int       `json:"activityCount"`
	Feed          *Feed     `json:"feed"`
	Messages      *Messages `json:"messages"`
}

type Feed struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Status               string `json:"status"`
	Usage                int    `json:"usage"`
	DaysRange            int    `json:"daysRange"`
	DetailedMessageShown bool   `json:"detailedMessageShown"`
	Type                 string `json:"type"`
}

type Messages struct {
	Message []Message `json:"message"`
}

type Message struct {
	ClientUnixTime string  `json:"@clientUnixTime"`
	ID             int     `json:"id"`
	MessengerID    string  `json:"messengerId"`
	MessengerName  string  `json:"messengerName"`
	UnixTime       int64   `json:"unixTime"`
	MessageType    string  `json:"messageType"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ModelID        string  `json:"modelId"`
	ShowCustomMsg  string  `json:"showCustomMsg"`
	DateTime       string  `json:"dateTime"`
	BatteryState   string  `json:"batteryState"`
	Hidden         int     `json:"hidden"`
	Altitude       int     `json:"altitude"`
	MessageContent string  `json:"messageContent"`
}

type Errors struct {
	Error *Error `json:"error"`
}

type Error struct {
	Code        string `json:"code"`
	Text        string `json:"text"`
	Description string `json:"description"`
}
