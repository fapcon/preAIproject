package models

type App struct {
	ID          int64  `json:"id" db:"id" db_type:"BIGINT primary key" db_default:"not null" db_index:"index,unique"`
	Name        string `json:"name" db:"name" db_type:"varchar(89)" db_default:"not null" db_index:"index,unique"`
	RedirectURL string `json:"redirect_url" db:"redirect_url" db_type:"varchar" db_default:"not null" db_index:"index,unique"`
}

func (a *App) TableName() string {
	return "apps"
}

func (a *App) OnCreate() []string {
	return []string{}
}
