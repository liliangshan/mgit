package main

import (
	"bufio"
	"crypto/md5"
	"database/sql"
	"fmt"
	"gitmanager/db"
	"gitmanager/git"
	"gitmanager/i18n"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"

	// 修改为 sqlite
	_ "modernc.org/sqlite"
)

// 添加全局变量
var appExePath string
var dbFilePath string
var databaseData *sql.DB

// 添加新的辅助函数来获取环境变量文件名
func getEnvFileName() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	baseName := filepath.Base(ex)
	// 移除扩展名
	baseName = strings.TrimSuffix(baseName, filepath.Ext(baseName))

	// 如果是 mgit，使用 .env
	if strings.ToLower(baseName) == "mgit" {
		return ".env"
	}

	// 其他情况使用 .应用名.env
	return fmt.Sprintf(".%s.env", strings.ToLower(baseName))
}

// 添加一个新的函数来获取数据库文件名
func getDbFileName() string {
	ex, err := os.Executable()
	if err != nil {
		return "projects.db" // 默认名称
	}

	exeName := filepath.Base(ex)
	// 移除 .exe 后缀 (Windows系统)
	exeName = strings.TrimSuffix(exeName, ".exe")

	// 如果是默认应用名 mgit，使用 projects.db
	if exeName == "mgit" {
		return "projects.db"
	}

	// 其他应用使用 应用名.db
	return exeName + ".db"
}

// 修改远程数据库目录名获取函数
func getRemoteDbDirName() string {

	//如果是远程数据库
	if os.Getenv("DB_ENABLED") == "true" {
		ex, err := os.Executable()
		if err != nil {
			return ".mgit_db" // 默认名称
		}

		exeName := filepath.Base(ex)
		exeName = strings.TrimSuffix(exeName, ".exe")

		if exeName == "mgit" {
			return filepath.Join(os.Getenv("APP_PATH"), ".mgit_db")
		}
		return filepath.Join(os.Getenv("APP_PATH"), "."+exeName+"_db")
	}
	return appExePath
}

