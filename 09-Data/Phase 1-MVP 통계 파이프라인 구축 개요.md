# Phase 1: MVP 통계 파이프라인 구축 개요

## 1. 개요 (Overview) 🎯

**목표:** Riot API를 통해 리그 오브 레전드 경기 데이터를 자동으로 수집하고, 분석 가능한 통계 형태로 가공하여 데이터베이스에 저장한다. 이 파이프라인은 향후 모든 심층 분석 기능의 안정적인 데이터 기반을 제공하는 것을 최우선으로 한다.

---

## 2. 핵심 기능 명세 (Core Feature Specifications)

- **데이터 수집기 (Data Collector):** Riot `League API`와 `Match API`를 호출하여 지정된 조건에 맞는 유저 및 경기 데이터를 가져온다.
    
- **데이터 처리기 (Data Processor):** 수집된 원본 경기 데이터(JSON)를 파싱하여, 개별 경기에 대한 주요 통계(KDA, CS, DPM 등)를 추출한다.
    
- **데이터 집계기 (Data Aggregator):** 개별 경기 데이터를 바탕으로, '챔피언별/포지션별/티어별' 평균 통계를 계산한다.
    
- **데이터 저장소 (Data Storage):** 처리 및 집계된 데이터를 PostgreSQL DB에 저장한다.
    
- **자동화 스케줄러 (Scheduler):** 위 모든 과정을 매일 정해진 시간에 자동으로 실행한다.
    
- **API 서버 (API Server):** 외부(클라이언트 앱)에서 집계된 통계 데이터를 요청할 수 있는 간단한 API 엔드포인트를 제공한다.
    

---

## 3. 기술 스택 (Technology Stack) 🛠️

- **언어:** **Python 3.9+**
    
- **백엔드 프레임워크:** **FastAPI** (빠르고 현대적인 API 서버 구축)
    
- **데이터베이스:** **PostgreSQL** (Supabase를 통해 관리)
    
- **데이터 처리:** **Pandas**
    
- **API 연동:** **Requests**
    
- **자동화:** **Cron** (리눅스 서버) 또는 **Supabase Cron Jobs / GitHub Actions** (클라우드 기반 스케줄링)
    

---

## 4. 데이터베이스 스키마 (Database Schema) 🗄️

- **`tracked_users`**: 데이터 수집 대상 유저 목록
    
    - `puuid` (PK), `summoner_id`, `tier`, `last_updated`
        
- **`matches`**: 수집된 개별 경기 데이터
    
    - `match_id` (PK), `user_puuid` (FK), `champion_id`, `position`, `win`, `kills`, `deaths`, `assists`, `cs_per_min`, `damage_per_min`, `game_version`, `created_at`
        
- **`aggregated_stats`**: 집계된 통계 데이터
    
    - `id` (PK), `game_version`, `tier_group`, `champion_id`, `position`, `total_games`, `win_rate`, `avg_kda`, `avg_cs_per_min`, `avg_damage_per_min`, `updated_at`
        

---

## 5. 구현 순서 (Implementation Steps) 🚶‍♂️

1. **환경 설정:** Supabase 프로젝트를 생성하고 위 스키마에 따라 DB 테이블을 정의한다.
    
2. **시드 유저 확보:** `League API`를 호출하여 최상위 티어 유저 목록을 가져오는 Python 스크립트를 작성한다. (`1_get_seed_users.py`)
    
3. **경기 데이터 수집:** 특정 유저의 `puuid`를 받아 `Match API`로 최근 경기 ID 목록과 각 경기의 상세 데이터를 가져오는 스크립트를 작성한다. (`2_collect_matches.py`)
    
4. **데이터 처리 및 저장:** 3번에서 얻은 경기 데이터를 Pandas로 가공하여 주요 통계를 계산하고, `matches` 테이블에 저장하는 로직을 구현한다. (`3_process_and_save.py`)
    
5. **통계 집계:** `matches` 테이블의 데이터를 읽어 `aggregated_stats` 테이블을 계산하고 업데이트하는 스크립트를 작성한다. (`4_aggregate_stats.py`)
    
6. **파이프라인 통합:** 1~5번 스크립트를 순차적으로 실행하는 메인 스크립트를 작성한다. (`run_daily_pipeline.py`)
    
