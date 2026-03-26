# AI Code Review Rules

This repository is written in Go.

When reviewing pull requests, follow the rules below.

---

# Documentation

All exported identifiers must have documentation comments.

Rules:

1. Every exported function, type, method, or variable must have a comment.
2. Comments must start with the identifier name (Go documentation style).
3. Preferred language for new comments is Chinese.
4. If a contributor already wrote comments in English, DO NOT request translation.
5. Comments should clearly explain what the function does.
6. If helpful, briefly describe parameters and return values.

Example:

// RegisterCharacter 注册一个新的角色到角色注册表中。
func RegisterCharacter(name string) error

Avoid comments like:

// 注册角色
func RegisterCharacter(...)

because Go documentation comments should start with the identifier name.

---

# Code Style

Follow idiomatic Go style.

General principles:

- Prefer simple and readable code
- Avoid unnecessary abstraction
- Avoid overly complex logic
- Keep functions focused on a single responsibility
- Prefer clarity to cleverness

Prefer this style:

if err != nil {
return err
}

Avoid deeply nested logic where possible.

---

# Error Handling

Errors should always be handled explicitly.

Rules:

- Do not ignore returned errors
- Do not use empty error checks
- Prefer returning errors to hiding them
- Error messages should be meaningful

Avoid patterns like:

if err != nil {
}

---

# Naming

Names should clearly describe their purpose.

Avoid vague names such as:

DoThing
HandleStuff
ProcessData
Manager2

Prefer descriptive names such as:

LoadCharacterData
RegisterPlayer
HandlePlayerInput

---

# Code Quality

When reviewing code, prioritize:

1. Missing documentation comments
2. Ignored errors
3. Potential bugs
4. Confusing logic
5. Poor naming
6. Unnecessary complexity

If possible, suggest clearer implementations.

---

# Performance

Avoid unnecessary allocations.

Prefer simple data structures.

Do not prematurely optimize, but point out obvious inefficiencies.

---

# Concurrency (if present)

When goroutines or shared data are used, check for:

- race conditions
- unsafe shared state
- missing synchronization

Suggest safer patterns when needed.

---

# Suggestions

When problems are found:

- Explain the issue clearly
- Suggest a better implementation
- Provide example code if helpful

Keep suggestions concise and actionable.

If the code already follows repository rules and no issues are found,
simply react with 👍.