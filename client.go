package profile_v1

import (
	"context"
	apic "github.com/antinvestor/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"

	"math"
)

const ctxKeyService = "profileClientKey"

func defaultProfileClientOptions() []apic.ClientOption {
	return []apic.ClientOption{
		apic.WithEndpoint("profile.api.antinvestor.com:443"),
		apic.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		apic.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func ToContext(ctx context.Context, profileClient *ProfileClient) context.Context {
	return context.WithValue(ctx, ctxKeyService, profileClient)
}

func FromContext(ctx context.Context) *ProfileClient {
	profileClient, ok := ctx.Value(ctxKeyService).(*ProfileClient)
	if !ok {
		return nil
	}

	return profileClient
}

// ProfileClient is a client for interacting with the profile service API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ProfileClient struct {
	// gRPC connection to the service.
	clientConn *grpc.ClientConn

	// The gRPC API client.
	profileClient ProfileServiceClient

	// The x-ant-* metadata to be sent with each request.
	xMetadata metadata.MD
}

// NewProfileClient creates a new notification client.
//
// The service that an application uses to send and access received messages
func NewProfileClient(ctx context.Context, opts ...apic.ClientOption) (*ProfileClient, error) {
	clientOpts := defaultProfileClientOptions()

	connPool, err := apic.DialConnection(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	cl := &ProfileClient{
		clientConn:    connPool,
		profileClient: NewProfileServiceClient(connPool),
	}

	cl.setClientInfo()

	return cl, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (nc *ProfileClient) Close() error {
	return nc.clientConn.Close()
}

// setClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (nc *ProfileClient) setClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", apic.VersionGo()}, keyval...)
	kv = append(kv, "grpc", grpc.Version)
	nc.xMetadata = metadata.Pairs("x-ai-api-client", apic.XAntHeader(kv...))
}

func (nc *ProfileClient) GetProfileByID(ctx context.Context, profileId string) (*ProfileObject, error) {

	profileCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	profileService := NewProfileServiceClient(nc.clientConn)

	profileRequest := ProfileIDRequest{
		ID: profileId,
	}

	return profileService.GetByID(profileCtx, &profileRequest)
}

func (nc *ProfileClient) CreateProfileByContactAndName(ctx context.Context, contact string, name string) (*ProfileObject, error) {

	profileCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	profileService := NewProfileServiceClient(nc.clientConn)

	properties := make(map[string]string)
	properties["name"] = name

	createProfileRequest := ProfileCreateRequest{
		Contact:    contact,
		Properties: properties,
	}

	return profileService.Create(profileCtx, &createProfileRequest)
}

func (nc *ProfileClient) GetProfileByContact(ctx context.Context, contact string) (*ProfileObject, error) {

	profileService := NewProfileServiceClient(nc.clientConn)

	profileCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	contactRequest := ProfileContactRequest{
		Contact: contact,
	}

	return profileService.GetByContact(profileCtx, &contactRequest)
}
