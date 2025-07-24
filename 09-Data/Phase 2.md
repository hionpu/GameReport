# Phase 2: 심층 분석 및 패턴 추출 파이프라인 개요

## 1. 개요 (Overview) 📈

**목표:** Phase 1에서 구축한 데이터 파이프라인을 기반으로, 게임 내 다양한 요소들 간의 **상관관계와 인과관계**를 분석한다. '무엇'이 일어났는지를 넘어 '왜' 그랬는지를 설명하고, 사용자의 실력 향상에 직접적으로 도움이 될 **의미 있는 패턴을 자동으로 추출**하는 것을 목표로 한다.

## 2. 핵심 기능 명세 (Core Feature Specifications)

- **타임라인 데이터 처리:** `MatchTimelineDTO`를 파싱하여 아이템 구매 순서, 스킬 빌드 순서 등 시간의 흐름에 따른 데이터를 추출한다.
    
- **상관관계 분석:** 특정 행동(예: '첫 아이템으로 특정 아이템 구매')과 결과(예: '승률', '15분 골드') 사이의 통계적 상관관계를 계산한다.
    
- **패턴 추출 및 저장:** 통계적으로 유의미하다고 판단된 패턴들(예: "제이스가 '월식'을 첫 코어로 갔을 때 승률 12% 상승")을 별도의 DB 테이블에 저장한다.
    
- **복합 조건 질의(Query):** "이즈리얼이 '삼위일체'와 '무라마나'를 완성했을 때의 평균 DPM"과 같이, 여러 조건이 결합된 복합적인 통계 조회가 가능해진다.
    

## 3. 주요 구현 과업 (Key Implementation Tasks)

1. **Develop High-Performance Rust Engine:**

    - Develop the Rust batch processing tool to perform heavy data aggregation from the raw `matches` and `match_events` tables collected in Phase 1.
    - This engine will be responsible for parsing raw data and calculating complex, large-scale statistics efficiently and safely.

2. **Python for Higher-Level Analysis:**

    - Use Python with `Pandas`, `SciPy`, and `Statsmodels` on the **pre-aggregated data** produced by the Rust engine.
    - This allows for rapid development of the final correlation analysis, feature engineering, and pattern discovery logic without being bottlenecked by raw data processing.

3. **패턴 DB 스키마 설계:**
    
    - 아이템 빌드별 통계를 위한 `item_build_stats` 테이블, 발견된 패턴을 저장할 `discovered_patterns` 테이블 등을 새롭게 설계한다. The Rust engine will populate these tables.
        
4. **분석 파이프라인 통합:**
    
    - The daily pipeline (`run_daily_pipeline.py`) will be updated to orchestrate the execution of the Rust engine first, and then run subsequent Python scripts for higher-level analysis.
        

## 4. 분석할 패턴 예시

- **아이템 빌드 경로 vs 승률:** 특정 1, 2, 3코어 아이템 순서에 따른 챔피언별 승률 및 KDA 변화.
    
- **스킬 마스터 순서 vs DPM:** Q, W, E 선마 순서에 따른 분당 데미지 효율.
    
- **첫 오브젝트 선택 vs 승률:** 첫 용과 첫 전령 중 무엇을 획득했는지가 게임 승패에 미치는 영향.
    
- **시야 점수와 생존율의 관계:** 포지션별로 분당 시야 점수와 평균 데스 수의 상관관계.


## 4. Rust Engine Detailed Specifications

### Core Rust Engine Architecture

```rust
// High-level Rust engine structure
mod gaming_analytics {
    use polars::prelude::*;
    
    pub struct BatchProcessor {
        aggregated_stats: DataFrame,
        discovered_patterns: Vec<Pattern>,
    }

    impl BatchProcessor {
        // Main batch processing functions
        pub fn process_match_data_batch(&mut self, matches: DataFrame) -> Result<()>;
        pub fn process_timeline_events_batch(&mut self, events: DataFrame) -> Result<()>;
        pub fn calculate_correlations(&self, analysis_type: &str) -> Result<DataFrame>;
        
        // Pattern discovery functions
        pub fn discover_item_build_patterns(&self) -> Result<Vec<Pattern>>;
        pub fn discover_skill_order_patterns(&self) -> Result<Vec<Pattern>>;
        pub fn discover_objective_patterns(&self) -> Result<Vec<Pattern>>;
    }
}
```

### Rust Engine Responsibilities in Phase 2

**1. Timeline Data Processing:**
- Parse `MatchTimelineDTO` events at scale using `serde_json`.
- Extract item purchase sequences and timing.
- Analyze skill level-up patterns.
- Process objective acquisition timing and team coordination.

**2. Statistical Correlation Engine:**
- Calculate Pearson/Spearman correlations between game variables.
- Identify statistically significant patterns (p-value < 0.05).
- Build multi-dimensional correlation matrices.
- Perform regression analysis for predictive modeling.

**3. Pattern Discovery Algorithms:**
- **Item Build Analysis:** Sequence mining for optimal item paths.
- **Skill Order Analysis:** N-gram analysis of skill leveling patterns.
- **Objective Priority Analysis:** Decision tree analysis for objective choices.
- **Team Coordination Analysis:** Synchronization analysis of team actions.

### Performance Benchmarks & Specifications

**Processing Capacity:**
- Timeline Events: 100,000+ events per second.
- Match Data: 20,000+ matches per batch.
- Correlation Analysis: 5,000+ variable pairs per minute.
- Pattern Discovery: Complete analysis of 100K+ matches in under 5 minutes.

**Memory Optimization:**
- Zero-copy data manipulation with Apache Arrow & Polars.
- Efficient memory management through Rust's ownership model.
- Streaming algorithms for large dataset processing.

### Integration with Python Analysis Layer

```python
# Python orchestration layer
import rust_gaming_analytics as rga

class AnalysisPipeline:
    def __init__(self):
        self.rust_engine = rga.BatchProcessor()
        
    def run_phase2_analysis(self):
        # Step 1: Rust engine processes raw data
        self.rust_engine.process_raw_matches()
        
        # Step 2: Python performs high-level analysis on dataframes from Rust
        aggregated_data_arrow = self.rust_engine.get_aggregated_stats()
        df = polars.from_arrow(aggregated_data_arrow)
        patterns = self.analyze_patterns_with_scipy(df)
        
        # Step 3: Store results in database
        self.store_discovered_patterns(patterns)
```