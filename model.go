package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ochanoco/ochano.co-auth/proxy/ent"
)

type Database struct {
	ctx    context.Context
	client *ent.Client
}

var db *Database

func initDB() (*Database, error) {
	err := errors.New("error")
	db := new(Database)

	client, err := ent.Open(DB_TYPE, DB_CONFIG)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	db = new(Database)
	db.ctx = ctx
	db.client = client

	return db, err
}

func migrateWhiteList() error {
	var urls []string

	b, _ := os.ReadFile(WHITELIST_FILE)
	err := json.Unmarshal(b, &urls)

	if err != nil {
		log.Fatalf("failed to load migrate.json: %v", err)
		return err
	}

	projc := createProject(db, AUTH_PAGE_DOMAIN, AUTH_PAGE_DESTINATION, "root", "root")
	proj, nil := projc.Save(db.ctx)

	if err != nil {
		fmt.Errorf("failed creating project: %v", err)
		return err
	}

	for _, url := range urls {
		wl := createWhiteList(db, url)
		proj, err = saveWhiteListOnProj(db, proj, wl)

		if err != nil {
			fmt.Errorf("failed add white list to project: %v", err)
			return err
		}
	}

	return nil
}

func createWhiteList(db *Database, url string) *ent.WhiteListCreate {
	wl := db.client.WhiteList.
		Create().
		SetURL(url)

	return wl
}

func createProject(db *Database, domain string, destination string, lineId string, name string) *ent.ProjectCreate {
	proj := db.client.Project.
		Create().
		SetDomain(domain).
		SetDestination(destination).
		SetLineID(lineId).
		SetName(name)

	return proj
}

func saveWhiteListOnProj(db *Database, projc *ent.Project, wlc *ent.WhiteListCreate) (*ent.Project, error) {
	wl, err := wlc.Save(db.ctx)

	if err != nil {
		return projc, err
	}

	proj, err := projc.
		Update().
		AddWhitelists(wl).
		Save(db.ctx)

	return proj, err
}
