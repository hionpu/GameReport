---
tags:
  - agent/design
  - agent/client
  - agent/server
  - system/party
  - system/social
  - system/progression
  - phase/design
  - doc/spec
  - tech/ui
  - tech/networking
  - tech/database
  - priority/high
  - status/planning
references:
  - "[[01-Main-GDD]]"
  - "[[01-Project-Overview]]"
---

# Party Member & Recruitment System Specification

## 📋 Document Information
- **Version**: 1.1
- **Last Updated**: 2025-05-27
- **Related**: Main-GDD.md

---

## 🎭 Class & Talent System

### Core Structure
- **Classes**: Each recruit belongs to a specific class
- **Talents**: Each class has 3 different talent specializations
- **Role Flexibility**: Talents determine actual combat role, not class

### Example Class Designs

#### Warrior Class
- **Talent 1: Guardian** (Tank) - Heavy armor, taunts, damage mitigation
- **Talent 2: Berserker** (DPS) - High damage, rage mechanics
- **Talent 3: Weaponmaster** (DPS) - Weapon expertise, combat techniques

#### Priest Class  
- **Talent 1: Holy** (Healer) - Direct heals, buffs
- **Talent 2: Discipline** (Support/Healer) - Shields, hybrid healing
- **Talent 3: Shadow** (DPS) - Dark magic, damage over time

#### Rogue Class
- **Talent 1: Assassination** (DPS) - Stealth, critical strikes
- **Talent 2: Combat** (DPS) - Dual wielding, sustained damage  
- **Talent 3: Subtlety** (DPS) - Stealth, utility, crowd control

#### Paladin Class
- **Talent 1: Protection** (Tank) - Holy defense, group protection
- **Talent 2: Retribution** (DPS) - Holy damage, two-handed weapons
- **Talent 3: Holy** (Healer) - Divine healing, group support

#### Mage Class
- **Talent 1: Fire** (DPS) - Burst damage, area effects
- **Talent 2: Frost** (DPS/Control) - Crowd control, damage over time
- **Talent 3: Arcane** (DPS) - High damage, mana management

#### Hunter Class
- **Talent 1: Beast Master** (DPS/Pet) - Animal companions, pet management
- **Talent 2: Marksman** (Ranged DPS) - Bow/crossbow expertise, long-range damage  
- **Talent 3: Survival** (DPS/Utility) - Traps, environmental damage, group utility

---

## 🎲 Recruitment System

### Randomness Elements

#### Available Class Pool
- **Daily Rotation**: 3-4 different classes available each day at tavern
- **Class Rarity**: Some classes appear less frequently
- **Unpredictable**: Cannot guarantee specific class availability

#### Individual Recruit Properties
- **Random Stats**: Each recruit has randomized stats within class ranges
- **Random Equipment**: Starting gear varies per recruit
- **Random Appearance**: Visual customization for personality

### Example Recruitment Scenario
```
Day 1 Tavern:
- Warrior (Available: 2 recruits, different stats)
- Priest (Available: 1 recruit)  
- Rogue (Available: 3 recruits, different stats)

Day 2 Tavern:
- Mage (Available: 1 recruit)
- Paladin (Available: 2 recruits)
- Warrior (Available: 1 recruit)
- Rogue (Available: 2 recruits)
```

---

## 📊 Progression System

### Individual Member Development
- **Level Progression**: Simple XP → Level → Stat increases
- **Equipment Slots**: Each member can equip:
  - Weapon (class-specific)
  - Armor (role-appropriate)  
  - Accessories (rings, trinkets)
- **Talent Selection**: Player chooses talent for each recruited member

### Equipment Independence
- **Separate Gear**: Each party member has individual equipment
- **Class Restrictions**: Warriors can't use mage staves, etc.
- **Role Optimization**: Tank gear different from DPS gear

---

## 🔒 Recruitment Limits

### Combined Limitation System

