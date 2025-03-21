package db

import (
	"database/sql"
	"fmt"
	"gitmanager/i18n"
	"time"

	_ "modernc.org/sqlite"
)

type Project struct {
	ID              int64
	Name            string
	RepositoryURL   string
	LocalBranch     string
	RemoteBranch    string
	PullBranch      string
	LastPushTime    time.Time
	LastMachineID   string
	LastPushMessage string
	IsAppSelf       bool
}

func CreateDB(dbName string) (*sql.DB, error) {

	// 创建数据库
	database, err := sql.Open("sqlite", dbName)
	if err != nil {
		return nil, err
	}
	defer database.Close()

	// 创建项目表
	createTableSQL := `
CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL,
			repository_url TEXT NOT NULL,
			local_branch TEXT NOT NULL,
			remote_branch TEXT NOT NULL,
			pull_branch TEXT,
			last_push_time TIMESTAMP,
			last_machine_id TEXT,
			last_push_message TEXT,
			is_app_self BOOLEAN DEFAULT 0
);`

	_, err = database.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.create_table"), err)
	}

	return database, nil
}

func InitDB(dbName string) (*sql.DB, error) {

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.open_db"), err)
	}

	// 创建项目表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL,
			repository_url TEXT NOT NULL,
			local_branch TEXT NOT NULL,
			remote_branch TEXT NOT NULL,
			pull_branch TEXT,
			last_push_time TIMESTAMP,
			last_machine_id TEXT,
			last_push_message TEXT,
			is_app_self BOOLEAN DEFAULT 0
		)
	`)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.create_table"), err)
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {
	// 创建项目表
	createTableSQL := `
CREATE TABLE IF NOT EXISTS projects (
	name TEXT PRIMARY KEY,
	repository_url TEXT,
	local_branch TEXT,
	remote_branch TEXT,
	last_machine_id TEXT,
	last_push_time DATETIME,
	is_app_self BOOLEAN DEFAULT 0,
	last_push_message TEXT,
	pull_branch TEXT
);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf(i18n.T("error.create_table"), err)
	}
	return nil
}

func InsertProject(db *sql.DB, project *Project) error {
	stmt, err := db.Prepare(`
		REPLACE INTO projects (
			name, repository_url, local_branch, remote_branch, pull_branch, last_push_time, is_app_self
		) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, ?)
	`)
	if err != nil {
		return fmt.Errorf(i18n.T("error.prepare_insert"), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(project.Name, project.RepositoryURL, project.LocalBranch, project.RemoteBranch, project.PullBranch, project.IsAppSelf)
	if err != nil {
		return fmt.Errorf(i18n.T("error.insert_project"), err)
	}
	return nil
}

func AddProject(db *sql.DB, name, repoURL, localBranch, remoteBranch, pullBranch string, isAppSelf bool) error {
	query := `
		INSERT INTO projects (name, repository_url, local_branch, remote_branch, pull_branch, last_push_time, last_machine_id, last_push_message, is_app_self)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, name, repoURL, localBranch, remoteBranch, pullBranch, time.Time{}, "", "", isAppSelf)
	return err
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
			   last_machine_id, last_push_time, is_app_self, 
			   pull_branch,last_push_message
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
		var pullBranch sql.NullString
		var lastPushMessage sql.NullString
		err := rows.Scan(
			&p.Name, &p.RepositoryURL, &p.LocalBranch, &p.RemoteBranch,
			&lastMachineID, &lastPushTime, &p.IsAppSelf, &pullBranch, &lastPushMessage)
		if err != nil {
			return nil, fmt.Errorf(i18n.T("error.scan_project"), err)
		}

		if lastMachineID.Valid {
			p.LastMachineID = lastMachineID.String
		}

		if lastPushTime.Valid {
			p.LastPushTime = lastPushTime.Time
		}

		if pullBranch.Valid {
			p.PullBranch = pullBranch.String
		} else {
			p.PullBranch = p.RemoteBranch
		}

		if lastPushMessage.Valid {
			p.LastPushMessage = lastPushMessage.String
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
	var pullBranch sql.NullString

	err := db.QueryRow(`
		SELECT name, repository_url, local_branch, remote_branch, 
			   last_machine_id, last_push_time, is_app_self, 
			   pull_branch
		FROM projects 
		WHERE name = ?`, name).Scan(
		&p.Name, &p.RepositoryURL, &p.LocalBranch, &p.RemoteBranch,
		&lastMachineID, &lastPushTime, &p.IsAppSelf, &pullBranch)

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

	if pullBranch.Valid {
		p.PullBranch = pullBranch.String
	} else {
		p.PullBranch = p.RemoteBranch
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
func UpdateLastPush(db *sql.DB, name, machineID string) error {
	query := `
		UPDATE projects 
		SET last_push_time = CURRENT_TIMESTAMP,
			last_machine_id = ?
		WHERE name = ?
	`
	_, err := db.Exec(query, machineID, name)
	return err
}

func UpdateProjectPullBranch(database *sql.DB, projectName, branchName string) error {

	query := `UPDATE projects SET pull_branch = ? WHERE name = ?`

	_, err := database.Exec(query, branchName, projectName)
	return err
}

func CheckAndUpdatePullBranch(database *sql.DB) error {
	// 检查 pull_branch 字段是否存在
	query := `PRAGMA table_info(projects)`
	rows, err := database.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	columnExists := false
	for rows.Next() {
		var cid, name, type_ string
		var notnull, dflt_value interface{}
		var pk int
		if err := rows.Scan(&cid, &name, &type_, &notnull, &dflt_value, &pk); err != nil {
			return err
		}
		if name == "pull_branch" {
			columnExists = true
			break
		}
	}

	// 如果字段不存在，添加字段并设置默认值为主分支
	if !columnExists {
		query := `ALTER TABLE projects ADD COLUMN pull_branch TEXT`
		_, err := database.Exec(query)
		if err != nil {
			return err
		}
	}

	// 更新 pull_branch 为主分支
	query = `UPDATE projects SET pull_branch = remote_branch`
	_, err = database.Exec(query)
	return err
}

// 添加新的 UpdateLastPushMessage 函数
func UpdateLastPushMessage(db *sql.DB, name, message string) error {
	query := `
		UPDATE projects 
		SET last_push_message = ?
		WHERE name = ?
	`
	_, err := db.Exec(query, message, name)
	return err
}

// 修改 UpdateLastPushWithMessage 函数
func UpdateLastPushWithMessage(db *sql.DB, name, machineID, message string) error {
	query := `
		UPDATE projects 
		SET last_push_time = CURRENT_TIMESTAMP,
			last_machine_id = ?,
			last_push_message = ?
		WHERE name = ?
	`
	_, err := db.Exec(query, machineID, message, name)
	return err
}