func main() {
	var database *sql.DB
	var err error
	var lang string
	// 初始化默认语言
	i18n.SetLanguage("zh-CN")

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// 设置全局变量
	appExePath = filepath.Dir(ex)

	// 尝试加载 .env 文件并设置语言
	envFile := getEnvFileName()
	if err := godotenv.Load(filepath.Join(appExePath, envFile)); err == nil {
		appLang := os.Getenv("MGIT_LANG")
		if appLang != "" {
			i18n.SetLanguage(lang)
		}
		lang = appLang
	}

	// 先检查是否是帮助命令
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "set":
			// 语言选择和初始设置
			if err := handleSettings(); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		case "h", "-h", "-help", "help":
			isTemp := false
			// 检查 .env 文件是否存在，如果不存在则让用户选择语言
			if _, err := os.Stat(filepath.Join(appExePath, envFile)); os.IsNotExist(err) {
				isTemp = true

				// 语言选择
				type Language struct {
					Name string
					Code string
				}
				languages := []Language{
					{Name: "简体中文", Code: "zh-CN"},
					{Name: "繁體中文", Code: "zh-TW"},
					{Name: "English", Code: "en-US"},
					{Name: "日本語", Code: "ja-JP"},
					{Name: "한국어", Code: "ko-KR"},
					{Name: "Français", Code: "fr-FR"},
				}

				prompt := promptui.Select{
					Label: "请选择语言 (Please select language)",
					Items: languages,
					Templates: &promptui.SelectTemplates{
						Label:    "{{ . }}",
						Active:   "> {{ .Name | cyan }}",
						Inactive: "  {{ .Name }}",
						Selected: "✓ {{ .Name | green }}",
					},
				}

				index, _, err := prompt.Run()
				if err != nil {
					log.Fatal("语言选择失败:", err)
				}

				// 1. 创建临时的 .env 文件，包含 MGIT_HOME
				envContent := fmt.Sprintf(`MGIT_LANG=%s
MGIT_HOME=%s`, languages[index].Code, appExePath)
				if err := os.WriteFile(filepath.Join(appExePath, envFile), []byte(envContent), 0644); err != nil {
					log.Fatal("创建临时语言文件失败:", err)
				}

				// 2. 加载 .env 文件
				if err := godotenv.Load(filepath.Join(appExePath, envFile)); err != nil {
					log.Fatal("加载语言设置失败:", err)
				}

				// 3. 设置语言
				i18n.SetLanguage(languages[index].Code)
			} else {
				// 如果 .env 文件存在，加载它
				if err := godotenv.Load(filepath.Join(appExePath, envFile)); err != nil {
					log.Fatal(i18n.T("msg.env_load_failed"))
				}
				// 重新设置语言
				lang := os.Getenv("MGIT_LANG")
				log.Println("lang is: ", os.Getenv("MGIT_HOME"))
				i18n.SetLanguage(lang)
			}

			showHelp(appExePath)
			// 如果是临时文件，显示帮助后删除
			if isTemp {
				time.Sleep(100)
				os.Remove(filepath.Join(appExePath, envFile))
			}
			return
		}
	}

	// 检查 .env 文件是否存在，如果不存在则让用户选择语言并执行 set 命令
	if _, err := os.Stat(filepath.Join(appExePath, envFile)); os.IsNotExist(err) {

		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		currentDir := filepath.Dir(ex)
		// 创建初始的.env文件,移除 MGIT_HOME
		envContent := fmt.Sprintf("# "+i18n.T("env.machine_id")+"\nMACHINE_ID=%s\n\n# "+i18n.T("env.app_path")+"\nAPP_PATH=%s\n\n# "+i18n.T("env.mgit_lang")+"\nMGIT_LANG=%s\n\n# "+i18n.T("prompt.db_enabled")+"\nDB_ENABLED=false\nDB_REPO=\nDB_BRANCH=",
			"machine-01",
			filepath.Dir(currentDir),
			lang)

		if err := os.WriteFile(filepath.Join(currentDir, envFile), []byte(envContent), 0644); err != nil {
			log.Fatal(i18n.T("msg.env_create_failed"), err)
		}

		//加载环境变量
		if err := godotenv.Load(filepath.Join(currentDir, envFile)); err != nil {
			log.Fatal(i18n.T("msg.env_load_failed"), err)
		}
		time.Sleep(100000)

		if err := updateEnvFile("MGIT_HOME", currentDir); err != nil {
			log.Fatal(i18n.T("msg.lang_update_failed"), err)
		}

		// 语言选择和初始设置
		if err := handleSettings(); err != nil {
			log.Fatal(err)
		}
		// 如果是 set 命令且没有额外参数，执行完就退出
		if len(os.Args) > 1 && strings.ToLower(os.Args[1]) == "set" && len(os.Args) == 2 {
			return
		}
	} else {
		// 加载 .env 文件
		if err := godotenv.Load(filepath.Join(appExePath, envFile)); err != nil {
			log.Fatal(i18n.T("msg.env_load_failed"))
		}
	}
	log.Println("appExePath is: ", filepath.Join(appExePath, envFile))
	//如果是远程仓库，获取远程数据库目录
	dbDir := getRemoteDbDirName()
	dbFilePath = filepath.Join(dbDir, getDbFileName())
	if databaseData == nil {
		//如果数据库不存在，则创建数据库
		if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {

			databaseData, err = db.CreateDB(dbFilePath)
			if err != nil {
				log.Fatal(i18n.T("msg.db_create_failed"), err)
			}
		}
	}
	// 初始化数据库
	database, err = db.InitDB(dbFilePath)

	if err != nil {
		log.Fatal(i18n.T("cmd.set"), err)
	}

	defer database.Close()

	lang = os.Getenv("MGIT_LANG")
	if lang == "" {
		lang = "zh-CN"
	}
	i18n.SetLanguage(lang)

	// 获取机器标识
	machineID := os.Getenv("MACHINE_ID")
	if machineID == "" {
		log.Fatal(i18n.T("msg.machine_id_not_set"))
	}

	// 根据命令行参数执行不同操作
	if len(os.Args) < 2 {
		showHelp(appExePath)
		return
	}

	switch os.Args[1] {
	case "h", "-h", "-help", "help":
		showHelp(appExePath)
		return
	case "init":
		var projectName string
		var repoURL string
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(i18n.T("msg.dir_get_failed"), err)
		}

		// 获取项目名称
		if len(os.Args) > 2 {
			projectName = os.Args[2]
		} else {
			fmt.Print(i18n.T("prompt.project_name"))
			projectName = readLine()
		}

		// 清理项目名称中的非法字符
		projectName = cleanProjectName(projectName)

		// 如果名称为空，返回错误
		if projectName == "" {
			log.Fatal(i18n.T("error.empty_project_name"))
		}

		// 检查项目目录是否已存在
		projectDir := filepath.Join(os.Getenv("APP_PATH"), projectName)
		if _, err := os.Stat(projectDir); err == nil {
			// 目录已存在，检查是否是 Git 仓库
			if err := os.Chdir(projectDir); err != nil {
				log.Fatal(i18n.T("msg.dir_change_failed"), err)
			}

			// 检查是否是 Git 仓库
			if err := git.GitOperation(true, "git", "rev-parse", "--git-dir"); err == nil {
				// 获取仓库地址
				cmd := exec.Command("git", "config", "--get", "remote.origin.url")
				output, err := cmd.Output()
				if err != nil {
					log.Fatal(i18n.T("msg.repo_url_get_failed"), err)
				}
				repoURL = strings.TrimSpace(string(output))

				// 获取当前分支
				cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
				output, err = cmd.Output()
				if err != nil {
					log.Fatal(i18n.T("msg.branch_get_failed"), err)
				}
				currentBranch := strings.TrimSpace(string(output))

				// 保存项目信息到数据库
				if err := db.AddProject(database, projectName, repoURL, currentBranch, currentBranch, currentBranch, false); err != nil {
					log.Fatal(i18n.T("msg.project_save_failed"), err)
				}

				// 切回原始目录
				if err := os.Chdir(currentDir); err != nil {
					log.Fatal(i18n.T("msg.dir_back_failed"), err)
				}

				fmt.Printf(i18n.T("msg.project_init_done"), projectName)
				fmt.Printf(i18n.T("msg.project_path"), projectDir)
				return
			}
			log.Fatalf(i18n.T("msg.dir_exists_not_git"), projectDir)
		}

		// 获取仓库地址
		if len(os.Args) > 3 {
			repoURL = os.Args[3]
		} else {
			fmt.Print(i18n.T("prompt.repo_url"))
			repoURL = readLine()
		}

		// 创建项目目录
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			log.Fatal(i18n.T("msg.dir_create_failed"), err)
		}

		// 切换到项目目录
		if err := os.Chdir(projectDir); err != nil {
			log.Fatal(i18n.T("msg.dir_change_failed"), err)
		}

		// 初始化 Git 仓库并拉取代码
		pullBranch, remoteBranch, err := git.GitInit(repoURL, "")
		if err != nil {
			log.Fatal(i18n.T("msg.git_init_failed"), err)
		}

		// 保存项目信息到数据库
		if err := db.AddProject(database, projectName, repoURL, remoteBranch, remoteBranch, pullBranch, false); err != nil {
			log.Fatal(i18n.T("msg.project_save_failed"), err)
		}

		// 切回原始目录
		if err := os.Chdir(currentDir); err != nil {
			log.Fatal(i18n.T("msg.dir_back_failed"), err)
		}

		fmt.Printf(i18n.T("msg.project_init_done"), projectName)
		fmt.Printf(i18n.T("msg.project_path"), projectDir)
		os.Exit(0)
	case "proxy":
		// 获取项目根目录
		parentDir := os.Getenv("APP_PATH")
		if parentDir == "" {
			log.Fatal(i18n.T("msg.dir_get_failed"))
		}

		// 获取目录中的所有文件夹
		entries, err := os.ReadDir(parentDir)
		if err != nil {
			log.Fatal(i18n.T("error.read_dir"), err)
		}

		// 从数据库获取所有项目
		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}

		// 筛选出实际存在的项目
		var existingProjects []db.Project
		for _, entry := range entries {
			if entry.IsDir() {
				for _, p := range projects {
					if p.Name == entry.Name() {
						existingProjects = append(existingProjects, p)
						break
					}
				}
			}
		}

		if len(existingProjects) == 0 {
			log.Fatal(i18n.T("error.no_projects"))
		}

		// 选择项目
		projectNames := make([]string, len(existingProjects))
		for i, p := range existingProjects {
			projectNames[i] = p.Name
		}

		prompt := promptui.Select{
			Label: i18n.T("msg.select_project"),
			Items: projectNames,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "> {{ . | cyan }}",
				Inactive: "  {{ . }}",
				Selected: "✓ {{ . | green }}",
			},
		}

		index, _, err := prompt.Run()
		if err != nil {
			log.Fatal(i18n.T("error.invalid_selection"))
		}
		selectedProject := existingProjects[index]

		// 切换到项目目录并设置代理
		err = executeInProjectDir(&selectedProject, func() error {
			return git.SetProxy()
		})
		if err != nil {
			log.Fatal(i18n.T("error.set_proxy"), err)
		}

		os.Exit(0)
	case "pull":
		// 重新初始化数据库连接
		database.Close()
		// 先同步数据库
		if err := syncDatabase("pull"); err != nil {
			log.Fatal(err)
		}

		database, err = db.InitDB(dbFilePath)
		if err != nil {
			log.Fatal(i18n.T("cmd.set"), err)
		}
		defer database.Close()
		// 检查并更新 pull_branch 字段
		if err := db.CheckAndUpdatePullBranch(database); err != nil {
			fmt.Printf(i18n.T("error.update_pull_branch"), "", err)
		}
		var projectName string
		if len(os.Args) > 2 {
			projectName = os.Args[2]
		} else {
			var err error

			projectName, err = selectProject(database, true)
			if err != nil {
				log.Fatal(err)
			}
		}

		if projectName == "ALL" {
			projects, err := db.ListProjects(database)
			if err != nil {
				log.Fatal(i18n.T("msg.project_list_failed"), err)
			}

			if len(projects) == 0 {
				log.Fatal(i18n.T("error.no_projects"))
			}

			for _, project := range projects {
				fmt.Printf(i18n.T("msg.pulling_project"), project.Name)

				// 检查项目目录是否存在，如果不存在则初始化
				projectDir := filepath.Join(os.Getenv("APP_PATH"), project.Name)
				if _, err := os.Stat(projectDir); os.IsNotExist(err) {
					fmt.Printf(i18n.T("msg.project_not_exist_init"), project.Name)
					if err := initProjectFromDB(&project, projectDir); err != nil {
						fmt.Printf(i18n.T("error.project_init_failed"), project.Name, err)
						continue
					}
				}

				err := executeInProjectDir(&project, func() error {
					return git.GitPull(project.LocalBranch, project.PullBranch, false)
				})
				if err != nil {
					fmt.Printf(i18n.T("error.pull_project_failed"), project.Name, err)
					continue
				}
				fmt.Printf(i18n.T("msg.pull_success"), project.Name)
			}
			return
		}

		project, err := db.GetProject(database, projectName)
		if err != nil {
			log.Fatal(i18n.T("error.project_not_found"))
		}

		// 检查项目目录是否存在，如果不存在则初始化
		projectDir := filepath.Join(os.Getenv("APP_PATH"), project.Name)
		if _, err := os.Stat(projectDir); os.IsNotExist(err) {
			fmt.Printf(i18n.T("msg.project_not_exist_init"), project.Name)
			if err := initProjectFromDB(project, projectDir); err != nil {
				log.Fatal(i18n.T("error.project_init_failed"), project.Name, err)
			}
		}

		// 添加强制拉取选项
		type PullOption struct {
			Name  string
			Force bool
		}
		pullOptions := []PullOption{
			{Name: i18n.T("msg.pull_normal"), Force: false},
			{Name: i18n.T("msg.pull_force"), Force: true},
		}

		pullPrompt := promptui.Select{
			Label: i18n.T("prompt.select_pull_mode"),
			Items: pullOptions,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "> {{ .Name | cyan }}",
				Inactive: "  {{ .Name }}",
				Selected: "✓ {{ .Name | green }}",
			},
		}

		pullIndex, _, err := pullPrompt.Run()
		if err != nil {
			log.Fatal(i18n.T("msg.pull_mode_select_failed"), err)
		}
		isForce := pullOptions[pullIndex].Force

		err = executeInProjectDir(project, func() error {
			// 获取默认分支
			remoteBranches, err := git.GetDefaultRemoteBranch(project.RepositoryURL)
			if err != nil {
				fmt.Printf(i18n.T("msg.fetch_branches_failed"), err)
				return err
			}
			if len(remoteBranches) == 0 {
				// 如果没有获取到分支，使用默认的 main 分支
				return git.GitPull(project.LocalBranch, project.RemoteBranch, isForce)
			}

			for i, branche := range remoteBranches {
				if branche == project.PullBranch {
					// 将找到的分支与第一个元素交换位置
					remoteBranches[0], remoteBranches[i] = remoteBranches[i], remoteBranches[0]
					break
				}
			}

			prompt := promptui.Select{
				Label: i18n.T("prompt.select_branch") + "(" + project.PullBranch + ")",
				Items: remoteBranches,
				Templates: &promptui.SelectTemplates{
					Label: "{{ . }}",

					Active:   "> {{ . | cyan }}",
					Inactive: "  {{ . }}",
					Selected: "✓ {{ . | green }}",
				},
			}
			index, _, err := prompt.Run()
			if err != nil {
				fmt.Printf(i18n.T("msg.branch_select_failed"), err)
				return err
			}
			var pullBranch = remoteBranches[index]
			if pullBranch != project.PullBranch {
				//更新数据库的pull_branch
				if err := db.UpdateProjectPullBranch(database, project.Name, pullBranch); err != nil {
					fmt.Printf(i18n.T("error.update_pull_branch"), project.Name, err)
					return err
				}
				//env判断是不是远程资料库
				if os.Getenv("DB_ENABLED") == "true" {
					//是远程资料库
					if err := syncDatabase("push"); err != nil {
						log.Fatal(err)
					}
				}
			}
			return git.GitPull(project.LocalBranch, pullBranch, isForce)
		})
		if err != nil {
			log.Fatal(i18n.T("error.pull_failed"), err)
		}
		os.Exit(0)
	case "push":
		var projectName string
		if len(os.Args) > 2 {
			projectName = os.Args[2]
		} else {
			var err error
			projectName, err = selectProject(database, true)
			if err != nil {
				log.Fatal(err)
			}
		}
		// 获取提交信息
		var message string
		if len(os.Args) > 3 {
			message = os.Args[3]
		} else {
			fmt.Println(i18n.T("prompt.commit_message") + "\n")
			message = readLine()
		}
		if projectName == "ALL" {
			projects, err := db.ListProjects(database)
			if err != nil {
				log.Fatal(i18n.T("msg.project_list_failed"), err)
			}

			if len(projects) == 0 {
				log.Fatal(i18n.T("error.no_projects"))
			}

			for _, project := range projects {
				fmt.Printf(i18n.T("msg.pushing_project"), project.Name)
				if message == "" {
					message = fmt.Sprintf(i18n.T("msg.push_by"), machineID)
				} else {
					message = fmt.Sprintf(i18n.T("msg.push_with_message"), message, machineID)
				}
				err := executeInProjectDir(&project, func() error {
					return git.GitPush(project.LocalBranch, project.RemoteBranch, message)
				})
				if err != nil {
					fmt.Printf(i18n.T("error.git_push_failed"), project.Name, err)
					continue
				}
				if err := db.UpdateLastPushWithMessage(database, project.Name, machineID, message); err != nil {
					fmt.Printf(i18n.T("error.update_commit_failed"), project.Name, err)
					continue
				}

				fmt.Printf(i18n.T("msg.push_success"), project.Name)
			}

			// 在成功推送后同步数据库
			if err := syncDatabase("push"); err != nil {
				log.Fatal(err)
			}
			return
		}

		project, err := db.GetProject(database, projectName)
		if err != nil {
			log.Fatal(i18n.T("error.project_not_found"))
		}

		if message == "" {
			message = fmt.Sprintf(i18n.T("msg.push_by"), machineID)
		} else {
			message = fmt.Sprintf(i18n.T("msg.push_with_message"), message, machineID)
		}

		err = executeInProjectDir(project, func() error {
			return git.GitPush(project.LocalBranch, project.RemoteBranch, message)
		})
		if err != nil {
			log.Fatal(i18n.T("error.push_failed"), err)
		}

		if err := db.UpdateLastPushWithMessage(database, projectName, machineID, message); err != nil {
			log.Fatal(i18n.T("error.update_last_push"), err)
		}

		// 在成功推送后同步数据库
		if err := syncDatabase("push"); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	case "list":
		// 获取项目根目录
		parentDir := os.Getenv("APP_PATH")

		// 获取目录中的所有文件夹
		entries, err := os.ReadDir(parentDir)
		if err != nil {
			log.Fatal(i18n.T("error.read_dir"), err)
		}

		// 从数据库获取所有项目
		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}

		// 筛选出实际存在的项目
		var existingProjects []db.Project
		for _, entry := range entries {
			if entry.IsDir() {
				for _, p := range projects {
					if p.Name == entry.Name() {
						existingProjects = append(existingProjects, p)
						break
					}
				}
			}
		}

		if len(existingProjects) == 0 {
			fmt.Printf(i18n.T("msg.no_projects_in_dir"), parentDir)
			return
		}

		fmt.Println(i18n.T("msg.project_list_header"))
		fmt.Println("----------------------------------------")
		for _, p := range existingProjects {
			fmt.Printf(i18n.T("msg.project_name"), p.Name)
			fmt.Printf(i18n.T("msg.repo_url"), p.RepositoryURL)
			fmt.Printf(i18n.T("msg.local_branch"), p.LocalBranch)
			fmt.Printf(i18n.T("msg.remote_branch_info"), p.RemoteBranch)
			if p.LastPushMessage != "" {
				fmt.Printf(i18n.T("msg.last_push_message"), p.LastPushMessage)
			}
			if p.LastPushTime.IsZero() {
				fmt.Println(i18n.T("msg.no_push_history"))
			} else {
				if p.LastMachineID != "" {
					fmt.Printf(i18n.T("msg.last_push_info"),
						p.LastPushTime.Format("2006-01-02 15:04:05"),
						p.LastMachineID)

				} else {
					fmt.Printf(i18n.T("msg.last_push_time"),
						p.LastPushTime.Format("2006-01-02 15:04:05"))
				}
			}
			fmt.Println("----------------------------------------")
		}

	case "delete":
		// 选择要删除的项目，不显示"全部项目"选项
		projectName, err := selectProject(database, false)
		if err != nil {
			log.Fatal(err)
		}

		// 获取项目信息用于显示
		project, err := db.GetProject(database, projectName)
		if err != nil {
			log.Fatal(i18n.T("msg.project_info_get_failed"), err)
		}

		// 显示项目信息并请求确认
		fmt.Println(i18n.T("msg.delete_confirm_header"))
		fmt.Printf(i18n.T("msg.project_name"), project.Name)
		fmt.Printf(i18n.T("msg.repo_url"), project.RepositoryURL)
		fmt.Printf(i18n.T("msg.delete_confirm_prompt"))

		var confirm string
		fmt.Scanln(&confirm)

		if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
			if err := db.DeleteProject(database, projectName); err != nil {
				log.Fatal(i18n.T("error.delete_project"), err)
			}
			fmt.Printf(i18n.T("msg.project_deleted"), projectName)
		} else {
			fmt.Println(i18n.T("msg.operation_cancelled"))
		}

		//env判断是不是远程资料库
		if os.Getenv("DB_ENABLED") == "true" {
			//是远程资料库
			if err := syncDatabase("push"); err != nil {
				log.Fatal(err)
			}
		}

	case "set":
		var setting string
		if len(os.Args) > 2 {
			setting = os.Args[2]
			switch setting {
			case "machine":
				machineID := os.Getenv("MACHINE_ID")
				if machineID == "" {
					machineID = "machine-01"
				}
				if len(os.Args) > 3 {
					machineID = os.Args[3]
				}
				if err := updateEnvFile("MACHINE_ID", machineID); err != nil {
					log.Fatal(i18n.T("msg.machine_id_update_failed"), err)
				}
				fmt.Printf(i18n.T("msg.machine_id_updated"), machineID)

			case "path":
				currentDir, err := os.Getwd()
				if err != nil {
					log.Fatal(i18n.T("msg.dir_get_failed"), err)
				}
				appPath := filepath.Dir(currentDir)
				if len(os.Args) > 3 {
					appPath = os.Args[3]
				}
				if err := updateEnvFile("APP_PATH", appPath); err != nil {
					log.Fatal(i18n.T("msg.app_path_update_failed"), err)
				}
				fmt.Printf(i18n.T("msg.app_path_updated_to"), appPath)

			case "lang":
				lang := os.Getenv("MGIT_LANG")
				if lang == "" {
					lang = "zh-CN"
				}
				if len(os.Args) > 3 {
					lang = os.Args[3]
				}
				if err := updateEnvFile("MGIT_LANG", lang); err != nil {
					log.Fatal(i18n.T("msg.lang_update_failed"), err)
				}
				i18n.SetLanguage(lang)
				fmt.Printf(i18n.T("msg.lang_updated"), lang)

			case "pull-branch":
				var projectName string
				var branchName string
				if len(os.Args) > 3 {
					projectName = os.Args[3]
					branchName = os.Args[4]
				} else {
					var err error
					projectName, err = selectProject(database, false)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Printf(i18n.T("prompt.branch_name") + " (" + projectName + "): ")
					branchName = readLine()
				}
				if projectName == "" || branchName == "" {
					log.Fatal(i18n.T("error.empty_project_name_or_branch"))
				}
				if err := db.UpdateProjectPullBranch(database, projectName, branchName); err != nil {
					log.Fatal(i18n.T("error.update_pull_branch"), err)
				}
				//env判断是不是远程资料库
				if os.Getenv("DB_ENABLED") == "true" {
					//是远程资料库
					if err := syncDatabase("push"); err != nil {
						log.Fatal(err)
					}
				}
				fmt.Printf(i18n.T("msg.pull_branch_updated"), projectName, branchName)

			default:
				fmt.Println(i18n.T("error.invalid_setting"))
			}
		} else {
			if err := handleSettings(); err != nil {
				log.Fatal(err)
			}
		}

		return

	case "update":
		if err := checkAndUpdateSelf(database, machineID); err != nil {
			log.Fatal(i18n.T("msg.app_update_failed"), err)
		}
		return

	case "v", "-v", "--version":
		fmt.Printf(i18n.T("msg.version"), "1.0.3")
		fmt.Println()
		fmt.Println(i18n.T("msg.version_aliases"))
		fmt.Println("https://github.com/liliangshan/mgit")
		return

	case "branch":
		var projectName string
		var localBranch string
		var remoteBranch string
		var project *db.Project

		// 如果命令行提供了完整参数
		if len(os.Args) > 4 {
			projectName = os.Args[2]
			localBranch = os.Args[3]
			remoteBranch = os.Args[4]

			// 获取项目信息
			var err error
			project, err = db.GetProject(database, projectName)
			if err != nil {
				log.Fatal(i18n.T("error.project_not_found"))
			}
			//env判断是不是远程资料库
			if os.Getenv("DB_ENABLED") == "true" {
				//是远程资料库
				if err := syncDatabase("push"); err != nil {
					log.Fatal(err)
				}
			}
		} else {
			// 交互式选择项目
			var err error
			projectName, err = selectProject(database, false)
			if err != nil {
				log.Fatal(err)
			}

			// 获取项目信息
			project, err = db.GetProject(database, projectName)
			if err != nil {
				log.Fatal(i18n.T("error.project_not_found"))
			}
			//获取项目目录
			projectDir := getProjectWorkDir(project)
			//切换到项目目录
			err = os.Chdir(projectDir)
			if err != nil {
				log.Fatal(i18n.T("error.chdir_failed"), err)
			}

			// 获取远程分支列表
			fmt.Println(i18n.T("msg.fetching_branches"))
			remoteBranches, err := getRemoteBranches(project.RepositoryURL)
			if err != nil {
				log.Fatal(i18n.T("msg.fetch_branches_failed"), err)
			}

			// 选择远程分支
			prompt := promptui.Select{
				Label: i18n.T("prompt.select_remote_branch"),
				Items: remoteBranches,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   "> {{ . | cyan }}",
					Inactive: "  {{ . }}",
					Selected: "✓ {{ . | green }}",
				},
			}

			index, _, err := prompt.Run()
			if err != nil {
				log.Fatal(i18n.T("msg.branch_select_failed"), err)
			}
			remoteBranch = remoteBranches[index]

			// 输入本地分支名称
			fmt.Print(i18n.T("prompt.local_branch_name") + " (" + remoteBranch + "): ")

			// 选择远程分支
			localPrompt := promptui.Select{
				Label: i18n.T("prompt.select_remote_branch"),
				Items: remoteBranches,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   "> {{ . | cyan }}",
					Inactive: "  {{ . }}",
					Selected: "✓ {{ . | green }}",
				},
			}

			localIndex, _, err := localPrompt.Run()
			if err != nil {
				log.Fatal(i18n.T("msg.branch_select_failed"), err)
			}

			localBranch = remoteBranches[localIndex]
			if localBranch == "" {
				localBranch = remoteBranch
			}
			//修改本地分支名和选择的一样
			if err := git.GitOperation(true, "git", "branch", "-m", localBranch); err != nil {
				log.Fatal(i18n.T("error.branch_rename"), err)
			}
		}

		// 更新数据库中的分支信息
		if err := db.UpdateProject(database, projectName, project.RepositoryURL, localBranch, remoteBranch); err != nil {
			log.Fatal(i18n.T("error.update_project_branches"), err)
		}
		//env判断是不是远程资料库
		if os.Getenv("DB_ENABLED") == "true" {
			//是远程资料库
			if err := syncDatabase("push"); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf(i18n.T("msg.branch_updated_success"), projectName, localBranch, remoteBranch)
		//pull-branch
		os.Exit(0)
	case "cd":
		var projectName string
		if len(os.Args) > 2 {
			projectName = os.Args[2]
		} else {
			var err error
			projectName, err = selectProject(database, false)
			if err != nil {
				log.Fatal(err)
			}
		}

		// 获取项目信息
		project, err := db.GetProject(database, projectName)
		if err != nil {
			log.Fatal(i18n.T("error.project_not_found"))
		}

		// 获取项目目录
		projectDir := getProjectWorkDir(project)

		// 对于 Windows，生成一个批处理文件来切换目录
		if runtime.GOOS == "windows" {
			batFile := filepath.Join(os.TempDir(), "mgit_cd.bat")
			content := fmt.Sprintf("@echo off\ncd /d %s\ncmd.exe", projectDir)
			if err := os.WriteFile(batFile, []byte(content), 0755); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.create_cd_bat_file"), err))
			}
			defer os.Remove(batFile)

			// 启动批处理文件
			cmd := exec.Command("cmd.exe", "/c", batFile)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.change_directory_failed"), err))
			}
		} else {
			// Unix-like 系统
			cmd := exec.Command("sh", "-c", fmt.Sprintf("cd %s && $SHELL", projectDir))
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.change_directory_failed"), err))
			}
		}

		return

	case "env":
		// 直接使用 appExePath 作为值
		if err := setSystemEnv("MGIT_PATH", appExePath); err != nil {
			log.Fatal(fmt.Errorf(i18n.T("error.set_env_failed"), err))
		}
		return

	case "new":
		// 获取新应用名称
		var newAppName string
		if len(os.Args) > 2 {
			newAppName = os.Args[2]
		} else {
			fmt.Print(i18n.T("prompt.new_app_name") + ": ")
			newAppName = readLine()
		}

		// 清理应用名称
		newAppName = cleanProjectName(newAppName)
		if newAppName == "" {
			log.Fatal(i18n.T("error.empty_app_name"))
		}

		// 获取当前可执行文件路径
		currentExe, err := os.Executable()
		if err != nil {
			log.Fatal(fmt.Errorf(i18n.T("error.get_current_exe"), err))
		}

		// 构建新应用路径
		newExePath := filepath.Join(filepath.Dir(currentExe), newAppName)
		if runtime.GOOS == "windows" {
			newExePath += ".exe"
		}

		// 检查文件是否已存在，如果存在则删除
		if _, err := os.Stat(newExePath); err == nil {
			fmt.Printf(i18n.T("msg.app_exists_deleting"), newAppName)
			if err := os.Remove(newExePath); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.delete_existing_app"), err))
			}
		}

		// 复制可执行文件
		if err := copyFile(currentExe, newExePath); err != nil {
			log.Fatal(fmt.Errorf(i18n.T("error.copy_app_failed"), err))
		}

		// 设置可执行权限
		if runtime.GOOS != "windows" {
			if err := os.Chmod(newExePath, 0755); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.set_permission_failed"), err))
			}
		}

		fmt.Printf(i18n.T("msg.app_copied_success"), newAppName)

		// 启动新应用
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "/c", "start", newExePath)
			if err := cmd.Run(); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.start_app_failed"), err))
			}
		} else {
			cmd := exec.Command(newExePath)
			if err := cmd.Start(); err != nil {
				log.Fatal(fmt.Errorf(i18n.T("error.start_app_failed"), err))
			}
		}
		return

	case "pull-all":
		// 重新初始化数据库连接
		database.Close()
		// 先同步数据库
		if err := syncDatabase("pull"); err != nil {
			log.Fatal(err)
		}

		database, err = db.InitDB(dbFilePath)
		if err != nil {
			log.Fatal(i18n.T("cmd.set"), err)
		}
		defer database.Close()

		// 检查并更新 pull_branch 字段
		if err := db.CheckAndUpdatePullBranch(database); err != nil {
			fmt.Printf(i18n.T("error.update_pull_branch"), "", err)
		}

		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}

		if len(projects) == 0 {
			log.Fatal(i18n.T("error.no_projects"))
		}

		for _, project := range projects {
			fmt.Printf(i18n.T("msg.pulling_project"), project.Name)

			// 检查项目目录是否存在，如果不存在则初始化
			projectDir := filepath.Join(os.Getenv("APP_PATH"), project.Name)
			if _, err := os.Stat(projectDir); os.IsNotExist(err) {
				fmt.Printf(i18n.T("msg.project_not_exist_init"), project.Name)
				if err := initProjectFromDB(&project, projectDir); err != nil {
					fmt.Printf(i18n.T("error.project_init_failed"), project.Name, err)
					continue
				}
			}

			err := executeInProjectDir(&project, func() error {
				return git.GitPull(project.LocalBranch, project.PullBranch, false)
			})
			if err != nil {
				fmt.Printf(i18n.T("error.pull_project_failed"), project.Name, err)
				continue
			}
			fmt.Printf(i18n.T("msg.pull_success"), project.Name)
		}
		return

	case "push-all":
		// 重新初始化数据库连接
		database.Close()
		// 先同步数据库
		if err := syncDatabase("pull"); err != nil {
			log.Fatal(err)
		}

		database, err = db.InitDB(dbFilePath)
		if err != nil {
			log.Fatal(i18n.T("cmd.set"), err)
		}
		defer database.Close()

		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}

		if len(projects) == 0 {
			log.Fatal(i18n.T("error.no_projects"))
		}

		// 获取提交信息
		var message string
		if len(os.Args) > 2 {
			message = os.Args[2]
		} else {
			fmt.Println(i18n.T("prompt.commit_message") + "\n")
			message = readLine()
		}

		for _, project := range projects {
			fmt.Printf(i18n.T("msg.pushing_project"), project.Name)
			if message == "" {
				message = fmt.Sprintf(i18n.T("msg.push_by"), machineID)
			} else {
				message = fmt.Sprintf(i18n.T("msg.push_with_message"), message, machineID)
			}

			err := executeInProjectDir(&project, func() error {
				return git.GitPush(project.LocalBranch, project.RemoteBranch, message)
			})
			if err != nil {
				fmt.Printf(i18n.T("error.git_push_failed"), project.Name, err)
				continue
			}

			if err := db.UpdateLastPushWithMessage(database, project.Name, machineID, message); err != nil {
				fmt.Printf(i18n.T("error.update_commit_failed"), project.Name, err)
				continue
			}

			fmt.Printf(i18n.T("msg.push_success"), project.Name)
		}

		// 在成功推送后同步数据库
		if err := syncDatabase("push"); err != nil {
			log.Fatal(err)
		}
		return

	default:
		fmt.Println(i18n.T("error.invalid_command"))
		showHelp(appExePath)
	}
}

