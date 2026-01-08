# CLI tool to run safely some npm, yarn & git commands

### Problem
On MacOS it happens often that **VSCode** editor loses permission because it's recognized as dangerous by OS and it can't use different npm/yarn/git commands without *chmod/chown*.

### Supported commands:
1. **npm install**
2. **npm i "package-name"**
3. **yarn**
4. **yarn add "package-name"**
5. **npm build**
6. **git commit -m "message"**
7. **git pull**
8. **git stash push**
9. **git stash push -m "message"**

### Supported OS for this program:
MacOS, Ubuntu.

### Usage:
1. Clone repository https://github.com/PseudoElement/cmd-executor.git
2. In terminal run **$ ./goexec --path=absolute/path/to/your/project**(example **$ ./goexec --path=/usr/desktop/my-app**)
**Note**: If you run git commands in **--path** argument you need to specify absolute path to directory where .git directory located, usually it's the same as root directory of your project **--path=/usr/desktop/my-app**.

Also you can create your own .env file, specify **PATH_TO_NPM_PROJECT** variable like *PATH_TO_NPM_PROJECT=/usr/desktop/my-app* and run executable file without need to pass **--path** variable.

#### Hint:
For better experience you can move executable in */usr/local/bin* directory and run cli command by executable name **goexec** from any directory without specifying absolute path to directory when it's placed.

### Why only 9 commands supported?
I wrote this tool for personal usage, cause I encountered problems with permission only with that commands. 
It's an open source project and if you want to extend a list of available command - send pull request to this repo, after check from my side your pr will be pushed.
