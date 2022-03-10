package simpleforce

import (
	"context"
	"testing"
)

func TestClient_Tooling_Query(t *testing.T) {
	client := requireClient(t, true)

	q := "SELECT Id, Name FROM Layout WHERE Name = 'Account Layout'"
	result, err := client.Tooling().Query(context.Background(), q)
	if err != nil {
		t.FailNow()
	}
	if len(result.Records) > 0 {
		case0 := &result.Records[0]
		if case0.StringField("Name") != "Account Layout" {
			t.FailNow()
		}
	}
}

func TestClient_ExecuteAnonymous(t *testing.T) {
	client := requireClient(t, true)

	apexBody := "System.debug('test');"
	result, err := client.ExecuteAnonymous(context.Background(), apexBody)
	if err != nil {
		t.FailNow()
	}
	if !result.Success {
		t.FailNow()
	}
}
