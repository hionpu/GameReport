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