7. **자동화 설정:** `cron`이나 다른 스케줄러를 사용하여 `run_daily_pipeline.py`가 매일 새벽 시간에 실행되도록 설정한다.
    
8. **API 서버 구축:** FastAPI를 사용하여 `aggregated_stats` 테이블의 데이터를 조회할 수 있는 간단한 GET 엔드포인트를 구현한다.
    

---

## 6. API 엔드포인트 명세 (API Endpoint Specification) 📡

- **Endpoint:** `GET /stats/champion/{champion_id}`
    
- **Description:** 특정 챔피언의 집계된 통계를 조회한다.
    
- **Query Parameters:**
    
    - `position` (string, optional): `TOP`, `JUNGLE` 등 포지션
        
    - `tier_group` (string, optional): `HIGH_ELO`, `MID_ELO` 등 티어 그룹
        
- **Success Response (200 OK):**
    
    JSON
    
    ```
    {
      "champion_id": 103,
      "position": "MIDDLE",
      "tier_group": "HIGH_ELO",
      "total_games": 1250,
      "win_rate": 0.52,
      "avg_kda": 4.1,
      "avg_cs_per_min": 8.5,
      "avg_damage_per_min": 650.0
    }
    ```


# 상세

# Phase 1: MVP 파이프라인 상세 구현 계획 (v1.1)

## 1. 개요 (Overview) 🎯

- **목표:** Riot API를 통해 리그 오브 레전드 경기 데이터를 자동으로 수집하고, 분석 가능한 통계 형태로 가공하여 데이터베이스에 저장한다. 이 파이프라인은 향후 모든 심층 분석 기능의 안정적인 데이터 기반을 제공하는 것을 최우선으로 한다.
    

## 2. 핵심 데이터 수집 전략: '1 Match, 10-Player Data'

- 한 번의 `Match API` 호출로 얻게 되는 **10명 참가자 모두의 데이터를 활용**하여 API 효율을 극대화한다.
    
- 경기 데이터에서 발견된 새로운 유저들을 지속적으로 추적 대상에 추가하는 **'Snowball' 방식으로 데이터베이스를 유기적으로 확장**한다.
    

---

## 3. 사용할 Riot API 명세

데이터 수집은 주로 Riot `LEAGUE-V4`, `SUMMONER-V4`, `MATCH-V5` API를 연쇄적으로 호출하여 이루어집니다.

#### A. 시드 유저 확보 (LEAGUE-V4)

최상위 티어 리그에 속한 소환사 목록을 가져와 데이터 수집의 시작점으로 삼습니다.

- `lol/league/v4/challengerleagues/by-queue/{queue}`
    
- `lol/league/v4/grandmasterleagues/by-queue/{queue}`
    
- `lol/league/v4/masterleagues/by-queue/{queue}`
    
- **결과물:** 소환사 목록 (`summonerId` 포함)
    

#### B. 유저 고유 식별자 변환 (SUMMONER-V4)

API V5에서 핵심 키로 사용되는 `PUUID`를 얻기 위해 사용합니다.

- `lol/summoner/v4/summoners/{encryptedSummonerId}`
    
- **결과물:** `puuid`, `accountId` 등
    

#### C. 경기 ID 목록 수집 (MATCH-V5)

특정 유저의 최근 경기 기록 ID 목록을 가져옵니다.

- `lol/match/v5/matches/by-puuid/{puuid}/ids`
    
- **결과물:** 경기 ID(`matchId`) 리스트
    

#### D. 경기 상세 데이터 수집 (MATCH-V5)

각 경기 ID를 사용하여 상세한 경기 데이터를 가져옵니다.

- `lol/match/v5/matches/{matchId}`
    
- **결과물:** 해당 경기의 모든 정보가 담긴 `MatchDTO` (JSON)
    

**※ 중요:** 모든 API 호출 시, Riot API가 제공하는 **요청 제한(Rate Limit)**을 반드시 준수해야 합니다. 요청 사이에 적절한 `time.sleep()`을 추가하는 로직이 필요합니다.

---

## 4. 원본 데이터(Raw Data) 저장 여부 결정

**결론: 초기에는 저장하지 않고, 장기적으로는 S3와 같은 오브젝트 스토리지에 저장하는 것을 추천합니다.**

