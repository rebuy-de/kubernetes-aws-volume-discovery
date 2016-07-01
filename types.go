package main

type PersistentVolume struct {
	ApiVersion string               `json:"apiVersion"`
	Kind       string               `json:"kind"`
	Metadata   Metadata             `json:"metadata"`
	Spec       PersistentVolumeSpec `json:"spec"`
}

type PersistentVolumeSpec struct {
	Capacity      Capacity `json:"capacity"`
	AccessModes   []string `json:"accessModes"`
	ReclaimPolicy string   `json:"persistentVolumeReclaimPolicy"`
	EBS           EBS      `json:"awsElasticBlockStore"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Capacity struct {
	Storage string `json:"storage"`
}

type EBS struct {
	ID   string `json:"volumeID"`
	Type string `json:"fsType"`
}
