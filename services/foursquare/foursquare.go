package foursquare

import (
	"speedcafe/interfaces"

	"github.com/elbuo8/4square-venues"
)

type FoursquareService struct {
	Client *fsvenues.FSClient
}

func NewFoursquareService(clientID, clientSecretKey string) interfaces.IFoursquareService {
	client := fsvenues.NewFSVenuesClient(clientID, clientSecretKey)

	return &FoursquareService{
		Client: &client,
	}
}