// 清理项目名称中的非法字符
func cleanProjectName(name string) string {
	// 移除前后空格
	name = strings.TrimSpace(name)

	// 替换非法字符为下划线
	name = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '-' || r == '_' {
			return r
		}
		return '_'
	}, name)

	return name
}

// 修改 getCurrentDir 函数,使用 appExePath 替代 MGIT_HOME
func getCurrentDir(isAppSelf bool) string {
	if isAppSelf {
		// 如果是应用自身，使用 appExePath
		if appExePath != "" {
			return appExePath
		}
		log.Fatal(i18n.T("msg.app_path_not_set"))
	}

	// 获取 APP_PATH 环境变量
	appPath := os.Getenv("APP_PATH")
	if appPath == "" {
		// 如果未设置，使用当前目录
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(i18n.T("msg.dir_get_failed"), err)
		}
		return currentDir
	}
	return appPath
}

// 修改 selectProject 函数
func selectProject(database *sql.DB, showAll bool) (string, error) {
	projects, err := db.ListProjects(database)
	if err != nil {
		return "", fmt.Errorf(i18n.T("msg.project_list_failed"), err)
	}

	var projectList []db.Project
	if showAll {
		projectList = append(projectList, db.Project{Name: i18n.T("msg.all_projects")})
	}

	// 添加所有项目
	for _, p := range projects {
		projectList = append(projectList, p)
	}

	if len(projectList) == 0 {
		return "", fmt.Errorf(i18n.T("error.no_projects"))
	}

	prompt := promptui.Select{
		Label: i18n.T("prompt.select_project"),
		Items: projectList,
		Size:  20,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "> {{ .Name | cyan }} {{ if .LastPushTime }}{{if .LastMachineID}}[{{.LastMachineID}}] {{end}}({{.LastPushTime.Format \"2006-01-02 15:04:05\"}}){{end}}",
			Inactive: "  {{ .Name }} {{ if .LastPushTime }}{{if .LastMachineID}}[{{.LastMachineID}}] {{end}}({{.LastPushTime.Format \"2006-01-02 15:04:05\"}}){{end}}",
			Selected: "✓ {{ .Name | green }} {{ if .LastPushTime }}{{if .LastMachineID}}[{{.LastMachineID}}] {{end}}({{.LastPushTime.Format \"2006-01-02 15:04:05\"}}){{end}}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf(i18n.T("msg.project_select_failed"), err)
	}

	if showAll && index == 0 {
		return "ALL", nil
	}

	if index < len(projectList) {
		return projectList[index].Name, nil
	}

	return "", fmt.Errorf(i18n.T("msg.invalid_selection"))
}

