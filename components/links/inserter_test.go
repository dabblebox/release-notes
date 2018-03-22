package links

import (
	"fmt"
	"testing"
)

func TestEnsureLinksAreInserted(t *testing.T) {
	const token = "{TICKET_NUMBER}"

	url := URL{
		ReplaceRegEx:  `[A-Z]{7}-\d*`,
		ReplaceToken:  token,
		TokenizedLink: fmt.Sprintf("<http://tickets.turner.com/browse/%s|%s>", token, token),
	}

	comment := "RGTMGMT-422: Feature or bug comments can be added here."

	expected := "<http://tickets.turner.com/browse/RGTMGMT-422|RGTMGMT-422>: Feature or bug comments can be added here."
	actual := Insert(url, comment)

	if actual != expected {
		t.Error("failed to insert link")
	}
}
