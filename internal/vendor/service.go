package vendor

import (
	"os"
)

type VendorService struct {
	collection *mongo.Collection
}

func NewUVendorService() *VendorService {
	client, err := database.DBinstance()
	if err != nil {
		panic(err)
	}
	return &UserService{
		collection: client.Database("VendorGrpc").Collection("vendor"),
	}
}
