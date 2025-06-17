---
tags:
  - agent/all
  - system/config
  - phase/planning
  - doc/guide
  - priority/critical
  - status/complete
---

# Obsidian Tag System - Game Development Project

## Tag Structure Overview

Each document should include tags from the following categories:
- **Agent Tags** - Responsible AI agent
- **System Tags** - Game/technical systems
- **Development Phase Tags** - Project stage and document type
- **Technical Tags** - Technology stack and implementation details
- **Priority Tags** - Importance level
- **Status Tags** - Current state

## 1. Agent Tags

### Primary Responsible Agent
- `#agent/design` - Design agent related
- `#agent/server` - C++ server agent related
- `#agent/client` - Unity client agent related
- `#agent/lead` - Integration management agent related

### Collaboration Tags
- `#agent/design-server` - Design-server collaboration needed
- `#agent/server-client` - Server-client collaboration needed
- `#agent/design-client` - Design-client collaboration needed
- `#agent/all` - All agents involved

## 2. System Tags

### Game Systems
- `#system/combat` - Combat system
- `#system/inventory` - Inventory system
- `#system/party` - Party system
- `#system/character` - Character system
- `#system/quest` - Quest system
- `#system/economy` - Economy system
- `#system/social` - Social system
- `#system/progression` - Progression/growth system
- `#system/world` - World/map system
- `#system/tutorial` - Tutorial system

### Technical Systems
- `#system/network` - Network system
- `#system/database` - Database system
- `#system/auth` - Authentication system
- `#system/security` - Security system
- `#system/performance` - Performance optimization
- `#system/logging` - Logging system
- `#system/config` - Configuration management system

## 3. Development Phase Tags

### Project Phases
- `#phase/concept` - Concept phase
- `#phase/planning` - Planning phase
- `#phase/design` - Design phase
- `#phase/prototype` - Prototype phase
- `#phase/implementation` - Implementation phase
- `#phase/testing` - Testing phase
- `#phase/optimization` - Optimization phase
- `#phase/polish` - Polish phase

### Document Types
- `#doc/gdd` - Game Design Document
- `#doc/spec` - Feature specification
- `#doc/api` - API documentation
- `#doc/architecture` - Architecture document
- `#doc/guide` - Guide/manual
- `#doc/meeting` - Meeting notes
- `#doc/decision` - Decision document
- `#doc/research` - Research document

## 4. Technical Tags

### Client Technology
- `#tech/unity` - Unity related
- `#tech/ui` - UI/UX related
- `#tech/graphics` - Graphics/rendering
- `#tech/animation` - Animation
- `#tech/audio` - Audio
- `#tech/input` - Input handling
- `#tech/mobile` - Mobile optimization
- `#tech/performance-client` - Client performance

### Server Technology
- `#tech/cpp` - C++ related
- `#tech/networking` - Networking
- `#tech/database` - Database
- `#tech/async` - Asynchronous programming
- `#tech/threading` - Multithreading
- `#tech/memory` - Memory management
- `#tech/security-server` - Server security
- `#tech/performance-server` - Server performance

### Integration Technology
- `#tech/protocol` - Communication protocol
- `#tech/serialization` - Serialization
- `#tech/sync` - Synchronization
- `#tech/testing` - Testing
- `#tech/deployment` - Deployment
- `#tech/monitoring` - Monitoring

## 5. Priority Tags

- `#priority/critical` - Critical importance
- `#priority/high` - High priority
- `#priority/medium` - Medium priority
- `#priority/low` - Low priority
- `#priority/backlog` - Backlog item

## 6. Status Tags

- `#status/planning` - Planning stage
- `#status/in-progress` - In progress
- `#status/review` - Under review
- `#status/testing` - Testing
- `#status/complete` - Completed
- `#status/blocked` - Blocked
- `#status/deprecated` - Deprecated

## Practical Tag Usage Examples

### Party Recruitment System Document
```yaml
tags:
  # Agent
  - agent/design
  - agent/client
  - agent/server
  
  # System
  - system/party
  - system/social
  - system/network
  
  # Development Phase
  - phase/design
  - doc/spec
  
  # Technology
  - tech/ui
  - tech/networking
  - tech/database
  
  # Priority & Status
  - priority/high
  - status/planning
```

### Inventory System API Document
```yaml
tags:
  # Agent
  - agent/server
  - agent/client
  
  # System
  - system/inventory
  - system/database
  
  # Development Phase
  - phase/implementation
  - doc/api
  
  # Technology
  - tech/cpp
  - tech/protocol
  - tech/serialization
  
  # Priority & Status
  - priority/critical
  - status/in-progress
```

### Combat System UI Design
```yaml
tags:
  # Agent
  - agent/design
  - agent/client
  
  # System
  - system/combat
  - system/character
  
  # Development Phase
  - phase/design
  - doc/spec
  
  # Technology
  - tech/unity
  - tech/ui
  - tech/animation
  
  # Priority & Status
  - priority/high
  - status/review
```

## Tag Utilization Tips

### 1. Search & Filtering
- `tag:#agent/design AND tag:#system/party` - Find party-related design documents
- `tag:#status/blocked` - Check blocked tasks
- `tag:#priority/critical AND tag:#status/in-progress` - Monitor critical ongoing work

### 2. Dashboard Creation
```dataview
TABLE priority, status, agent
FROM #system/inventory
SORT priority DESC, status ASC
```

### 3. Progress Tracking
```dataview
TABLE count(rows) AS "Document Count"
FROM ""
GROUP BY status
```

## Tag Management Guidelines

1. **Consistency**: Include tags from at least 3 categories in every document
2. **Specificity**: Use specific tags when possible
3. **Updates**: Update tags immediately when status changes
4. **Minimize Duplicates**: Focus on core tags rather than over-tagging
5. **Scalability**: Extend tag system when adding new systems/features

## Benefits for Solo Development

1. **Quick Context Switching** - Rapidly find relevant documentation when working with different AI agents
2. **Progress Tracking** - Monitor what's blocked, in progress, or completed at a glance
3. **Smart Queries** - Use Dataview plugin to create dynamic project dashboards
4. **Reduced Token Usage** - Find specific documents faster, reducing need to re-explain context to AI agents
5. **Knowledge Management** - Maintain organized documentation as project grows

This tag system enables efficient project management and seamless collaboration with AI agents throughout your game development journey.