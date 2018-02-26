package permission

import (
	db "sujor.com/leo/sujor-api/database"
)

type Permission struct {
	Id 				int 	`json:"id" form:"id"`
	PermissionName 	string 	`json:"permission_name" form:"permission_name"`
	Description		string	`json:"description" form:"description"`
}

// GetPermissionsByName  获取权限列表
func (p Permission) GetPermissionsByName(username string) (permissions []Permission, err error) {
	permissions = make([]Permission, 0)
	sql := `SELECT p.id, p.permission_name, p.description
			FROM
			t_permission p,role_permission rp,t_role r
			WHERE
			r.id = rp.role_id AND rp.permission_id = p.id AND r.id
			IN
			(SELECT r.id
			FROM
   			t_user u,t_role r,user_role ur
			WHERE
  			u.username = ? AND u.id = ur.user_id AND ur.role_id = r.id)`
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()

	rows, err := stmt.Query(username)
	// 遍历rows
	for rows.Next() {
		var permission Permission
		rows.Scan(&permission.Id, &permission.PermissionName, &permission.Description)
		permissions = append(permissions, permission)
	}
	rows.Close()
	return
}