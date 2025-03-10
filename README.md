# MGit - Git 项目管理工具

[简体中文](#简体中文) | [English](#english) | [日本語](#japanese) | [한국어](#korean) | [繁體中文](#traditional-chinese) | [Français](#français)

GitHub: [https://github.com/liliangshan/mgit](https://github.com/liliangshan/mgit)

## 简体中文

MGit 是一个用于管理多个 Git 项目的命令行工具。它可以帮助您有效地管理、同步和更新多个 Git 仓库。

## 功能特点

- 初始化和管理多个 Git 项目
- 自动同步所有项目的代码
- 跟踪最后提交的机器和时间
- 支持管理工具自身的版本控制
- 灵活配置项目存放路径


## 特性说明

- 多项目管理：一键管理多个 Git 项目，提高工作效率
- 多语言支持：支持中文（简体/繁体）、英文、日文、韩文、法文等多种语言
- 远程数据库同步：支持多设备间的项目配置同步，团队协作更便捷
- 版本控制：内置版本管理，支持自动更新
- 系统集成：支持添加到系统环境变量，随时随地使用
- 智能分支管理：自动检测和切换分支，避免冲突
- 批量操作：支持同时拉取/推送多个项目的更改

### 版本更新说明
1.0.5 
- 更新远程数据库每次强制更新（如果开启远程数据库）
- 单个pull时如果远程有多个分支可以选择要pull的分支，批量pull仍然以数据库的为准备
- 增加切换本地分支和远程分支的功能

### 远程数据库同步

MGit 支持通过远程 Git 仓库同步数据库，实现多设备间的配置共享：

1. 启用数据库同步：
```bash
./mgit set
# 选择 "启用数据库仓库同步"
# 输入数据库仓库地址
```

2. 同步机制：
- 每次 push 操作后自动同步数据库
- 每次 pull 操作前自动获取最新配置
- 支持多人协作，自动合并配置

3. 数据同步内容：
- 项目配置信息
- 分支设置
- 最后提交记录
- 设备标识

### 系统环境变量配置

将 MGit 添加到系统 PATH，实现全局访问：

1. Windows:
```powershell
# 将 MGit 目录添加到用户 PATH
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# 编辑 ~/.bashrc 或 ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

配置后，可在任意目录使用 `mgit` 命令。

### 版本信息

查看当前版本：
```bash
mgit -v
# 或
mgit --version
```

当前版本：1.0.5
项目主页：https://github.com/liliangshan/mgit

### 为什么选择 MGit？

1. 高效协作
- 多设备配置自动同步
- 团队项目统一管理
- 批量操作节省时间

2. 用户友好
- 交互式命令行界面
- 多语言本地化支持
- 智能提示和帮助

3. 安全可靠
- 自动备份配置
- 版本控制保护
- 冲突智能处理

4. 扩展性强
- 支持自定义配置
- 灵活的项目管理
- 持续更新迭代

## 安装

1. 克隆仓库：
```bash
git clone [repository_url]
```

2. 编译：
```bash
go build -o mgit
```

3. 初始化工具：
```bash
# 初始化工具自身
./mgit init
# 或者
./mgit init mgit
```

4. 配置环境：
```bash
# 设置机器标识
./mgit set machine your-machine-name

# 设置项目存放路径
./mgit set path /your/custom/path

# 查看当前配置
./mgit set
```

## 配置说明

工具使用 .env 文件存储配置：

```env
# 机器标识
MACHINE_ID=machine-01

# 应用路径（存放所有项目的父目录）
APP_PATH=/path/to/projects

# 应用安装路径（自动设置，不要手动修改）
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID`: 用于标识当前机器，可通过 `set machine` 命令修改
- `APP_PATH`: 指定项目存放的根目录，可通过 `set path` 命令修改
- `MGIT_HOME`: 工具自身的安装路径（自动管理）

## 使用方法

### 初始化项目

```bash
# 初始化新项目
./mgit init project_name https://github.com/user/repo.git

# 初始化或更新工具自身
./mgit init mgit
# 或者直接
./mgit init
```

### 拉取代码

```bash
# 拉取单个项目
./mgit pull project_name
# 或者
./mgit pull
# 然后在交互式菜单中选择要同步的项目或者选择要同步的项目

# 拉取所有项目
./mgit pull-all

```

项目选择菜单会显示：
- 项目名称和仓库地址
- 最后提交的机器和时间
- "== 全部项目 ==" 选项（用于同步所有项目）

### 推送代码

```bash
# 推送单个项目
./mgit push project_name
# 或者
./mgit push
# 然后在交互式菜单中选择要推送的项目
# 之后可以输入提交说明（留空则使用机器名）

# 推送所有项目
./mgit push-all
# 使用相同的提交说明推送所有项目
```

提交说明示例：
```
请输入提交说明: 更新配置文件
# 最终提交信息：更新配置文件 (由 machine-01 推送)

请输入提交说明: 
# 最终提交信息：由 machine-01 推送
```

交互式菜单示例：
```
选择项目:
  ▸ == 全部项目 == (同步/推送所有项目)
    project1 (https://github.com/user/project1.git) [最后提交: machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### 查看项目列表

```bash
./mgit list
```

### 删除项目

```bash
./mgit delete
# 通过交互式菜单选择要删除的项目
```

### 查看帮助

```bash
./mgit help
# 或者使用以下别名
./mgit h
./mgit -h
./mgit -help
```

### 设置项目分支

```bash
# 交互式设置项目分支
./mgit branch

# 直接设置指定项目的本地和远程分支
./mgit branch project_name local_branch remote_branch
```

交互式分支设置流程：
1. 选择要设置分支的项目
2. 选择远程分支
3. 输入本地分支名称

直接设置示例：
```bash
# 将 project1 的本地分支设置为 develop，远程分支设置为 origin/main
./mgit branch project1 develop origin/main
```

分支设置可以帮助您：
- 快速切换和管理项目分支
- 同步本地和远程分支配置
- 在多项目间统一分支管理

### 设置拉取分支

```bash
# 交互式设置项目拉取分支
./mgit set pull-branch

# 为特定项目直接设置拉取分支
./mgit set pull-branch project_name branch_name
```

拉取分支设置的好处：
- 精确控制每个项目的拉取分支
- 简化多分支项目的代码同步
- 避免意外拉取错误的分支

### 帮助命令

```bash
# 显示完整帮助信息
./mgit help
# 或使用以下别名
./mgit h
./mgit -h
./mgit -help
```

帮助命令提供：
- 所有可用命令列表
- 每个命令的使用示例
- 详细的操作指南和最佳实践

## English

MGit is a command-line tool for managing multiple Git projects. It helps you efficiently manage, synchronize, and update multiple Git repositories.

## Features

- Initialize and manage multiple Git projects
- Automatically synchronize code for all projects
- Track the machine and time of the last commit
- Support version control for the tool itself
- Flexibly configure project storage paths

## Feature Description

- Multi-project management: Manage multiple Git projects with one click, improving work efficiency
- Multi-language support: Supports Chinese (Simplified/Traditional), English, Japanese, Korean, French, and other languages
- Remote database synchronization: Supports project configuration synchronization between devices, facilitating team collaboration
- Version control: Built-in version management, supports automatic updates
- System integration: Can be added to system environment variables, usable anytime and anywhere
- Smart branch management: Automatically detect and switch branches, avoiding conflicts
- Batch operations: Support simultaneous pulling/pushing of multiple projects

### Version Update Notes
1.0.5 
- Force update remote database each time (if remote database is enabled)
- When pulling a single project, choose which branch to pull if multiple remote branches exist. Batch pull still follows database settings
- Added functionality to switch local and remote branches

### Remote Database Synchronization

MGit supports database synchronization through remote Git repositories, achieving configuration sharing between multiple devices:

1. Enable database synchronization:
```bash
./mgit set
# Select "Enable database repository synchronization"
# Enter database repository address
```

2. Synchronization mechanism:
- Automatically synchronize database after each push operation
- Automatically fetch the latest configuration before each pull operation
- Supports multi-user collaboration, automatic configuration merging

3. Synchronized data:
- Project configuration information
- Branch settings
- Last commit records
- Device identifiers

### System Environment Variable Configuration

Add MGit to system PATH for global access:

1. Windows:
```powershell
# Add MGit directory to user PATH
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# Edit ~/.bashrc or ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

After configuration, you can use the `mgit` command from any directory.

### Version Information

View current version:
```bash
mgit -v
# or
mgit --version
```

Current version: 1.0.5
Project homepage: https://github.com/liliangshan/mgit

### Why Choose MGit?

1. Efficient Collaboration
- Multi-device configuration automatic synchronization
- Unified team project management
- Batch operations save time

2. User-Friendly
- Interactive command-line interface
- Multi-language localization support
- Smart prompts and help

3. Secure and Reliable
- Automatic configuration backup
- Version control protection
- Intelligent conflict handling

4. Highly Extensible
- Support for custom configurations
- Flexible project management
- Continuous updates and iterations

## Installation

1. Clone repository:
```bash
git clone [repository_url]
```

2. Compile:
```bash
go build -o mgit
```

3. Initialize tool:
```bash
# Initialize tool itself
./mgit init
# or
./mgit init mgit
```

4. Configure environment:
```bash
# Set machine identifier
./mgit set machine your-machine-name

# Set project storage path
./mgit set path /your/custom/path

# View current configuration
./mgit set
```

## Configuration Description

The tool stores configurations in .env file:

```env
# Machine identifier
MACHINE_ID=machine-01

# Application path (parent directory of all projects)
APP_PATH=/path/to/projects

# Application installation path (automatically set, do not modify manually)
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID`: Used to identify the current machine, can be modified via `set machine` command
- `APP_PATH`: Specify the root directory for project storage, can be modified via `set path` command
- `MGIT_HOME`: Tool's own installation path (automatically managed)

## Usage Methods

### Initialize Project

```bash
# Initialize new project
./mgit init project_name https://github.com/user/repo.git

# Initialize or update tool itself
./mgit init mgit
# or directly
./mgit init
```

### Pull Code

```bash
# Pull single project
./mgit pull project_name
# or
./mgit pull
# Then select project to sync in interactive menu

# Pull all projects
./mgit pull-all
```

Project selection menu will display:
- Project name and repository address
- Last commit machine and time
- "== All Projects ==" option (for synchronizing all projects)

### Push Code

```bash
# Push single project
./mgit push project_name
# or
./mgit push
# Then select project to push in interactive menu
# Then enter commit message (leave blank to use machine name)

# Push all projects
./mgit push-all
# Use the same commit message to push all projects
```

Commit message example:
```
Enter commit message: Update configuration file
# Final commit message: Update configuration file (pushed by machine-01)

Enter commit message: 
# Final commit message: Pushed by machine-01
```

Interactive menu example:
```
Select project:
  ▸ == All Projects == (Sync/Push all projects)
    project1 (https://github.com/user/project1.git) [Last commit: machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### View Project List

```bash
./mgit list
```

### Delete Project

```bash
./mgit delete
# Select project to delete through interactive menu
```

### View Help

```bash
./mgit help
# Or use the following aliases
./mgit h
./mgit -h
./mgit -help
```

### Set Project Branches

```bash
# Interactively set project branches
./mgit branch

# Directly set local and remote branches for a specific project
./mgit branch project_name local_branch remote_branch
```

Interactive branch setting process:
1. Select project to set branches for
2. Choose remote branch
3. Enter local branch name

Direct setting example:
```bash
# Set local branch of project1 to develop, remote branch to origin/main
./mgit branch project1 develop origin/main
```

Branch setting helps you:
- Quickly switch and manage project branches
- Synchronize local and remote branch configurations
- Unify branch management across multiple projects

### Setting Pull Branch

```bash
# Interactively set project pull branch
./mgit set pull-branch

# Directly set pull branch for a specific project
./mgit set pull-branch project_name branch_name
```

Benefits of Pull Branch Setting:
- Precisely control pull branch for each project
- Simplify code synchronization for multi-branch projects
- Avoid accidentally pulling from the wrong branch

### Help Command

```bash
# Display full help information
./mgit help
# Or use these aliases
./mgit h
./mgit -h
./mgit -help
```

Help command provides:
- List of all available commands
- Usage examples for each command
- Detailed operation guide and best practices

## 繁體中文

MGit 是一個用於管理多個 Git 專案的命令列工具。它可以幫助您有效地管理、同步和更新多個 Git 倉庫。

## 功能特點

- 初始化和管理多個 Git 專案
- 自動同步所有專案的程式碼
- 追蹤最後提交的機器和時間
- 支援管理工具本身的版本控制
- 靈活配置專案存放路徑

## 特性說明

- 多專案管理：一鍵管理多個 Git 專案，提高工作效率
- 多語言支援：支援中文（簡體/繁體）、英文、日文、韓文、法文等多種語言
- 遠端資料庫同步：支援多設備間的專案配置同步，團隊協作更便捷
- 版本控制：內建版本管理，支援自動更新
- 系統整合：支援添加到系統環境變數，隨時隨地使用
- 智慧分支管理：自動檢測和切換分支，避免衝突
- 批次操作：支援同時拉取/推送多個專案的變更

### 版本更新說明
1.0.5 
- 更新遠端資料庫每次強制更新（如果開啟遠端資料庫）
- 單個 pull 時如果遠端有多個分支可以選擇要 pull 的分支，批次 pull 仍然以資料庫的為準
- 增加切換本地分支和遠端分支的功能

### 遠端資料庫同步

MGit 支援通過遠端 Git 倉庫同步資料庫，實現多設備間的配置共享：

1. 啟用資料庫同步：
```bash
./mgit set
# 選擇「啟用資料庫倉庫同步」
# 輸入資料庫倉庫地址
```

2. 同步機制：
- 每次 push 操作後自動同步資料庫
- 每次 pull 操作前自動獲取最新配置
- 支援多人協作，自動合併配置

3. 資料同步內容：
- 專案配置資訊
- 分支設定
- 最後提交記錄
- 設備識別

### 系統環境變數配置

將 MGit 添加到系統 PATH，實現全域存取：

1. Windows:
```powershell
# 將 MGit 目錄添加到使用者 PATH
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# 編輯 ~/.bashrc 或 ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

配置後，可在任意目錄使用 `mgit` 命令。

### 版本資訊

查看當前版本：
```bash
mgit -v
# 或
mgit --version
```

當前版本：1.0.5
專案主頁：https://github.com/liliangshan/mgit

### 為什麼選擇 MGit？

1. 高效協作
- 多設備配置自動同步
- 團隊專案統一管理
- 批次操作節省時間

2. 使用者友好
- 互動式命令列介面
- 多語言本地化支援
- 智慧提示和幫助

3. 安全可靠
- 自動備份配置
- 版本控制保護
- 衝突智慧處理

4. 擴展性強
- 支援自定義配置
- 靈活的專案管理
- 持續更新迭代

## 安裝

1. 克隆倉庫：
```bash
git clone [repository_url]
```

2. 編譯：
```bash
go build -o mgit
```

3. 初始化工具：
```bash
# 初始化工具本身
./mgit init
# 或者
./mgit init mgit
```

4. 配置環境：
```bash
# 設定機器識別碼
./mgit set machine your-machine-name

# 設定專案存放路徑
./mgit set path /your/custom/path

# 查看當前配置
./mgit set
```

## 配置說明

工具使用 .env 檔案儲存配置：

```env
# 機器識別碼
MACHINE_ID=machine-01

# 應用程式路徑（存放所有專案的父目錄）
APP_PATH=/path/to/projects

# 應用程式安裝路徑（自動設定，請勿手動修改）
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID`: 用於識別當前機器，可通過 `set machine` 命令修改
- `APP_PATH`: 指定專案存放的根目錄，可通過 `set path` 命令修改
- `MGIT_HOME`: 工具本身的安裝路徑（自動管理）

## 使用方法

### 初始化專案

```bash
# 初始化新專案
./mgit init project_name https://github.com/user/repo.git

# 初始化或更新工具本身
./mgit init mgit
# 或者直接
./mgit init
```

### 拉取程式碼

```bash
# 拉取單個專案
./mgit pull project_name
# 或者
./mgit pull
# 然後在互動式選單中選擇要同步的專案

# 拉取所有專案
./mgit pull-all
```

專案選擇選單會顯示：
- 專案名稱和倉庫地址
- 最後提交的機器和時間
- "== 全部專案 ==" 選項（用於同步所有專案）

### 推送程式碼

```bash
# 推送單個專案
./mgit push project_name
# 或者
./mgit push
# 然後在互動式選單中選擇要推送的專案
# 之後可以輸入提交說明（留空則使用機器名稱）

# 推送所有專案
./mgit push-all
# 使用相同的提交說明推送所有專案
```

提交說明範例：
```
請輸入提交說明：更新配置檔案
# 最終提交訊息：更新配置檔案（由 machine-01 推送）

請輸入提交說明：
# 最終提交訊息：由 machine-01 推送
```

互動式選單範例：
```
選擇專案：
  ▸ == 全部專案 == (同步/推送所有專案)
    project1 (https://github.com/user/project1.git) [最後提交：machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### 查看專案列表

```bash
./mgit list
```

### 刪除專案

```bash
./mgit delete
# 通過互動式選單選擇要刪除的專案
```

### 查看說明

```bash
./mgit help
# 或者使用以下別名
./mgit h
./mgit -h
./mgit -help
```

### 設定專案分支

```bash
# 互動式設定專案分支
./mgit branch

# 直接設定指定專案的本地和遠端分支
./mgit branch project_name local_branch remote_branch
```

互動式分支設定流程：
1. 選擇要設定分支的專案
2. 選擇遠端分支
3. 輸入本地分支名稱

直接設定範例：
```bash
# 將 project1 的本地分支設定為 develop，遠端分支設定為 origin/main
./mgit branch project1 develop origin/main
```

分支設定可以幫助您：
- 快速切換和管理專案分支
- 同步本地和遠端分支配置
- 在多專案間統一分支管理

### 設定拉取分支

```bash
# 互動式設定專案拉取分支
./mgit set pull-branch

# 為特定專案直接設定拉取分支
./mgit set pull-branch project_name branch_name
```

拉取分支設定的好處：
- 精確控制每個專案的拉取分支
- 簡化多分支專案的程式碼同步
- 避免意外拉取錯誤的分支

### 說明命令

```bash
# 顯示完整說明資訊
./mgit help
# 或使用以下別名
./mgit h
./mgit -h
./mgit -help
```

說明命令提供：
- 所有可用命令列表
- 每個命令的使用範例
- 詳細的操作指南和最佳實踐

## Japanese

MGit は複数の Git プロジェクトを管理するためのコマンドラインツールです。複数の Git リポジトリを効率的に管理、同期、更新することができます。

## 機能の特徴

- 複数の Git プロジェクトの初期化と管理
- すべてのプロジェクトのコードを自動同期
- 最終コミットのマシンと時刻を追跡
- ツール自体のバージョン管理をサポート
- プロジェクトの保存パスを柔軟に設定可能

## 機能説明

- マルチプロジェクト管理：複数の Git プロジェクトをワンクリックで管理し、作業効率を向上
- 多言語サポート：中国語（簡体字/繁体字）、英語、日本語、韓国語、フランス語など多言語対応
- リモートデータベース同期：デバイス間のプロジェクト設定同期をサポート、チームコラボレーションを促進
- バージョン管理：内蔵のバージョン管理、自動更新をサポート
- システム統合：システム環境変数に追加可能、いつでもどこでも使用可能
- スマートブランチ管理：ブランチの自動検出と切り替え、コンフリクトを回避
- バッチ操作：複数のプロジェクトの同時プル/プッシュをサポート

### バージョン更新情報
1.0.5 
- リモートデータベースの強制更新（リモートデータベースが有効な場合）
- 単一プルの場合、リモートに複数のブランチがある場合はプルするブランチを選択可能
- ローカルブランチとリモートブランチの切り替え機能を追加

### リモートデータベース同期

MGit はリモート Git リポジトリを通じてデータベースを同期し、複数のデバイス間で設定を共有できます：

1. データベース同期を有効にする：
```bash
./mgit set
# "データベースリポジトリ同期を有効にする"を選択
# データベースリポジトリアドレスを入力
```

2. 同期メカニズム：
- プッシュ操作後に自動的にデータベースを同期
- プル操作前に自動的に最新の設定を取得
- マルチユーザーコラボレーション、自動設定マージをサポート

3. 同期データ：
- プロジェクト設定情報
- ブランチ設定
- 最終コミット記録
- デバイス識別子

### システム環境変数の設定

MGitをシステム PATHに追加してグローバルアクセスを実現：

1. Windows:
```powershell
# MGit ディレクトリをユーザー PATH に追加
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# ~/.bashrc または ~/.zshrc を編集
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

設定後、任意のディレクトリで `mgit` コマンドを使用できます。

### バージョン情報

現在のバージョンを確認：
```bash
mgit -v
# または
mgit --version
```

現在のバージョン：1.0.5
プロジェクトホームページ：https://github.com/liliangshan/mgit

### MGitを選ぶのか？

1. 効率的なコラボレーション
- 複数デバイス間の設定自動同期
- チームプロジェクトの統一管理
- バッチ操作で時間を節約

2. ユーザーフレンドリー
- インタラクティブなコマンドラインインターフェース
- 多言語ローカライゼーションサポート
- スマートなプロンプトとヘルプ

3. 安全で信頼性が高い
- 自動設定バックアップ
- バージョン管理による保護
- インテリジェントな競合処理

4. 高い拡張性
- カスタム設定のサポート
- 柔軟なプロジェクト管理
- 継続的な更新と改良

## インストール

1. リポジトリをクローン：
```bash
git clone [repository_url]
```

2. コンパイル：
```bash
go build -o mgit
```

3. ツールの初期化：
```bash
# ツール自体を初期化
./mgit init
# または
./mgit init mgit
```

4. 環境の設定：
```bash
# マシン識別子を設定
./mgit set machine your-machine-name

# プロジェクト保存パスを設定
./mgit set path /your/custom/path

# 現在の設定を表示
./mgit set
```

## 設定の説明

ツールは .env ファイルに設定を保存します：

```env
# マシン識別子
MACHINE_ID=machine-01

# アプリケーションパス（すべてのプロジェクトの親ディレクトリ）
APP_PATH=/path/to/projects

# アプリケーションインストールパス（自動設定、手動で変更しないでください）
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID`: 現在のマシンを識別するために使用、`set machine` コマンドで変更可能
- `APP_PATH`: プロジェクト保存のルートディレクトリを指定、`set path` コマンドで変更可能
- `MGIT_HOME`: ツール自体のインストールパス（自動管理）

## 使用方法

### プロジェクトの初期化

```bash
# 新しいプロジェクトを初期化
./mgit init project_name https://github.com/user/repo.git

# ツール自体を初期化または更新
./mgit init mgit
# または直接
./mgit init
```

### コードのプル

```bash
# 単一プロジェクトをプル
./mgit pull project_name
# または
./mgit pull
# その後、インタラクティブメニューで同期するプロジェクトを選択

# すべてのプロジェクトをプル
./mgit pull-all
```

プロジェクト選択メニューには以下が表示されます：
- プロジェクト名とリポジトリアドレス
- 最後のコミットマシンと時間
- "== すべてのプロジェクト ==" オプション（すべてのプロジェクトを同期するため）

### コードのプッシュ

```bash
# 単一プロジェクトをプッシュ
./mgit push project_name
# または
./mgit push
# その後、インタラクティブメニューでプッシュするプロジェクトを選択
# その後、コミットメッセージを入力（空白の場合はマシン名を使用）

# すべてのプロジェクトをプッシュ
./mgit push-all
# 同じコミットメッセージですべてのプロジェクトをプッシュ
```

コミットメッセージの例：
```
コミットメッセージを入力：設定ファイルを更新
# 最終コミットメッセージ：設定ファイルを更新（machine-01 によってプッシュ）

コミットメッセージを入力：
# 最終コミットメッセージ：machine-01 によってプッシュ
```

インタラクティブメニューの例：
```
プロジェクトを選択：
  ▸ == すべてのプロジェクト == （すべてのプロジェクトを同期/プッシュ）
    project1 (https://github.com/user/project1.git) [最終コミット：machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### プロジェクトリストの表示

```bash
./mgit list
```

### プロジェクトの削除

```bash
./mgit delete
# インタラクティブメニューで削除するプロジェクトを選択
```

### ヘルプの表示

```bash
./mgit help
# または以下のエイリアスを使用
./mgit h
./mgit -h
./mgit -help
```

### プロジェクトブランチの設定

```bash
# インタラクティブにプロジェクトブランチを設定
./mgit branch

# 特定のプロジェクトのローカルブランチとリモートブランチを直接設定
./mgit branch project_name local_branch remote_branch
```

インタラクティブなブランチ設定プロセス：
1. ブランチを設定するプロジェクトを選択
2. リモートブランチを選択
3. ローカルブランチ名を入力

直接設定の例：
```bash
# project1のローカルブランチを develop に、リモートブランチを origin/main に設定
./mgit branch project1 develop origin/main
```

ブランチ設定は以下に役立ちます：
- プロジェクトブランチの迅速な切り替えと管理
- ローカルとリモートのブランチ設定の同期
- 複数プロジェクト間でのブランチ管理の統一

### プル元ブランチの設定

```bash
# 対話形式でプロジェクトのプル元ブランチを設定
./mgit set pull-branch

# 特定のプロジェクトのプル元ブランチを直接設定
./mgit set pull-branch project_name branch_name
```

プル元ブランチ設定のメリット：
- 各プロジェクトのプル元ブランチを正確に制御
- マルチブランチプロジェクトのコード同期を簡素化
- 誤ったブランチからのプルを防止

### ヘルプコマンド

```bash
# 完全なヘルプ情報を表示
./mgit help
# または以下のエイリアスを使用
./mgit h
./mgit -h
./mgit -help
```

ヘルプコマンドが提供する情報：
- 利用可能なすべてのコマンドのリスト
- 各コマンドの使用例
- 詳細な操作ガイドとベストプラクティス


## Korean

MGit은 여러 Git 프로젝트를 관리하기 위한 명령줄 도구입니다. 여러 Git 저장소를 효율적으로 관리, 동기화 및 업데이트할 수 있습니다.

## 기능 특징

- 여러 Git 프로젝트 초기화 및 관리
- 모든 프로젝트 코드 자동 동기화
- 마지막 커밋 머신 및 시간 추적
- 도구 자체 버전 관리 지원
- 프로젝트 저장 경로 유연한 설정

## 기능 설명

- 다중 프로젝트 관리: 여러 Git 프로젝트를 원클릭으로 관리하여 작업 효율성 향상
- 다국어 지원: 중국어(간체/번체), 영어, 일본어, 한국어, 프랑스어 등 다양한 언어 지원
- 원격 데이터베이스 동기화: 여러 장치 간 프로젝트 설정 동기화 지원, 팀 협업 편의성 향상
- 버전 관리: 내장 버전 관리, 자동 업데이트 지원
- 시스템 통합: 시스템 환경 변수에 추가 가능, 언제 어디서나 사용 가능
- 스마트 브랜치 관리: 브랜치 자동 감지 및 전환, 충돌 방지
- 일괄 작업: 여러 프로젝트 동시 풀/푸시 지원

### 버전 업데이트 정보
1.0.5 
- 원격 데이터베이스 강제 업데이트(원격 데이터베이스 활성화 시)
- 단일 풀 시 원격에 여러 브랜치가 있는 경우 풀할 브랜치 선택 가능
- 로컬 브랜치와 원격 브랜치 전환 기능 추가

### 리모트 데이터베이스 동기화

MGit은 원격 Git 저장소를 통해 데이터베이스를 동기화하여 여러 장치 간 설정을 공유합니다:

1. 데이터베이스 동기화 활성화:
```bash
./mgit set
# "데이터베이스 저장소 동기화 활성화" 선택
# 데이터베이스 저장소 주소 입력
```

2. 동기화 메커니즘:
- 푸시 작업 후 자동 데이터베이스 동기화
- 풀 작업 전 자동 최신 설정 가져오기
- 다중 사용자 협업, 자동 설정 병합 지원

3. 동기화 데이터:
- 프로젝트 설정 정보
- 브랜치 설정
- 마지막 커밋 기록
- 장치 식별자

### 시스템 환경 변수 설정

MGit을 시스템 PATH에 추가하여 전역 액세스 실현:

1. Windows:
```powershell
# MGit 디렉토리를 사용자 PATH에 추가
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# ~/.bashrc 또는 ~/.zshrc 편집
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

설정 후 어느 디렉토리에서나 `mgit` 명령어를 사용할 수 있습니다.

### 버전 정보

현재 버전 확인:
```bash
mgit -v
# 또는
mgit --version
```

현재 버전: 1.0.5
프로젝트 홈페이지: https://github.com/liliangshan/mgit

### MGit을 선택하는 이유?

1. 효율적인 협업
- 다중 장치 구성 자동 동기화
- 팀 프로젝트 통합 관리
- 일괄 작업으로 시간 절약

2. 사용자 친화적
- 대화형 명령줄 인터페이스
- 다국어 현지화 지원
- 스마트 프롬프트 및 도움말

3. 안전하고 신뢰할 수 있음
- 자동 구성 백업
- 버전 관리 보호
- 지능적인 충돌 처리

4. 높은 확장성
- 사용자 정의 구성 지원
- 유연한 프로젝트 관리
- 지속적인 업데이트 및 반복

## 설치

1. 저장소 복제:
```bash
git clone [repository_url]
```

2. 컴파일:
```bash
go build -o mgit
```

3. 도구 초기화:
```bash
# 도구 자체 초기화
./mgit init
# 또는
./mgit init mgit
```

4. 환경 구성:
```bash
# 머신 식별자 설정
./mgit set machine your-machine-name

# 프로젝트 저장 경로 설정
./mgit set path /your/custom/path

# 현재 구성 보기
./mgit set
```

## 구성 설명

도구는 .env 파일에 구성을 저장합니다:

```env
# 머신 식별자
MACHINE_ID=machine-01

# 애플리케이션 경로(모든 프로젝트의 상위 디렉토리)
APP_PATH=/path/to/projects

# 애플리케이션 설치 경로(자동 설정, 수동으로 수정하지 마세요)
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID`: 현재 머신을 식별하는 데 사용, `set machine` 명령으로 수정 가능
- `APP_PATH`: 프로젝트 저장을 위한 루트 디렉토리 지정, `set path` 명령으로 수정 가능
- `MGIT_HOME`: 도구 자체의 설치 경로(자동 관리)

## 사용 방법

### 프로젝트 초기화

```bash
# 새 프로젝트 초기화
./mgit init project_name https://github.com/user/repo.git

# 도구 자체 초기화 또는 업데이트
./mgit init mgit
# 또는 직접
./mgit init
```

### 코드 풀

```bash
# 단일 프로젝트 풀
./mgit pull project_name
# 또는
./mgit pull
# 그런 다음 대화형 메뉴에서 동기화할 프로젝트 선택

# 모든 프로젝트 풀
./mgit pull-all
```

프로젝트 선택 메뉴에 표시됩니다:
- 프로젝트 이름 및 저장소 주소
- 마지막 커밋 머신 및 시간
- "== 모든 프로젝트 ==" 옵션(모든 프로젝트 동기화용)

### 코드 푸시

```bash
# 단일 프로젝트 푸시
./mgit push project_name
# 또는
./mgit push
# 그런 다음 대화형 메뉴에서 푸시할 프로젝트 선택
# 그런 다음 커밋 메시지 입력(비워두면 머신 이름 사용)

# 모든 프로젝트 푸시
./mgit push-all
# 동일한 커밋 메시지로 모든 프로젝트 푸시
```

커밋 메시지 예:
```
커밋 메시지 입력: 구성 파일 업데이트
# 최종 커밋 메시지: 구성 파일 업데이트(machine-01에 의해 푸시됨)

커밋 메시지 입력:
# 최종 커밋 메시지: machine-01에 의해 푸시됨
```

대화형 메뉴 예:
```
프로젝트 선택:
  ▸ == 모든 프로젝트 == (모든 프로젝트를 동기화/푸시)
    project1 (https://github.com/user/project1.git) [마지막 커밋: machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### 프로젝트 목록 보기

```bash
./mgit list
```

### 프로젝트 삭제

```bash
./mgit delete
# 대화형 메뉴를 통해 삭제할 프로젝트 선택
```

### 도움말 보기

```bash
./mgit help
# 또는 다음 별칭 사용
./mgit h
./mgit -h
./mgit -help
```

### 프로젝트 브랜치 설정

```bash
# 대화형으로 프로젝트 브랜치 설정
./mgit branch

# 특정 프로젝트의 로컬 및 원격 브랜치 직접 설정
./mgit branch project_name local_branch remote_branch
```

대화형 브랜치 설정 과정:
1. 브랜치를 설정할 프로젝트 선택
2. 원격 브랜치 선택
3. 로컬 브랜치 이름 입력

직접 설정 예:
```bash
# project1의 로컬 브랜치를 develop으로, 원격 브랜치를 origin/main으로 설정
./mgit branch project1 develop origin/main
```

브랜치 설정은 다음과 같은 도움을 줍니다:
- 프로젝트 브랜치의 빠른 전환 및 관리
- 로컬 및 원격 브랜치 구성 동기화
- 여러 프로젝트 간 브랜치 관리 통합

### 풀 브랜치 설정

```bash
# 대화형으로 프로젝트의 풀 브랜치 설정
./mgit set pull-branch

# 특정 프로젝트의 풀 브랜치를 직접 설정
./mgit set pull-branch 프로젝트_이름 브랜치_이름
```

풀 브랜치 설정의 이점:
- 각 프로젝트의 풀 브랜치를 정확하게 제어
- 다중 브랜치 프로젝트의 코드 동기화 간소화
- 잘못된 브랜치에서 풀하는 것을 방지


















## Français

### Informations de version

Voir la version actuelle :
```bash
mgit -v
# ou
mgit --version
```

Version actuelle : 1.0.5
Page d'accueil du projet : https://github.com/liliangshan/mgit

### Pourquoi choisir MGit ?

1. Collaboration efficace
- Synchronisation automatique de la configuration multi-appareils
- Gestion unifiée des projets d'équipe
- Les opérations par lots font gagner du temps

2. Convivial
- Interface de ligne de commande interactive
- Support de localisation multilingue
- Invites et aide intelligentes

3. Sûr et fiable
- Sauvegarde automatique de la configuration
- Protection par contrôle de version
- Gestion intelligente des conflits

4. Hautement extensible
- Support pour les configurations personnalisées
- Gestion flexible des projets
- Mises à jour et itérations continues

## Installation

1. Cloner le dépôt :
```bash
git clone [repository_url]
```

2. Compiler :
```bash
go build -o mgit
```

3. Initialiser l'outil :
```bash
# Initialiser l'outil lui-même
./mgit init
# ou
./mgit init mgit
```

4. Configurer l'environnement :
```bash
# Définir l'identifiant de la machine
./mgit set machine your-machine-name

# Définir le chemin de stockage du projet
./mgit set path /your/custom/path

# Voir la configuration actuelle
./mgit set
```

## Description de la configuration

L'outil stocke les configurations dans un fichier .env :

```env
# Identifiant de la machine
MACHINE_ID=machine-01

# Chemin de l'application (répertoire parent de tous les projets)
APP_PATH=/path/to/projects

# Chemin d'installation de l'application (défini automatiquement, ne pas modifier manuellement)
MGIT_HOME=/path/to/mgit
```

- `MACHINE_ID` : Utilisé pour identifier la machine actuelle, peut être modifié via la commande `set machine`
- `APP_PATH` : Spécifie le répertoire racine pour le stockage du projet, peut être modifié via la commande `set path`
- `MGIT_HOME` : Chemin d'installation de l'outil (géré automatiquement)

## Méthodes d'utilisation

### Initialiser un projet

```bash
# Initialiser un nouveau projet
./mgit init project_name https://github.com/user/repo.git

# Initialiser ou mettre à jour l'outil lui-même
./mgit init mgit
# ou directement
./mgit init
```

### Tirer (pull) du code

```bash
# Tirer un seul projet
./mgit pull project_name
# ou
./mgit pull
# Puis sélectionner le projet à synchroniser dans le menu interactif

# Tirer tous les projets
./mgit pull-all
```

Le menu de sélection de projet affichera :
- Nom du projet et adresse du dépôt
- Machine et heure du dernier commit
- Option "== Tous les projets ==" (pour synchroniser tous les projets)

### Pousser (push) du code

```bash
# Pousser un seul projet
./mgit push project_name
# ou
./mgit push
# Puis sélectionner le projet à pousser dans le menu interactif
# Puis entrer un message de commit (laisser vide pour utiliser le nom de la machine)

# Pousser tous les projets
./mgit push-all
# Utiliser le même message de commit pour pousser tous les projets
```

Exemple de message de commit :
```
Entrez le message de commit : Mise à jour du fichier de configuration
# Message de commit final : Mise à jour du fichier de configuration (poussé par machine-01)

Entrez le message de commit : 
# Message de commit final : Poussé par machine-01
```

Exemple de menu interactif :
```
Sélectionnez un projet :
  ▸ == Tous les projets == (Synchroniser/Pousser tous les projets)
    project1 (https://github.com/user/project1.git) [Dernier commit : machine-02 @ 2024-01-02 13:00:00]
    project2 (https://github.com/user/project2.git)
```

### Voir la liste des projets

```bash
./mgit list
```

### Supprimer un projet

```bash
./mgit delete
# Sélectionner le projet à supprimer via le menu interactif
```

### Voir l'aide

```bash
./mgit help
# Ou utiliser les alias suivants
./mgit h
./mgit -h
./mgit -help
```

### Définir les branches du projet

```bash
# Définir interactivement les branches du projet
./mgit branch

# Définir directement les branches locales et distantes pour un projet spécifique
./mgit branch project_name local_branch remote_branch
```

Processus de configuration interactive des branches :
1. Sélectionner le projet pour lequel définir les branches
2. Choisir la branche distante
3. Entrer le nom de la branche locale

Exemple de configuration directe :
```bash
# Définir la branche locale de project1 à develop, la branche distante à origin/main
./mgit branch project1 develop origin/main
```

La configuration des branches vous aide à :
- Basculer et gérer rapidement les branches du projet
- Synchroniser les configurations de branches locales et distantes
- Unifier la gestion des branches entre plusieurs projets

### Configuration de la branche de pull

```bash
# Configuration interactive de la branche de pull du projet
./mgit set pull-branch

# Définir directement la branche de pull pour un projet spécifique
./mgit set pull-branch nom_projet nom_branche
```

Avantages de la configuration de la branche de pull :
- Contrôle précis de la branche de pull pour chaque projet
- Simplification de la synchronisation du code pour les projets multi-branches
- Éviter de pull accidentellement à partir de la mauvaise branche

