# Vertical Auto-Scaling using K8s
[![CI](https://github.com/dipankardas011/VMWare/actions/workflows/main.yml/badge.svg)](https://github.com/dipankardas011/VMWare/actions/workflows/main.yml)

## Group members
1. Dipankar Das  `Kalinga Institute Industrial Technology`

## Tech Stack

1. Go
2. Simple csv file for the load
3. Docker

## Demo Link
[*`Youtube Video Link`*](https://youtu.be/szN6RZKE51I)

## References
1. VPA [Kubernetes autoscaler](https://github.com/kubernetes/autoscaler) and some of the [Google docs](https://cloud.google.com/kubernetes-engine/docs/concepts/verticalpodautoscaler)
2. Bigdata package for go[BigData](https://github.com/go-gota/gota/)
3. Charts for go[Charts](https://blog.logrocket.com/visualizing-data-go-echarts/)
4. Big data from [UCI Machine Learning Repository](https://archive.ics.uci.edu/ml/index.php)

## How to Run

Container Image: `dipugodocker/vmware-sol:latest`
```bash
kubectl apply -f VPA/manifest.yml
kubectl describe deploy my-big-data-app
```

## VPA

VPA docs [here]("https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler#install-command") <br/>
Let's install the VPA from a container that can access our cluster

```bash
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler/

./hack/vpa-up.sh

# after few seconds, we can see the VPA components in:

kubectl -n kube-system get pods

cd -

kubectl apply -f VPA/AutoScaler.yml

# to view the recommendations
kubectl get vpa
kubectl describe vpa <Name> # to get the recommendations
# when we are Satisfied with the result being update it to the `auto` mode
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
