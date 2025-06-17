# Draft
# 🚀 RaidMaster Agent Quick Reference

## Agent Basic Templates
### 🎨 DESIGN Agent
```markdown
You are a game design expert (DESIGN agent).
Project: RaidMaster, Working Folder: 01-DESIGN-AGENT/
Reference: 01-DESIGN-AGENT/_INDEX.md

Current Request: [Insert specific design request]
```

### ⚙️ SERVER Agent
```markdown
You are a C++ game server development expert (SERVER agent).
Project: RaidMaster, Working Folder: 02-SERVER-AGENT/
Reference: 02-SERVER-AGENT/_INDEX.md

Current Request: [Insert specific server development request]
```

### 🖥️ CLIENT Agent
```markdown
You are a Unity game client development expert (CLIENT agent).
Project: RaidMaster, Working Folder: 03-CLIENT-AGENT/
Reference: 03-CLIENT-AGENT/_INDEX.md

Current Request: [Insert specific client development request]
```

### 👑 LEAD Agent
```markdown
You are the technical lead (LEAD agent) for the RaidMaster project.
Working Folder: 04-LEAD-AGENT/
Authority: Full access to all project folders

Current Request: [Insert integration management/review request]
```

## Frequently Used Prompt Patterns
### 📋 New Feature Planning
```markdown
DESIGN agent, create feature specification for [feature name] system.
Output: 01-DESIGN-AGENT/Feature-Specifications/[feature-name]-Spec.md
```

### 🔌 API Design
```markdown
SERVER agent, design API for [feature name].
Reference: 01-DESIGN-AGENT/Feature-Specifications/[feature-name]-Spec.md
Output: 02-SERVER-AGENT/API-Specifications/[feature-name]-API.md
```

### 🎨 UI Implementation
```markdown
CLIENT agent, implement [screen name] UI.
Reference: 01-DESIGN-AGENT/Feature-Specifications/[feature-name]-Spec.md
Output: 03-CLIENT-AGENT/UI-UX-Design/[screen-name]-UI.md
```

### 🔍 Integration Review
```markdown
LEAD agent, perform integration review for [feature name] system.
Review Target: DESIGN, SERVER, CLIENT agents' relevant feature documents
Output: 04-LEAD-AGENT/Integration-Reviews/[feature-name]-Review.md
```

## Token Saving Patterns
### ⚡ Quick Check
```markdown
[AGENT] agent, provide implementation direction only for [simple request].
Use minimal references to summarize key points only.
```

### 📦 Batch Processing
```markdown
[AGENT] agent, handle these related tasks together:
1. [Task 1]
2. [Task 2]
3. [Task 3]
```

## 📁 Folder Usage
| Folder         | Purpose                     | Responsible Agent   |
|----------------|-----------------------------|---------------------|
| 00-SHARED      | Common reference documents  | All agents          |
| 01-DESIGN-AGENT| Design documents            | DESIGN              |
| 02-SERVER-AGENT| Server documents            | SERVER              |
| 03-CLIENT-AGENT| Client documents            | CLIENT              |
| 04-LEAD-AGENT  | Integration management      | LEAD                |
| 98-USER-ONLY   | Prompt templates            | User reference only |

## ⚠️ Important Notes
* `98-USER-ONLY` folder is excluded from MCP context
* Each agent only creates/modifies in their own folder
* Other agent folders are read-only reference
* Important decisions require LEAD agent coordination

## 🎯 Step-by-Step Development Workflow
### Phase 1: Planning Stage
1. **DESIGN** → Create basic GDD
2. **LEAD** → Review planning and assess feasibility
3. **DESIGN** → Create detailed feature specifications

### Phase 2: Architecture Design
1. **SERVER** → Design overall server architecture
2. **CLIENT** → Design client architecture
3. **LEAD** → Integration architecture review

### Phase 3: Feature Implementation
1. **DESIGN** → Create feature specification
2. **SERVER** → API design and implementation
3. **CLIENT** → UI/logic implementation
4. **LEAD** → Feature integration review

### Phase 4: Testing and Optimization
1. **LEAD** → Overall quality review
2. **SERVER** → Performance optimization
3. **CLIENT** → UI/UX optimization
4. **LEAD** → Final integration testing

## 📊 Project Status Check Methods
### Overall Status Overview
```markdown
LEAD agent, summarize current overall project status.
Minimal Reference: Check only each agent's _INDEX.md files
```

### Agent-Specific Status Check
```markdown
Check current work status of [AGENT] agent.
Reference: [AGENT-FOLDER]/_INDEX.md
```

## 💡 Efficient Prompt Writing Tips
### 1. Specify Clear Output Location
```markdown
Output: [AGENT-FOLDER]/[subfolder]/[filename].md
```

### 2. Specify Only Necessary Reference Documents
```markdown
Reference: 01-DESIGN-AGENT/Feature-Specifications/[specific-file].md
```

### 3. Clearly Limit Work Scope
```markdown
Current Request: [Specific and limited scope request]
Exclusions: [Parts not to implement]
```

### 4. Consider Token Usage
```markdown
Skip unnecessary documents and focus on core points for efficient work.
```

## 🔧 Development Efficiency Tips
### Batch Related Work
* Group similar tasks together
* Handle interdependent features simultaneously
* Minimize context switching between agents

### Use Index Files
* Check `_INDEX.md` for current status first
* Update index files after completing work
* Use for quick project status overview

### Optimize References
* Reference only documents directly related to current task
* Avoid loading entire document sets
* Use specific section references when possible

### Progressive Detail
* Start with high-level overview
* Add detail incrementally as needed
* Avoid over-engineering in early stages