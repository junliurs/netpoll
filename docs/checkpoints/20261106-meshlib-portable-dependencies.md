# meshlib netpoll 固定远端依赖

范围仅 go.mod/go.sum：删除八条本地路径 replace，保留 gopkg 的固定远端 fork 映射。不改产品或测试源码，不升级无关依赖。

保留远端源码起点 7771cff887475877e19a13a2195b17e46af0d00a 和本地依赖适配起点 1728921575054ed6d917cb76bf3f4a4edd637216。仍在原 worktree、test/meshlib-local-integration-20261105 分支提交。

实际 gopkg 来源：github.com/junliurs/cloudwego-gopkg v0.0.0-20261105192813-e9461b4bc999，提交 e9461b4bc9990522821afc9f462add46ee59ac19；对应提交已在 GitHub 远端可获取，未用分支浮动版本。

Go 1.26.1、GOWORK=off、mod=readonly 下 go mod download、go build -p=8 ./...、测试代码 test-link 均通过。test-link 不代表执行测试。缓存和临时文件在工作区，GOGC100、不设 GOMEMLIMIT/GODEBUG。

凭证仅在 Git 子进程中临时复用原仓库 credential helper，不改原仓库或全局设置。下一步上层使用本提交的规范模块版本。
