package http

type DeviceResponse struct {
	Device struct {
		FrameTVSupport    string `json:"FrameTVSupport"`
		GamePadSupport    string `json:"GamePadSupport"`
		ImeSyncedSupport  string `json:"ImeSyncedSupport"`
		Os                string `json:"OS"`
		TokenAuthSupport  string `json:"TokenAuthSupport"`
		VoiceSupport      string `json:"VoiceSupport"`
		CountryCode       string `json:"countryCode"`
		Description       string `json:"description"`
		DeveloperIP       string `json:"developerIP"`
		DeveloperMode     string `json:"developerMode"`
		Duid              string `json:"duid"`
		FirmwareVersion   string `json:"firmwareVersion"`
		ID                string `json:"id"`
		IP                string `json:"ip"`
		Model             string `json:"model"`
		ModelName         string `json:"modelName"`
		Name              string `json:"name"`
		NetworkType       string `json:"networkType"`
		Resolution        string `json:"resolution"`
		SmartHubAgreement string `json:"smartHubAgreement"`
		Type              string `json:"type"`
		Udn               string `json:"udn"`
		WifiMac           string `json:"wifiMac"`
	} `json:"device"`
	ID        string `json:"id"`
	IsSupport string `json:"isSupport"`
	Name      string `json:"name"`
	Remote    string `json:"remote"`
	Type      string `json:"type"`
	URI       string `json:"uri"`
	Version   string `json:"version"`
}

type ApplicationResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Running bool   `json:"running"`
	Version string `json:"version"`
	// If the current application is up and running on the screen, is the
	// application running and can the people see it?
	Visible bool `json:"visible"`
}
