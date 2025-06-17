---
tags:
  - agent/all
  - system/config
  - phase/planning
  - doc/guide
  - tech/unity
  - tech/cpp
  - priority/critical
  - status/complete
references:
  - "[[01-Project-Overview]]"
  - "[[04-Coding-Standards]]"
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
---

# RaidMaster MCP Agent Access Control & Workflow Guide

## 📁 Project Structure Overview

This project implements the **Human-AI Parallel Documentation System (HAPDS)** with dual folder structure:

### 001-Human/ (Human-Readable Documentation)
```
001-Human/
├── 00-SHARED/              # Common reference documents
├── 01-DESIGN-AGENT/        # Design agent workspace
├── 02-SERVER-AGENT/        # Server agent workspace  
├── 03-CLIENT-AGENT/        # Client agent workspace
├── 04-LEAD-AGENT/          # Lead agent workspace
└── 98-USER-ONLY/           # User-only (excluded from MCP context)
```

### 002-AI-Context/ (Compressed for AI Consumption)
```
002-AI-Context/
├── 00-SHARED/              # Compressed common reference documents
├── 01-DESIGN-AGENT/        # Compressed design agent workspace
├── 02-SERVER-AGENT/        # Compressed server agent workspace  
├── 03-CLIENT-AGENT/        # Compressed client agent workspace
├── 04-LEAD-AGENT/          # Compressed lead agent workspace
└── 98-USER-ONLY/           # Compressed user-only documents
```

### HAPDS Usage Guidelines
- **Human Work**: Use `001-Human/` folder for daily work, reading, editing, and collaboration
- **AI Context**: Use `002-AI-Context/` folder when providing project context to AI agents
- **Synchronization**: Both folders maintain identical semantic content in different formats
- **Token Optimization**: AI context folder enables sharing comprehensive project context without token limitations

### HAPDS Implementation Requirements
- **Content Synchronization**: When updating documents in `001-Human/`, corresponding compressed versions must be created/updated in `002-AI-Context/`
- **Compression Standards**: All AI context documents must follow `[[02-General Compression Rules (GenCompRules)]]` and maintain proper YAML frontmatter with `---` delimiters
- **Reference Integrity**: Wikilink references must be maintained in both human and compressed formats
- **Quality Assurance**: Regular verification that both versions contain identical semantic information
- **Documentation**: See `[[03-Human-AI Parallel Documentation System (HAPDS)]]` for complete implementation guidelines

## 🔐 Agent Access Control Matrix

### DESIGN Agent (Code: DESIGN)

**✅ Read Access:**
- `00-SHARED/` - Project standards, coding guidelines, document standards
- `01-DESIGN-AGENT/` - Own workspace (full read/write)
- `04-LEAD-AGENT/Integration-Reviews/` - Integration feedback affecting design decisions

**🎯 Primary Responsibilities:**
- Game design documents (GDD) and feature specifications
- User experience design and game balance
- Requirements definition for other agents

### SERVER Agent (Code: SERVER)

**✅ Read Access:**
- `00-SHARED/` - Project standards, coding guidelines, document standards
- `02-SERVER-AGENT/` - Own workspace (full read/write)
- `01-DESIGN-AGENT/Feature-Specifications/` - Requirements for API implementation

**🎯 Primary Responsibilities:**
- C++ server architecture and API implementation
- Database design and performance optimization
- Server-side game logic and security

### CLIENT Agent (Code: CLIENT)

**✅ Read Access:**
- `00-SHARED/` - Project standards, coding guidelines, document standards
- `03-CLIENT-AGENT/` - Own workspace (full read/write)
- `01-DESIGN-AGENT/Feature-Specifications/` - UI/UX requirements and user stories
- `02-SERVER-AGENT/API-Specifications/` - Server communication interfaces

**🎯 Primary Responsibilities:**
- Unity client implementation and UI/UX
- Server communication and data synchronization
- Platform optimization and user experience

