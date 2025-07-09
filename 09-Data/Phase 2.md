# Phase 2: 심층 분석 및 패턴 추출 파이프라인 개요

## 1. 개요 (Overview) 📈

**목표:** Phase 1에서 구축한 데이터 파이프라인을 기반으로, 게임 내 다양한 요소들 간의 **상관관계와 인과관계**를 분석한다. '무엇'이 일어났는지를 넘어 '왜' 그랬는지를 설명하고, 사용자의 실력 향상에 직접적으로 도움이 될 **의미 있는 패턴을 자동으로 추출**하는 것을 목표로 한다.

## 2. 핵심 기능 명세 (Core Feature Specifications)

- **타임라인 데이터 처리:** `MatchTimelineDTO`를 파싱하여 아이템 구매 순서, 스킬 빌드 순서 등 시간의 흐름에 따른 데이터를 추출한다.
    
- **상관관계 분석:** 특정 행동(예: '첫 아이템으로 특정 아이템 구매')과 결과(예: '승률', '15분 골드') 사이의 통계적 상관관계를 계산한다.
    
- **패턴 추출 및 저장:** 통계적으로 유의미하다고 판단된 패턴들(예: "제이스가 '월식'을 첫 코어로 갔을 때 승률 12% 상승")을 별도의 DB 테이블에 저장한다.
    
- **복합 조건 질의(Query):** "이즈리얼이 '삼위일체'와 '무라마나'를 완성했을 때의 평균 DPM"과 같이, 여러 조건이 결합된 복합적인 통계 조회가 가능해진다.
    

## 3. 주요 구현 과업 (Key Implementation Tasks)

1. **Develop High-Performance C++ Engine:**

    - Develop the C++ batch processing tool to perform heavy data aggregation from the raw `matches` and `match_events` tables collected in Phase 1.
    - This engine will be responsible for parsing raw data and calculating complex, large-scale statistics efficiently.

2. **Python for Higher-Level Analysis:**

    - Use Python with `Pandas`, `SciPy`, and `Statsmodels` on the **pre-aggregated data** produced by the C++ engine.
    - This allows for rapid development of the final correlation analysis, feature engineering, and pattern discovery logic without being bottlenecked by raw data processing.

3. **패턴 DB 스키마 설계:**
    
    - 아이템 빌드별 통계를 위한 `item_build_stats` 테이블, 발견된 패턴을 저장할 `discovered_patterns` 테이블 등을 새롭게 설계한다. The C++ engine will populate these tables.
        
4. **분석 파이프라인 통합:**
    
    - The daily pipeline (`run_daily_pipeline.py`) will be updated to orchestrate the execution of the C++ engine first, and then run subsequent Python scripts for higher-level analysis.
        

## 4. 분석할 패턴 예시

- **아이템 빌드 경로 vs 승률:** 특정 1, 2, 3코어 아이템 순서에 따른 챔피언별 승률 및 KDA 변화.
    
- **스킬 마스터 순서 vs DPM:** Q, W, E 선마 순서에 따른 분당 데미지 효율.
    
- **첫 오브젝트 선택 vs 승률:** 첫 용과 첫 전령 중 무엇을 획득했는지가 게임 승패에 미치는 영향.
    
- **시야 점수와 생존율의 관계:** 포지션별로 분당 시야 점수와 평균 데스 수의 상관관계.