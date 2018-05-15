package omfjsongenerator

import "strconv"

func GenerateType(typeId, propertyName string) string {
	var json = "[ { " +
		"\"id\": \"" + typeId + "\"," +
		"\"type\": \"object\"," +

		// friendly name - we'll just use the typeId for now
		"\"name\": \"" + typeId + "\"," +
		"\"classification\": \"dynamic\"," +

		// add in the property
		"\"properties\": {" +
		"\"" + propertyName + "\": {" +
		"\"type\": \"integer\"," +
		"\"format\": \"int32\"," +
		"\"name\": \"" + propertyName + "\"" +
		"}," +
		//  and the timestamp -
		"\"TimeId\": {" +
		"\"type\": \"string\"," +
		"\"format\": \"date-time\"," +

		// note this isindex field set to true for the timestamp property
		"\"isindex\": true," +
		"\"name\": \"TimeId\" }}} ]"

	return json
}

func GenerateStream(streamId, typeId string) string {
	var json = "[ { \"id\":\"" + streamId + "\", \"typeId\":\"" + typeId + "\" } ]"
	return json
}

func GenerateValues( containerId, propertyName, timestamp string, value int) string {
            var json = " [ { " +
                                  " \"containerid\": \"" + containerId + "\", " +
                                  " \"values\": [ { " +
                                            " \"" + propertyName + "\":" + strconv.Itoa(value) + ", " +
                                            " \"TimeId\": \"" + timestamp + "\" } ] " +
                                  " } ] ";
            return json;
}