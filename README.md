<a href="https://volcano.sh/">
    <img src="https://raw.githubusercontent.com/volcano-sh/volcano/master/docs/images/volcano-horizontal-color.png"/>
</a>

-------

[![Build Status](https://travis-ci.org/volcano-sh/volcano.svg?branch=master)](https://travis-ci.org/volcano-sh/volcano)
[![Go Report Card](https://goreportcard.com/badge/github.com/volcano-sh/volcano)](https://goreportcard.com/report/github.com/volcano-sh/volcano)
[![RepoSize](https://img.shields.io/github/repo-size/volcano-sh/volcano.svg)](https://github.com/volcano-sh/volcano)
[![Release](https://img.shields.io/github/release/volcano-sh/volcano.svg)](https://github.com/volcano-sh/volcano/releases)
[![LICENSE](https://img.shields.io/github/license/volcano-sh/volcano.svg)](https://github.com/volcano-sh/volcano/blob/master/LICENSE)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B162%2Fgithub.com%2Fvolcano-sh%2Fvolcano.svg?type=shield)](https://app.fossa.com/projects/custom%2B162%2Fgithub.com%2Fvolcano-sh%2Fvolcano)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3012/badge)](https://bestpractices.coreinfrastructure.org/projects/3012)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/volcano-sh/volcano/badge)](https://scorecard.dev/viewer/?uri=github.com/volcano-sh/volcano)
[![Gurubase](https://img.shields.io/badge/Gurubase-Ask%20Volcano%20Guru-006BFF)](https://gurubase.io/g/volcano)
[![CLOMonitor](https://img.shields.io/endpoint?url=https://clomonitor.io/api/projects/cncf/volcano/badge)](https://clomonitor.io/projects/cncf/volcano)



[Volcano](https://volcano.sh/) is a Kubernetes-native batch scheduling system, extending and enhancing the capabilities of the standard kube-scheduler. It provides a comprehensive set of features specifically designed to manage and optimize various batch and elastic workloads, including Artificial Intelligence (AI) / machine learning (ML) / deep learning (DL), bioinformatics / genomics, and other "Big Data" applications.

These workloads commonly leverage AI, Big Data, and HPC frameworks such as Spark, Flink, Ray, TensorFlow, PyTorch, Argo, MindSpore, PaddlePaddle, Kubeflow, MPI, Horovod, MXNet, KubeGene, and others, with which Volcano offers robust integration.

Volcano incorporates over fifteen years of collective experience in operating diverse high-performance workloads at scale across multiple systems and platforms. It combines proven best practices and innovative concepts from the open-source community to deliver a powerful and flexible scheduling solution.

As of 2026, Volcano has seen widespread adoption across numerous industries globally, including Internet/Cloud, Finance, Manufacturing, and Medical sectors. Many organizations and institutions are not only end-users but also active contributors to the project. Hundreds of contributors actively participate in code commits, pull request reviews, issue discussions, documentation updates, and design proposals. We encourage your participation in the ongoing development and growth of the Volcano project.


> [!NOTE]
> the scheduler is built based on [kube-batch](https://github.com/kubernetes-sigs/kube-batch);
refer to [#241](https://github.com/volcano-sh/volcano/issues/241) and [#288](https://github.com/volcano-sh/volcano/pull/288) for more detail.

![cncf_logo](docs/images/cncf-logo.png)

Volcano is an incubating project of the [Cloud Native Computing Foundation](https://cncf.io/) (CNCF). Please consider joining the CNCF if you are an organization that wants to take an active role in supporting the growth and evolution of the cloud native ecosystem. 

## Overall Architecture

![volcano](docs/images/volcano-architecture.png)

## Talks

- [Intro: Kubernetes Batch Scheduling @ KubeCon 2019 EU](https://sched.co/MPi7)
- [Volcano 在 Kubernetes 中运行高性能作业实践 @ ArchSummit 2019](https://archsummit.infoq.cn/2019/shenzhen/presentation/1817)
- [Volcano：基于云原生的高密计算解决方案 @ Huawei Connection 2019](https://e.huawei.com/cn/material/event/HC/09099dce0070415e9f26ada51b2216d7)
- [Improving Performance of Deep Learning Workloads With Volcano @ KubeCon 2019 NA](https://sched.co/UaZi)
- [Batch Capability of Kubernetes Intro @ KubeCon 2019 NA](https://sched.co/Uajv)
- [Optimizing Knowledge Distillation Training With Volcano @ KubeCon 2021 EU](https://www.youtube.com/watch?v=cDPGmhVcj7Y&t=143s)
- [Exploration About Mixing Technology of Online Services and Offline Jobs Based On Volcano @ KubeCon 2021 China](https://www.youtube.com/watch?v=daqkUlT5ReY)
- [Volcano - Cloud Native Batch System for AI, Big Data and HPC @ KubeCon 2022 EU](https://www.youtube.com/watch?v=wjy35HfIP_k)
- [How to Leverage Volcano to Improve the Resource Utilization of AI Pharmaceuticals, Autonomous Driving, and Smart Buildings @ KubeCon 2023 EU](https://www.youtube.com/watch?v=ujHDV5xteqU)
- [Run Your AI Workloads and Microservices on Kubernetes More Easily and Efficiently @ KubeCon 2023 China](https://www.youtube.com/watch?v=OO7zpyf7fgs)
- [Optimize LLM Workflows with Smart Infrastructure Enhanced by Volcano @ KubeCon 2024 China](https://www.youtube.com/watch?v=77Qn1-I-muQ)
- [How Volcano Enable Next Wave of Intelligent Applications @ KubeCon 2024 China](https://www.youtube.com/watch?v=IzR7zJQ8vMw)
- [Leverage Topology Modeling and Topology-Aware Scheduling to Accelerate LLM Training @ KubeCon 2024 China](https://www.youtube.com/watch?v=IB54LHQQ8lI)
- [Cloud Native AI: Harness the Power of Advanced Scheduling for High-Performance AI/ML Training @ KubeCon 2025 EU](https://www.youtube.com/watch?v=yCyezOTVU_Y)
- [Scaling Model Training with Volcano: iFlytek's Kubernetes Breakthrough @ KubeCon 2025 China](https://www.youtube.com/watch?v=JfN73WyvPbE)
- [Intelligent Topology for AI Power: Network-Aware Scheduling Optimization with Volcano HyperNode @ KubeCon 2025 NA](https://www.youtube.com/watch?v=KwPHDI4qusc)
- [Project Lightning Talk: Next-Gen AI Orchestration with Volcano on Kubernetes @ KubeCon 2026 EU](https://www.youtube.com/watch?v=0vYUzJ7eO8M)
- [Volcano: Orchestrating the Full AI Lifecycle from Training to Inference and Agents @ KubeCon 2026 EU](https://www.youtube.com/watch?v=mzA0yuZ4QuY)


## Ecosystem

- [Spark Operator](https://www.kubeflow.org/docs/components/spark-operator/user-guide/volcano-integration/)
- [Native Spark](https://spark.apache.org/docs/latest/running-on-kubernetes.html#using-volcano-as-customized-scheduler-for-spark-on-kubernetes)
- [Flink](https://github.com/GoogleCloudPlatform/flink-on-k8s-operator/blob/master/docs/volcano_integration.md)
- [KubeRay](https://docs.ray.io/en/master/cluster/kubernetes/k8s-ecosystem/volcano.html)
- [Grove](https://github.com/ai-dynamo/grove/blob/main/docs/proposals/376-volcano-scheduler-backend/README.md)
- [PyTorch](https://github.com/volcano-sh/volcano/blob/master/docs/user-guide/how_to_use_pytorch_plugin.md)
- [TensorFlow](https://github.com/volcano-sh/volcano/tree/master/example/integrations/tensorflow)
- [kubeflow/trainer (v2)](https://www.kubeflow.org/docs/components/trainer/operator-guides/job-scheduling/volcano/)
- [kubeflow/training-operator (v1)](https://www.kubeflow.org/docs/components/trainer/legacy-v1/user-guides/job-scheduling/#volcano-scheduler)
- [kubeflow/arena](https://github.com/kubeflow/arena/blob/master/docs/training/volcanojob/volcanojob.md)
- [MPI](https://github.com/volcano-sh/volcano/tree/master/example/integrations/mpi)
- [Horovod](https://github.com/volcano-sh/volcano/blob/master/example/kubecon-2019-china/horovod-sample/lm-horovod-tf-mnist-v0.5.yaml)
- [PaddlePaddle](https://github.com/volcano-sh/volcano/tree/master/example/integrations/paddlepaddle)
- [Cromwell](https://github.com/broadinstitute/cromwell/blob/develop/docs/backends/Volcano.md)
- [MindSpore](https://github.com/volcano-sh/volcano/tree/master/example/MindSpore-example)
- [MXNet](https://github.com/volcano-sh/volcano/tree/master/example/integrations/mxnet/train)
- [Argo](https://github.com/volcano-sh/volcano/tree/master/example/integrations/argo)
- [KubeGene](https://github.com/volcano-sh/kubegene)
- [LeaderWorkerSet](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/sample/gang-scheduling#using-volcano-for-gang-scheduling)
- [Kthena](https://github.com/volcano-sh/kthena/tree/main/examples)

## Use Cases
- [How Feedzai Built a Self-Service Spark on Kubernetes Platform with Volcano (KubeCon 2026 EU)](https://www.youtube.com/watch?v=jmz8c3oEmPY)
- [iFLYTEK Wins CNCF End User Case Study Contest for Scalable AI Infrastructure Breakthroughs with Volcano](https://www.cncf.io/announcements/2025/06/09/iflytek-wins-cncf-end-user-case-study-contest-for-scalable-ai-infrastructure-breakthroughs-with-volcano/)
- [Why Spark chooses Volcano as built-in batch scheduler on Kubernetes?](https://www.cncf.io/blog/2022/06/30/why-spark-chooses-volcano-as-built-in-batch-scheduler-on-kubernetes/)
- [ING Bank: How Volcano empowers its big data analytics platform](https://www.cncf.io/blog/2023/02/21/ing-bank-how-volcano-empowers-its-big-data-analytics-platform/)
- [Using Volcano as a custom scheduler for Apache Spark on Amazon EMR on EKS](https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/tutorial-volcano.html)
- [Deploy Azure Machine Learning extension on AKS or Arc Kubernetes cluster](https://learn.microsoft.com/en-us/azure/machine-learning/how-to-deploy-kubernetes-extension?view=azureml-api-2&tabs=deploy-extension-with-cli)
- [Practical Tips for Preventing GPU Fragmentation for Volcano Scheduler](https://developer.nvidia.com/blog/practical-tips-for-preventing-gpu-fragmentation-for-volcano-scheduler/)
- [Using Volcano in Large-Scale, Distributed Offline Computing](https://volcano.sh/en/blog/ruitian2-en/)
- [OpenI-Octopus: How to Avoid Resource Preemption in Kubernetes Clusters](https://volcano.sh/en/blog/pengcheng-en/)
- [How Does Volcano Empower a Content Recommendation Engine in Xiaohongshu](https://volcano.sh/en/blog/xiaohongshu-en/)
- [How Ruitian Used Volcano to Run Large-Scale Offline HPC Jobs](https://volcano.sh/en/blog/ruitian-en/)
- [Integrating Volcano into the Leinao Cloud OS](https://volcano.sh/en/blog/leinao-en/)
- [HPC on Volcano: How Containers Support HPC Applications in the Meteorological Industry](https://volcano.sh/en/blog/hpc-en/)
- [iQIYI:Volcano-based Cloud Native Migration Practices](https://volcano.sh/en/blog/aiqiyi-en/)
- [PaddlePaddle Distributed Training on Volcano](https://volcano.sh/en/blog/paddlepaddle-en/)

## Quick Start Guide

### Prerequisites

- Kubernetes 1.12+ with CRD support


You can try Volcano by one of the following two ways.

> [!NOTE]
> * For Kubernetes v1.17 and above, use CRDs under config/crd/bases (recommended)
> * For Kubernetes v1.16 and below, use CRDs under config/crd/v1beta1 (deprecated)

### Install with YAML files

Install Volcano on an existing Kubernetes cluster. This way is both available for x86_64 and arm64 architecture.

```
kubectl apply -f https://raw.githubusercontent.com/volcano-sh/volcano/master/installer/volcano-development.yaml
```

Enjoy! Volcano will create the following resources in `volcano-system` namespace.


```
NAME                                       READY   STATUS      RESTARTS   AGE
pod/volcano-admission-5bd5756f79-dnr4l     1/1     Running     0          96s
pod/volcano-controllers-687948d9c8-nw4b4   1/1     Running     0          96s
pod/volcano-scheduler-94998fc64-4z8kh      1/1     Running     0          96s

NAME                                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/volcano-admission-service   ClusterIP   10.98.152.108   <none>        443/TCP   96s

NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/volcano-admission     1/1     1            1           96s
deployment.apps/volcano-controllers   1/1     1            1           96s
deployment.apps/volcano-scheduler     1/1     1            1           96s

NAME                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/volcano-admission-5bd5756f79     1         1         1       96s
replicaset.apps/volcano-controllers-687948d9c8   1         1         1       96s
replicaset.apps/volcano-scheduler-94998fc64      1         1         1       96s

NAME                               COMPLETIONS   DURATION   AGE
job.batch/volcano-admission-init   1/1           48s        96s

```

### Install via helm

To install official release, please visit [helm-charts](https://github.com/volcano-sh/helm-charts) for details.

```bash
helm repo add volcano-sh https://volcano-sh.github.io/helm-charts
helm install volcano volcano-sh/volcano -n volcano-system --create-namespace
```

Install from source code for developers:

```bash
helm install volcano installer/helm/chart/volcano --namespace volcano-system --create-namespace

# list helm release
helm list -n volcano-system
```

### Install from code

If you don't have a kubernetes cluster, try one-click install from code base:

```bash
./hack/local-up-volcano.sh
```

### Install volcano agent

Please follow the guide [Volcano Agent](https://volcano.sh/en/docs/colocation) to install volcano agent.

### Install monitoring system

If you want to get prometheus and grafana volcano dashboard after volcano installed, try following commands:

```bash
kubectl create -f installer/volcano-monitoring.yaml
```

### Install dashboard

Please follow the guide [Volcano Dashboard](https://github.com/volcano-sh/dashboard#volcano-dashboard) to install volcano dashboard.

## HCU/DCU vHCU 资源前缀（resource_prefix）使用说明

本仓库在 volcano 调度器中新增了 **Hygon HCU/DCU vHCU 虚拟化调度支持**，通过 `deviceshare.ResourcePrefix` 配置项实现资源名前缀可配置，使调度器可适配 **hcu、dcu、*cu、acu、bcu、ccu** 等任意含 `cu` 的产品族前缀。

### 1. 配置项

调度器配置文件（ConfigMap `volcano-scheduler-configmap`）中 deviceshare 插件的参数：

```yaml
tiers:
- plugins:
  - name: deviceshare
    arguments:
      deviceshare.VHCUEnable: true          # 启用 vHCU 虚拟化调度
      deviceshare.SchedulePolicy: binpack   # 装箱策略（可选 spread）
      deviceshare.ResourcePrefix: hcu       # 资源名前缀，见下表
```

**前缀校验规则**：任意包含 `cu` 的小写前缀均为合法值（`hcu`、`*cu`、`dcu`、`acu`、`bcu`、`ccu` 等）；不含 `cu` 的值（如 `gpu`）为非法值，调度器启动时直接拒绝（`Fatal: invalid deviceshare.ResourcePrefix: ... want a product family containing "cu"`）。取值必须与 HCU Device Plugin 的 `RESOURCE_PREFIX` 环境变量保持一致。

### 2. 前缀对应关系

| `deviceshare.ResourcePrefix` | 扩展资源名 | Pod 分配注解 | 节点注册注解 |
| ---------------------------- | --------------------------------------- | ---------------------------------------------- | ------------------------ |
| `hcu` | `hygon.com/hcunum`、`hygon.com/hcucores`、`hygon.com/hcumem` | `hami.io/hcu-devices-to-allocate`、`hami.io/hcu-devices-allocated` | `hami.io/node-hcu-register`、`hami.io/node-handshake-hcu` |
| `dcu` | `hygon.com/dcunum`、`hygon.com/dcucores`、`hygon.com/dcumem` | `hami.io/dcu-devices-to-allocate`、`hami.io/dcu-devices-allocated` | `hami.io/node-dcu-register`、`hami.io/node-handshake-dcu` |
| 其他含 `cu` 前缀（`acu`/`bcu`/`ccu`） | `hygon.com/<前缀>num` 等 | `hami.io/<前缀>-devices-*` | `hami.io/node-<前缀>-register` 等 |

> **注意**：`*cu` 前缀可通过调度器校验，但 `*` 不是 Kubernetes 资源名/注解名的合法字符（仅允许字母数字及 `-`、`_`、`.`），设备插件的节点注册注解与 Pod 资源请求会被 apiserver 拒绝。若需支持旧集群 `*cu` 工作负载，设备插件侧需做资源名映射。

### 3. 切换前缀（调度器 + 设备插件两处同步修改）

```bash
# ① 调度器：修改 ConfigMap 前缀并重启
kubectl -n volcano-system get cm volcano-scheduler-configmap -o yaml \
  | sed 's/deviceshare.ResourcePrefix: .*/deviceshare.ResourcePrefix: dcu/' | kubectl apply -f -
kubectl -n volcano-system rollout restart deploy/volcano-scheduler

# ② 设备插件：修改 RESOURCE_PREFIX 环境变量（节点资源随之重注册）
kubectl -n kube-system set env ds/hcu-device-plugin-hami RESOURCE_PREFIX=dcu

# ③ 验证
kubectl -n volcano-system logs deploy/volcano-scheduler | grep "RESOURCE_PREFIX=" | tail -1
# 预期：RESOURCE_PREFIX=dcu (count=hygon.com/dcunum core=hygon.com/dcucores memory=hygon.com/dcumem)
kubectl get node <节点> -o jsonpath='{.status.allocatable}' | tr ',' '\n' | grep dcunum
# 预期：dcunum=4、dcucores=100、dcumem=65520（数值随物理卡配置）
```

### 4. 工作负载 YAML 示例

**虚拟卡（HCU，申请 1 个 vHCU 单元 + 2048MB 显存 + 30% 算力）：**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: vhcu-demo
spec:
  schedulerName: volcano
  containers:
  - name: app
    image: <your-image>
    resources:
      limits:
        hygon.com/hcunum: 1      # vHCU 单元数
        hygon.com/hcumem: 2048   # 显存（MB）
        hygon.com/hcucores: 30   # 算力（百分比）
```

**虚拟卡（DCU，同结构换前缀资源名）：**

```yaml
    resources:
      limits:
        hygon.com/dcunum: 1
        hygon.com/dcumem: 2048
        hygon.com/dcucores: 30
```

**整卡独占（100% 算力 + 全部显存触发独占；`<前缀>num` 语义为物理卡数量）：**

```yaml
    resources:
      limits:
        hygon.com/hcunum: 1      # 物理卡数量（单卡节点填 1）
        hygon.com/hcumem: 65520  # 全部显存（MB）
        hygon.com/hcucores: 100  # 100% 算力 → 独占
```

### 5. 实测结论（129 环境，K100_AI 单卡）

| 测试项 | 结果 |
| --- | --- |
| 前缀矩阵 hcu/dcu/acu/bcu/ccu | ✅ 各前缀下虚拟卡 Pod 均正常调度运行（调度器前缀日志、节点资源、分配注解全部正确） |
| gpu 非法前缀 | ✅ 调度器启动即 Fatal 拒绝 |
| 虚拟卡/整卡/共享/VolcanoJob 用例（hcu） | ✅ 全部通过（虚拟卡 30%→36CU/2GB 精确匹配；整卡独占反证 Pending；20+20 同卡共享；2 副本成组） |
| 训练冒烟（hcu + dcu，虚拟卡 + 整卡） | ✅ 4 组全部 TRAIN_OK（虚拟卡 ~129 ms/step、整卡 ~50 ms/step，fp16 matmul 35/104 TFLOPS，算力限制与显存限制均真实生效） |

更多细节见 [docs/hcu/README.md](docs/hcu/README.md) 与示例 [docs/hcu/vhcu-demo.yaml](docs/hcu/vhcu-demo.yaml)。

This way is only available for x86_64 temporarily.
## Kubernetes compatibility
|                       | Kubernetes 1.36 | Kubernetes 1.35 | Kubernetes 1.34 | Kubernetes 1.33 | Kubernetes 1.32 | Kubernetes 1.31 | Kubernetes 1.30 | Kubernetes 1.29 | Kubernetes 1.28 | Kubernetes 1.27 | Kubernetes 1.26 | Kubernetes 1.25 | Kubernetes 1.24 | Kubernetes 1.23 | Kubernetes 1.22 | Kubernetes 1.21 |
|-----------------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|-----------------|
| Volcano HEAD (master) | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | -               | -               | -               |
| Volcano v1.15         | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | -               | -               | -               |
| Volcano v1.14         | -               | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | -               | -               |
| Volcano v1.13         | -               | -               | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | -               | -               |
| Volcano v1.12         | -               | -               | -               | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               |
| Volcano v1.11         | -               | -               | -               | -               | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               |
| Volcano v1.10         | -               | -               | -               | -               | -               | -               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               | ✓               |

Key:
* `✓` Volcano and the Kubernetes version are exactly compatible.
* `+` Volcano has features or API objects that may not be present in the Kubernetes version.
* `-` The Kubernetes version has features or API objects that Volcano can't use.


## Citing Volcano

If Volcano helps your research, we appreciate your citations. Here is the BibTeX entry:

```bibtex
@misc{volcano2026,
  title={Volcano: A Cloud Native Batch System},
  author={Klaus Ma and Kevin Wang and others},
  year={2026},
  howpublished={\url{https://github.com/volcano-sh/volcano}},
}
```

## Meeting

We hold community meetings for different timezones. See [Volcano Community Meeting Info](https://github.com/volcano-sh/community#community-meeting).

Resources:
- [Meeting notes, agenda, and the meeting link](https://docs.google.com/document/d/1YLbF8zjZBiR9PbXQPB22iuc_L0Oui5A1lddVfRnZrqs/edit)
- [Meeting Calendar](https://calendar.google.com/calendar/b/1/embed?src=volcano.sh.bot@gmail.com) | [Subscribe](https://calendar.google.com/calendar/b/1?cid=dm9sY2Fuby5zaC5ib3RAZ21haWwuY29t)

## Contact

If you have any question, feel free to reach out to us in the following ways:

> [!TIP]
> New to the Volcano community? First, join the CNCF Slack workspace using the **Join CNCF Slack** link below. After signing in, search for the `#volcano` channel in the CNCF Slack workspace and join it to participate in community discussions.

**CNCF Slack Workspace Signup:** [Domain Signup](https://cloud-native.slack.com/signup#/domain-signup) or [Community Inviter](https://communityinviter.com/apps/cloud-native/cncf)

**Volcano Slack Channel:** [#volcano](https://cloud-native.slack.com/archives/C011GJDQS0N)

**Mailing List:** [volcano-sh Google Group](https://groups.google.com/forum/#!forum/volcano-sh)

**WeChat:** Please add WeChat account `k8s2222` and request an invitation to the group chat.
