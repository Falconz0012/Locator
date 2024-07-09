package cmd

import(
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/kelvins/geocoder"
  "strings"
	"github.com/spf13/cobra"
	"strconv"
  "errors"
)

func setGeoCodeToken() string {

	godotenv.Load()
	token:= os.Getenv("GEOCODE_API_TOKEN")

	geocoder.ApiKey = token

	return token
}

func parseInt(number string)int {
	var newNumber int
	newNumber,err := strconv.Atoi(number)
	
	if err !=nil{
		fmt.Println("Invalid street number")
		os.Exit(2)
	}
	return newNumber
}

func convertAddress(cmd *cobra.Command, args []string){
  addressArray := strings.Split(address," ")

  if len(args) < 5 {
    err := errors.New("All 5 arguments needed in order to geocode address")
    fmt.Println(err) 
    os.Exit(22)
  }
	setGeoCodeToken()

	newAddress:= geocoder.Address{
		Street: addressArray[1],
		Number: parseInt(addressArray[2]),
		City  : addressArray[3],
		State : addressArray[4],
		Country: addressArray[5]}

	
	location, err := geocoder.Geocoding(newAddress)
	
	if err != nil {
		fmt.Println("Could not get the location ", err)
	}else{
    fmt.Println("Latitude: ", location.Latitude)
    fmt.Println("Longitude: ", location.Longitude)
	}
	
}