### LEAD Agent (Code: LEAD)

**✅ Read Access:**
- **ALL folders except `98-USER-ONLY/`** - Full project oversight

**🎯 Primary Responsibilities:**
- Cross-agent integration and conflict resolution
- Technical architecture decisions and quality assurance
- Project timeline management and risk mitigation

## 🏷️ MANDATORY DOCUMENT STANDARDS

### Required YAML Frontmatter (Every Document Must Include)

All documents **MUST** maintain proper YAML frontmatter with `---` delimiters:

```yaml
---
tags:
  # Agent (required - choose primary responsible agent)
  - agent/design          # Design specifications, game mechanics
  - agent/server          # Server implementation, APIs
  - agent/client          # Client implementation, UI/UX
  - agent/lead            # Integration, architecture decisions
  - agent/all             # Shared resources, standards
  
  # System (required - choose main focus area)
  - system/combat         # Combat mechanics
  - system/inventory      # Item management
  - system/party          # Party/recruitment systems
  - system/character      # Character progression
  - system/network        # Networking systems
  - system/database       # Database systems
  - system/config         # Configuration/setup
  
  # Priority & Status (required)
  - priority/critical     # Must have for launch
  - priority/high         # Should have for launch
  - priority/medium       # Nice to have
  - priority/low          # Future consideration
  - status/planning       # Planning stage
  - status/in-progress    # Active development
  - status/complete       # Finished
  - status/blocked        # Waiting on dependencies

# CRITICAL: Wikilink References (MANDATORY)
references:
  - "[[Document-Name]]"    # Related documents using Obsidian wikilinks
  - "[[Another-Document]]" # Enables graph view and automated processing
---
```

### ⚠️ CRITICAL NEW RULE: Wikilink References

**MANDATORY FOR ALL DOCUMENTS:**

Every document MUST include a `references` property in the YAML frontmatter with proper `---` delimiters, using Obsidian wikilinks to related documents with `[[Doc-Name]]` format.

### Required Document Structure Template

```markdown
---
tags:
  - agent/[primary-agent]
  - system/[main-system] 
  - priority/[level]
  - status/[current-status]
references:
  - "[[Related-Document-1]]"
  - "[[Related-Document-2]]"
---

# [Document Title]

## 📋 Document References

### Related Documents:
- `[folder]/[filename].md` - [brief description]

### Dependencies:
- **Requires**: [list prerequisite documents]
- **Blocks**: [list documents waiting on this]

### Cross-Agent Impact:
- **DESIGN**: [how this affects design decisions]
- **SERVER**: [how this affects server implementation]  
- **CLIENT**: [how this affects client implementation]

[Main content...]
```

## 🔄 Efficient Workflow Patterns

### 1. Feature Development Workflow

1. **DESIGN** → Create specification in `01-DESIGN-AGENT/Feature-Specifications/` → Reference `00-SHARED/` standards and related feature specs → Include proper tags, wikilink references, cross-agent impact → Output complete feature requirements with dependencies

2. **SERVER** → Read feature spec + referenced dependencies → Create API design in `02-SERVER-AGENT/API-Specifications/` → Reference design spec, coding standards, related APIs → Output server implementation with proper documentation

3. **CLIENT** → Read feature spec + API spec + referenced dependencies → Create UI implementation in `03-CLIENT-AGENT/` → Reference design requirements, server APIs, UI standards → Output complete client-side feature implementation

4. **LEAD** integration review → Review all outputs from steps 1-3 → Check cross-references, dependency conflicts, standards compliance → Output integration review with recommendations

### 2. Document Creation Checklist (MANDATORY)

Every new document MUST include:
- [ ] **Proper YAML frontmatter with `---` delimiters and required tags**
- [ ] **🔥 CRITICAL: Wikilink references property with `[[Document-Name]]` format**
- [ ] Document References section with related files
- [ ] Dependencies (requires/blocks) clearly stated
- [ ] Cross-agent impact analysis
- [ ] Brief overview/summary
- [ ] References to 00-SHARED standards when applicable

