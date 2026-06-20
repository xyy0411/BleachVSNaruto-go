# 架构改进计划

## 一、当前架构问题分析

### 1. 系统耦合度高

**问题表现：**
- `BaseCharacter.Update()` 方法承担了太多职责（事件处理、状态转换、连招处理、移动锁定、朝向处理、音频事件）
- `handleMovementLock()` 和 `handleFacing()` 硬编码了大量状态判断
- `Runtime` 结构体混合了物理、动画、音频、连招等不同层次的数据

**具体代码位置：**
- [base.go:41-118](core/charactor/base.go#L41-L118) - Update 方法职责过多
- [base.go:146-157](core/charactor/base.go#L146-L157) - 硬编码的 12 种状态判断
- [runtime.go:13-36](core/charactor/runtime.go#L13-L36) - Runtime 数据层次混乱

### 2. 数据流不清晰

**问题表现：**
- `Intent` 从哪里生成？如何流转？
- `Events` 在 `updateEvents()` 生成，但在多处使用，生命周期不明确
- 状态转换逻辑分散在 `handleComboCommand()` 和 `Update()` 的 switch 中

**具体代码位置：**
- [controller/system.go:30-71](core/controller/system.go#L30-L71) - Intent 生成逻辑
- [base.go:120-130](core/charactor/base.go#L120-L130) - Events 生成
- [base.go:71-96](core/charactor/base.go#L71-L96) - 状态转换逻辑分散

### 3. 扩展性差

**问题表现：**
- 添加新状态需要修改：`handleMovementLock()`、`handleFacing()`、`Update()` 的 switch
- 添加新功能时不知道该在哪里写代码
- 状态判断逻辑重复出现在多个地方

**具体代码位置：**
- [base.go:148-152](core/charactor/base.go#L148-L152) - 状态判断重复
- [physics/system.go:92-95](core/physics/system.go#L92-L95) - 相同的状态判断

---

## 二、改进目标

### 1. 降低耦合度
- 将 `BaseCharacter.Update()` 拆分成多个独立的系统
- 每个系统只负责一个职责
- 系统之间通过明确的接口通信

### 2. 清晰化数据流
- 明确数据从输入到输出的流转路径
- 使用事件驱动架构
- 引入清晰的状态机

### 3. 提高扩展性
- 添加新功能时只需添加新系统或新状态
- 不需要修改现有代码
- 支持配置化扩展

---

## 三、新架构设计

### 1. 系统拆分方案

将 `BaseCharacter.Update()` 拆分成以下独立系统：

```
core/
├── state/              # 新增：状态管理系统
│   ├── machine.go      # 状态机
│   ├── transitions.go  # 状态转换规则
│   └── context.go      # 状态上下文
├── event/              # 新增：事件系统
│   ├── bus.go          # 事件总线
│   └── types.go        # 事件类型定义
├── movement/           # 新增：移动系统
│   ├── system.go       # 移动逻辑
│   └── constraints.go  # 移动约束
├── facing/             # 新增：朝向系统
│   └── system.go       # 朝向逻辑
├── audio/              # 已存在：音频系统
│   └── system.go       # 音频逻辑
├── controller/         # 已存在：输入控制
│   └── system.go       # 输入转换
├── physics/            # 已存在：物理系统
│   └── system.go       # 物理模拟
└── charactor/          # 已存在：角色系统
    ├── base.go         # 简化为数据容器
    ├── runtime.go      # 简化数据结构
    └── char.go          # 角色接口
```

### 2. 数据流设计

```
输入层 → 控制层 → 状态层 → 物理层 → 渲染层
  ↓        ↓        ↓        ↓        ↓
Input → Controller → State → Physics → Animation
           ↓          ↓        ↓
         Intent    Events   Body
                      ↓
                   Event Bus
                      ↓
              [Movement, Facing, Audio]
```

**数据流转路径：**
1. **输入层**：`input.System` 读取键盘输入 → `InputFrame`
2. **控制层**：`controller.System` 转换输入 → `Intent`
3. **状态层**：`state.Machine` 根据意图和当前状态 → 新状态 + `Events`
4. **物理层**：`physics.System` 根据状态和意图 → 物理位置
5. **事件层**：`event.Bus` 分发事件 → 各子系统响应
6. **子系统**：
   - `movement.System` 处理移动约束
   - `facing.System` 处理朝向
   - `audio.System` 处理音效

### 3. 状态机设计

```go
// state/machine.go
type Machine struct {
    current state.State
    rules   TransitionRules
}

// state/transitions.go
type TransitionRules map[state.State][]Transition

type Transition struct {
    Condition func(ctx Context) bool
    To        state.State
    Event     event.Type  // 触发的事件
}

// 示例：状态转换规则配置
var defaultRules = TransitionRules{
    state.Idle: {
        {Condition: isJumping, To: state.JumpStart, Event: event.JumpStart},
        {Condition: isDashing, To: state.Dash, Event: event.Dash},
        {Condition: isMoving, To: state.Run, Event: event.None},
    },
    state.JumpStart: {
        {Condition: animationFinished, To: state.Jump, Event: event.None},
    },
    // ...
}
```

### 4. 事件系统设计

```go
// event/bus.go
type Bus struct {
    subscribers map[Type][]Handler
}

type Handler func(Event)

type Event struct {
    Type      Type
    Character string
    Data      interface{}
}

// 示例：事件订阅
bus.Subscribe(event.JumpStart, func(e Event) {
    audioSystem.Play(e.Character, "jump_start")
})
```

---

## 四、具体实施步骤

### 阶段一：引入事件系统（1-2天）

**目标：** 建立事件驱动基础

**步骤：**
1. 创建 `core/event/` 包
2. 实现 `Bus` 和事件类型定义
3. 将 `BaseCharacter` 中的音频事件改为通过事件总线发布
4. 创建 `audio.System` 订阅音频事件

**验证：** 音频功能正常工作，代码更清晰

### 阶段二：引入状态机（2-3天）

**目标：** 集中管理状态转换逻辑

**步骤：**
1. 创建 `core/state/` 包
2. 实现状态机和转换规则
3. 将 `BaseCharacter.Update()` 中的状态转换逻辑迁移到状态机
4. 状态转换时发布事件

**验证：** 状态转换正常，逻辑集中

### 阶段三：拆分移动和朝向系统（1-2天）

**目标：** 分离关注点

**步骤：**
1. 创建 `core/movement/` 包
2. 将 `handleMovementLock()` 迁移到 `movement.System`
3. 创建 `core/facing/` 包
4. 将 `handleFacing()` 迁移到 `facing.System`
5. 这些系统订阅状态变更事件

**验证：** 移动和朝向功能正常

### 阶段四：简化 BaseCharacter（1天）

**目标：** 让 BaseCharacter 只做协调

**步骤：**
1. `BaseCharacter` 只保留数据容器和系统协调
2. `Update()` 方法变为调用各系统
3. 移除所有业务逻辑

**验证：** 所有功能正常，代码更清晰

### 阶段五：优化数据结构（1天）

**目标：** 清晰化数据层次

**步骤：**
1. 将 `Runtime` 拆分成多个专注的数据结构
2. 引入 `StateData`、`PhysicsData`、`AnimationData` 等
3. 更新所有系统使用新的数据结构

**验证：** 数据流更清晰

---

## 五、示例代码

### 1. 事件系统示例

```go
// core/event/types.go
package event

type Type int

const (
    None Type = iota
    JumpStart
    JustLanded
    Dash
    RunStep
)

type Event struct {
    Type      Type
    Character string
    Data      interface{}
}

// core/event/bus.go
package event

type Handler func(Event)

type Bus struct {
    subscribers map[Type][]Handler
}

func NewBus() *Bus {
    return &Bus{
        subscribers: make(map[Type][]Handler),
    }
}

func (b *Bus) Subscribe(eventType Type, handler Handler) {
    b.subscribers[eventType] = append(b.subscribers[eventType], handler)
}

func (b *Bus) Publish(event Event) {
    for _, handler := range b.subscribers[event.Type] {
        handler(event)
    }
}
```

### 2. 状态机示例

```go
// core/state/machine.go
package state

import (
    "github.com/xyy0411/bleachVSnaruto/common/state"
    "github.com/xyy0411/bleachVSnaruto/core/event"
)

type Context struct {
    Intent    models.Intent
    Body      *models.PhysicsBody
    PrevState state.State
    Events    *charactor.Events
}

type Machine struct {
    current state.State
    rules   TransitionRules
    bus     *event.Bus
}

func NewMachine(rules TransitionRules, bus *event.Bus) *Machine {
    return &Machine{
        current: state.Idle,
        rules:   rules,
        bus:     bus,
    }
}

func (m *Machine) Update(ctx Context) state.State {
    transitions, exists := m.rules[m.current]
    if !exists {
        return m.current
    }

    for _, transition := range transitions {
        if transition.Condition(ctx) {
            m.current = transition.To
            if transition.Event != event.None {
                m.bus.Publish(event.Event{
                    Type:  transition.Event,
                    Data:  ctx,
                })
            }
            return m.current
        }
    }

    return m.current
}
```

### 3. 简化后的 BaseCharacter

```go
// core/charactor/base.go
package charactor

type BaseCharacter struct {
    id       string
    name     string
    Runtime  *Runtime
    Data     *Data

    // 系统
    stateMachine *state.Machine
    eventBus     *event.Bus
}

func NewBaseCharacter(id, name string, runtime *Runtime, data *Data) *BaseCharacter {
    eventBus := event.NewBus()

    return &BaseCharacter{
        id:           id,
        name:         name,
        Runtime:      runtime,
        Data:         data,
        stateMachine: state.NewMachine(state.DefaultRules(), eventBus),
        eventBus:     eventBus,
    }
}

func (b *BaseCharacter) Update() {
    // 更新事件
    b.updateEvents()

    // 状态转换
    ctx := state.Context{
        Intent:    b.Runtime.Intent,
        Body:      b.Runtime.Body,
        PrevState: b.Runtime.State,
        Events:    &b.Runtime.Events,
    }
    b.Runtime.State = b.stateMachine.Update(ctx)

    // 更新动画
    b.Runtime.AnimPlayer.Play(b.Data.Animations.ByState[b.Runtime.State])

    // 保存前一帧状态
    b.savePrevState()
}
```

---

## 六、扩展性示例

### 添加新状态（如 "受伤"）

**旧架构：** 需要修改 3-5 个地方
**新架构：** 只需修改 1 个地方

```go
// 只需在状态转换规则中添加
state.Idle: {
    // ... 现有规则
    {Condition: isHurt, To: state.Hurt, Event: event.Hurt},
}

// 添加事件处理
bus.Subscribe(event.Hurt, func(e Event) {
    audioSystem.Play(e.Character, "hurt")
})
```

### 添加新系统（如 "特效系统"）

**旧架构：** 需要修改 BaseCharacter.Update()
**新架构：** 只需添加新系统并订阅事件

```go
// core/effect/system.go
package effect

type System struct {
    bus *event.Bus
}

func NewSystem(bus *event.Bus) *System {
    s := &System{bus: bus}

    // 订阅感兴趣的事件
    bus.Subscribe(event.JumpStart, s.onJumpStart)
    bus.Subscribe(event.Dash, s.onDash)

    return s
}

func (s *System) onJumpStart(e event.Event) {
    // 播放跳跃特效
    s.playEffect(e.Character, "jump_dust")
}
```

---

## 七、迁移策略

### 兼容性保证

1. **渐进式迁移**：每个阶段都能独立工作
2. **功能验证**：每个阶段完成后进行完整测试
3. **回滚机制**：每个阶段都可以独立回滚

### 测试策略

1. **单元测试**：为每个新系统编写单元测试
2. **集成测试**：验证系统间的协作
3. **回归测试**：确保现有功能不受影响

---

## 八、预期收益

### 1. 降低耦合度
- 每个系统职责单一，易于理解和修改
- 系统间通过事件通信，依赖关系清晰

### 2. 清晰化数据流
- 输入 → 意图 → 状态 → 物理 → 渲染，路径清晰
- 事件总线让数据流向可视化

### 3. 提高扩展性
- 添加新状态：修改状态转换规则
- 添加新系统：订阅相关事件
- 添加新功能：添加新系统或新事件处理器

### 4. 提高可测试性
- 每个系统可以独立测试
- 可以模拟事件进行测试

---

## 九、风险和注意事项

### 1. 性能考虑
- 事件总线可能引入轻微性能开销
- 状态机查表比 switch 稍慢
- **缓解措施**：使用事件池、优化状态机实现

### 2. 学习曲线
- 团队需要理解新的架构
- **缓解措施**：提供详细文档和示例

### 3. 迁移成本
- 需要重写部分代码
- **缓解措施**：渐进式迁移，保证每个阶段可用

---

## 十、总结

这个改进计划通过引入事件驱动架构和状态机模式，将现有的单体 `BaseCharacter.Update()` 拆分成多个独立的系统，从而：

1. **降低耦合度**：每个系统职责单一
2. **清晰化数据流**：通过事件总线明确数据流向
3. **提高扩展性**：添加新功能只需添加新系统或新状态

建议按照阶段逐步实施，每个阶段完成后进行验证，确保功能正常。