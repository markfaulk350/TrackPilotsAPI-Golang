package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/markfaulk350/TrackPilotsAPI/helpers"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func (svc ServiceImpl) RetreiveDataFromGarmin(user entity.User, whenToQueryFrom int64) error {
	whenToQueryFromFormated := time.Unix(whenToQueryFrom+1, 0).Format("2006-01-02T15:04z")
	finalLink := (user.Trklink + "?d1=" + whenToQueryFromFormated)
	fmt.Println("Get Data From Garmin has fired! ", user.ID, finalLink)

	resp, err := http.Get(finalLink)
	if err != nil {
		fmt.Println(finalLink, "Get request to Garmin tracking API failed.")
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body from Garmin API.")
		return err
	}

	var kmlFile entity.GarminDataStructure
	err = xml.Unmarshal(body, &kmlFile)
	if err != nil {
		fmt.Printf("Problem unmarshalling data from Garmin kml file.")
		return err
	}

	// If there is no data update that we tried to get new data at a specific time, but nothing was avaliable and return
	if kmlFile.Document.Folder == nil {
		fmt.Println("The location data request seems to be empty, the pilot might be inactive, might not have new data avaliable yet, or the tracking link might not work.")
		// need to update the pilot table to reflect when the api call was made
		sqlStatement := `UPDATE pilots SET lastApiCall=? WHERE id=?;`
		_, err = svc.DBClient.Exec(sqlStatement, time.Now().Unix(), user.ID)
		if err != nil {
			fmt.Println("Not able to update pilots lastApiCall when the response is empty")
		}
		return err
	}

	// We have a struct containing all the data we need, now we have to create a query string to insert into DB
	var b bytes.Buffer
	b.WriteString("INSERT IGNORE INTO pings(unixTime, lat, lng, alt, velocity, heading, txtMsg, isEmergency, pilot_id) VALUES")

	short := kmlFile.Document.Folder.PlacemarkList
	for i := 0; i < len(short)-1; i++ {
		s := strings.Split(short[i].Point.Coordinates, ",")
		b.WriteString("(" + getUnixTimestamp(short[i].TimeStamp.When) + ", " + s[1] + ", " + s[0] + ", " + s[2] + ", '" + short[i].ExtendedData.DataList[11].Value + "', '" + short[i].ExtendedData.DataList[12].Value + "', '" + short[i].ExtendedData.DataList[15].Value + "', '" + short[i].ExtendedData.DataList[14].Value + "', " + helpers.IntToString(user.ID) + "),")
	}
	finalQueryString := b.String()
	finalQueryString = finalQueryString[:len(finalQueryString)-1]
	fmt.Println(finalQueryString)

	_, err = svc.DBClient.Exec(finalQueryString)
	if err != nil {
		fmt.Println("Unable to update database with tracking data")
		return err
	}

	// Then once successful we need to add the last location ping and api call timestamp
	latestPing := getUnixTimestamp(short[len(short)-2].TimeStamp.When)
	// Might have to convert tring to int64 or reverse when patching
	//fmt.Println("This is where we patch the lastApiCall and lastLocationPing.", pilotsID, latestPing, whenToQueryTill)

	sqlStatement := `UPDATE pilots SET lastLocationPing=?, lastApiCall=? WHERE id=?;`
	_, err = svc.DBClient.Exec(sqlStatement, latestPing, time.Now().Unix(), user.ID)
	if err != nil {
		fmt.Println("Unable to set pilots last location ping or update last api call.")
		return err
	}

	return nil
}

func getUnixTimestamp(input string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z", input)
	if err != nil {
		fmt.Println("Error:", err)
	}
	s := strconv.FormatInt(t.Unix(), 10)
	return s
}
