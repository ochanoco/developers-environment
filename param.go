package main

var DB_TYPE = "sqlite3"
var DB_CONFIG = "file:./db.sqlite3?_fk=1"
var WHITELIST_FILE = "./whitelist.json"

var LOGIN_REDIRECT_PAGE_URL = "http://localhost:3000/redirect"
var ERROR_PAGE_URL = "http://localhost:3000/error"

var AUTH_PAGE_DOMAIN = "localhost:9000"
var AUTH_PAGE_DESTINATION = "localhost:3000"
