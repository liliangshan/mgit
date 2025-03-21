package git

import (
	"bufio"
	"fmt"
	"gitmanager/i18n"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/manifoldco/promptui"
)

// 添加全局变量
var isGitConnectionClosed bool

// GitOperation 执行 git 命令并实时输出日志
func GitOperation(showMessage bool, command string, args ...string) error {
	// 检查连接状态
	if isGitConnectionClosed {
		return fmt.Errorf(i18n.T("error.git_connection_closed"))
	}

	// 输出完整命令，过滤掉 HEAD
	cmdStr := strings.Join(args, " ")
	cmdStr = strings.TrimSuffix(cmdStr, " HEAD")
	if showMessage {
		fmt.Printf("\n"+i18n.T("msg.executing_command"), command, cmdStr)
	}

	cmd := exec.Command(command, args...)
	var wg sync.WaitGroup
	wg.Add(2)
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {

		os.Exit(1)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		if isGitConnectionClosed {
			return
		}
		GetOutput(readout, "")
	}()

	//捕获标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		if isGitConnectionClosed {
			return
		}
		defer wg.Done()

		GetOutput(readerr, "")
	}()
	//执行命令

	cmd.Run()
	wg.Wait()

	return nil
}

func GetOutput(reader *bufio.Reader, localBranch string) {
	var sumOutput string
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			if strings.Contains(err.Error(), "file already closed") {
				fmt.Println(i18n.T("error.file_closed"))
				isGitConnectionClosed = true // 设置全局标志
				return
			}
			if strings.Contains(err.Error(), "overwritten by merge") {
				fmt.Println(i18n.T("error.pull_merge"))
				sumOutput += i18n.T("error.pull_merge")
			} else {
				fmt.Println(err)
				sumOutput += err.Error()
			}

		}

		output := string(outputBytes[:n])
		if strings.Contains(output, "overwritten by merge") {
			output = i18n.T("error.pull_merge")
		}
		if strings.Contains(output, "file already closed") {
			fmt.Println(i18n.T("error.file_closed"))
			isGitConnectionClosed = true // 设置全局标志
			return
		}
		fmt.Print(output)
		sumOutput += output
	}
	return
}

// GitPull 执行 git pull 操作
func GitPull(localBranch, remoteBranch string, force bool) error {
	if force {
		// 强制拉取，强制覆盖本地分支，使用git fetch --all
		err := GitOperation(true, "git", "fetch", "--all")
		if err != nil {
			return fmt.Errorf(i18n.T("error.fetch_all"), err)
		}
		// 强制拉取，强制覆盖本地分支，使用git reset --hard
		err = GitOperation(true, "git", "reset", "--hard", "origin/"+remoteBranch)
		if err != nil {
			return fmt.Errorf(i18n.T("error.reset_hard"), err)
		}
	}

	// 普通拉取
	err := GitOperation(true, "git", "pull", "origin", remoteBranch)
	if err != nil && strings.Contains(err.Error(), "merge conflict") {
		// 如果遇到合并冲突，建议使用强制拉取
		fmt.Println(i18n.T("msg.merge_conflict_detected"))
		fmt.Println(i18n.T("msg.suggest_force_pull"))
		return fmt.Errorf(i18n.T("error.pull_conflict"))
	}
	return err
}

// 添加检查未提交更改的函数
func hasUncommittedChanges() (bool, error) {
	// 使用 git status --porcelain 检查是否有未提交的更改
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf(i18n.T("error.check_changes"), err)
	}

	// 如果输出不为空，说明有未提交的更改
	return len(strings.TrimSpace(string(output))) > 0, nil
}

