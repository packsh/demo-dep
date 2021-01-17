/*
Copyright The Kubepack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

const (
	BundleRepoURL = "https://kubernetes-charts.storage.googleapis.com"

	StripeAPIKey = "STRIPE_API_KEY"

	StripeMetadataKeyUserID = "client_id"
)

const (
	LabelPlanID         = "byte.builders/plan-id"
	LabelPlanPhase      = "byte.builders/plan-phase"
	LabelProductID      = "byte.builders/product-id"
	LabelProductKey     = "byte.builders/product-key"
	LabelProductPhase   = "byte.builders/product-phase"
	LabelProductOwnerID = "byte.builders/product-owner-id"
)

const (
	LabelChartURL           = "helm.sh/chart-url"
	LabelChartName          = "helm.sh/chart-name"
	LabelChartVersion       = "helm.sh/chart-version"
	LabelChartFirstDeployed = "helm.sh/first-deployed"
	LabelChartLastDeployed  = "helm.sh/last-deployed"
)

const (
	YAMLHost   = "https://pkg.bytebuilders.xyz"
	YAMLBucket = "gs://pkg.bytebuilders.xyz"
)
