# Figma + UXPressia 워크플로우 가이드

## 🎯 전체 워크플로우 개요

```
📋 1단계: 사용자 리서치 & 페르소나 정의 (UXPressia)
    ↓
🗺️ 2단계: 사용자 여정 매핑 (UXPressia)
    ↓
💡 3단계: 아이디어 구체화 & 기능 정의 (UXPressia + Figma)
    ↓
📱 4단계: 와이어프레임 & 프로토타입 (Figma)
    ↓
🎨 5단계: 시각적 디자인 & 인터랙션 (Figma)
    ↓
🔄 6단계: 피드백 & 개선 (UXPressia + Figma)
```

## 📋 1단계: 사용자 리서치 & 페르소나 정의 (UXPressia)

### UXPressia에서 페르소나 생성
```
🎯 목표: 타겟 사용자 명확히 정의

📝 페르소나 구성 요소:
- 기본 정보 (이름, 나이, 직업, 지역)
- 기술적 특성 (사용 기기, 기술 수준)
- 목표 & 동기 (무엇을 원하는가?)
- 불만 & 문제점 (현재 어려움)
- 선호도 (인터페이스, 사용 패턴)
```

### 예시: Go + Chi + Flutter 프로젝트
```
페르소나: "효율적인 팀장 김철수"

👤 기본 정보:
- 이름: 김철수 (32세)
- 직업: IT 회사 팀장
- 거주지: 서울
- 팀 규모: 8명

💻 기술적 특성:
- 주 사용 기기: 데스크톱 + 아이폰
- 기술 수준: 중급 (새로운 도구 학습 시간 부족)
- 선호 환경: 데스크톱(상세 작업), 모바일(빠른 확인)

🎯 목표 & 동기:
- 팀원 관리 간소화
- 업무 효율성 향상
- 실시간 팀 상태 파악

😤 불만 & 문제점:
- 복잡한 관리 시스템
- 여러 도구 사용의 번거로움
- 모바일에서 제한적인 기능
```

## 🗺️ 2단계: 사용자 여정 매핑 (UXPressia)

### 여정 맵 생성 과정
```
🎯 목표: 사용자의 전체 경험 플로우 이해

📍 여정 단계 정의:
1. 인지 (문제 인식)
2. 탐색 (해결책 찾기)
3. 사용 (실제 사용)
4. 목표 달성 (성공 경험)
5. 지속 사용 (재사용)
```

### 예시: 사용자 관리 여정
```
김철수의 "새 팀원 온보딩" 여정:

1️⃣ 인지 단계:
- 상황: 새 팀원 입사 예정
- 행동: 기존 관리 도구 확인
- 감정: 😐 "또 복잡한 과정이 시작되네"
- 생각: "더 간단한 방법이 없을까?"
- 접촉점: 이메일, 기존 시스템

2️⃣ 탐색 단계:
- 행동: 새로운 관리 도구 검색
- 감정: 🤔 "이 도구가 우리에게 맞을까?"
- 생각: "웹과 모바일 둘 다 써야 하는데..."
- 접촉점: 구글 검색, 우리 랜딩 페이지
- 💡 기회: 명확한 기능 설명 필요

3️⃣ 사용 단계:
- 행동: 계정 생성 및 팀원 등록
- 감정: 😊 "생각보다 쉽네!"
- 생각: "모바일에서도 확인할 수 있겠다"
- 접촉점: 회원가입 폼, 사용자 생성 화면
- 💡 기회: 직관적인 UI 디자인 중요

4️⃣ 목표 달성:
- 행동: 팀원 정보 입력 완료
- 감정: 😌 "팀원이 바로 사용할 수 있겠다"
- 생각: "다음에도 이 도구를 쓸 것 같다"
- 접촉점: 성공 메시지, 대시보드
- 💡 기회: 성공 경험 강화

5️⃣ 지속 사용:
- 행동: 정기적인 팀 관리
- 감정: 😎 "관리가 훨씬 편해졌다"
- 생각: "다른 기능도 써볼까?"
- 접촉점: 모바일 알림, 웹 대시보드
- 💡 기회: 추가 기능 추천
```

## 💡 3단계: 아이디어 구체화 & 기능 정의

### UXPressia에서 기능 우선순위 정의
```
여정 맵에서 파악한 기회점들:
1. 👑 핵심 기능 (Must Have)
2. 🎯 중요 기능 (Should Have)  
3. 💡 개선 기능 (Could Have)
4. 🔮 미래 기능 (Won't Have This Time)
```

### Figma로 넘어가기 전 체크리스트
```
✅ 페르소나 명확히 정의됨
✅ 핵심 여정 맵 완성
✅ 주요 기능 목록 정리
✅ 기능별 우선순위 설정
✅ 웹/모바일 요구사항 구분
```

## 📱 4단계: 와이어프레임 & 프로토타입 (Figma)

### Figma 프로젝트 구조 설정
```
📁 프로젝트: 사용자 관리 시스템
├── 📄 01_Style_Guide (색상, 폰트, 컴포넌트)
├── 📄 02_Web_Wireframes (웹 와이어프레임)
├── 📄 03_Mobile_Wireframes (모바일 와이어프레임)
├── 📄 04_Web_UI (웹 시각 디자인)
├── 📄 05_Mobile_UI (모바일 시각 디자인)
└── 📄 06_Prototypes (인터랙션 프로토타입)
```