// 添加计算文件 MD5 的函数
func calculateFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// 检查并更新应用自身
func checkAndUpdateSelf(database *sql.DB, machineID string) error {
	// 删除获取 MGIT_HOME 的代码
	if appExePath == "" {
		return fmt.Errorf(i18n.T("msg.mgit_home_not_set"))
	}

	// 获取应用信息
	app, err := db.GetProject(database, "mgit")
	if err != nil {
		return fmt.Errorf(i18n.T("msg.app_info_get_failed"), err)
	}

	// 使用 appExePath 路径下的可执行文件
	execPath := filepath.Join(appExePath, "mgit")
	if runtime.GOOS == "windows" {
		execPath += ".exe"
	}

	// 计算当前可执行文件的 MD5
	currentMD5, err := calculateFileMD5(execPath)
	if err != nil {
		return fmt.Errorf(i18n.T("msg.md5_calc_failed"), err)
	}

	// 切换到应用目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
	}

	if err := os.Chdir(appExePath); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	// 拉取最新代码
	if err := git.GitPull(app.LocalBranch, app.RemoteBranch, false); err != nil {
		return fmt.Errorf(i18n.T("error.pull_latest"), err)
	}

	// 更新最后提交信息
	if err := db.UpdateLastPushWithMessage(database, "mgit", machineID, "update"); err != nil {
		return fmt.Errorf(i18n.T("error.update_last_push"), err)
	}

	// 切回原始目录
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
	}

	// 重新计算 MD5
	newMD5, err := calculateFileMD5(execPath)
	if err != nil {
		return fmt.Errorf(i18n.T("msg.md5_calc_failed"), err)
	}

	// 如果 MD5 不同，说明需要重新编译
	if currentMD5 != newMD5 {
		fmt.Println(i18n.T("msg.app_updated"))
		os.Exit(0)
	}

	fmt.Println(i18n.T("msg.app_unchanged"))
	return nil
}

