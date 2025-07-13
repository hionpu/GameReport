#!/usr/bin/env python3
"""
Enhanced ctags tree generator with ultra-compact single-line format
Reads a tags file and creates a structured tree view with maximum token efficiency
"""

import os
import sys
import re
from typing import Dict, List, Optional, Tuple

class TreeNode:
    def __init__(self, name: str, node_type: str = "folder"):
        self.name = name
        self.type = node_type  # folder, file, function, method, class, struct, variable, etc.
        self.children: Dict[str, TreeNode] = {}
        self.parent: Optional[TreeNode] = None
        self.signature = ""  # Store function/method signature
        self.return_type = ""  # Store return type information
        
    def add_child(self, child: 'TreeNode'):
        self.children[child.name] = child
        child.parent = self
        
    def get_or_create_child(self, name: str, node_type: str = "folder") -> 'TreeNode':
        if name not in self.children:
            self.children[name] = TreeNode(name, node_type)
            self.children[name].parent = self
        return self.children[name]

def get_symbol_map():
    """Get the symbol mapping for node types"""
    return {
        "folder": "D",      # Directory
        "file": "F",        # File
        "function": "f",    # function
        "method": "m",      # method
        "field": "v",       # variable/field
        "class": "c",       # class
        "struct": "s",      # struct
        "interface": "i",   # interface
        "type": "t",        # type
        "variable": "v",    # variable
        "constant": "C",    # Constant
        "css_class": ".",   # CSS class (.)
        "css_id": "#",      # CSS ID (#)
        "html_id": "#",     # HTML ID (#)
        "heading": "h",     # heading
        "header": "H",      # Header
        "namespace": "n",   # namespace
        "enum": "e",        # enum
        "enum_member": "E", # Enum member
        "alias": "a",       # alias
        "local_var": "l",   # local variable
        "parameter": "p",   # parameter
        "file_marker": "F", # file marker
        "macro": "M",       # macro
        "property": "P",    # property
        "object": "o",      # object
        "external": "x",    # external
        "reference": "r",   # reference
        "typedef": "T",     # typedef
        "block": "b",       # block
        "keyword": "k",     # keyword
        "widget": "w",      # widget
        "union": "u",       # union
        "array": "A",       # array
        "css_selector": "S", # CSS selector
        "css_rule": "R",    # CSS rule
        "label": "L",       # label
        "define": "d",      # define
        "event": "!",       # event
        "getter": "g",      # getter
        "setter": "S",      # setter
        "unknown": "?",     # unknown
        "qualified": "q",   # qualified
        "name": "N",        # name
        "binding": "B",     # binding
        "json_key": "j",    # JSON key
        "key": "K",         # key
        "value": "V",       # value
        "wrapper": "W",     # wrapper
        "zone": "z"         # zone
    }

def format_signature_display(node: TreeNode) -> str:
    """Format the signature and return type for display"""
    display_parts = []
    
    if node.signature:
        display_parts.append(node.signature)
    
    if node.return_type:
        display_parts.append(f"->{node.return_type}")
    
    return "".join(display_parts)

def format_node_compact(node: TreeNode) -> str:
    """Format a single node and its children in ultra-compact format"""
    symbol_map = get_symbol_map()
    symbol = symbol_map.get(node.type, "?")
    
    # Format name with signature if available
    display_name = node.name
    signature_display = format_signature_display(node)
    if signature_display:
        display_name += signature_display
    
    # Start with the node representation
    result = f"{symbol}:{display_name}"
    
    # Add children if they exist
    if node.children:
        child_parts = []
        for child_name, child_node in sorted(node.children.items()):
            child_parts.append(format_node_compact(child_node))
        
        if child_parts:
            result += "{" + " ".join(child_parts) + "}"
    
    return result

def print_tree_compact(node: TreeNode, file_obj=None):
    """Print tree structure in ultra-compact single-line format"""
    if node.name == "root":
        # Print all top-level nodes in compact format
        compact_parts = []
        for child_name, child_node in sorted(node.children.items()):
            compact_parts.append(format_node_compact(child_node))
        
        if compact_parts:
            print(" ".join(compact_parts), file=file_obj)
        return

