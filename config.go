package soda

// config.go
type DriverType string
type DriverContext string

const (
	ContextBlock     DriverContext = "block"
	ContextFile      DriverContext = "file"
	ContextObject    DriverContext = "object"
	ContextTelemetry DriverContext = "telemetry"

	AWSDriverType                  DriverType = "aws"
	CinderDriverType               DriverType = "cinder"
	CephDriverType                 DriverType = "ceph"
	LVMDriverType                  DriverType = "lvm"
	NFSDriverType                  DriverType = "nfs"
	IBMSpectrumScaleDriverType     DriverType = "spectrumscale"
	HuaweiOceanStorBlockDriverType DriverType = "huawei_oceanstor_block"
	HuaweiFusionStorageDriverType  DriverType = "huawei_fusionstorage"
	HuaweiOceanStorFileDriverType  DriverType = "huawei_oceanstor_file"
	HpeNimbleDriverType            DriverType = "hpe_nimble"
	DRBDDriverType                 DriverType = "drbd"
	ScutechCMSDriverType           DriverType = "scutech_cms"
	ManilaDriverType               DriverType = "manila"
	FujitsuEternusDriverType       DriverType = "fujitsu_eternus"
	ChubaofsDriverType             DriverType = "chubaofs"
	NetappOntapSanDriverType       DriverType = "netapp_ontap_san"
)