// 修改项目操作时的目录切换逻辑
func executeInProjectDir(project *db.Project, action func() error) error {
	// 获取项目工作目录
	workDir := getProjectWorkDir(project)

	// 确定切换目录的命令
	var cdCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows 使用 /d 参数
		cdCmd = exec.Command("cmd", "/c", "cd", "/d", workDir)
	} else {
		// Linux/macOS 直接使用 cd
		cdCmd = exec.Command("sh", "-c", "cd "+workDir)
	}

	// 执行切换目录命令
	if err := cdCmd.Run(); err != nil {
		return fmt.Errorf("无法切换到项目目录: %v", err)
	}

	// 切换当前工作目录
	if err := os.Chdir(workDir); err != nil {
		return fmt.Errorf("无法切换到项目目录: %v", err)
	}

	// 执行传入的操作
	return action()
}

// 修改 getProjectWorkDir 函数
func getProjectWorkDir(project *db.Project) string {
	if project.IsAppSelf {
		// 如果是应用自身，直接返回 appExePath
		if appExePath != "" {
			return filepath.Clean(appExePath)
		}
		log.Fatal(i18n.T("msg.mgit_home_not_set"))
	}
	// 其他项目使用 APP_PATH/项目名
	return filepath.Join(os.Getenv("APP_PATH"), project.Name)
}

// 修改 updateEnvFile 函数
func updateEnvFile(key, value string) error {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	exePath := filepath.Dir(ex)
	envFile := getEnvFileName()

	// 读取当前环境变量文件内容
	content, err := os.ReadFile(filepath.Join(exePath, envFile))
	if err != nil {
		return fmt.Errorf(i18n.T("msg.env_read_failed"), err)
	}

	lines := strings.Split(string(content), "\n")
	found := false
	var newLines []string

	// 更新指定的键值
	for _, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			newLines = append(newLines, fmt.Sprintf("%s=%s", key, value))
			found = true
		} else {
			newLines = append(newLines, line)
		}
	}

	// 如果没有找到键，则添加
	if !found {
		newLines = append(newLines, fmt.Sprintf("%s=%s", key, value))
	}

	// 写回文件
	if err := os.WriteFile(filepath.Join(exePath, envFile), []byte(strings.Join(newLines, "\n")), 0644); err != nil {
		return fmt.Errorf(i18n.T("msg.env_write_failed"), err)
	}
	os.Setenv(key, value)
	return nil
}

