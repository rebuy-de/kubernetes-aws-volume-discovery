package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func aws2kube(vol *ec2.Volume) *PersistentVolume {
	var name string

	for _, tag := range vol.Tags {
		if *tag.Key == "Name" {
			name = *tag.Value
		}
	}

	if name == "" {
		log.Printf("Skip volume %s without name.", *vol.VolumeId)
		return nil
	}

	log.Printf("Mapping volume %s with name '%s'", *vol.VolumeId, name)

	pv := PersistentVolume{
		ApiVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: Metadata{
			Name: name,
		},
		Spec: PersistentVolumeSpec{
			Capacity: Capacity{
				Storage: fmt.Sprintf("%dGi", *vol.Size),
			},
			AccessModes:   []string{"ReadWriteOnce"},
			ReclaimPolicy: "Recycle",
			EBS: EBS{
				ID:   *vol.VolumeId,
				Type: "ext4",
			},
		},
	}

	return &pv
}
