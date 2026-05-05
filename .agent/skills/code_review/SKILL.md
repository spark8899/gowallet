---
name: Code Review
description: A comprehensive guide and checklist for performing code reviews, with a focus on Go best practices and security.
---

# Code Review Skill

This skill guides you through performing a thorough code review. It is designed to ensure code quality, security (especially for cryptocurrency applications), and maintainability.

## 1. Preparation & Context
- **Understand the Objective**: Before looking at code, clarify *what* is being solved. Read the issue description or task prompt.
- **Scope Analysis**: Identify which files are modified and how they interact with the rest of the system.
- **Dependencies**: Check if `go.mod` or `go.sum` has changed.

## 2. Functional Correctness
- **Logic Verification**: Trace the execution path. Does the logic hold?
- **Edge Cases**:
  - Empty inputs/slices/maps.
  - Zero values for structs/primitives.
  - Boundary conditions (min/max values).
- **Error Handling**:
  - Are errors checked? (Never use `_` to ignore errors without a strong reason).
  - Are errors propagated with context (e.g., `fmt.Errorf("failed to x: %w", err)`)?

## 3. Security (High Priority for `gowallet`)
- **Memory Safety**:
  - **CRITICAL**: Sensitive data (private keys, mnemonics, seeds) **MUST** be zeroed out after use. Look for `defer security.ZeroBytes(data)` or similar calls.
- **Input Validation**:
  - Check lengths, formats, and ranges for all user inputs.
  - Validate derivation paths (BIP32/44 compliance).
- **Randomness**:
  - Ensure `crypto/rand` is used for key generation, NEVER `math/rand`.

## 4. Code Quality & Go Idioms
- **Style**:
  - Adhere to `duber/uber-go-style` or standard Go conventions.
  - Names should be short but descriptive (`ctx`, `err` are fine; `data` is vague).
- **Concurrency**:
  - Check for race conditions.
  - Ensure channels are closed properly.
  - Verify strict usage of `sync.Mutex` or `sync.RWMutex`.
- **Maintainability**:
  - Avoid large functions. If a function is >50 lines, can it be split?
  - Avoid deep nesting (guard clauses preferred).

## 5. Performance
- **Allocations**: Be mindful of allocations in hot loops.
- **Resources**: Are file handles, connections, and iterators closed? (Check `defer`).

## 6. Testing
- **Coverage**: Do the new tests cover the changes?
- **Scenarios**: Are failure modes tested?
- **Integration**: Does it break existing `cmd` logic?

## 7. Review Output Format
When generating a review, structure it as follows:

1.  **Summary**: High-level impression (LGTM / Request Changes).
2.  **Critical Issues**: Security risks or bugs that *must* be fixed.
3.  **Suggestions**: Improvements that are nice to have but not blocking.
4.  **Questions**: clarifications needed.

---
**Tip**: If you are unsure about a specific Go behavior, suggest writing a small test case to verify.
