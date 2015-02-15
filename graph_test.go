package transportdata

import "testing"

func TestGraphStopNode(t *testing.T) {
	hopEdges := make([]HopEdge, 0)
	stop := StopNode{"Central", hopEdges}
	if stop.Name != "Central" {
		t.Errorf("Expected Central but got: %#v", stop.Name)
	}
}