def extract_signature_info(line: str) -> Tuple[str, str]:
    """Extract function signature and return type from ctags line"""
    signature = ""
    return_type = ""
    
    # Extract the pattern (function definition)
    pattern_match = re.search(r'/\^(.+?)\$/', line)
    if pattern_match:
        pattern = pattern_match.group(1)
        
        # For Go functions, extract parameters from pattern
        if 'func ' in pattern:
            # Extract function signature from Go pattern
            func_match = re.search(r'func\s+(?:\([^)]*\)\s+)?(\w+)\s*\(([^)]*)\)', pattern)
            if func_match:
                params = func_match.group(2).strip()
                if params:
                    signature = f"({params})"
        
        # For Python methods, extract parameters from pattern
        elif 'def ' in pattern:
            # Extract method signature from Python pattern
            def_match = re.search(r'def\s+\w+\s*\(([^)]*)\)', pattern)
            if def_match:
                params = def_match.group(1).strip()
                if params:
                    signature = f"({params})"
        
        # For TypeScript/JavaScript functions
        elif 'function ' in pattern:
            # Extract function signature from JS/TS pattern
            func_match = re.search(r'function\s+\w+\s*\(([^)]*)\)', pattern)
            if func_match:
                params = func_match.group(1).strip()
                if params:
                    signature = f"({params})"
        
        # For arrow functions and method definitions
        elif '=>' in pattern or '(' in pattern:
            # Try to extract parameters from various JS/TS patterns
            # Pattern like: name(params) => return or name(params): return
            param_match = re.search(r'\w+\s*\(([^)]*)\)', pattern)
            if param_match:
                params = param_match.group(1).strip()
                if params:
                    signature = f"({params})"
            
            # Extract return type from arrow function or type annotation
            arrow_match = re.search(r'=>\s*([^{;]+)', pattern)
            if arrow_match:
                return_type = arrow_match.group(1).strip()
            else:
                # Try type annotation pattern: ): ReturnType
                type_match = re.search(r'\):\s*([^{;=]+)', pattern)
                if type_match:
                    return_type = type_match.group(1).strip()
    
    # Extract return type from typeref
    typeref_match = re.search(r'typeref:typename:(.+)', line)
    if typeref_match:
        return_type = typeref_match.group(1).strip()
        # Clean up Go return types
        if return_type.startswith('(') and return_type.endswith(')'):
            return_type = return_type[1:-1]  # Remove outer parentheses
    
    # Extract return type from signature field
    signature_match = re.search(r'signature:\(([^)]*)\)\s*(.+)', line)
    if signature_match:
        if not signature:
            params = signature_match.group(1).strip()
            if params:
                signature = f"({params})"
        if not return_type:
            return_type = signature_match.group(2).strip()
    
    return signature, return_type

