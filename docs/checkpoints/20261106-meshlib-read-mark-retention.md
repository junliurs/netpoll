# meshlib ReadMark 在 Release 前保留节点

范围仅修复 review 第三项：MarkRead 后通过 connection.Read（内部 readCopy）跨节点读取，会在 Release 前丢失标记节点。继续在本仓库原 worktree 修改。

增加一个 reader 状态位：MarkRead 设置，Release 清除；存在标记时 readCopy 暂停提前回收，沿途节点统一到 Release 回收。无标记的复制读取仍按原逻辑释放，不增加独立缓存，不改接口、线程、GC 或 Kitex/SDK 源码。

回归覆盖多个标记、多次复制读取跨中间节点、字节计数与完整借用数据、Release 清理，以及未标记读取的原回收行为。针对性普通/race 测试通过，已有 Read/Copy/Release/Slice/LinkBuffer 测试及 go build ./... 通过。原始 12 字节读取后误报 0 的 review 复现位于共同工作空间 .cache/meshlib-review-20261106/，本次验证日志 .cache/meshlib-review-fixes-20261106/。

原 Header map 复用、Native context 池化和 client 多 key 缓存保留。本提交没有本地路径 replace，后续只按依赖顺序更新引用本提交的固定版本并验证组合，不追加其他功能。