#### Roster Size Limit
- **Maximum Party Members**: 20 total recruited members
- **Dismissal Required**: Must dismiss old members for new recruits
- **Dismissal Cost**: Small gold fee to prevent spam recruiting

#### Daily Recruitment Limit  
- **Daily Attempts**: 3 recruitment attempts per day
- **Selective Choice**: Must choose wisely from available pool
- **Reset Timer**: Resets at daily reset (midnight)

#### Resource Cost
- **Gold Cost**: Each recruitment requires gold payment
- **Scaling Cost**: Higher level recruits cost more gold
- **Premium Currency**: Special recruits may require gems/tokens

#### Recruitment Tokens
- **Token System**: Earn tokens from completing raids
- **Quality Tiers**: Better tokens = access to rare classes
- **Token Types**: 
  - Common Token (basic classes)
  - Rare Token (uncommon classes)  
  - Epic Token (rare classes with better stats)

---

## 🎯 Strategic Implications

### Team Composition Planning
- **Role Requirements**: Need 1 Tank, 2 Healers, 2 DPS for raids
- **Class Diversity**: Different classes offer different tactical options
- **Talent Synergy**: Some talent combinations work better together

### Daily Decision Making
- **Recruitment Planning**: Check tavern daily for needed classes
- **Resource Management**: Balance gold/tokens for optimal recruits
- **Long-term Strategy**: Build roster depth for different raid challenges

---

## ✅ Design Decisions

### Total Class Count at Launch
**Decision**: Launch with **6 core classes**
- **Rationale**: Provides sufficient variety without overwhelming new players
- **Classes**: Warrior, Priest, Rogue, Paladin, Mage, Hunter
- **Post-Launch**: Additional classes (Warlock, Druid, Death Knight) in future updates
- **Balance**: Each class fills distinct tactical roles with 3 talent variations

### Talent Respecialization System
**Decision**: **Limited Respec Available**
- **Respec Tokens**: Earned through high-level raid completions (rare resource)
- **Gold Alternative**: High gold cost (scales with member level) for immediate respec
- **Free Respec**: One free respec per member at level milestones (10, 20, 30)
- **Strategic Value**: Makes initial talent choice meaningful while allowing flexibility

### Member Relationship Mechanics  
**Decision**: **Loyalty System Implementation**
- **Loyalty Points**: Members gain loyalty through successful raids and equipment upgrades
- **Loyalty Benefits**: 
  - High loyalty = stat bonuses and special combat abilities
  - Low loyalty = chance to leave party or underperform
- **Loyalty Actions**: 
  - Positive: Winning raids, getting equipment, being used in party
  - Negative: Losing raids, being benched long-term, dismissing friends
- **Friendship Bonds**: Some recruits arrive as pairs/groups with loyalty connections

### Rare Recruit Variants
**Decision**: **Named Elite Recruits System**
- **Elite Spawns**: 5% chance for named recruits with unique appearances
- **Elite Benefits**:
  - +2 stat points above normal range
  - Unique visual customization (special armor, colors, effects)
  - Memorable names and backstories
  - Higher loyalty starting value
- **Elite Cost**: Require Epic Recruitment Tokens + additional gold
- **Collection Incentive**: Achievement system for recruiting specific named elites

### Loyalty System Details
- **Loyalty Scale**: 0-100 points
- **Loyalty Effects**:
  - 0-25: -10% stats, 5% chance to refuse orders
  - 26-50: Normal performance
  - 51-75: +5% stats, improved combat AI
  - 76-100: +10% stats, chance for bonus actions in combat

### Elite Recruit Examples
- **"Thorgar the Unbreakable"** (Warrior/Guardian) - Legendary tank with unique shield
- **"Whisper"** (Rogue/Assassination) - Master assassin with shadow effects
- **"Lightbringer Sera"** (Priest/Holy) - Divine healer with golden aura
- **"Frostweaver Kael"** (Mage/Frost) - Ice mage with crystalline appearance

---