package git

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"gitmanager/i18n"
)

// GitOperation 执行 git 命令并实时输出日志
func GitOperation(command string, args ...string) error {
	// 输出完整命令，过滤掉 HEAD
	cmdStr := strings.Join(args, " ")
	cmdStr = strings.TrimSuffix(cmdStr, " HEAD")
	fmt.Printf(i18n.T("msg.executing_command"), command, cmdStr)
	
	cmd := exec.Command(command, args...)
	
	// 获取命令的输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf(i18n.T("error.create_output_pipe"), err)
	}
	
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf(i18n.T("error.create_error_pipe"), err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		return fmt.Errorf(i18n.T("error.start_command"), err)
	}

	// 实时读取并打印输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf(i18n.T("msg.output"), scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Printf(i18n.T("msg.error"), scanner.Text())
		}
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf(i18n.T("error.command_failed"), err)
	}

	return nil
}

// GitPull 执行 git pull 操作
func GitPull(localBranch, remoteBranch string) error {
	// 清理分支名称，只取第一个部分
	remoteBranch = strings.Fields(remoteBranch)[0]
	
	// 调试输出
	fmt.Printf(i18n.T("msg.remote_branch"), remoteBranch)
	
	return GitOperation("git", "pull", "origin", remoteBranch)
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
	if err := GitPull(localBranch, remoteBranch); err != nil {
		return fmt.Errorf(i18n.T("error.pull_latest"), err)
	}

	// 如果有未提交的更改，先提交
	if hasChanges, err := hasUncommittedChanges(); err != nil {
		return err
	} else if hasChanges {
		// 添加所有更改
		if err := GitOperation("git", "add", "."); err != nil {
			return fmt.Errorf(i18n.T("error.add_changes"), err)
		}

		// 提交更改
		if err := GitOperation("git", "commit", "-m", message); err != nil {
			return fmt.Errorf(i18n.T("error.commit_changes"), err)
		}
	}

	// 推送到远程
	if err := GitOperation("git", "push", "origin", fmt.Sprintf("%s:%s", localBranch, remoteBranch)); err != nil {
		return fmt.Errorf(i18n.T("error.push_failed"), err)
	}

	return nil
}

// GitInit 初始化仓库并拉取代码
func GitInit(repoURL string) error {
	// 初始化仓库
	if err := GitOperation("git", "init"); err != nil {
		return fmt.Errorf(i18n.T("error.init_repo"), err)
	}

	// 添加远程仓库
	if err := GitOperation("git", "remote", "add", "origin", repoURL); err != nil {
		return fmt.Errorf(i18n.T("error.add_remote"), err)
	}

	// 获取默认分支
	defaultBranch, err := GetDefaultRemoteBranch(repoURL)
	if err != nil {
		defaultBranch = "main" // 如果获取失败，使用 main 作为默认值
	}
	
	// 清理分支名称，移除可能的空格
	defaultBranch = strings.TrimSpace(defaultBranch)

	// 拉取默认分支代码并创建本地分支
	if err := GitOperation("git", "fetch", "origin"); err != nil {
		return fmt.Errorf(i18n.T("error.fetch_code"), err)
	}

	// 创建并切换到本地分支，设置上游分支
	if err := GitOperation("git", "checkout", "-b", defaultBranch, "origin/"+defaultBranch); err != nil {
		return fmt.Errorf(i18n.T("error.set_default_branch"), err)
	}

	return nil
}

// GetDefaultRemoteBranch 根据仓库地址判断默认分支名称
func GetDefaultRemoteBranch(repoURL string) (string, error) {
	// 如果是 GitHub 仓库，使用 main
	if strings.Contains(strings.ToLower(repoURL), "github.com") {
		return "main", nil
	}
	
	// 其他 Git 服务默认使用 master
	return "master", nil
} 