// Copyright 2020-2023 Project Capsule Authors.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
)

func (in *CapsuleConfiguration) SetupWebhookWithManager(mgr ctrl.Manager) error {
	certData, _ := os.ReadFile("/tmp/k8s-webhook-server/serving-certs/tls.crt")
	if len(certData) == 0 {
		return nil
	}

	return ctrl.NewWebhookManagedBy(mgr).
		For(in).
		Complete()
}
