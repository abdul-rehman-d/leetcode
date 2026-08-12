package time_based_key_value_store

import (
	"testing"
)

func assertEqual(t *testing.T, operation string, expected string, actual string) {
	t.Helper()
	if actual != expected {
		t.Fatalf("\n%s failed:\nExpected: %+v\nActual: %+v", operation, expected, actual)
	}
}

func TestTimeMap(t *testing.T) {
	// ["TimeMap", "set", "get", "get", "set", "get", "get"]
	// [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
	// [null, null, "bar", "bar", null, "bar2", "bar2"]
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	assertEqual(t, "Get(foo, 1)", "bar", obj.Get("foo", 1))
	assertEqual(t, "Get(foo, 3)", "bar", obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	assertEqual(t, "Get(foo, 4)", "bar2", obj.Get("foo", 4))
	assertEqual(t, "Get(foo, 5)", "bar2", obj.Get("foo", 5))
}
