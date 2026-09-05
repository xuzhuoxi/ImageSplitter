# ReleaseNote 模板

生成或补充 `notes/release/ReleaseNotes_{tag}.md` 前阅读本文件。节名与 `notes/release/ReleaseNotes.md` 一致；文风与现有 `notes/release/ReleaseNotes_v*.md` 一致。

开头写一两句亮点总述，再按 Improvements / API Changes / Changes / Fixes 分节，有依赖变更时再写 Library Changes。

## 新文件结构

```markdown
## Release Notes

+ 一两句总述（自上一版本到本 tag 的要点）。

### Known Issues

### Improvements

### API Changes

### Changes

### Fixes

## Library Changes

### library Updated

#### Updated

#### No Longer Available

#### Added
```

## 填写规则

- 不要写 `# ImageSplitter` 大标题，文档从 `## Release Notes` 起头
- 标题用模板英文；条目用中文，`+ ` 开头
- 模板中部分条目没有内容时，生成的文档中可以不包含该条目：无条目的 `###` / `####` 整节不要写出；`## Library Changes` 下没有任何库变更时整节不要写出。不要留空标题
- 总述写本版本最重要的一两件事，不要把下面条目再抄一遍
- **Improvements**：新能力、参数增强、脚本、文档完善
- **API Changes**：公开参数或接口的增删改
- **Changes**：对外行为、配置、默认策略等非修复变更
- **Fixes**：缺陷修复
- **Known Issues**：明确未解决的限制；无则省略
- `Library Changes`：只写 `go.mod` 里实际变化的模块与版本（`旧 → 新`）
- `### library Updated` 中的 `library` 换成真实模块名，或不用该占位、直接在 `## Library Changes` 下写 `+ 模块：旧 → 新`
