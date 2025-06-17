---
tags:
  - agent/design
  - system/combat
  - system/party
  - system/progression
  - phase/concept
  - doc/gdd
  - tech/unity
  - priority/critical
  - status/in-progress
references:
  - "[[Project-Overview]]"
  - "[[01-Party-Recruitment-System]]"
---

# RaidMaster - Main Game Design Document

## 📋 Document Information
- **Version**: 1.0
- **Last Updated**: 2025-05-26
- **Status**: In Development
- **Next Review**: TBD

---

## 🎯 Game Overview

### Core Concept
**RaidMaster** is a 3D isometric action RPG focused on strategic raid encounters. Players hire party members from a tavern to form balanced 5-person teams (1 Tank, 2 Healers, 2 Dealers) and challenge increasingly difficult raid bosses. The game features daily progression through material collection, equipment crafting, and party optimization.

### Genre & Platform
- **Genre**: 3D Isometric Action RPG
- **Camera**: Fixed quarter-view (isometric)
- **Platform**: Mobile + PC (Cross-platform)
- **Target Rating**: T for Teen (Combat, mild fantasy violence)

### Key Selling Points
- **Strategic Team Building**: Recruit and develop unique party members
- **Daily Progression**: Meaningful advancement through limited daily attempts
- **Complex Crafting**: Deep equipment customization system
- **Scalable Challenge**: From solo 5-man raids to group 25-man events

---

## 🎮 Core Gameplay

### Primary Game Loop
**Daily Raid Cycle:**
1. **Planning Phase**: Check available raids, review party composition
2. **Recruitment Phase**: Visit tavern to recruit/manage party members (with limitations)
3. **Raid Execution**: Challenge bosses with 5-person team (1 Tank, 2 Healers, 2 Dealers)
4. **Progression Phase**: Collect materials, craft/upgrade equipment, level party members
5. **Reset**: Daily attempts refresh, new recruitment opportunities

### Multiplayer Scope
- **Primary Content**: Solo player with hired AI party (5-man raids)
- **Secondary Content**: 10-man cooperative raids (2 players, future feature)
- **Special Events**: 25-man raids (5 players, limited-time monthly events)

### Core Systems

#### Daily Access System
- **Daily Reset Model**: Each raid has limited daily attempts
- **Attempt Count**: Varies by raid difficulty (3-5 attempts per raid)
- **Reset Time**: Midnight local time

#### Party Member System
- **Progression**: Party members gain experience, levels, and skills over time
- **Recruitment Limits**: [DETAILS NEEDED]
- **Randomness Elements**: [DETAILS NEEDED]

#### Equipment & Crafting
- **Complex Crafting System**: 
  - Base materials + crafting stats + random properties
  - Multiple upgrade paths and enchanting options
  - Rare material drops from higher-tier raids

---

## 🏗️ Technical Foundation

### Client-Server Architecture
- **Client**: Unity 3D (Isometric view)
- **Server**: C++ backend
- **Network**: Real-time synchronization for multiplayer raids
- **Offline Support**: Solo content can be played offline

### Performance Targets
- **Target FPS**: 60 FPS on mobile, 120 FPS on PC
- **Platform Optimization**: Scalable graphics settings
- **Memory Usage**: <2GB on mobile devices

---

## 📊 Development Roadmap

### Phase 1: Foundation (MVP) - 3 months
- Core raid system (5-man solo)
- Basic party recruitment
- 4-6 starter raids
- Simple equipment system

### Phase 2: Core Features - 2 months
- Complex crafting system
- Party member progression
- 8-12 additional raids
- Daily reset system

### Phase 3: Polish & Expansion - 2 months
- 10-man cooperative raids
- Advanced UI/UX polish
- Balance testing
- Launch preparation

---

## ❓ Pending Decisions

**Need Details For:**

1. **Party Member Recruitment Limits**:
   - How many party members can player recruit total?
   - Recruitment cost system (gold, tokens, etc.)?
   - Dismissal mechanics?

2. **Recruitment Randomness**:
   - Random stats on recruitment?
   - Random skills/abilities?
   - Rarity tiers for recruits?

3. **Raid Content Planning**:
   - How many raids per difficulty tier for launch?
   - Boss mechanics complexity level?
   - Material reward distribution?

---

*Document will be updated as decisions are finalized.*

## 📋 Document References

### Related Documents:
- `00-SHARED/Project-Overview.md` - Technical stack and development timeline
- `01-DESIGN-AGENT/Feature-Specifications/Party-Recruitment-System.md` - Core game system specification
- `02-SERVER-AGENT/` - Server architecture must support game requirements (future)
- `03-CLIENT-AGENT/` - Client implementation must match game design (future)

### Dependencies:
- **Requires**: `Project-Overview.md` for technical constraints and project scope
- **Blocks**: All feature specifications and technical implementations depend on this GDD

### Cross-Agent Impact:
- **DESIGN**: Foundation for all feature specifications and game balance decisions
- **SERVER**: Defines server requirements, player data models, and API needs
- **CLIENT**: Specifies Unity implementation requirements, UI systems, and user experience
- **LEAD**: Provides scope for integration planning and technical architecture decisions
