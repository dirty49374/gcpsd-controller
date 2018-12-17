package verdb

import "log"
import appsv1 "k8s.io/api/apps/v1"

type ImageVersion struct {
	FullName string
	Numbers  []int
}

type PodRef struct {
	Namespace string
	Name      string
}

// POD --> VER MAP
type DeployDB struct {
	pods     map[string]*appsv1.Deployment
	patterns map[string]*DeploySpec
}

func NewDeployDB() *DeployDB {
	db := &DeployDB{
		patterns: make(map[string]*DeploySpec, 0),
	}
	return db
}

func (db *DeployDB) OnNewImage(version string) {

}

func (db *DeployDB) OnDeployment(deployment *appsv1.Deployment) {

}

func (db *DeployDB) OnNewDeploySpec(pattern string) {
	spec, err := NewDeploySpec(pattern)
	if err != nil {
		log.Print(err)
		return
	}

	if _, ok := db.patterns[pattern]; ok {
		return
	}

	db.patterns[pattern] = spec
}