- **초기 MVP 단계:**
    
    - **전략:** 저장하지 않는다.
        
    - **이유:** 파이프라인을 최대한 빠르고 단순하게 구축하는 데 집중합니다. 원본 데이터를 저장하고 관리하는 것은 추가적인 복잡성을 야기합니다. 가공된 데이터만 DB에 저장하여 빠르게 핵심 기능을 구현합니다.
        
- **장기적인 관점:**
    
    - **전략:** AWS S3, Google Cloud Storage 등 저렴한 오브젝트 스토리지에 `matchId`를 파일명으로 하여 JSON 그대로 저장한다.
        
    - **이유:**
        
        1. **재처리 가능성:** 데이터 처리 로직에 버그를 발견하거나, 나중에 새로운 통계(예: 시야 점수)를 추출하고 싶을 때, API를 다시 호출할 필요 없이 저장된 원본 데이터로 모든 과거 데이터를 재처리할 수 있습니다.
            
        2. **머신러닝 데이터셋:** 향후 3단계에서 ML 모델을 학습시킬 때, 가공되지 않은 풍부한 정보가 담긴 원본 데이터는 매우 귀중한 자산이 됩니다.

---

## 5. 데이터베이스 테이블 상세 설계 (v1.1)

#### `matches` (개별 경기 데이터)

#### `matches` (개별 경기 데이터)

수집된 모든 경기의 핵심 정보를 저장합니다.

| 컬럼명              | 데이터 타입         | 설명                                      |
| ---------------- | -------------- | --------------------------------------- |
| `match_id`       | `VARCHAR(32)`  | PK, 경기 고유 ID                            |
| `user_puuid`     | `VARCHAR(128)` | **해당 경기 기록의 주체**가 되는 유저의 PUUID (인덱스 필요) |
| `game_version`   | `VARCHAR(32)`  | 게임 버전 (예: "12.1.123") (인덱스 필요)          |
| `game_duration`  | `INT`          | 게임 시간 (초)                               |
| `champion_id`    | `INT`          | 플레이한 챔피언 ID (인덱스 필요)                    |
| `position`       | `VARCHAR(16)`  | 플레이한 포지션 (TOP, JUNGLE 등)                |
| `win`            | `BOOLEAN`      | 승리 여부                                   |
| `kills`          | `INT`          | 킬                                       |
| `deaths`         | `INT`          | 데스                                      |
| `assists`        | `INT`          | 어시스트                                    |
| `cs_per_min`     | `FLOAT`        | 분당 CS                                   |
| `damage_per_min` | `FLOAT`        | 분당 데미지                                  |
| `gold_per_min`   | `FLOAT`        | 분당 골드                                   |
| `created_at`     | `TIMESTAMP`    | 데이터 생성 시각                               |

Sheets로 내보내기

#### `aggregated_stats` (집계 통계 데이터)

`matches` 테이블을 기반으로 주기적으로 계산되어 저장됩니다. API가 직접 조회하는 테이블입니다.

|컬럼명|데이터 타입|설명|
|---|---|---|
|`id`|`SERIAL`|PK, 자동 증가|
|`game_version`|`VARCHAR(32)`|기준 게임 버전|
|`tier_group`|`VARCHAR(16)`|티어 그룹 (예: HIGH_ELO, MID_ELO)|
|`champion_id`|`INT`|기준 챔피언 ID|
|`position`|`VARCHAR(16)`|기준 포지션|
|`total_games`|`INT`|총 게임 수|
|`win_rate`|`FLOAT`|승률|
|`avg_kda`|`FLOAT`|평균 KDA|
|`avg_cs_per_min`|`FLOAT`|평균 분당 CS|
|`avg_damage_per_min`|`FLOAT`|평균 분당 데미지|
|`updated_at`|`TIMESTAMP`|마지막 업데이트 시각|
(참고: `tier_group`, `champion_id`, `position` 등은 복합 인덱스(composite index)를 설정하면 조회 성능이 향상됩니다.)
#### `tracked_users` (수집 대상 유저 목록 - 역할 변경)

- 초기 시드 유저 목록뿐만 아니라, 경기에서 **새롭게 발견된 유저들을 저장하고 관리**하는 역할이 추가됨.
    

