### 1. **Content Negotiation (콘텐츠 협상)**

go

```go
func isJSONRequest(r *http.Request) bool {
    accept := r.Header.Get("Accept")
    return strings.Contains(accept, "application/json")
}
```

요청 헤더의 `Accept` 값을 확인해서 JSON 요청인지 HTML 요청인지 판별합니다.

### 2. **웹 브라우저 요청 (HDA)**

```
GET /users
Accept: text/html
→ HTML 응답 (HTMX 사용)
```

### 3. **모바일 앱 요청 (Traditional API)**

```
GET /users  
Accept: application/json
→ JSON 응답
```

### 4. **사용 예시**

**웹에서:**

html

```html
<div hx-get="/users" hx-target="#user-list">
  <!-- 서버가 HTML 조각을 반환 -->
</div>
```

**Flutter에서:**

dart

```dart
final response = await http.get(
  Uri.parse('/users'),
  headers: {'Accept': 'application/json'}
);
```

### 5. **프로젝트 구조**

```
project/
├── main.go              # 서버 코드
├── components.templ     # Templ 컴포넌트
├── components_templ.go  # 생성된 Go 코드
├── static/
│   └── style.css       # CSS 파일
└── go.mod
```

### 6. **빌드 및 실행**

bash

```bash
# Templ 컴파일
templ generate

# 서버 실행
go run .
```

이렇게 하면 **같은 API 엔드포인트**에서 웹과 모바일 모두 지원하는 하이브리드 아키텍처를 구현할 수 있습니다. 웹은 HDA의 장점을, 모바일은 전통적인 REST API의 장점을 각각 활용할 수 있습니다.

# Go Chi + Templ api 예시 코드
```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// User 모델
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// 임시 데이터
var users = []User{
	{ID: 1, Name: "김철수", Email: "kim@example.com", Age: 25},
	{ID: 2, Name: "이영희", Email: "lee@example.com", Age: 30},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 사용자 관련 라우트
	r.Route("/users", func(r chi.Router) {
		r.Get("/", getUsersHandler)
		r.Get("/{id}", getUserHandler)
		r.Post("/", createUserHandler)
		r.Put("/{id}", updateUserHandler)
		r.Delete("/{id}", deleteUserHandler)
	})

	// 정적 파일 서빙 (CSS, JS 등)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Content-Type 판별 함수
func isJSONRequest(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "application/json") || 
		   strings.Contains(r.Header.Get("Content-Type"), "application/json")
}

// 사용자 목록 조회
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if isJSONRequest(r) {
		// JSON 응답 (모바일용)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	} else {
		// HTML 응답 (웹용)
		w.Header().Set("Content-Type", "text/html")
		userListComponent := UserList(users)
		userListComponent.Render(r.Context(), w)
	}
}

// 특정 사용자 조회
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user *User
	for _, u := range users {
		if u.ID == id {
			user = &u
			break
		}
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if isJSONRequest(r) {
		// JSON 응답 (모바일용)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	} else {
		// HTML 응답 (웹용)
		w.Header().Set("Content-Type", "text/html")
		userDetailComponent := UserDetail(*user)
		userDetailComponent.Render(r.Context(), w)
	}
}

// 사용자 생성
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	
	if isJSONRequest(r) {
		// JSON 요청 처리
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	} else {
		// 폼 데이터 처리
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		
		age, _ := strconv.Atoi(r.FormValue("age"))
		user = User{
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
			Age:   age,
		}
	}

	// 새 ID 생성
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1

	// 사용자 추가
	users = append(users, user)

	if isJSONRequest(r) {
		// JSON 응답
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	} else {
		// HTML 응답 (HTMX용)
		w.Header().Set("Content-Type", "text/html")
		userItemComponent := UserItem(user)
		userItemComponent.Render(r.Context(), w)
	}
}

// 사용자 업데이트
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var userIndex = -1
	for i, u := range users {
		if u.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser User
	if isJSONRequest(r) {
		// JSON 요청 처리
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	} else {
		// 폼 데이터 처리
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		
		age, _ := strconv.Atoi(r.FormValue("age"))
		updatedUser = User{
			ID:    id,
			Name:  r.FormValue("name"),
			Email: r.FormValue("email"),
			Age:   age,
		}
	}

	updatedUser.ID = id
	users[userIndex] = updatedUser

	if isJSONRequest(r) {
		// JSON 응답
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedUser)
	} else {
		// HTML 응답 (HTMX용)
		w.Header().Set("Content-Type", "text/html")
		userItemComponent := UserItem(updatedUser)
		userItemComponent.Render(r.Context(), w)
	}
}

// 사용자 삭제
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var userIndex = -1
	for i, u := range users {
		if u.ID == id {
			userIndex = i
			break
		}
	}

	if userIndex == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 사용자 삭제
	users = append(users[:userIndex], users[userIndex+1:]...)

	if isJSONRequest(r) {
		// JSON 응답
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	} else {
		// HTML 응답 (HTMX용 - 빈 응답으로 요소 제거)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		// 빈 응답을 보내면 HTMX가 해당 요소를 제거함
	}
}
```

