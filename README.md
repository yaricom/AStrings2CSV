# AStrings2CSV
Converter of Android string resources file to CSV in order to have it for external translations

# Overview
Android string resources provided as XML file which is not always handy for linguist folks who do translate it. 
This tools allows convertation of Android string XML resources to CSV file which can be opened by any Office suite for further processing.

# Usage
**$ go run main.go inputFile [outputFile]**

where:
* inputFile - the path to the Android String XML file
* outputFile - the path to the output file (optional). If missed than name of input will be used by appending .csv suffix
