package article

type Article struct {
	Id		  	int		`json:"id" form:"id"`
	Title     	string 	`json:"title" form:"title"`
	AuthorName	string	`json:"author_name" form:"author_name"`
	Desc 	  	string 	`json:"desc" form:"desc"`
	CreatedAt 	string 	`json:"created_at" form:"created_at"`
	UpdatedAt 	string 	`json:"updated_at" form:"updated_at"`
}

