// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EndpointEncryptionSpec struct {
	/* Immutable. Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource. Has the form: 'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'. The key needs to be in the same region as where the compute resource is created. */
	KmsKeyName string `json:"kmsKeyName"`
}

type VertexAIEndpointSpec struct {
	/* The description of the Endpoint. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters. */
	DisplayName string `json:"displayName"`

	/* Immutable. Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key. */
	// +optional
	EncryptionSpec *EndpointEncryptionSpec `json:"encryptionSpec,omitempty"`

	/* Immutable. The location for the resource. */
	Location string `json:"location"`

	/* Immutable. The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the Endpoint should be peered. Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network. Only one of the fields, network or enable_private_service_connect, can be set. [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): 'projects/{project}/global/networks/{network}'. Where '{project}' is a project number, as in '12345', and '{network}' is network name. */
	// +optional
	Network *string `json:"network,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type EndpointAutomaticResourcesStatus struct {
	/* The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, a no upper bound for scaling under heavy traffic will be assume, though Vertex AI may be unable to scale beyond certain replica number. */
	// +optional
	MaxReplicaCount *int `json:"maxReplicaCount,omitempty"`

	/* The minimum number of replicas this DeployedModel will be always deployed on. If traffic against it increases, it may dynamically be deployed onto more replicas up to max_replica_count, and as traffic decreases, some of these extra replicas may be freed. If the requested value is too large, the deployment will error. */
	// +optional
	MinReplicaCount *int `json:"minReplicaCount,omitempty"`
}

type EndpointAutoscalingMetricSpecsStatus struct {
	/* The resource metric name. Supported metrics: * For Online Prediction: * 'aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle' * 'aiplatform.googleapis.com/prediction/online/cpu/utilization'. */
	// +optional
	MetricName *string `json:"metricName,omitempty"`

	/* The target resource utilization in percentage (1% - 100%) for the given metric; once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%) if not provided. */
	// +optional
	Target *int `json:"target,omitempty"`
}

type EndpointDedicatedResourcesStatus struct {
	/* The metric specifications that overrides a resource utilization metric (CPU utilization, accelerator's duty cycle, and so on) target value (default to 60 if not set). At most one entry is allowed per metric. If machine_spec.accelerator_count is above 0, the autoscaling will be based on both CPU utilization and accelerator's duty cycle metrics and scale up when either metrics exceeds its target value while scale down if both metrics are under their target value. The default target value is 60 for both metrics. If machine_spec.accelerator_count is 0, the autoscaling will be based on CPU utilization metric only with default target value 60 if not explicitly set. For example, in the case of Online Prediction, if you want to override target CPU utilization to 80, you should set autoscaling_metric_specs.metric_name to 'aiplatform.googleapis.com/prediction/online/cpu/utilization' and autoscaling_metric_specs.target to '80'. */
	// +optional
	AutoscalingMetricSpecs []EndpointAutoscalingMetricSpecsStatus `json:"autoscalingMetricSpecs,omitempty"`

	/* The specification of a single machine used by the prediction. */
	// +optional
	MachineSpec []EndpointMachineSpecStatus `json:"machineSpec,omitempty"`

	/* The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, will use min_replica_count as the default value. The value of this field impacts the charge against Vertex CPU and GPU quotas. Specifically, you will be charged for max_replica_count * number of cores in the selected machine type) and (max_replica_count * number of GPUs per replica in the selected machine type). */
	// +optional
	MaxReplicaCount *int `json:"maxReplicaCount,omitempty"`

	/* The minimum number of machine replicas this DeployedModel will be always deployed on. This value must be greater than or equal to 1. If traffic against the DeployedModel increases, it may dynamically be deployed onto more replicas, and as traffic decreases, some of these extra replicas may be freed. */
	// +optional
	MinReplicaCount *int `json:"minReplicaCount,omitempty"`
}

type EndpointDeployedModelsStatus struct {
	/* A description of resources that to large degree are decided by Vertex AI, and require only a modest additional configuration. */
	// +optional
	AutomaticResources []EndpointAutomaticResourcesStatus `json:"automaticResources,omitempty"`

	/* Output only. Timestamp when the DeployedModel was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* A description of resources that are dedicated to the DeployedModel, and that need a higher degree of manual configuration. */
	// +optional
	DedicatedResources []EndpointDedicatedResourcesStatus `json:"dedicatedResources,omitempty"`

	/* The display name of the DeployedModel. If not provided upon creation, the Model's display_name is used. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* These logs are like standard server access logs, containing information like timestamp and latency for each prediction request. Note that Stackdriver logs may incur a cost, especially if your project receives prediction requests at a high queries per second rate (QPS). Estimate your costs before enabling this option. */
	// +optional
	EnableAccessLogging *bool `json:"enableAccessLogging,omitempty"`

	/* If true, the container of the DeployedModel instances will send 'stderr' and 'stdout' streams to Stackdriver Logging. Only supported for custom-trained Models and AutoML Tabular Models. */
	// +optional
	EnableContainerLogging *bool `json:"enableContainerLogging,omitempty"`

	/* The ID of the DeployedModel. If not provided upon deployment, Vertex AI will generate a value for this ID. This value should be 1-10 characters, and valid characters are /[0-9]/. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* The name of the Model that this is the deployment of. Note that the Model may be in a different location than the DeployedModel's Endpoint. */
	// +optional
	Model *string `json:"model,omitempty"`

	/* Output only. The version ID of the model that is deployed. */
	// +optional
	ModelVersionId *string `json:"modelVersionId,omitempty"`

	/* Output only. Provide paths for users to send predict/explain/health requests directly to the deployed model services running on Cloud via private services access. This field is populated if network is configured. */
	// +optional
	PrivateEndpoints []EndpointPrivateEndpointsStatus `json:"privateEndpoints,omitempty"`

	/* The service account that the DeployedModel's container runs as. Specify the email address of the service account. If this service account is not specified, the container runs as a service account that doesn't have access to the resource project. Users deploying the Model must have the 'iam.serviceAccounts.actAs' permission on this service account. */
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	/* The resource name of the shared DeploymentResourcePool to deploy on. Format: projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}. */
	// +optional
	SharedResources *string `json:"sharedResources,omitempty"`
}

type EndpointMachineSpecStatus struct {
	/* The number of accelerators to attach to the machine. */
	// +optional
	AcceleratorCount *int `json:"acceleratorCount,omitempty"`

	/* The type of accelerator(s) that may be attached to the machine as per accelerator_count. See possible values [here](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/MachineSpec#AcceleratorType). */
	// +optional
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	/* The type of the machine. See the [list of machine types supported for prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types) See the [list of machine types supported for custom training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types). For DeployedModel this field is optional, and the default value is 'n1-standard-2'. For BatchPredictionJob or as part of WorkerPoolSpec this field is required. TODO(rsurowka): Try to better unify the required vs optional. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`
}

