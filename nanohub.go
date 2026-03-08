package nats

import "time"

// NanoHubDeviceInfoResponse maps the full MDM DeviceInformation plist response.
// Used on NATS subject: nanohub.deviceinfo
type NanoHubDeviceInfoResponse struct {
	UDID           string                          `plist:"UDID"`
	CommandUUID    string                          `plist:"CommandUUID"`
	Status         string                          `plist:"Status"`
	QueryResponses NanoHubDeviceInfoQueryResponse `plist:"QueryResponses"`
}

type NanoHubDeviceInfoQueryResponse struct {
	ActiveManagedUsers               []string                  `plist:"ActiveManagedUsers"`
	AvailableDeviceCapacity          float64                   `plist:"AvailableDeviceCapacity"`
	AwaitingConfiguration            bool                      `plist:"AwaitingConfiguration"`
	BatteryLevel                     float64                   `plist:"BatteryLevel"`
	BluetoothMAC                     string                    `plist:"BluetoothMAC"`
	BuildVersion                     string                    `plist:"BuildVersion"`
	CurrentConsoleManagedUser        string                    `plist:"CurrentConsoleManagedUser"`
	DeviceCapacity                   float64                   `plist:"DeviceCapacity"`
	DeviceName                       string                    `plist:"DeviceName"`
	EACSPreflight                    string                    `plist:"EACSPreflight"`
	EthernetMAC                      string                    `plist:"EthernetMAC"`
	WiFiMAC                          string                    `plist:"WiFiMAC"`
	HasBattery                       bool                      `plist:"HasBattery"`
	HostName                         string                    `plist:"HostName"`
	IsActivationLockEnabled          bool                      `plist:"IsActivationLockEnabled"`
	IsActivationLockSupported        bool                      `plist:"IsActivationLockSupported"`
	IsAppleSilicon                   bool                      `plist:"IsAppleSilicon"`
	IsSupervised                     bool                      `plist:"IsSupervised"`
	LocalHostName                    string                    `plist:"LocalHostName"`
	Model                            string                    `plist:"Model"`
	ModelName                        string                    `plist:"ModelName"`
	ModelNumber                      string                    `plist:"ModelNumber"`
	OSVersion                        string                    `plist:"OSVersion"`
	OSUpdateSettings                 NanoHubOSUpdateSettings   `plist:"OSUpdateSettings"`
	PINRequiredForDeviceLock         bool                      `plist:"PINRequiredForDeviceLock"`
	PINRequiredForEraseDevice        bool                      `plist:"PINRequiredForEraseDevice"`
	ProductName                      string                    `plist:"ProductName"`
	ProvisioningUDID                 string                    `plist:"ProvisioningUDID"`
	SerialNumber                     string                    `plist:"SerialNumber"`
	SoftwareUpdateDeviceID           string                    `plist:"SoftwareUpdateDeviceID"`
	SupplementalBuildVersion         string                    `plist:"SupplementalBuildVersion"`
	SupportsLOMDevice                bool                      `plist:"SupportsLOMDevice"`
	SupportsiOSAppInstalls           bool                      `plist:"SupportsiOSAppInstalls"`
	SystemIntegrityProtectionEnabled bool                      `plist:"SystemIntegrityProtectionEnabled"`
	TimeZone                         string                    `plist:"TimeZone"`
	UDID                             string                    `plist:"UDID"`
}

type NanoHubOSUpdateSettings struct {
	AutoCheckEnabled                bool   `plist:"AutoCheckEnabled"`
	AutomaticAppInstallationEnabled bool   `plist:"AutomaticAppInstallationEnabled"`
	AutomaticOSInstallationEnabled  bool   `plist:"AutomaticOSInstallationEnabled"`
	AutomaticSecurityUpdatesEnabled bool   `plist:"AutomaticSecurityUpdatesEnabled"`
	BackgroundDownloadEnabled       bool   `plist:"BackgroundDownloadEnabled"`
	CatalogURL                      string `plist:"CatalogURL"`
	IsDefaultCatalog                bool   `plist:"IsDefaultCatalog"`
	PreviousScanDate                time.Time `plist:"PreviousScanDate"`
	PreviousScanResult              int    `plist:"PreviousScanResult"`
}

// NanoHubApplicationListItem maps an InstalledApplicationList plist entry.
// Used on NATS subject: nanohub.installedapplicationslist
type NanoHubApplicationListItem struct {
	BundleSize   int64  `plist:"BundleSize"`
	DynamicSize  int64  `plist:"DynamicSize"`
	Identifier   string `plist:"Identifier"`
	Name         string `plist:"Name"`
	ShortVersion string `plist:"ShortVersion"`
	Version      string `plist:"Version"`
	IsValidated  bool   `plist:"IsValidated"`
}

// NanoHubUser maps a UserList plist entry.
// Used on NATS subject: nanohub.userslist
type NanoHubUser struct {
	DataQuota      int64  `plist:"DataQuota"`
	DataUsed       int64  `plist:"DataUsed"`
	HasDataToSync  bool   `plist:"HasDataToSync"`
	HasSecureToken bool   `plist:"HasSecureToken"`
	IsLoggedIn     bool   `plist:"IsLoggedIn"`
	UserName       string `plist:"UserName"`
	FullName       string `plist:"FullName"`
	MobileAccount  bool   `plist:"MobileAccount"`
	UID            int64  `plist:"UID"`
	UserGUID       string `plist:"UserGUID"`
}
