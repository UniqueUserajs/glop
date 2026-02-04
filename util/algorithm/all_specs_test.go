package algorithm_test

import (
  "github.com/orfjackal/gospec/src/gospec"
  "testing"
)

func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(DijkstraSpec)
  r.AddSpec(ReachableSpec)
  r.AddSpec(ReachableDestinationsSpec)
  r.AddSpec(ChooserSpec)
  // Choose2 and Map remove in 90245e6c9bbc54104534f9db4a857b9a64189932
  // r.AddSpec(Chooser2Spec)
  // r.AddSpec(MapperSpec)
  r.AddSpec(Mapper2Spec)
  r.AddSpec(TopoSpec)
  gospec.MainGoTest(r, t)
}
