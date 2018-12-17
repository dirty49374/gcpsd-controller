package verdb

type Xxx struct {
	*VersionPattern
	CurrentVersion string
}

type DeployXxxx struct {
	NamespacedName string
	Specs          []DeploySpec
}

type DeploySpec struct {
	*VersionPattern
	CurrentImage string
}

func NewDeploySpec(pattern string) (*DeploySpec, error) {

	vp, err := NewVersionPattern(pattern)
	if err != nil {
		return nil, err
	}

	return &DeploySpec{
		VersionPattern: vp,
		CurrentImage:   "",
	}, nil
}
