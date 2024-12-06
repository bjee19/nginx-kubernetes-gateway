package upstreamsettings

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ngfAPI "github.com/nginxinc/nginx-gateway-fabric/apis/v1alpha1"
	"github.com/nginxinc/nginx-gateway-fabric/internal/framework/helpers"
	"github.com/nginxinc/nginx-gateway-fabric/internal/mode/static/nginx/config/policies"
)

func TestProcess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		expUpstreamSettings policies.UpstreamSettings
		policies            []policies.Policy
	}{
		{
			name: "all fields populated",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						ZoneSize: helpers.GetPointer[ngfAPI.Size]("2m"),
						KeepAlive: helpers.GetPointer(ngfAPI.UpstreamKeepAlive{
							Connections: helpers.GetPointer(int32(1)),
							Requests:    helpers.GetPointer(int32(1)),
							Time:        helpers.GetPointer[ngfAPI.Duration]("5s"),
							Timeout:     helpers.GetPointer[ngfAPI.Duration]("10s"),
						}),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				ZoneSize:             "2m",
				KeepAliveConnections: 1,
				KeepAliveRequests:    1,
				KeepAliveTime:        "5s",
				KeepAliveTimeout:     "10s",
			},
		},
		{
			name: "zoneSize set",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						ZoneSize: helpers.GetPointer[ngfAPI.Size]("2m"),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				ZoneSize: "2m",
			},
		},
		{
			name: "keepAlive Connections set",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						KeepAlive: helpers.GetPointer(ngfAPI.UpstreamKeepAlive{
							Connections: helpers.GetPointer(int32(1)),
						}),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				KeepAliveConnections: 1,
			},
		},
		{
			name: "keepAlive Requests set",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						KeepAlive: helpers.GetPointer(ngfAPI.UpstreamKeepAlive{
							Requests: helpers.GetPointer(int32(1)),
						}),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				KeepAliveRequests: 1,
			},
		},
		{
			name: "keepAlive Time set",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						KeepAlive: helpers.GetPointer(ngfAPI.UpstreamKeepAlive{
							Time: helpers.GetPointer[ngfAPI.Duration]("5s"),
						}),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				KeepAliveTime: "5s",
			},
		},
		{
			name: "keepAlive Timeout set",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{
						KeepAlive: helpers.GetPointer(ngfAPI.UpstreamKeepAlive{
							Timeout: helpers.GetPointer[ngfAPI.Duration]("10s"),
						}),
					},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{
				KeepAliveTimeout: "10s",
			},
		},
		{
			name: "no fields populated",
			policies: []policies.Policy{
				&ngfAPI.UpstreamSettingsPolicy{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "usp",
						Namespace: "test",
					},
					Spec: ngfAPI.UpstreamSettingsPolicySpec{},
				},
			},
			expUpstreamSettings: policies.UpstreamSettings{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			processor := NewProcessor()

			g.Expect(processor.Process(test.policies)).To(Equal(test.expUpstreamSettings))
		})
	}
}