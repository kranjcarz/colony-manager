package controller

import (
	// "archive/zip"
	// "encoding/json"
	// "fmt"
	// "github.com/eaciit/colony-core/v0"
	// "github.com/eaciit/colony-manager/helper"
	// "github.com/eaciit/dbox"
	// _ "github.com/eaciit/dbox/dbc/jsons"
	"github.com/eaciit/knot/knot.v1"
	// "github.com/eaciit/toolkit"
	// "io"
	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "strings"
	// "time"
)

type GroupController struct {
	App
}

func CreateGroupController(s *knot.Server) *GroupController {
	var controller = new(GroupController)
	controller.Server = s
	return controller
}