## 🎯 Practical Usage Examples

### Example 1: New Feature Implementation

**Scenario**: DESIGN agent creating inventory system specification

**Access**: 
- `00-SHARED/` (project standards, document standards)  
- `01-DESIGN-AGENT/` (workspace)

**Task Requirements**:
- Follow Document-Standards.md template
- **MUST include wikilink references in YAML frontmatter with `---` delimiters**
- Reference Party-Recruitment-System.md for consistency
- Include proper tags and cross-agent impact analysis
- Specify dependencies and requirements clearly

**Expected Output**: `Feature-Specifications/Inventory-System.md`

```yaml
---
tags:
  - agent/design
  - system/inventory
  - priority/high
  - status/planning
references:
  - "[[01-Main-GDD]]"
  - "[[01-Party-Recruitment-System]]"
  - "[[01-Project-Overview]]"
---
```

### Example 2: Server Implementation

**Scenario**: SERVER agent implementing inventory system API

**Access**:
- `00-SHARED/` (standards)
- `02-SERVER-AGENT/` (workspace)
- `01-DESIGN-AGENT/Feature-Specifications/Inventory-System.md` (requirements)

**Requirements**:
- Follow coding standards and documentation template
- **MUST include wikilink references with proper YAML frontmatter**
- Reference design specification and include API documentation
- Include proper tags and specify integration points with other systems

**Expected Output**: `API-Specifications/Inventory-API.md` + implementation code

```yaml
---
tags:
  - agent/server
  - system/inventory
  - priority/high
  - status/in-progress
references:
  - "[[Inventory-System]]"
  - "[[04-Coding-Standards]]"
  - "[[01-Project-Overview]]"
---
```

### Example 3: Integration Review

**Scenario**: LEAD agent reviewing inventory system implementation integration

**Access**: Full project access

**Requirements**:
- Check design specification, server API, and client implementation
- Verify proper documentation standards compliance
- **MUST verify all documents have wikilink references with proper YAML frontmatter**
- Identify potential conflicts and missing dependencies
- Ensure cross-agent references are accurate

**Expected Output**: `Integration-Reviews/Inventory-System-Review.md`

```yaml
---
tags:
  - agent/lead
  - system/inventory
  - priority/critical
  - status/complete
references:
  - "[[Inventory-System]]"
  - "[[Inventory-API]]"
  - "[[Inventory-UI]]"
  - "[[Document-Standards]]"
---
```

## ⚠️ Critical Guidelines

### For All Agents

1. **Always check Document-Standards.md** before creating new documents
2. **Include proper tags** - minimum 3 categories (agent, system, priority)
3. **🔥 MANDATORY: Add wikilink references** - use `[[Document-Name]]` format in YAML frontmatter with `---` delimiters
4. **Add reference sections** - clearly state related documents and dependencies
5. **Cross-reference accurately** - ensure referenced documents exist and are current
6. **Update existing references** when creating new documents that others depend on

### Document Quality Standards

- **Consistency**: Follow established templates and naming conventions
- **Traceability**: Clear references enable easy navigation between related documents
- **Completeness**: Include all required sections and metadata
- **Maintainability**: Update references when documents are modified or moved
- **🔥 Wikilink Compliance**: All documents must have proper YAML references property with `---` delimiters

### Token Optimization

- **Selective Reading**: Only access documents directly relevant to the task
- **Reference Efficiently**: Link to existing documents instead of duplicating content
- **Use Tags**: Leverage tag system for quick document discovery
- **Batch Related Work**: Group similar tasks to maintain context

## 🚀 Success Metrics

### Documentation Quality

