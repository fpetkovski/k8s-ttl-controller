package ttl_controller

import (
	"bytes"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/util/jsonpath"
)

func GetTTLValue(object *unstructured.Unstructured, ttlValueField string) (time.Duration, error) {
	buf, err := getValue(object, ttlValueField)
	if err != nil {
		return 0, err
	}

	if duration, err := time.ParseDuration(buf.String()); err != nil {
		return 0, err
	} else {
		return duration, nil
	}
}

func GetExpirationValue(object *unstructured.Unstructured, expirationValueField string) (time.Time, error) {
	buf, err := getValue(object, expirationValueField)
	if err != nil {
		return time.Time{}, err
	}

	if t, err := time.Parse(time.RFC3339, buf.String()); err != nil {
		return time.Time{}, err
	} else {
		return t, nil
	}
}

func IsExpired(ttl time.Duration, createdAt time.Time) bool {
	return createdAt.Add(ttl).Before(time.Now())
}

func getValue(object *unstructured.Unstructured, field string) (*bytes.Buffer, error) {
	jp := jsonpath.New("ttl").AllowMissingKeys(true)
	if err := jp.Parse(`{` + field + `}`); err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jp.Execute(buf, object.UnstructuredContent()); err != nil {
		return nil, err
	}
	return buf, nil
}
