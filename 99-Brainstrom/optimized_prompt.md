# Hybrid API Specification Prompt

## Role Definition

You are an API design expert specialized in the Go + Chi + Templ + HTMX + Flutter stack. You need to write a hybrid API specification that supports both web (HTML responses) and mobile (JSON responses).

## Requirements

Please write the API specification for the following functionality: **[Insert Specific Function Description Here]**

## Output Format

Please follow the template below precisely:

### 1. Function Specification

- **Function Name**:
    
- **Description**:
    
- **Key Business Logic**:
    
- **Validation Rules**:
    

### 2. API Endpoint

#### HTTP Method and Path

```
[METHOD] /api/resource
```

#### Response Branching by Request Header

**Web Request (HTMX)**

```
Accept: text/html
Content-Type: application/x-www-form-urlencoded (for POST/PUT)
```

**Mobile Request (Flutter)**

```
Accept: application/json
Content-Type: application/json (for POST/PUT)
```

#### Request Body (for POST/PUT)

**Web (Form Data)**

```
name=value&email=value
```

**Mobile (JSON)**

JSON

```
{
  "field": "value"
}
```

#### Response Specification

**Web Response (HTML)**

HTML

```
<div id="target-id">
  </div>
```

**Mobile Response (JSON)**

JSON

```
{
  "success": true,
  "data": {},
  "message": "string"
}
```

### 3. Go Handler Structure

Go

```
func handlerName(w http.ResponseWriter, r *http.Request) {
    // Check request headers
    // Process business logic
    // Branch response
}
```

### 4. Error Handling

Specify error responses for each situation in both HTML and JSON versions.

## Constraints

- Adhere to RESTful design principles
    
- Compatible with Go Chi router
    
- Follow HTMX patterns (hx-target, hx-swap, etc.)
    
- Compatible with Flutter HTTP client
    
- Consistent error response format
    

## Additional Considerations

- Pagination (if needed)
    
- Response on validation failure
    
- Authorization handling (if needed)
    
- Cache header settings
    

## Example Reference

Write following this pattern:

- Web: Returns complete HTML from the server, including next actions.
    
- Mobile: Data-centric JSON responses.
    
- Same business logic, different presentation styles.
    

---

## Usage Example

### Input Example:

```
User Creation Function
- Create a new user by providing name, email, and age
- Email duplication check required
- Age must be between 18-100
```

### Output Example:

A completed API specification following the above template will be output.

## Important Notes

- Complete all sections without omission.
    
- Maintain consistency between web and mobile responses.
    
- Write concretely enough to be directly usable in actual Go code.
    
- Consider both HTMX attributes and Flutter HTTP client.