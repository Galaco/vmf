package vmf


type Vmf struct {
	VersionInfo VersionInfo
	ViewSettings ViewSettings
	VisGroup []VisGroup
	World World
	Cameras Cameras
	Cordon Cordon // Pre-L4D only
	Cordons Cordons // Post-L4D only
}