// 添加显示帮助信息的函数
func showHelp(exePath string) {
	envFile := getEnvFileName()
	// 确保已加载语言文件
	if err := godotenv.Load(filepath.Join(exePath, envFile)); err == nil {
		if lang := os.Getenv("MGIT_LANG"); lang != "" {
			i18n.SetLanguage(lang)
		}
	}

	fmt.Println(i18n.T("msg.usage"))
	fmt.Println(i18n.T("msg.available_commands"))

	fmt.Printf("\n  init      - %s\n", i18n.T("cmd.init"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.example"), "mgit init project_name https://github.com/user/repo.git"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.or"), "mgit init mgit"))

	fmt.Printf("\n  pull      - %s\n", i18n.T("cmd.pull"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.example"), "mgit pull [project_name]"))
	fmt.Printf("             %s\n", i18n.T("msg.no_args_menu"))

	fmt.Printf("\n  push      - %s\n", i18n.T("cmd.push"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.example"), "mgit push [project_name]"))
	fmt.Printf("             %s\n", i18n.T("msg.no_args_menu"))

	fmt.Printf("\n  update    - %s\n", i18n.T("cmd.update"))

	fmt.Printf("\n  pull-all  - %s\n", i18n.T("cmd.pull_all"))
	fmt.Printf("  push-all  - %s\n", i18n.T("cmd.push_all"))
	fmt.Printf("  list      - %s\n", i18n.T("cmd.list"))
	fmt.Printf("  delete    - %s\n", i18n.T("cmd.delete"))

	fmt.Printf("\n  set       - %s\n", i18n.T("cmd.set"))
	fmt.Printf("             %s\n", i18n.T("msg.set_example"))
	fmt.Printf("             %s\n", i18n.T("msg.set_path_example"))
	fmt.Printf("             %s\n", i18n.T("msg.set_interactive"))

	fmt.Printf("\n  v         - %s\n", i18n.T("cmd.version"))
	fmt.Printf("             %s\n", i18n.T("msg.version_aliases"))

	fmt.Printf("\n  help      - %s\n", i18n.T("cmd.help"))
	fmt.Printf("             %s\n", i18n.T("msg.help_aliases"))

	// 添加 branch 命令的帮助信息
	fmt.Printf("\n  branch    - %s\n", i18n.T("cmd.branch"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.example"), "mgit branch project_name local_branch remote_branch"))
	fmt.Printf("             %s\n", i18n.T("msg.no_args_interactive"))

	fmt.Printf("\n  env       - %s\n", i18n.T("cmd.env"))
	fmt.Printf("             %s\n", i18n.T("msg.env_example"))

	// 在 showHelp 函数中添加
	fmt.Printf("\n  proxy     - %s\n", i18n.T("cmd.proxy"))
	fmt.Printf("             %s\n", i18n.T("msg.proxy_example"))

	// 在 showHelp 函数中添加
	fmt.Printf("\n  copy      - %s\n", i18n.T("cmd.copy"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.copy_example"), "mgit copy new_app_name"))

	// 在 showHelp 函数中修改
	fmt.Printf("\n  new      - %s\n", i18n.T("cmd.new"))
	fmt.Printf("             %s\n", fmt.Sprintf(i18n.T("msg.new_example"), "mgit new new_app_name"))
}

// 添加辅助函数用于读取用户输入
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// 添加复制文件的辅助函数
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

// 添加设置处理函数
func handleSettings() error {

	lang := os.Getenv("MGIT_LANG")
	if lang == "" {
		lang = "zh-CN"

	}
	// 语言选择
	if err := handleLanguageSelection(lang, true); err != nil {
		return err
	}
	dbRepo := os.Getenv("DB_REPO")
	dbEnabled := os.Getenv("DB_ENABLED")
	if dbEnabled == "" {
		dbEnabled = "false"
	}

	// 显示当前设置
	machineID := os.Getenv("MACHINE_ID")
	if machineID == "" {
		machineID = "machine-01"
	}
	appPath := os.Getenv("APP_PATH")
	if appPath == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
		}
		appPath = filepath.Dir(currentDir)
	}
	fmt.Printf(i18n.T("msg.current_settings") + "\n")
	fmt.Printf("machine: %s\n", machineID)
	fmt.Printf("path: %s\n", appPath)
	fmt.Printf("lang: %s\n", lang)
	fmt.Printf("db_enabled: %s\n", dbEnabled)
	if dbEnabled == "true" {
		fmt.Printf("db_repo: %s\n\n", dbRepo)
	}

	// 让用户依次输入设置
	fmt.Printf(i18n.T("prompt.machine_id")+" (%s): ", machineID)
	newMachineID := readLine()
	if newMachineID != "" {
		if err := updateEnvFile("MACHINE_ID", newMachineID); err != nil {
			return fmt.Errorf(i18n.T("msg.machine_id_update_failed"), err)
		}
		fmt.Printf(i18n.T("msg.machine_id_updated"), newMachineID)
	}

	fmt.Printf(i18n.T("prompt.app_path")+" (%s): ", appPath)
	newAppPath := readLine()
	if newAppPath != "" {
		if err := updateEnvFile("APP_PATH", newAppPath); err != nil {
			return fmt.Errorf(i18n.T("msg.app_path_update_failed"), err)
		}
	} else {
		if err := updateEnvFile("APP_PATH", appPath); err != nil {
			return fmt.Errorf(i18n.T("msg.app_path_update_failed"), err)
		}
	}

	// 数据库仓库设置 - 使用下拉选择
	type Option struct {
		Name  string
		Value string
	}
	options := []Option{
		{Name: i18n.T("msg.disabled"), Value: "false"},
		{Name: i18n.T("msg.enabled"), Value: "true"},
	}

	// 从环境变量获取当前设置，如果未设置则默认为false
	currentDbEnabled := os.Getenv("DB_ENABLED")
	if currentDbEnabled == "" {
		currentDbEnabled = "false"
	}

	// 根据当前设置调整选项顺序
	for i, option := range options {
		if option.Value == currentDbEnabled && i > 0 {
			// 将当前选中的选项移到第一位
			options[0], options[i] = options[i], options[0]
			break
		}
	}

	prompt := promptui.Select{
		Label: i18n.T("prompt.db_enabled"),
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "> {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "✓ {{ .Name | green }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.db_enabled_select_failed"), err)
	}

	newDbEnabled := options[index].Value
	if err := updateEnvFile("DB_ENABLED", newDbEnabled); err != nil {
		return fmt.Errorf(i18n.T("msg.db_enabled_update_failed"), err)
	}

	if newDbEnabled == "true" {
		// 如果启用数据库仓库，先检查是否已有配置的仓库地址
		dbRepo := os.Getenv("DB_REPO")
		if dbRepo != "" {
			// 使用下拉选择是否使用现有地址
			type Option struct {
				Name  string
				Value string
			}
			options := []Option{
				{Name: fmt.Sprintf(i18n.T("msg.use_existing_repo"), dbRepo), Value: dbRepo},
				{Name: i18n.T("msg.input_new_repo"), Value: "new"},
			}

			prompt := promptui.Select{
				Label: i18n.T("prompt.use_existing_db_repo"),
				Items: options,

				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   "> {{ .Name | cyan }}",
					Inactive: "  {{ .Name }}",
					Selected: "✓ {{ .Name | green }}",
				},
			}

			index, _, err := prompt.Run()
			if err != nil {
				return fmt.Errorf(i18n.T("msg.repo_select_failed"), err)
			}

			if index == 0 {
				// 使用现有地址初始化仓库
				if err := initDbRepo(dbRepo); err != nil {
					return err
				}
				return nil
			}
		}

		// 如果没有配置地址或用户选择输入新地址，询问新地址
		fmt.Printf(i18n.T("prompt.db_repo")+" (%s): ", dbRepo)
		newDbRepo := readLine()
		if newDbRepo != "" {
			if err := updateEnvFile("DB_REPO", newDbRepo); err != nil {
				return fmt.Errorf(i18n.T("msg.db_repo_update_failed"), err)
			}

			// 初始化数据库仓库
			if err := initDbRepo(newDbRepo); err != nil {
				return err
			}
		} else if dbRepo != "" {
			// 如果用户没有输入新地址，但存在旧地址，使用旧地址
			if err := initDbRepo(dbRepo); err != nil {
				return err
			}
		}
	}

	return nil
}

// 添加语言选择处理函数
func handleLanguageSelection(lang string, isFirst bool) error {
	// 获取i18n路径

	type Language struct {
		Name string
		Code string
	}
	languages := []Language{
		{Name: "简体中文", Code: "zh-CN"},
		{Name: "繁體中文", Code: "zh-TW"},
		{Name: "English", Code: "en-US"},
		{Name: "日本語", Code: "ja-JP"},
		{Name: "한국어", Code: "ko-KR"},
		{Name: "Français", Code: "fr-FR"},
	}

	// 根据传入的lang参数调整语言顺序
	if lang != "" {
		// 找到匹配的语言并移到第一位
		for i, language := range languages {
			if language.Code == lang {
				// 将找到的语言与第一个元素交换位置
				languages[0], languages[i] = languages[i], languages[0]
				break
			}
		}
	}

	prompt := promptui.Select{
		Label: i18n.T("prompt.select_language"),
		Items: languages,
		Size:  20,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "> {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "✓ {{ .Name | green }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.lang_select_failed"), err)
	}

	newLang := languages[index].Code
	i18n.SetLanguage(newLang)
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	exePath := filepath.Dir(ex)
	envFile := getEnvFileName()
	// 如果是第一次选择语言，需要先创建.env文件
	// 检查是否存在.env文件
	if _, err := os.Stat(filepath.Join(exePath, envFile)); os.IsNotExist(err) {
		fmt.Println(i18n.T("msg.env_not_exit"), err)

		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		currentDir := filepath.Dir(ex)
		// 创建初始的.env文件,移除 MGIT_HOME
		envContent := fmt.Sprintf("# "+i18n.T("env.machine_id")+"\nMACHINE_ID=%s\n\n# "+i18n.T("env.app_path")+"\nAPP_PATH=%s\n\n# "+i18n.T("env.mgit_lang")+"\nMGIT_LANG=%s\n\n# "+i18n.T("prompt.db_enabled")+"\nDB_ENABLED=false\nDB_REPO=\nDB_BRANCH=",
			"machine-01",
			filepath.Dir(currentDir),
			newLang)

		if err := os.WriteFile(filepath.Join(currentDir, envFile), []byte(envContent), 0644); err != nil {
			return fmt.Errorf(i18n.T("msg.env_create_failed"), err)
		}

		//加载环境变量
		if err := godotenv.Load(filepath.Join(currentDir, envFile)); err != nil {
			return fmt.Errorf(i18n.T("msg.env_load_failed"), err)
		}
	}
	if err := updateEnvFile("MGIT_LANG", newLang); err != nil {
		return fmt.Errorf(i18n.T("msg.lang_update_failed"), err)
	}

	fmt.Printf(i18n.T("msg.lang_updated"), newLang)
	return nil
}

// 添加获取远程分支的函数
func getRemoteBranches(repoUrl string) ([]string, error) {
	return git.GetDefaultRemoteBranch(repoUrl)
}

// 修改 initDbRepo 函数，简化分支选择逻辑
func initDbRepo(repoUrl string) error {

	// 获取或设置分支
	dbBranch := os.Getenv("DB_BRANCH")
	if dbBranch == "" {
		fmt.Println(i18n.T("msg.fetching_branches"))
		remoteBranches, err := getRemoteBranches(repoUrl)
		if err != nil {
			return fmt.Errorf(i18n.T("msg.fetch_branches_failed"), err)
		}

		if len(remoteBranches) == 0 {
			// 如果没有获取到分支，使用默认的 main 分支
			dbBranch = "main"
		} else {
			prompt := promptui.Select{
				Label: i18n.T("prompt.select_branch"),
				Items: remoteBranches,
				Templates: &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   "> {{ . | cyan }}",
					Inactive: "  {{ . }}",
					Selected: "✓ {{ . | green }}",
				},
			}

			index, _, err := prompt.Run()
			if err != nil {
				return fmt.Errorf(i18n.T("msg.branch_select_failed"), err)
			}

			dbBranch = remoteBranches[index]
			// 保存分支设置到环境变量
			if err := updateEnvFile("DB_BRANCH", dbBranch); err != nil {
				return fmt.Errorf(i18n.T("msg.branch_update_failed"), err)
			}
		}

		fmt.Printf(i18n.T("msg.branch_updated"), dbBranch)
	}

	dbDir := os.Getenv("APP_PATH")
	if dbDir == "" {
		fmt.Println("dbDir: ")
		return fmt.Errorf(i18n.T("msg.dir_get_failed"))
	}
	// 使用 getRemoteDbDirName() 获取正确的远程数据库目录名
	dbDir = getRemoteDbDirName()
	// 使用 appExePath 替换 selfPath
	if appExePath == "" {
		fmt.Println("selfPath: ")
		return fmt.Errorf(i18n.T("msg.dir_get_failed"))
	}

	// 检查数据库仓库目录是否已存在
	if _, err := os.Stat(dbDir); err == nil {
		// 目录存在，检查是否已经是 Git 仓库
		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
		}

		if err := os.Chdir(dbDir); err != nil {
			return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
		}

		// 检查是否是 Git 仓库
		if err := git.GitOperation(true, "git", "rev-parse", "--git-dir"); err == nil {
			// 检查远程仓库地址
			cmd := exec.Command("git", "config", "--get", "remote.origin.url")
			output, err := cmd.Output()

			if err == nil {

				existingRepo := strings.TrimSpace(string(output))

				//同时判断有没有.git文件夹
				if _, err := os.Stat(filepath.Join(dbDir, ".git")); err != nil {
					existingRepo = "null"
				}
				if existingRepo == repoUrl {

					// 如果仓库地址相同，直接返回
					if err := os.Chdir(currentDir); err != nil {
						return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
					}
					//判断数据库是否存在
					if _, err := os.Stat(filepath.Join(dbDir, getDbFileName())); err == nil {
						fmt.Println(i18n.T("msg.db_repo_exists"))
						return fmt.Errorf(i18n.T("msg.db_repo_exists"))

					} else {

						if err := git.GitPull(dbBranch, dbBranch, false); err != nil {
							return fmt.Errorf(i18n.T("msg.db_repo_pull_failed"), err)
						}

						// 复制数据库文件
						if err := copyFile(filepath.Join(appExePath, getDbFileName()), filepath.Join(dbDir, getDbFileName())); err != nil {
							return fmt.Errorf(i18n.T("msg.db_file_copy_failed"), err)
						}

						// 提交并推送数据库文件
						if err := git.GitPush(dbBranch, dbBranch, "Initial database commit"); err != nil {
							return fmt.Errorf(i18n.T("msg.db_repo_push_failed"), err)
						}

						// 切回原始目录
						if err := os.Chdir(currentDir); err != nil {
							return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
						}

						return fmt.Errorf(i18n.T("msg.db_repo_initialized"))
					}
				} else {

					if _, err := os.Stat(filepath.Join(dbDir, getDbFileName())); err == nil {

						// 判断filepath.Join(appExePath, "projects.db")是否存在
						if _, err := os.Stat(filepath.Join(appExePath, getDbFileName())); err == nil {
							// 删除数据库文件
							if err := os.Remove(filepath.Join(appExePath, getDbFileName())); err != nil {
								return fmt.Errorf(i18n.T("msg.db_file_delete_failed"), err)
							}
						}

						if err := copyFile(filepath.Join(dbDir, getDbFileName()), filepath.Join(appExePath, getDbFileName())); err != nil {
							return fmt.Errorf(i18n.T("msg.db_file_copy_failed"), err)
						}
					}
					// 复制数据库文件

					//datebase.Close()
					//删除dbDir目录
					os.RemoveAll(dbDir)

					return nil
				}

			}
		}

		// 如果不是 Git 仓库，初始化它
		_, remoteBranch, err := git.GitInit(repoUrl, dbBranch)
		if err != nil {
			return fmt.Errorf(i18n.T("msg.db_repo_init_failed"), err)
		}
		log.Println("================")
		// 判断filepath.Join(appExePath, "projects.db")是否存在
		if _, err := os.Stat(filepath.Join(appExePath, getDbFileName())); err != nil {
			// 创建数据库文件
			databaseData, err = db.CreateDB(filepath.Join(appExePath, getDbFileName()))
			if err != nil {
				return fmt.Errorf(i18n.T("msg.db_create_failed"), err)
			}
		}
		if _, err := os.Stat(filepath.Join(dbDir, getDbFileName())); err == nil {
			// 删除数据库文件
			if err := os.Remove(filepath.Join(dbDir, getDbFileName())); err != nil {
				return fmt.Errorf(i18n.T("msg.db_file_delete_failed"), err)
			}
		}
		// 复制数据库文件
		if err := copyFile(filepath.Join(appExePath, getDbFileName()), filepath.Join(dbDir, getDbFileName())); err != nil {
			return fmt.Errorf(i18n.T("msg.db_file_copy_failed"), err)
		}

		// 提交并推送数据库文件
		if err := git.GitPush(remoteBranch, dbBranch, "Initial database commit"); err != nil {
			return fmt.Errorf(i18n.T("msg.db_repo_push_failed"), err)
		}

		// 切回原始目录
		if err := os.Chdir(currentDir); err != nil {
			return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
		}

		fmt.Println(i18n.T("msg.db_repo_initialized"))
		return nil
	}

	// 如果目录不存在，创建新的仓库
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf(i18n.T("msg.db_dir_create_failed"), err)
	}

	// 切换到数据库目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
	}
	if err := os.Chdir(dbDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	// 初始化 Git 仓库
	_, iremoteBranch, err := git.GitInit(repoUrl, dbBranch)
	if err != nil {
		return fmt.Errorf(i18n.T("msg.db_repo_init_failed"), err)
	}

	// 复制数据库文件
	if err := copyFile(filepath.Join(appExePath, getDbFileName()), filepath.Join(dbDir, getDbFileName())); err != nil {
		return fmt.Errorf(i18n.T("msg.db_file_copy_failed"), err)
	}

	// 提交并推送数据库文件
	if err := git.GitPush(dbBranch, iremoteBranch, "Initial database commit"); err != nil {
		return fmt.Errorf(i18n.T("msg.db_repo_push_failed"), err)
	}

	// 切回原始目录
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
	}

	fmt.Println(i18n.T("msg.db_repo_initialized"))
	return nil
}

// 修改 syncDatabase 函数使用环境变量中的分支
func syncDatabase(operation string) error {
	// 检查是否启用了数据库仓库
	dbEnabled := os.Getenv("DB_ENABLED")
	if dbEnabled != "true" {
		return nil
	}

	// 获取配置的分支
	dbBranch := os.Getenv("DB_BRANCH")
	if dbBranch == "" {
		dbBranch = "main" // 默认使用 main 分支
	}

	// 使用 getRemoteDbDirName() 获取正确的远程数据库目录名
	dbDir := getRemoteDbDirName() //filepath.Join(os.Getenv("APP_PATH"), )
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
	}

	if err := os.Chdir(dbDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	if operation == "pull" {

		if err := git.GitOperation(true, "git", "reset", "--hard"); err != nil {
			os.Chdir(currentDir)
			return fmt.Errorf(i18n.T("msg.db_push_failed"), err)
		}
		if err := git.GitOperation(true, "git", "pull", "--rebase", "origin", dbBranch); err != nil {
			os.Chdir(currentDir)
			return fmt.Errorf(i18n.T("msg.db_push_failed"), err)
		}
		fmt.Println(i18n.T("msg.db_pulled"))
	} else if operation == "push" {

		// 提交并推送数据库更新
		if err := git.GitPush(dbBranch, dbBranch, "update-mgit-database"); err != nil {
			os.Chdir(currentDir)
			return fmt.Errorf(i18n.T("msg.db_push_failed"), err)
		}

		fmt.Println(i18n.T("msg.db_pushed"))
	}

	// 切回原始目录
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
	}

	return nil
}

// 添加新的项目初始化函数
func initProjectFromDB(project *db.Project, projectDir string) error {
	// 创建项目目录
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_create_failed"), err)
	}

	// 切换到项目目录
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
	}
	if err := os.Chdir(projectDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	// 初始化 Git 仓库并拉取代码
	if _, _, err := git.GitInit(project.RepositoryURL, project.RemoteBranch); err != nil {
		return fmt.Errorf(i18n.T("msg.git_init_failed"), err)
	}

	// 切回原始目录
	if err := os.Chdir(currentDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_back_failed"), err)
	}

	return nil
}

// changeToAppDirectory 切换到应用程序所在目录
func changeToAppDirectory() error {
	// 获取可执行文件的目录
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("无法获取可执行文件路径: %v", err)
	}

	// 获取目录
	appDir := filepath.Dir(exePath)

	// 切换目录
	if runtime.GOOS == "windows" {
		// Windows 使用 cmd /c cd /d
		cmd := exec.Command("cmd", "/c", "cd", "/d", appDir)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("无法切换到应用目录: %v", err)
		}
	} else {
		// Linux/macOS 直接使用 cd
		cmd := exec.Command("sh", "-c", "cd "+appDir)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("无法切换到应用目录: %v", err)
		}
	}

	// 使用 Go 的 os.Chdir 确保当前进程的工作目录也被更新
	if err := os.Chdir(appDir); err != nil {
		return fmt.Errorf("无法切换到应用目录: %v", err)
	}

	return nil
}