- **100% Tag Compliance**: All documents include required tags
- **🔥 100% Wikilink Compliance**: All documents have proper references property with `---` delimiters
- **Complete References**: All documents properly reference dependencies
- **Cross-Agent Clarity**: Clear impact analysis for all agent interactions
- **Dependency Tracking**: No missing or broken document references

### Workflow Efficiency

- **Reduced Context Switching**: Clear references minimize need to search for related documents
- **Faster Onboarding**: New agents can quickly understand project structure
- **Better Integration**: Clear dependencies prevent conflicts between agent outputs
- **Maintained Quality**: Consistent standards across all project documentation
- **🔥 Graph Connectivity**: All documents connected through Obsidian graph view

## 🔥 ENFORCEMENT RULES

### For LEAD Agent

**Mandatory Quality Checks:**
1. **Verify all new documents have wikilink references in YAML frontmatter with `---` delimiters**
2. **Reject any document submissions without proper references property or missing YAML delimiters**
3. **Update existing documents when new dependencies are created**
4. **Maintain graph connectivity across all project documents**

### For All Agents

**Before Document Creation:**
1. Identify which documents this new document relates to
2. Add those documents to the `references` property using `[[Document-Name]]` format with proper YAML frontmatter
3. Verify referenced documents exist
4. Update referenced documents if they need to reference back

## 🔄 Automatic Synchronization Rule

**MANDATORY AUTO-SYNC**: When any document in `001-Human/` folder is updated, the corresponding compressed version in `002-AI-Context/` folder **MUST** be immediately updated.

### Sync Triggers
- **Content Modifications**: Any section added, edited, or deleted
- **Status Changes**: Tag updates for status, priority, or phase
- **Reference Changes**: Adding or removing wikilink references  
- **Version Changes**: Document version increments

### Auto-Sync Process
1. **Detect** human document change
2. **Identify** corresponding AI-Context file path
3. **Apply** GenCompRules compression to human content
4. **Update** AI-Context file with compressed version
5. **Verify** YAML frontmatter integrity maintained
6. **Confirm** sync completion

### Bypass Prevention
- Manual sync disable only for temporary bulk edits
- **Must re-enable** sync after bulk operations complete
- No permanent bypass allowed - sync integrity critical

## 🚫 CRITICAL: Never Use Active Document APIs

**MANDATORY RULE: Always Target Specific Documents**

**NEVER** use active document APIs (`get_active_file`, `update_active_file`, `patch_active_file`, `append_to_active_file`). The user may change the active document while you're working, causing you to modify the wrong file.

### Required Approach (MANDATORY)
**ALWAYS** use specific file path APIs to target the exact document you intend to modify:

- ✅ **USE**: `get_vault_file("path/to/specific/file.md")` 
- ✅ **USE**: `create_vault_file("path/to/specific/file.md", content)`
- ✅ **USE**: `patch_vault_file("path/to/specific/file.md", ...)`
- ✅ **USE**: `append_to_vault_file("path/to/specific/file.md", content)`

### Prevention Rules
- ❌ **NEVER** use `get_active_file`, `update_active_file`, `patch_active_file`, or `append_to_active_file`
- ❌ **NEVER** assume you know which file is currently active
- ❌ **NEVER** rely on active document state for file operations
- ✅ **ALWAYS** specify the complete file path for the document you want to modify
- ✅ **ALWAYS** confirm the target file path before any modification

### Why This Rule Exists
**Problem**: User changes active document while AI is working → AI modifies wrong file
**Solution**: Always specify exact file paths → No dependency on active document state
**Benefit**: Guaranteed correct file targeting regardless of user's current document focus

### File Path Examples
```
✅ CORRECT: get_vault_file("001-Human/00-SHARED/05-MCP-Usage-Guide.md")
✅ CORRECT: create_vault_file("002-AI-Context/01-DESIGN-AGENT/New-Feature.md", content)
❌ WRONG: update_active_file(content)  // Could modify any file!
❌ WRONG: patch_active_file(...)       // Dangerous - unknown target!
```