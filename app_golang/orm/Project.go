package orm

import (
	"gorm.io/gorm"
)

type Project struct {
  gorm.Model
  screen_IMG string
  project_name string
  project_Path string
  WorkspaceID uint
  Workspace   Workspace `gorm:"foreignKey:WorkspaceID"`

}
