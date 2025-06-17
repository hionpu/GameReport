---
tags:
  - agent/all
  - system/config
  - doc/guide
  - priority/critical
  - status/planning
references:
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
  - "[[05-MCP-Usage-Guide]]"
---

# Co-Located HAPDS (CL-HAPDS) Implementation Plan

## Purpose

Evolve the Human-AI Parallel Documentation System to maintain compressed versions alongside human documents in the same folder, using special file extensions that Obsidian ignores while enabling intelligent API-based context selection.

## Core Concept

Replace the dual folder structure (`001-Human/` and `002-AI-Context/`) with co-located files:

```
OLD STRUCTURE:
001-Human/00-SHARED/Project-Overview.md
002-AI-Context/00-SHARED/Project-Overview.md

NEW STRUCTURE:
00-SHARED/Project-Overview.md          <- Human readable (Obsidian native)
00-SHARED/Project-Overview.md.aicomp   <- AI compressed (hidden from Obsidian)
```

## Benefits of Co-Location

### ✅ **Simplified Structure**
- Single folder hierarchy eliminates dual-folder complexity
- Documents stay together for easier maintenance
- No risk of folder structure mismatches

### ✅ **Obsidian Integration**  
- Graph view remains clean (only shows .md files)
- Search and linking work normally
- Compressed files invisible to Obsidian workflows

### ✅ **Intelligent API**
- API automatically selects appropriate version based on context
- AI agents get compressed versions, humans get readable versions
- Transparent switching without user intervention

### ✅ **Maintenance Efficiency**
- Compressed file right next to source document
- Easy to verify synchronization
- Simpler update workflows

## File Extension Strategy

### Recommended Extensions

| Extension | Purpose | Obsidian Visibility |
|-----------|---------|-------------------|
| `.aicomp` | AI compressed version | Hidden |
| `.ctx` | Context optimized version | Hidden |
| `.tok` | Token optimized version | Hidden |

**Recommended**: Use `.aicomp` for clarity and consistency.

## Obsidian Configuration

### Hide Compressed Files

Add to Obsidian settings or configuration:

```json
{
  "excludedFiles": "*.aicomp,*.ctx,*.tok",
  "showUnsupportedFiles": false
}
```

This ensures compressed files don't appear in:
- File explorer
- Graph view  
- Search results
- Link suggestions

## API Implementation

### Enhanced MCP Functions

```typescript
// Intelligent context file retrieval
async function get_context_file(filename: string): Promise<string> {
  const compressedPath = filename + '.aicomp';
  
  if (await fileExists(compressedPath)) {
    // Return compressed version for AI context
    return await readFile(compressedPath);
  } else {
    // Fall back to human version
    return await readFile(filename);
  }
}

// Automatic compression generation
async function create_compressed_version(filename: string): Promise<void> {
  const humanContent = await readFile(filename);
  const compressedContent = await applyGenCompRules(humanContent);
  const compressedPath = filename + '.aicomp';
  
  await writeFile(compressedPath, compressedContent);
}

// Synchronization check
async function sync_versions(filename: string): Promise<void> {
  const compressedPath = filename + '.aicomp';
  const humanTimestamp = await getModifiedTime(filename);
  const compressedTimestamp = await getModifiedTime(compressedPath);
  
  if (humanTimestamp > compressedTimestamp) {
    await create_compressed_version(filename);
  }
}
```

### MCP Function Modifications

Enhance existing MCP functions to automatically use compressed versions for AI context:

```typescript
// Modified get_vault_file for AI context
get_vault_file(filename) {
  if (requestingAgent === 'AI') {
    return get_context_file(filename);
  } else {
    return get_standard_file(filename);
  }
}
```

## Migration Strategy

### Phase 1: Single Document Test
1. Choose one document (e.g., `01-Project-Overview.md`)
2. Create compressed version as `01-Project-Overview.md.aicomp`
3. Test API selection logic
4. Verify Obsidian ignores compressed file

### Phase 2: 00-SHARED Migration
1. For each file in `002-AI-Context/00-SHARED/`:
   - Move to corresponding `00-SHARED/[filename].aicomp`
   - Delete `002-AI-Context/00-SHARED/` folder
2. Update API to use new co-located structure
3. Test all MCP functions work correctly

### Phase 3: Agent Folders
1. Apply same migration to all agent folders:
   - `01-DESIGN-AGENT/`
   - `02-SERVER-AGENT/`
   - `03-CLIENT-AGENT/`
   - `04-LEAD-AGENT/`
2. Remove entire `002-AI-Context/` folder structure

### Phase 4: Documentation Update
1. Update all references to old dual-folder system
2. Update MCP usage guides
3. Update compression workflows

## Workflow Integration

### Document Creation
1. Create human-readable `.md` file normally in Obsidian
2. API automatically generates `.aicomp` version when needed
3. Compressed version invisible to human workflows

### Document Updates
1. Edit human `.md` file in Obsidian as normal
2. API detects timestamp difference
3. Automatically regenerates `.aicomp` version
4. AI context always stays synchronized

### Quality Assurance
- Regular automated sync checks
- Validation that compressed versions retain semantic content
- Monitoring that Obsidian properly ignores compressed files

## Implementation Checklist

### Technical Setup
- [ ] Configure Obsidian to ignore `.aicomp` files
- [ ] Implement `get_context_file()` API function
- [ ] Implement automatic compression generation
- [ ] Implement sync checking logic
- [ ] Test AI context selection works correctly

### Migration Tasks  
- [ ] Test with single document first
- [ ] Migrate 00-SHARED folder
- [ ] Migrate all agent folders
- [ ] Remove old 002-AI-Context structure
- [ ] Update all documentation references

### Validation
- [ ] Verify compressed files hidden from Obsidian
- [ ] Confirm AI agents receive compressed context
- [ ] Test synchronization automation
- [ ] Validate semantic equivalence maintained
- [ ] Check all MCP functions work with new structure

## Expected Token Savings

With the same 59.4% token reduction achieved by current HAPDS, but with significantly reduced structural complexity and maintenance overhead.

## Success Metrics

- **Structural Simplification**: Single folder hierarchy maintained
- **Obsidian Integration**: Compressed files completely invisible to users
- **API Intelligence**: Automatic version selection based on context
- **Maintenance Efficiency**: Compressed files stay synchronized automatically
- **Token Optimization**: Same ~60% reduction in AI context consumption

This Co-Located HAPDS approach provides all the benefits of the original HAPDS while dramatically simplifying the file structure and maintenance requirements.