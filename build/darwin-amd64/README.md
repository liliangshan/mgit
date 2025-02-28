# MGit - Git 项目管理工具

[简体中文](#简体中文) | [English](#english) | [日本語](#japanese) | [한국어](#korean) | [繁體中文](#traditional-chinese) | [Français](#french)

## 简体中文

MGit 是一个用于管理多个 Git 项目的命令行工具。它可以帮助您有效地管理、同步和更新多个 Git 仓库。

### 功能特点
- 初始化和管理多个 Git 项目
- 自动同步所有项目的代码
- 跟踪最后提交的机器和时间
- 支持管理工具自身的版本控制
- 灵活配置项目存放路径

### 命令
- `mgit init [项目名称]` - 初始化新项目
- `mgit pull [项目名称]` - 拉取项目更新
- `mgit push [项目名称]` - 推送项目变更
- `mgit pull-all` - 拉取所有项目更新
- `mgit push-all` - 推送所有项目变更
- `mgit list` - 列出所有项目
- `mgit delete` - 删除项目
- `mgit set` - 配置设置
- `mgit update` - 更新 MGit 自身
- `mgit help` - 显示帮助信息

### 安装

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

### 配置说明

工具使用 .env 文件存储配置：

```env
# 机器标识
MACHINE_ID=machine-01

# 应用路径（存放所有项目的父目录）
APP_PATH=/path/to/projects

# 应用安装路径（自动设置，不要手动修改）
MGIT_HOME=/path/to/mgit
```

### 使用方法

#### 初始化项目
```bash
# 初始化新项目
./mgit init project_name https://github.com/user/repo.git

# 初始化或更新工具自身
./mgit init mgit
```

#### 拉取和推送代码
```bash
# 拉取/推送单个项目
./mgit pull project_name
./mgit push project_name

# 拉取/推送所有项目
./mgit pull-all
./mgit push-all
```

## English

MGit is a command-line tool for managing multiple Git projects. It helps you efficiently manage, synchronize, and update multiple Git repositories.

### Features
- Initialize and manage multiple Git projects
- Automatic code synchronization for all projects
- Track last commit machine and time
- Version control support for the tool itself
- Flexible project path configuration

### Commands
- `mgit init [project_name]` - Initialize a new project
- `mgit pull [project_name]` - Pull updates for a project
- `mgit push [project_name]` - Push changes for a project
- `mgit pull-all` - Pull updates for all projects
- `mgit push-all` - Push changes for all projects
- `mgit list` - List all projects
- `mgit delete` - Delete a project
- `mgit set` - Configure settings
- `mgit update` - Update MGit itself
- `mgit help` - Show help information

### Installation

1. Clone repository:
```bash
git clone [repository_url]
```

2. Build:
```bash
go build -o mgit
```

3. Initialize the tool:
```bash
# Initialize the tool itself
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

### Configuration

The tool uses a .env file to store configuration:

```env
# Machine identifier
MACHINE_ID=machine-01

# Application path (parent directory for all projects)
APP_PATH=/path/to/projects

# Application installation path (auto-set, do not modify manually)
MGIT_HOME=/path/to/mgit
```

### Usage

#### Initialize Projects
```bash
# Initialize new project
./mgit init project_name https://github.com/user/repo.git

# Initialize or update the tool itself
./mgit init mgit
```

#### Pull and Push Code
```bash
# Pull/push single project
./mgit pull project_name
./mgit push project_name

# Pull/push all projects
./mgit pull-all
./mgit push-all
```

## Japanese

MGitは、複数のGitプロジェクトを管理するためのコマンドラインツールです。複数のGitリポジトリの効率的な管理、同期、更新をサポートします。

### 機能
- 複数のGitプロジェクトの初期化と管理
- すべてのプロジェクトの自動コード同期
- 最終コミットのマシンと時間の追跡
- ツール自体のバージョン管理サポート
- 柔軟なプロジェクトパス設定

### コマンド
- `mgit init [プロジェクト名]` - 新規プロジェクトの初期化
- `mgit pull [プロジェクト名]` - プロジェクトの更新を取得
- `mgit push [プロジェクト名]` - プロジェクトの変更をプッシュ
- `mgit pull-all` - 全プロジェクトの更新を取得
- `mgit push-all` - 全プロジェクトの変更をプッシュ
- `mgit list` - プロジェクト一覧の表示
- `mgit delete` - プロジェクトの削除
- `mgit set` - 設定の構成
- `mgit update` - MGit自体の更新
- `mgit help` - ヘルプ情報の表示

### インストール

1. リポジトリのクローン：
```bash
git clone [repository_url]
```

2. ビルド：
```bash
go build -o mgit
```

3. ツールの初期化：
```bash
# ツール自体の初期化
./mgit init
# または
./mgit init mgit
```

4. 環境設定：
```bash
# マシン識別子の設定
./mgit set machine your-machine-name

# プロジェクト保存パスの設定
./mgit set path /your/custom/path

# 現在の設定を表示
./mgit set
```

### 設定

ツールは.envファイルで設定を保存します：

```env
# マシン識別子
MACHINE_ID=machine-01

# アプリケーションパス（全プロジェクトの親ディレクトリ）
APP_PATH=/path/to/projects

# アプリケーションインストールパス（自動設定、手動で変更しないでください）
MGIT_HOME=/path/to/mgit
```

### 使用方法

#### プロジェクトの初期化
```bash
# 新規プロジェクトの初期化
./mgit init project_name https://github.com/user/repo.git

# ツール自体の初期化または更新
./mgit init mgit
```

#### コードの取得とプッシュ
```bash
# 単一プロジェクトの取得/プッシュ
./mgit pull project_name
./mgit push project_name

# 全プロジェクトの取得/プッシュ
./mgit pull-all
./mgit push-all
```

## Korean

MGit은 여러 Git 프로젝트를 관리하기 위한 명령줄 도구입니다. 여러 Git 저장소의 효율적인 관리, 동기화 및 업데이트를 지원합니다.

### 기능
- 여러 Git 프로젝트 초기화 및 관리
- 모든 프로젝트의 자동 코드 동기화
- 마지막 커밋 머신 및 시간 추적
- 도구 자체의 버전 관리 지원
- 유연한 프로젝트 경로 구성

### 명령어
- `mgit init [프로젝트명]` - 새 프로젝트 초기화
- `mgit pull [프로젝트명]` - 프로젝트 업데이트 가져오기
- `mgit push [프로젝트명]` - 프로젝트 변경사항 푸시
- `mgit pull-all` - 모든 프로젝트 업데이트 가져오기
- `mgit push-all` - 모든 프로젝트 변경사항 푸시
- `mgit list` - 프로젝트 목록 표시
- `mgit delete` - 프로젝트 삭제
- `mgit set` - 설정 구성
- `mgit update` - MGit 자체 업데이트
- `mgit help` - 도움말 표시

### 설치

1. 저장소 복제:
```bash
git clone [repository_url]
```

2. 빌드:
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

4. 환경 설정:
```bash
# 머신 식별자 설정
./mgit set machine your-machine-name

# 프로젝트 저장 경로 설정
./mgit set path /your/custom/path

# 현재 설정 보기
./mgit set
```

### 설정

도구는 .env 파일에 설정을 저장합니다:

```env
# 머신 식별자
MACHINE_ID=machine-01

# 애플리케이션 경로 (모든 프로젝트의 상위 디렉토리)
APP_PATH=/path/to/projects

# 애플리케이션 설치 경로 (자동 설정, 수동으로 수정하지 마세요)
MGIT_HOME=/path/to/mgit
```

### 사용 방법

#### 프로젝트 초기화
```bash
# 새 프로젝트 초기화
./mgit init project_name https://github.com/user/repo.git

# 도구 자체 초기화 또는 업데이트
./mgit init mgit
```

#### 코드 가져오기 및 푸시
```bash
# 단일 프로젝트 가져오기/푸시
./mgit pull project_name
./mgit push project_name

# 모든 프로젝트 가져오기/푸시
./mgit pull-all
./mgit push-all
```

## Traditional Chinese

MGit 是一個用於管理多個 Git 專案的命令列工具。它可以幫助您有效地管理、同步和更新多個 Git 儲存庫。

### 功能特點
- 初始化和管理多個 Git 專案
- 自動同步所有專案的程式碼
- 追蹤最後提交的機器和時間
- 支援管理工具自身的版本控制
- 靈活配置專案存放路徑

### 命令
- `mgit init [專案名稱]` - 初始化新專案
- `mgit pull [專案名稱]` - 拉取專案更新
- `mgit push [專案名稱]` - 推送專案變更
- `mgit pull-all` - 拉取所有專案更新
- `mgit push-all` - 推送所有專案變更
- `mgit list` - 列出所有專案
- `mgit delete` - 刪除專案
- `mgit set` - 配置設定
- `mgit update` - 更新 MGit 本身
- `mgit help` - 顯示說明資訊

### 安裝

1. 克隆儲存庫：
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
# 設定機器標識
./mgit set machine your-machine-name

# 設定專案存放路徑
./mgit set path /your/custom/path

# 查看當前配置
./mgit set
```

### 配置說明

工具使用 .env 檔案儲存配置：

```env
# 機器標識
MACHINE_ID=machine-01

# 應用路徑（存放所有專案的父目錄）
APP_PATH=/path/to/projects

# 應用安裝路徑（自動設定，請勿手動修改）
MGIT_HOME=/path/to/mgit
```

### 使用方法

#### 初始化專案
```bash
# 初始化新專案
./mgit init project_name https://github.com/user/repo.git

# 初始化或更新工具本身
./mgit init mgit
```

#### 拉取和推送程式碼
```bash
# 拉取/推送單個專案
./mgit pull project_name
./mgit push project_name

# 拉取/推送所有專案
./mgit pull-all
./mgit push-all
```

## Français

MGit est un outil en ligne de commande pour gérer plusieurs projets Git. Il vous aide à gérer, synchroniser et mettre à jour efficacement plusieurs dépôts Git.

### Fonctionnalités
- Initialisation et gestion de plusieurs projets Git
- Synchronisation automatique du code pour tous les projets
- Suivi de la dernière machine et heure de commit
- Support de gestion de version pour l'outil lui-même
- Configuration flexible du chemin des projets

### Commandes
- `mgit init [nom_projet]` - Initialiser un nouveau projet
- `mgit pull [nom_projet]` - Récupérer les mises à jour d'un projet
- `mgit push [nom_projet]` - Pousser les modifications d'un projet
- `mgit pull-all` - Récupérer les mises à jour de tous les projets
- `mgit push-all` - Pousser les modifications de tous les projets
- `mgit list` - Lister tous les projets
- `mgit delete` - Supprimer un projet
- `mgit set` - Configurer les paramètres
- `mgit update` - Mettre à jour MGit lui-même
- `mgit help` - Afficher l'aide

### Installation

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

# Définir le chemin de stockage des projets
./mgit set path /your/custom/path

# Voir la configuration actuelle
./mgit set
```

### Configuration

L'outil utilise un fichier .env pour stocker la configuration :

```env
# Identifiant de la machine
MACHINE_ID=machine-01

# Chemin de l'application (répertoire parent pour tous les projets)
APP_PATH=/path/to/projects

# Chemin d'installation de l'application (auto-configuré, ne pas modifier manuellement)
MGIT_HOME=/path/to/mgit
```

### Utilisation

#### Initialiser les projets
```bash
# Initialiser un nouveau projet
./mgit init project_name https://github.com/user/repo.git

# Initialiser ou mettre à jour l'outil lui-même
./mgit init mgit
```

#### Récupérer et pousser le code
```bash
# Récupérer/pousser un projet unique
./mgit pull project_name
./mgit push project_name

# Récupérer/pousser tous les projets
./mgit pull-all
./mgit push-all
```

## 功能特点

- 初始化和管理多个 Git 项目
- 自动同步所有项目的代码
- 跟踪最后提交的机器和时间
- 支持管理工具自身的版本控制
- 灵活配置项目存放路径

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
- 应用自身会标记为 [应用自身]

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
请输入提交说明（留空则使用机器名）: 更新配置文件
# 最终提交信息：更新配置文件 (由 machine-01 推送)

请输入提交说明（留空则使用机器名）: 
# 最终提交信息：由 machine-01 推送
```

交互式菜单示例：
```
选择项目:
  ▸ == 全部项目 == (同步/推送所有项目)
    mgit (https://github.com/user/mgit.git) [最后提交: machine-01 @ 2024-01-01 12:00:00] [应用自身]
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

## 特性说明

1. **工具自更新**：
   - 工具会自动检查其他机器的更新
   - 发现更新时会自动拉取并提示重启

2. **项目路径管理**：
   - 普通项目存放在 `APP_PATH`