type EndpointPrivateEndpointsStatus struct {
	/* Output only. Http(s) path to send explain requests. */
	// +optional
	ExplainHttpUri *string `json:"explainHttpUri,omitempty"`

	/* Output only. Http(s) path to send health check requests. */
	// +optional
	HealthHttpUri *string `json:"healthHttpUri,omitempty"`

	/* Output only. Http(s) path to send prediction requests. */
	// +optional
	PredictHttpUri *string `json:"predictHttpUri,omitempty"`

	/* Output only. The name of the service attachment resource. Populated if private service connect is enabled. */
	// +optional
	ServiceAttachment *string `json:"serviceAttachment,omitempty"`
}

type VertexAIEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIEndpoint's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. Timestamp when this Endpoint was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The models deployed in this Endpoint. To add or remove DeployedModels use EndpointService.DeployModel and EndpointService.UndeployModel respectively. Models can also be deployed and undeployed using the [Cloud Console](https://console.cloud.google.com/vertex-ai/). */
	// +optional
	DeployedModels []EndpointDeployedModelsStatus `json:"deployedModels,omitempty"`

	/* Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'. */
	// +optional
	ModelDeploymentMonitoringJob *string `json:"modelDeploymentMonitoringJob,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. Timestamp when this Endpoint was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIEndpoint is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIEndpointList contains a list of VertexAIEndpoint
type VertexAIEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIEndpoint{}, &VertexAIEndpointList{})
}
