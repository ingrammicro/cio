package testdata

import "github.com/ingrammicro/concerto/api/types"

// GetCloudProviderData loads test data
func GetCloudProviderData() *[]types.CloudProvider {

	testCloudProviders := []types.CloudProvider{
		{
			Id:                  "fakeID0",
			Name:                "fakeName0",
			ProvidedServices:    []string{"fakeService01", "fakeService02", "fakeService03"},
			RequiredCredentials: []string{"fakeCredential01", "fakeCredential02"},
		},
		{
			Id:                  "fakeID1",
			Name:                "fakeName1",
			ProvidedServices:    []string{"fakeService11", "fakeService12"},
			RequiredCredentials: []string{"fakeCredential11", "fakeCredential12", "fakeCredential13"},
		},
	}

	return &testCloudProviders
}
