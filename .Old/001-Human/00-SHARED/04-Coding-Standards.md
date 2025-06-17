---
tags:
  - agent/all
  - system/config
  - phase/planning
  - doc/guide
  - tech/cpp
  - tech/unity
  - priority/high
  - status/complete
references:
  - "[[01-Project-Overview]]"
---

# Coding Standards and Guidelines

## General Principles
### 1. Naming Conventions
- **File Names**: PascalCase (e.g., `PlayerController.cs`, `DatabaseManager.cpp`)
- **Class Names**: PascalCase (e.g., `GameManager`, `NetworkClient`)
- **Method Names**: PascalCase (C#), camelCase (C++)
- **Variable Names**: camelCase (e.g., `playerHealth`, `maxConnections`)
- **Constants**: UPPER_SNAKE_CASE (e.g., `MAX_PLAYERS`, `DEFAULT_TIMEOUT`)

### 2. Documentation and Comments
- All public methods must have summary comments
- Complex logic requires inline comments
- Use TODO, FIXME, HACK tags actively
- Update changelog for API changes

### 3. Error Handling
- All exceptions must be logged
- Provide meaningful error messages to users
- Implement retry logic for transient errors (network, etc.)

## C# (Unity Client) Standards

### Namespace Structure
```csharp
namespace RaidMaster
{
    namespace UI { /* UI-related classes */ }
    namespace Network { /* Network communication */ }
    namespace Data { /* Data models */ }
    namespace Utils { /* Utility classes */ }
}
```

### Class Structure Order
```csharp
public class ExampleClass : MonoBehaviour
{
    // 1. Constants
    private const int MAX_ITEMS = 100;
    
    // 2. SerializeField (Inspector exposure)
    [SerializeField] private Button actionButton;
    
    // 3. Public fields/properties
    public int Health { get; private set; }
    
    // 4. Private fields
    private NetworkClient networkClient;
    
    // 5. Unity lifecycle methods
    private void Awake() { }
    private void Start() { }
    private void Update() { }
    
    // 6. Public methods
    public void TakeDamage(int damage) { }
    
    // 7. Private methods
    private void Die() { }
}
```

### Async Patterns
```csharp
// Use async/await pattern
public async Task<ApiResponse> CallServerAsync()
{
    try
    {
        var response = await networkClient.SendRequestAsync(request);
        return response;
    }
    catch (Exception ex)
    {
        Debug.LogError($"API call failed: {ex.Message}");
        throw;
    }
}
```

## C++ (Server) Standards

### Namespace Structure
```cpp
namespace raid_master {
    namespace network { /* Network-related */ }
    namespace database { /* Database-related */ }
    namespace logic { /* Game logic */ }
    namespace utils { /* Utility functions */ }
}
```

### Header File Structure
```cpp
#pragma once

#include <vector>
#include <memory>
#include "base_class.h"

namespace raid_master {

class ExampleClass : public BaseClass {
public:
    // Constructor/Destructor
    ExampleClass();
    virtual ~ExampleClass();
    
    // Public methods
    void Initialize();
    bool ProcessRequest(const Request& request);
    
private:
    // Private methods
    void HandleError(const std::string& error);
    
    // Member variables
    std::vector<int> data_;
    std::unique_ptr<NetworkManager> network_manager_;
};

} // namespace raid_master
```

### Memory Management
```cpp
// Prefer smart pointers
std::unique_ptr<Player> player = std::make_unique<Player>();
std::shared_ptr<GameSession> session = std::make_shared<GameSession>();

// Use RAII pattern
class DatabaseConnection {
public:
    DatabaseConnection() { Connect(); }
    ~DatabaseConnection() { Disconnect(); }
private:
    void Connect();
    void Disconnect();
};
```

## Git Commit Standards

### Commit Message Format
```
<type>(<scope>): <subject>

<body>

<footer>
```

### Type Categories
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code formatting
- `refactor`: Code refactoring
- `test`: Adding/modifying tests
- `chore`: Build tasks, package manager config

### Example
```
feat(inventory): add item drag and drop functionality

- Enable drag and drop between inventory slots
- Integrate with server API for real-time sync
- Include item quantity limit validation

Closes #123
```

## Project Structure Standards

### Unity Project
```
Assets/
├── Scripts/
│   ├── UI/
│   ├── Network/
│   ├── Data/
│   └── Utils/
├── Prefabs/
├── Scenes/
├── Materials/
└── Resources/
```

### C++ Server Project
```
server/
├── src/
│   ├── network/
│   ├── database/
│   ├── logic/
│   └── utils/
├── include/
├── tests/
├── config/
└── CMakeLists.txt
```

## Performance Guidelines

### Unity
- Use Object Pooling pattern
- Minimize unnecessary Update() calls
- Update UI only when changes occur
- Minimize garbage collection

### C++ Server
- Minimize memory allocations
- Use asynchronous I/O actively
- Implement database connection pooling
- Apply caching strategies

## Security Guidelines
- Always validate client input on server
- Prevent SQL injection (use Prepared Statements)
- Store passwords as hashes
- Apply rate limiting to API requests

## 📋 Document References

### Related Documents:
- `00-SHARED/01-Project-Overview.md` - Technical stack and development principles
- `00-SHARED/Document-Standards.md` - Documentation formatting and tagging standards
- `02-SERVER-AGENT/` - Server implementations must follow C++ standards
- `03-CLIENT-AGENT/` - Client implementations must follow Unity/C# standards

### Dependencies:
- **Requires**: `01-Project-Overview.md` for technical stack definitions
- **Blocks**: All code implementation across SERVER and CLIENT agents

### Cross-Agent Impact:
- **DESIGN**: Must consider coding constraints when creating specifications
- **SERVER**: Must implement all C++ code following these standards
- **CLIENT**: Must implement all Unity/C# code following these standards
- **LEAD**: Uses these standards for code quality review and integration