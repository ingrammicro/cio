package api

//func initConfigAndClientAPI() (*ClientAPI, *configuration.Config, error) {
//	config := initConfig()
//	//config := new(configuration.Config) // para forzar error de config
//	ds, err := NewHTTPClient(config)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	svc := new(ClientAPI)
//	svc.HTTPClient = *ds
//	return svc, config, nil
//}

//func TestClientAPI(t *testing.T) {
//	events := pathAuditEvents
//	svc, config, err := InitConfigAndClientAPI()
//	if err != nil && svc == nil && config == nil {
//		t.Errorf("Unexpected error: %v\n", err)
//		return
//	}
//}
