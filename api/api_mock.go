// Copyright (c) 2017-2021 Ingram Micro Inc.

package api

//import "github.com/stretchr/testify/mock"
//
//// MockIMCOClient service manager
//type MockIMCOClient struct {
//	mock.Mock
//}
//
//// Post mocks POST request to IMCO API
//func (m *MockIMCOClient) Post(path string, payload *map[string]interface{}) ([]byte, int, error) {
//	args := m.Called(path, payload)
//	return args.Get(0).([]byte), args.Int(1), args.Error(2)
//}
//
//// Put mocks PUT request to IMCO API
//func (m *MockIMCOClient) Put(path string, payload *map[string]interface{}) ([]byte, int, error) {
//	args := m.Called(path, payload)
//	return args.Get(0).([]byte), args.Int(1), args.Error(2)
//}
//
//// Delete mocks DELETE request to IMCO API
//func (m *MockIMCOClient) Delete(path string) ([]byte, int, error) {
//	args := m.Called(path)
//	return args.Get(0).([]byte), args.Int(1), args.Error(2)
//}
//
//// Get mocks GET request to IMCO API
//func (m *MockIMCOClient) Get(path string) ([]byte, int, error) {
//	args := m.Called(path)
//	return args.Get(0).([]byte), args.Int(1), args.Error(2)
//}
//
//// GetFile sends GET request to IMCO API and receives a file
//func (m *MockIMCOClient) GetFile(url string, filePath string, discoveryFileName bool) (string, int, error) {
//	args := m.Called(url, filePath)
//	return args.String(0), args.Int(1), args.Error(2)
//}
//
//// PutFile sends PUT request to send a file
//func (m *MockIMCOClient) PutFile(sourceFilePath string, targetURL string) ([]byte, int, error) {
//	args := m.Called(sourceFilePath, targetURL)
//	return args.Get(0).([]byte), args.Int(1), args.Error(2)
//}
