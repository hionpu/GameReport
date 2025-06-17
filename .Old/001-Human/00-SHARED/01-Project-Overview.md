---
tags:
  - agent/all
  - phase/planning
  - doc/gdd
  - tech/unity
  - tech/cpp
  - priority/critical
  - status/complete
references:
  - "[[04-Coding-Standards]]"
  - "[[05-MCP-Usage-Guide]]"
  - "[[03-Human-AI Parallel Documentation System (HAPDS)]]"
  - "[[02-General Compression Rules (GenCompRules)]]"
---

# RaidMaster Project Overview

## Project Information
- **Project Name**: RaidMaster
- **Genre**: 2D Action RPG
- **Target Platform**: PC (Windows/Mac/Linux)
- **Development Timeline**: 6 months (solo development)
- **Developer**: Solo developer (2 years WPF experience, new to game development)

## Technical Stack
### Client
- **Engine**: Unity 2023.x LTS
- **Language**: C#
- **UI**: Unity UI Toolkit
- **Rendering**: Universal Render Pipeline (URP)

### Server
- **Language**: C++ 17/20
- **Networking**: Boost.Asio / libuv
- **Database**: MySQL/PostgreSQL
- **Cache**: Redis
- **Build System**: CMake

### Common
- **Version Control**: Git
- **Documentation**: Obsidian + Claude MCP
- **CI/CD**: GitHub Actions (planned)

## Project Principles
1. **Practicality First**: Working code over perfect code
2. **Incremental Development**: MVP first, feature addition later
3. **Documentation Focus**: All decisions and designs must be documented
4. **AI-Assisted Development**: Leverage Claude agents for efficient development

## Agent Workflow Structure
### Agent Responsibilities
- **DESIGN**: Game design documents, feature specifications, user stories
- **SERVER**: C++ server architecture, API design, database schema
- **CLIENT**: Unity implementation, UI/UX, client-server communication
- **LEAD**: Integration management, technical decisions, quality assurance

### Folder Access Rights
- **DESIGN**: Works in `01-DESIGN-AGENT/`, reads `00-SHARED/`, `04-LEAD-AGENT/Integration-Reviews/`
- **SERVER**: Works in `02-SERVER-AGENT/`, reads `00-SHARED/`, `01-DESIGN-AGENT/Feature-Specifications/`
- **CLIENT**: Works in `03-CLIENT-AGENT/`, reads `00-SHARED/`, `01-DESIGN-AGENT/Feature-Specifications/`, `02-SERVER-AGENT/API-Specifications/`
- **LEAD**: Full access to all folders except `98-USER-ONLY/`

## MCP Context Management
### Included in MCP Context
- `Projects/RaidMaster/00-SHARED/` - Common project information
- `Projects/RaidMaster/01-DESIGN-AGENT/` - Design documents
- `Projects/RaidMaster/02-SERVER-AGENT/` - Server documents
- `Projects/RaidMaster/03-CLIENT-AGENT/` - Client documents
- `Projects/RaidMaster/04-LEAD-AGENT/` - Integration management

### Excluded from MCP Context
- `Projects/RaidMaster/98-USER-ONLY/` - User-only prompt templates
  - **Purpose**: Token optimization through template reference
  - **Usage**: User manually references for prompt creation

## Current Status
- **Phase**: Initial project setup
- **Completed**: Obsidian MCP integration, folder structure design
- **Next Steps**: Game concept finalization, GDD creation

## Development Goals
- **Primary**: Complete playable 2D Action RPG in 6 months
- **Secondary**: Establish efficient AI-assisted development workflow
- **Quality**: Focus on core gameplay over visual polish
- **Learning**: Gain practical game development experience

## 📋 Document References

### Related Documents:
- `00-SHARED/Document-Standards.md` - Documentation and tagging requirements
- `00-SHARED/04-Coding-Standards.md` - Implementation standards for all agents
- `00-SHARED/05-MCP-Usage-Guide.md` - Agent workflow and access patterns
- `01-DESIGN-AGENT/Game-Design-Documents/Main-GDD.md` - Detailed game design

### Dependencies:
- **Requires**: None (foundation document)
- **Blocks**: All other project documents depend on this overview

### Cross-Agent Impact:
- **DESIGN**: Provides technical constraints and project scope for game design
- **SERVER**: Defines server technology stack and performance requirements
- **CLIENT**: Specifies Unity version, platform targets, and client requirements
- **LEAD**: Establishes project timeline and development methodology
Use code with caution.
Yaml
And here's the second part of the document, assuming it's a separate file or was intended to be merged. If it's a separate file, it would have its own frontmatter. If it's part of the same file, it wouldn't have a second frontmatter block.
Let's assume for now it's meant to be part of the same logical document, and the references in the first frontmatter are intended to cover these as well (which they do).
# MCP Context Management
## MCP Context Management

### HAPDS Implementation
This project follows the **Human-AI Parallel Documentation System (HAPDS)** with dual folder structure:

#### 001-Human/ (Human-Readable Documentation)
- `001-Human/00-SHARED/` - Common project information
- `001-Human/01-DESIGN-AGENT/` - Design documents  
- `001-Human/02-SERVER-AGENT/` - Server documents
- `001-Human/03-CLIENT-AGENT/` - Client documents
- `001-Human/04-LEAD-AGENT/` - Integration management
- `001-Human/98-USER-ONLY/` - User-only prompt templates (excluded from MCP context)

#### 002-AI-Context/ (Compressed for AI Consumption)
- `002-AI-Context/00-SHARED/` - Compressed common project information
- `002-AI-Context/01-DESIGN-AGENT/` - Compressed design documents
- `002-AI-Context/02-SERVER-AGENT/` - Compressed server documents  
- `002-AI-Context/03-CLIENT-AGENT/` - Compressed client documents
- `002-AI-Context/04-LEAD-AGENT/` - Compressed integration management
- `002-AI-Context/98-USER-ONLY/` - Compressed user-only documents

### MCP Context Strategy
- **For AI Agents**: Use `002-AI-Context/` folder for token-optimized context
- **For Human Work**: Use `001-Human/` folder for daily work and editing
- **Synchronization**: Both folders maintain identical semantic content in different formats
- **Reference**: See `[[03-Human-AI Parallel Documentation System (HAPDS)]]` and `[[02-General Compression Rules (GenCompRules)]]` for detailed implementation