// 添加 isPowerShell 函数到文件中
func isPowerShell() bool {
	// 检查父进程名称
	pid := os.Getppid()
	cmd := exec.Command("wmic", "process", "where", fmt.Sprintf("ProcessId=%d", pid), "get", "Name")
	output, err := cmd.Output()
	if err == nil {
		processName := strings.TrimSpace(string(output))
		return strings.Contains(strings.ToLower(processName), "powershell")
	}

	// 检查 PSModulePath 环境变量
	if os.Getenv("PSModulePath") != "" {
		return true
	}

	// 检查 PROMPT 环境变量
	prompt := os.Getenv("PROMPT")
	if strings.Contains(prompt, "PS") {
		return true
	}

	return false
}

// 修改 setSystemEnv 函数,使用应用名称作为环境变量前缀
func setSystemEnv(key, value string) error {
	// 获取应用程序名称
	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf(i18n.T("error.get_app_name_failed"), err)
	}
	appName := strings.ToUpper(filepath.Base(ex))
	// 移除扩展名
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	// 构造环境变量名称
	envKey := fmt.Sprintf("%s_PATH", appName)

	// 如果没有提供值，使用应用程序所在目录
	if value == "" {
		value = appExePath
	}

	// 检查环境变量是否已存在
	if runtime.GOOS == "windows" {
		// Windows 系统检查环境变量
		cmd := exec.Command("reg", "query", `HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, "/v", envKey)
		output, err := cmd.CombinedOutput()
		if err == nil && strings.Contains(string(output), envKey) {
			existingValue := strings.TrimSpace(strings.Split(string(output), "REG_SZ")[1])
			if existingValue == value {
				fmt.Printf(i18n.T("msg.env_already_exists"), envKey, value)
				return nil
			}
		}

		// 设置环境变量
		if isPowerShell() {
			cmd := exec.Command("powershell", "-Command",
				fmt.Sprintf(`[System.Environment]::SetEnvironmentVariable('%s', '%s', 'Machine')`, envKey, value))
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf(i18n.T("error.set_env_failed"), err, string(out))
			}
		} else {
			cmd := exec.Command("setx", "/M", envKey, value)
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf(i18n.T("error.set_env_failed"), err, string(out))
			}
		}
	} else {
		// Linux/macOS 检查 /etc/environment
		content, err := os.ReadFile("/etc/environment")
		if err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, envKey+"=") {
					existingValue := strings.Trim(strings.TrimPrefix(line, envKey+"="), `"`)
					if existingValue == value {
						fmt.Printf(i18n.T("msg.env_already_exists"), envKey, value)
						return nil
					}
				}
			}
		}

		// 设置环境变量
		cmd := exec.Command("sudo", "sh", "-c",
			fmt.Sprintf(`echo '%s="%s"' >> /etc/environment`, envKey, value))
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf(i18n.T("error.set_env_failed"), err, string(out))
		}
	}

	// 立即生效
	if err := os.Setenv(envKey, value); err != nil {
		return fmt.Errorf(i18n.T("error.set_current_env_failed"), err)
	}

	fmt.Printf(i18n.T("msg.env_set_success"), envKey, value)
	fmt.Println(i18n.T("msg.env_restart_required"))
	return nil
}

func syncDatabaseToRemote() error {
	remoteDbDir := getRemoteDbDirName()
	dbName := getDbFileName()

	// 检查远程数据库目录是否存在
	if _, err := os.Stat(remoteDbDir); os.IsNotExist(err) {
		return fmt.Errorf("remote database directory not found: %v", err)
	}

	// 复制数据库文件到远程目录
	srcPath := filepath.Join(appExePath, dbName)
	dstPath := filepath.Join(remoteDbDir, dbName)

	if err := copyFile(srcPath, dstPath); err != nil {
		return fmt.Errorf("failed to sync database to remote: %v", err)
	}

	return nil
}

func pullDatabaseFromRemote() error {
	remoteDbDir := getRemoteDbDirName()
	dbName := getDbFileName()

	// 检查远程数据库文件是否存在
	remotePath := filepath.Join(remoteDbDir, dbName)
	if _, err := os.Stat(remotePath); os.IsNotExist(err) {
		return fmt.Errorf("remote database file not found: %v", err)
	}

	// 复制远程数据库文件到本地
	localPath := filepath.Join(appExePath, dbName)
	if err := copyFile(remotePath, localPath); err != nil {
		return fmt.Errorf("failed to pull database from remote: %v", err)
	}

	return nil
}
