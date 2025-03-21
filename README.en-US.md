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