|컬럼명|데이터 타입|설명|
|---|---|---|
|`puuid`|`VARCHAR(128)`|PK, 유저 고유 ID|
|`tier`|`VARCHAR(16)`|티어 정보|
|`is_seed`|`BOOLEAN`|초기 시드 유저인지 여부|
|`last_crawled_at`|`TIMESTAMP`|마지막으로 이 유저의 경기 목록을 수집한 시각|

Sheets로 내보내기

---

## 6. 파이썬 스크립트 구조 및 주요 함수 (v1.1)

프로젝트를 모듈화하여 관리의 용이성을 높입니다.

```
lol_pipeline/
├── main.py                 # 전체 파이프라인을 실행하는 메인 스크립트
├── api/                    # FastAPI 서버 코드
│   └── server.py
├── collector/              # 데이터 수집 관련 모듈
│   └── riot_api.py
├── processor/              # 데이터 처리 및 집계 모듈
│   ├── process_match.py
│   └── aggregate_stats.py
├── database/               # 데이터베이스 연결 및 쿼리
│   └── db_handler.py
└── config.py               # API 키, DB 접속 정보 등 설정 파일
```

### 주요 함수 명세

#### `collector/riot_api.py`

- `get_high_elo_summoner_ids()`: LEAGUE API를 호출하여 최상위 티어 소환사 ID 목록 반환.
    
- `get_puuid_by_summoner_id(summoner_id)`: SUMMONER API를 호출하여 `summonerId`로 `puuid` 반환.
    
- `get_match_ids_by_puuid(puuid)`: MATCH API를 호출하여 `puuid`로 최근 경기 ID 리스트 반환.
    
- `get_match_data_by_match_id(match_id)`: MATCH API를 호출하여 `matchId`로 상세 경기 데이터 JSON 반환.
#### `processor/process_match.py`

- **`process_all_participants_in_match(match_json)`**:
    
    - `match_json`을 입력받아 **`participants` 리스트를 순회**.
        
    - **10명의 참가자 각각에 대해** KDA, 분당 CS, 분당 데미지 등 주요 KPI를 계산.
        
    - `matches` 테이블에 저장될 수 있는 형태의 **딕셔너리 리스트 (10개)**를 반환.
        
#### `processor/aggregate_stats.py`

- `update_aggregated_stats()`:
    
    - DB에서 `matches` 테이블 데이터를 읽어옴 (Pandas DataFrame 활용).
        
    - `groupby`를 사용하여 `game_version`, `champion_id`, `position` 별로 통계 집계.
        
    - `aggregated_stats` 테이블에 최신 통계 데이터를 업데이트(UPSERT)함.

#### `database/db_handler.py`

- **`bulk_insert_matches(match_data_list)`**:
    
    - `process_all_participants_in_match`가 반환한 10개의 경기 기록 리스트를 받아 `matches` 테이블에 한 번에 삽입(Bulk Insert)하여 DB 부하를 줄임.
        
- **`get_users_to_crawl(limit)`**:
    
    - `tracked_users` 테이블에서 마지막 수집 시간(`last_crawled_at`)이 오래된 순서로 지정된 수만큼 유저를 가져옴.
        
- **`add_new_puuids_to_tracked_users(puuid_list)`**:
    
    - 새로 발견된 puuid들을 `tracked_users` 테이블에 추가함. (이미 존재하면 무시)
        

#### `main.py`

- **`run_daily_pipeline()`**:
    
    1. `db_handler.get_users_to_crawl()`로 오늘 처리할 유저 목록을 가져온다.
        
    2. 각 유저에 대해 `riot_api.get_match_ids_by_puuid()`로 경기 ID 목록을 가져온다.
        
    3. DB를 확인하여 이미 `matches` 테이블에 저장된 `match_id`는 처리 대상에서 제외한다.
        
    4. 새로운 `match_id`에 대해 `riot_api.get_match_data_by_match_id()`를 호출한다.
        
    5. `processor.process_all_participants_in_match()`를 호출하여 **10명의 데이터를 한 번에 가공**한다.
        
    6. `db_handler.bulk_insert_matches()`로 **10개의 경기 기록을 DB에 저장**한다.
        
    7. 해당 경기의 참가자 목록에서 새로운 puuid들을 `db_handler.add_new_puuids_to_tracked_users()`로 `tracked_users` 목록에 추가한다. (Snowball Effect)
        
    8. 모든 작업이 끝나면 `processor.update_aggregated_stats()`를 호출하여 전체 통계를 업데이트한다.