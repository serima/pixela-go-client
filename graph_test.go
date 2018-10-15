package pixela

import (
	"reflect"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	c := NewClient()
	//Initialize test account
	c.DeleteUser("test-gainings", "testtest")
	c.DeleteUser("test-gainings", "testhogehoge")
	c.DeleteGraph("test-gainings", "testtest", "hoge1")
	c.DeleteGraph("test-gainings", "testtest", "hoge2")
	c.RegisterUser("test-gainings", "testtest", "yes", "yes")

	gi1 := GraphInfo{
		ID:   "hoge1",
		Name: "fuga1",
		Unit: "Kg",
		//Invalid unit type
		UnitType: "string",
		//invalid color type
		Color: "skyblue",
	}
	err := gi1.Validate()
	if err == nil {
		t.Fatalf("want invalid unit type error")
	}
	gi1.UnitType = "float"
	err = gi1.Validate()
	if err == nil {
		t.Fatalf("want invalid color error")
	}
	gi1.Color = "shibafu"
	err = gi1.Validate()
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}

	err = c.CreateGraph("test-gainings", "testtest", gi1)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}

	gi2 := GraphInfo{
		ID:       "hoge2",
		Name:     "fuga2",
		Unit:     "Kg",
		UnitType: "float",
		Color:    "shibafu",
	}
	err = c.CreateGraph("test-gainings", "testtest", gi2)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}

	gs, err := c.ListGraph("test-gainings", "testtest")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if len(gs) != 2 {
		t.Fatalf("want 2, but got %v", len(gs))
	}
	if !reflect.DeepEqual(gs[0], gi1) {
		t.Fatalf("want %#v, but got %#v", gs[0], gi1)
	}
	if !reflect.DeepEqual(gs[1], gi2) {
		t.Fatalf("want %#v, but got %#v", gs[1], gi2)
	}

	_, err = c.GetGraph("test-gainings", "testtest", "hoge3", "")
	if err == nil {
		t.Fatalf("want error, got nil")
	}

	svg, err := c.GetGraph("test-gainings", "testtest", "hoge2", "")
	if err != nil {
		t.Fatalf("want nil, got %v", err)
	}
	if svg == "" {
		t.Fatalf("want svg url")
	}

	newGi1 := GraphInfo{
		ID:       "hoge1",
		Name:     "newFuga1",
		Unit:     "Kg",
		UnitType: "float",
		Color:    "shibafu",
	}
	newGi2 := GraphInfo{
		ID:       "hoge2",
		Name:     "newFuga2",
		Unit:     "commit",
		UnitType: "int",
		Color:    "shibafu",
	}
	err = c.UpdateGraph("test-gainings", "testtest", newGi1)
	if err != nil {
		t.Fatalf("want nil, got %v", err)
	}
	err = c.UpdateGraph("test-gainings", "testtest", newGi2)
	if err != nil {
		t.Fatalf("want nil, got %v", err)
	}
	gs, err = c.ListGraph("test-gainings", "testtest")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if !reflect.DeepEqual(gs[0], newGi1) {
		t.Fatalf("want %#v, but got %#v", gs[0], newGi1)
	}
	c.DeleteUser("test-gainings", "testtest")
}