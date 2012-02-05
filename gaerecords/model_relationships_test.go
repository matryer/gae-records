package gaerecords

import (
	"testing"
)

func TestParentModel(t *testing.T) {
	
	Parent := NewModel("parent1")
	
	if Parent.ParentModel() != nil {
		t.Error("ParentModel should be nil")
	}
	assertEqual(t, false, Parent.HasParentModel())
	
	RootModel := NewModel("root")
	
	assertEqual(t, Parent, Parent.SetParentModel(RootModel))
	
	assertEqual(t, RootModel, Parent.ParentModel())
	assertEqual(t, RootModel, Parent.parentModel)
	assertEqual(t, true, Parent.HasParentModel())

}

func TestModelHasMany(t *testing.T) {
	
	Parent := NewModel("parent2")
	Child := Parent.HasMany("children")

	if assertNotNil(t, Child, "Child") {

		// ensure the child knows about the parent model
		assertEqual(t, Parent, Child.ParentModel())
	
	}
	
}