# Templ component

```go
package main

import "strconv"

// 사용자 목록 컴포넌트
templ UserList(users []User) {
	<div id="user-list">
		<h2>사용자 목록</h2>
		
		<!-- 사용자 추가 폼 -->
		<form hx-post="/users" hx-target="#user-list" hx-swap="outerHTML">
			<input type="text" name="name" placeholder="이름" required/>
			<input type="email" name="email" placeholder="이메일" required/>
			<input type="number" name="age" placeholder="나이" required/>
			<button type="submit">추가</button>
		</form>

		<!-- 사용자 목록 -->
		<div id="users">
			for _, user := range users {
				@UserItem(user)
			}
		</div>
	</div>
}

// 개별 사용자 아이템 컴포넌트
templ UserItem(user User) {
	<div id={ "user-" + strconv.Itoa(user.ID) } class="user-item">
		<h3>{ user.Name }</h3>
		<p>이메일: { user.Email }</p>
		<p>나이: { strconv.Itoa(user.Age) }세</p>
		
		<button 
			hx-get={ "/users/" + strconv.Itoa(user.ID) }
			hx-target={ "#user-" + strconv.Itoa(user.ID) }
			hx-swap="outerHTML">
			상세보기
		</button>
		
		<button 
			hx-delete={ "/users/" + strconv.Itoa(user.ID) }
			hx-target={ "#user-" + strconv.Itoa(user.ID) }
			hx-swap="outerHTML"
			hx-confirm="정말 삭제하시겠습니까?">
			삭제
		</button>
	</div>
}

// 사용자 상세 컴포넌트
templ UserDetail(user User) {
	<div id={ "user-" + strconv.Itoa(user.ID) } class="user-detail">
		<h3>{ user.Name } 상세 정보</h3>
		
		<!-- 편집 폼 -->
		<form hx-put={ "/users/" + strconv.Itoa(user.ID) } 
			  hx-target={ "#user-" + strconv.Itoa(user.ID) } 
			  hx-swap="outerHTML">
			<input type="text" name="name" value={ user.Name } required/>
			<input type="email" name="email" value={ user.Email } required/>
			<input type="number" name="age" value={ strconv.Itoa(user.Age) } required/>
			<button type="submit">수정</button>
		</form>

		<button 
			hx-get={ "/users/" + strconv.Itoa(user.ID) }
			hx-target={ "#user-" + strconv.Itoa(user.ID) }
			hx-swap="outerHTML">
			취소
		</button>
		
		<button 
			hx-delete={ "/users/" + strconv.Itoa(user.ID) }
			hx-target={ "#user-" + strconv.Itoa(user.ID) }
			hx-swap="outerHTML"
			hx-confirm="정말 삭제하시겠습니까?">
			삭제
		</button>
	</div>
}

// 메인 페이지 레이아웃
templ Layout(title string, content templ.Component) {
	<!DOCTYPE html>
	<html>
	<head>
		<title>{ title }</title>
		<script src="https://unpkg.com/htmx.org@1.9.10"></script>
		<link rel="stylesheet" href="/static/style.css"/>
	</head>
	<body>
		<div class="container">
			@content
		</div>
	</body>
	</html>
}
```


