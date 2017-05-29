# kubernetes-aws-volume-discovery

Scans AWS EBS volumes and adds them as PersistentVolume to Kubernetes.

## Warning: *Deprecated!*
Instead we are now using Storage Classes (https://kubernetes.io/docs/concepts/storage/persistent-volumes/#storageclasses).

## Usage

*You need access to our AWS account `074509403805` for these steps.*

1. Get access to ECR: `eval $(aws ecr get-login --profile ecr-master --region eu-west-1)`
2. Build and push Docker image: `hack/push.sh`
3. Let Kubernetes access ECR: https://github.com/rebuy-de/kubernetes-example-application/blob/master/hack/create-ecr-imagepullsecret
4. Deploy image: `kcr apply -f kubernetes/`
5. Verify: `kcr get pods`, `kcr get pv`
