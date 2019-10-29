/*
Copyright The Stash Authors.

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
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	store "kmodules.xyz/objectstore-api/api/v1"
)

const (
	ResourceKindRepository     = "Repository"
	ResourcePluralRepository   = "repositories"
	ResourceSingularRepository = "repository"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=repositories,singular=repository,shortName=repo,categories={stash,appscode}
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Integrity",type="boolean",JSONPath=".status.integrity"
// +kubebuilder:printcolumn:name="Size",type="string",JSONPath=".status.size"
// +kubebuilder:printcolumn:name="Snapshot-Count",type="integer",JSONPath=".status.snapshotCount"
// +kubebuilder:printcolumn:name="Last-Successful-Backup",type="date",format="date-time",JSONPath=".status.lastBackupTime"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec,omitempty"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

type RepositorySpec struct {
	// Backend specify the storage where backed up snapshot will be stored
	Backend store.Backend `json:"backend,omitempty"`
	// If true, delete respective restic repository
	// +optional
	WipeOut bool `json:"wipeOut,omitempty"`
}

type RepositoryStatus struct {
	// ObservedGeneration is the most recent generation observed for this Repository. It corresponds to the
	// Repository's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// FirstBackupTime indicates the timestamp when the first backup was taken
	FirstBackupTime *metav1.Time `json:"firstBackupTime,omitempty"`
	// LastBackupTime indicates the timestamp when the latest backup was taken
	LastBackupTime *metav1.Time `json:"lastBackupTime,omitempty"`
	// Integrity shows result of repository integrity check after last backup
	Integrity *bool `json:"integrity,omitempty"`
	// Size show size of repository after last backup
	Size string `json:"size,omitempty"`
	// SnapshotCount shows number of snapshots stored in the repository
	SnapshotCount int `json:"snapshotCount,omitempty"`
	// SnapshotsRemovedOnLastCleanup shows number of old snapshots cleaned up according to retention policy on last backup session
	SnapshotsRemovedOnLastCleanup int `json:"snapshotsRemovedOnLastCleanup,omitempty"`

	// Deprecated
	LastSuccessfulBackupTime *metav1.Time `json:"lastSuccessfulBackupTime,omitempty"`
	// Deprecated
	LastBackupDuration string `json:"lastBackupDuration,omitempty"`
	// Deprecated
	BackupCount int64 `json:"backupCount,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items,omitempty"`
}
