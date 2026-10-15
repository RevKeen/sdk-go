//go:build sandbox

package revkeen

import (
	"context"
	"os"
	"strings"
	"testing"
)

func TestSandboxContract(t *testing.T) {
	key := os.Getenv("SDK_SANDBOX_API_KEY")
	if !strings.HasPrefix(key, "rk_sandbox_") || len(key) == len("rk_sandbox_") {
		t.Fatal("SDK_SANDBOX_API_KEY must be a complete rk_sandbox_ secret key")
	}

	client, err := NewSecretClient(key, EnvironmentSandbox)
	if err != nil {
		t.Fatalf("create sandbox client: %v", err)
	}

	result, response, err := client.ProductsAPI.ProductsList(context.Background()).Limit(1).Execute()
	if err != nil {
		t.Fatalf("list sandbox products: %v", err)
	}
	if response == nil || response.StatusCode < 200 || response.StatusCode >= 300 {
		t.Fatalf("unexpected sandbox response: %#v", response)
	}
	if result == nil {
		t.Fatal("sandbox products response is nil")
	}
}
