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

### コマンド説明

#### プロキシ設定
```bash
mgit proxy
```
Git プロキシサーバーを設定します。HTTP と HTTPS プロキシをサポートしています。コマンドを実行すると、システムが以下のプロンプトを表示します：
1. プロキシを使用するかどうか
2. プロキシタイプの選択（HTTP/HTTPS）
3. プロキシサーバーの IP アドレスの入力
4. プロキシサーバーのポート番号の入力

例：
```bash
$ mgit proxy
プロキシを使用しますか？(y/N): y
プロキシタイプを選択してください：
1) HTTP
2) HTTPS
選択してください: 1
プロキシIPを入力してください: 127.0.0.1
プロキシポートを入力してください: 7890
プロキシ設定が完了しました
```