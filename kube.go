package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func kubeApply(vol PersistentVolume) error {
	sts, err := kubeHttp("POST", "", vol)

	if err != nil {
		return err
	}

	if sts == 409 {
		sts, err = kubeHttp("PUT", vol.Metadata.Name, vol)
		if err != nil {
			return errors.Wrap(err, "PUT request failed")
		}
	}

	if sts != 200 && sts != 201 {
		return errors.New(fmt.Sprintf("unexpected HTTP status code %d", sts))
	}

	return nil
}

func kubeHttp(method string, path string, vol PersistentVolume) (int, error) {
	var (
		caPath         = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
		tokenPath      = "/var/run/secrets/kubernetes.io/serviceaccount/token"
		kubernetesHost = os.Getenv("KUBERNETES_SERVICE_HOST")
		kubernetesPort = os.Getenv("KUBERNETES_SERVICE_PORT")
	)

	cacert, err := ioutil.ReadFile(caPath)
	if err != nil {
		return 0, err
	}

	token, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return 0, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(cacert))
	if !ok {
		return 0, fmt.Errorf("failed to parse root certificate")
	}
	tlsConf := &tls.Config{RootCAs: roots}
	tr := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: tr}

	data, err := json.Marshal(vol)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(method, fmt.Sprintf(
		"https://%s:%s/api/v1/persistentvolumes/%s",
		kubernetesHost, kubernetesPort, path),
		bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", string(token)))

	log.Printf("%+v", req)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%T %s", err, err.Error())
		return 0, err
	}

	defer resp.Body.Close()

	log.Printf("%+v", resp)
	return resp.StatusCode, nil
}
