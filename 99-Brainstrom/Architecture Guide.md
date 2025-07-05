이 문서는 'Daily Gaming Report Card' 프로젝트의 최종 기술 아키텍처를 정의합니다. 본 아키텍처는 **안정적인 웹 서비스**와 **복잡한 데이터 분석**의 책임을 분리하고, 향후 **다양한 게임을 지원**할 수 있는 확장성을 확보하는 것을 목표로 합니다.

---

## 🎯 1. 핵심 아키텍처

**듀얼 서버(마이크로서비스) 모델**을 채택하여 각 서비스의 역할을 명확히 구분하고, **전략 패턴(Strategy Pattern)**을 도입하여 다중 게임 지원을 위한 확장성을 확보합니다.

- **Go 웹 서버 (프론트엔드 서버)**
    
    - **역할**: 사용자 인터페이스(UI) 제공, 요청 오케스트레이션
        
    - **기술**: Go, Chi, Templ, HTMX
        
    - **책임**: 사용자 요청 처리, Python 분석 서버 호출, 응답 데이터를 HTML로 렌더링
        
- **Python 분석 서버 (백엔드 서버)**
    
    - **역할**: 데이터 수집, 심층 분석, AI 기반 리포트 생성
        
    - **기술**: Python, FastAPI, Pandas
        
    - **책임**: 게임별 API 호출, 데이터베이스 연동, 통계 분석, AI 인사이트 생성, 분석 결과를 JSON API로 제공
        

### 🌊 데이터 흐름

```
사용자 ↔ [게임 선택 & 소환사 검색] ↔ Go 웹 서버 ↔ [REST API] ↔ Python 분석 서버 ↔ [게임별 분석기] ↔ (Riot API, DB, AI API)
```

---

## 📁 2. 프로젝트 구조

두 개의 독립된 서비스는 각각의 디렉토리로 관리되며, `docker-compose`를 통해 통합 운영됩니다.

### 2.1 Go 웹 서버 (`web_server/`)

```
web_server/
├── cmd/server/
│   └── main.go                 # 애플리케이션 엔트리포인트
├── internal/
│   ├── analysis/
│   │   └── client.go           # Python 분석 서버 API 클라이언트
│   ├── handler/
│   │   └── report_handler.go   # '/{game}/report' 요청 처리 핸들러
│   └── middleware/
│       └── htmx.go
├── web/
│   ├── templates/
│   │   ├── layouts/
│   │   │   └── base.templ
│   │   ├── components/
│   │   │   ├── shared/         # 공통 UI 컴포넌트
│   │   │   └── reports/
│   │   │       ├── lol/        # 게임별 리포트 컴포넌트
│   │   │       └── valorant/
│   │   └── pages/
│   │       ├── home.templ
│   │       └── report/
│   │           └── show.templ  # 동적 리포트 결과 페이지
│   └── static/
│       ├── css/
│       └── js/
├── go.mod
└── Makefile
```

### 2.2 Python 분석 서버 (`analysis_server/`)

```
analysis_server/
├── app/
│   ├── api/v1/
│   │   └── reports.py          # /{game_name} API 엔드포인트
│   ├── schemas/
│   │   └── report.py           # Pydantic 데이터 모델
│   ├── services/
│   │   ├── base_analyzer.py    # 분석기 기본 인터페이스 (ABC)
│   │   ├── analyzers/          # 게임별 실제 분석기 구현
│   │   │   ├── lol_analyzer.py
│   │   │   └── valorant_analyzer.py
│   │   └── analysis_factory.py # 요청에 맞는 분석기를 생성하는 팩토리
│   └── db/
│       └── models.py
├── main.py                     # FastAPI 애플리케이션 시작점
├── Dockerfile
└── requirements.txt
```

### 2.2 Python 분석 서버 (`analysis_server/`)

```
analysis_server/
├── app/
│   ├── api/v1/
│   │   └── reports.py          # /{game_name} API 엔드포인트
│   ├── schemas/
│   │   └── report.py           # Pydantic 데이터 모델
│   ├── services/
│   │   ├── base_analyzer.py    # 분석기 기본 인터페이스 (ABC)
│   │   ├── analyzers/          # 게임별 실제 분석기 구현
│   │   │   ├── lol_analyzer.py
│   │   │   └── valorant_analyzer.py
│   │   └── analysis_factory.py # 요청에 맞는 분석기를 생성하는 팩토리
│   └── db/
│       └── models.py
├── main.py                     # FastAPI 애플리케이션 시작점
├── Dockerfile
└── requirements.txt
```

---

## 🛠️ 3. 핵심 구현 패턴

### 3.1 Python: 분석기 팩토리 (Strategy Pattern)