def parse_ctags_file(file_path: str, subpath_filter: str = "") -> TreeNode:
    """Parse ctags file and build tree structure"""
    root = TreeNode("root", "folder")
    nested_definitions = {}  # Track nested definitions by their scope
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            lines = f.readlines()
    except FileNotFoundError:
        print(f"Error: Tags file '{file_path}' not found.")
        return root
    
    for line in lines:
        line = line.strip()
        if not line or line.startswith('!'):
            continue
            
        parts = line.split('\t')
        if len(parts) < 4:
            continue
            
        symbol = parts[0]
        file_path_raw = parts[1]
        pattern = parts[2]
        
        # Find the tag type - it's usually the first single character after the pattern
        tag_type = ""
        for i, part in enumerate(parts[3:], 3):
            if len(part) == 1 and part.isalpha():
                tag_type = part
                break
            elif part.endswith(';"') and len(part) > 2:
                # Handle format like 'pattern$/;"' where type comes after
                if i + 1 < len(parts) and len(parts[i + 1]) == 1 and parts[i + 1].isalpha():
                    tag_type = parts[i + 1]
                    break
        
        # Fallback: check if parts[3] contains the type
        if not tag_type:
            if ';' in parts[3]:
                after_semicolon = parts[3].split(';')[1]
                if len(after_semicolon) >= 2 and after_semicolon[1].isalpha():
                    tag_type = after_semicolon[1]
            else:
                tag_type = parts[3]
        
        # Apply subpath filter
        if subpath_filter and not file_path_raw.startswith(f"./{subpath_filter}"):
            continue
            
        # Clean up file path
        file_path_clean = file_path_raw.replace('./', '')
        
        # Extract signature and return type information
        signature, return_type = extract_signature_info(line)
        
        # Parse additional tag information
        scope = ""
        struct_name = ""
        class_name = ""
        
        # Extract scope information for nested definitions
        for part in parts[4:]:
            if part.startswith('struct:'):
                struct_name = part.split(':', 1)[1]
                scope = struct_name
            elif part.startswith('class:'):
                class_name = part.split(':', 1)[1]
                scope = class_name
        
        # Build file path structure
        path_parts = file_path_clean.split('/')
        current_node = root
        
        # Create directory structure
        for i, part in enumerate(path_parts[:-1]):
            current_node = current_node.get_or_create_child(part, "folder")
        
        # Skip package declarations and imports
        if tag_type == 'p':  # package declaration
            continue
        elif tag_type == 'P':  # package import
            continue
        
        # Create file node
        file_node = current_node.get_or_create_child(path_parts[-1], "file")
        
        # Determine symbol type and emoji
        symbol_type = "unknown"
        emoji = "❓"
        
        # Comprehensive tag type mapping based on Universal Ctags documentation
        # Reference: https://docs.ctags.io/en/latest/man/ctags.1.html
        
        if tag_type == 'f':  # function
            symbol_type = "function"
            emoji = "🔧"
        elif tag_type == 'm':  # method or member
            if scope:
                # Check if this is a struct field vs method by looking at the scope info
                is_struct_field = False
                for part in parts[4:]:
                    if part.startswith('struct:'):
                        is_struct_field = True
                        break
                
                if is_struct_field:
                    symbol_type = "field"
                    emoji = "📋"
                else:
                    symbol_type = "method"
                    emoji = "⚙️"
            else:
                symbol_type = "method"
                emoji = "⚙️"
        elif tag_type == 's':  # struct
            symbol_type = "struct"
            emoji = "🏗️"
        elif tag_type == 'c':  # class
            symbol_type = "class"
            emoji = "📦"
        elif tag_type == 'v':  # variable
            # Check if it's a constant by looking at the pattern
            if '/^const ' in line or '/^let ' in line or '/^var ' in line:
                symbol_type = "constant"
                emoji = "🔒"
            else:
                symbol_type = "variable"
                emoji = "📋"
        elif tag_type == 'C':  # constant
            symbol_type = "constant"
            emoji = "🔒"
        elif tag_type == 'd':  # define (macro)
            symbol_type = "macro"
            emoji = "🔧"
        elif tag_type == 'e':  # enumerator
            symbol_type = "enum_member"
            emoji = "🔸"
        elif tag_type == 'g':  # enumeration name
            symbol_type = "enum"
            emoji = "🔢"
        elif tag_type == 'i':  # interface/id
            if file_path_clean.endswith('.css'):
                symbol_type = "css_id"
                emoji = "🆔"
            else:
                symbol_type = "interface"
                emoji = "🔌"
        elif tag_type == 'j':  # heading (HTML)
            symbol_type = "heading"
            emoji = "🔌"
        elif tag_type == 'h':  # heading or header
            if file_path_clean.endswith(('.h', '.hpp', '.hxx', '.h++', '.hh')):
                symbol_type = "header"
                emoji = "📄"
            else:
                symbol_type = "heading"
                emoji = "🔌"
        elif tag_type == 'n':  # namespace
            symbol_type = "namespace"
            emoji = "📦"
        elif tag_type == 't':  # typedef
            symbol_type = "typedef"
            emoji = "🏷️"
        elif tag_type == 'u':  # union
            symbol_type = "union"
            emoji = "🤝"
        elif tag_type == 'l':  # local variable
            symbol_type = "local_var"
            emoji = "📍"
        elif tag_type == 'z':  # parameter
            symbol_type = "parameter"
            emoji = "📝"
        elif tag_type == 'F':  # file
            symbol_type = "file_marker"
            emoji = "📄"
        elif tag_type == 'M':  # macro
            symbol_type = "macro"
            emoji = "🔧"
        elif tag_type == 'P':  # property
            symbol_type = "property"
            emoji = "🔗"
        elif tag_type == 'o':  # object
            symbol_type = "object"
            emoji = "🧩"
        elif tag_type == 'x':  # external
            symbol_type = "external"
            emoji = "🌐"
        elif tag_type == 'r':  # reference
            symbol_type = "reference"
            emoji = "📎"
        elif tag_type == 'T':  # typedef
            symbol_type = "typedef"
            emoji = "🏷️"
        elif tag_type == 'b':  # block
            symbol_type = "block"
            emoji = "🧱"
        elif tag_type == 'k':  # keyword
            symbol_type = "keyword"
            emoji = "🔑"
        elif tag_type == 'w':  # widget
            symbol_type = "widget"
            emoji = "🎛️"
        elif tag_type == 'A':  # array
            symbol_type = "array"
            emoji = "📊"
        elif tag_type == 'S':  # selector (CSS) or setter
            if file_path_clean.endswith('.css'):
                symbol_type = "css_selector"
                emoji = "🎨"
            else:
                symbol_type = "setter"
                emoji = "📥"
        elif tag_type == 'R':  # rule (CSS)
            symbol_type = "css_rule"
            emoji = "📐"
        elif tag_type == 'L':  # label
            symbol_type = "label"
            emoji = "🏷️"
        elif tag_type == 'D':  # define
            symbol_type = "define"
            emoji = "🔧"
        elif tag_type == 'E':  # event
            symbol_type = "event"
            emoji = "⚡"
        elif tag_type == 'G':  # getter
            symbol_type = "getter"
            emoji = "📤"
        elif tag_type == 'I':  # HTML ID
            symbol_type = "html_id"
            emoji = "🆔"
        elif tag_type == 'a':  # alias
            symbol_type = "alias"
            emoji = "🔗"
        elif tag_type == 'y':  # type (alternative)
            symbol_type = "type"
            emoji = "🏷️"
        elif tag_type == 'Y':  # unknown
            symbol_type = "unknown"
            emoji = "❓"
        elif tag_type == 'q':  # qualified
            symbol_type = "qualified"
            emoji = "🔗"
        elif tag_type == 'X':  # external (alternative)
            symbol_type = "external"
            emoji = "🌐"
        elif tag_type == 'N':  # name
            symbol_type = "name"
            emoji = "🏷️"
        elif tag_type == 'B':  # binding
            symbol_type = "binding"
            emoji = "🔗"
        elif tag_type == 'J':  # JSON key
            symbol_type = "json_key"
            emoji = "🔑"
        elif tag_type == 'K':  # key
            symbol_type = "key"
            emoji = "🔑"
        elif tag_type == 'V':  # value
            symbol_type = "value"
            emoji = "💎"
        elif tag_type == 'W':  # wrapper
            symbol_type = "wrapper"
            emoji = "📦"
        elif tag_type == 'Z':  # zone
            symbol_type = "zone"
            emoji = "🌐"
        elif tag_type.startswith('c') and file_path_clean.endswith('.css'):
            symbol_type = "css_class"
            emoji = "🎨"
        
        # Skip certain tag types that are noise
        skip_types = {'p', 'P'}  # package declarations and imports
        if tag_type in skip_types:
            continue
        
        # Skip common noise patterns
        skip_patterns = [
            'export default',
            'export {',
            'import ',
            'from ',
            'require(',
        ]
        
        # Check if this is a noise pattern
        is_noise = False
        for skip_pattern in skip_patterns:
            if skip_pattern in line:
                is_noise = True
                break
        
        if is_noise:
            continue
        
        # Create symbol node
        symbol_node = TreeNode(symbol, symbol_type)
        symbol_node.signature = signature
        symbol_node.return_type = return_type
        
        # Handle nested definitions (struct members, class methods)
        if scope:
            # Find the parent definition
            parent_key = f"{file_path_clean}:{scope}"
            if parent_key not in nested_definitions:
                nested_definitions[parent_key] = []
            nested_definitions[parent_key].append(symbol_node)
        else:
            # Top-level definition
            file_node.add_child(symbol_node)
    
    # Add nested definitions to their parents
    for parent_key, children in nested_definitions.items():
        file_path_clean, parent_name = parent_key.split(':', 1)
        
        # Find the parent node
        path_parts = file_path_clean.split('/')
        current_node = root
        
        for part in path_parts[:-1]:
            if part in current_node.children:
                current_node = current_node.children[part]
        
        if path_parts[-1] in current_node.children:
            file_node = current_node.children[path_parts[-1]]
            
            # Find the parent struct/class
            parent_simple_name = parent_name.split('.')[-1]  # Get last part after dot
            if parent_simple_name in file_node.children:
                parent_node = file_node.children[parent_simple_name]
                for child in children:
                    parent_node.add_child(child)
    
    return root

