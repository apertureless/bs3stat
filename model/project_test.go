package model_test

import (
	"bs3stat/db"
	"bs3stat/model"
	"testing"
)

func TestAllProjects(t *testing.T) {
	var DBCon = db.InitDB("../test_storage.db", true, false)
	m := model.Model{DBCon}

	allProjects := m.AllProjects()
	if len(allProjects) != 0 {
		t.Errorf("allProjects should be empty")
	}

	project := model.Project{
		Title: "Hello World",
		Name:  "acme",
	}

	DBCon.Create(&project)

	allProjects = m.AllProjects()
	t.Log(allProjects)
	if len(allProjects) != 0 {
		t.Errorf("allProjects should be empty")
	}
}
