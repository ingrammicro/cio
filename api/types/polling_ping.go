// Copyright (c) 2017-2021 Ingram Micro Inc.

package types

type PollingPing struct {
	PendingCommands bool `json:"pending_commands" header:"PENDING_COMMANDS"`
}
