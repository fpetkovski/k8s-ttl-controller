package ttl_controller

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
	"time"
)

func TestGetTTLValueFromMetadata(t *testing.T) {
	content := map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]string{
				"ttl": "30s",
			},
		},
	}
	object := unstructured.Unstructured{}
	object.SetUnstructuredContent(content)

	ttl, err := GetTTLValue(object, `.metadata.annotations.ttl`)
	assert.NoError(t, err)
	assert.Equal(t, 30*time.Second, ttl)
}

func TestGetExpirationValueFromCondition(t *testing.T) {
	lastUpdateTime, err := time.Parse(time.RFC3339, "2020-12-15T12:54:31Z")
	if err != nil {
		t.Fatal(err)
	}

	deployment := v1.Deployment{
		Status: v1.DeploymentStatus{
			Conditions: []v1.DeploymentCondition{
				{
					Type:   v1.DeploymentAvailable,
					Status: "True",
					LastUpdateTime: metav1.Time{
						lastUpdateTime,
					},
				},
			},
		},
	}
	jsonContent, _ := json.Marshal(&deployment)
	object := unstructured.Unstructured{}
	_ = object.UnmarshalJSON(jsonContent)

	expirationValue, err := GetExpirationValue(object, ".status.conditions[0].lastUpdateTime")
	assert.NoError(t, err)
	assert.True(t, lastUpdateTime.Equal(expirationValue))
}
