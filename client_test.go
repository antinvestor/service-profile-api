package profile_v1

import (
	"context"
	"github.com/antinvestor/apis"
	"google.golang.org/grpc"
	"testing"
)

func TestNewProfileClient(t *testing.T) {

	ctx := context.Background()
	_, err := NewProfileClient(ctx, apis.WithEndpoint("127.0.0.1:7005"))
	if err != nil {
		t.Errorf("Could not setup profile service : %v", err)
	}

}



func TestProfileClient_GetProfileByID(t *testing.T) {

	ctx := context.Background()



	//pcli, err := NewProfileClient(ctx, apis.WithEndpoint("127.0.0.1:7005"), apis.WithoutAuthentication())
	//if err != nil {
	//	t.Errorf("Could not setup profile service : %v", err)
	//}

	serviceConnection, err := grpc.Dial("localhost:7005", grpc.WithInsecure())

	profileService := NewProfileServiceClient(serviceConnection)

	profileRequest := ProfileIDRequest{
		ID: "text_id",
	}

	_, err = profileService.GetByID(ctx, &profileRequest)

	if err != nil{
		t.Errorf("There was an error obtaining the profile : %v", err)
	}



}
