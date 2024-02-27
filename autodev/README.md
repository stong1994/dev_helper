快速提交、部署脚本
简介
开发过程中，需要频繁切换到开发环境合并代码并部署：
1. 切换分支
2. 拉取代码
3. 合并代码
4. 推送代码
5. 部署代码
6. 切回分支
   这个过程可以通过脚本来执行。
   功能
1. 在当前开发分支
   shell脚本
   脚本可放置在全局可执行目录。
   完整命令：`deploy.sh [branch] [version]`
   version未指定下的默认规则：
- 目标分支是dev，则默认值为1.0.0
- 目标分支是test，则默认值为2.0.0
- 目标分支是master，则默认值为3.0.0
  branch未指定，默认为dev