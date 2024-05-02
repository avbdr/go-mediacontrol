package upnp

import "encoding/xml"

type GetDeviceVolumeResponse struct {
	Envelope struct {
		Body struct {
			GetVolumeResponse struct {
				CurrentVolume string `xml:"CurrentVolume"`
			} `xml:"GetVolumeResponse"`
		} `xml:"Body"`
	} `xml:"Envelope"`
}

type GetCurrentMuteStatusResponse struct {
	Envelope struct {
		Body struct {
			GetMuteResponse struct {
				CurrentMute string `xml:"CurrentMute"`
			} `xml:"GetMuteResponse"`
		} `xml:"Body"`
	} `xml:"Envelope"`
}

type GetPositionInfoResponse struct {
	Envelope struct {
		Body struct {
			GetPositionResponse struct {
				Track         string
				RelTime       string
				TrackDuration string
				TrackMetaData string
				TrackURI      string
			} `xml:"GetPositionInfoResponse"`
		} `xml:"Body"`
	} `xml:"Envelope"`
}

type upnpDescribeDevice_XML struct {
	XMLNamespace string           `xml:"xmlns,attr"`
	Device       []upnpDevice_XML `xml:"device"`
}

type upnpDevice_XML struct {
	FriendlyName     string `xml:"friendlyName"`
	Manufacturer     string `xml:"manufacturer"`
	ModelNumber      string `xml:"modelNumber"`
	ModelDescription string `xml:"modelDescription"`
	ModelName        string `xml:"modelName"`
	RoomName         string `xml:"roomName"`
	MacAddress       string `xml:"MACAddress"`
	Udn				 string `xml:"UDN"`
}

type Item_XML struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Album       string   `xml:"album"`
	AlbumArtUri string   `xml:"albumArtURI"`
	Creator     string   `xml:"creator"`
}

type TrackMetaData_XML struct {
	XMLName xml.Name `xml:"DIDL-Lite"`
	Item    Item_XML `xml:"item"`
}