def main():
    tags_file = "tags"
    subpath_filter = ""
    
    if len(sys.argv) > 1:
        subpath_filter = sys.argv[1]
    
    if not os.path.exists(tags_file):
        print(f"Error: Tags file '{tags_file}' not found.")
        print("Please run 'ctags -R .' to generate the tags file first.")
        return
    
    print("Parsing ctags file...")
    root = parse_ctags_file(tags_file, subpath_filter)
    
    # Generate tree output
    output_file = "code_tree.tree"
    
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write("# Code Tree Structure\n\n")
        f.write("Generated from ctags analysis of the codebase.\n\n")
        
        if subpath_filter:
            f.write(f"**Filtered for subpath:** `{subpath_filter}`\n\n")
        
        # Add complete symbol legend
        f.write("## Symbol Legend\n\n")
        f.write("| Symbol | Type | Description |\n")
        f.write("|--------|------|-------------|\n")
        f.write("| D | Folder | Directory |\n")
        f.write("| F | File/Header | Source file or header file |\n")
        f.write("| f | Function/Macro | Standalone function or macro |\n")
        f.write("| m | Method | Class/struct method |\n")
        f.write("| v | Field/Variable | Struct field or variable |\n")
        f.write("| C | Constant | Constant value |\n")
        f.write("| c | Class/Namespace | Class, namespace, or wrapper |\n")
        f.write("| s | Struct | Struct definition |\n")
        f.write("| i | Interface/Heading | Interface or heading |\n")
        f.write("| t | Type/Typedef | Type definition or typedef |\n")
        f.write("| e | Enum | Enumeration |\n")
        f.write("| E | Enum Member | Enumeration member |\n")
        f.write("| a | Alias/Property | Type alias, property, or binding |\n")
        f.write("| l | Local Variable | Local variable |\n")
        f.write("| p | Parameter | Function parameter |\n")
        f.write("| o | Object | Object definition |\n")
        f.write("| x | External/Zone | External reference or zone |\n")
        f.write("| r | Reference | Reference |\n")
        f.write("| b | Block | Code block |\n")
        f.write("| k | Keyword/Key | Keyword, key, or JSON key |\n")
        f.write("| w | Widget | Widget |\n")
        f.write("| u | Union | Union type |\n")
        f.write("| A | Array | Array |\n")
        f.write("| . | CSS Class | CSS class selector |\n")
        f.write("| S | CSS Selector | CSS selector |\n")
        f.write("| R | CSS Rule | CSS rule |\n")
        f.write("| # | CSS/HTML ID | CSS or HTML ID selector |\n")
        f.write("| ! | Event | Event |\n")
        f.write("| g | Getter | Getter method |\n")
        f.write("| h | Heading | Heading |\n")
        f.write("| M | Macro | Macro definition |\n")
        f.write("| P | Property | Property |\n")
        f.write("| T | Typedef | Type definition |\n")
        f.write("| L | Label | Label |\n")
        f.write("| H | Header | Header file |\n")
        f.write("| n | Namespace | Namespace |\n")
        f.write("| d | Define | Define/macro |\n")
        f.write("| ? | Unknown | Unknown or unrecognized type |\n")
        f.write("\n**Compact Format:** `symbol:name{children}` where `->` indicates return type\n\n")
        
        # Write the ultra-compact tree structure
        print_tree_compact(root, f)
    
    print(f"Code tree generated: {output_file}")
    
    # Print summary
    total_files = count_nodes_by_type(root, "file")
    total_functions = count_nodes_by_type(root, "function")
    total_methods = count_nodes_by_type(root, "method")
    total_classes = count_nodes_by_type(root, "class")
    total_structs = count_nodes_by_type(root, "struct")
    
    print(f"\nSummary:")
    print(f"  Files: {total_files}")
    print(f"  Functions: {total_functions}")
    print(f"  Methods: {total_methods}")
    print(f"  Classes: {total_classes}")
    print(f"  Structs: {total_structs}")

def count_nodes_by_type(node: TreeNode, target_type: str) -> int:
    """Count nodes of a specific type in the tree"""
    count = 0
    if node.type == target_type:
        count += 1
    
    for child in node.children.values():
        count += count_nodes_by_type(child, target_type)
    
    return count

if __name__ == "__main__":
    main()