# MGit - Git 项目管理工具

[简体中文](#简体中文) | [English](#english) | [日本語](#japanese) | [한국어](#korean) | [繁體中文](#traditional-chinese) | [Français](#français)

GitHub: [https://github.com/liliangshan/mgit](https://github.com/liliangshan/mgit)

## 简体中文

# MGit - Git 项目管理工具

## 简介
MGit 是一个用于管理多个 Git 项目的命令行工具。它可以帮助您有效地管理、同步和更新多个 Git 仓库。

## 主要功能
- 多项目管理：一键管理多个 Git 项目，提高工作效率
- 多语言支持：支持中文（简体/繁体）、英文、日文、韩文、法文等多种语言
- 远程数据库同步：支持多设备间的项目配置同步，团队协作更便捷
- 版本控制：内置版本管理，支持自动更新
- 系统集成：支持添加到系统环境变量，随时随地使用
- 智能分支管理：自动检测和切换分支，避免冲突
- 批量操作：支持同时拉取/推送多个项目的更改

## 特性说明

### 1. 高效协作
- 多设备配置自动同步，确保团队成员使用相同的项目配置
- 团队项目统一管理，避免配置不一致
- 批量操作节省时间，提高工作效率
- 智能分支管理，自动处理分支切换和同步

### 2. 用户友好
- 交互式命令行界面，操作直观简单
- 多语言本地化支持，无语言障碍
- 智能提示和帮助，降低使用门槛
- 统一的命令格式，易于记忆和使用

### 3. 安全可靠
- 自动备份配置，防止意外丢失
- 版本控制保护，支持回滚操作
- 冲突智能处理，确保数据安全
- 权限管理，防止误操作

### 4. 扩展性强
- 支持自定义配置，满足不同需求
- 灵活的项目管理，适应各种场景
- 持续更新迭代，不断优化功能
- 支持插件扩展，功能可扩展

## 版本更新说明
1.0.16
- 新增应用实例管理功能，支持通过 `mgit new` 创建新的应用实例
- 优化配置文件和数据库命名规则，根据应用名自动生成对应文件名
- 新增环境变量自动设置功能，支持将应用目录添加到系统环境变量
- 更新远程数据库每次强制更新（如果开启远程数据库）
- 单个pull时如果远程有多个分支可以选择要pull的分支，批量pull仍然以数据库的为准
- 增加切换本地分支和远程分支的功能

## 新增功能说明

### 1. 应用实例管理
- 创建新的应用实例：
  ```bash
  mgit new newapp  # 创建一个名为 newapp 的新应用实例
  mgit new         # 交互式创建新应用实例
  ```
  - 新应用将继承当前应用的所有功能
  - 每个应用实例拥有独立的配置文件和数据库
  - 应用名称会自动清理非法字符
  - 如果应用已存在，将提示直接使用现有应用

### 2. 配置文件命名规则
应用会根据程序名称自动选择配置文件：
- mgit.exe：使用 `.env` 文件
- 其他应用：使用 `.应用名.env` 文件
例如：gitmanager.exe -> .gitmanager.env

### 3. 数据库文件命名规则
数据库文件同样遵循应用名称规则：
- mgit.exe：使用 `projects.db` 文件
- 其他应用：使用 `应用名.db` 文件
例如：gitmanager.exe -> gitmanager.db

### 4. 远程数据库目录
远程数据库目录也采用应用名称命名：
- mgit.exe：使用 `.mgit_db` 目录
- 其他应用：使用 `.应用名_db` 目录
例如：gitmanager.exe -> .gitmanager_db

### 5. 环境变量设置
自动设置当前应用目录到系统环境变量：
```bash
mgit env  # 将自动创建 MGIT_PATH 或 应用名_PATH 环境变量
```
- 环境变量名称根据应用名自动生成
- 自动检查是否已存在相同的环境变量
- Windows 系统需要管理员权限
- 某些设置可能需要重启系统后生效

### 注意事项
1. 创建新应用实例时：
   - 会自动启动新创建的应用
   - 新应用拥有独立的配置和数据
   - 如果同名应用已存在则无法创建

2. 环境变量设置：
   - Windows 系统需要管理员权限
   - 可能需要重启系统使环境变量生效
   - 建议使用英文路径避免编码问题

3. 数据存储：
   - 所有配置文件和数据库都存放在应用所在目录
   - 远程数据库目录作为独立的 Git 仓库管理
   - 支持多设备间的配置自动同步

## 安装说明

### 初始化配置
```bash
# 初始化工具
./mgit init
# 或者
./mgit init mgit
```

### 环境配置
```bash
# 设置机器标识
./mgit set machine your-machine-name

# 设置项目存放路径
./mgit set path /your/custom/path

# 查看当前配置
./mgit set
```

## 基础配置
工具使用 .env 文件存储配置：

```env
# 机器标识
MACHINE_ID=machine-01

# 应用路径（存放所有项目的父目录）
APP_PATH=/path/to/projects

# 环境变量设置
MGIT_LANG=zh-CN
```

## 数据库和配置文件说明

### 文件命名规则
应用会根据程序名称自动选择配置文件和数据库文件：

1. 配置文件：
   - mgit.exe：使用 `.env` 文件
   - 其他应用：使用 `.应用名.env` 文件
   例如：gitmanager.exe -> .gitmanager.env

2. 本地数据库：
   - mgit.exe：使用 `projects.db` 文件
   - 其他应用：使用 `应用名.db` 文件
   例如：gitmanager.exe -> gitmanager.db

3. 远程数据库目录：
   - mgit.exe：使用 `.mgit_db` 目录
   - 其他应用：使用 `.应用名_db` 目录
   例如：gitmanager.exe -> .gitmanager_db

### 远程数据库同步
```bash
./mgit set
# 选择 "启用数据库仓库同步"
# 输入数据库仓库地址
```

### 数据同步机制
- 每次 push 操作后自动同步数据库
- 每次 pull 操作前自动获取最新配置
- 支持多人协作，自动合并配置

### 同步内容
- 项目配置信息（保存在 应用名.db 中）
- 分支设置
- 最后提交记录
- 设备标识

### 数据存储位置
1. 本地数据库：
   - 存放在应用程序所在目录
   - 自动根据应用名称创建对应的数据库文件

2. 远程数据库：
   - 存放在 APP_PATH 目录下的 .应用名_db 目录中
   - 作为独立的 Git 仓库进行管理
   - 自动同步和合并多设备间的配置

### 重要注意事项
- 数据库文件首次使用时自动创建
- 远程数据库同步需要有效的 Git 仓库地址
- 数据库文件会自动备份以防止数据丢失
- 切换远程数据库时会自动迁移现有数据

## 命令使用说明

### 创建应用实例
```bash
# 创建新的应用实例（指定名称）
./mgit new project_manager

# 创建新的应用实例（交互式）
./mgit new
# 根据提示输入新应用名称
```

### 初始化项目
```bash
# 初始化新项目
./mgit init project_name https://github.com/user/repo.git

# 新建项目
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
# 然后在交互式菜单中选择要同步的项目

# 拉取所有项目
./mgit pull-all
```

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
```

### 分支管理
```bash
# 交互式设置项目分支
./mgit branch

# 直接设置指定项目的本地和远程分支
./mgit branch project_name local_branch remote_branch
```

### 设置拉取分支
```bash
# 交互式设置项目拉取分支
./mgit set pull-branch

# 为特定项目直接设置拉取分支
./mgit set pull-branch project_name branch_name
```

### 项目管理
```bash
# 查看项目列表
./mgit list

# 删除项目
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

## 交互式菜单示例

### 项目选择菜单
1. 拉取/推送单个项目：
```bash
? 选择要操作的项目（使用方向键）
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [取消]
```

2. 分支管理：
```bash
? 选择要设置分支的项目
❯ project1（当前：main -> origin/main）
  project2（当前：develop -> origin/develop）
  [取消]

? 输入本地分支名称：main
? 输入远程分支名称：origin/main
```

3. 项目删除：
```bash
? 选择要删除的项目
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [取消]

? 确认删除 project1？(Y/n)
```

## 环境变量设置说明

### 命令用法
```bash
mgit env
```

### 功能说明
- 自动将应用程序所在目录设置为系统环境变量
- 环境变量名称会根据应用程序名称自动生成，格式为 `应用名_PATH`
- 例如：
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### 系统环境变量配置

将 MGit 添加到系统 PATH，实现全局访问：

1. Windows：
```powershell
# 将 MGit 目录添加到用户 PATH
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS：
```bash
# 编辑 ~/.bashrc 或 ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

配置完成后，可以在任意目录使用 `mgit` 命令。工具会根据您的配置自动处理项目管理、同步和环境设置。

## English

# MGit - Git Project Management Tool

## English

# MGit - Git Project Management Tool

## Introduction
MGit is a command-line tool for managing multiple Git projects. It helps you efficiently manage, synchronize, and update multiple Git repositories.

## Main Features
- Multi-project Management: Manage multiple Git projects with one click, improving work efficiency
- Multi-language Support: Supports multiple languages including Chinese (Simplified/Traditional), English, Japanese, Korean, French
- Remote Database Sync: Supports project configuration synchronization across multiple devices for better team collaboration
- Version Control: Built-in version management with automatic updates
- System Integration: Can be added to system environment variables for use anywhere
- Smart Branch Management: Automatically detects and switches branches to avoid conflicts
- Batch Operations: Supports pulling/pushing changes for multiple projects simultaneously

## Feature Details

### 1. Efficient Collaboration
- Automatic configuration sync across multiple devices ensures team members use the same project settings
- Unified team project management prevents configuration inconsistencies
- Batch operations save time and improve work efficiency
- Smart branch management automatically handles branch switching and synchronization

### 2. User-Friendly
- Interactive command-line interface with intuitive operations
- Multi-language localization support eliminates language barriers
- Smart prompts and help reduce learning curve
- Unified command format, easy to remember and use

### 3. Safe and Reliable
- Automatic configuration backup prevents accidental loss
- Version control protection supports rollback operations
- Intelligent conflict handling ensures data safety
- Permission management prevents misoperations

### 4. Highly Extensible
- Supports custom configurations to meet different needs
- Flexible project management adapts to various scenarios
- Continuous updates and iterations for ongoing optimization
- Supports plugin extensions for expandable functionality

## Version Update Notes
1.0.16
- Added application instance management with `mgit new` command
- Optimized configuration and database file naming rules based on application name
- Added automatic environment variable setting feature
- Updated remote database with forced updates (if enabled)
- Added branch selection for single pull when multiple remote branches exist
- Added local and remote branch switching functionality

## New Feature Details

### 1. Application Instance Management
- Create new application instances:
  ```bash
  mgit new newapp  # Create a new application instance named newapp
  mgit new         # Create interactively
  ```
  - New applications inherit all features from the current application
  - Each instance has independent configuration files and database
  - Application names are automatically cleaned of illegal characters
  - If the application exists, you'll be prompted to use the existing one

### 2. Configuration File Naming Rules
The application automatically selects configuration files based on program name:
- mgit.exe: Uses `.env` file
- Other apps: Uses `.appname.env` file
Example: gitmanager.exe -> .gitmanager.env

### 3. Database File Naming Rules
Database files follow the same application name convention:
- mgit.exe: Uses `projects.db` file
- Other apps: Uses `appname.db` file
Example: gitmanager.exe -> gitmanager.db

### 4. Remote Database Directory
Remote database directories also use application name:
- mgit.exe: Uses `.mgit_db` directory
- Other apps: Uses `.appname_db` directory
Example: gitmanager.exe -> .gitmanager_db

### 5. Environment Variable Settings
Automatically set current application directory to system environment variables:
```bash
mgit env  # Will automatically create MGIT_PATH or appname_PATH environment variable
```
- Environment variable names are generated based on application name
- Automatically checks for existing identical environment variables
- Windows systems require administrator privileges
- Some settings may require system restart to take effect

### Notes
1. When creating new application instances:
   - The new application will start automatically
   - New applications have independent configurations and data
   - Cannot create if an application with the same name exists

2. Environment Variable Settings:
   - Windows systems require administrator privileges
   - System restart may be required for environment variables to take effect
   - English paths recommended to avoid encoding issues

3. Data Storage:
   - All configuration files and databases are stored in the application directory
   - Remote database directory is managed as an independent Git repository
   - Supports automatic configuration sync across multiple devices

## Installation Guide

### Initial Configuration
```bash
# Initialize tool
./mgit init
# Or
./mgit init mgit
```

### Environment Configuration
```bash
# Set machine identifier
./mgit set machine your-machine-name

# Set project storage path
./mgit set path /your/custom/path

# View current configuration
./mgit set
```

## Basic Configuration
The tool uses .env file to store configuration:

```env
# Machine identifier
MACHINE_ID=machine-01

# Application path (parent directory for all projects)
APP_PATH=/path/to/projects

# Environment variable settings
MGIT_LANG=en-US
```

## Database and Configuration File Description

### File Naming Rules
The application automatically selects configuration files and database files based on program name:

1. Configuration Files:
   - mgit.exe: Uses `.env` file
   - Other apps: Uses `.appname.env` file
   Example: gitmanager.exe -> .gitmanager.env

2. Local Database:
   - mgit.exe: Uses `projects.db` file
   - Other apps: Uses `appname.db` file
   Example: gitmanager.exe -> gitmanager.db

3. Remote Database Directory:
   - mgit.exe: Uses `.mgit_db` directory
   - Other apps: Uses `.appname_db` directory
   Example: gitmanager.exe -> .gitmanager_db

### Remote Database Synchronization
```bash
./mgit set
# Select "Enable database repository sync"
# Enter database repository address
```

### Data Synchronization Mechanism
- Automatically syncs database after each push operation
- Automatically gets latest configuration before each pull operation
- Supports multi-user collaboration with automatic configuration merging

### Synchronized Content
- Project configuration information (stored in appname.db)
- Branch settings
- Last commit records
- Device identifiers

### Data Storage Location
1. Local Database:
   - Stored in application directory
   - Database files automatically created based on application name

2. Remote Database:
   - Stored in .appname_db directory under APP_PATH
   - Managed as independent Git repository
   - Automatic sync and merge of configurations across devices

### Important Notes
- Database files are created automatically on first use
- Remote database sync requires valid Git repository address
- Database files are automatically backed up to prevent data loss
- Existing data is automatically migrated when switching remote databases

## Command Usage Guide

### Create Application Instance
```bash
# Create new application instance (specify name)
./mgit new project_manager

# Create new application instance (interactive)
./mgit new
# Follow prompts to enter new application name
```

### Initialize Project
```bash
# Initialize new project
./mgit init project_name https://github.com/user/repo.git

# Create new project
./mgit init mgit
# Or simply
./mgit init
```

### Pull Code
```bash
# Pull single project
./mgit pull project_name
# Or
./mgit pull
# Then select project to sync from interactive menu

# Pull all projects
./mgit pull-all
```

### Push Code
```bash
# Push single project
./mgit push project_name
# Or
./mgit push
# Then select project to push from interactive menu
# Enter commit message (leave empty to use machine name)

# Push all projects
./mgit push-all
```

### Branch Management
```bash
# Interactive branch setting
./mgit branch

# Set local and remote branches for specific project
./mgit branch project_name local_branch remote_branch
```

### Set Pull Branch
```bash
# Interactive pull branch setting
./mgit set pull-branch

# Set pull branch for specific project
./mgit set pull-branch project_name branch_name
```

### Project Management
```bash
# View project list
./mgit list

# Delete project
./mgit delete
# Select project to delete from interactive menu
```

### View Help
```bash
./mgit help
# Or use aliases
./mgit h
./mgit -h
./mgit -help
```

## Interactive Menu Examples

### Project Selection Menu
1. Pull/Push Single Project:
```bash
? Select project to operate (Use arrow keys)
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [Cancel]
```

2. Branch Management:
```bash
? Select project to set branch
❯ project1 (current: main -> origin/main)
  project2 (current: develop -> origin/develop)
  [Cancel]

? Input local branch name: main
? Input remote branch name: origin/main
```

3. Project Deletion:
```bash
? Select project to delete
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [Cancel]

? Confirm to delete project1? (Y/n)
```

## Environment Variable Settings Guide

### Command Usage
```bash
mgit env
```

### Function Description
- Automatically sets application directory as system environment variable
- Environment variable names are generated based on application name, format: `appname_PATH`
- Examples:
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

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

After completing all configurations, you can use MGit from any directory. The tool will automatically handle project management, synchronization, and environment settings based on your configuration.


## Japanese

# MGit - Git プロジェクト管理ツール

# MGit - Git プロジェクト管理ツール

## はじめに
MGitは、複数のGitプロジェクトを管理するためのコマンドラインツールです。複数のGitリポジトリを効率的に管理、同期、更新することができます。

## 主な機能
- マルチプロジェクト管理：複数のGitプロジェクトをワンクリックで管理し、作業効率を向上
- 多言語対応：日本語、中国語（簡体字/繁体字）、英語、韓国語、フランス語など多言語をサポート
- リモートデータベース同期：複数デバイス間のプロジェクト設定同期をサポート
- バージョン管理：自動更新機能付きの内蔵バージョン管理
- システム統合：システム環境変数に追加可能で、どこでも使用可能
- スマートブランチ管理：ブランチの自動検出と切り替えで競合を回避
- 一括操作：複数プロジェクトの変更を同時にプル/プッシュ可能

## 機能の詳細

### 1. 効率的なコラボレーション
- 複数デバイス間での設定自動同期により、チームメンバーが同じプロジェクト設定を使用可能
- チームプロジェクトの統一管理で設定の不一致を防止
- 一括操作で時間を節約し、作業効率を向上
- スマートブランチ管理でブランチの切り替えと同期を自動処理

### 2. ユーザーフレンドリー
- 直感的な操作が可能なインタラクティブなコマンドラインインターフェース
- 多言語ローカライゼーションで言語の壁を解消
- スマートなプロンプトとヘルプで学習曲線を緩和
- 統一されたコマンド形式で覚えやすく使いやすい

### 3. 安全性と信頼性
- 設定の自動バックアップで偶発的な損失を防止
- バージョン管理保護でロールバック操作をサポート
- インテリジェントな競合処理でデータの安全性を確保
- 権限管理で誤操作を防止

### 4. 高い拡張性
- カスタム設定で様々なニーズに対応
- 柔軟なプロジェクト管理で様々なシナリオに適応
- 継続的な更新と改善で最適化を継続
- プラグイン拡張をサポートし、機能を拡張可能

## バージョン更新情報
1.0.16
- `mgit new` コマンドでアプリケーションインスタンス管理機能を追加
- アプリケーション名に基づく設定とデータベースファイルの命名規則を最適化
- 環境変数自動設定機能を追加
- リモートデータベースの強制更新機能を追加（有効な場合）
- 複数のリモートブランチが存在する場合の単一プル時のブランチ選択機能を追加
- ローカルとリモートのブランチ切り替え機能を追加

## 新機能の詳細

### 1. アプリケーションインスタンス管理
- 新しいアプリケーションインスタンスの作成：
  ```bash
  mgit new newapp  # newappという名前の新しいアプリケーションインスタンスを作成
  mgit new         # 対話的に作成
  ```
  - 新しいアプリケーションは現在のアプリケーションの全機能を継承
  - 各インスタンスは独立した設定ファイルとデータベースを保持
  - アプリケーション名から不正な文字を自動的に除去
  - 既存のアプリケーションが存在する場合は、それを使用するよう促す

### 2. 設定ファイルの命名規則
プログラム名に基づいて設定ファイルを自動選択：
- mgit.exe：`.env` ファイルを使用
- その他のアプリ：`.アプリ名.env` ファイルを使用
例：gitmanager.exe -> .gitmanager.env

### 3. データベースファイルの命名規則
データベースファイルもアプリケーション名の規則に従う：
- mgit.exe：`projects.db` ファイルを使用
- その他のアプリ：`アプリ名.db` ファイルを使用
例：gitmanager.exe -> gitmanager.db

### 4. リモートデータベースディレクトリ
リモートデータベースディレクトリもアプリケーション名を使用：
- mgit.exe：`.mgit_db` ディレクトリを使用
- その他のアプリ：`.アプリ名_db` ディレクトリを使用
例：gitmanager.exe -> .gitmanager_db

### 5. 環境変数設定
現在のアプリケーションディレクトリをシステム環境変数に自動設定：
```bash
mgit env  # MGIT_PATH またはアプリ名_PATH 環境変数を自動作成
```
- 環境変数名はアプリケーション名に基づいて生成
- 同一の環境変数が存在するか自動チェック
- Windowsシステムでは管理者権限が必要
- 一部の設定はシステム再起動後に有効化

### 注意事項
1. 新しいアプリケーションインスタンス作成時：
   - 新しいアプリケーションは自動的に起動
   - 新しいアプリケーションは独立した設定とデータを持つ
   - 同名のアプリケーションが存在する場合は作成不可

2. 環境変数設定：
   - Windowsシステムでは管理者権限が必要
   - 環境変数の有効化にはシステム再起動が必要な場合がある
   - エンコーディングの問題を避けるため英語のパスを推奨

3. データストレージ：
   - すべての設定ファイルとデータベースはアプリケーションディレクトリに保存
   - リモートデータベースディレクトリは独立したGitリポジトリとして管理
   - 複数デバイス間の設定自動同期をサポート

## インストールガイド

### 初期設定
```bash
# ツールの初期化
./mgit init
# または
./mgit init mgit
```

### 環境設定
```bash
# マシン識別子の設定
./mgit set machine your-machine-name

# プロジェクト保存パスの設定
./mgit set path /your/custom/path

# 現在の設定を表示
./mgit set
```

## 基本設定
ツールは .env ファイルで設定を保存：

```env
# マシン識別子
MACHINE_ID=machine-01

# アプリケーションパス（全プロジェクトの親ディレクトリ）
APP_PATH=/path/to/projects

# 環境変数設定
MGIT_LANG=ja-JP
```

## データベースと設定ファイルの説明

### ファイルの命名規則
プログラム名に基づいて設定ファイルとデータベースファイルを自動選択：

1. 設定ファイル：
   - mgit.exe：`.env` ファイルを使用
   - その他のアプリ：`.アプリ名.env` ファイルを使用
   例：gitmanager.exe -> .gitmanager.env

2. ローカルデータベース：
   - mgit.exe：`projects.db` ファイルを使用
   - その他のアプリ：`アプリ名.db` ファイルを使用
   例：gitmanager.exe -> gitmanager.db

3. リモートデータベースディレクトリ：
   - mgit.exe：`.mgit_db` ディレクトリを使用
   - その他のアプリ：`.アプリ名_db` ディレクトリを使用
   例：gitmanager.exe -> .gitmanager_db

### リモートデータベース同期
```bash
./mgit set
# "データベースリポジトリ同期を有効にする" を選択
# データベースリポジトリアドレスを入力
```

### データ同期メカニズム
- プッシュ操作後に自動的にデータベースを同期
- プル操作前に自動的に最新の設定を取得
- 複数ユーザーのコラボレーションをサポートし、設定を自動マージ

### 同期内容
- プロジェクト設定情報（アプリ名.dbに保存）
- ブランチ設定
- 最終コミット記録
- デバイス識別子

### データ保存場所
1. ローカルデータベース：
   - アプリケーションディレクトリに保存
   - アプリケーション名に基づいて自動的にデータベースファイルを作成

2. リモートデータベース：
   - APP_PATHディレクトリ下の.アプリ名_dbディレクトリに保存
   - 独立したGitリポジトリとして管理
   - 複数デバイス間の設定を自動同期およびマージ

### 重要な注意事項
- データベースファイルは初回使用時に自動作成
- リモートデータベース同期には有効なGitリポジトリアドレスが必要
- データベースファイルは自動的にバックアップされ、データ損失を防止
- リモートデータベース切り替え時に既存データは自動的に移行

## コマンド使用ガイド

### アプリケーションインスタンスの作成
```bash
# 新しいアプリケーションインスタンスを作成（名前を指定）
./mgit new project_manager

# 新しいアプリケーションインスタンスを作成（対話的）
./mgit new
# プロンプトに従って新しいアプリケーション名を入力
```

### プロジェクトの初期化
```bash
# 新しいプロジェクトを初期化
./mgit init project_name https://github.com/user/repo.git

# 新規プロジェクト作成
./mgit init mgit
# または単に
./mgit init
```

### コードのプル
```bash
# 単一プロジェクトのプル
./mgit pull project_name
# または
./mgit pull
# 対話式メニューから同期するプロジェクトを選択

# すべてのプロジェクトをプル
./mgit pull-all
```

### コードのプッシュ
```bash
# 単一プロジェクトのプッシュ
./mgit push project_name
# または
./mgit push
# 対話式メニューからプッシュするプロジェクトを選択
# コミットメッセージを入力（空の場合はマシン名を使用）

# すべてのプロジェクトをプッシュ
./mgit push-all
```

### ブランチ管理
```bash
# 対話的なブランチ設定
./mgit branch

# 特定のプロジェクトのローカルとリモートブランチを設定
./mgit branch project_name local_branch remote_branch
```

### プルブランチの設定
```bash
# 対話的なプルブランチ設定
./mgit set pull-branch

# 特定のプロジェクトのプルブランチを設定
./mgit set pull-branch project_name branch_name
```

### プロジェクト管理
```bash
# プロジェクトリストの表示
./mgit list

# プロジェクトの削除
./mgit delete
# 対話式メニューから削除するプロジェクトを選択
```

### ヘルプの表示
```bash
./mgit help
# または以下のエイリアスを使用
./mgit h
./mgit -h
./mgit -help
```

## インタラクティブメニューの例

### プロジェクト選択メニュー
1. 単一プロジェクトのプル/プッシュ：
```bash
? 操作するプロジェクトを選択 (矢印キーを使用)
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [キャンセル]
```

2. ブランチ管理：
```bash
? ブランチを設定するプロジェクトを選択
❯ project1 (現在: main -> origin/main)
  project2 (現在: develop -> origin/develop)
  [キャンセル]

? ローカルブランチ名を入力: main
? リモートブランチ名を入力: origin/main
```

3. プロジェクト削除：
```bash
? 削除するプロジェクトを選択
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [キャンセル]

? project1を削除しますか？ (Y/n)
```

## 環境変数設定ガイド

### コマンド使用法
```bash
mgit env
```

### 機能説明
- アプリケーションディレクトリを自動的にシステム環境変数に設定
- 環境変数名はアプリケーション名に基づいて生成、形式：`アプリ名_PATH`
- 例：
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### システム環境変数の設定

グローバルアクセス用にMGitをシステムPATHに追加：

1. Windows:
```powershell
# MGitディレクトリをユーザーPATHに追加
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# ~/.bashrcまたは~/.zshrcを編集
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

設定完了後、任意のディレクトリから `mgit` コマンドを使用できます。ツールは設定に基づいてプロジェクト管理、同期、環境設定を自動的に処理します。
```


## Korean

# MGit - Git 프로젝트 관리 도구

# MGit - Git 프로젝트 관리 도구

## 소개
MGit은 여러 Git 프로젝트를 관리하기 위한 명령줄 도구입니다. 여러 Git 저장소를 효율적으로 관리, 동기화 및 업데이트할 수 있습니다.

## 주요 기능
- 다중 프로젝트 관리: 여러 Git 프로젝트를 원클릭으로 관리하여 작업 효율성 향상
- 다국어 지원: 한국어, 중국어(간체/번체), 영어, 일본어, 프랑스어 등 다양한 언어 지원
- 원격 데이터베이스 동기화: 여러 장치 간의 프로젝트 설정 동기화 지원
- 버전 관리: 자동 업데이트가 포함된 내장 버전 관리
- 시스템 통합: 시스템 환경 변수에 추가하여 어디서나 사용 가능
- 스마트 브랜치 관리: 브랜치 자동 감지 및 전환으로 충돌 방지
- 일괄 작업: 여러 프로젝트의 변경 사항을 동시에 풀/푸시 지원

## 기능 상세

### 1. 효율적인 협업
- 여러 장치 간 설정 자동 동기화로 팀원들이 동일한 프로젝트 설정 사용
- 팀 프로젝트 통합 관리로 설정 불일치 방지
- 일괄 작업으로 시간 절약 및 작업 효율성 향상
- 스마트 브랜치 관리로 브랜치 전환 및 동기화 자동 처리

### 2. 사용자 친화적
- 직관적인 조작이 가능한 대화형 명령줄 인터페이스
- 다국어 현지화로 언어 장벽 해소
- 스마트한 프롬프트와 도움말로 학습 곡선 완화
- 통일된 명령어 형식으로 기억하고 사용하기 쉬움

### 3. 안전성과 신뢰성
- 설정 자동 백업으로 우발적 손실 방지
- 버전 관리 보호로 롤백 작업 지원
- 지능적인 충돌 처리로 데이터 안전성 보장
- 권한 관리로 오조작 방지

### 4. 높은 확장성
- 사용자 정의 설정으로 다양한 요구 사항 충족
- 유연한 프로젝트 관리로 다양한 시나리오 대응
- 지속적인 업데이트와 개선으로 최적화 유지
- 플러그인 확장 지원으로 기능 확장 가능

## 버전 업데이트 정보
1.0.16
- `mgit new` 명령으로 애플리케이션 인스턴스 관리 기능 추가
- 애플리케이션 이름 기반 설정 및 데이터베이스 파일 명명 규칙 최적화
- 환경 변수 자동 설정 기능 추가
- 원격 데이터베이스 강제 업데이트 기능 추가(활성화된 경우)
- 여러 원격 브랜치가 있는 경우 단일 풀 시 브랜치 선택 기능 추가
- 로컬 및 원격 브랜치 전환 기능 추가

## 새로운 기능 상세

### 1. 애플리케이션 인스턴스 관리
- 새로운 애플리케이션 인스턴스 생성:
  ```bash
  mgit new newapp  # newapp이라는 이름의 새 애플리케이션 인스턴스 생성
  mgit new         # 대화형으로 생성
  ```
  - 새 애플리케이션은 현재 애플리케이션의 모든 기능 상속
  - 각 인스턴스는 독립적인 설정 파일과 데이터베이스 보유
  - 애플리케이션 이름에서 잘못된 문자 자동 제거
  - 동일한 이름의 애플리케이션이 존재하는 경우 해당 앱 사용 안내

### 2. 설정 파일 명명 규칙
프로그램 이름을 기반으로 설정 파일 자동 선택:
- mgit.exe: `.env` 파일 사용
- 기타 앱: `.앱이름.env` 파일 사용
예: gitmanager.exe -> .gitmanager.env

### 3. 데이터베이스 파일 명명 규칙
데이터베이스 파일도 애플리케이션 이름 규칙을 따름:
- mgit.exe: `projects.db` 파일 사용
- 기타 앱: `앱이름.db` 파일 사용
예: gitmanager.exe -> gitmanager.db

### 4. 원격 데이터베이스 디렉터리
원격 데이터베이스 디렉터리도 애플리케이션 이름 사용:
- mgit.exe: `.mgit_db` 디렉터리 사용
- 기타 앱: `.앱이름_db` 디렉터리 사용
예: gitmanager.exe -> .gitmanager_db

### 5. 환경 변수 설정
현재 애플리케이션 디렉터리를 시스템 환경 변수에 자동 설정:
```bash
mgit env  # MGIT_PATH 또는 앱이름_PATH 환경 변수 자동 생성
```
- 환경 변수 이름은 애플리케이션 이름 기반으로 생성
- 동일한 환경 변수 존재 여부 자동 확인
- Windows 시스템에서는 관리자 권한 필요
- 일부 설정은 시스템 재시작 후 적용

### 주의 사항
1. 새로운 애플리케이션 인스턴스 생성 시:
   - 새 애플리케이션이 자동으로 시작됨
   - 새 애플리케이션은 독립적인 설정과 데이터를 가짐
   - 동일한 이름의 애플리케이션이 존재하는 경우 생성 불가

2. 환경 변수 설정:
   - Windows 시스템에서는 관리자 권한 필요
   - 환경 변수 적용을 위해 시스템 재시작이 필요할 수 있음
   - 인코딩 문제를 피하기 위해 영문 경로 권장

3. 데이터 저장:
   - 모든 설정 파일과 데이터베이스는 애플리케이션 디렉터리에 저장
   - 원격 데이터베이스 디렉터리는 독립 Git 저장소로 관리
   - 여러 장치 간 설정 자동 동기화 지원

## 설치 가이드

### 초기 설정
```bash
# 도구 초기화
./mgit init
# 또는
./mgit init mgit
```

### 환경 설정
```bash
# 기기 식별자 설정
./mgit set machine your-machine-name

# 프로젝트 저장 경로 설정
./mgit set path /your/custom/path

# 현재 설정 확인
./mgit set
```

## 기본 설정
도구는 .env 파일로 설정을 저장:

```env
# 기기 식별자
MACHINE_ID=machine-01

# 애플리케이션 경로(모든 프로젝트의 상위 디렉터리)
APP_PATH=/path/to/projects

# 환경 변수 설정
MGIT_LANG=ko-KR
```

## 데이터베이스와 설정 파일 설명

### 파일 명명 규칙
프로그램 이름을 기반으로 설정 파일과 데이터베이스 파일 자동 선택:

1. 설정 파일:
   - mgit.exe: `.env` 파일 사용
   - 기타 앱: `.앱이름.env` 파일 사용
   예: gitmanager.exe -> .gitmanager.env

2. 로컬 데이터베이스:
   - mgit.exe: `projects.db` 파일 사용
   - 기타 앱: `앱이름.db` 파일 사용
   예: gitmanager.exe -> gitmanager.db

3. 원격 데이터베이스 디렉터리:
   - mgit.exe: `.mgit_db` 디렉터리 사용
   - 기타 앱: `.앱이름_db` 디렉터리 사용
   예: gitmanager.exe -> .gitmanager_db

### 원격 데이터베이스 동기화
```bash
./mgit set
# "데이터베이스 저장소 동기화 활성화" 선택
# 데이터베이스 저장소 주소 입력
```

### 데이터 동기화 메커니즘
- 푸시 작업 후 자동으로 데이터베이스 동기화
- 풀 작업 전 자동으로 최신 설정 가져오기
- 여러 사용자 협업 지원, 설정 자동 병합

### 동기화 내용
- 프로젝트 설정 정보(앱이름.db에 저장)
- 브랜치 설정
- 마지막 커밋 기록
- 장치 식별자

### 데이터 저장 위치
1. 로컬 데이터베이스:
   - 애플리케이션 디렉터리에 저장
   - 애플리케이션 이름 기반으로 자동 데이터베이스 파일 생성

2. 원격 데이터베이스:
   - APP_PATH 디렉터리 아래 .앱이름_db 디렉터리에 저장
   - 독립 Git 저장소로 관리
   - 여러 장치 간 설정 자동 동기화 및 병합

### 중요 주의 사항
- 데이터베이스 파일은 첫 사용 시 자동 생성됨
- 원격 데이터베이스 동기화에는 유효한 Git 저장소 주소 필요
- 데이터베이스 파일은 자동으로 백업되어 데이터 손실 방지
- 원격 데이터베이스 전환 시 기존 데이터는 자동으로 이전됨

## 명령어 사용 가이드

### 애플리케이션 인스턴스 생성
```bash
# 새 애플리케이션 인스턴스 생성(이름 지정)
./mgit new project_manager

# 새 애플리케이션 인스턴스 생성(대화형)
./mgit new
# 프롬프트에 따라 새 애플리케이션 이름 입력
```

### 프로젝트 초기화
```bash
# 새 프로젝트 초기화
./mgit init project_name https://github.com/user/repo.git

# 새 프로젝트 생성
./mgit init mgit
# 또는 간단히
./mgit init
```

### 코드 풀
```bash
# 단일 프로젝트 풀
./mgit pull project_name
# 또는
./mgit pull
# 대화형 메뉴에서 동기화할 프로젝트 선택

# 모든 프로젝트 풀
./mgit pull-all
```

### 코드 푸시
```bash
# 단일 프로젝트 푸시
./mgit push project_name
# 또는
./mgit push
# 대화형 메뉴에서 푸시할 프로젝트 선택
# 커밋 메시지 입력(비워두면 기기 이름 사용)

# 모든 프로젝트 푸시
./mgit push-all
```

### 브랜치 관리
```bash
# 대화형 브랜치 설정
./mgit branch

# 특정 프로젝트의 로컬 및 원격 브랜치 설정
./mgit branch project_name local_branch remote_branch
```

### 풀 브랜치 설정
```bash
# 대화형 풀 브랜치 설정
./mgit set pull-branch

# 특정 프로젝트의 풀 브랜치 설정
./mgit set pull-branch project_name branch_name
```

### 프로젝트 관리
```bash
# 프로젝트 목록 보기
./mgit list

# 프로젝트 삭제
./mgit delete
# 대화형 메뉴에서 삭제할 프로젝트 선택
```

### 도움말 보기
```bash
./mgit help
# 또는 별칭 사용
./mgit h
./mgit -h
./mgit -help
```

## 대화형 메뉴 예시

### 프로젝트 선택 메뉴
1. 단일 프로젝트 풀/푸시:
```bash
? 작업할 프로젝트 선택 (방향키 사용)
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [취소]
```

2. 브랜치 관리:
```bash
? 브랜치를 설정할 프로젝트 선택
❯ project1 (현재: main -> origin/main)
  project2 (현재: develop -> origin/develop)
  [취소]

? 로컬 브랜치 이름 입력: main
? 원격 브랜치 이름 입력: origin/main
```

3. 프로젝트 삭제:
```bash
? 삭제할 프로젝트 선택
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [취소]

? project1을 삭제하시겠습니까? (Y/n)
```

## 환경 변수 설정 안내

### 명령어 사용법
```bash
mgit env
```

### 기능 설명
- 애플리케이션 디렉터리를 자동으로 시스템 환경 변수로 설정
- 환경 변수 이름은 애플리케이션 이름을 기반으로 생성되며, 형식은 `앱이름_PATH`
- 예시:
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### 시스템 환경 변수 설정

MGit을 시스템 PATH에 추가하여 전역 액세스 실현:

1. Windows:
```powershell
# MGit 디렉터리를 사용자 PATH에 추가
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS:
```bash
# ~/.bashrc 또는 ~/.zshrc 편집
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

설정 완료 후, 어느 디렉터리에서나 `mgit` 명령어를 사용할 수 있습니다. 도구는 설정에 따라 프로젝트 관리, 동기화, 환경 설정을 자동으로 처리합니다.

## Traditional Chinese

# MGit - Git 專案管理工具

# MGit - Git 專案管理工具

## 簡介
MGit 是一個用於管理多個 Git 專案的命令列工具。它可以幫助您有效地管理、同步和更新多個 Git 儲存庫。

## 主要功能
- 多專案管理：一鍵管理多個 Git 專案，提高工作效率
- 多語言支援：支援中文（簡體/繁體）、英文、日文、韓文、法文等多種語言
- 遠端資料庫同步：支援多設備間的專案配置同步，團隊協作更便捷
- 版本控制：內建版本管理，支援自動更新
- 系統整合：支援加入到系統環境變數，隨時隨地使用
- 智慧分支管理：自動檢測和切換分支，避免衝突
- 批次操作：支援同時拉取/推送多個專案的更改

## 功能詳細說明

### 1. 高效協作
- 多設備配置自動同步，確保團隊成員使用相同的專案配置
- 團隊專案統一管理，避免配置不一致
- 批次操作節省時間，提高工作效率
- 智慧分支管理，自動處理分支切換和同步

### 2. 用戶友好
- 互動式命令列介面，操作直觀簡單
- 多語言本地化支援，無語言障礙
- 智慧提示和幫助，降低使用門檻
- 統一的命令格式，易於記憶和使用

### 3. 安全可靠
- 自動備份配置，防止意外丟失
- 版本控制保護，支援回滾操作
- 衝突智慧處理，確保資料安全
- 權限管理，防止誤操作

### 4. 擴展性強
- 支援自定義配置，滿足不同需求
- 靈活的專案管理，適應各種場景
- 持續更新迭代，不斷優化功能
- 支援外掛擴展，功能可擴展

## 版本更新說明
1.0.16
- 新增應用實例管理功能，支援通過 `mgit new` 創建新的應用實例
- 優化配置文件和資料庫命名規則，根據應用名自動生成對應文件名
- 新增環境變數自動設置功能，支援將應用目錄添加到系統環境變數
- 更新遠端資料庫每次強制更新（如果開啟遠端資料庫）
- 單個 pull 時如果遠端有多個分支可以選擇要 pull 的分支，批次 pull 仍然以資料庫的為準
- 增加切換本地分支和遠端分支的功能

## 新增功能說明

### 1. 應用實例管理
- 創建新的應用實例：
  ```bash
  mgit new newapp  # 創建一個名為 newapp 的新應用實例
  mgit new         # 互動式創建新應用實例
  ```
  - 新應用將繼承當前應用的所有功能
  - 每個應用實例擁有獨立的配置文件和資料庫
  - 應用名稱會自動清理非法字符
  - 如果應用已存在，將提示直接使用現有應用

### 2. 配置文件命名規則
應用會根據程式名稱自動選擇配置文件：
- mgit.exe：使用 `.env` 文件
- 其他應用：使用 `.應用名.env` 文件
例如：gitmanager.exe -> .gitmanager.env

### 3. 資料庫文件命名規則
資料庫文件同樣遵循應用名稱規則：
- mgit.exe：使用 `projects.db` 文件
- 其他應用：使用 `應用名.db` 文件
例如：gitmanager.exe -> gitmanager.db

### 4. 遠端資料庫目錄
遠端資料庫目錄也採用應用名稱命名：
- mgit.exe：使用 `.mgit_db` 目錄
- 其他應用：使用 `.應用名_db` 目錄
例如：gitmanager.exe -> .gitmanager_db

### 5. 環境變數設置
自動設置當前應用目錄到系統環境變數：
```bash
mgit env  # 將自動創建 MGIT_PATH 或 應用名_PATH 環境變數
```
- 環境變數名稱根據應用名自動生成
- 自動檢查是否已存在相同的環境變數
- Windows 系統需要管理員權限
- 某些設置可能需要重啟系統後生效

### 注意事項
1. 創建新應用實例時：
   - 會自動啟動新創建的應用
   - 新應用擁有獨立的配置和資料
   - 如果同名應用已存在則無法創建

2. 環境變數設置：
   - Windows 系統需要管理員權限
   - 可能需要重啟系統使環境變數生效
   - 建議使用英文路徑避免編碼問題

3. 資料存儲：
   - 所有配置文件和資料庫都存放在應用所在目錄
   - 遠端資料庫目錄作為獨立的 Git 儲存庫管理
   - 支援多設備間的配置自動同步

## 安裝說明

### 初始化配置
```bash
# 初始化工具
./mgit init
# 或者
./mgit init mgit
```

### 環境配置
```bash
# 設置機器標識
./mgit set machine your-machine-name

# 設置專案存放路徑
./mgit set path /your/custom/path

# 查看當前配置
./mgit set
```

## 基礎配置
工具使用 .env 文件存儲配置：

```env
# 機器標識
MACHINE_ID=machine-01

# 應用路徑（存放所有專案的父目錄）
APP_PATH=/path/to/projects

# 環境變數設置
MGIT_LANG=zh-TW
```

## 資料庫和配置文件說明

### 文件命名規則
應用會根據程式名稱自動選擇配置文件和資料庫文件：

1. 配置文件：
   - mgit.exe：使用 `.env` 文件
   - 其他應用：使用 `.應用名.env` 文件
   例如：gitmanager.exe -> .gitmanager.env

2. 本地資料庫：
   - mgit.exe：使用 `projects.db` 文件
   - 其他應用：使用 `應用名.db` 文件
   例如：gitmanager.exe -> gitmanager.db

3. 遠端資料庫目錄：
   - mgit.exe：使用 `.mgit_db` 目錄
   - 其他應用：使用 `.應用名_db` 目錄
   例如：gitmanager.exe -> .gitmanager_db

### 遠端資料庫同步
```bash
./mgit set
# 選擇 "啟用資料庫儲存庫同步"
# 輸入資料庫儲存庫地址
```

### 資料同步機制
- 每次 push 操作後自動同步資料庫
- 每次 pull 操作前自動獲取最新配置
- 支援多人協作，自動合併配置

### 同步內容
- 專案配置信息（保存在 應用名.db 中）
- 分支設置
- 最後提交記錄
- 設備標識

### 資料存儲位置
1. 本地資料庫：
   - 存放在應用程式所在目錄
   - 自動根據應用名稱創建對應的資料庫文件

2. 遠端資料庫：
   - 存放在 APP_PATH 目錄下的 .應用名_db 目錄中
   - 作為獨立的 Git 儲存庫進行管理
   - 自動同步和合併多設備間的配置

### 重要注意事項
- 資料庫文件首次使用時自動創建
- 遠端資料庫同步需要有效的 Git 儲存庫地址
- 資料庫文件會自動備份以防止資料丟失
- 切換遠端資料庫時會自動遷移現有資料

## 命令使用說明

### 創建應用實例
```bash
# 創建新的應用實例（指定名稱）
./mgit new project_manager

# 創建新的應用實例（互動式）
./mgit new
# 根據提示輸入新應用名稱
```

### 初始化專案
```bash
# 初始化新專案
./mgit init project_name https://github.com/user/repo.git

# 新建專案
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

### 推送程式碼
```bash
# 推送單個專案
./mgit push project_name
# 或者
./mgit push
# 然後在互動式選單中選擇要推送的專案
# 之後可以輸入提交說明（留空則使用機器名）

# 推送所有專案
./mgit push-all
```

### 分支管理
```bash
# 互動式設置專案分支
./mgit branch

# 直接設置指定專案的本地和遠端分支
./mgit branch project_name local_branch remote_branch
```

### 設置拉取分支
```bash
# 互動式設置專案拉取分支
./mgit set pull-branch

# 為特定專案直接設置拉取分支
./mgit set pull-branch project_name branch_name
```

### 專案管理
```bash
# 查看專案列表
./mgit list

# 刪除專案
./mgit delete
# 通過互動式選單選擇要刪除的專案
```

### 查看幫助
```bash
./mgit help
# 或者使用以下別名
./mgit h
./mgit -h
./mgit -help
```

## 互動式選單範例

### 專案選擇選單
1. 拉取/推送單一專案：
```bash
? 選擇要操作的專案（使用方向鍵）
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [取消]
```

2. 分支管理：
```bash
? 選擇要設定分支的專案
❯ project1（目前：main -> origin/main）
  project2（目前：develop -> origin/develop）
  [取消]

? 輸入本地分支名稱：main
? 輸入遠端分支名稱：origin/main
```

3. 專案刪除：
```bash
? 選擇要刪除的專案
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [取消]

? 確認刪除 project1？(Y/n)
```

## 環境變數設置說明

### 命令用法
```bash
mgit env
```

### 功能說明
- 自動將應用程式所在目錄設置為系統環境變數
- 環境變數名稱會根據應用程式名稱自動生成，格式為 `應用名_PATH`
- 例如：
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### 系統環境變數配置

將 MGit 添加到系統 PATH，實現全域存取：

1. Windows：
```powershell
# 將 MGit 目錄添加到使用者 PATH
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS：
```bash
# 編輯 ~/.bashrc 或 ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

配置完成後，可以在任意目錄使用 `mgit` 命令。工具會根據您的配置自動處理專案管理、同步和環境設置。

## Français

# MGit - Outil de Gestion de Projets Git

# MGit - Outil de Gestion de Projets Git

## Introduction
MGit est un outil en ligne de commande pour gérer plusieurs projets Git. Il vous aide à gérer, synchroniser et mettre à jour efficacement plusieurs dépôts Git.

## Fonctionnalités Principales
- Gestion Multi-projets : Gérez plusieurs projets Git en un clic, améliorant l'efficacité du travail
- Support Multilingue : Prend en charge plusieurs langues dont le chinois (simplifié/traditionnel), l'anglais, le japonais, le coréen et le français
- Synchronisation de Base de Données Distante : Prend en charge la synchronisation des configurations de projet entre plusieurs appareils
- Contrôle de Version : Gestion de version intégrée avec mises à jour automatiques
- Intégration Système : Peut être ajouté aux variables d'environnement système pour une utilisation partout
- Gestion Intelligente des Branches : Détecte et change automatiquement les branches pour éviter les conflits
- Opérations par Lots : Prend en charge le pull/push simultané des modifications pour plusieurs projets

## Détails des Fonctionnalités

### 1. Collaboration Efficace
- Synchronisation automatique des configurations entre appareils pour une utilisation uniforme par l'équipe
- Gestion unifiée des projets d'équipe pour éviter les incohérences de configuration
- Les opérations par lots économisent du temps et améliorent l'efficacité
- Gestion intelligente des branches pour le changement et la synchronisation automatiques

### 2. Convivial
- Interface en ligne de commande interactive avec des opérations intuitives
- Support multilingue pour éliminer les barrières linguistiques
- Invites et aide intelligentes pour réduire la courbe d'apprentissage
- Format de commande unifié, facile à mémoriser et à utiliser

### 3. Sûr et Fiable
- Sauvegarde automatique des configurations pour prévenir les pertes accidentelles
- Protection par contrôle de version avec support des opérations de rollback
- Gestion intelligente des conflits pour assurer la sécurité des données
- Gestion des permissions pour prévenir les erreurs de manipulation

### 4. Hautement Extensible
- Prend en charge les configurations personnalisées pour répondre aux différents besoins
- Gestion flexible des projets adaptée à divers scénarios
- Mises à jour et améliorations continues pour une optimisation permanente
- Support des extensions plugin pour des fonctionnalités extensibles

## Notes de Version
1.0.16
- Ajout de la gestion d'instances d'application avec la commande `mgit new`
- Optimisation des règles de nommage des fichiers de configuration et de base de données
- Ajout de la configuration automatique des variables d'environnement
- Mise à jour forcée de la base de données distante (si activée)
- Ajout de la sélection de branche pour les pulls individuels
- Ajout de la fonctionnalité de changement de branches locales et distantes

## Nouvelles Fonctionnalités

### 1. Gestion des Instances d'Application
- Créer de nouvelles instances d'application :
  ```bash
  mgit new newapp  # Créer une nouvelle instance nommée newapp
  mgit new         # Création interactive
  ```
  - Les nouvelles applications héritent de toutes les fonctionnalités
  - Chaque instance possède ses propres fichiers de configuration et base de données
  - Les noms d'application sont automatiquement nettoyés des caractères invalides
  - Si l'application existe déjà, vous serez invité à utiliser l'existante

### 2. Règles de Nommage des Fichiers de Configuration
L'application sélectionne automatiquement les fichiers de configuration :
- mgit.exe : Utilise le fichier `.env`
- Autres applications : Utilise `.nomapp.env`
Exemple : gitmanager.exe -> .gitmanager.env

### 3. Règles de Nommage des Fichiers de Base de Données
Les fichiers de base de données suivent la même convention :
- mgit.exe : Utilise `projects.db`
- Autres applications : Utilise `nomapp.db`
Exemple : gitmanager.exe -> gitmanager.db

### 4. Répertoire de Base de Données Distante
Les répertoires de base de données distante utilisent également le nom de l'application :
- mgit.exe : Utilise le répertoire `.mgit_db`
- Autres applications : Utilise `.nomapp_db`
Exemple : gitmanager.exe -> .gitmanager_db

### 5. Configuration des Variables d'Environnement
Configuration automatique du répertoire de l'application dans les variables d'environnement :
```bash
mgit env  # Crée automatiquement MGIT_PATH ou nomapp_PATH
```
- Les noms des variables sont générés selon le nom de l'application
- Vérifie automatiquement l'existence de variables identiques
- Nécessite des droits administrateur sous Windows
- Certains paramètres peuvent nécessiter un redémarrage

### Notes Importantes
1. Lors de la création de nouvelles instances :
   - La nouvelle application démarre automatiquement
   - Les nouvelles applications ont des configurations et données indépendantes
   - Impossible de créer si une application du même nom existe

2. Configuration des Variables d'Environnement :
   - Nécessite des droits administrateur sous Windows
   - Peut nécessiter un redémarrage pour l'application des variables
   - Chemins en anglais recommandés pour éviter les problèmes d'encodage

3. Stockage des Données :
   - Tous les fichiers de configuration et bases de données sont stockés dans le répertoire de l'application
   - Le répertoire de base de données distante est géré comme un dépôt Git indépendant
   - Supporte la synchronisation automatique entre appareils

## Guide d'Installation

### Configuration Initiale
```bash
# Initialiser l'outil
./mgit init
# Ou
./mgit init mgit
```

### Configuration de l'Environnement
```bash
# Définir l'identifiant de la machine
./mgit set machine your-machine-name

# Définir le chemin de stockage des projets
./mgit set path /your/custom/path

# Voir la configuration actuelle
./mgit set
```

## Configuration de Base
L'outil utilise un fichier .env pour stocker la configuration :

```env
# Identifiant de la machine
MACHINE_ID=machine-01

# Chemin de l'application (répertoire parent pour tous les projets)
APP_PATH=/path/to/projects

# Configuration des variables d'environnement
MGIT_LANG=fr-FR
```

## Description de la Base de Données et des Fichiers de Configuration

### Règles de Nommage des Fichiers
L'application sélectionne automatiquement les fichiers selon le nom du programme :

1. Fichiers de Configuration :
   - mgit.exe : Utilise `.env`
   - Autres applications : Utilise `.nomapp.env`
   Exemple : gitmanager.exe -> .gitmanager.env

2. Base de Données Locale :
   - mgit.exe : Utilise `projects.db`
   - Autres applications : Utilise `nomapp.db`
   Exemple : gitmanager.exe -> gitmanager.db

3. Répertoire de Base de Données Distante :
   - mgit.exe : Utilise `.mgit_db`
   - Autres applications : Utilise `.nomapp_db`
   Exemple : gitmanager.exe -> .gitmanager_db

### Synchronisation de la Base de Données Distante
```bash
./mgit set
# Sélectionner "Activer la synchronisation du dépôt de base de données"
# Entrer l'adresse du dépôt de base de données
```

### Mécanisme de Synchronisation
- Synchronisation automatique après chaque push
- Récupération automatique de la dernière configuration avant chaque pull
- Support de la collaboration multi-utilisateurs avec fusion automatique

### Contenu Synchronisé
- Informations de configuration des projets (stockées dans nomapp.db)
- Paramètres des branches
- Derniers enregistrements de commit
- Identifiants des appareils

### Emplacements de Stockage
1. Base de Données Locale :
   - Stockée dans le répertoire de l'application
   - Fichiers de base de données créés automatiquement selon le nom de l'application

2. Base de Données Distante :
   - Stockée dans le répertoire .nomapp_db sous APP_PATH
   - Gérée comme un dépôt Git indépendant
   - Synchronisation et fusion automatiques entre appareils

### Notes Importantes
- Les fichiers de base de données sont créés automatiquement à la première utilisation
- La synchronisation distante nécessite une adresse de dépôt Git valide
- Les fichiers sont sauvegardés automatiquement pour prévenir les pertes
- Les données existantes sont migrées automatiquement lors du changement de base distante

## Guide d'Utilisation des Commandes

### Créer une Instance d'Application
```bash
# Créer une nouvelle instance (nom spécifié)
./mgit new project_manager

# Créer une nouvelle instance (interactif)
./mgit new
# Suivre les instructions pour entrer le nom
```

### Initialiser un Projet
```bash
# Initialiser un nouveau projet
./mgit init project_name https://github.com/user/repo.git

# Créer un nouveau projet
./mgit init mgit
# Ou simplement
./mgit init
```

### Tirer le Code (Pull)
```bash
# Tirer un projet unique
./mgit pull project_name
# Ou
./mgit pull
# Puis sélectionner le projet dans le menu

# Tirer tous les projets
./mgit pull-all
```

### Pousser le Code (Push)
```bash
# Pousser un projet unique
./mgit push project_name
# Ou
./mgit push
# Puis sélectionner le projet dans le menu
# Entrer le message de commit (vide pour utiliser le nom de la machine)

# Pousser tous les projets
./mgit push-all
```

### Gestion des Branches
```bash
# Configuration interactive des branches
./mgit branch

# Configurer les branches locale et distante d'un projet
./mgit branch project_name local_branch remote_branch
```

### Configurer la Branche de Pull
```bash
# Configuration interactive
./mgit set pull-branch

# Configurer pour un projet spécifique
./mgit set pull-branch project_name branch_name
```

### Gestion des Projets
```bash
# Voir la liste des projets
./mgit list

# Supprimer un projet
./mgit delete
# Sélectionner le projet à supprimer dans le menu
```

### Voir l'Aide
```bash
./mgit help
# Ou utiliser les alias
./mgit h
./mgit -h
./mgit -help
```

## Exemples de Menus Interactifs

### Menu de Sélection de Projet
1. Pull/Push d'un Projet Unique :
```bash
? Sélectionner le projet à opérer (Utiliser les flèches)
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  project3 - /path/to/project3
  [Annuler]
```

2. Gestion des Branches :
```bash
? Sélectionner le projet pour configurer les branches
❯ project1 (actuel : main -> origin/main)
  project2 (actuel : develop -> origin/develop)
  [Annuler]

? Entrer le nom de la branche locale : main
? Entrer le nom de la branche distante : origin/main
```

3. Suppression de Projet :
```bash
? Sélectionner le projet à supprimer
❯ project1 - /path/to/project1
  project2 - /path/to/project2
  [Annuler]

? Confirmer la suppression de project1 ? (O/n)
```

## Guide des Variables d'Environnement

### Utilisation de la Commande
```bash
mgit env
```

### Description des Fonctionnalités
- Configure automatiquement le répertoire de l'application comme variable d'environnement
- Le nom de la variable est généré selon le nom de l'application : `nomapp_PATH`
- Exemples :
  - mgit.exe -> MGIT_PATH
  - gitmanager.exe -> GITMANAGER_PATH

### Configuration des Variables d'Environnement Système

Ajouter MGit au PATH système pour un accès global :

1. Windows :
```powershell
# Ajouter le répertoire MGit au PATH utilisateur
setx PATH "%PATH%;D:\path\to\mgit"
```

2. Linux/macOS :
```bash
# Éditer ~/.bashrc ou ~/.zshrc
echo 'export PATH="$PATH:/path/to/mgit"' >> ~/.bashrc
source ~/.bashrc
```

Une fois la configuration terminée, vous pouvez utiliser la commande `mgit` depuis n'importe quel répertoire. L'outil gère automatiquement la gestion des projets, la synchronisation et les paramètres d'environnement selon votre configuration.

```tool_call:save_file
path: README.fr-FR.md
content: [以上法文版本的完整内容]
```