### 와이어프레임 작성 순서
```
1. 핵심 화면부터 시작
   - 로그인 화면
   - 사용자 목록 화면
   - 사용자 생성 화면

2. 사용자 여정 따라 화면 연결
   - UXPressia 여정 맵 참조
   - 각 단계별 필요 화면 설계

3. 웹/모바일 버전 구분
   - 웹: 상세 정보, 복잡한 작업
   - 모바일: 빠른 확인, 간단한 작업
```

### 와이어프레임 예시
```
📱 사용자 생성 화면 (모바일):
┌─────────────────────┐
│ ← 새 사용자 추가      │
├─────────────────────┤
│ 이름 [입력필드]       │
│ 이메일 [입력필드]     │
│ 역할 [드롭다운]       │
│                     │
│ [저장하기 버튼]       │
└─────────────────────┘

💻 사용자 관리 화면 (웹):
┌─────────────────────────────────────┐
│ 사용자 관리               [+ 추가]    │
├─────────────────────────────────────┤
│ 검색: [___________] [필터▼]          │
├─────────────────────────────────────┤
│ 이름    이메일        역할    액션    │
│ 김철수  kim@...      팀장    [수정]   │
│ 이영희  lee@...      팀원    [수정]   │
└─────────────────────────────────────┘
```

## 🎨 5단계: 시각적 디자인 & 인터랙션 (Figma)

### 디자인 시스템 구축
```
🎨 색상 팔레트:
- Primary: #2563EB (파란색)
- Secondary: #64748B (회색)
- Success: #16A34A (초록색)
- Error: #DC2626 (빨간색)
- Background: #F8FAFC (연한 회색)

📝 타이포그래피:
- 제목: Inter Bold 24px
- 본문: Inter Regular 16px
- 캡션: Inter Regular 14px

🔘 컴포넌트:
- 버튼 (Primary, Secondary, Ghost)
- 입력 필드 (Default, Focus, Error)
- 카드 (기본, 호버, 선택)
```

### 인터랙션 정의
```
🖱️ 웹 인터랙션 (HTMX 고려):
- 버튼 클릭 → 로딩 상태 → 결과 표시
- 폼 제출 → 검증 → 성공/실패 메시지
- 목록 업데이트 → 새 항목 추가 애니메이션

📱 모바일 인터랙션 (Flutter 고려):
- 스와이프 제스처 → 삭제 액션
- 풀투리프레시 → 목록 새로고침
- 플로팅 액션 버튼 → 새 항목 추가
```

## 🔄 6단계: 피드백 & 개선

### UXPressia에서 피드백 정리
```
📊 사용자 테스트 결과:
- 어떤 단계에서 어려움을 겪었는지
- 예상과 다른 행동 패턴
- 개선이 필요한 부분

💡 개선 아이디어:
- 여정 맵 업데이트
- 새로운 기회점 발견
- 추가 기능 아이디어
```

### Figma에서 디자인 개선
```
🎨 디자인 반복:
- 사용자 피드백 반영
- A/B 테스트 버전 생성
- 프로토타입 업데이트
```

## 🔗 도구 간 연결 방법

### 1. 링크 연결
```
UXPressia 여정 맵 → Figma 프로토타입
- 각 여정 단계에 Figma 링크 첨부
- 기능 ID로 연결점 명시
```

### 2. 에셋 공유
```
Figma → UXPressia
- 스크린샷을 UXPressia 여정 맵에 첨부
- 실제 UI 모습을 여정에 시각화
```

### 3. 문서화
```
공통 문서:
- 기능 ID 매핑 테이블
- 디자인 결정 사항 로그
- 사용자 피드백 정리
```

## 🎯 실제 적용 예시

### 프로젝트 시작 (1주차)
```
월: UXPressia 페르소나 정의
화: UXPressia 여정 맵 생성
수: 기능 우선순위 정의
목: Figma 프로젝트 설정
금: 핵심 와이어프레임 작성
```

### 디자인 개발 (2-3주차)
```
Figma 중심 작업:
- 와이어프레임 → 시각 디자인
- 프로토타입 제작
- 인터랙션 정의
- 컴포넌트 시스템 구축
```

### 피드백 & 개선 (4주차)
```
UXPressia + Figma 연동:
- 프로토타입 테스트 결과 정리
- 여정 맵 업데이트
- 디자인 개선 반영
- 다음 스프린트 계획
```

## 💡 실전 팁

### 효율적인 작업 방법
```
1. 🎯 작은 단위로 시작
   - 1개 페르소나, 1개 핵심 여정부터
   - 점진적 확장

2. 🔄 빠른 반복
   - 저해상도 프로토타입으로 빠른 검증
   - 큰 틀부터 세부사항 순서

3. 📱 플랫폼 특성 고려
   - 웹: 정보 밀도 높게
   - 모바일: 단순하고 직관적으로

4. 🔗 연결점 명확히
   - 기능 ID로 일관성 유지
   - 도구 간 링크 적극 활용
```

이렇게 UXPressia와 Figma를 연계해서 사용하면 사용자 중심의 체계적인 UI/UX 설계가 가능합니다!