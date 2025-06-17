---
tags:
  - agent/all
  - system/config
  - doc/guide
  - priority/critical
  - status/complete
references:
  - "[[01-Project-Overview]]"
  - "[[05-MCP-Usage-Guide]]"
---

# Numbering Reference System

## Purpose

Establish a consistent numbering system for quick document reference and automated path resolution across the project structure.

## Reference Format

**Pattern**: `XX/YY` or `XX/YY/ZZ` automatically resolves to numbered items

**Components**:
- `XX`: Folder level (00 = 00-SHARED, 01 = 01-DESIGN-AGENT, etc.)
- `YY`: Item number within that folder (01-, 02-, 03-, etc.)
- `ZZ`: Sub-item number (for nested folders/files)

## Auto-Resolution Rules

### Basic Pattern
- `00/01` → `00-SHARED/01-Project-Overview.md`
- `00/02` → `00-SHARED/02-General-Compression-Rules.md`
- `00/06` → `00-SHARED/06-Numbering-Reference-System.md`

### Nested Pattern  
- `01/01` → `01-DESIGN-AGENT/01-Game-Design-Documents/`
- `01/01/01` → `01-DESIGN-AGENT/01-Game-Design-Documents/01-Main-GDD.md`
- `01/02/01` → `01-DESIGN-AGENT/02-Feature-Specifications/01-Party-Recruitment-System.md`

### Folder vs File Resolution
- System automatically determines if reference points to folder or file
- `XX/YY` can resolve to either `XX-FOLDER/YY-ITEM` (folder or file)
- Context and usage determine the appropriate resolution

## Current Structure Map

**Note**: This map shows the numbering pattern and examples. For the complete current structure, use Obsidian's file explorer or search functionality.

### Numbering Pattern Examples

#### 00-SHARED (Base Reference Documents)
- `00/01` → `01-Project-Overview.md`
- `00/02` → `02-General Compression Rules (GenCompRules).md`  
- `00/03` → `03-Human-AI Parallel Documentation System (HAPDS).md`
- `00/04` → `04-Coding-Standards.md`
- `00/05` → `05-MCP-Usage-Guide.md`
- `00/06` → `06-Numbering-Reference-System.md`
- `00/XX` → (Additional shared documents as they are created)

#### 01-DESIGN-AGENT
- `01/01` → `01-Game-Design-Documents/` (folder)
  - `01/01/01` → `01-Main-GDD.md`
  - `01/01/XX` → (Additional GDD documents)
- `01/02` → `02-Feature-Specifications/` (folder)
  - `01/02/01` → `01-Party-Recruitment-System.md`
  - `01/02/XX` → (Additional feature specifications)
- `01/XX` → (Additional design agent folders)

#### Other Agent Folders
- `02/XX/XX` → SERVER-AGENT documents (structure to be established)
- `03/XX/XX` → CLIENT-AGENT documents (structure to be established)
- `04/XX/XX` → LEAD-AGENT documents (structure to be established)
- `98/XX/XX` → USER-ONLY documents

### Finding Current Structure
Instead of maintaining this map manually:
1. **Use Obsidian File Explorer**: Browse the actual folder structure
2. **Use Obsidian Search**: Search for specific numbered documents
3. **Use Graph View**: Visualize document relationships
4. **Check Folder Contents**: Use MCP tools to list current files

## Usage Patterns

### Quick Reference Format
- **Short Form**: `00/01` for quick verbal/written reference
- **Context Aware**: System determines folder vs file based on context
- **Hierarchical**: `01/02/01` for nested items

### AI Agent Recognition
- AI agents automatically recognize `XX/YY/ZZ` pattern
- Expand to full path: `002-AI-Context/XX-FOLDER/YY-SUBFOLDER/ZZ-FILE.md`
- Human agents use same shorthand for consistency

### Cross-Reference Examples
```
"See 00/01 for project overview"
"Reference 01/01/01 for main game design"  
"Check 00/05 for MCP usage guidelines"
"Party system details in 01/02/01"
```

## Implementation Guidelines

### Document Creation
1. **Assign Next Number**: New items get next sequential number in their folder
2. **Follow Convention**: All items use two-digit prefix (01-, 02-, etc.)
3. **No Manual Map Updates**: Use Obsidian tools for current structure instead of maintaining this map

### Folder Reorganization
1. **Maintain Numbering**: Keep sequential order during restructuring
2. **Update Cross-References**: Batch update all affected references
3. **Consistency Check**: Verify both 001-Human and 002-AI-Context match

### Reference Updates
1. **Propagate Changes**: Update references in both folder systems
2. **Verify Resolution**: Ensure all `XX/YY/ZZ` references resolve correctly
3. **Document Updates**: Update this reference system when structure changes (but not the map)

## Maintenance Rules

### Sequential Numbering
- New documents assigned next available number in sequence
- Gaps in numbering indicate deleted/moved documents
- Maintain chronological order when possible

### Consistency Requirements  
- **Both Folders**: 001-Human and 002-AI-Context must have identical numbering
- **Cross-References**: All `XX/YY/ZZ` references must resolve to existing items
- **No Manual Map Maintenance**: Use automated tools to discover current structure

### Update Procedures
1. **Add New Item**: Assign next number, update both folders *(no map update needed)*
2. **Move Item**: Update number if needed, update all references *(no map update needed)*
3. **Delete Item**: Remove from both folders, update references, note gap in sequence *(no map update needed)*

## HAPDS Format Compliance Rules

### CRITICAL PREVENTION RULE
**001-Human/ folders MUST contain human-readable markdown format ONLY**
- ❌ **NEVER** use compressed GenCompRules syntax in 001-Human/ documents
- ✅ **ALWAYS** use full markdown with proper headers, lists, and explanations
- ✅ **ALWAYS** verify format before saving documents in 001-Human/

### Format Verification Checklist
Before saving any document in 001-Human/:
- [ ] **Headers**: Uses `#`, `##`, `###` (not `H1:`, `Sec:`, `SubSec:`)
- [ ] **Lists**: Uses `-` or `1.` format (not comma-separated strings)
- [ ] **Text**: Uses full sentences and paragraphs (not abbreviated syntax)
- [ ] **Structure**: Uses standard markdown structure (not compressed key-value pairs)

### Quality Control Process
1. **Before Creating**: Determine target folder (001-Human vs 002-AI-Context)
2. **During Writing**: Use appropriate format for target folder
3. **Before Saving**: Verify format compliance with folder type
4. **After Saving**: Cross-check that human version is readable, AI version is compressed

## 📋 Document References

### Related Documents:
- `00-SHARED/01-Project-Overview.md` - Project structure and organization
- `00-SHARED/05-MCP-Usage-Guide.md` - Agent workflow and access patterns  
- `00-SHARED/03-Human-AI Parallel Documentation System (HAPDS).md` - Dual folder system

### Dependencies:
- **Requires**: Consistent numbering across all project documents
- **Blocks**: Efficient cross-referencing and navigation system

### Cross-Agent Impact:
- **ALL AGENTS**: Use `XX/YY/ZZ` format for quick document references
- **LEAD**: Maintains numbering consistency during reorganizations  
- **AI AGENTS**: Automatically resolve shorthand references to full paths
- **HUMAN USERS**: Use shorthand for quick communication and documentation