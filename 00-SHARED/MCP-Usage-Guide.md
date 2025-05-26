# RaidMaster MCP Usage Guide

## Project Structure
```
Projects/RaidMaster/
├── 00-SHARED/              # Common reference documents
├── 01-DESIGN-AGENT/        # Design agent workspace
├── 02-SERVER-AGENT/        # Server agent workspace  
├── 03-CLIENT-AGENT/        # Client agent workspace
├── 04-LEAD-AGENT/          # Lead agent workspace
└── 98-USER-ONLY/           # User-only (excluded from MCP context)
```

## MCP Context Management

### ✅ Included Folders (MCP accessible)
- `Projects/RaidMaster/00-SHARED/` - Common project information
- `Projects/RaidMaster/01-DESIGN-AGENT/` - Design documents
- `Projects/RaidMaster/02-SERVER-AGENT/` - Server documents
- `Projects/RaidMaster/03-CLIENT-AGENT/` - Client documents
- `Projects/RaidMaster/04-LEAD-AGENT/` - Integration management documents

### ❌ Excluded Folders (MCP context excluded)
- `Projects/RaidMaster/98-USER-ONLY/` - User-only prompt templates
  - **Purpose**: Token optimization through reference-only prompt storage
  - **Contents**: Agent-specific prompt patterns and templates
  - **Usage**: User manually references to create prompts

## Agent-Specific Workflows

### DESIGN Agent
- **Primary Workspace**: `Projects/RaidMaster/01-DESIGN-AGENT/`
- **Read Access**: `00-SHARED/`, `04-LEAD-AGENT/Integration-Reviews/`
- **Prompt Templates**: `98-USER-ONLY/DESIGN-Agent-Prompts.md` (user reference only)

### SERVER Agent
- **Primary Workspace**: `Projects/RaidMaster/02-SERVER-AGENT/`
- **Read Access**: `00-SHARED/`, `01-DESIGN-AGENT/Feature-Specifications/`
- **Prompt Templates**: `98-USER-ONLY/SERVER-Agent-Prompts.md` (user reference only)

### CLIENT Agent
- **Primary Workspace**: `Projects/RaidMaster/03-CLIENT-AGENT/`
- **Read Access**: `00-SHARED/`, `01-DESIGN-AGENT/Feature-Specifications/`, `02-SERVER-AGENT/API-Specifications/`
- **Prompt Templates**: `98-USER-ONLY/CLIENT-Agent-Prompts.md` (user reference only)

### LEAD Agent
- **Primary Workspace**: `Projects/RaidMaster/04-LEAD-AGENT/`
- **Read Access**: All folders (except 98-USER-ONLY)
- **Prompt Templates**: `98-USER-ONLY/LEAD-Agent-Prompts.md` (user reference only)

## Token Optimization Strategies

### 1. Selective Document Reference
```markdown
# Efficient approach
SERVER agent, implement inventory API.
Reference: Projects/RaidMaster/01-DESIGN-AGENT/Feature-Specifications/Inventory-System.md

# Inefficient approach
SERVER agent, implement inventory API.
(References all documents indiscriminately)
```

### 2. Batch Processing
```markdown
# Efficient method
DESIGN agent, plan these 3 related features together:
1. Player stats system
2. Level up system  
3. Experience gain system

# Inefficient method
Separate requests for each of the 3 features
```

### 3. Index File Priority
```markdown
# Status check
First check Projects/RaidMaster/01-DESIGN-AGENT/_INDEX.md

# Detailed work
Selectively reference only necessary documents
```

## Workflow Process

### Step 1: Template Reference
1. Open `Projects/RaidMaster/98-USER-ONLY/[AGENT]-Agent-Prompts.md`
2. Select appropriate prompt pattern for the task
3. Customize with specific request details

### Step 2: MCP Agent Call
1. Use customized prompt to call Claude
2. Agent performs work in designated folder
3. Results are created/updated in appropriate location

### Step 3: Review and Next Steps
1. Review generated/modified documents
2. Use LEAD agent for integration review if needed
3. Call next agent for follow-up work

## Important Notes

### MCP Usage
- `98-USER-ONLY` folder is NOT directly referenced by Claude
- User manually references templates in that folder for prompt creation
- Avoid unnecessary document references to save tokens

### Document Management
- Keep each agent's `_INDEX.md` files up to date
- Document all important decisions
- Use LEAD agent for conflict resolution between agents

### Development Efficiency
- Group related tasks for batch processing
- Use index files for quick status overview
- Approach complex decisions step by step (design → review → implement)