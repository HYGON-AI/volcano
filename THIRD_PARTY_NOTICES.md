# Third-Party Notices / 第三方来源清单

本文件登记本仓库（Volcano HCU 二次开发分支）相对上游基线的来源、版权、许可证与 HYGON 修改说明。
仓库根目录许可证见 [LICENSE](LICENSE)（Apache-2.0）。上游归属见 [README.md](README.md#upstream-attribution--上游归属)。

## 1. Upstream / 上游项目

| 字段 | 内容 |
| --- | --- |
| 项目 | Volcano |
| 仓库 URL | https://github.com/volcano-sh/volcano.git |
| 上游分支 | `release-1.15` |
| 固定 Tag | `v1.15.2` |
| 固定 Commit | `1462fb7b4835970708717456e3aed85e697ec2eb` |
| Copyright | The Volcano Authors |
| 许可证 | Apache-2.0 |
| 本地路径 | 仓库根目录（fork 全树） |
| HYGON 修改 | 新增 Hygon HCU/vHCU `deviceshare` 调度支持、相关配置与文档；对既有调度插件做接入性修改 |

## 2. HYGON 新增原创源码（非第三方）

以下文件由 HYGON 在本 fork 中新增，适用 Apache-2.0，文件头使用 H1（HYGON Copyright + SPDX）。
实现参考了同仓库上游 `pkg/scheduler/api/devices/nvidia/vgpu` 的 deviceshare 设备模型，但作为 HYGON HCU 适配原创实现维护，**不作为独立第三方组件引入**。

| 本地路径 | Copyright | 许可证 | HYGON 修改 |
| --- | --- | --- | --- |
| `pkg/scheduler/api/devices/config/vhcu.go` | Copyright (c) 2026 Hygon Information Technology Co., Ltd. | Apache-2.0 | HYGON 原创新增 |
| `pkg/scheduler/api/devices/hygon/device_info.go` | Copyright (c) 2026 Hygon Information Technology Co., Ltd. | Apache-2.0 | HYGON 原创新增 |
| `pkg/scheduler/api/devices/hygon/metrics.go` | Copyright (c) 2026 Hygon Information Technology Co., Ltd. | Apache-2.0 | HYGON 原创新增 |
| `pkg/scheduler/api/devices/hygon/type.go` | Copyright (c) 2026 Hygon Information Technology Co., Ltd. | Apache-2.0 | HYGON 原创新增 |
| `pkg/scheduler/api/devices/hygon/utils.go` | Copyright (c) 2026 Hygon Information Technology Co., Ltd. | Apache-2.0 | HYGON 原创新增 |

配套文档与安装配置（非源码头文件义务范围）：

- `docs/hcu/README.md`
- `docs/hcu/vhcu-demo.yaml`
- `installer/helm/chart/volcano/config/volcano-scheduler.conf`
- `installer/volcano-development.yaml`
- `installer/volcano-development-vap.yaml`
- `installer/volcano-development-vap-map.yaml`

## 3. HYGON 修改的上游源码

保留上游 Volcano Authors 原声明，并按 Apache-2.0 增加 HYGON 修改说明（实质修改追加 HYGON Copyright + SPDX）。

| 本地路径 | 性质 | 模板 |
| --- | --- | --- |
| `pkg/scheduler/api/devices/config/config.go` | 实质修改：接入 HygonConfig 默认值 | H2 |
| `pkg/scheduler/api/devices/util.go` | 实质修改：资源请求算力补齐逻辑 | H2 |
| `pkg/scheduler/api/node_info.go` | 实质修改：节点 HCU 设备注册与资源加减 | H2 |
| `pkg/scheduler/plugins/deviceshare/deviceshare.go` | 实质修改：VHCUEnable 与会话内设备刷新 | H2 |
| `pkg/scheduler/plugins/util/nodelock/nodelock.go` | 实质修改：HCU 节点锁 API | H2 |
| `pkg/scheduler/api/shared_device_pool.go` | 非实质修改：接口编译期注册 | H3 |

## 4. 上游附带许可证文本（待法务/合规确认）

下列路径来自上游 Volcano 基线树中的 `licenses/` 与许可证扫描配置，**本 HCU 分支未改写其许可证文本**。
自动准入范围仅为 MIT、BSD-3-Clause、Apache-2.0；以下组件需在公开发布前完成法务/合规审批，批准前阻断对外发布。

### 4.1 `licenses/github.com/davecgh/go-spew/LICENSE`

| 字段 | 内容 |
| --- | --- |
| 项目 | github.com/davecgh/go-spew |
| 仓库 URL | https://github.com/davecgh/go-spew |
| 固定版本 | go.mod 间接依赖 `v1.1.2-0.20180830191138-d8f796af33cc` |
| Copyright | Copyright (c) 2012-2016 Dave Collins \<dave@davec.name\> |
| 许可证 | ISC（文件正文为 ISC License） |
| 本地路径 | `licenses/github.com/davecgh/go-spew/LICENSE` |
| HYGON 修改 | 无（上游原样） |
| 合规状态 | ISC 不在当前自动准入列表；需法务/合规审批 |

### 4.2 `licenses/github.com/hashicorp/errwrap/LICENSE`

| 字段 | 内容 |
| --- | --- |
| 项目 | github.com/hashicorp/errwrap |
| 仓库 URL | https://github.com/hashicorp/errwrap |
| 固定版本 | go.mod 间接依赖 `v1.1.0` |
| 许可证 | MPL-2.0（文件正文为 Mozilla Public License 2.0） |
| 本地路径 | `licenses/github.com/hashicorp/errwrap/LICENSE` |
| HYGON 修改 | 无（上游原样） |
| 合规状态 | MPL-2.0 不在自动准入列表；需法务/合规审批。扫描器曾提示 GPL-3.0/MPL 复合风险，以文件正文 MPL-2.0 为准提交审批材料 |

### 4.3 `licenses/github.com/hashicorp/go-multierror/LICENSE`

| 字段 | 内容 |
| --- | --- |
| 项目 | github.com/hashicorp/go-multierror |
| 仓库 URL | https://github.com/hashicorp/go-multierror |
| 固定版本 | go.mod 直接依赖 `v1.1.1` |
| 许可证 | MPL-2.0（文件正文为 Mozilla Public License 2.0） |
| 本地路径 | `licenses/github.com/hashicorp/go-multierror/LICENSE` |
| HYGON 修改 | 无（上游原样） |
| 合规状态 | MPL-2.0 不在自动准入列表；需法务/合规审批。扫描器曾提示 GPL-3.0/MPL 复合风险，以文件正文 MPL-2.0 为准提交审批材料 |

### 4.4 `licenses/k8s.io/kubernetes/logo/LICENSE`

| 字段 | 内容 |
| --- | --- |
| 项目 | k8s.io/kubernetes logo assets |
| 仓库 URL | https://github.com/kubernetes/kubernetes |
| 许可证 | 文件声明可在 Apache-2.0 或 CC-BY-4.0 中选择 |
| 本地路径 | `licenses/k8s.io/kubernetes/logo/LICENSE` |
| HYGON 修改 | 无（上游原样） |
| 合规状态 | 双许可表述导致自动识别失败；需法务确认实际使用范围并审批 |

### 4.5 `config/license-lint.yaml`

| 字段 | 内容 |
| --- | --- |
| 说明 | 上游 Volcano 的许可证 lint 策略配置（列有 unrestricted / reciprocal / restricted 许可证名），**不是**某个第三方组件的 LICENSE 正文 |
| 本地路径 | `config/license-lint.yaml` |
| HYGON 修改 | 无（上游原样） |
| 合规状态 | 配置内出现 GPL/MPL 等名称仅表示策略分类；需法务确认该配置本身不构成额外许可证授予，并与真实依赖审批结论一致 |

## 5. 备注

- 平台命名（HCU/DCU/AMD/XGMI）与运行时可见文案不在本开源合规清单范围，需另行执行 `audit-hygon-platform`。
- 未在本文件登记的上游未改动路径，沿用上游 Volcano 原有版权与 Apache-2.0 许可。
