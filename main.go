package main

import (
	"fmt"
)

func main() {
	age := 16
	isAdult := age >= 18
	hasPermission := IsLoggedIn && hasRole
	canAccess := isAdmin || hasPermission
	isRestricted := !canAccess
}
