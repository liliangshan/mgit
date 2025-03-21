package db

import (
	"database/sql"
	"fmt"
	"gitmanager/i18n"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Project struct {
	Name          string
	RepositoryURL string
	LocalBranch   string
	RemoteBranch  string
	LastMachineID string
	LastPushTime  time.Time
	IsAppSelf     bool // 新增字段：标记是否为应用自身
	PullBranch    string
}

func CreateDB(dbName string) error {
	// 获取当前可执行文件路径
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	currentDir := filepath.Dir(ex)

	// 使用当前目录作为数据库路径
	dbPath := filepath.Join(currentDir, dbName)

	// 创建数据库
	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer database.Close()

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
	pull_branch TEXT
);`

	_, err = database.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf(i18n.T("error.create_table"), err)
	}

	return nil
}

func InitDB(dbName string) (*sql.DB, error) {
	// 获取当前可执行文件路径
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	currentDir := filepath.Dir(ex)

	// 使用当前目录作为数据库路径
	dbPath := filepath.Join(currentDir, dbName)

	// 打开数据库
	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
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
	is_app_self BOOLEAN DEFAULT 0,
	pull_branch TEXT
);`

	_, err = database.Exec(createTableSQL)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("error.create_table"), err)
	}

	return database, nil
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
	stmt, err := db.Prepare(`
		REPLACE INTO projects (
			name, repository_url, local_branch, remote_branch, pull_branch, last_push_time, is_app_self
		) VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, ?)
	`)
	if err != nil {
		return fmt.Errorf(i18n.T("error.prepare_insert"), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, repoURL, localBranch, remoteBranch, pullBranch, isAppSelf)
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
			   last_machine_id, last_push_time, is_app_self, 
			   pull_branch
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

		err := rows.Scan(
			&p.Name, &p.RepositoryURL, &p.LocalBranch, &p.RemoteBranch,
			&lastMachineID, &lastPushTime, &p.IsAppSelf, &pullBranch)
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
