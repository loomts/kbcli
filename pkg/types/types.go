/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package types

import (
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	workloads "github.com/apecloud/kubeblocks/apis/workloads/v1alpha1"
)

const (
	// CliDefaultHome defines kbcli default home name
	CliDefaultHome = ".kbcli"

	// CliClusterTypeConfigs defines kbcli cluster-type config file name
	CliClusterTypeConfigs = "cluster_types"

	//	CliChartsCache defines kbcli charts cache dir name
	CliChartsCache = "charts"

	// CliLogDir defines kbcli log dir name
	CliLogDir = "logs"

	// CliHomeEnv defines kbcli home system env
	CliHomeEnv = "KBCLI_HOME"

	// DefaultLogFilePrefix is the default log file prefix
	DefaultLogFilePrefix = "kbcli"

	// AddonIndexDirEnv defines kbcli addon index dir
	AddonIndexDirEnv = "KBCLI_ADDON_INDEX_DIR"

	// DefaultIndexName defines the kbcli addon default index name
	DefaultIndexName = "kubeblocks"

	// DefaultNamespace is the namespace where kubeblocks is installed if
	// no other namespace is specified
	DefaultNamespace = "kb-system"

	// GoosLinux is os.GOOS linux string
	GoosLinux = "linux"
	// GoosDarwin is os.GOOS darwin string
	GoosDarwin = "darwin"
	// GoosWindows is os.GOOS windows string
	GoosWindows = "windows"
)

// K8s core API group
const (
	K8sCoreAPIVersion    = "v1"
	ResourceDeployments  = "deployments"
	ResourceConfigmaps   = "configmaps"
	ResourceStatefulSets = "statefulsets"
	ResourceDaemonSets   = "daemonsets"
	ResourceSecrets      = "secrets"
)

// K8s batch API group
const (
	K8SBatchAPIGroup   = batchv1.GroupName
	K8sBatchAPIVersion = "v1"
	ResourceJobs       = "jobs"
	ResourceCronJobs   = "cronjobs"
)

// K8s webhook API group
const (
	WebhookAPIGroup                         = "admissionregistration.k8s.io"
	K8sWebhookAPIVersion                    = "v1"
	ResourceValidatingWebhookConfigurations = "validatingwebhookconfigurations"
	ResourceMutatingWebhookConfigurations   = "mutatingwebhookconfigurations"
)

// Apps API group
const (
	AppsAPIGroup                        = "apps.kubeblocks.io"
	AppsAPIVersion                      = "v1alpha1"
	AppsAPIBetaVersion                  = "v1beta1"
	ResourcePods                        = "pods"
	ResourceClusters                    = "clusters"
	ResourceClusterDefs                 = "clusterdefinitions"
	ResourceClusterVersions             = "clusterversions"
	ResourceComponentDefs               = "componentdefinitions"
	ResourceComponents                  = "components"
	ResourceOpsRequests                 = "opsrequests"
	ResourceOpsDefinitions              = "opsdefinitions"
	ResourceConfigConstraintVersions    = "configconstraints"
	ResourceConfigurationVersions       = "configurations"
	ResourceComponentResourceConstraint = "componentresourceconstraints"
	ResourceComponentClassDefinition    = "componentclassdefinitions"
	KindCluster                         = "Cluster"
	KindComponentClassDefinition        = "ComponentClassDefinition"
	KindClusterDef                      = "ClusterDefinition"
	KindClusterVersion                  = "ClusterVersion"
	KindConfigConstraint                = "ConfigConstraint"
	KindConfiguration                   = "Configuration"
	KindBackup                          = "Backup"
	KindRestore                         = "Restore"
	KindBackupPolicy                    = "BackupPolicy"
	KindOps                             = "OpsRequest"
	KindBackupSchedule                  = "BackupSchedule"
	KindBackupPolicyTemplate            = "BackupPolicyTemplate"
	KindStatefulSet                     = "StatefulSet"
	KindDeployment                      = "Deployment"
	KindRSM                             = "ReplicatedStateMachine"
	KindConfigMap                       = "ConfigMap"
)

// K8S rbac API group
const (
	RBACAPIGroup        = rbacv1.GroupName
	RBACAPIVersion      = "v1"
	ClusterRoles        = "clusterroles"
	ClusterRoleBindings = "clusterrolebindings"
	Roles               = "roles"
	RoleBindings        = "rolebindings"
	ServiceAccounts     = "serviceaccounts"
)

// Annotations
const (
	ServiceHAVIPTypeAnnotationKey   = "service.kubernetes.io/kubeblocks-havip-type"
	ServiceHAVIPTypeAnnotationValue = "private-ip"
	ServiceFloatingIPAnnotationKey  = "service.kubernetes.io/kubeblocks-havip-floating-ip"

	ReloadConfigMapAnnotationKey = "kubeblocks.io/reload-configmap" // mark an annotation to load configmap

	KBVersionValidateAnnotationKey = "addon.kubeblocks.io/kubeblocks-version"
)