# Flutter 클라이언트 예시

```dart
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

// User 모델
class User {
  final int id;
  final String name;
  final String email;
  final int age;

  User({
    required this.id,
    required this.name,
    required this.email,
    required this.age,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      name: json['name'],
      email: json['email'],
      age: json['age'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'email': email,
      'age': age,
    };
  }
}

// API 서비스
class UserService {
  static const String baseUrl = 'http://localhost:8080';

  // JSON 요청을 위한 헤더
  static const Map<String, String> jsonHeaders = {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  };

  // 사용자 목록 조회
  Future<List<User>> getUsers() async {
    final response = await http.get(
      Uri.parse('$baseUrl/users'),
      headers: jsonHeaders,
    );

    if (response.statusCode == 200) {
      List<dynamic> jsonList = json.decode(response.body);
      return jsonList.map((json) => User.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load users');
    }
  }

  // 특정 사용자 조회
  Future<User> getUser(int id) async {
    final response = await http.get(
      Uri.parse('$baseUrl/users/$id'),
      headers: jsonHeaders,
    );

    if (response.statusCode == 200) {
      return User.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load user');
    }
  }

  // 사용자 생성
  Future<User> createUser(User user) async {
    final response = await http.post(
      Uri.parse('$baseUrl/users'),
      headers: jsonHeaders,
      body: json.encode(user.toJson()),
    );

    if (response.statusCode == 201) {
      return User.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to create user');
    }
  }

  // 사용자 업데이트
  Future<User> updateUser(User user) async {
    final response = await http.put(
      Uri.parse('$baseUrl/users/${user.id}'),
      headers: jsonHeaders,
      body: json.encode(user.toJson()),
    );

    if (response.statusCode == 200) {
      return User.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to update user');
    }
  }

  // 사용자 삭제
  Future<void> deleteUser(int id) async {
    final response = await http.delete(
      Uri.parse('$baseUrl/users/$id'),
      headers: jsonHeaders,
    );

    if (response.statusCode != 204) {
      throw Exception('Failed to delete user');
    }
  }
}

// 메인 앱
void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: '사용자 관리',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: UserListScreen(),
    );
  }
}

// 사용자 목록 화면
class UserListScreen extends StatefulWidget {
  @override
  _UserListScreenState createState() => _UserListScreenState();
}

class _UserListScreenState extends State<UserListScreen> {
  final UserService _userService = UserService();
  List<User> users = [];
  bool isLoading = false;

  @override
  void initState() {
    super.initState();
    loadUsers();
  }

  Future<void> loadUsers() async {
    setState(() {
      isLoading = true;
    });

    try {
      final loadedUsers = await _userService.getUsers();
      setState(() {
        users = loadedUsers;
        isLoading = false;
      });
    } catch (e) {
      setState(() {
        isLoading = false;
      });
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('사용자 목록을 불러오는데 실패했습니다: $e')),
      );
    }
  }

  Future<void> deleteUser(int id) async {
    try {
      await _userService.deleteUser(id);
      await loadUsers(); // 목록 새로고침
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('사용자가 삭제되었습니다')),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('사용자 삭제에 실패했습니다: $e')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('사용자 목록'),
        actions: [
          IconButton(
            icon: Icon(Icons.refresh),
            onPressed: loadUsers,
          ),
        ],
      ),
      body: isLoading
          ? Center(child: CircularProgressIndicator())
          : users.isEmpty
              ? Center(child: Text('사용자가 없습니다'))
              : ListView.builder(
                  itemCount: users.length,
                  itemBuilder: (context, index) {
                    final user = users[index];
                    return Card(
                      margin: EdgeInsets.all(8.0),
                      child: ListTile(
                        title: Text(user.name),
                        subtitle: Text('${user.email} • ${user.age}세'),
                        trailing: Row(
                          mainAxisSize: MainAxisSize.min,
                          children: [
                            IconButton(
                              icon: Icon(Icons.edit),
                              onPressed: () {
                                Navigator.push(
                                  context,
                                  MaterialPageRoute(
                                    builder: (context) => UserFormScreen(user: user),
                                  ),
                                ).then((_) => loadUsers());
                              },
                            ),
                            IconButton(
                              icon: Icon(Icons.delete),
                              onPressed: () {
                                showDialog(
                                  context: context,
                                  builder: (context) => AlertDialog(
                                    title: Text('삭제 확인'),
                                    content: Text('정말 삭제하시겠습니까?'),
                                    actions: [
                                      TextButton(
                                        onPressed: () => Navigator.pop(context),
                                        child: Text('취소'),
                                      ),
                                      TextButton(
                                        onPressed: () {
                                          Navigator.pop(context);
                                          deleteUser(user.id);
                                        },
                                        child: Text('삭제'),
                                      ),
                                    ],
                                  ),
                                );
                              },
                            ),
                          ],
                        ),
                      ),
                    );
                  },
                ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          Navigator.push(
            context,
            MaterialPageRoute(
              builder: (context) => UserFormScreen(),
            ),
          ).then((_) => loadUsers());
        },
        child: Icon(Icons.add),
      ),
    );
  }
}

// 사용자 폼 화면
class UserFormScreen extends StatefulWidget {
  final User? user;

  UserFormScreen({this.user});

  @override
  _UserFormScreenState createState() => _UserFormScreenState();
}

class _UserFormScreenState extends State<UserFormScreen> {
  final _formKey = GlobalKey<FormState>();
  final _nameController = TextEditingController();
  final _emailController = TextEditingController();
  final _ageController = TextEditingController();
  final UserService _userService = UserService();

  bool get isEditMode => widget.user != null;

  @override
  void initState() {
    super.initState();
    if (isEditMode) {
      _nameController.text = widget.user!.name;
      _emailController.text = widget.user!.email;
      _ageController.text = widget.user!.age.toString();
    }
  }

  Future<void> saveUser() async {
    if (_formKey.currentState!.validate()) {
      try {
        final user = User(
          id: isEditMode ? widget.user!.id : 0,
          name: _nameController.text,
          email: _emailController.text,
          age: int.parse(_ageController.text),
        );

        if (isEditMode) {
          await _userService.updateUser(user);
        } else {
          await _userService.createUser(user);
        }

        Navigator.pop(context);
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('사용자가 ${isEditMode ? '수정' : '생성'}되었습니다')),
        );
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('사용자 ${isEditMode ? '수정' : '생성'}에 실패했습니다: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(isEditMode ? '사용자 수정' : '사용자 생성'),
      ),
      body: Form(
        key: _formKey,
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            children: [
              TextFormField(
                controller: _nameController,
                decoration: InputDecoration(labelText: '이름'),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return '이름을 입력해주세요';
                  }
                  return null;
                },
              ),
              TextFormField(
                controller: _emailController,
                decoration: InputDecoration(labelText: '이메일'),
                keyboardType: TextInputType.emailAddress,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return '이메일을 입력해주세요';
                  }
                  if (!value.contains('@')) {
                    return '올바른 이메일 형식을 입력해주세요';
                  }
                  return null;
                },
              ),
              TextFormField(
                controller: _ageController,
                decoration: InputDecoration(labelText: '나이'),
                keyboardType: TextInputType.number,
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return '나이를 입력해주세요';
                  }
                  if (int.tryParse(value) == null) {
                    return '숫자를 입력해주세요';
                  }
                  return null;
                },
              ),
              SizedBox(height: 20),
              ElevatedButton(
                onPressed: saveUser,
                child: Text(isEditMode ? '수정' : '생성'),
              ),
            ],
          ),
        ),
      ),
    );
  }

  @override
  void dispose() {
    _nameController.dispose();
    _emailController.dispose();
    _ageController.dispose();
    super.dispose();
  }
}
}
```