package main

import (
	"io"
	"time"

	"github.com/aquasecurity/table"

	"github.com/crowdsecurity/crowdsec/pkg/database/ent"
	"github.com/crowdsecurity/crowdsec/pkg/emoji"
)

func getBouncersTable(out io.Writer, bouncers []*ent.Bouncer) {
	t := newLightTable(out)
	t.SetHeaders("Name", "IP Address", "Valid", "Last API pull", "Type", "Version", "Auth Type")
	t.SetHeaderAlignment(table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft)
	t.SetAlignment(table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft, table.AlignLeft)

	for _, b := range bouncers {
		revoked := emoji.CheckMark
		if b.Revoked {
			revoked = emoji.Prohibited
		}

		t.AddRow(b.Name, b.IPAddress, revoked, b.LastPull.Format(time.RFC3339), b.Type, b.Version, b.AuthType)
	}

	t.Render()
}
