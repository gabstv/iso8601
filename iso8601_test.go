package iso8601

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"
)

func TestISO8601Time(t *testing.T) {
	now := New(time.Now().UTC())

	data, err := json.Marshal(now)
	if err != nil {
		t.Fatal(err)
	}

	_, err = time.Parse(`"`+Format+`"`, string(data))
	if err != nil {
		t.Fatal(err)
	}

	var now2 Time
	err = json.Unmarshal(data, &now2)
	if err != nil {
		t.Fatal(err)
	}

	if now != now2 {
		t.Fatalf("Time %s does not equal expected %s", now2, now)
	}

	if now.String() != now2.String() {
		t.Fatalf("String format for %s does not equal expected %s", now2, now)
	}
}

func TestISOString(t *testing.T) {
	tt := New(time.Date(1993, time.June, 23, 12, 15, 35, 0, time.UTC))
	ex := "1993-06-23T12:15:35"
	if tt.ISOString() != ex {
		t.Errorf("Expected ISOString() to return %s but got %s", ex, tt.ISOString())
	}
}

func TestURLEncode(t *testing.T) {
	isotime := New(time.Date(1993, time.June, 23, 12, 15, 35, 0, time.UTC))
	v := &url.Values{}

	isotime.EncodeValues("created", v)
	expected := "1993-06-23T12:15:35"
	if r := v.Get("created"); r != expected {
		t.Errorf("Error encoding URL value: expected %s but got %s\n", expected, r)
	}
}
