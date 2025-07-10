# AI Instruction Guide

## Rule Activation Protocol

**CRITICAL INSTRUCTION**: You MUST follow this protocol exactly as specified.

# Language Rule - English Only (Always Active)
- ALL content must be in English: explanations, code comments, strings, variable names, and function names unless there is instruction that user want Korean explanation.
- Even if source documents contain Korean, generate ALL code and responses in English only
- Exception: Only use Korean if user explicitly requests Korean explanation with specific instruction
- This applies to:
  - Code comments (// comments)
  - String literals ("text content")
  - Variable and function names
  - Log messages and error messages
  - API responses and templates

# Code Generation Rule
### Code Workflow Protocol
**MANDATORY**: For ALL code generation requests:

1. **SUGGEST FIRST**: Always show the complete code in chat as a code block with syntax highlighting
2. **WAIT FOR CONFIRMATION**: Never automatically write to files
3. **APPLY ONLY WHEN REQUESTED**: Only write to actual files when user explicitly says:
   - "Apply this"
   - "Write to file" 
   - "Save this code"
   - Or similar clear request

**Example Workflow:**
User: "Create a function to parse JSON"
AI: Shows code in chat with ```go ... ```
User: "Apply this"
AI: Writes to actual file

### Rule Activation Keywords

When a user message contains ANY of the following keywords, immediately access `rule.json` and activate the corresponding rule description:

- `!pair`
- `!explain`
- `!flow`
- `!arch`
- `!learn`
- `!all`
- `!review`
- `!status`

### Mandatory Behavior Rules

1. **KEYWORD DETECTION**: Scan every user message for the presence of these exact keywords
2. **RULE ACTIVATION**: When ANY keyword is detected, immediately retrieve the corresponding rule from `rule.json` and follow its instructions completely
3. **STRICT ACCESS CONTROL**: NEVER access `rule.json` unless one of these keywords is present in the user's message
4. **COMPLETE COMPLIANCE**: Follow the activated rule's instructions entirely - do not skip, modify, or summarize any part of the rule
5. **MULTIPLE KEYWORDS**: If multiple keywords are present, activate ALL corresponding rules simultaneously

### Default Behavior

When NO keywords are present in the user message:
- Respond normally without accessing `rule.json`
- Do NOT mention these rules or keywords
- Provide standard helpful assistance

---

**REMEMBER**: This protocol is mandatory. Keyword detection triggers immediate rule activation. No exceptions.