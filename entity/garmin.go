package entity

import "encoding/xml"

type GarminDataStructure struct {
	XMLName  xml.Name  `xml:"kml"`
	Document *Document `xml:"Document"`
}

type Document struct {
	XMLName xml.Name `xml:"Document"`
	Folder  *Folder  `xml:"Folder"`
	Name    string   `xml:"name"`
}

type Folder struct {
	XMLName       xml.Name    `xml:"Folder"`
	Name          string      `xml:"name"`
	PlacemarkList []Placemark `xml:"Placemark"`
}

type Placemark struct {
	Visibility   string        `xml:"visibility"`
	ExtendedData *ExtendedData `xml:"ExtendedData"`
	Point        *Point        `xml:"Point"`
	TimeStamp    *TimeStamp    `xml:"TimeStamp"`
}

type ExtendedData struct {
	DataList []Data `xml:"Data"`
}

type Data struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type Point struct {
	Coordinates string `xml:"coordinates"`
}

type TimeStamp struct {
	When string `xml:"when"`
}
