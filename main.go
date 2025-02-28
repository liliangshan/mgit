package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"database/sql"
	"os/exec"
	"crypto/md5"
	"io"
	"bufio"
	"runtime"

	"github.com/joho/godotenv"
	"gitmanager/git"
	"gitmanager/db"
	"github.com/manifoldco/promptui"
	"gitmanager/i18n"
)

func main() {
	
	// 先检查是否是帮助命令
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "h", "-h", "-help", "help":
			isTemp := false
			// 检查 .env 文件是否存在，如果不存在则让用户选择语言
			if _, err := os.Stat(".env"); os.IsNotExist(err) {
				isTemp = true
				// 获取当前目录作为 MGIT_HOME
				currentDir, err := os.Getwd()
				if err != nil {
					log.Fatal("获取当前目录失败:", err)
				}

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
MGIT_HOME=%s`, languages[index].Code, currentDir)
				if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
					log.Fatal("创建临时语言文件失败:", err)
				}

				// 2. 加载 .env 文件
				if err := godotenv.Load(); err != nil {
					log.Fatal("加载语言设置失败:", err)
				}

				// 3. 设置语言
				i18n.SetLanguage(languages[index].Code)
			} else {
				// 如果 .env 文件存在，加载它
				if err := godotenv.Load(); err != nil {
					log.Fatal(i18n.T("msg.env_load_failed"))
				}
				// 重新设置语言
				lang := os.Getenv("MGIT_LANG")
				i18n.SetLanguage(lang)
			}

			showHelp()
			// 如果是临时文件，显示帮助后删除
			if isTemp {
				os.Remove(".env")
			}
			return
		}
	}

	// 检查 .env 文件是否存在，如果不存在则让用户选择语言并执行 set 命令
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
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
		
		// 设置语言
		i18n.SetLanguage(languages[index].Code)
		var machineID string
		// 执行设置
		fmt.Println(i18n.T("prompt.machine_id")+" (machine-01): ")
		machineID = readLine()
		if machineID == "" {
			machineID = "machine-01"
		}
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(i18n.T("msg.dir_get_failed"), err)
		}
		fmt.Println(i18n.T("prompt.app_path")+" ("+filepath.Dir(currentDir)+"): ", )
		
		appPath := readLine()
		if appPath == "" {
			appPath = filepath.Dir(currentDir)
		}

		// 创建 .env 文件内容
		envContent := fmt.Sprintf(`# %s
MACHINE_ID=%s

# %s
APP_PATH=%s

# %s
MGIT_HOME=%s

# %s
MGIT_LANG=%s`,
			i18n.T("env.machine_id"),
			machineID,
			i18n.T("env.app_path"),
			appPath,
			i18n.T("env.mgit_home"),
			currentDir,
			i18n.T("env.mgit_lang"),
			languages[index].Code)

		// 写入文件
		if err := os.WriteFile(".env", []byte(envContent), 0644); err != nil {
			log.Fatal(i18n.T("msg.env_create_failed"), err)
		}
	}
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Fatal(i18n.T("msg.env_load_failed"))
	}
	lang := os.Getenv("MGIT_LANG")
	if lang == "" {
		lang = "zh-CN"
	}
	i18n.SetLanguage(lang)	
	// 初始化数据库
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(i18n.T("msg.db_init_failed"), err)
	}
	defer database.Close()

	

	// 获取机器标识
	machineID := os.Getenv("MACHINE_ID")
	if machineID == "" {
		log.Fatal(i18n.T("msg.machine_id_not_set"))
	}

	// 检查是否已初始化自身应用
	if _, err := db.GetProject(database, "mgit"); err != nil {
		if len(os.Args) > 1 && os.Args[1] == "init" {
			// 如果是 init 命令，让它继续执行
			goto HANDLE_COMMAND
		}
		fmt.Println(i18n.T("msg.app_not_init"))
		fmt.Println(i18n.T("msg.app_not_init_or"))
		return
	}

HANDLE_COMMAND:
	// 根据命令行参数执行不同操作
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	switch os.Args[1] {
	case "h", "-h", "-help", "help":
		showHelp()
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
		
		// 如果名称为空或者是 "mgit"，则初始化/更新自身
		if projectName == "" || projectName == "mgit" {
			// 获取当前目录的 Git 仓库地址
			if err := git.GitOperation("git", "rev-parse", "--git-dir"); err != nil {
				log.Fatal(i18n.T("msg.not_git_repo"))
			}

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

			// 更新应用安装路径
			if err := updateEnvFile("MGIT_HOME", currentDir); err != nil {
				log.Fatal(i18n.T("msg.env_update_failed"), err)
			}
			fmt.Println(i18n.T("msg.app_path_updated"))

			// 更新或添加应用信息到数据库
			if err := db.UpdateProject(database, "mgit", repoURL, currentBranch, currentBranch); err != nil {
				if err := db.AddProject(database, "mgit", repoURL, currentBranch, currentBranch, true); err != nil {
					log.Fatal(i18n.T("msg.app_save_failed"), err)
				}
				fmt.Println(i18n.T("msg.app_init_done"))
			} else {
				fmt.Println(i18n.T("msg.app_info_updated"))
			}
			return
		} else {
			// 检查项目目录是否已存在
			projectDir := filepath.Join(getCurrentDir(false), projectName)
			if _, err := os.Stat(projectDir); err == nil {
				// 目录已存在，检查是否是 Git 仓库
				if err := os.Chdir(projectDir); err != nil {
					log.Fatal(i18n.T("msg.dir_change_failed"), err)
				}
				
				// 检查是否是 Git 仓库
				if err := git.GitOperation("git", "rev-parse", "--git-dir"); err == nil {
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
					if err := db.AddProject(database, projectName, repoURL, currentBranch, currentBranch, false); err != nil {
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

			// 创建项目目录
			if err := os.MkdirAll(projectDir, 0755); err != nil {
				log.Fatal(i18n.T("msg.dir_create_failed"), err)
			}

			// 切换到项目目录
			if err := os.Chdir(projectDir); err != nil {
				log.Fatal(i18n.T("msg.dir_change_failed"), err)
			}

			// 初始化 Git 仓库并拉取代码
			if err := git.GitInit(repoURL); err != nil {
				log.Fatal(i18n.T("msg.git_init_failed"), err)
			}

			// 获取远程分支
			cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
			output, err := cmd.Output()
			if err != nil {
				log.Fatal(i18n.T("msg.branch_get_failed"), err)
			}
			remoteBranch := strings.TrimSpace(string(output))

			// 保存项目信息到数据库
			if err := db.AddProject(database, projectName, repoURL, remoteBranch, remoteBranch, false); err != nil {
				log.Fatal(i18n.T("msg.project_save_failed"), err)
			}

			// 切回原始目录
			if err := os.Chdir(currentDir); err != nil {
				log.Fatal(i18n.T("msg.dir_back_failed"), err)
			}

			fmt.Printf(i18n.T("msg.project_init_done"), projectName)
			fmt.Printf(i18n.T("msg.project_path"), projectDir)
		}

	case "pull":
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
				err := executeInProjectDir(&project, func() error {
					return git.GitPull(project.LocalBranch, project.RemoteBranch)
				})
				if err != nil {
					fmt.Printf(i18n.T("error.pull_project_failed"), project.Name, err)
				} else {
					fmt.Printf(i18n.T("msg.pull_success"), project.Name)
				}
			}
			return
		}

		project, err := db.GetProject(database, projectName)
		if err != nil {
			log.Fatal(i18n.T("error.project_not_found"))
		}

		err = executeInProjectDir(project, func() error {
			return git.GitPull(project.LocalBranch, project.RemoteBranch)
		})
		if err != nil {
			log.Fatal(i18n.T("error.pull_failed"), err)
		}

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
			
			fmt.Println(i18n.T("prompt.commit_message")+"\n")
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
				message := fmt.Sprintf(i18n.T("msg.push_by"), machineID)
				err := executeInProjectDir(&project, func() error {
					return git.GitPush(project.LocalBranch, project.RemoteBranch, message)
				})
				if err != nil {
					fmt.Printf(i18n.T("error.git_push_failed"), project.Name, err)
					continue
				}
				if err := db.UpdateLastPush(database, project.Name, machineID); err != nil {
					fmt.Printf(i18n.T("error.update_commit_failed"), project.Name, err)
					continue
				}
				fmt.Printf(i18n.T("msg.push_success"), project.Name)
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

		if err := db.UpdateLastPush(database, projectName, machineID); err != nil {
			log.Fatal(i18n.T("error.update_last_push"), err)
		}

	case "pull-all":
		// 获取所有本地存在的项目
		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}
		
		if len(projects) == 0 {
			log.Fatal(i18n.T("error.no_projects"))
		}
		
		// 遍历所有项目执行 pull
		for _, project := range projects {
			fmt.Printf(i18n.T("msg.pulling_project"), project.Name)
			
			err := executeInProjectDir(&project, func() error {
				return git.GitPull(project.LocalBranch, project.RemoteBranch)
			})
			
			if err != nil {
				fmt.Printf(i18n.T("error.pull_project_failed"), project.Name, err)
				continue
			}
			
			fmt.Printf(i18n.T("msg.pull_success"), project.Name)
		}

	case "push-all":
		// 获取所有本地存在的项目
		projects, err := db.ListProjects(database)
		if err != nil {
			log.Fatal(i18n.T("msg.project_list_failed"), err)
		}
		
		if len(projects) == 0 {
			log.Fatal(i18n.T("error.no_projects"))
		}

		// 获取提交说明
		fmt.Println(i18n.T("prompt.commit_message"))
		message := readLine()
		if message == "" {
			message = fmt.Sprintf(i18n.T("msg.push_by"), machineID)
		} else {
			message = fmt.Sprintf(i18n.T("msg.push_with_message"), message, machineID)
		}
		
		// 遍历所有项目执行 push
		for _, project := range projects {
			fmt.Printf(i18n.T("msg.pushing_project"), project.Name)
			
			err := executeInProjectDir(&project, func() error {
				return git.GitPush(project.LocalBranch, project.RemoteBranch, message)
			})
			if err != nil {
				fmt.Printf(i18n.T("error.git_push_failed"), project.Name, err)
				continue
			}

			if err := db.UpdateLastPush(database, project.Name, machineID); err != nil {
				fmt.Printf(i18n.T("error.update_commit_failed"), project.Name, err)
				continue
			}
			
			fmt.Printf(i18n.T("msg.push_success"), project.Name)
		}

	case "list":
		// 获取项目根目录
		parentDir := getCurrentDir(false)

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

	case "set":
		var setting string
		if len(os.Args) > 2 {
			setting = os.Args[2]
		} else {
			// 显示当前设置
			machineID := os.Getenv("MACHINE_ID")
			if machineID == "" {
				machineID = "machine-01"
			}
			appPath := os.Getenv("APP_PATH")
			if appPath == "" {
				currentDir, err := os.Getwd()
				if err != nil {
					log.Fatal(i18n.T("msg.dir_get_failed"), err)
				}
				appPath = filepath.Dir(currentDir)
			}
			lang := os.Getenv("MGIT_LANG")
			if lang == "" {
				lang = "zh-CN"
			}

			fmt.Printf(i18n.T("msg.current_settings")+"\n")
			fmt.Printf("machine: %s\n", machineID)
			fmt.Printf("path: %s\n", appPath)
			fmt.Printf("lang: %s\n\n", lang)

			// 让用户依次输入设置
			fmt.Printf(i18n.T("prompt.machine_id")+" (%s): ", machineID)
			newMachineID := readLine()
			if newMachineID != "" {
				if err := updateEnvFile("MACHINE_ID", newMachineID); err != nil {
					log.Fatal(i18n.T("msg.machine_id_update_failed"), err)
				}
				fmt.Printf(i18n.T("msg.machine_id_updated"), newMachineID)
			}

			fmt.Printf(i18n.T("prompt.app_path")+" (%s): ", appPath)
			newAppPath := readLine()
			if newAppPath != "" {
				if err := updateEnvFile("APP_PATH", newAppPath); err != nil {
					log.Fatal(i18n.T("msg.app_path_update_failed"), err)
				}
				fmt.Printf(i18n.T("msg.app_path_updated_to"), newAppPath)
			}

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
				Label: i18n.T("prompt.select_language"),
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
				log.Fatal(i18n.T("msg.lang_select_failed"), err)
			}

			newLang := languages[index].Code
			if err := updateEnvFile("MGIT_LANG", newLang); err != nil {
				log.Fatal(i18n.T("msg.lang_update_failed"), err)
			}
			i18n.SetLanguage(newLang)
			fmt.Printf(i18n.T("msg.lang_updated"), newLang)
			return
		}

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

		default:
			fmt.Println(i18n.T("error.invalid_setting"))
		}

	case "update":
		if err := checkAndUpdateSelf(database, machineID); err != nil {
			log.Fatal(i18n.T("msg.app_update_failed"), err)
		}
		return

	default:
		fmt.Println(i18n.T("error.invalid_command"))
		showHelp()
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

// 修改 getCurrentDir 函数
func getCurrentDir(isAppSelf bool) string {
	if isAppSelf {
		// 如果是应用自身，使用 MGIT_HOME
		mgitHome := os.Getenv("MGIT_HOME")
		if mgitHome != "" {
			return mgitHome
		}
		log.Fatal(i18n.T("msg.mgit_home_not_set"))
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
	// 获取应用安装路径
	mgitHome := os.Getenv("MGIT_HOME")
	if mgitHome == "" {
		return fmt.Errorf(i18n.T("msg.mgit_home_not_set"))
	}

	// 获取应用信息
	app, err := db.GetProject(database, "mgit")
	if err != nil {
		return fmt.Errorf(i18n.T("msg.app_info_get_failed"), err)
	}

	// 使用 MGIT_HOME 路径下的可执行文件
	execPath := filepath.Join(mgitHome, "mgit")
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

	if err := os.Chdir(mgitHome); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	// 拉取最新代码
	if err := git.GitPull(app.LocalBranch, app.RemoteBranch); err != nil {
		return fmt.Errorf(i18n.T("error.pull_latest"), err)
	}

	// 更新最后提交信息
	if err := db.UpdateLastPush(database, "mgit", machineID); err != nil {
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
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf(i18n.T("msg.dir_get_failed"), err)
	}

	workDir := getProjectWorkDir(project)
	if err := os.Chdir(workDir); err != nil {
		return fmt.Errorf(i18n.T("msg.dir_change_failed"), err)
	}

	err = action()

	if chErr := os.Chdir(originalDir); chErr != nil {
		return fmt.Errorf(i18n.T("msg.dir_back_failed"), chErr)
	}

	return err
}

// 修改 getProjectWorkDir 函数
func getProjectWorkDir(project *db.Project) string {
	if project.IsAppSelf {
		// 如果是应用自身，直接返回 MGIT_HOME
		if mgitHome := os.Getenv("MGIT_HOME"); mgitHome != "" {
			return filepath.Clean(mgitHome)
		}
		log.Fatal(i18n.T("msg.mgit_home_not_set"))
	}
	// 其他项目使用 APP_PATH/项目名
	return filepath.Join(getCurrentDir(false), project.Name)
}

// 修改 updateEnvFile 函数
func updateEnvFile(key, value string) error {
	// 读取当前 .env 文件内容
	content, err := os.ReadFile(".env")
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
	if err := os.WriteFile(".env", []byte(strings.Join(newLines, "\n")), 0644); err != nil {
		return fmt.Errorf(i18n.T("msg.env_write_failed"), err)
	}

	return nil
}

// 添加显示帮助信息的函数
func showHelp() {
	// 确保已加载语言文件
	if err := godotenv.Load(); err == nil {
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
	
	fmt.Printf("\n  help      - %s\n", i18n.T("cmd.help"))
	fmt.Printf("             %s\n", i18n.T("msg.help_aliases"))
}

// 添加辅助函数用于读取用户输入
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
} 