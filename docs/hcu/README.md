# Volcano中HCU虚拟化支持安装部署说明

本文档基于 Volcano 官方vGPU虚拟化机制，并结合 HCU虚拟化用法，说明在 Volcano 中实现 HCU 虚拟化支持后的完整安装与部署流程，适用于 AI / HPC 混合算力资源调度场景。

## 1. 前提条件

| 项目             | 要求                  |
| ---------------- |---------------------|
| Kubernetes       | \>= 1.20            |
| 容器运行时       | containerd / Docker |
| HCU 驱动         | \>=6.3.8            |

节点需提前完成 HCU 驱动与运行时环境安装，并确认以下命令可正常执行：

```bash
hy-smi
hy-smi virtual -show-device-info
```

---

## 2. 架构说明

整体架构复用 Volcano GPU 虚拟化模型，通过 deviceshare 插件扩展支持 HCU 虚拟化：

```
+------------------+
|   Volcano Job    |
+--------+---------+
         |
         v
+------------------+
| Volcano Scheduler|
|   deviceshare    |
+--------+---------+
         |
         v
+-------------------+
| HCU Device Plugin   |
+--------+----------+
         |
         v
+--------------------------+
| HCU Driver + Runtime     |
+--------------------------+
```

---

## 3. 环境准备

### 3.1 节点打标签

为包含 HCU 的节点打标签：

```bash
kubectl label node <node-name> hygon.com/hcu=true --overwrite
```

---

## 4. 安装 Volcano

```bash
git clone https://github.com/volcano-sh/volcano.git
cd volcano
kubectl apply -f manifests/
```

验证：

```bash
kubectl get pods -n volcano-system
```

所有组件应处于 `Running` 状态。

## 5. 安装 HAMi 模式 HCU-Device-Plugin

```bash
kubectl apply -f k8s-hcu-plugin-hami.yaml
```

检查插件状态：

```bash
kubectl get pods -n kube-system | grep hcu-device-plugin-hami
```

---

## 6. 配置 Volcano 调度器支持 HCU 虚拟化

编辑 Volcano Scheduler 配置：

```bash
kubectl edit cm volcano-scheduler-configmap -n volcano-system
```

示例配置：

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: volcano-scheduler-configmap
  namespace: volcano-system
data:
  volcano-scheduler.conf: |
    actions: "enqueue, allocate, backfill"
    tiers:
    - plugins:
      - name: gang
      - name: priority
      - name: predicates
      - name: deviceshare
        arguments:
          deviceshare.VHCUEnable: true
          deviceshare.SchedulePolicy: binpack
```

重启 scheduler 使配置生效：

```bash
kubectl delete pod -n volcano-system -l app=volcano-scheduler
```

---

## 7. 验证资源注册

```bash
kubectl get node <node-name> -o yaml | grep -A20 allocatable
```

期望看到类似输出：

```yaml
allocatable:
  hygon.com/hcunum: "8"
  hygon.com/hcumem: "512"
  hygon.com/hcucores: "800"
```
hygon.com/hcunum 标识可用于vHCU切分的最大数量，hygon.com/hcumem 标识可用于vHCU切分的最大显存（按照GiB单位计数），hygon.com/hcucores 标识可用于vHCU切分的最大算力（按照百分比计数）。
## 8. vHCU 作业示例

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: vhcu-demo
spec:
  schedulerName: volcano
  containers:
  - name: hcu-app
    image: your-hcu-runtime-image
    command: ["bash", "-c", "hy-smi virtual -show-device-info && sleep 3600"]
    resources:
      limits:
        hygon.com/hcunum: 1     # 申请1个vHCU
        hygon.com/hcumem: 2     # 指定vHCU显存为2GiB
        hygon.com/hcucores: 50  # 指定vHCU算力为整卡的50%
```
提交：

```bash
kubectl apply -f vhcu-demo.yaml
```

容器内验证：

```bash
hy-smi virtual -show-device-info
```

输出示例：

```
Device 0:
  Actual Device: 1
  Compute Units: 60
  Global Memory: 2147483648 bytes
```

表示 vHCU 虚拟化生效。

---

## 9. 常见问题与注意事项

### 9.1 HCU 虚拟化限制

- 每个容器最多申请 **1 张 vHCU**
- 不支持在 `initContainer` 中使用 vHCU
- 不支持 HCU MIG 与 vHCU 同时开启

---

### 9.2 调度策略建议

- 推荐使用 `binpack` 提高单卡利用率
- Volcano gang scheduling 可与 vHCU 混合使用

---

### 9.3 资源未注册排查

```bash
kubectl logs -n kube-system <hcu-device-plugin-pod>
hy-smi
```

确认驱动、插件与内核版本匹配。