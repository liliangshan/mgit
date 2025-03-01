# MGit - Git 项目管理工具

[简体中文](#简体中文) | [English](#english) | [日本語](#japanese) | [한국어](#korean) | [繁體中文](#traditional-chinese) | [Français](#français)

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

当前版本：1.0.3
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

## English

MGit is a command-line tool for managing multiple Git projects. It helps you efficiently manage, synchronize, and update multiple Git repositories.

### Features
- Initialize and manage multiple Git projects
- Automatic code synchronization for all projects
- Track last commit machine and time
- Version control support for the tool itself
- Flexible project path configuration
- Multi-language support
- Remote database synchronization
- System PATH integration

### Remote Database Synchronization

MGit supports database synchronization through remote Git repositories for multi-device configuration sharing:

1. Enable database sync:
```bash
./mgit set
# Select "Enable database repository sync"
# Enter database repository address
```

2. Sync mechanism:
- Automatic database sync after each push
- Latest configuration fetch before each pull
- Multi-user collaboration with automatic merging

3. Synced data:
- Project configurations
- Branch settings
- Last commit records
- Device identifiers

### System PATH Configuration

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

After configuration, use `mgit` command from any directory.

### Version Information

Check current version:
```bash
mgit -v
# or
mgit --version
```

Current version: 1.0.3
Project homepage: https://github.com/liliangshan/mgit

### Why Choose MGit?

1. Efficient Collaboration
- Multi-device configuration sync
- Unified team project management
- Time-saving batch operations

2. User-Friendly
- Interactive command-line interface
- Multi-language localization
- Smart prompts and help

3. Secure and Reliable
- Automatic configuration backup
- Version control protection
- Smart conflict handling

4. Highly Extensible
- Custom configuration support
- Flexible project management
- Continuous updates

## Japanese

MGitは、複数のGitプロジェクトを管理するためのコマンドラインツールです。複数のGitリポジトリの効率的な管理、同期、更新をサポートします。

### 主な特徴
- 複数のGitプロジェクトの初期化と管理
- すべてのプロジェクトの自動コード同期
- 最終コミットのマシンと時間の追跡
- ツール自体のバージョン管理サポート
- 柔軟なプロジェクトパス設定
- 多言語サポート
- リモートデータベース同期
- システムPATH統合

### リモートデータベース同期

MGitはリモートGitリポジトリを通じてデータベースを同期し、複数デバイス間の設定共有を実現します：

1. データベース同期の有効化：
```bash
./mgit set
# 「データベースリポジトリ同期を有効にする」を選択
# データベースリポジトリアドレスを入力
```

2. 同期メカニズム：
- プッシュ後に自動的にデータベースを同期
- プル前に最新の設定を取得
- 複数ユーザーの協力と自動マージをサポート

3. 同期データ：
- プロジェクト設定
- ブランチ設定
- 最終コミット記録
- デバイス識別子

### システムPATH設定

MGitをシステムPATHに追加してグローバルアクセスを実現：

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

設定後、任意のディレクトリで`mgit`コマンドを使用できます。

### バージョン情報

現在のバージョンを確認：
```bash
mgit -v
# または
mgit --version
```

現在のバージョン：1.0.3
プロジェクトホームページ：https://github.com/liliangshan/mgit

### なぜMGitを選ぶのか？

1. 効率的な協力
- 複数デバイスの設定同期
- チームプロジェクトの統一管理
- 時間節約のバッチ操作

2. ユーザーフレンドリー
- 対話式コマンドラインインターフェース
- 多言語ローカライゼーション
- スマートなプロンプトとヘルプ

3. 安全で信頼性が高い
- 自動設定バックアップ
- バージョン管理による保護
- スマートな競合処理

4. 高い拡張性
- カスタム設定のサポート
- 柔軟なプロジェクト管理
- 継続的な更新

### インストール方法

1. リポジトリのクローン：
```bash
git clone [repository_url]
```

2. ビルド：
```bash
go build -o mgit
```

3. 初期化：
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

## Traditional Chinese

MGit 是一個用於管理多個 Git 專案的命令列工具。它支援多個 Git 倉庫的高效管理、同步和更新。

### 主要特點
- 多個 Git 專案的初始化和管理
- 所有專案的自動程式碼同步
- 最後提交的機器和時間追蹤
- 工具本身的版本管理支援
- 靈活的專案路徑配置
- 多語言支援
- 遠端資料庫同步
- 系統 PATH 整合

### 遠端資料庫同步

MGit 通過遠端 Git 倉庫同步資料庫，實現多設備間的設定共享：

1. 啟用資料庫同步：
```bash
./mgit set
# 選擇「啟用資料庫倉庫同步」
# 輸入資料庫倉庫地址
```

2. 同步機制：
- 推送後自動同步資料庫
- 拉取前獲取最新設定
- 支援多用戶協作和自動合併

3. 同步資料：
- 專案設定
- 分支設定
- 最後提交記錄
- 設備識別碼

### 系統 PATH 設定

將 MGit 添加到系統 PATH 以實現全域存取：

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

設定後，可在任意目錄使用 `mgit` 命令。

### 版本資訊

查看當前版本：
```bash
mgit -v
# 或
mgit --version
```

當前版本：1.0.3
專案主頁：https://github.com/liliangshan/mgit

### 為什麼選擇 MGit？

1. 高效協作
- 多設備設定同步
- 團隊專案統一管理
- 批次操作節省時間

2. 使用者友好
- 互動式命令列介面
- 多語言本地化支援
- 智慧提示和幫助

3. 安全可靠
- 自動備份設定
- 版本控制保護
- 智慧衝突處理

4. 擴展性強
- 支援自定義設定
- 靈活的專案管理
- 持續更新迭代

### 安裝方法

1. 克隆倉庫：
```bash
git clone [repository_url]
```

2. 建置：
```bash
go build -o mgit
```

3. 初始化：
```bash
# 工具本身初始化
./mgit init
# 或
./mgit init mgit
```

4. 環境設定：
```bash
# 設定機器識別碼
./mgit set machine your-machine-name

# 設定專案儲存路徑
./mgit set path /your/custom/path

# 顯示當前設定
./mgit set
```

## Français

MGit est un outil en ligne de commande pour gérer plusieurs projets Git. Il prend en charge la gestion, la synchronisation et la mise à jour efficaces de plusieurs dépôts Git.

### Caractéristiques principales
- Initialisation et gestion de plusieurs projets Git
- Synchronisation automatique du code pour tous les projets
- Suivi de la machine et du temps du dernier commit
- Support de gestion de version pour l'outil lui-même
- Configuration flexible du chemin des projets
- Support multilingue
- Synchronisation de base de données distante
- Intégration du PATH système

### Synchronisation de base de données distante

MGit synchronise la base de données via un dépôt Git distant, permettant le partage des configurations entre plusieurs appareils :

1. Activer la synchronisation de la base de données :
```bash
./mgit set
# Sélectionner "Activer la synchronisation du dépôt de base de données"
# Saisir l'adresse du dépôt de base de données
```

2. Mécanisme de synchronisation :
- Synchronisation automatique de la base de données après push
- Récupération des derniers paramètres avant pull
- Support de la collaboration multi-utilisateurs et fusion automatique

3. Données synchronisées :
- Configurations des projets
- Paramètres des branches
- Enregistrements des derniers commits
- Identifiants des appareils

### Configuration du PATH système

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

Après configuration, la commande `mgit` peut être utilisée dans n'importe quel répertoire.

### Informations de version

Vérifier la version actuelle :
```bash
mgit -v
# ou
mgit --version
```

Version actuelle : 1.0.3
Page du projet : https://github.com/liliangshan/mgit

### Pourquoi choisir MGit ?

1. Collaboration efficace
- Synchronisation des configurations multi-appareils
- Gestion unifiée des projets d'équipe
- Opérations par lots pour gagner du temps

2. Convivial
- Interface en ligne de commande interactive
- Support de localisation multilingue
- Invites et aide intelligentes

3. Sûr et fiable
- Sauvegarde automatique des configurations
- Protection par contrôle de version
- Gestion intelligente des conflits

4. Haute extensibilité
- Support des configurations personnalisées
- Gestion flexible des projets
- Mises à jour continues

### Méthode d'installation

1. Cloner le dépôt :
```bash
git clone [repository_url]
```

2. Compiler :
```bash
go build -o mgit
```

3. Initialisation :
```bash
# Initialisation de l'outil lui-même
./mgit init
# ou
./mgit init mgit
```

4. Configuration de l'environnement :
```bash
# Définir l'identifiant de la machine
./mgit set machine your-machine-name

# Définir le chemin de stockage des projets
./mgit set path /your/custom/path

# Afficher la configuration actuelle
./mgit set
```

