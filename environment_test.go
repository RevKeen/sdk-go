package revkeen

import "testing"

func TestNamedEnvironmentsResolveExplicitly(t *testing.T) {
	for _, test := range []struct {
		environment Environment
		want        string
	}{
		{environment: EnvironmentSandbox, want: "https://staging-api.revkeen.com/v2"},
		{environment: EnvironmentProduction, want: "https://api.revkeen.com/v2"},
	} {
		got, err := test.environment.baseURL()
		if err != nil {
			t.Fatal(err)
		}
		if got != test.want {
			t.Fatalf("environment %q URL = %q; want %q", test.environment, got, test.want)
		}
	}
}

func TestStandardConstructorsRejectUnknownEnvironmentWithoutProductionFallback(t *testing.T) {
	unknown := Environment("prodution")
	if _, err := NewClient("rk_live_test", unknown); err == nil {
		t.Fatal("NewClient accepted an unknown environment")
	}
	if _, err := NewSecretClient("rk_live_test", unknown); err == nil {
		t.Fatal("NewSecretClient accepted an unknown environment")
	}
	if _, err := NewPublishableClient("rk_pk_live_test", unknown); err == nil {
		t.Fatal("NewPublishableClient accepted an unknown environment")
	}
}

func TestStandardClientUsesNamedEnvironmentURL(t *testing.T) {
	client, err := NewClient("rk_sandbox_test", EnvironmentSandbox)
	if err != nil {
		t.Fatal(err)
	}
	if got := client.cfg.Servers[0].URL; got != sandboxBaseURL {
		t.Fatalf("sandbox client URL = %q; want %q", got, sandboxBaseURL)
	}

	client, err = NewSecretClient("rk_live_test", EnvironmentProduction)
	if err != nil {
		t.Fatal(err)
	}
	if got := client.cfg.Servers[0].URL; got != productionBaseURL {
		t.Fatalf("production client URL = %q; want %q", got, productionBaseURL)
	}
}

func TestCustomBaseURLRequiresExplicitConstructor(t *testing.T) {
	client, err := NewClientWithCustomBaseURL("rk_sandbox_test", "https://localhost:8443/v2/")
	if err != nil {
		t.Fatal(err)
	}
	if got := client.cfg.Servers[0].URL; got != "https://localhost:8443/v2" {
		t.Fatalf("custom client URL = %q", got)
	}

	for _, rawURL := range []string{
		"api.revkeen.test/v2",
		"ftp://api.revkeen.test/v2",
		"https://user:secret@api.revkeen.test/v2",
		"https://api.revkeen.test/v2?tenant=unsafe",
		"https://api.revkeen.test/v2#fragment",
	} {
		if _, err := NewSecretClientWithCustomBaseURL("rk_sandbox_test", rawURL); err == nil {
			t.Fatalf("unsafe custom base URL %q was accepted", rawURL)
		}
	}
}
