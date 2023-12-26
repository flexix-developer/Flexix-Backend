// package orm

// import (
// 	"gorm.io/gorm"
// )

// type Project struct {
//   gorm.Model
//   screen_IMG string
//   project_name string
//   project_Path string
//   WorkspaceID uint
//   Workspace   Workspace `gorm:"foreignKey:WorkspaceID"`

// }

package orm

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	WorkspaceID  uint
	Workspace    Workspace `gorm:"foreignKey:WorkspaceID"`
  ScreenIMG    string    `gorm:"column:screen_img"`
	ProjectName  string    `gorm:"column:project_name"`
	ProjectPath  string    `gorm:"column:project_path"`
}
