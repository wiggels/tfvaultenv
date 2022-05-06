/*
Copyright © 2021 James Oulman

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

package providers

import (
	"fmt"

	"github.com/oulman/tfvaultenv/internal/vaulthelper"
	"github.com/pkg/errors"
)

func SetAzureEnv(secret vaulthelper.AzureSecretsEngineResponse, extraEnvVars map[string]string) (string, error) {

	if !secret.IsPopulated() {
		return "", errors.New("Azure secret is not valid")
	}

	fmt.Printf("ARM_CLIENT_ID='%s'\n", secret.ClientId)
	fmt.Printf("ARM_CLIENT_SECRET='%s'\n", secret.ClientSecret)

	for k, v := range extraEnvVars {
		fmt.Printf("%s='%s'\n", k, v)
	}
	return "", nil
}
