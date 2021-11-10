// Copyright (c) 2017-2021 Ingram Micro Inc.

package agentsecret

import "testing"

func TestRetrieveSecret(t *testing.T) {
	RetrieveSecretVersionMocked(t, "616d42ab704af004a4976917", "/tmp/616d42ab704af004a4976917")
	RetrieveSecretVersionFailErrMocked(t, "616d430a704af004a4976918", "/tmp/616d430a704af004a4976918")
}
