// Copyright Red Hat

package helpers

import "fmt"

func ManagedClusterSetNameForWorkspace(workspaceName string) string {
	// For now, workspaces are uniquely identified by their name. This may change.
	return workspaceName
}

func GetSyncerPrefix() string {
	return "kcp-syncer"
}

func GetSyncerName(regClusterName string) string { // Should be passing in the SyncTarget
	//TODO - Adjust to match https://github.com/robinbobbitt/kcp/blob/b6314f86a563a354eddde44f1a7038042090df9e/pkg/cliplugins/workload/plugin/sync.go#L141 once we have SyncTarget
	return fmt.Sprintf("%s-%s", GetSyncerPrefix(), regClusterName)
}

func GetSyncerServiceAccountName() string {
	return fmt.Sprintf("kcp-syncer-sa")
}
