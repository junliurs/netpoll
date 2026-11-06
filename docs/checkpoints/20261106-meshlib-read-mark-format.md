# meshlib ReadMark 格式检查

仅对 5ff953d 新增 readMarked 字段的注释执行 gofmt 对齐，生产文件只有一行空白差异，没有逻辑变化。保留同一 worktree 和既定三项修复范围。之前的针对性普通/race、读取回归和空模块缓存构建结论仍适用；最终 benchmark 固定本提交。