// GitPush 执行 git add、commit 和 push 操作
func GitPush(localBranch, remoteBranch, message string) error {
	// 先尝试获取最新代码
	/*if err := GitPull(localBranch, remoteBranch); err != nil {
		return fmt.Errorf(i18n.T("error.pull_latest"), err)
	}*/
	var showMessage bool
	showMessage = true
	if message == "update-mgit-database" {
		isShowMsg := false
		showMessage = isShowMsg
	}
	// 如果有未提交的更改，先提交
	if hasChanges, err := hasUncommittedChanges(); err != nil {
		return err
	} else if hasChanges {

		// 添加所有更改
		if err := GitOperation(showMessage, "git", "add", "."); err != nil {
			return fmt.Errorf(i18n.T("error.add_changes"), err)
		}

		// 提交更改
		if err := GitOperation(showMessage, "git", "commit", "-m", message); err != nil {
			return fmt.Errorf(i18n.T("error.commit_changes"), err)
		}
	}

	// 推送到远程
	if err := GitPustProgress(showMessage, localBranch, remoteBranch); err != nil {
		return fmt.Errorf(i18n.T("error.push_failed"), err)
	}

	return nil
}
func GitPustProgress(showMessage bool, localBranch string, remoteBranch string) error {
	//return GitOperation(showMessage, "git", "push", "origin", fmt.Sprintf("%s:%s", localBranch, remoteBranch), "--progress")
	command := "git"
	args := []string{"push", "origin", fmt.Sprintf("%s:%s", localBranch, remoteBranch), "--progress"}
	cmdStr := strings.Join(args, " ")
	cmdStr = strings.TrimSuffix(cmdStr, " HEAD")
	if showMessage {
		fmt.Printf("\n"+i18n.T("msg.executing_command"), command, cmdStr)
	}

	cmd := exec.Command(command, args...)
	var wg sync.WaitGroup
	wg.Add(2)
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {

		os.Exit(1)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()

		GetOutput(readout, localBranch)
	}()

	//捕获标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		os.Exit(1)
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		GetOutput(readerr, localBranch)
	}()
	//执行命令

	cmd.Run()
	wg.Wait()

	return nil

}

// GitInit 初始化仓库并拉取代码
func GitInit(repoURL string, remoteBranch string) (string, string, error) {
	// 初始化仓库
	if err := GitOperation(true, "git", "init"); err != nil {
		return "", "", fmt.Errorf(i18n.T("error.init_repo"), err)
	}

	// 添加远程仓库
	if err := GitOperation(true, "git", "remote", "add", "origin", repoURL); err != nil {
		return "", "", fmt.Errorf(i18n.T("error.add_remote"), err)
	}

	// 设置代理
	if err := SetProxy(); err != nil {
		return "", "", fmt.Errorf(i18n.T("error.set_proxy"), err)
	}

	// 获取默认分支
	remoteBranches, err := GetDefaultRemoteBranch(repoURL)
	if err != nil {
		return "", "", fmt.Errorf(i18n.T("msg.fetch_branches_failed"), err)
	}
	if len(remoteBranches) == 0 {
		// 如果没有获取到分支，使用默认的 main 分支
		return "", "", fmt.Errorf(i18n.T("msg.fetch_branches_failed"), err)
	}
	if remoteBranch == "" {

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
			return "", "", fmt.Errorf(i18n.T("msg.branch_select_failed"), err)
		}

		remoteBranch = remoteBranches[index]
		//修改本地分支名和选择的一样
		if err := GitOperation(true, "git", "branch", "-m", remoteBranch); err != nil {
			return "", "", fmt.Errorf(i18n.T("error.branch_rename"), err)
		}
	}

	// 拉取默认分支代码并创建本地分支
	if err := GitOperation(true, "git", "pull", "origin", remoteBranch, "--progress"); err != nil {
		return "", "", fmt.Errorf(i18n.T("error.fetch_code"), err)
	}
	var pullBranch string = remoteBranch
	if len(remoteBranches) > 0 {
		prompt := promptui.Select{
			Label: i18n.T("prompt.select_push_branch"),
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
			return "", "", fmt.Errorf(i18n.T("msg.branch_select_failed"), err)
		}

		remoteBranch = remoteBranches[index]
		//修改本地分支名和选择的一样
		if err := GitOperation(true, "git", "branch", "-m", remoteBranch); err != nil {
			return "", "", fmt.Errorf(i18n.T("error.branch_rename"), err)
		}
	}
	return pullBranch, remoteBranch, nil
}

