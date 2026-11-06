# ConsumedLen race 包装一致性修复

- 仅在 nocopy_linkbuffer_race.go 为 SafeLinkBuffer.ConsumedLen 增加与 MarkRead/ConsumedSlices 相同的 Mutex 包装。
- 不改非 race 实现、节点生命周期或读写算法。此修复补齐包装一致性，不声称已经复现 reactor 与合法单 reader 之间的数据竞争。
- 验证通过：go build ./...；go test -race -count=1 -run='Read|Copy|Release|Slice|LinkBuffer|Meshlib' .。日志在共同工作空间 .cache/meshlib-confirmed-fixes-20261106/netpoll-*。
