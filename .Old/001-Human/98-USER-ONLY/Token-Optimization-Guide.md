---
tags:
  - agent/all
  - system/config
  - phase/optimization
  - doc/guide
  - priority/high
  - status/complete
---

# RaidMaster Token Optimization Guide

## Language Usage Strategy for Maximum Token Efficiency

### Document Type Language Selection

#### 1. Design Documents (English Primary)
- **Rationale**: English provides better token efficiency for technical concepts
- **Token Saving**: Use concise technical terminology
```markdown
# Combat System
## Core Mechanics
- Combo System: consecutive attacks increase damage multiplier
- Dodge Roll: 0.5s iframe, 2s cooldown
- Skill Tree: 3 branches (attack/defense/utility)
```

#### 2. Technical Documents (English Only)
- **Rationale**: APIs, code, architecture are most token-efficient in English
```markdown
# Server Architecture

## API Endpoints
- POST /api/player/inventory/add
- GET /api/player/inventory/list
- PUT /api/player/inventory/move

## Database Schema
```sql
CREATE TABLE inventory_items (
    player_id INT,
    item_id INT,
    slot_index INT,
    quantity INT
);
```

**Notes**: slot_index range 0-47, no duplicates allowed
```

#### 3. Integration Documents (English Primary)
```markdown
# Inventory Integration Specification

## Game Design Requirements
- Max slots: 48 (8x6 grid)
- Item types: equipment, consumable, material, misc

## Technical Implementation
```cpp
enum class ItemType {
    EQUIPMENT = 1,
    CONSUMABLE = 2,
    MATERIAL = 3,
    MISC = 4
};
```

## UI/UX Guidelines
- Drag feedback: visual highlight + cursor change
- Empty slot display: gray border, semi-transparent
- Item tooltip: detailed info on hover
```

### Token-Efficient Writing Patterns

#### 1. Abbreviations and Symbols
```markdown
# Efficient (token-optimized)
- HP: 100 (base) → 150 (max)
- MP: 50 (base) → 100 (max)
- ATK: 10~15 (weapon dependent)
- DEF: 5 (armor bonus)

# Inefficient (token-heavy)
- Health Points: increases from base 100 to maximum 150
- Mana Points: increases from base 50 to maximum 100
- Attack Power: ranges from 10 to 15 depending on weapon
- Defense Power: receives 5 additional points from armor
```

#### 2. Structured Information
```markdown
# Player Stats Structure

| Stat | Base | Max | Growth |
|------|------|-----|--------|
| HP   | 100  | 999 | +10/lv |
| MP   | 50   | 500 | +5/lv  |
| STR  | 10   | 99  | +1/lv  |
| DEX  | 10   | 99  | +1/lv  |

**Formula**: final_value = base + (level * growth)
```

#### 3. Code-Centric Documentation
```markdown
# Combat Calculation

```cpp
// Damage calculation
int calculateDamage(int attack, int defense) {
    int baseDamage = attack - defense;
    if (baseDamage < 1) baseDamage = 1;
    
    // Critical hit (10% chance, 1.5x damage)
    if (random() % 100 < 10) {
        baseDamage = static_cast<int>(baseDamage * 1.5f);
    }
    
    return baseDamage;
}
```

**Balance Notes**: min damage 1, critical rate 10%
```

### Token Usage Monitoring

#### Document Type Token Efficiency (estimated)
- **Pure English**: 1.0x (baseline)
- **English with minimal explanations**: 1.1x
- **Code-heavy documentation**: 0.8x (most efficient)
- **Mixed content**: 1.2x

#### Recommended Language Ratios by Agent
```markdown
01-DESIGN-AGENT/
├── Game-Design-Documents/     # English 90%
├── Feature-Specifications/    # English 95%
└── User-Stories/             # English 85%

02-SERVER-AGENT/
├── Architecture/             # English 98%
├── API-Specifications/       # English 99%
└── Database-Design/          # English 98%

03-CLIENT-AGENT/
├── UI-UX-Design/            # English 90%
├── Component-Architecture/   # English 95%
└── Scene-Management/        # English 95%

04-LEAD-AGENT/
├── Integration-Reviews/      # English 92%
├── Risk-Management/         # English 88%
└── Quality-Assurance/       # English 90%
```

## Practical Token Saving Tips

### 1. Priority-Based Detail Level
```markdown
# High Priority Features (detailed documentation)
- Combat System: complete specification
- Inventory System: detailed implementation guide

# Medium Priority Features (moderate detail)  
- Quest System: core functionality only
- Shop System: basic structure definition

# Low Priority Features (minimal documentation)
- Achievement System: simple overview only
- Settings Menu: reference standard UI patterns
```

### 2. Template-Based Writing
```markdown
# Feature Template (reusable)

## Overview
[1-2 sentence summary]

## Core Mechanics  
- Mechanic 1: [brief description]
- Mechanic 2: [brief description]

## Technical Requirements
- Server: [key points]
- Client: [key points]
- Database: [schema changes]

## Implementation Priority
- P0: [must have]
- P1: [should have]  
- P2: [nice to have]
```

### 3. Efficient Reference Patterns
```markdown
# Reference other documents efficiently
See Combat-System.md for damage calculations
Refer to API-Auth.md for authentication flow

# Instead of duplicating content
[Avoid copying entire sections from other documents]
```

## Expected Token Savings
- **Baseline (mixed language)**: 100% token usage
- **English-optimized**: 60-70% token usage (**30-40% savings**)
- **Code-heavy docs**: 50-60% token usage (**40-50% savings**)
- **Template-based**: 70-80% token usage (**20-30% savings**)

## Implementation Strategy
1. **Phase 1**: Convert all existing documents to English
2. **Phase 2**: Implement token-efficient writing patterns
3. **Phase 3**: Create reusable templates for common document types
4. **Phase 4**: Monitor and optimize based on actual usage patterns