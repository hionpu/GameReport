
# Genesis Rulebook: Algorithmic Markdown Compression

---
version: 1.0.0
author: Project-AI-Integration
description: A deterministic, algorithmic ruleset for converting verbose Markdown into a compact, machine-readable format for AI consumption.
---

## 1. Overview

This document defines the rules for a lossless, algorithmic compression of Markdown (`.md`) files. The goal is to create a consistent, token-efficient format that can be reliably parsed by an AI model.

## 2. Core Principle: Parse, Don't Interpret

The conversion process is based on parsing the Markdown structure and applying fixed transformations. It does not rely on subjective AI interpretation. The process MUST be 100% deterministic.

## 3. Compression Rules

### 3.1. Structure & Delimiters

-   **Source:** Markdown Headers (`#`, `##`, `###`, etc.).
-   **Rule:** Every Header line initiates a new Key-Value block. The Header text becomes the `Key`. The content following it, until the next header of the same or higher level, becomes the `Value`.
-   **Key Transformation:**
    1.  Remove `#` symbols and leading/trailing whitespace.
    2.  Convert the text to `PascalCase`.
    3.  Append a colon (`:`).
-   **Delimiters:**
    -   `|` (Pipe): Separates major blocks derived from `H1` or `H2` headers.
    -   `;` (Semicolon): Separates sub-blocks or complex key-value pairs within a larger value.
    -   `,` (Comma): Separates list items or multiple simple values.

**Example:**
```

## My Section

Some text.

### My Subsection

More text.

```
**Becomes:**
```

MySection:Some text.;MySubsection:More text

```

### 3.2. Markdown Element Conversion

#### **Lists (`-`, `*`, `1.`)**

-   **Rule:** Collapse all list items into a single, comma-separated string.

**Example:**
```

- First item
- Second item

```
**Becomes:**
```

First item,Second item

```

#### **Emphasis (Bold/Italics)**

-   **Rule:** All emphasis formatting (`*`, `_`, `**`) MUST be stripped by default. The only exception is for predefined keywords which are converted into special tags.
-   **Keyword List:** `["MUST", "NOT", "WARNING", "IMPORTANT", "CRITICAL", "TBD", "DONE", "WIP"]`
-   **Logic:** If emphasized text is a case-insensitive match for a keyword, convert it to `!KEYWORD`. Otherwise, remove the markdown symbols.

**Example:**
```

This is _important_. You **MUST** proceed with caution. The status is **WIP**.

```
**Becomes:**
```

This is important. You !MUST proceed with caution. The status is !WIP

````

#### **Tables**

-   **Rule:** Table conversion is based on a complexity metric.
-   **Simple Table (Flatten):** If a table has **4 or fewer columns** AND no single cell contains more than **60 characters**, it WILL be flattened.
    -   **Flattening Logic:** `Tbl:Row(ColHdr1:Val1,ColHdr2:Val2);Row(ColHdr1:Val3,ColHdr2:Val4)`
-   **Complex Table (Reference):** If it exceeds either metric, it WILL be referenced.
    -   **Reference Logic:** `TblRef:OriginalDocID/TableName_or_SectionID`

#### **Links & Code**

-   **Links:** The format `[Text](URL)` WILL be converted to `Text[URL]`.
-   **Code:** Inline `code` is preserved. For code blocks:
    -   If a block has **fewer than 3 lines**, it WILL be inlined: `Code(language:"lang",src:"source_code")`.
    -   If a block has **3 or more lines**, it WILL be referenced: `CodeRef:OriginalDocID/SectionID`.

### 3.3. Abbreviations & Symbology

#### **Abbreviations**

-   **Rule:** The parser MUST use a project-wide `glossary.json` file to perform a case-sensitive, whole-word search and replace. This occurs after structural conversion.

#### **Symbols**

-   **Rule:** Symbols are used ONLY when specific trigger phrases are present in the source sentence.
-   `→` (Dependency): Used for trigger phrases like "depends on", "is blocked by", "requires".
-   `?=` (Uncertainty/Query): Used for trigger phrases like "TBD", "unconfirmed", "potential issue", "to be investigated".

### 3.4. Final Processing

#### **YAML Frontmatter**

-   **Rule:** YAML frontmatter MUST be preserved. Arrays SHOULD be converted to an inline format.

**Example:**
```yaml
tags:
  - tag1
  - tag2
````

**Becomes:**

YAML

```
tags:[tag1,tag2]
```

#### **Whitespace**

- **Rule:** After all other transformations, all non-semantic whitespace (including newlines, indents, and extra spaces) between blocks WILL be removed.