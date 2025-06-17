# Day 2: 라우터 구조 설계 & 미들웨어 설정

## 🎯 Today's Goals
- [ ] API 엔드포인트 계층 구조 설계
- [ ] 미들웨어 설정 (로깅, CORS, 에러 핸들링)
- [ ] 헬스체크 엔드포인트 구현
- [ ] JSON 응답 표준화
- [ ] 구조화된 로깅 시스템

## 📋 API 엔드포인트 구조 설계

### Health & Status
```
GET /health                    - 서버 상태 확인
GET /api/status               - API 상태 및 버전 정보
```

### Player Endpoints (Week 1 목표)
```
GET /api/player/{gameName}/{tagLine}           - 플레이어 기본 정보 및 통계
GET /api/player/{gameName}/{tagLine}/matches   - 플레이어 최근 매치 목록
GET /api/player/{gameName}/{tagLine}/stats     - 플레이어 통계만 조회
```

### AI Insights (Week 2 목표)
```
GET /api/insights/{gameName}/{tagLine}         - AI 기반 인사이트
POST /api/insights/analyze                     - 커스텀 통계 분석 요청
```

### Future Endpoints (Phase 2 준비)
```
GET /api/leaderboard                           - 리더보드
GET /api/meta/agents                          - 에이전트 메타 정보
GET /api/meta/maps                            - 맵 정보
```

## 🔧 미들웨어 스택 설계

### 1. 로깅 미들웨어
- 모든 요청/응답 로깅
- 요청 ID 추적
- 성능 메트릭 (응답 시간)

### 2. CORS 미들웨어
- 프론트엔드 도메인 허용
- 개발 환경에서 모든 오리진 허용

### 3. 에러 핸들링 미들웨어
- 패닉 복구
- 구조화된 에러 응답
- 에러 로깅

### 4. 레이트 리미팅 (추후)
- Riot API 레이트 제한 준수
- 사용자별 요청 제한

## 📦 표준 JSON 응답 구조

### 성공 응답
```json
{
  "success": true,
  "data": {
    // 실제 데이터
  },
  "timestamp": "2025-06-16T10:30:00Z",
  "request_id": "uuid-string"
}
```

### 에러 응답
```json
{
  "success": false,
  "error": {
    "code": "PLAYER_NOT_FOUND",
    "message": "플레이어를 찾을 수 없습니다",
    "details": {}
  },
  "timestamp": "2025-06-16T10:30:00Z",
  "request_id": "uuid-string"
}
```

## 📝 구현 체크리스트

### 미들웨어 구현
- [ ] 로깅 미들웨어 구현
- [ ] CORS 설정
- [ ] 에러 핸들링 미들웨어
- [ ] 요청 ID 생성 및 추적

### 라우터 설정
- [ ] Chi 라우터 그룹 구성
- [ ] 미들웨어 체인 적용
- [ ] 라우트 등록

### 헬스체크
- [ ] `/health` 엔드포인트 구현
- [ ] 시스템 상태 확인 로직
- [ ] JSON 응답 구현

### 응답 구조
- [ ] 응답 구조체 정의
- [ ] 헬퍼 함수 구현
- [ ] 에러 코드 정의

## 🎯 완료 기준
- `/health` 엔드포인트가 정상 작동
- 모든 요청이 로깅됨
- CORS가 올바르게 설정됨
- 에러가 구조화된 형태로 반환됨
- 응답 시간이 로깅됨

---
작성일: 2025-06-16
상태: 진행 중