// Labels
const (
	AddonProviderLabelKey = "addon.kubeblocks.io/provider"
	// ProviderLabelKey was used as the label for addon providers before version 0.8.0
	ProviderLabelKey     = "kubeblocks.io/provider"
	AddonVersionLabelKey = "addon.kubeblocks.io/version"
	AddonNameLabelKey    = "addon.kubeblocks.io/name"
	AddonModelLabelKey   = "addon.kubeblocks.io/model"
)

// DataProtection API group
const (
	DPAPIGroup              = "dataprotection.kubeblocks.io"
	DPAPIVersion            = "v1alpha1"
	ResourceBackups         = "backups"
	ResourceActionSets      = "actionsets"
	ResourceRestores        = "restores"
	ResourceBackupPolicies  = "backuppolicies"
	ResourceBackupRepos     = "backuprepos"
	ResourceBackupSchedules = "backupschedules"
	ResourceBackupTemplates = "backuppolicytemplates"
)

// Extensions API group
const (
	ExtensionsAPIGroup   = "extensions.kubeblocks.io"
	ExtensionsAPIVersion = "v1alpha1"
	ResourceAddons       = "addons"
)

// Storage API group
const (
	StorageAPIGroup          = "storage.kubeblocks.io"
	StorageAPIVersion        = "v1alpha1"
	ResourceStorageProviders = "storageproviders"
)

// Migration API group
const (
	MigrationAPIGroup          = "datamigration.apecloud.io"
	MigrationAPIVersion        = "v1alpha1"
	ResourceMigrationTasks     = "migrationtasks"
	ResourceMigrationTemplates = "migrationtemplates"
)

// Crd Api group
const (
	CustomResourceDefinitionAPIGroup   = "apiextensions.k8s.io"
	CustomResourceDefinitionAPIVersion = "v1"
	ResourceCustomResourceDefinition   = "customresourcedefinitions"
)

// Kubebench API group
const (
	KubebenchAPIGroup   = "benchmark.apecloud.io"
	KubebenchAPIVersion = "v1alpha1"
	ResourcePgBench     = "pgbenches"
	ResourceSysBench    = "sysbenches"
	ResourceYcsb        = "ycsbs"
	ResourceTpcc        = "tpccs"
	ResourceTpch        = "tpches"
	ResourceTpcds       = "tpcds"
	ResourceRedisBench  = "redisbenches"
)

// Workload API group
const (
	ResourceRSM = "replicatedstatemachines"
)

const (
	None = "<none>"

	// AddonReleasePrefix is the prefix of addon release name
	AddonReleasePrefix = "kb-addon"

	// AddonResourceNamePrefix is the prefix for the names of all K8s resources rendered by the addon.
	AddonResourceNamePrefix = "resourceNamePrefix"
	// ImageRegistryKey is the image registry key in KB resource helm chart values.yaml
	ImageRegistryKey = "image.registry"
)

// Migrate some const from kubeblocks to kbcli
const (
	// KBDefaultClusterVersionAnnotationKey specifies the default cluster version.
	KBDefaultClusterVersionAnnotationKey = "kubeblocks.io/is-default-cluster-version"

	// KBAddonProviderLabelKey marks the addon provider
	KBAddonProviderLabelKey = "kubeblocks.io/provider"
)

var (
	// KubeBlocksRepoName helm repo name for kubeblocks
	KubeBlocksRepoName = "kubeblocks"

	// KubeBlocksChartName helm chart name for kubeblocks
	KubeBlocksChartName = "kubeblocks"

	// KubeBlocksReleaseName helm release name for kubeblocks
	KubeBlocksReleaseName = "kubeblocks"

	// KubeBlocksChartURL the helm chart repo for installing kubeblocks
	KubeBlocksChartURL = "https://apecloud.github.io/helm-charts"

	// GitLabHelmChartRepo the helm chart repo in GitLab
	GitLabHelmChartRepo = "https://jihulab.com/api/v4/projects/85949/packages/helm/stable"

	// KubeBlocksHelmLabel name=kubeblocks,owner-helm, for helm secret
	KubeBlocksHelmLabel = fmt.Sprintf("%s=%s,%s=%s", "name", KubeBlocksChartName, "owner", "helm")

	// KubeBlocksManagerConfigMapName the kubeblocks manager configMap name
	KubeBlocksManagerConfigMapName = fmt.Sprintf("%s-manager-config", KubeBlocksChartName)

	// DefaultAddonIndexURL points to the upstream index.
	DefaultAddonIndexURL = "https://github.com/apecloud/block-index.git"

	// AddonIndexDir is the default addon index dir
	AddonIndexDir = filepath.Join("addon", "index")
)

// Playground
var (
	// K3dClusterName is the k3d cluster name for playground
	K3dClusterName = "kb-playground"
)

type ConfigTemplateInfo struct {
	Name  string
	TPL   appsv1alpha1.ComponentConfigSpec
	CMObj *corev1.ConfigMap
}

func PodGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: "", Version: K8sCoreAPIVersion, Resource: ResourcePods}
}

func ClusterGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceClusters}
}

func ClusterDefGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceClusterDefs}
}

func ClusterVersionGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceClusterVersions}
}

func CompDefGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceComponentDefs}
}

func ComponentGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceComponents}
}

func OpsDefinitionGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceOpsDefinitions}
}

func OpsGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceOpsRequests}
}

func BackupGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceBackups}
}

func BackupPolicyGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceBackupPolicies}
}

func BackupPolicyTemplateGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: DPAPIVersion, Resource: ResourceBackupTemplates}
}

func BackupScheduleGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceBackupSchedules}
}

func ActionSetGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceActionSets}
}

func BackupRepoGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceBackupRepos}
}

func RestoreGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: DPAPIGroup, Version: DPAPIVersion, Resource: ResourceRestores}
}

func ConfigurationGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: DPAPIVersion, Resource: ResourceConfigurationVersions}
}

func AddonGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: ExtensionsAPIGroup, Version: ExtensionsAPIVersion, Resource: ResourceAddons}
}

func StorageProviderGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: StorageAPIGroup, Version: StorageAPIVersion, Resource: ResourceStorageProviders}
}

func ComponentResourceConstraintGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceComponentResourceConstraint}
}

func ComponentClassDefinitionGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIVersion, Resource: ResourceComponentClassDefinition}
}

func CRDGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  K8sCoreAPIVersion,
		Resource: "customresourcedefinitions",
	}
}

func ConfigmapGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: ResourceConfigmaps}
}

func SecretGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: ResourceSecrets}
}

func StatefulSetGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: appsv1.GroupName, Version: K8sCoreAPIVersion, Resource: ResourceStatefulSets}
}

func RSMGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: workloads.GroupVersion.Group, Version: workloads.GroupVersion.Version, Resource: ResourceRSM}
}

func DaemonSetGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: appsv1.GroupName, Version: K8sCoreAPIVersion, Resource: ResourceDaemonSets}
}

func DeployGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: appsv1.GroupName, Version: K8sCoreAPIVersion, Resource: ResourceDeployments}
}

func ServiceGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: "services"}
}

func PVCGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: "persistentvolumeclaims"}
}

func PVGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: "persistentvolumes"}
}

func ConfigConstraintGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: AppsAPIGroup, Version: AppsAPIBetaVersion, Resource: ResourceConfigConstraintVersions}
}

func StorageClassGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "storage.k8s.io",
		Version:  K8sCoreAPIVersion,
		Resource: "storageclasses",
	}
}

func VolumeSnapshotClassGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "snapshot.storage.k8s.io",
		Version:  K8sCoreAPIVersion,
		Resource: "volumesnapshotclasses",
	}
}

func ValidatingWebhookConfigurationGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    WebhookAPIGroup,
		Version:  K8sWebhookAPIVersion,
		Resource: ResourceValidatingWebhookConfigurations,
	}
}

func MutatingWebhookConfigurationGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    WebhookAPIGroup,
		Version:  K8sWebhookAPIVersion,
		Resource: ResourceMutatingWebhookConfigurations,
	}
}

func ClusterRoleGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: RBACAPIGroup, Version: RBACAPIVersion, Resource: ClusterRoles}
}
func ClusterRoleBindingGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: RBACAPIGroup, Version: RBACAPIVersion, Resource: ClusterRoleBindings}
}

func RoleGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: RBACAPIGroup, Version: RBACAPIVersion, Resource: Roles}
}

func RoleBindingGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: RBACAPIGroup, Version: RBACAPIVersion, Resource: RoleBindings}
}

func ServiceAccountGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: corev1.GroupName, Version: K8sCoreAPIVersion, Resource: ServiceAccounts}
}

func MigrationTaskGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    MigrationAPIGroup,
		Version:  MigrationAPIVersion,
		Resource: ResourceMigrationTasks,
	}
}

func MigrationTemplateGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    MigrationAPIGroup,
		Version:  MigrationAPIVersion,
		Resource: ResourceMigrationTemplates,
	}
}

func CustomResourceDefinitionGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    CustomResourceDefinitionAPIGroup,
		Version:  CustomResourceDefinitionAPIVersion,
		Resource: ResourceCustomResourceDefinition,
	}
}

func JobGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: K8SBatchAPIGroup, Version: K8sBatchAPIVersion, Resource: ResourceJobs}
}
func CronJobGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: K8SBatchAPIGroup, Version: K8sBatchAPIVersion, Resource: ResourceCronJobs}
}

func PgBenchGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourcePgBench}
}

func SysbenchGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceSysBench}
}

func YcsbGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceYcsb}
}

func TpccGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceTpcc}
}

func TpchGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceTpch}
}

func TpcdsGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceTpcds}
}

func RedisBenchGVR() schema.GroupVersionResource {
	return schema.GroupVersionResource{Group: KubebenchAPIGroup, Version: KubebenchAPIVersion, Resource: ResourceRedisBench}
}
