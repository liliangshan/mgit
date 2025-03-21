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

### 代理設定
```bash
mgit proxy
```
設定 Git 代理伺服器，支援 HTTP 和 HTTPS 代理。執行命令後，系統會提示：
1. 是否使用代理
2. 選擇代理類型（HTTP/HTTPS）
3. 輸入代理伺服器 IP 位址
4. 輸入代理伺服器端口

範例：
```bash
$ mgit proxy
是否使用代理？(y/N): y
請選擇代理類型：
1) HTTP
2) HTTPS
請選擇: 1
請輸入代理IP: 127.0.0.1
請輸入代理端口: 7890
代理設定成功
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