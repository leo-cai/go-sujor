package project

import (
	db "sujor.com/leo/sujor-api/database"
	"log"
)

type Project struct {
	Id		  int		`json:"id" form:"id"`
	Title     string 	`json:"title" form:"title"`
	Author	  string	`json:"author" form:"author"`
	Content   string 	`json:"content" form:"content"`
	CreatedAt string 	`json:"created_at" form:"created_at"`
	UpdatedAt string 	`json:"updated_at" form:"updated_at"`
}

func (p *Project) GetProjects(limit int, page int) (projects []Project, err error) {
	projects = make([]Project, 0)
	sql := "SELECT id, title, author, content, created_at, updated_at FROM project WHERE id >= (SELECT id FROM project ORDER BY id LIMIT ?, 1) LIMIT ?"
	// stmt及错误处理
	stmt, err := db.SqlDB.Prepare(sql)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
	// rows及错误处理
	rows, err := stmt.Query((page-1)*limit, limit)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	// 遍历
	for rows.Next() {
		var project Project
		rows.Scan(&project.Id, &project.Title, &project.Author, &project.Content, &project.CreatedAt, &project.UpdatedAt)
		projects = append(projects, project)
		//log.Println(projects)
	}

	if err = rows.Err(); err != nil {
		rows.Close()
		return
	}
	return
}
