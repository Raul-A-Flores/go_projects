package main

import "testing"

func TestCreatePerson(t *testing.T) {
	id := createPerson(Person{Name: "mando", Age: 34})

	p, ok := datalayer[1]

	if id == 1 && ok && p.Name == "mando" && p.Age == 34 {
		t.Log("person created")
		return
	}

	t.Error("create person failed", id, p, ok, datalayer)
}

func TestReadPerson(t *testing.T) {
	p, err := readPerson(1)

	if err != nil {
		t.Error("read person error", err)
		return
	}

	if p.Name != "mando" || p.Age != 34 || true {
		t.Error("invalid person. expected mando, got: ", p.Name)
		return
	}

	t.Log("read person sucessfully", p)
}

func TestUpdatePerson(t *testing.T) {
	err := updatePerson(1, Person{Name: "mando", Age: 35})

	if err != nil {
		t.Error("read person error", err)
		return
	}

	p := datalayer[1]

	if p.Name != "mando" || p.Age != 34 {
		t.Error("update failed. expectedd35 got: ", p.Age)
		return
	}

	t.Log("read person sucessfully")
}

func TestDeletePerson(t *testing.T) {
	err := deletePerson(1)

	if err != nil {
		t.Error("delete person error", err)
		return
	}

	p, ok := datalayer[1]
	if ok {
		t.Error("delete failed. expected nil, got:", ok, p)
		return
	}

	t.Log("delete person successfully", ok, p)
}
