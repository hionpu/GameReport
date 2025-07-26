# DevContainer 빌드 문제 해결 보고서

## 문제 상황
DevContainer 빌드 시 `postCreateCommand`가 실패하며 다음과 같은 오류들이 발생했습니다:

1. **파일 경로 문제**: `/workspace/.devcontainer/post-create.sh: not found`
2. **Windows 줄바꿈 문제**: `\r': command not found`, `syntax error near unexpected token`
3. **실행 권한 문제**: `cannot execute: required file not found`

## 근본 원인 분석

### 1. 절대 경로 vs 상대 경로
- **문제**: `devcontainer.json`에서 `/workspace/.devcontainer/post-create.sh` 절대 경로 사용
- **원인**: 컨테이너 환경에서 workspace 폴더가 `/workspace`로 마운트되기 전에 스크립트 실행 시도
- **해결**: `.devcontainer/post-create.sh` 상대 경로로 변경

### 2. Windows 줄바꿈 문자 (CRLF)
- **문제**: Windows에서 작성된 셸 스크립트가 CRLF(`\r\n`) 줄바꿈 사용
- **원인**: Linux 컨테이너에서는 LF(`\n`)만 인식, `\r` 문자가 명령어로 해석됨
- **증상**: 
  ```bash
  set: -: invalid option
  \r': command not found
  syntax error near unexpected token `{\r'
  ```
- **해결**: `sed -i 's/\r$//'` 명령으로 CRLF → LF 변환

### 3. 복잡한 ElixirLS 설치 로직
- **문제**: 과도한 재시도 로직과 복잡한 오류 처리
- **해결**: 다른 LSP와 동일한 단순한 설치 방식으로 변경

## 해결 과정

### 1단계: 경로 문제 해결
```json
// 변경 전
"postCreateCommand": "/workspace/.devcontainer/post-create.sh"

// 변경 후  
"postCreateCommand": ".devcontainer/post-create.sh"
```

### 2단계: 줄바꿈 문제 해결
```bash
# Windows CRLF를 Unix LF로 변환
sed -i 's/\r$//' /path/to/script.sh
```

### 3단계: 스크립트 최적화
- ElixirLS 설치 로직 단순화
- MCP 설정 전 기존 `.mcp.json` 삭제 추가
- 모든 스크립트 파일에 동일한 줄바꿈 처리 적용

## 교훈

### 크로스 플랫폼 개발 시 주의사항
1. **줄바꿈 문자 통일**: Git의 `core.autocrlf` 설정 또는 `.gitattributes` 활용
2. **상대 경로 사용**: 컨테이너 환경에서는 절대 경로보다 상대 경로가 안전
3. **스크립트 실행 권한**: `chmod +x` 보다는 `bash script.sh` 형태가 더 안정적

### 개발 환경 일관성
- 다른 PC에서 잘 작동했던 이유: 해당 환경에서는 줄바꿈이 이미 LF로 설정되어 있었음
- Windows 환경에서 작업 시 에디터의 줄바꿈 설정 확인 필요

### 디버깅 전략
1. **오류 메시지 패턴 인식**: `\r': command not found`는 즉시 CRLF 문제로 판단 가능
2. **파일 형식 확인**: `file` 명령어로 줄바꿈 형식 확인
3. **단계별 검증**: 각 수정사항을 개별적으로 테스트

## 예방 방법

### Git 설정
```bash
# 전역 설정
git config --global core.autocrlf input

# 또는 .gitattributes 파일 사용
*.sh text eol=lf
```

### VSCode 설정
```json
{
  "files.eol": "\n",
  "files.insertFinalNewline": true
}
```

이 문제는 Windows와 Linux 환경 간의 차이에서 발생하는 전형적인 크로스 플랫폼 이슈였으며, 향후 유사한 문제 예방을 위해 개발 환경 설정의 표준화가 필요합니다.