// GetDefaultRemoteBranch 根据仓库地址判断默认分支名称
func GetDefaultRemoteBranch(repoUrl string) ([]string, error) {

	// 修改成利用 git ls-remote --heads 获取远程分支列表
	cmd := exec.Command("git", "ls-remote", "--heads", repoUrl)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	// 解析输出，获取分支列表
	branches := []string{}
	for _, line := range strings.Split(string(output), "\n") {
		branch := strings.TrimSpace(line)
		if branch != "" && strings.Contains(branch, "refs/heads/") {
			//以refs/heads/分割获取后面的分支名
			branchName := strings.Split(branch, "refs/heads/")[1]
			branches = append(branches, branchName)
		}
	}
	return branches, nil

}

// SetProxy 设置Git代理
func SetProxy() error {
	// 让用户选择是否设置代理
	type ProxyOption struct {
		Name  string
		Value bool
	}
	options := []ProxyOption{
		{Name: i18n.T("msg.no"), Value: false},
		{Name: i18n.T("msg.yes"), Value: true},
	}

	prompt := promptui.Select{
		Label: i18n.T("msg.set_proxy_confirm"),
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
		return fmt.Errorf(i18n.T("error.invalid_selection"))
	}

	// 如果选择否
	if !options[index].Value {
		RemoveProxy() // 直接调用，不处理错误
		return nil
	}

	// 选择代理类型
	proxyTypes := []string{"HTTP", "SOCKS5"}
	prompt = promptui.Select{
		Label: i18n.T("msg.select_proxy_type"),
		Items: proxyTypes,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "> {{ . | cyan }}",
			Inactive: "  {{ . }}",
			Selected: "✓ {{ . | green }}",
		},
	}

	index, _, err = prompt.Run()
	if err != nil {
		return fmt.Errorf(i18n.T("error.invalid_proxy_type"))
	}

	var proxyType string
	switch index {
	case 0:
		proxyType = "http"
	case 1:
		proxyType = "socks5"
	default:
		return fmt.Errorf(i18n.T("error.invalid_proxy_type"))
	}

	// 输入代理IP
	fmt.Print(i18n.T("prompt.proxy_ip") + " (127.0.0.1): ")
	var ip string
	fmt.Scanln(&ip)
	if ip == "" {
		ip = "127.0.0.1"
	}

	// 输入代理端口
	fmt.Print(i18n.T("prompt.proxy_port") + ": ")
	var port string
	fmt.Scanln(&port)

	// 检查端口是否为空
	if port == "" {
		return fmt.Errorf(i18n.T("error.empty_port"))
	}

	// 设置代理
	proxyURL := ""
	if proxyType == "http" {
		proxyURL = fmt.Sprintf("http://%s:%s", ip, port)
	} else if proxyType == "socks5" {
		proxyURL = fmt.Sprintf("socks5://%s:%s", ip, port)
	}

	// 设置http代理
	if err := exec.Command("git", "config", "http.proxy", proxyURL).Run(); err != nil {
		return fmt.Errorf(i18n.T("error.set_http_proxy"), err)
	}

	// 设置https代理
	if err := exec.Command("git", "config", "https.proxy", proxyURL).Run(); err != nil {
		return fmt.Errorf(i18n.T("error.set_https_proxy"), err)
	}

	return nil
}

// RemoveProxy 移除Git代理
func RemoveProxy() error {
	// 移除http代理
	if err := exec.Command("git", "config", "--unset", "http.proxy").Run(); err != nil {
		return fmt.Errorf(i18n.T("error.remove_http_proxy"), err)
	}

	// 移除https代理
	if err := exec.Command("git", "config", "--unset", "https.proxy").Run(); err != nil {
		return fmt.Errorf(i18n.T("error.remove_https_proxy"), err)
	}

	return nil
}
