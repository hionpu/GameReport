
```go
// 모바일 친화적 HDA API 설계 예시
package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// 1. 단순화된 액션 정의
type ActionLink struct {
	Href   string            `json:"href"`
	Method string            `json:"method"`
	Title  string            `json:"title,omitempty"`
	Fields map[string]string `json:"fields,omitempty"` // 단순화된 필드 정의
}

// 2. 플랫폼별 응답 구조
type MobileResponse struct {
	Data    interface{}            `json:"data"`
	Actions map[string]ActionLink  `json:"actions"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

type WebResponse struct {
	Data    interface{}            `json:"data"`
	Links   map[string]interface{} `json:"_links"`
	Embeds  map[string]string      `json:"_embedded,omitempty"`
}

// 3. 계정 정보 API 예시
func getAccount(w http.ResponseWriter, r *http.Request) {
	account := map[string]interface{}{
		"id":      123,
		"name":    "김철수",
		"balance": 50000,
		"status":  "active",
	}

	if isMobileClient(r) {
		// 모바일용 단순화된 응답
		response := MobileResponse{
			Data: account,
			Actions: map[string]ActionLink{
				"deposit": {
					Href:   "/accounts/123/deposit",
					Method: "POST",
					Title:  "입금",
					Fields: map[string]string{
						"amount": "number",
						"memo":   "text",
					},
				},
				"withdraw": {
					Href:   "/accounts/123/withdraw",
					Method: "POST",
					Title:  "출금",
					Fields: map[string]string{
						"amount": "number",
						"memo":   "text",
					},
				},
				"transfer": {
					Href:   "/accounts/123/transfer",
					Method: "POST",
					Title:  "이체",
					Fields: map[string]string{
						"to_account": "text",
						"amount":     "number",
						"memo":       "text",
					},
				},
			},
			Meta: map[string]interface{}{
				"can_withdraw": account["balance"].(int) > 0,
				"daily_limit":  100000,
			},
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// 웹용 전통적 HDA 응답
		response := WebResponse{
			Data: account,
			Links: map[string]interface{}{
				"self": map[string]string{
					"href": "/accounts/123",
				},
				"deposit": map[string]interface{}{
					"href":   "/accounts/123/deposit",
					"method": "POST",
					"fields": []map[string]interface{}{
						{"name": "amount", "type": "number", "required": true},
						{"name": "memo", "type": "text", "required": false},
					},
				},
				"withdraw": map[string]interface{}{
					"href":      "/accounts/123/withdraw",
					"method":    "POST",
					"condition": "balance > 0",
					"fields": []map[string]interface{}{
						{"name": "amount", "type": "number", "required": true, "max": account["balance"]},
						{"name": "memo", "type": "text", "required": false},
					},
				},
			},
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// 4. 주문 워크플로우 API 예시
func checkoutFlow(w http.ResponseWriter, r *http.Request) {
	step := r.URL.Query().Get("step")
	if step == "" {
		step = "shipping"
	}

	if isMobileClient(r) {
		// 모바일용: 단계별 독립적인 엔드포인트
		switch step {
		case "shipping":
			response := MobileResponse{
				Data: map[string]interface{}{
					"step":       1,
					"total_step": 3,
					"title":      "배송 정보",
				},
				Actions: map[string]ActionLink{
					"submit": {
						Href:   "/checkout/shipping",
						Method: "POST",
						Title:  "다음 단계",
						Fields: map[string]string{
							"address": "text",
							"phone":   "text",
						},
					},
					"back": {
						Href:   "/cart",
						Method: "GET",
						Title:  "장바구니로",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
			
		case "payment":
			response := MobileResponse{
				Data: map[string]interface{}{
					"step":       2,
					"total_step": 3,
					"title":      "결제 정보",
				},
				Actions: map[string]ActionLink{
					"submit": {
						Href:   "/checkout/payment",
						Method: "POST",
						Title:  "결제하기",
						Fields: map[string]string{
							"payment_method": "select",
							"card_number":    "text",
						},
					},
					"back": {
						Href:   "/checkout?step=shipping",
						Method: "GET",
						Title:  "이전 단계",
					},
				},
			}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		// 웹용: 복잡한 워크플로우 지원
		response := WebResponse{
			Data: map[string]interface{}{
				"current_step": step,
				"progress":     33,
			},
			Links: map[string]interface{}{
				"next": map[string]interface{}{
					"href":      "/checkout?step=payment",
					"method":    "GET",
					"condition": "shipping_info.valid",
				},
				"previous": map[string]interface{}{
					"href":   "/checkout?step=cart",
					"method": "GET",
				},
				"skip": map[string]interface{}{
					"href":      "/checkout?step=payment",
					"method":    "GET",
					"condition": "user.has_default_address",
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}

// 5. 파일 다운로드 API 예시
func downloadInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceID := r.URL.Query().Get("id")
	
	if isMobileClient(r) {
		// 모바일용: 직접 다운로드 URL 제공
		response := MobileResponse{
			Data: map[string]interface{}{
				"invoice_id": invoiceID,
				"filename":   "invoice_123.pdf",
				"size":       "245KB",
			},
			Actions: map[string]ActionLink{
				"download": {
					Href:   "/files/invoices/123.pdf",
					Method: "GET",
					Title:  "다운로드",
				},
				"share": {
					Href:   "/invoices/123/share",
					Method: "POST",
					Title:  "공유",
					Fields: map[string]string{
						"email": "text",
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	} else {
		// 웹용: 브라우저 내장 기능 활용
		response := WebResponse{
			Data: map[string]interface{}{
				"invoice_id": invoiceID,
			},
			Links: map[string]interface{}{
				"download": map[string]interface{}{
					"href":    "/invoices/123/download",
					"method":  "GET",
					"type":    "application/pdf",
					"target":  "_blank",
				},
				"preview": map[string]interface{}{
					"href":   "/invoices/123/preview",
					"method": "GET",
					"type":   "text/html",
					"embed":  true,
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}

// 클라이언트 타입 판별
func isMobileClient(r *http.Request) bool {
	userAgent := r.Header.Get("User-Agent")
	accept := r.Header.Get("Accept")
	
	// 모바일 앱인지 판별하는 로직
	return strings.Contains(userAgent, "Mobile") || 
		   strings.Contains(accept, "application/vnd.api+json") ||
		   r.Header.Get("X-Client-Type") == "mobile"
}

// 6. 에러 응답도 플랫폼별로 다르게
func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if isMobileClient(r) {
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "INSUFFICIENT_BALANCE",
				"message": "잔액이 부족합니다",
				"details": map[string]interface{}{
					"current_balance": 5000,
					"required_amount": 10000,
				},
			},
			"actions": map[string]ActionLink{
				"deposit": {
					Href:   "/accounts/123/deposit",
					Method: "POST",
					Title:  "입금하기",
				},
				"retry": {
					Href:   "/accounts/123/withdraw",
					Method: "POST",
					Title:  "다시 시도",
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	} else {
		// 웹용 에러 응답
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"message": "잔액이 부족합니다",
				"code":    "INSUFFICIENT_BALANCE",
			},
			"_links": map[string]interface{}{
				"deposit": map[string]interface{}{
					"href":   "/accounts/123/deposit",
					"method": "POST",
					"title":  "입금하기",
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}
}
```

# **모바일에서 문제가 되는 HDA 패턴:**

1. **복잡한 폼 정의** - 동적 UI 생성 어려움
2. **HTML 콘텐츠 포함** - 네이티브 UI 변환 복잡
3. **조건부 로직** - 클라이언트 평가 부담
4. **복합 미디어 타입** - 브라우저 의존 기능들
5. **웹 특화 워크플로우** - 모바일 UX 패턴과 상충

# **해결 방안:**

1. **단순화된 액션 정의** - 필드 타입만 명시
2. **플랫폼별 응답 구조** - 헤더로 클라이언트 타입 판별
3. **메타데이터 분리** - 조건부 로직을 서버에서 처리
4. **직접 리소스 링크** - 모바일 친화적 URL 제공
5. **단계별 독립 엔드포인트** - 모바일 네비게이션 최적화

이렇게 하면 **같은 백엔드**에서 웹과 모바일 모두 지원하면서, 각 플랫폼의 특성에 맞는 HDA를 구현할 수 있습니다.