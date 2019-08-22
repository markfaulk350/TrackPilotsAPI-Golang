package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/helpers"
)

func (svc ServiceImpl) RetreiveDataFromSpot(user entity.User, whenToQueryFrom int64) error {
	// https://api.findmespot.com/spot-main-web/consumer/rest-api/2.0/public/feed/FEED_ID_HERE/message.xml?startDate=2012-07-03T13:08:55-0000&endDate=20012-08-02T:08:55-0000
	//finalLink := "https://api.findmespot.com/spot-main-web/consumer/rest-api/2.0/public/feed/0fTO7dcPGBEIRy3x8VlDPURJF1IYcYxJx/message.json"
	spotFormatedDate := time.Unix(whenToQueryFrom, 0).Format("2006-01-02T15:04:05-0000")

	finalLink := (user.Trklink + "/message.json?startDate=" + spotFormatedDate)

	// fmt.Println("This is the spot tracking link. See if the format is correct!")
	// fmt.Println(finalLink)

	resp, err := http.Get(finalLink)
	if err != nil {
		fmt.Println(finalLink, "http request to this spot tracking link is not working.")
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the spot request")
		return err
	}

	spotTrackerResponse := entity.SpotDataStructure{}
	error := json.Unmarshal(body, &spotTrackerResponse)
	if error != nil {
		fmt.Println("Error marshaling the spot tracking request")
		return err
	}

	if spotTrackerResponse.Response.Errors != nil {
		shortErr := spotTrackerResponse.Response.Errors.Error
		fmt.Println("An error has been thrown while grabbing data from spot. Response may be empty.", shortErr.Code, shortErr.Text, shortErr.Description)

		sqlStatement := `UPDATE pilots SET lastApiCall=? WHERE id=?;`
		_, err = svc.DBClient.Exec(sqlStatement, time.Now().Unix(), user.ID)
		if err != nil {
			fmt.Println("Not able to update pilots lastApiCall when the response is empty")
			return err
		}
	}

	// Need to run some checks to see if the response is valid and has usable data, else return and print an err to console
	var b bytes.Buffer
	b.WriteString("INSERT IGNORE INTO pings(unixTime, lat, lng, alt, txtMsg, isEmergency, pilot_id) VALUES")

	short := spotTrackerResponse.Response.FeedMessageResponse.Messages.Message
	for i := 0; i < len(short); i++ {
		fmt.Println(short[i].UnixTime, short[i].DateTime, short[i].BatteryState, short[i].Latitude, short[i].Longitude, short[i].Altitude, short[i].MessageContent, short[i].MessageType)
		var status string
		if short[i].MessageType != "HELP" {
			status = "False"
		} else {
			status = "True"
		}
		b.WriteString("(" + helpers.Int64ToString(short[i].UnixTime) + ", " + helpers.Float64ToString(short[i].Latitude) + ", " + helpers.Float64ToString(short[i].Longitude) + ", " + helpers.IntToString(short[i].Altitude) + ", '" + short[i].MessageContent + "', '" + status + "', " + helpers.IntToString(user.ID) + "),")
	}
	finalQueryString := b.String()
	finalQueryString = finalQueryString[:len(finalQueryString)-1]
	fmt.Println(finalQueryString)

	_, err = svc.DBClient.Exec(finalQueryString)
	if err != nil {
		fmt.Println("Unable to update database with tracking data")
	}

	latestPing := short[0].UnixTime

	// update pilot table
	sqlStatement := `UPDATE pilots SET lastLocationPing=?, lastApiCall=? WHERE id=?;`
	_, err = svc.DBClient.Exec(sqlStatement, latestPing, time.Now().Unix(), user.ID)
	if err != nil {
		fmt.Println("Unable to update pilot table to reflect last api call")
	}

	return nil
}
