package db

import (
	"database/sql"
	"fmt"
	"time"
	_ "modernc.org/sqlite"
	"gitmanager/i18n"
)

type Project struct {
	Name          string
	RepositoryURL string
	LocalBranch   string
	RemoteBranch  string
	LastMachineID string
	LastPushTime  time.Time
	IsAppSelf     bool    // 新增字段：标记是否为应用自身
}

func InitDB() (*sql.DB, error) {
	// 使用 modernc.org/sqlite 驱动
	db, err := sql.Open("sqlite", "projects.db")
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.db_open"), err)
	}

	// 创建项目表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS projects (
		name TEXT PRIMARY KEY,
		repository_url TEXT,
		local_branch TEXT,
		remote_branch TEXT,
		last_machine_id TEXT,
		last_push_time DATETIME,
		is_app_self BOOLEAN DEFAULT 0
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.create_table"), err)
	}

	return db, nil
}

func AddProject(db *sql.DB, name, repoURL, localBranch, remoteBranch string, isAppSelf bool) error {
	stmt, err := db.Prepare(`
		INSERT INTO projects (
			name, repository_url, local_branch, remote_branch, last_push_time, is_app_self
		) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, ?)
	`)
	if err != nil {
		return fmt.Errorf(i18n.T("error.prepare_insert"), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, repoURL, localBranch, remoteBranch, isAppSelf)
	if err != nil {
		return fmt.Errorf(i18n.T("error.insert_project"), err)
	}
	return nil
}

func UpdateProjectPush(db *sql.DB, name, machineID string) error {
	_, err := db.Exec(`
		UPDATE projects 
		SET last_machine_id = ?, last_push_time = ?
		WHERE name = ?`,
		machineID, time.Now(), name)
	if err != nil {
		return fmt.Errorf(i18n.T("error.update_project_push"), err)
	}
	return nil
}

func ListProjects(db *sql.DB) ([]Project, error) {
	rows, err := db.Query(`
		SELECT name, repository_url, local_branch, remote_branch, 
			   last_machine_id, last_push_time, is_app_self 
		FROM projects`)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.query_projects"), err)
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var p Project
		var lastMachineID sql.NullString
		var lastPushTime sql.NullTime
		
		err := rows.Scan(
			&p.Name, &p.RepositoryURL, &p.LocalBranch, &p.RemoteBranch,
			&lastMachineID, &lastPushTime, &p.IsAppSelf)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("error.scan_project"), err)
		}
		
		if lastMachineID.Valid {
			p.LastMachineID = lastMachineID.String
		}
		
		if lastPushTime.Valid {
			p.LastPushTime = lastPushTime.Time
		}
		
		projects = append(projects, p)
	}
	return projects, nil
}

// GetProject 获取指定项目的信息
func GetProject(db *sql.DB, name string) (*Project, error) {
	var p Project
	var lastMachineID sql.NullString
	var lastPushTime sql.NullTime
	
	err := db.QueryRow(`
		SELECT name, repository_url, local_branch, remote_branch, 
			   last_machine_id, last_push_time, is_app_self 
		FROM projects 
		WHERE name = ?`, name).Scan(
		&p.Name, &p.RepositoryURL, &p.LocalBranch, &p.RemoteBranch,
		&lastMachineID, &lastPushTime, &p.IsAppSelf)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(i18n.T("error.project_not_found"), name)
		}
		return nil, fmt.Errorf(i18n.T("error.query_project"), err)
	}
	
	if lastMachineID.Valid {
		p.LastMachineID = lastMachineID.String
	}
	
	if lastPushTime.Valid {
		p.LastPushTime = lastPushTime.Time
	}
	
	return &p, nil
}

// DeleteProject 从数据库中删除项目
func DeleteProject(db *sql.DB, name string) error {
	result, err := db.Exec("DELETE FROM projects WHERE name = ?", name)
	if err != nil {
		return fmt.Errorf(i18n.T("error.delete_project"), err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(i18n.T("error.check_deleted"), err)
	}

	if rows == 0 {
		return fmt.Errorf(i18n.T("error.project_not_found"), name)
	}

	return nil
}

// UpdateProject 更新项目信息
func UpdateProject(db *sql.DB, name, repositoryURL, localBranch, remoteBranch string) error {
	result, err := db.Exec(`
		UPDATE projects 
		SET repository_url = ?, local_branch = ?, remote_branch = ?
		WHERE name = ?`,
		repositoryURL, localBranch, remoteBranch, name)
	if err != nil {
		return fmt.Errorf(i18n.T("error.update_project"), err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(i18n.T("error.check_updated"), err)
	}

	if rows == 0 {
		return fmt.Errorf(i18n.T("error.project_not_found"), name)
	}

	return nil
}

// UpdateLastPush 更新项目的最后提交信息
func UpdateLastPush(db *sql.DB, projectName, machineID string) error {
	result, err := db.Exec(`
		UPDATE projects 
		SET last_machine_id = ?, last_push_time = CURRENT_TIMESTAMP 
		WHERE name = ?`,
		machineID, projectName)
	if err != nil {
		return fmt.Errorf(i18n.T("error.update_last_push"), err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(i18n.T("error.check_updated"), err)
	}

	if rows == 0 {
		return fmt.Errorf(i18n.T("error.project_not_found"), projectName)
	}

	return nil
} 