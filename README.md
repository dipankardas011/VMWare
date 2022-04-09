# Vertical Auto-Scaling using K8s
[![CI](https://github.com/dipankardas011/VMWare/actions/workflows/main.yml/badge.svg)](https://github.com/dipankardas011/VMWare/actions/workflows/main.yml)

## Group members
1. Dipankar Das  `Kalinga Institute Industrial Technology`

## Tech Stack

1. Go
2. Simple csv file for the load
3. Docker

## References
1. Kuberenetes/AutoScaler
2. Charts [Charts](https://blog.logrocket.com/visualizing-data-go-echarts/)

## How to Run

Container Image: `dipugodocker/vmware-sol:latest`

## VPA

VPA docs [here]("https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#install-command") <br/>
Let's install the VPA from a container that can access our cluster

```shell
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler/

./hack/vpa-up.sh

# after few seconds, we can see the VPA components in:

kubectl -n kube-system get pods

# to view the recommendations
kubectl get vpa
kubectl describe vpa <Name>
```

## VPA installation
<img src="./01.png" width="80%" height=auto>

## Recommendation provided by the autoscaler
<img src="./02.png" width="80%" height=auto>

## Current State of VPA and POD
<img src="./03.png" width="80%" height=auto>

## Metric for CPU util.
<img src="./04.png" width="80%" height=auto>

<!-- ## Recommendation provided by the autoscaler
<img src="./05.png" width="80%" height=auto> -->
