// No Sale event handler — compiled to a WASI command (GOOS=wasip1 GOARCH=wasm)
// and executed in-process by the till's wazero runtime. The till passes the
// event as JSON on stdin; stdout (JSON) is logged to the audit trail and
// stderr goes to the POS log.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type event struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Payload   struct {
		PluginID string `json:"plugin_id"`
		EntryKey string `json:"entry_key"`
		Label    string `json:"label"`
	} `json:"payload"`
}

func main() {
	var ev event
	if err := json.NewDecoder(os.Stdin).Decode(&ev); err != nil {
		fmt.Fprintf(os.Stderr, "nosale: bad event: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "nosale: drawer release requested via %s/%s at %s\n",
		ev.Payload.PluginID, ev.Payload.EntryKey, ev.Timestamp)

	// v1: record the request for the audit trail. Driving a physical drawer
	// needs a hardware/device plugin (runtime "go") wired to this event.
	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{
		"handled":      true,
		"action":       "drawer.open",
		"reason":       "no_sale",
		"requested_at": ev.Timestamp,
		"event_id":     ev.ID,
	})
}
