package main

import (
	"fmt"
	"github.com/rglass-OSIsoft/OMF/OMFHelpers"
	"github.com/rglass-OSIsoft/OMF/OMFHttpClient"
	"github.com/rglass-OSIsoft/OMF/OMFJsonGenerator"
)

func main() {
	fmt.Println("Hello OMF")

	typeId := "Aircraft_Altitude"
	propertyName := "Altitude"
	typeJson := omfjsongenerator.GenerateType(typeId, propertyName)
	omfhttpclient.CreateOMFType(typeJson)

	streamId := "N123GG" 
	streamJson := omfjsongenerator.GenerateStream(streamId, typeId)
	omfhttpclient.CreateOMFStream(streamJson)

	altitude := omfhelpers.GetAltitude()
	timestamp := omfhelpers.GetCurrentTimeFormatted()
	dataJson := omfjsongenerator.GenerateValues(streamId, propertyName, timestamp, altitude)
	omfhttpclient.SendOMFData(dataJson)

	fmt.Println(dataJson)

	fmt.Println("Done")
}