요청된 게임 이름(`game_name`)에 따라 적절한 분석기(Analyzer)를 동적으로 선택하고 반환합니다.


```python
# analysis_server/app/services/analysis_factory.py
from app.services.analyzers.lol_analyzer import LolAnalyzer
from app.services.analyzers.valorant_analyzer import ValorantAnalyzer

ANALYZERS = {"lol": LolAnalyzer, "valorant": ValorantAnalyzer}

class AnalysisFactory:
    @staticmethod
    def get_analyzer(game_name: str):
        analyzer_class = ANALYZERS.get(game_name.lower())
        if not analyzer_class:
            raise ValueError(f"'{game_name}'은(는) 지원하지 않는 게임입니다.")
        return analyzer_class()

# analysis_server/app/api/v1/reports.py
@router.post("/{game_name}")
async def create_daily_report(game_name: str, request: ReportRequest):
    try:
        analyzer = AnalysisFactory.get_analyzer(game_name)
        report = await analyzer.generate_report(summoner_name=request.summoner_name)
        return report
    except ValueError as e:
        raise HTTPException(status_code=404, detail=str(e))
```

### 3.2 Go & Templ: 동적 라우팅 및 렌더링

Chi 라우터에서 `{game}`을 URL 파라미터로 받아 처리하고, `Templ`의 `type switch`를 사용하여 게임별로 다른 UI 컴포넌트를 렌더링합니다.


```go
// web_server/internal/handler/report_handler.go
func (h *ReportHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
    game := chi.URLParam(r, "game")
    summonerName := r.FormValue("summonerName")

    // Python 클라이언트 호출 시 game 이름 전달
    reportData, err := h.AnalysisClient.GetDailyReport(r.Context(), game, summonerName)
    // ...
    h.RenderTemplate(w, "pages/report/show", reportData)
}
```

코드 스니펫

```go
// web_server/web/templates/pages/report/show.templ
templ Show(reportData any) { // reportData는 any(interface{}) 타입
    <div id="report-container">
        switch data := reportData.(type) {
        case analysis.LolReportData:
            @lol.ReportComponent(data)
        case analysis.ValorantReportData:
            @valorant.ReportComponent(data)
        default:
            <p>리포트를 표시할 수 없습니다.</p>
        }
    </div>
}
```

---

## 🔀 4. API 및 라우팅 설계

- **Go Web Server (User-Facing)**
    
    - `GET /`: 메인 페이지 (게임 선택 및 검색)
        
    - `POST /{game}/report`: 특정 게임의 분석 리포트 생성 요청 (HTMX)
        
- **Python Analysis Server (Internal API)**
    
    - `POST /api/v1/reports/{game_name}`: 특정 게임의 분석 로직을 트리거
        

---

## 🗄️ 5. 데이터베이스 스키마

여러 게임의 데이터를 유연하게 저장하기 위해 `game_type` 컬럼과 **JSONB** 타입을 활용합니다.



```sql
CREATE TABLE matches (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    match_api_id VARCHAR(255) NOT NULL,
    game_type VARCHAR(50) NOT NULL, -- 'lol', 'valorant' 등 게임 식별자
    match_timestamp TIMESTAMPTZ NOT NULL,
    -- 게임별 고유 스탯을 JSON 형태로 저장하여 스키마 변경 없이 확장 가능
    game_specific_stats JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 자주 사용하는 조회 조건을 위한 인덱스 생성
CREATE INDEX idx_matches_user_game ON matches(user_id, game_type);
```

---

## 🚀 6. 배포 및 운영

`docker-compose`를 사용하여 두 서비스를 통합 관리하고 배포합니다.



```yaml
# docker-compose.yml
version: '3.8'

services:
  web-server:
    build: ./web_server
    ports:
      - "8080:8080"
    environment:
      - ANALYSIS_API_URL=http://analysis-server:8000/api/v1
    depends_on:
      - analysis-server
    networks:
      - app-network

  analysis-server:
    build: ./analysis_server
    environment:
      - DATABASE_URL=...
      - RIOT_API_KEY=...
      - AI_API_KEY=...
    networks:
      - app-network

  db:
    image: postgres:15-alpine
    # ... (DB 설정)
    networks:
      - app-network

networks:
  app-network:
    driver: bridge
```

### 개발 워크플로우 (Makefile)



```makefile
# 최상위 디렉토리의 Makefile
.PHONY: dev build test

# Docker Compose를 사용하여 모든 개발 서버 실행
dev:
	docker-compose up --build

# 모든 서비스 빌드
build:
	docker-compose build

# 서비스별 테스트 실행
test: test-go test-python

test-go:
	cd web_server && go test ./...

test-python:
	cd analysis_server && pytest
```